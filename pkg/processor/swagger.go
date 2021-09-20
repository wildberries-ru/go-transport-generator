package processor

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/pkg/errors"
	"github.com/vetcher/go-astra"
	"github.com/vetcher/go-astra/types"

	v1 "github.com/wildberries-ru/go-transport-generator/example/api/v1"
	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

const (
	inPath   = "path"
	inHeader = "header"
	inQuery  = "query"
	inCookie = "cookie"

	errStructNotFound = "struct not found %s %s"
)

type swaggerTagsParser interface {
	Parse(info *api.SwaggerInfo, firstTag string, tags ...string) (err error)
}

type mod interface {
	PkgModPath(pkgName string) (mod string)
}

type swagger struct {
	tagMark                        string
	httpMethodProcessor            HTTPMethod
	swaggerTagsParser              swaggerTagsParser
	mod                            mod
	goGeneratedAutomaticallyPrefix []byte
}

// Process ...
func (s *swagger) Process(info *api.GenerationInfo, iface *api.Interface) (err error) {
	var (
		ok          bool
		swaggerPath *v1.Path
		prop        v1.Schema
	)
	if info.Swagger == nil {
		info.Swagger = &v1.Swagger{
			Info: v1.Info{
				Title:          info.SwaggerInfo.Title,
				Description:    info.SwaggerInfo.Description,
				TermsOfService: nil,
				Contact:        nil,
				License:        nil,
				Version:        info.SwaggerInfo.Version,
			},
			Servers: info.SwaggerInfo.Servers,
		}
	}
	for _, method := range iface.Iface.Methods {
		httpMethod := api.HTTPMethod{}
		err = s.httpMethodProcessor.Process(&httpMethod, iface, method)
		if err != nil {
			err = errors.Wrap(err, "[swagger]s.httpMethodProcessor.Process error")
			return
		}
		for _, doc := range method.Docs {
			doc = strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(doc), "//"))
			if strings.HasPrefix(doc, s.tagMark) {
				words := strings.Split(strings.TrimSpace(doc[len(s.tagMark):]), " ")
				err = s.swaggerTagsParser.Parse(&httpMethod.SwaggerInfo, words[0], words[1:]...)
				if err != nil {
					err = errors.Wrap(err, "[swagger]s.swaggerTagsParser.Parse error")
					return
				}
			}
		}
		if info.Swagger.Paths == nil {
			info.Swagger.Paths = make(map[string]*v1.Path)
		}
		if swaggerPath, ok = info.Swagger.Paths[httpMethod.RawURIPath]; !ok {
			swaggerPath = &v1.Path{}
			info.Swagger.Paths[httpMethod.RawURIPath] = swaggerPath
		}
		httpMethodQueryPlaceholders := make(map[string]string, len(httpMethod.QueryPlaceholders))
		for k, v := range httpMethod.QueryPlaceholders {
			httpMethodQueryPlaceholders[k] = v.Name
		}
		params := s.fillParamsFromSlice(iface.RelOutputPath, method.Args, httpMethod.URIPathPlaceholders, inPath, true)
		params = append(params, s.fillParamsFromMap(iface.RelOutputPath, method.Args, httpMethod.HeaderPlaceholders, inHeader, false)...)
		params = append(params, s.fillParamsFromMap(iface.RelOutputPath, method.Args, httpMethodQueryPlaceholders, inQuery, false)...)
		params = append(params, s.fillParamsFromMap(iface.RelOutputPath, method.Args, httpMethod.CookiePlaceholders, inCookie, false)...)

		var req *v1.RequestBody
		if len(httpMethod.Body) > 0 {
			reqSchema := v1.Schema{}
			reqSchema.Type = "object"
			reqSchema.Nullable = false
			reqSchema.Properties = make(map[string]v1.Schema)
			for to := range httpMethod.Body {
				for _, arg := range method.Args {
					if arg.Name == to {
						prop, err = s.makeType(iface.RelOutputPath, arg, arg.Type) //todo pkgPath
						if err != nil {
							err = errors.Wrap(err, "[swagger.RequestBody]s.makeType error")
							return
						}
						reqSchema.Properties[arg.Name] = prop
					}
				}
			}
			req = &v1.RequestBody{
				Description: "",
				Content: v1.Content{
					httpMethod.ContentType: v1.Media{
						Schema: reqSchema,
					},
				},
			}
		}

		resSchema := v1.Schema{}
		resSchema.Type = "object"
		resSchema.Nullable = false
		resSchema.Properties = make(map[string]v1.Schema)
		for from := range httpMethod.ResponseBody {
			for _, arg := range method.Results {
				if arg.Name == from {
					prop, err = s.makeType(iface.RelOutputPath, arg, arg.Type)
					if err != nil {
						err = errors.Wrap(err, "[swagger.resSchema]s.makeType error")
						return
					}
					resSchema.Properties[httpMethod.ResponseJSONTags[arg.Name]] = prop
				}
			}
		}
		responses := make(map[string]v1.Response)

		response := v1.Response{
			Description: "",
			Content: v1.Content{
				httpMethod.ResponseContentType: v1.Media{
					Schema: resSchema,
				},
			},
		}

		if len(httpMethod.ResponseHeaders) > 0 {
			response.Headers = make(map[string]v1.Header, len(httpMethod.ResponseHeaders))
			for to, from := range httpMethod.ResponseHeaders {
				for _, r := range method.Results {
					if r.Name == from {
						prop, err = s.makeType(iface.RelOutputPath, r, r.Type)
						if err != nil {
							err = errors.Wrap(err, "[swagger.ResponseHeaders]s.makeType error")
							return
						}
						response.Headers[to] = v1.Header{
							Description: "",
							Schema:      prop,
						}
					}
				}
			}
		}

		responses[s.castStatusConst(httpMethod.ResponseStatus)] = response

		op := v1.Operation{
			Tags:        []string{path.Base(iface.AbsOutputPath) + "/" + iface.Iface.Name},
			Summary:     httpMethod.SwaggerInfo.Summary,
			Description: httpMethod.SwaggerInfo.Description,
			Parameters:  params,
			RequestBody: req,
			Responses:   responses,
		}
		reflect.ValueOf(swaggerPath).Elem().FieldByName(strings.Title(strings.ToLower(httpMethod.Method))).Set(reflect.ValueOf(&op))
	}
	return
}

