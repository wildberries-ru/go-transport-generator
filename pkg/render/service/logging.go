package service

import (
	"os"
	"path"
	"runtime"
	"strings"
	"text/template"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

const loggingTpl = `// Package {{.PkgName}} ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package {{.PkgName}}

import (
	"context"
	"time"

	"github.com/wildberries-ru/go-transport-generator/log/logger"
)

// loggingMiddleware wraps Service and logs request information to the provided logger
type loggingMiddleware struct {
	logger logger.Logger
	svc    {{ .Iface.Name }}
}
{{$methods := .HTTPMethods}}
{{range .Iface.Methods -}}
	{{$method := index $methods .Name}}
	// {{.Name}} ...
	func (s *loggingMiddleware) {{.Name}}({{joinFullVariables .Args ","}}) ({{joinFullVariables .Results ","}}) {
		defer func(begin time.Time) {
			lg := s.logger.WithError(err).WithFields(
				map[string]interface{} {
					{{$args := popFirst .Args -}}
					{{range $arg := $args -}}
						{{if notin $method.LogIgnores $arg.Name}}"{{$arg.Name}}": {{$arg.Name}},{{end}} 
					{{end -}}
					{{$args := popLast .Results -}}
					{{range $arg := $args -}}
						{{if notin $method.LogIgnores $arg.Name}}"{{$arg.Name}}": {{$arg.Name}},{{end}}
					{{end -}}
					"elapsed": time.Since(begin),
				},
			)
			if err != nil {
				lg.Debug("{{.Name}}")
			} else {
				lg.Error("{{.Name}}")
			}
		}(time.Now())
		return s.svc.{{.Name}}({{joinVariableNamesWithEllipsis .Args ","}})
	}
{{end}}

// NewLoggingMiddleware ...
func NewLoggingMiddleware(logger logger.Logger, svc {{ .Iface.Name }}) {{ .Iface.Name }} {
	return &loggingMiddleware{
		logger: logger,
		svc:    svc,
	}
}
`

// Logging ...
type Logging struct {
	*template.Template
	filePath []string
	imports  imports
}

// Generate ...
func (s *Logging) Generate(info api.Interface) (err error) {
	if runtime.GOOS == "windows" {
		info.AbsOutputPath = strings.Replace(info.AbsOutputPath, `\`, "/", -1)
	}
	info.PkgName = path.Base(info.AbsOutputPath)
	info.AbsOutputPath = strings.Join(append(strings.Split(info.AbsOutputPath, "/"), s.filePath...), "/")
	dir, _ := path.Split(info.AbsOutputPath)
	err = os.MkdirAll(dir, 0750)
	if err != nil {
		return
	}
	file, err := os.Create(info.AbsOutputPath)
	defer func() {
		_ = file.Close()
	}()
	t := template.Must(s.Parse(loggingTpl))
	if err = t.Execute(file, info); err != nil {
		return
	}
	err = s.imports.GoImports(info.AbsOutputPath)
	return
}

// NewLogging ...
func NewLogging(template *template.Template, filePath []string, imports imports) *Logging {
	return &Logging{
		Template: template,
		filePath: filePath,
		imports:  imports,
	}
}
