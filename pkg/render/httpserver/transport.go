package httpserver

import (
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"

	"github.com/pkg/errors"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

const serverTransportTpl = `// Package {{.PkgName}} ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package {{.PkgName}}
{{$methods := .HTTPMethods}}
import (
	"bytes"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

var (
	emptyBytes = []byte("")
)
 
{{range .Iface.Methods}}
	{{$ct := index $methods .Name}}
	{{$uriPlaceholders := $ct.URIPathPlaceholders}}
	{{$queryPlaceholders := $ct.QueryPlaceholders}}
	{{$isIntQueryPlaceholders := $ct.IsIntQueryPlaceholders}}
	{{$headerPlaceholders := $ct.HeaderPlaceholders}}
	{{$body := $ct.Body}}
	{{$bodyPlaceholders := $ct.BodyPlaceholders}}
	{{$isIntBodyPlaceholders := $ct.IsIntBodyPlaceholders}}
	{{$contentType := $ct.ContentType}}
	{{$jsonTags := $ct.JsonTags}}
	{{$multipartValueTags := $ct.MultipartValueTags}}
	{{$multipartFileTags := $ct.MultipartFileTags}}
	{{$responseHeaderPlaceholders := $ct.ResponseHeaders}}
	{{$responseStatus := $ct.ResponseStatus}}
	{{$responseContentType := $ct.ResponseContentType}}
	{{$responseContentEncoding := $ct.ResponseContentEncoding}}
	{{$responseJsonTags := $ct.ResponseJsonTags}}
	{{$responseFile := $ct.ResponseFile}}
	{{$responseFileName := $ct.ResponseFileName}}
	{{$responseBody := $ct.ResponseBody}}
	{{$responseBodyField := $ct.ResponseBodyField}}
	{{$responseBodyType := index $responseBody $ct.ResponseBodyField}}
	{{$responseBodyTypeIsSlice := isSliceType $responseBodyType}}
	{{$responseBodyTypeIsMap := isMapType $responseBodyType}}

	{{if eq $contentType "application/json"}}
		{{if lenMap $body}}
			type {{low .Name}}Request struct {
				{{range $name, $tp := $body}}
					{{up $name}} {{$tp}}{{$tag := index $jsonTags $name}}{{if $tag}} ` + "`" + `json:"{{$tag}}"` + "`" + `{{end}}
				{{end}}
			}
		{{end}}
	{{end}}

	{{if eq $responseContentType "application/json"}}
		{{if lenMap $responseBody}}
			//easyjson:json
			type {{low .Name}}Response {{if or $responseBodyTypeIsSlice $responseBodyTypeIsMap}}{{$responseBodyType}}{{else}} struct {
  				{{if $responseBodyType}}
    				{{$responseBodyType}}
	  			{{else}}
    				{{range $name, $tp := $responseBody}}
						{{up $name}} {{$tp}}{{$tag := index $responseJsonTags $name}}{{if $tag}} ` + "`" + `json:"{{$tag}}"` + "`" + `{{end}}
					{{end}}
	  			{{end}}
			}
			{{end}}
		{{end}}
	{{end}}

	// {{.Name}}Transport transport interface
	type {{.Name}}Transport interface {
		DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) ({{$args := popFirst .Args}}{{joinFullVariables $args "," "err error"}})
		EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, {{$args := popLast .Results}}{{joinFullVariables $args ","}}) (err error)
	}

	type {{low .Name}}Transport struct {
		{{if eq $contentType "application/json"}}{{if lenMap $body}}decodeJSONErrorCreator errorCreator{{end}}{{end}}
		{{if eq $responseContentType "application/json"}}{{if lenMap $responseBody}}encodeJSONErrorCreator errorCreator{{end}}{{end}}
		{{if or $isIntQueryPlaceholders $isIntBodyPlaceholders}}decodeTypeIntErrorCreator errorCreator{{end}}
	}

	// DecodeRequest method for decoding requests on server side
	func (t *{{low .Name}}Transport) DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) ({{$args := popFirst .Args}}{{joinFullVariables $args "," "err error"}}) {
		{{range $i, $sc := $uriPlaceholders}}
			{{$sc}} = ctx.UserValue("{{$sc}}").(string)
		{{end}}
		{{range $from, $to := $queryPlaceholders}}
			{{if $to.IsString}}
				{{if eq $to.IsPointer true}}
					{{$to.Name}} = ptr(ctx.QueryArgs().Peek("{{$from}}"))
				{{else}}
					{{$to.Name}} = string(ctx.QueryArgs().Peek("{{$from}}"))
				{{end}}
			{{else if $to.IsInt}}
				{{if $to.IsPointer}}
					_{{$to.Name}} := ctx.QueryArgs().Peek("{{$from}}")
					if !bytes.Equal(_{{$to.Name}}, emptyBytes) {
						var i int
						i, err = strconv.Atoi(string(_{{$to.Name}}))
						if err != nil {
							err = t.decodeTypeIntErrorCreator(err)
							return
						}
						{{if eq $to.Type "int"}}
							{{$to.Name}} = &i
						{{else}}
							ii := {{$to.Type}}(i) 
							{{$to.Name}} = &ii
						{{end}}
					}
				{{else}}
					_{{$to.Name}} := ctx.QueryArgs().Peek("{{$from}}")
					if !bytes.Equal(_{{$to.Name}}, emptyBytes) {
						var i int	
						i, err = strconv.Atoi(string(_{{$to.Name}}))
						if err != nil {
							err = t.decodeTypeIntErrorCreator(err)
							return
						}
						{{if eq $to.Type "int"}}
							{{$to.Name}} = i
						{{else}} 
							{{$to.Name}} = {{$to.Type}}(i)
						{{end}}
					}
				{{end}}
			{{end}}
		{{end}}
		{{range $from, $to := $headerPlaceholders}}
			{{$to}} = ptr(r.Header.Peek("{{$from}}"))
		{{end}}
		{{if eq $contentType "application/json"}}
			{{if lenMap $body}}var request {{low .Name}}Request
				if err = request.UnmarshalJSON(r.Body()); err != nil {
					err = t.decodeJSONErrorCreator(err)
					return
				}
				{{range $name, $tp := $body}}
					{{$name}} = request.{{up $name}}
				{{end}}
			{{end}}
		{{else if eq $contentType "multipart/form-data"}}
			var form *multipart.Form
			if form, err = ctx.MultipartForm(); err != nil {
				err = errors.Wrap(err, "failed to read MultipartForm")
				return
			}
			{{range $t, $from := $multipartValueTags}}
				{{$to := index $bodyPlaceholders $from}}
				_{{$to.Name}} := form.Value["{{$from}}"]
				{{if $to.IsPointer}}
					if len(_{{$to.Name}}) == 1 {
						{{if $to.IsString}}
							_{{up $to.Name}} := _{{$to.Name}}[0]
							{{$to.Name}} = &_{{up $to.Name}}
						{{else if $to.IsInt}}
							_{{up $to.Name}} := _{{$to.Name}}[0]
							if _{{up $to.Name}} != "" {
								var i int
								i, err = strconv.Atoi(_{{up $to.Name}})
								if err != nil {
									err = t.decodeTypeIntErrorCreator(err)
									return
								}
								{{if eq $to.Type "int"}}
									{{$to.Name}} = &i
								{{else}}
									ii := {{$to.Type}}(i) 
									{{$to.Name}} = &ii
								{{end}}
							}
						{{end}}
					}
				{{else}}
					if len(_{{$to.Name}}) != 1 {
						err = errors.New("failed to read {{$from}} in MultipartForm")
						return
					}
					{{if $to.IsString}}
						{{$to.Name}} = _{{$to.Name}}[0]
					{{else if $to.IsInt}}
						_{{up $to.Name}} := _{{$to.Name}}[0]
						if _{{up $to.Name}} != "" {
							var i int
							i, err = strconv.Atoi(_{{up $to.Name}})
							if err != nil {
								err = t.decodeTypeIntErrorCreator(err)
								return
							}
							{{if eq $to.Type "int"}}
								{{$to.Name}} = i
							{{else}} 
								{{$to.Name}} = {{$to.Type}}(i)
							{{end}}
						}
					{{end}}
				{{end}}
			{{end}}
			{{range $to, $from := $multipartFileTags}}
				_{{$to}} := form.File["{{$from}}"]
				if _{{$to}} == nil {
					err = errors.New("failed to read {{$from}} in MultipartForm")
					return
				}
				if len(_{{$to}}) != 1 {
					err = errors.New("failed to read file in MultipartForm: too many files in {{$from}}")
					return
				}
				{{$to}}, err = _{{$to}}[0].Open()
				if err != nil {
					err = errors.Wrap(err, "failed to open file {{$from}} in MultipartForm")
					return
				}
			{{end}}
		{{end}}
		return
	}

	// EncodeResponse method for encoding response on server side
	func (t *{{low .Name}}Transport) EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, {{$args := popLast .Results}}{{joinFullVariables $args ","}}) (err error) {
		{{if eq $responseContentType "application/json"}}
			r.Header.Set("Content-Type", "application/json{{if $responseContentEncoding}}; charset={{$responseContentEncoding}}{{end}}")
			{{if lenMap $responseBody}}var theResponse {{low .Name}}Response
				{{if $responseBodyType }}
					theResponse{{if or $responseBodyTypeIsSlice $responseBodyTypeIsMap}}{{else}}.{{stripType $responseBodyType}}{{end}} = {{$responseBodyField}}
				{{else}}
					{{range $name, $tp := $responseBody}}
						theResponse.{{up $name}} = {{$name}}
					{{end}}
				{{end}}
				body, err := theResponse.MarshalJSON()
				if err != nil {
					err = t.encodeJSONErrorCreator(err)
					return
				}
				r.SetBody(body)
			{{end}}
		{{end}}
		{{range $to, $from := $responseHeaderPlaceholders}}
			r.Header.Set("{{$to}}", "{{$from}}")
		{{end}}
		{{if $responseFile}}r.SetBody({{$responseFile}}){{end}}
		{{if $responseFileName}}if len(fileName) > 0 {
		r.Header.Set("Content-Disposition", "attachment; filename=\""+fileName+"\"")
			}{{end}}
		r.Header.SetStatusCode({{$responseStatus}})
		return
	}

	// New{{.Name}}Transport the transport creator for http requests
	func New{{.Name}}Transport(
		{{if eq $contentType "application/json"}}
			{{if lenMap $body}}decodeJSONErrorCreator errorCreator,{{end}}
		{{end}}
		{{if eq $responseContentType "application/json"}}
			{{if lenMap $responseBody}}encodeJSONErrorCreator errorCreator,{{end}}
		{{end}}
		{{if or $isIntQueryPlaceholders $isIntBodyPlaceholders}}decodeTypeIntErrorCreator errorCreator,{{end}}
	) {{.Name}}Transport {
		return &{{low .Name}}Transport{
			{{if eq $contentType "application/json"}}
				{{if lenMap $body}}decodeJSONErrorCreator: decodeJSONErrorCreator,{{end}}
			{{end}}
			{{if eq $responseContentType "application/json"}}
				{{if lenMap $responseBody}}encodeJSONErrorCreator: encodeJSONErrorCreator,{{end}}
			{{end}}
			{{if or $isIntQueryPlaceholders $isIntBodyPlaceholders}}decodeTypeIntErrorCreator: decodeTypeIntErrorCreator,{{end}}
		}
	}

{{end}}

func ptr(in []byte) *string {
	if bytes.Equal(in, emptyBytes) {
		return nil
	}
	i := string(in)
	return &i
}
`

// Transport ...
type Transport struct {
	*template.Template
	packageName string
	filePath    []string
	imports     imports
}

// Generate ...
func (s *Transport) Generate(info api.Interface) (err error) {
	info.PkgName = s.packageName
	info.AbsOutputPath = strings.Join(append(strings.Split(info.AbsOutputPath, "/"), s.filePath...), "/")
	dir, _ := path.Split(info.AbsOutputPath)
	err = os.MkdirAll(dir, 0750)
	if err != nil {
		err = errors.Wrap(err, "[httpserver.Transport]os.MkdirAll error")
		return
	}
	serverFile, err := os.Create(info.AbsOutputPath)
	defer func() {
		_ = serverFile.Close()
	}()
	t := template.Must(s.Parse(serverTransportTpl))
	if err = t.Execute(serverFile, info); err != nil {
		err = errors.Wrap(err, "[httpserver.Transport]t.Execute error")
		return
	}
	err = s.imports.GoImports(info.AbsOutputPath)
	// easyJSON generator
	cmd := exec.Command("/bin/sh", "-c", "easyjson -all "+info.AbsOutputPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		err = errors.Wrapf(err, "[httpserver.Transport]cmd.Output error\nCMD: %s\noutput: %s\n", cmd.String(), string(output))
	}

	return
}

// NewTransport ...
func NewTransport(template *template.Template, packageName string, filePath []string, imports imports) *Transport {
	return &Transport{
		Template:    template,
		packageName: packageName,
		filePath:    filePath,
		imports:     imports,
	}
}