func (s *swagger) fillParamsFromSlice(pkg string, args []types.Variable, placeholders []string, in string, required bool) (params []v1.Parameter) {
	for _, p := range placeholders {
		for _, arg := range args {
			if arg.Name == p {
				schema, err := s.makeType(pkg, arg, arg.Type)
				if err != nil {
					return
				}
				params = append(params, v1.Parameter{
					In:       in,
					Name:     p,
					Required: required,
					Schema:   schema,
				})
				break
			}
		}
	}
	return
}

func (s *swagger) fillParamsFromMap(pkg string, args []types.Variable, placeholders map[string]string, in string, required bool) (params []v1.Parameter) {
	for _, to := range placeholders {
		for _, arg := range args {
			if arg.Name == to {
				schema, err := s.makeType(pkg, arg, arg.Type)
				if err != nil {
					return
				}
				params = append(params, v1.Parameter{
					In:       in,
					Name:     to,
					Required: required,
					Schema:   schema,
				})
				break
			}
		}
	}
	return
}

func (s *swagger) makeType(pkgPath string, field types.Variable, fieldType types.Type) (schema v1.Schema, err error) {
	var (
		structInfo types.Struct
		items      v1.Schema
	)
	for fieldType == nil {
		return
	}
	switch f := fieldType.(type) {
	case types.TName:
		if s.isBuiltin(fieldType) {
			schema.Type, schema.Format = s.castBuiltinType(fieldType)
			return
		}
		structInfo, err = s.searchStructInfo(pkgPath, f.TypeName)
		if err != nil {
			err = errors.Wrap(err, "[swagger.TName]s.searchStructInfo error")
			return
		}
		schema.Type = "object"
		schema.Nullable = false
		schema.Properties, err = s.fillProps(structInfo, pkgPath)
		if err != nil {
			err = errors.Wrap(err, "[swagger.TName]s.fillProps error")
			return
		}
	case types.Struct:
		if s.isBuiltin(fieldType) {
			schema.Type, schema.Format = s.castBuiltinType(fieldType)
			return
		}
		structInfo, err = s.searchStructInfo(pkgPath, f.Name)
		if err != nil {
			err = errors.Wrap(err, "[swagger.Struct]s.searchStructInfo error")
			return
		}
		schema.Type = "object"
		schema.Nullable = false
		schema.Properties, err = s.fillProps(structInfo, pkgPath)
		if err != nil {
			err = errors.Wrap(err, "[swagger.Struct]s.fillProps error")
			return
		}
	case types.TImport:
		if s.isBuiltin(fieldType) {
			schema.Type, schema.Format = s.castBuiltinType(fieldType)
			return
		}
		structInfo, err = s.searchStructInfo(f.Import.Package, f.Next.String())
		if err != nil {
			err = errors.Wrap(err, "[swagger.TImport]s.searchStructInfo error")
			return
		}
		schema.Type = "object"
		schema.Properties, err = s.fillProps(structInfo, f.Import.Package)
		if err != nil {
			schema.Properties, err = s.fillProps(structInfo, pkgPath)
			if err != nil {
				err = errors.Wrap(err, "[swagger.TImport]s.fillProps error")
				return
			}
		}
	case types.TArray:
		schema.Type = "array"
		items, err = s.makeType(pkgPath, field, f.Next)
		if err != nil {
			err = errors.Wrap(err, "[swagger.TArray]s.makeType error")
			return
		}
		schema.Items = &items
	case types.TEllipsis:
		// todo
		schema, err = s.makeType(pkgPath, field, f.Next)
		if err != nil {
			err = errors.Wrap(err, "[swagger.TEllipsis]s.makeType error")
			return
		}
	case types.TMap:
		// todo
		schema.Type = "object"
		//key, err := s.makeType("", field, f.Key)
		//if err != nil {
		//	return
		//}
		//val, err := s.makeType("", field, f.Value)
		//if err != nil {
		//	return
		//}
	case types.TPointer:
		schema, err = s.makeType(pkgPath, field, f.Next)
		if err != nil {
			err = errors.Wrap(err, "[swagger.TPointer]s.makeType error")
			return
		}
		schema.Nullable = true
	case types.TInterface:
	default:
		err = errors.New("unknown type " + fieldType.String())
	}
	return
}

