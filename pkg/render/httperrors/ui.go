package httperrors

import (
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

type imports interface {
	GoImports(path string) (err error)
}

// UI ...
type UI struct {
	*template.Template
	packageName string
	filePath    []string
	imports     imports
	uiTpl       string
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
	t := template.Must(s.Parse(s.uiTpl))
	if err = t.Execute(serverFile, info); err != nil {
		return
	}
	err = s.imports.GoImports(info.AbsOutputPath)
	return
}

// NewUI ...
func NewUI(template *template.Template, packageName string, filePath []string, imports imports, uiTpl string) *UI {
	return &UI{
		Template:    template,
		packageName: packageName,
		filePath:    filePath,
		imports:     imports,
		uiTpl:       uiTpl,
	}
}
