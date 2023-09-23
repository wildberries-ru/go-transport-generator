package httpserver

import (
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

// Builder ...
type Builder struct {
	*template.Template
	packageName string
	filePath    []string
	imports     imports
	builderTpl  string
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
	t := template.Must(s.Parse(s.builderTpl))
	if err = t.Execute(serverFile, info); err != nil {
		return
	}
	err = s.imports.GoImports(info.AbsOutputPath)
	return
}

// NewBuilder ...
func NewBuilder(template *template.Template, packageName string, filePath []string, imports imports, builderTpl string) *Builder {
	return &Builder{
		Template:    template,
		packageName: packageName,
		filePath:    filePath,
		imports:     imports,
		builderTpl:  builderTpl,
	}
}
