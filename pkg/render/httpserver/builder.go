package httpserver

import (
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

const builderTpl = `// Package {{.PkgName}} ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package {{.PkgName}}
{{$methods := .HTTPMethods}}
import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"


	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

const (
	{{range .Iface.Methods}}
	{{$ct := index $methods .Name}}httpMethod{{.Name}} = "{{$ct.Method}}" 
	uriPath{{.Name}} = "{{$ct.URIPath}}"{{end}}
)

type errorProcessor interface {
	Encode(ctx context.Context, r *fasthttp.Response, err error)
}

type errorCreator func(err error) error

// New ...
func New(router *fasthttprouter.Router, svc service, decodeJSONErrorCreator errorCreator, encodeJSONErrorCreator errorCreator, decodeTypeIntErrorCreator errorCreator, errorProcessor errorProcessor) {
	{{range .Iface.Methods}}
		{{$ct := index $methods .Name}}
		{{$body := $ct.Body}}
		{{$contentType := $ct.ContentType}}
		{{$isIntQueryPlaceholders := $ct.IsIntQueryPlaceholders}}
		{{$isIntBodyPlaceholders := $ct.IsIntBodyPlaceholders}}
		{{$responseContentType := $ct.ResponseContentType}}
		{{$responseBody := $ct.ResponseBody}}{{low .Name}}Transport := New{{.Name}}Transport(
			{{if eq $contentType "application/json"}}{{if lenMap $body}}decodeJSONErrorCreator, {{end}}{{end}}
			{{if eq $responseContentType "application/json"}}{{if lenMap $responseBody}}encodeJSONErrorCreator, {{end}}{{end}}
			{{if or $isIntQueryPlaceholders $isIntBodyPlaceholders}}decodeTypeIntErrorCreator,{{end}}
		)
		router.Handle(httpMethod{{.Name}}, uriPath{{.Name}}, New{{.Name}}({{low .Name}}Transport, svc, errorProcessor))
	{{end}}
	router.Handle("GET", "/debug/pprof/", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Index))
	router.Handle("GET", "/debug/pprof/profile", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Profile))
}
`

// Builder ...
type Builder struct {
	*template.Template
	packageName string
	filePath    []string
	imports     imports
}

// Generate ...
func (s *Builder) Generate(info api.Interface) (err error) {
	info.PkgName = s.packageName
	info.AbsOutputPath = strings.Join(append(strings.Split(info.AbsOutputPath, "/"), s.filePath...), "/")
	dir, _ := path.Split(info.AbsOutputPath)
	err = os.MkdirAll(dir, 0750)
	if err != nil {
		return
	}
	serverFile, err := os.Create(info.AbsOutputPath)
	if err != nil {
		return
	}
	defer func() {
		_ = serverFile.Close()
	}()
	t := template.Must(s.Parse(builderTpl))
	if err = t.Execute(serverFile, info); err != nil {
		return
	}
	err = s.imports.GoImports(info.AbsOutputPath)
	return
}

// NewBuilder ...
func NewBuilder(template *template.Template, packageName string, filePath []string, imports imports) *Builder {
	return &Builder{
		Template:    template,
		packageName: packageName,
		filePath:    filePath,
		imports:     imports,
	}
}
