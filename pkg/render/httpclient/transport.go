package httpclient

import (
	"os"
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
	"encoding/json"

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
	{{$jsonTags := $ct.JSONTags}}
	{{$plainObject := $ct.PlainObject}}
	{{$multipartValueTags := $ct.MultipartValueTags}}
	{{$multipartFileTags := $ct.MultipartFileTags}}
	{{$formUrlencodedTags := $ct.FormUrlencodedTags}}
	{{$responseJSONTags := $ct.ResponseJSONTags}}
	{{$responseCookiesPlaceholders := $ct.ResponseCookies}}
	{{$responseHeaderPlaceholders := $ct.ResponseHeaders}}
	{{$responseStatus := $ct.ResponseStatus}}
	{{$responseContentType := $ct.ResponseContentType}}
	{{$responseFile := $ct.ResponseFile}}
	{{$responseBody := $ct.ResponseBody}}
	{{$responseBodyField := $ct.ResponseBodyField}}
	{{$responseBodyType := index $responseBody $ct.ResponseBodyField}}
	{{$responseBodyTypeIsSlice := isSliceType $responseBodyType}}
	{{$responseBodyTypeIsMap := isMapType $responseBodyType}}

	{{if lenMap $body}}
		{{$bodyLen := lenMap $body}}{{$n := .Name}}{{if $plainObject}}{{range $name, $tp := $body}}type {{low $n}}Request {{$tp}}
			{{end}}
		{{else}}
			type {{low .Name}}Request struct {
				{{range $name, $tp := $body}}{{up $name}} {{$tp}}{{$tag := index $jsonTags $name}}{{if $tag}} ` + "`" + `json:"{{$tag}}"` + "`" + `{{end}}
				{{end}}
			}
		{{end}}
	{{end}}

	{{if lenMap $responseBody}}
		type {{low .Name}}Response {{if or $responseBodyTypeIsSlice $responseBodyTypeIsMap}}{{$responseBodyType}}{{else}} struct {
  			{{if $responseBodyType}}
    			{{$responseBodyType}}
  			{{else}}
			    {{$respLen := lenMap $responseBody}}
			    {{range $name, $tp := $responseBody}}{{up $name}} {{$tp}}{{$tag := index $responseJSONTags $name}}{{if $tag}} ` + "`" + `json:"{{$tag}}"` + "`" + `{{end}}
    			{{end}}
  			{{end}}
		}{{end}}
	{{end}}

	// {{.Name}}Transport transport interface
	type {{.Name}}Transport interface {
		EncodeRequest(ctx context.Context, r *fasthttp.Request, {{$args := popFirst .Args}}{{joinFullVariables $args ","}}) (err error)
		DecodeResponse(ctx context.Context, r *fasthttp.Response) ({{$args := popLast .Results}}{{joinFullVariables $args "," "err error"}})
	}

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
		{{$qstr := ""}}
		{{$qc:="1"}}
		{{range $from, $to := $queryPlaceholders}}
			{{if eq $qc "1"}} 
				{{$qstr = concat $qstr "\""}}
			{{else}}
				{{$qstr = concat $qstr "+\"&"}}
			{{end}}
			{{$qstr = concat $qstr $from}}
			{{$qstr = concat $qstr "=\"+"}}
			{{$qstr = concat $qstr $to.Name}}
			{{$qc = concat $qc "1"}}
		{{end}}
		{{if ne $qc "1"}}
			r.URI().SetQueryString({{$qstr}})
		{{end}}
		{{$clen := lenMap $cookiePlaceholders}}
		{{if gt $clen 0}}
			// cookies must be the *string type
			{{range $to, $from := $cookiePlaceholders}}
				if {{$from}} != nil {
					r.Header.SetCookie("{{$to}}", *{{$from}})
				}
			{{end}}
		{{end}}
		{{range $to, $from := $headerPlaceholders}}
			r.Header.Set("{{$to}}", *{{$from}})
		{{end}}
		{{if eq $contentType "application/json"}}r.Header.Set("Content-Type", "application/json")
		    {{$requestName := low .Name}}
			{{if lenMap $body}}var request {{$requestName}}Request
				{{if $plainObject}}
					{{range $name, $tp := $body}}
						request = {{$requestName}}Request({{$name}})
					{{end}}
				{{else}}
					{{range $name, $tp := $body}}
						request.{{up $name}} = {{$name}}
					{{end}}
				{{end}}
				body, err := json.Marshal(request)
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
				{{if $from.IsSlice}}
					// todo generate then {{$from.Name}} is slice
				{{else}}
					{{if $from.IsPointer}}
						if {{$from.Name}} != nil {
							{{if $from.IsString}}
								writer.WriteField("{{$to}}", *{{$from.Name}})
							{{else if $from.IsInt}}
								writer.WriteField("{{$to}}", fmt.Sprintf("%d", *{{$from.Name}}))
							{{else if $from.IsBool}}
								writer.WriteField("{{$to}}", fmt.Sprintf("%t", *{{$from.Name}}))
							{{end}}
						}
					{{else}}
						{{if $from.IsString}}
							writer.WriteField("{{$to}}", {{$from.Name}})
						{{else if $from.IsInt}}
							writer.WriteField("{{$to}}", fmt.Sprintf("%d", {{$from.Name}}))
						{{else if $from.IsBool}}
							writer.WriteField("{{$to}}", fmt.Sprintf("%t", {{$from.Name}}))
						{{end}}
					{{end}}
				{{end}}
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

		{{if eq $contentType "application/octet-stream"}}
			r.Header.Set("Content-Type", "application/octet-stream")
			{{if lenMap $body}}
				{{range $name, $tp := $body}}
					r.SetBody({{$name}})
				{{end}}
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
		{{$clen := lenMap $responseCookiesPlaceholders}}
		{{if gt $clen 0}}
			cookie := fasthttp.AcquireCookie()
			{{range $from, $to := $responseCookiesPlaceholders}}
				// cookies must be a *string type
				_{{$to}}:=string(r.Header.PeekCookie("{{$from}}"))
				{{$to}} = &_{{$to}}
			{{end}}
			fasthttp.ReleaseCookie(cookie)
		{{end}}
		{{range $from, $to := $responseHeaderPlaceholders}}
			{{$to}} = ptr(r.Header.Peek("{{$from}}"))
		{{end}}
		{{if eq $responseContentType "application/json"}}
			{{if lenMap $responseBody}}var theResponse {{low .Name}}Response
				if err = json.Unmarshal(r.Body(), &theResponse); err != nil {
					return
				}
				{{if $responseBodyType}}
					{{$responseBodyField}} = theResponse{{if or $responseBodyTypeIsSlice $responseBodyTypeIsMap}}{{else}}.{{stripType $responseBodyType}}{{end}}
				{{else}}
					{{range $name, $tp := $responseBody}}
						{{$name}} = theResponse.{{up $name}}
					{{end}}
				{{end}}
			{{end}}
		{{end}}
		{{if eq $responseContentType "application/octet-stream"}}
			{{if lenMap $responseBody}}
				b := r.Body()
				// fasthttp reuses body memory, we have to copy a response
				{{range $name, $tp := $responseBody}}{{$name}} = make([]byte, len(b))
					copy({{$name}}, b)
				{{end}}
			{{end}}
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
	t, err := s.Parse(clientTransportTpl)
	if err != nil {
		err = errors.Wrap(err, "[httpclient.Transport]t.Parse error")
		return
	}
	t = template.Must(t, err)
	if err = t.Execute(serverFile, info); err != nil {
		err = errors.Wrap(err, "[httpclient.Transport]t.Execute error")
		return
	}
	err = s.imports.GoImports(info.AbsOutputPath)
	if err != nil {
		err = errors.Wrap(err, "[httpclient.Transport]t.Go imports error")
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
