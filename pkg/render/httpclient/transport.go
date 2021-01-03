package httpclient

import (
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"

	"github.com/pkg/errors"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

const clientTransportTpl = `// Package {{.PkgName}} ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package {{.PkgName}}
{{$methods := .HTTPMethods}}
{{$isTLS := .IsTLSClient}}
import (
	"context"
	"net/http"
	"strconv"
	"mime/multipart"

	"github.com/valyala/fasthttp"
)
 
{{range .Iface.Methods}}

	{{$ct := index $methods .Name}}
	{{$method := $ct.Method}}
	{{$bodyPlaceholders := $ct.BodyPlaceholders}}
	{{$uriPlaceholders := $ct.URIPathPlaceholders}}
	{{$queryPlaceholders := $ct.QueryPlaceholders}}
	{{$headerPlaceholders := $ct.HeaderPlaceholders}}
	{{$cookiePlaceholders := $ct.CookiePlaceholders}}
	{{$body := $ct.Body}}
	{{$contentType := $ct.ContentType}}
	{{$jsonTags := $ct.JsonTags}}
	{{$multipartValueTags := $ct.MultipartValueTags}}
	{{$multipartFileTags := $ct.MultipartFileTags}}
	{{$formUrlencodedTags := $ct.FormUrlencodedTags}}
	{{$responseJsonTags := $ct.ResponseJsonTags}}
	{{$responseHeaderPlaceholders := $ct.ResponseHeaders}}
	{{$responseStatus := $ct.ResponseStatus}}
	{{$responseContentType := $ct.ResponseContentType}}
	{{$responseFile := $ct.ResponseFile}}
	{{$responseBody := $ct.ResponseBody}}
	{{$responseBodyField := $ct.ResponseBodyField}}
	{{$responseBodyType := index $responseBody $ct.ResponseBodyField}}
	{{$responseBodyTypeIsSlice := isSliceType $responseBodyType}}
	{{$responseBodyTypeIsMap := isMapType $responseBodyType}}

	{{if lenMap $body}}{{if eq $contentType "application/json"}}//easyjson:json{{else}}//easyjson:skip{{end}}
		{{$bodyLen := lenMap $body}}
		{{$n := .Name}}
			type {{low .Name}}Request struct {
				{{range $name, $tp := $body}}{{up $name}} {{$tp}}{{$tag := index $jsonTags $name}}{{if $tag}} ` + "`" + `json:"{{$tag}}"` + "`" + `{{end}}
				{{end}}
			}
	{{end}}

	{{if lenMap $responseBody}}{{if eq $responseContentType "application/json"}}//easyjson:json{{else}}//easyjson:skip{{end}}
		type {{low .Name}}Response {{if or $responseBodyTypeIsSlice $responseBodyTypeIsMap}}{{$responseBodyType}}{{else}} struct {
  			{{if $responseBodyType}}
    			{{$responseBodyType}}
  			{{else}}
			    {{$respLen := lenMap $responseBody}}	
			    {{range $name, $tp := $responseBody}}{{if eq $respLen 1}}{{if contains $tp "."}}{{$tp}}{{else}}{{up $name}} {{$tp}}{{end}}{{else}}{{up $name}} {{$tp}}{{end}}{{$tag := index $responseJsonTags $name}}{{if $tag}} ` + "`" + `json:"{{$tag}}"` + "`" + `{{end}}
    			{{end}}
  			{{end}}
		}{{end}}
	{{end}}

	// {{.Name}}Transport transport interface
	type {{.Name}}Transport interface {
		EncodeRequest(ctx context.Context, r *fasthttp.Request, {{$args := popFirst .Args}}{{joinFullVariables $args ","}}) (err error)
		DecodeResponse(ctx context.Context, r *fasthttp.Response) ({{$args := popLast .Results}}{{joinFullVariables $args "," "err error"}})
	}

	{{if ne $contentType "application/json"}}//easyjson:skip{{end}}
	type {{low .Name}}Transport struct {
		errorProcessor errorProcessor
		pathTemplate   string
		method         string
	}

	// EncodeRequest method for decoding requests on server side
	func (t *{{low .Name}}Transport) EncodeRequest(ctx context.Context, r *fasthttp.Request, {{$args := popFirst .Args}}{{joinFullVariables $args ","}}) (err error) {
		r.Header.SetMethod(t.method)
		{{if len $uriPlaceholders}}r.SetRequestURI(fmt.Sprintf(t.pathTemplate, {{join $uriPlaceholders ","}})){{else}}r.SetRequestURI(t.pathTemplate){{end}}
		{{if $isTLS}}
			{{if not $queryPlaceholders}}
				_ = r.URI()
			{{end}}
		{{end}}
		{{range $from, $to := $queryPlaceholders}}
			{{if eq $to.IsString true}}
				{{if eq $to.IsPointer true}}if {{$to.Name}} != nil { {{end}}
				r.URI().QueryArgs().Set("{{$from}}", {{if eq $to.IsPointer true}}*{{end}}{{$to.Name}})
				{{if eq $to.IsPointer true}} } {{end}}
			{{else if eq $to.IsInt true}}
				{{if eq $to.IsPointer true}}if {{$to.Name}} != nil { {{end}}
				r.URI().QueryArgs().Set("{{$from}}", strconv.Itoa({{if ne $to.Type "int"}}int({{if eq $to.IsPointer true}}*{{end}}{{$to.Name}}){{else}}{{if eq $to.IsPointer true}}*{{end}}{{$to.Name}}{{end}}))
				{{if eq $to.IsPointer true}} } {{end}}
			{{end}}
		{{end}}
		{{range $from, $to := $headerPlaceholders}}
			r.Header.Set("{{$from}}", *{{$to}})
		{{end}}
		{{range $from, $to := $cookiePlaceholders}}
			r.Header.SetCookie("{{$from}}", *{{$to}})
		{{end}}
		{{if eq $contentType "application/json"}}r.Header.Set("Content-Type", "application/json")
			{{if lenMap $body}}var request {{low .Name}}Request
				{{$bodyLen := lenMap $body}}
					{{range $name, $tp := $body}}
						request.{{up $name}} = {{$name}}
					{{end}}
				body, err := request.MarshalJSON()
				if err != nil {
					return
				}
				r.SetBody(body)
			{{end}}
		{{end}}
		
		{{if eq $contentType "multipart/form-data"}}r.Header.Set("Content-Type", "multipart/form-data")
			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			
			{{range $fr, $to := $multipartValueTags}}
				{{$from := index $bodyPlaceholders $fr}}
				{{if $from.IsPointer}}
					if {{$from.Name}}) != nil {
						{{if $from.IsString}}
							_{{$from.Name}} := *{{$from.Name}}
						{{else if $from.IsInt}}
							_{{$from.Name}} := fmt.Sprintf("%d", *{{$from.Name}})
						{{end}}
					}
				{{else}}
					{{if $from.IsString}}
						_{{$from.Name}} := {{$from.Name}}
					{{else if $from.IsInt}}
						_{{$from.Name}} := fmt.Sprintf("%d", {{$from.Name}})
					{{end}}
				{{end}}
				writer.WriteField("{{$to}}", _{{$from.Name}})
			{{end}}

			{{range $fr, $to := $multipartFileTags}}
				{{$from := index $bodyPlaceholders $fr}}
				{{if isFileHeaderPlaceholder $from.Type}}
					part{{up $from.Name}}, _ := writer.CreateFormFile("{{$to}}", {{$from.Name}}.Filename)
					var _{{$from.Name}} multipart.File
					_{{$from.Name}}, err = {{$from.Name}}.Open()
					if err != nil {
						return
					}
					io.Copy(part{{up $from.Name}}, _{{$from.Name}})
				{{else if isBytesPlaceholder $from.Type}}
					// we don't have any name at this time, receiver should know what he receives
					part{{up $from.Name}}, _ := writer.CreateFormFile("{{$to}}", "{{$from.Name}}") 
					io.Copy(part{{up $from.Name}}, bytes.NewReader({{$from.Name}}))
				{{end}}
			{{end}}

			writer.Close()
			r.Header.Set("Content-Type", writer.FormDataContentType())
			r.SetBody(body.Bytes())
		{{end}}

		{{if eq $contentType "application/x-www-form-urlencoded"}}r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			{{range $from, $to := $formUrlencodedTags}}
				r.PostArgs().Add("{{$to}}", {{$from}})
			{{end}}
		{{end}}
		return
	}

	// DecodeResponse method for decoding response on server side
	func (t *{{low .Name}}Transport) DecodeResponse(ctx context.Context, r *fasthttp.Response) ({{$args := popLast .Results}}{{joinFullVariables $args "," "err error"}}) {
		if r.StatusCode() != {{$responseStatus}} {
			err = t.errorProcessor.Decode(r)
			return
		}
		{{if eq $responseContentType "application/json"}}
			{{if lenMap $responseBody}}var theResponse {{low .Name}}Response
				if err = theResponse.UnmarshalJSON(r.Body()); err != nil {
					return
				}
				{{if $responseBodyType }}
					{{$responseBodyField}} = theResponse{{if or $responseBodyTypeIsSlice $responseBodyTypeIsMap}}{{else}}.{{stripType $responseBodyType}}{{end}}
				{{else}}
					{{range $name, $tp := $responseBody}}
						{{$name}} = theResponse.{{up $name}}
					{{end}}
				{{end}}
			{{end}}
		{{end}}
		{{range $to, $from := $responseHeaderPlaceholders}}
			{{$from}} = ptr(r.Header.Peek("{{$to}}"))
		{{end}}
		{{if $responseFile}}{{$responseFile}} = r.Body(){{end}}
		return
	}

	// New{{.Name}}Transport the transport creator for http requests
	func New{{.Name}}Transport(
		errorProcessor errorProcessor,
		pathTemplate string,
		method string,
	) {{.Name}}Transport {
		return &{{low .Name}}Transport{
			errorProcessor: errorProcessor,
			pathTemplate:   pathTemplate,
			method:         method,
		}
	}
{{end}}

func ptr(in []byte) *string {
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
		err = errors.Wrap(err, "[httpclient.Transport]os.MkdirAll error")
		return
	}
	serverFile, err := os.Create(info.AbsOutputPath)
	defer func() {
		_ = serverFile.Close()
	}()
	t := template.Must(s.Parse(clientTransportTpl))
	if err = t.Execute(serverFile, info); err != nil {
		err = errors.Wrap(err, "[httpclient.Transport]t.Execute error")
		return
	}
	err = s.imports.GoImports(info.AbsOutputPath)
	// easyJSON generator
	cmd := exec.Command("/bin/sh", "-c", "easyjson -all "+info.AbsOutputPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		err = errors.Wrapf(err, "[httpclient.Transport]cmd.Output error\nCMD: %s\noutput: %s\n", cmd.String(), string(output))
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
