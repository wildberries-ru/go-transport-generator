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

const clientTestsTpl = `// Package {{.PkgName}} ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package {{.PkgName}}

import (
	"context"
	"encoding/json"
	"github.com/bxcodec/faker/v3"
	"github.com/valyala/fasthttp"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

{{range .Iface.Methods}}
func Test_client_{{.Name}}(t *testing.T) {
	{{range $index, $tp := .Args}}
	{{if ne $index 0}}
	var {{$tp.Name}} {{$tp.Type}}
	_ = faker.FakeData(&{{$tp.Name}})
	{{end}}
	{{end}}

	{{range $index, $tp := .Results}}
	{{$isErr := isError $tp.Type}}
	{{if not $isErr}}
	var {{$tp.Name}} {{$tp.Type}}
	_ = faker.FakeData(&{{$tp.Name}})
	{{end}}
	{{end}}
		
	maxConns := rand.Int() + 1
	opts := map[interface{}]Option{}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		{{$lenRes := lenVariables .Results}}
		{{if eq $lenRes 2}}
		{{$first := index .Results 0}}
		result := {{$first.Name}}
		{{else}}
		result := struct{
			{{range $index, $tp := .Results}}
			{{$isErr := isError $tp.Type}}
			{{if not $isErr}}{{up $tp.Name}} {{$tp.Type}} ` + "`" + `json:"{{up $tp.Name}}"` + "`" + `{{end}}
			{{end}}
		}{
			{{range $index, $tp := .Results}}
			{{$isErr := isError $tp.Type}}
			{{if not $isErr}}
			{{up $tp.Name}}: {{$tp.Name}},
			{{end}}
			{{end}}
		}
		{{end}}
		b, _ := json.Marshal(result)
		w.Write(b)
	}))
	defer ts.Close()

	parsedServerURL, _ := url.Parse(ts.URL)

	hostClient := &fasthttp.HostClient{
		Addr:     parsedServerURL.Host,
		MaxConns: maxConns,
	}

	transport{{.Name}} := New{{.Name}}Transport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClient{{.Name}},
		httpMethod{{.Name}},
	)

	type fields struct {
		cli *fasthttp.HostClient
		transport{{.Name}} {{.Name}}Transport
		options map[interface{}]Option
	}
	type args struct {
		{{range $index, $tp := .Args}}{{$tp}}
		{{end}}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		{{range $index, $tp := .Results}}
		{{$isErr := isError $tp.Type}}
		{{if not $isErr}}
		want{{up $tp.Name}} {{$tp.Type}}
		{{end}}{{end}}
		wantErr bool
	}{
		{
			"test {{.Name}}", 
			fields{hostClient, transport{{.Name}}, opts}, 
			args{context.Background(), {{range $index, $tp := .Args}}{{if ne $index 0}}{{$tp.Name}},{{end}}{{end}}}, 
			{{range $index, $tp := .Results}}{{$isErr := isError $tp.Type}}{{if not $isErr}}{{$tp.Name}},{{end}}
			{{end}}
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &client{
				cli:                  tt.fields.cli,
				transport{{.Name}}:   tt.fields.transport{{.Name}},
				options:              tt.fields.options,
			}
			{{range $index, $tp := .Results}}{{$isErr := isError $tp.Type}}{{if not $isErr}}got{{up $tp.Name}},{{end}}{{end}} err := s.{{.Name}}(tt.args.ctx, {{range $index, $tp := .Args}}{{if ne $index 0}}tt.args.{{$tp.Name}},{{end}}{{end}})
			if (err != nil) != tt.wantErr {
				t.Errorf("client.{{.Name}}() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			{{range $index, $tp := .Results}}
			{{$isErr := isError $tp.Type}}
			{{if not $isErr}}
			if !reflect.DeepEqual(got{{up $tp.Name}}, tt.want{{up $tp.Name}}) {
				t.Errorf("client.{{.Name}}() = %v, want %v", got{{up $tp.Name}}, tt.want{{up $tp.Name}})
			}
			{{end}}
			{{end}}
		})
	}
}

{{end}}

func TestNewClient(t *testing.T) {
	serverURL := fmt.Sprintf("https://%v.com", time.Now().UnixNano())
	parsedServerURL, _ := url.Parse(serverURL)
	hostClient := &fasthttp.HostClient{
		Addr:     parsedServerURL.Host,
		MaxConns: rand.Int(),
	}
	opts := map[interface{}]Option{}

	{{range .Iface.Methods}}transport{{.Name}} := New{{.Name}}Transport(
		&testErrorProcessor{},
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClient{{.Name}},
		httpMethod{{.Name}},
	)
	{{end}}

	cl := &client{
		hostClient,
		{{range .Iface.Methods}}transport{{.Name}},
		{{end}}opts,
	}

	type args struct {
		cli *fasthttp.HostClient
		{{range .Iface.Methods}}transport{{.Name}} {{.Name}}Transport
		{{end}}
		options map[interface{}]Option
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		{"test new client", args{hostClient, {{range .Iface.Methods}}transport{{.Name}},{{end}} opts}, cl},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.cli, {{range .Iface.Methods}}tt.args.transport{{.Name}},{{end}} tt.args.options); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
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

// Generate ...
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

	absTestPath := strings.Replace(info.AbsOutputPath, ".go", "_test.go", 1)

	serverTestFile, err := os.Create(absTestPath)
	if err != nil {
		return
	}
	defer func() {
		_ = serverTestFile.Close()
	}()
	t = template.Must(s.Parse(clientTestsTpl))
	if err = t.Execute(serverTestFile, info); err != nil {
		return
	}
	err = s.imports.GoImports(absTestPath)

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
