package service

import (
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

const instrumentingTpl = `// Package {{.PkgName}} ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package {{.PkgName}}

import (
	"context"
	"strconv"
	"time"

	"github.com/go-kit/kit/metrics"
)

// instrumentingMiddleware wraps Service and enables request metrics
type instrumentingMiddleware struct {
	reqCount    metrics.Counter
	reqDuration metrics.Histogram
	svc         {{ .Iface.Name }}
}

{{range .Iface.Methods -}}
// {{.Name}} ...
func (s *instrumentingMiddleware) {{.Name}}({{joinFullVariables .Args ","}}) ({{joinFullVariables .Results ","}}) {
	defer s.recordMetrics("{{.Name}}", time.Now(), err)
	return s.svc.{{.Name}}({{joinVariableNames .Args ","}})
}
{{end}}

func (s *instrumentingMiddleware) recordMetrics(method string, startTime time.Time, err error) {
	labels := []string{
		"method", method,
		"error", strconv.FormatBool(err != nil),
	}
	s.reqCount.With(labels...).Add(1)
	s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
}

// NewInstrumentingMiddleware ...
func NewInstrumentingMiddleware(reqCount metrics.Counter, reqDuration metrics.Histogram, svc {{ .Iface.Name }}) {{ .Iface.Name }} {
	return &instrumentingMiddleware{
		reqCount:    reqCount,
		reqDuration: reqDuration,
		svc:         svc,
	}
}
`

type imports interface {
	GoImports(path string) (err error)
}

// Instrumenting ...
type Instrumenting struct {
	*template.Template
	filePath []string
	imports  imports
}

// Generate ...
func (s *Instrumenting) Generate(info api.Interface) (err error) {
	info.PkgName = path.Base(info.AbsOutputPath)
	info.AbsOutputPath = strings.Join(append(strings.Split(info.AbsOutputPath, "/"), s.filePath...), "/")
	dir, _ := path.Split(info.AbsOutputPath)
	err = os.MkdirAll(dir, 0750)
	if err != nil {
		return
	}
	serverFile, err := os.Create(info.AbsOutputPath)
	defer func() {
		_ = serverFile.Close()
	}()
	t := template.Must(s.Parse(instrumentingTpl))
	if err = t.Execute(serverFile, info); err != nil {
		return
	}
	err = s.imports.GoImports(info.AbsOutputPath)
	return
}

// NewInstrumenting ...
func NewInstrumenting(template *template.Template, filePath []string, imports imports) *Instrumenting {
	return &Instrumenting{
		Template: template,
		filePath: filePath,
		imports:  imports,
	}
}
