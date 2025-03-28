package httpclient

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

// Client ...
type Client struct {
	*template.Template
	packageName string
	filePath    []string
	imports     imports
	clientTpl   string
}

// Generate ...
func (s *Client) Generate(info api.Interface) (err error) {
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
	t := template.Must(s.Parse(s.clientTpl))
	if err = t.Execute(serverFile, info); err != nil {
		return
	}
	err = s.imports.GoImports(info.AbsOutputPath)

	return
}

// NewClient ...
func NewClient(template *template.Template, packageName string, filePath []string, imports imports, clientTpl string) *Client {
	return &Client{
		Template:    template,
		packageName: packageName,
		filePath:    filePath,
		imports:     imports,
		clientTpl:   clientTpl,
	}
}
