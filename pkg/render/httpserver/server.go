package httpserver

import (
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

const serverTpl = `// Package {{.PkgName}} ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package {{.PkgName}}

import (
	"context"

	"github.com/valyala/fasthttp"
)

type service interface {
{{range .Iface.Methods}}
	{{.Name}}({{joinFullVariables .Args ","}}) ({{joinFullVariables .Results ","}}){{end}}
}

{{range .Iface.Methods}}
type {{low .Name}} struct {
	transport      {{.Name}}Transport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *{{low .Name}}) ServeHTTP(ctx *fasthttp.RequestCtx) {
	var (
		{{$args := popFirst .Args -}}
		{{range $i, $arg := $args -}}
			{{$arg.String}}
		{{end -}}
		{{$args := popLast .Results -}}
		{{range $i, $arg := $args -}}
			{{$arg.String}}
		{{end -}}
		err error
	)
	{{$args := popFirst .Args -}}
	{{joinVariableNames $args "," "err"}} = s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	{{joinVariableNames .Results ","}} = s.service.{{.Name}}({{joinVariableNames .Args ","}})
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	{{ $args := popLast .Results -}}
	if err = s.transport.EncodeResponse(ctx, &ctx.Response, {{joinVariableNames $args ","}}); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// New{{.Name}} the server creator
func New{{.Name}}(transport {{.Name}}Transport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := {{low .Name}}{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}{{end}}
`

type imports interface {
	GoImports(path string) (err error)
}

// Server ...
type Server struct {
	*template.Template
	packageName string
	filePath    []string
	imports     imports
}

// Generate ...
func (s *Server) Generate(info api.Interface) (err error) {
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
	t := template.Must(s.Parse(serverTpl))
	if err = t.Execute(serverFile, info); err != nil {
		return
	}
	err = s.imports.GoImports(info.AbsOutputPath)
	return
}

// NewServer ...
func NewServer(template *template.Template, packageName string, filePath []string, imports imports) *Server {
	return &Server{
		Template:    template,
		packageName: packageName,
		filePath:    filePath,
		imports:     imports,
	}
}
