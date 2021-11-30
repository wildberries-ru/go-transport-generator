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

	"github.com/pkg/errors"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

const (
	{{range .Iface.Methods}}
	{{$ct := index $methods .Name}}httpMethod{{.Name}} = "{{$ct.Method}}" 
	uriPathClient{{.Name}} = "{{$ct.ClientURIPath}}"{{end}}
)

type errorProcessor interface {
	Decode(r *fasthttp.Response) error
}

// Config ...
type Config struct {
	ServerURL           string
	MaxConns            *int
	MaxConnDuration     *time.Duration
	MaxIdleConnDuration *time.Duration
	ReadBufferSize      *int
	WriteBufferSize     *int
	ReadTimeout         *time.Duration
	WriteTimeout        *time.Duration
	MaxResponseBodySize *int
}

// New ...
func New(
	config Config,
	errorProcessor errorProcessor, 
	options map[interface{}]Option,
) (client {{ .Iface.Name }}, err error) {
	parsedServerURL, err := url.Parse(config.ServerURL)
	if err != nil {
		err = errors.Wrap(err, "failed to parse server url")
		return
	}
	{{range .Iface.Methods}}transport{{.Name}} := New{{.Name}}Transport(
		errorProcessor,
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClient{{.Name}},
		httpMethod{{.Name}},
	)
	{{end}}

	cli := fasthttp.HostClient{
		Addr: parsedServerURL.Host,
	}
	if config.MaxConns != nil {
		cli.MaxConns = *config.MaxConns
	}
	if config.MaxConnDuration != nil {
		cli.MaxConnDuration = *config.MaxConnDuration
	}
	if config.MaxIdleConnDuration != nil {
		cli.MaxIdleConnDuration = *config.MaxIdleConnDuration
	}
	if config.ReadBufferSize != nil {
		cli.ReadBufferSize = *config.ReadBufferSize
	}
	if config.WriteBufferSize != nil {
		cli.WriteBufferSize = *config.WriteBufferSize
	}
	if config.ReadTimeout != nil {
		cli.ReadTimeout = *config.ReadTimeout
	}
	if config.WriteTimeout != nil {
		cli.WriteTimeout = *config.WriteTimeout
	}
	if config.MaxResponseBodySize != nil {
		cli.MaxResponseBodySize = *config.MaxResponseBodySize
	}
	{{if .IsTLSClient}}cli.IsTLS = true{{if .IsInsecureTLS}}
	cli.TLSConfig = &tls.Config{InsecureSkipVerify: true}{{end}}{{end}}

	client = NewClient(
		&cli,
		{{range .Iface.Methods}}transport{{.Name}},
		{{end}}options,
	)
	return
}
`

const builderTestsTpl = `// Package {{.PkgName}} ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package {{.PkgName}}

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"math/rand"
	"net/url"
	"reflect"
	"testing"
	"time"
)

type testErrorProcessor struct{}

func TestNew(t *testing.T) {
	serverURL := fmt.Sprintf("https://%v.com", time.Now().UnixNano())
	parsedServerURL, _ := url.Parse(serverURL)
	maxConns := rand.Int()
	opts := map[interface{}]Option{}

	{{range .Iface.Methods}}transport{{.Name}} := New{{.Name}}Transport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+parsedServerURL.Path+uriPathClient{{.Name}},
		httpMethod{{.Name}},
	)
	{{end}}

	cl := client{
		&fasthttp.HostClient{
			Addr:     parsedServerURL.Host,
			MaxConns: maxConns,
			{{if .IsTLSClient}}IsTLS:    true,{{if .IsInsecureTLS}}
			TLSConfig: &tls.Config{InsecureSkipVerify: true},{{end}}{{end}}
		},
		{{range .Iface.Methods}}transport{{.Name}},
		{{end}}opts,
	}

	type args struct {
		serverURL      string
		maxConns       int
		errorProcessor errorProcessor
		options        map[interface{}]Option
	}
	tests := []struct {
		name       string
		args       args
		wantClient {{ .Iface.Name }}
		wantErr    bool
	}{
		{"test new builder", args{serverURL, maxConns, &testErrorProcessor{}, opts}, &cl, false},
		{"test new builder incorrect URL", args{" http:example%20.com", maxConns, &testErrorProcessor{}, opts}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotClient, err := New(tt.args.serverURL, tt.args.maxConns, tt.args.errorProcessor, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotClient, tt.wantClient) {
				t.Errorf("New() = %v, want %v", gotClient, tt.wantClient)
			}
		})
	}
}

func (ep *testErrorProcessor) Decode(r *fasthttp.Response) error {
	return nil
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

	absTestPath := strings.Replace(info.AbsOutputPath, ".go", "_test.go", 1)

	serverTestFile, err := os.Create(absTestPath)
	if err != nil {
		return
	}
	defer func() {
		_ = serverTestFile.Close()
	}()
	t = template.Must(s.Parse(builderTestsTpl))
	if err = t.Execute(serverTestFile, info); err != nil {
		return
	}
	err = s.imports.GoImports(absTestPath)

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
