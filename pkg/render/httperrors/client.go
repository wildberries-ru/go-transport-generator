package httperrors

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

// ClientErrorProcessor ...
type ClientErrorProcessor struct {
	defaultCode    int
	defaultMessage string
}

// Encode writes a svc error to the given http.ResponseWriter.
func (e *ClientErrorProcessor) Encode(ctx context.Context, r *fasthttp.Response, err error) {
	code := e.defaultCode
	message := e.defaultMessage
	if err, ok := err.(*httpError); ok {
		if err.Code != e.defaultCode {
			code = err.Code
			message = err.Message
		}
	}
	r.SetStatusCode(code)
	r.SetBodyString(message)
	return
}

// Decode reads a Service error from the given *http.Response.
func (e *ClientErrorProcessor) Decode(r *fasthttp.Response) error {
	msgBytes := r.Body()
	msg := strings.TrimSpace(string(msgBytes))
	if msg == "" {
		msg = http.StatusText(r.StatusCode())
	}
	return &httpError{
		Code:    r.StatusCode(),
		Message: msg,
	}
}

// NewClientErrorProcessor ...
func NewClientErrorProcessor(defaultCode int, defaultMessage string) *ClientErrorProcessor {
	return &ClientErrorProcessor{
		defaultCode:    defaultCode,
		defaultMessage: defaultMessage,
	}
}
`

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
