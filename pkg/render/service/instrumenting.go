package service

import (
	"os"
	"path"
	"runtime"
	"strings"
	"text/template"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

type imports interface {
	GoImports(path string) (err error)
}

// Instrumenting ...
type Instrumenting struct {
	*template.Template
	filePath         []string
	imports          imports
	instrumentingTpl string
}

// Generate ...
func (s *Instrumenting) Generate(info api.Interface) (err error) {
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
	serverFile, err := os.Create(info.AbsOutputPath)
	defer func() {
		_ = serverFile.Close()
	}()
	t := template.Must(s.Parse(s.instrumentingTpl))
	if err = t.Execute(serverFile, info); err != nil {
		return
	}
	err = s.imports.GoImports(info.AbsOutputPath)
	return
}

// NewInstrumenting ...
func NewInstrumenting(template *template.Template, filePath []string, imports imports, instrumentingTpl string) *Instrumenting {
	return &Instrumenting{
		Template:         template,
		filePath:         filePath,
		imports:          imports,
		instrumentingTpl: instrumentingTpl,
	}
}
