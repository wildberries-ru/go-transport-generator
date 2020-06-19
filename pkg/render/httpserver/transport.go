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

	"github.com/valyala/fasthttp"
)

var (
	emptyBytes = []byte("")
)
 
{{range .Iface.Methods}}

{{$ct := getValueMap $methods .Name}}
{{$uriPlaceholders := $ct.URIPathPlaceholders}}
{{$queryPlaceholders := $ct.RawQueryPlaceholders}}
{{$isIntQueryPlaceholders := $ct.IsIntQueryPlaceholders}}
{{$headerPlaceholders := $ct.HeaderPlaceholders}}
{{$body := $ct.Body}}
{{$contentType := $ct.ContentType}}
{{$jsonTags := $ct.JsonTags}}
{{$responseHeaderPlaceholders := $ct.ResponseHeaders}}
{{$responseStatus := $ct.ResponseStatus}}
{{$responseContentType := $ct.ResponseContentType}}
{{$responseContentEncoding := $ct.ResponseContentEncoding}}
{{$responseJsonTags := $ct.ResponseJsonTags}}
{{$responseBody := $ct.ResponseBody}}
{{$responseBodyField := $ct.ResponseBodyField}}
{{$responseBodyType := index $responseBody $ct.ResponseBodyField}}
{{$responseBodyTypeIsSlice := isSliceType $responseBodyType}}

{{if lenMap $body}}type {{low .Name}}Request struct {
{{range $name, $tp := $body}}{{up $name}} {{$tp}}{{$tag := index $jsonTags $name}}{{if $tag}} ` + "`" + `json:"{{$tag}}"` + "`" + `{{end}}
{{end}}
}{{end}}

{{if lenMap $responseBody}}//easyjson:json
type {{low .Name}}Response {{if $responseBodyTypeIsSlice}}{{$responseBodyType}}{{else}} struct {
  {{if $responseBodyType}}
    {{$responseBodyType}}
  {{else}}
    {{range $name, $tp := $responseBody}}{{up $name}} {{$tp}}{{$tag := index $responseJsonTags $name}}{{if $tag}} ` + "`" + `json:"{{$tag}}"` + "`" + `{{end}}
    {{end}}
  {{end}}
}{{end}}
{{end}}

// {{.Name}}Transport transport interface
type {{.Name}}Transport interface {
	DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) ({{$args := popFirst .Args}}{{joinFullVariables $args "," "err error"}})
	EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, {{$args := popLast .Results}}{{joinFullVariables $args ","}}) (err error)
}

type {{low .Name}}Transport struct {
	{{if eq $contentType "application/json"}}{{if lenMap $body}}decodeJSONErrorCreator errorCreator{{end}}{{end}}
	{{if eq $responseContentType "application/json"}}{{if lenMap $responseBody}}encodeJSONErrorCreator errorCreator{{end}}{{end}}
	{{if $isIntQueryPlaceholders}}encodeQueryTypeIntErrorCreator errorCreator{{end}}
}

// DecodeRequest method for decoding requests on server side
func (t *{{low .Name}}Transport) DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) ({{$args := popFirst .Args}}{{joinFullVariables $args "," "err error"}}) {
	{{range $i, $sc := $uriPlaceholders}}{{$sc}} = ctx.UserValue("{{$sc}}").(string)
	{{end}}{{range $from, $to := $queryPlaceholders}}{{if eq $to.IsString true}}{{$to.Name}} = {{if eq $to.IsPointer true}}ptr{{else}}string{{end}}(ctx.QueryArgs().Peek("{{$from}}"))
	{{else if eq $to.IsInt true}}{{if eq $to.Type "int"}}{{$to.Name}}, err = atoi{{if eq $to.IsPointer true}}ptr{{end}}(ctx.QueryArgs().Peek("{{$from}}"))
	{{else}}_{{$to.Name}}, err := atoi(ctx.QueryArgs().Peek("{{$from}}")){{end}}
	if err != nil {
		err = t.encodeQueryTypeIntErrorCreator(err)
		return
	}
	{{if ne $to.Type "int"}}{{if eq $to.IsPointer true}}__{{$to.Name}} := {{$to.Type}}(_{{$to.Name}})
	{{$to.Name}} = &__{{$to.Name}}
	{{else}}{{$to.Name}} = {{$to.Type}}(_{{$to.Name}}){{end}}{{end}}
	{{end}}{{end}}{{range $from, $to := $headerPlaceholders}}{{$to}} = ptr(r.Header.Peek("{{$from}}"))
	{{end}}{{if eq $contentType "application/json"}}{{if lenMap $body}}var request {{low .Name}}Request
	if err = request.UnmarshalJSON(r.Body()); err != nil {
		err = t.decodeJSONErrorCreator(err)
		return
	}
	{{range $name, $tp := $body}}{{$name}} = request.{{up $name}}
	{{end}}{{end}}{{end}}return
}

// EncodeResponse method for encoding response on server side
func (t *{{low .Name}}Transport) EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, {{$args := popLast .Results}}{{joinFullVariables $args ","}}) (err error) {
	{{if eq $responseContentType "application/json"}}r.Header.Set("Content-Type", "application/json{{if $responseContentEncoding}}; charset={{$responseContentEncoding}}{{end}}")
	{{if lenMap $responseBody}}var theResponse {{low .Name}}Response
	{{if $responseBodyType }}theResponse{{if not $responseBodyTypeIsSlice}}.{{stripType $responseBodyType}}{{end}} = {{$responseBodyField}}
	{{else}}{{range $name, $tp := $responseBody}}theResponse.{{up $name}} = {{$name}}
	{{end}}{{end}}body, err := theResponse.MarshalJSON()
	if err != nil {
		err = t.encodeJSONErrorCreator(err)
		return
	}
	r.SetBody(body)
	{{end}}{{end}}{{range $to, $from := $responseHeaderPlaceholders}}r.Header.Set("{{$to}}", "{{$from}}")
	{{end}}r.Header.SetStatusCode({{$responseStatus}})
	return
}

// New{{.Name}}Transport the transport creator for http requests
func New{{.Name}}Transport({{if eq $contentType "application/json"}}{{if lenMap $body}}decodeJSONErrorCreator errorCreator,{{end}}{{end}} {{if eq $responseContentType "application/json"}}{{if lenMap $responseBody}}encodeJSONErrorCreator errorCreator,{{end}}{{end}}{{if eq $isIntQueryPlaceholders true}}encodeQueryTypeIntErrorCreator errorCreator{{end}}) {{.Name}}Transport {
	return &{{low .Name}}Transport{
		{{if eq $contentType "application/json"}}{{if lenMap $body}}decodeJSONErrorCreator: decodeJSONErrorCreator,
		{{end}}{{end}}{{if eq $responseContentType "application/json"}}{{if lenMap $responseBody}}encodeJSONErrorCreator: encodeJSONErrorCreator,
		{{end}}{{end}}{{if $isIntQueryPlaceholders}}encodeQueryTypeIntErrorCreator: encodeQueryTypeIntErrorCreator,{{end}}
	}
}{{end}}

func ptr(in []byte) *string {
	i := string(in)
	return &i
}

func atoiptr(in []byte) (out *int, err error) {
	var (
		o int
		i = string(in)
	)
	if i != "" {
		if o, err = strconv.Atoi(i); err == nil {
			out = &o
		}
	}
	return
}

func atoi(in []byte) (out int, err error) {
	if bytes.Equal(in, emptyBytes) {
		return 
	}
	return strconv.Atoi(string(in))
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
