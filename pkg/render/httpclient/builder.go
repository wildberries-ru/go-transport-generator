package httpclient

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
	"net/url"
	"os"
	"strings"


	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

const (
	{{range .Iface.Methods}}
	{{$ct := getValueMap $methods .Name}}httpMethod{{.Name}} = "{{$ct.Method}}" 
	uriPathClient{{.Name}} = "{{$ct.ClientURIPath}}"{{end}}
)

type errorProcessor interface {
	Decode(r *fasthttp.Response) error
}

// New ...
func New(
	serverURL string,
	maxConns int, 
	errorProcessor errorProcessor, 
	options map[interface{}]Option,
) (client {{ .Iface.Name }}, err error) {
	parsedServerURL, err := url.Parse(serverURL)
	if err != nil {
		err = fmt.Errorf("failed to parse apiserver url", err)
		return
	}
	{{range .Iface.Methods}}transport{{.Name}} := New{{.Name}}Transport(
		errorProcessor,
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClient{{.Name}},
		httpMethod{{.Name}},
	)
	{{end}}
	client = NewClient(
		&fasthttp.HostClient{
			Addr:     parsedServerURL.Host,
			MaxConns: maxConns,
			{{if .IsTLSClient}}IsTLS:    true,{{end}}
		},
		{{range .Iface.Methods}}transport{{.Name}},
		{{end}}options,
	)
	return
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
