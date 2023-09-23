package service

import (
	"os"
	"path"
	"runtime"
	"strings"
	"text/template"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

// Logging ...
type Logging struct {
	*template.Template
	filePath   []string
	imports    imports
	loggingTpl string
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
	t := template.Must(s.Parse(s.loggingTpl))
	if err = t.Execute(file, info); err != nil {
		return
	}
	err = s.imports.GoImports(info.AbsOutputPath)
	return
}

// NewLogging ...
func NewLogging(template *template.Template, filePath []string, imports imports, loggingTpl string) *Logging {
	return &Logging{
		Template:   template,
		filePath:   filePath,
		imports:    imports,
		loggingTpl: loggingTpl,
	}
}