func (s *swagger) fillProps(structInfo types.Struct, pkgPath string) (properties map[string]v1.Schema, err error) {
	props := make(map[string]v1.Schema)
	defer func() {
		if err != nil {
			return
		}
		if len(props) > 0 {
			properties = props
		}
	}()
	for _, field := range structInfo.Fields {
		if jsonTags, found := field.Tags["json"]; found {
			fieldName := jsonTags[0]
			if fieldName != "-" {
				props[fieldName], err = s.makeType(pkgPath, field.Variable, field.Type)
				if err != nil {
					return
				}
			}
			continue
		}
		fieldName := field.Name
		if len(fieldName) > 0 {
			if !s.fieldIsPrivate(fieldName) {
				props[fieldName], err = s.makeType(pkgPath, field.Variable, field.Type)
				if err != nil {
					return
				}
			}
		}
	}
	return
}

func (s *swagger) castBuiltinType(tp types.Type) (typeName, format string) {
	switch tp.String() {
	case "bool":
		typeName = "boolean"
	case "time.Time":
		format = "date-time"
		typeName = "string"
	case "byte":
		format = "byte"
		typeName = "string"
	case "uuid.UUID":
		format = "uuid"
		typeName = "string"
	case "float32", "float64":
		format = "float"
		typeName = "number"
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		typeName = "number"
	case "multipart.FileHeader":
		format = "binary"
		typeName = "string"
	default:
		typeName = tp.String()
	}
	return
}

