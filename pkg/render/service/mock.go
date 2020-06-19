package service

import (
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

const mockTpl = `// Package {{.PkgName}} ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package {{.PkgName}}

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// MockService ...
type MockService struct {
	mock.Mock
}

{{range .Iface.Methods -}}
// {{.Name}} ...
func (s *MockService) {{.Name}}({{joinFullVariables .Args ","}}) ({{joinFullVariables .Results ","}}) {
	{{$args := popFirst .Args -}}
	{{$res := popLast .Results -}}
	args := s.Called(context.Background(), {{joinVariableNames $args ","}})
	return {{range $i, $a := $res}}args.Get({{$i}}).({{$a.Type}}), {{end}}args.Error({{lenVariables $res}})
}
{{end}}
`

// Mock ...
type Mock struct {
	*template.Template
	packageName string
	filePath    []string
	imports     imports
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
	t := template.Must(s.Parse(mockTpl))
	if err = t.Execute(file, info); err != nil {
		return
	}
	err = s.imports.GoImports(info.AbsOutputPath)
	return
}

// NewMock ...
func NewMock(template *template.Template, packageName string, filePath []string, imports imports) *Mock {
	return &Mock{
		Template:    template,
		packageName: packageName,
		filePath:    filePath,
		imports:     imports,
	}
}
