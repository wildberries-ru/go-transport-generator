package httpclient

import (
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

// ClientTests ...
type ClientTests struct {
	*template.Template
	packageName    string
	filePath       []string
	imports        imports
	clientTestsTpl string
}

// Generate ...
func (s *ClientTests) Generate(info api.Interface) (err error) {
	info.PkgName = s.packageName
	info.AbsOutputPath = strings.Join(append(strings.Split(info.AbsOutputPath, "/"), s.filePath...), "/")
	dir, _ := path.Split(info.AbsOutputPath)
	err = os.MkdirAll(dir, 0750)
	if err != nil {
		return
	}

	absTestPath := strings.Replace(info.AbsOutputPath, ".go", "_test.go", 1)

	serverTestFile, err := os.Create(absTestPath)
	if err != nil {
		return
	}
	defer func() {
		_ = serverTestFile.Close()
	}()
	t := template.Must(s.Parse(s.clientTestsTpl))
	if err = t.Execute(serverTestFile, info); err != nil {
		return
	}
	err = s.imports.GoImports(absTestPath)

	return
}

// NewClientTests ...
func NewClientTests(
	template *template.Template,
	packageName string,
	filePath []string,
	imports imports,
	clientTestsTpl string,
) *ClientTests {
	return &ClientTests{
		Template:       template,
		packageName:    packageName,
		filePath:       filePath,
		imports:        imports,
		clientTestsTpl: clientTestsTpl,
	}
}
