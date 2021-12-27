package service

import (
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

// Mock ...
type Mock struct {
	*template.Template
	packageName string
	filePath    []string
	imports     imports
	mockTpl     string
}

// Generate ...
func (s *Mock) Generate(info api.Interface) (err error) {
	info.PkgName = s.packageName
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
	t := template.Must(s.Parse(s.mockTpl))
	if err = t.Execute(file, info); err != nil {
		return
	}
	err = s.imports.GoImports(info.AbsOutputPath)
	return
}

// NewMock ...
func NewMock(template *template.Template, packageName string, filePath []string, imports imports, mockTpl string) *Mock {
	return &Mock{
		Template:    template,
		packageName: packageName,
		filePath:    filePath,
		imports:     imports,
		mockTpl:     mockTpl,
	}
}