func (s *swagger) castStatusConst(v string) string {
	switch v {
	default:
		return v
	case "http.StatusContinue":
		return "100"
	case "http.StatusSwitchingProtocols":
		return "101"
	case "http.StatusProcessing":
		return "102"
	case "http.StatusEarlyHints":
		return "103"

	case "http.StatusOK":
		return "200"
	case "http.StatusCreated":
		return "201"
	case "http.StatusAccepted":
		return "202"
	case "http.StatusNonAuthoritativeInfo":
		return "203"
	case "http.StatusNoContent":
		return "204"
	case "http.StatusResetContent":
		return "205"
	case "http.StatusPartialContent":
		return "206"
	case "http.StatusMultiStatus":
		return "207"
	case "http.StatusAlreadyReported":
		return "208"
	case "http.StatusIMUsed":
		return "226"

	case "http.StatusMultipleChoices":
		return "300"
	case "http.StatusMovedPermanently":
		return "301"
	case "http.StatusFound":
		return "302"
	case "http.StatusSeeOther":
		return "303"
	case "http.StatusNotModified":
		return "304"
	case "http.StatusUseProxy":
		return "305"

	case "http.StatusTemporaryRedirect":
		return "307"
	case "http.StatusPermanentRedirect":
		return "308"

	case "http.StatusBadRequest":
		return "400"
	case "http.StatusUnauthorized":
		return "401"
	case "http.StatusPaymentRequired":
		return "402"
	case "http.StatusForbidden":
		return "403"
	case "http.StatusNotFound":
		return "404"
	case "http.StatusMethodNotAllowed":
		return "405"
	case "http.StatusNotAcceptable":
		return "406"
	case "http.StatusProxyAuthRequired":
		return "407"
	case "http.StatusRequestTimeout":
		return "408"
	case "http.StatusConflict":
		return "409"
	case "http.StatusGone":
		return "410"
	case "http.StatusLengthRequired":
		return "411"
	case "http.StatusPreconditionFailed":
		return "412"
	case "http.StatusRequestEntityTooLarge":
		return "413"
	case "http.StatusRequestURITooLong":
		return "414"
	case "http.StatusUnsupportedMediaType":
		return "415"
	case "http.StatusRequestedRangeNotSatisfiable":
		return "416"
	case "http.StatusExpectationFailed":
		return "417"
	case "http.StatusTeapot":
		return "418"
	case "http.StatusMisdirectedRequest":
		return "421"
	case "http.StatusUnprocessableEntity":
		return "422"
	case "http.StatusLocked":
		return "423"
	case "http.StatusFailedDependency":
		return "424"
	case "http.StatusTooEarly":
		return "425"
	case "http.StatusUpgradeRequired":
		return "426"
	case "http.StatusPreconditionRequired":
		return "428"
	case "http.StatusTooManyRequests":
		return "429"
	case "http.StatusRequestHeaderFieldsTooLarge":
		return "431"
	case "http.StatusUnavailableForLegalReasons":
		return "451"

	case "http.StatusInternalServerError":
		return "500"
	case "http.StatusNotImplemented":
		return "501"
	case "http.StatusBadGateway":
		return "502"
	case "http.StatusServiceUnavailable":
		return "503"
	case "http.StatusGatewayTimeout":
		return "504"
	case "http.StatusHTTPVersionNotSupported":
		return "505"
	case "http.StatusVariantAlsoNegotiates":
		return "506"
	case "http.StatusInsufficientStorage":
		return "507"
	case "http.StatusLoopDetected":
		return "508"
	case "http.StatusNotExtended":
		return "510"
	case "http.StatusNetworkAuthenticationRequired":
		return "511"
	}
}

