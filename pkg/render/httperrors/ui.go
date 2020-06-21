package httperrors

import (
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

const uiTpl = `// Package {{.PkgName}} ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package {{.PkgName}}

type httpError struct {
	Code    int
	Message string
}

// Error returns a text message corresponding to the given error.
func (e *httpError) Error() string {
	return e.Message
}

// StatusCode returns an HTTP status code corresponding to the given error.
func (e *httpError) StatusCode() int {
	return e.Code
}

type errorResponse struct {
	Error            bool              ` + "`" + `json:"error"` + "`" + `
	ErrorText        string            ` + "`" + `json:"errorText"` + "`" + `
	AdditionalErrors map[string]string ` + "`" + `json:"additionalErrors"` + "`" + `
	Data             *struct{}         ` + "`" + `json:"data"` + "`" + `
}

// ErrorProcessor ...
type ErrorProcessor struct {
	errors map[string]string
}

//Encode writes a svc error to the given http.ResponseWriter.
func (e *ErrorProcessor) Encode(ctx context.Context, r *fasthttp.Response, err error) {
	errorText := err.Error()
	if idx := strings.Index(err.Error(), ":"); idx != -1 {
		numberOfError := err.Error()[:idx]
		if text, ok := e.errors[numberOfError]; ok {
			errorText = text
		}
	}
	res := errorResponse{
		Error:     true,
		ErrorText: errorText,
	}
	r.SetStatusCode(200)
	r.Header.Set("Content-Type", "application/json")
	body, err := json.Marshal(res)
	if err != nil {
		return
	}
	r.SetBody(body)
}

// Decode reads a Service error from the given *http.Response.
func (e *ErrorProcessor) Decode(r *fasthttp.Response) error {
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

// NewErrorProcessor ...
func NewErrorProcessor(errors map[string]string) *ErrorProcessor {
	return &ErrorProcessor{
		errors: errors,
	}
}
`

type imports interface {
	GoImports(path string) (err error)
}

// UI ...
type UI struct {
	*template.Template
	packageName string
	filePath    []string
	imports     imports
}

// Generate ...
func (s *UI) Generate(info api.Interface) (err error) {
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
	t := template.Must(s.Parse(uiTpl))
	if err = t.Execute(serverFile, info); err != nil {
		return
	}
	err = s.imports.GoImports(info.AbsOutputPath)
	return
}

// NewUI ...
func NewUI(template *template.Template, packageName string, filePath []string, imports imports) *UI {
	return &UI{
		Template:    template,
		packageName: packageName,
		filePath:    filePath,
		imports:     imports,
	}
}
