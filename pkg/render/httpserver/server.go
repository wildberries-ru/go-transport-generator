package httpserver

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

// Server ...
type Server struct {
	*template.Template
	packageName string
	filePath    []string
	imports     imports
	serverTpl   string
}

// Generate ...
func (s *Server) Generate(info api.Interface) (err error) {
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
	t := template.Must(s.Parse(s.serverTpl))
	if err = t.Execute(serverFile, info); err != nil {
		return
	}
	err = s.imports.GoImports(info.AbsOutputPath)
	return
}

// NewServer ...
func NewServer(template *template.Template, packageName string, filePath []string, imports imports, serverTpl string) *Server {
	return &Server{
		Template:    template,
		packageName: packageName,
		filePath:    filePath,
		imports:     imports,
		serverTpl:   serverTpl,
	}
}