func (s *swagger) searchStructInfo(pkg, name string) (structInfo types.Struct, err error) {
	var (
		tmpStructInfo *types.Struct
	)
	pkgPath := s.mod.PkgModPath(pkg)
	if tmpStructInfo, err = s.getStructInfo(pkgPath, name); err == nil && tmpStructInfo != nil {
		return *tmpStructInfo, nil
	}
	pkgPath = path.Join("./vendor", pkg)
	if tmpStructInfo, err = s.getStructInfo(pkgPath, name); err == nil && tmpStructInfo != nil {
		return *tmpStructInfo, nil
	}
	pkgPath = s.trimLocalPkg(pkg)
	if tmpStructInfo, err = s.getStructInfo(pkgPath, name); err == nil && tmpStructInfo != nil {
		return *tmpStructInfo, nil
	}
	if err != nil {
		err = errors.Wrapf(err, errStructNotFound, pkg, name)
	} else {
		err = fmt.Errorf(errStructNotFound, pkg, name)
	}
	return
}

func (s *swagger) getStructInfo(relPath, name string) (structInfo *types.Struct, err error) {
	var (
		srcFile  *types.File
		body     []byte
		files    []os.FileInfo
		filePath string
		st       types.Struct
	)
	pkgPath, _ := filepath.Abs(relPath)
	if files, err = ioutil.ReadDir(pkgPath); err != nil {
		return
	}

	for _, file := range files {
		filePath = path.Join(pkgPath, file.Name())
		if file.IsDir() {
			structInfo, err = s.getStructInfo(filePath, name)
			if err != nil || structInfo != nil {
				return
			}
			continue
		}
		if !strings.HasSuffix(file.Name(), ".go") {
			continue
		}
		body, err = ioutil.ReadFile(filePath)
		if err != nil {
			return
		}
		if bytes.HasPrefix(body, s.goGeneratedAutomaticallyPrefix) {
			continue
		}
		if srcFile, err = astra.ParseFile(filePath); err != nil {
			err = errors.Wrap(err, fmt.Sprintf("%s,%s", filePath, name))
			return
		}
		for _, st = range srcFile.Structures {
			if st.Name == name {
				structInfo = &st
				return
			}
		}
	}

	return
}

func (s *swagger) fieldIsPrivate(name string) bool {
	return string([]rune(name)[0]) != strings.ToUpper(string([]rune(name)[0]))
}

func (s *swagger) getModName() (module string) {
	modFile, err := os.OpenFile("go.mod", os.O_RDONLY, os.ModePerm)
	if err != nil {
		return
	}
	defer func() {
		_ = modFile.Close()
	}()
	rd := bufio.NewReader(modFile)
	if module, err = rd.ReadString('\n'); err != nil {
		return ""
	}
	module = strings.Trim(module, "\n")
	moduleTokens := strings.Split(module, " ")
	if len(moduleTokens) == 2 {
		module = strings.TrimSpace(moduleTokens[1])
	}
	return
}

func (s *swagger) trimLocalPkg(pkg string) (pgkPath string) {
	module := s.getModName()
	if module == "" {
		return pkg
	}
	moduleTokens := strings.Split(module, "/")
	pkgTokens := strings.Split(pkg, "/")
	if len(pkgTokens) < len(moduleTokens) {
		return pkg
	}
	pgkPath = path.Join(strings.Join(pkgTokens[len(moduleTokens):], "/"))
	return
}

func (s *swagger) isBuiltin(t types.Type) bool {
	if types.IsBuiltin(t) {
		return true
	}
	switch strings.TrimPrefix(t.String(), "*") {
	case "uuid.UUID", "UUID", "json.RawMessage", "bson.ObjectId", "time.Time", "multipart.FileHeader":
		return true
	default:
		return false
	}
}

// NewSwagger ...
func NewSwagger(
	tagMark string,
	httpMethodProcessor HTTPMethod,
	swaggerTagsParser swaggerTagsParser,
	mod mod,
	goGeneratedAutomaticallyPrefix []byte,
) Processor {
	return &swagger{
		tagMark:                        tagMark,
		httpMethodProcessor:            httpMethodProcessor,
		swaggerTagsParser:              swaggerTagsParser,
		mod:                            mod,
		goGeneratedAutomaticallyPrefix: goGeneratedAutomaticallyPrefix,
	}
}
