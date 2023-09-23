package processor

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/vetcher/go-astra/types"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

const (
	errHTTPMethodGETCouldNotHaveRequestBody = "http method GET could not have request body in %s interface %s method %s"
	errTooManyReturnOctetParams             = "http method with application/octet-stream content type expects 1 bosy or return parameter, but got %v in %s interface %s method %s"
	errTypeToMetrics                        = "only strings, ints and pointers on them are allowed as addition metrics labels: method: %s %s variable: %s"
	errParameterDoNotExist                  = "parameter %s does not exist in method %s"
)

// HTTPMethod ...
type HTTPMethod interface {
	Process(httpInfoMethod *api.HTTPMethod, iface *api.Interface, method *types.Function) (err error)
}

type httpMethod struct {
	tagMark    string
	tagsParser httpServerTagsParser
}

// Process ...
func (s *httpMethod) Process(httpMethod *api.HTTPMethod, iface *api.Interface, method *types.Function) (err error) {
	var (
		from, to      string
		toPlaceholder *api.Placeholder
		arg           types.Variable
		args          []types.Variable
		res           types.Variable
		results       []types.Variable
		diff          map[string]string
	)
	for _, doc := range method.Docs {
		doc = strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(doc), "//"))
		if strings.HasPrefix(doc, s.tagMark) {
			words := strings.Split(strings.TrimSpace(doc[len(s.tagMark):]), " ")
			err = s.tagsParser.Parse(httpMethod, words[0], words[1:]...)
			if err != nil {
				return
			}
		}
	}
	args = method.Args[1:]
	diff = make(map[string]string, len(args))
	for _, arg = range args {
		diff[arg.Name] = arg.Type.String()
	}

	for from, toMetricsPlaceholder := range httpMethod.AdditionalMetricsLabels {
		for _, arg = range args {
			if _, ok := diff[from]; !ok {
				return fmt.Errorf(errParameterDoNotExist, from, method.Name)
			}
			if arg.Name == toMetricsPlaceholder.Name {
				err := s.castMetricsLabelPlaceholder(from, arg.Type, httpMethod)
				if err != nil {
					return err
				}
			}
		}
	}

	for from, toPlaceholder = range httpMethod.QueryPlaceholders {
		delete(diff, toPlaceholder.Name)
		for _, arg = range args {
			if arg.Name == toPlaceholder.Name {
				s.castQueryPlaceholder(from, arg, arg.Type, httpMethod)
			}
		}
	}

	for _, to = range httpMethod.HeaderPlaceholders {
		delete(diff, to)
	}
	for _, to = range httpMethod.CookiePlaceholders {
		delete(diff, to)
	}
	for _, to = range httpMethod.URIPathPlaceholders {
		delete(diff, to)
	}
	if httpMethod.Method == http.MethodGet && len(diff) > 0 {
		return fmt.Errorf(errHTTPMethodGETCouldNotHaveRequestBody, iface.RelOutputPath, iface.Iface.Name, method.Name)
	}
	httpMethod.Body = diff
	if len(diff) > 0 {
		httpMethod.BodyPlaceholders = make(map[string]*api.Placeholder, len(diff))
		for _, arg = range args {
			if _, ok := diff[arg.Name]; ok {
				httpMethod.BodyPlaceholders[arg.Name] = &api.Placeholder{
					Name: arg.Name,
				}
				s.castBodyPlaceholder(arg.Name, arg, arg.Type, httpMethod)
			}
		}
	}

	results = method.Results[:len(method.Results)-1]

	if httpMethod.ContentType == "application/octet-stream" && len(httpMethod.BodyPlaceholders) != 1 {
		return fmt.Errorf(errTooManyReturnOctetParams, len(results), iface.RelOutputPath, iface.Iface.Name, method.Name)
	}

	if httpMethod.ResponseContentType == "application/octet-stream" && len(results) != 1 {
		return fmt.Errorf(errTooManyReturnOctetParams, len(results), iface.RelOutputPath, iface.Iface.Name, method.Name)
	}

	diff = make(map[string]string, len(results))
	for _, res = range results {
		diff[res.Name] = res.Type.String()
	}
	for _, from = range httpMethod.ResponseCookies {
		delete(diff, from)
	}
	for _, from = range httpMethod.ResponseHeaders {
		delete(diff, from)
	}
	httpMethod.ResponseBody = diff
	return
}

func (s *httpMethod) castQueryPlaceholder(from string, arg types.Variable, argType types.Type, httpMethod *api.HTTPMethod) {
	switch tp := argType.(type) {
	case types.TName:
		if s.isInt(tp) {
			httpMethod.QueryPlaceholders[from].IsInt = true
			httpMethod.IsIntQueryPlaceholders = true
		} else if s.isString(tp) {
			httpMethod.QueryPlaceholders[from].IsString = true
		}
		httpMethod.QueryPlaceholders[from].Type = tp.String()
		return
	case types.TPointer:
		httpMethod.QueryPlaceholders[from].IsPointer = true
		s.castQueryPlaceholder(from, arg, tp.Next, httpMethod)
		return
	}
}

func (s *httpMethod) castMetricsLabelPlaceholder(from string, argType types.Type, httpMethod *api.HTTPMethod) (err error) {
	switch tp := argType.(type) {
	case types.TName:
		if s.isInt(tp) {
			httpMethod.AdditionalMetricsLabels[from].IsInt = true
		} else if s.isString(tp) {
			httpMethod.AdditionalMetricsLabels[from].IsString = true
		}
	case types.TPointer:
		httpMethod.AdditionalMetricsLabels[from].IsPointer = true
		if s.isInt(tp.Next) {
			httpMethod.AdditionalMetricsLabels[from].IsInt = true
		} else if s.isString(tp.Next) {
			httpMethod.AdditionalMetricsLabels[from].IsString = true
		}
	default:
		return fmt.Errorf(errTypeToMetrics, httpMethod.Method, httpMethod.URIPath, from)
	}

	return
}

func (s *httpMethod) castBodyPlaceholder(to string, arg types.Variable, argType types.Type, httpMethod *api.HTTPMethod) {
	switch tp := argType.(type) {
	case types.TName:
		if s.isInt(tp) {
			httpMethod.BodyPlaceholders[to].IsInt = true
			if httpMethod.ContentType == "multipart/form-data" {
				httpMethod.IsIntBodyPlaceholders = true
			}
		} else if s.isString(tp) {
			httpMethod.BodyPlaceholders[to].IsString = true
		} else if s.isBool(tp) {
			httpMethod.BodyPlaceholders[to].IsBool = true
		}
		httpMethod.BodyPlaceholders[to].Type = tp.String()
		return
	case types.TPointer:
		httpMethod.BodyPlaceholders[to].IsPointer = true
		s.castBodyPlaceholder(to, arg, tp.Next, httpMethod)
		return
	case types.TArray:
		httpMethod.BodyPlaceholders[to].IsSlice = true
		s.castBodyPlaceholder(to, arg, tp.Next, httpMethod)
		return
	}
}

func (s *httpMethod) isInt(tp types.Type) bool {
	switch tp.String() {
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		return true
	}
	return false
}

func (s *httpMethod) isString(tp types.Type) bool {
	switch tp.String() {
	case "string":
		return true
	}
	return false
}

func (s *httpMethod) isBool(tp types.Type) bool {
	switch tp.String() {
	case "bool":
		return true
	}
	return false
}

// NewHTTPMethod ...
func NewHTTPMethod(tagMark string, tagsParser httpServerTagsParser) HTTPMethod {
	return &httpMethod{
		tagMark:    tagMark,
		tagsParser: tagsParser,
	}
}
