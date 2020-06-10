package httpclient

import (
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

const clientTpl = `// Package {{.PkgName}} ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package {{.PkgName}}

import (
	"context"

	"github.com/valyala/fasthttp"
)

// Options ...
var (
{{range .Iface.Methods}}
	{{.Name}} = option{}{{end}}
)

type option struct{}

// Option ...
type Option interface {
	Prepare(ctx context.Context, r *fasthttp.Request)
}

// {{ .Iface.Name }} ...
type {{ .Iface.Name }} interface {
{{range .Iface.Methods}}
	{{.Name}}({{joinFullVariables .Args ","}}) ({{joinFullVariables .Results ","}}){{end}}
}

type client struct {
	cli *fasthttp.HostClient
	{{range .Iface.Methods}}transport{{.Name}} {{.Name}}Transport
	{{end}}options map[interface{}]Option
}

{{range .Iface.Methods}}// {{.Name}} ...
func (s *client) {{.Name}}({{joinFullVariables .Args ","}}) ({{joinFullVariables .Results ","}}) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[{{.Name}}]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transport{{.Name}}.EncodeRequest(ctx, req, {{$args := popFirst .Args}}{{joinVariableNames $args ","}}); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transport{{.Name}}.DecodeResponse(ctx, res)
}
{{end}}

// NewClient the client creator
func NewClient(
	cli *fasthttp.HostClient,
	{{range .Iface.Methods}}transport{{.Name}} {{.Name}}Transport,
	{{end}}options map[interface{}]Option,
) {{ .Iface.Name }} {
	return &client{
		cli: cli,
		{{range .Iface.Methods}}transport{{.Name}}: transport{{.Name}},
		{{end}}options: options,
	}
}
`

type imports interface {
	GoImports(path string) (err error)
}

// Client ...
type Client struct {
	*template.Template
	packageName string
	filePath    []string
	imports     imports
}

func (s *Client) Generate(info api.Interface) (err error) {
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
	t := template.Must(s.Parse(clientTpl))
	if err = t.Execute(serverFile, info); err != nil {
		return
	}
	err = s.imports.GoImports(info.AbsOutputPath)
	return
}

// NewClient ...
func NewClient(template *template.Template, packageName string, filePath []string, imports imports) *Client {
	return &Client{
		Template:    template,
		packageName: packageName,
		filePath:    filePath,
		imports:     imports,
	}
}
