package httpserver

import (
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/pkg/errors"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

// Transport ...
type Transport struct {
	*template.Template
	packageName        string
	filePath           []string
	imports            imports
	serverTransportTpl string
}

// Generate ...
func (s *Transport) Generate(info api.Interface) (err error) {
	info.PkgName = s.packageName
	info.AbsOutputPath = strings.Join(append(strings.Split(info.AbsOutputPath, "/"), s.filePath...), "/")
	dir, _ := path.Split(info.AbsOutputPath)
	err = os.MkdirAll(dir, 0750)
	if err != nil {
		err = errors.Wrap(err, "[httpserver.Transport]os.MkdirAll error")
		return
	}
	serverFile, err := os.Create(info.AbsOutputPath)
	defer func() {
		_ = serverFile.Close()
	}()
	t := template.Must(s.Parse(s.serverTransportTpl))
	if err = t.Execute(serverFile, info); err != nil {
		err = errors.Wrap(err, "[httpserver.Transport]t.Execute error")
		return
	}
	err = s.imports.GoImports(info.AbsOutputPath)

	return
}

// NewTransport ...
func NewTransport(template *template.Template, packageName string, filePath []string, imports imports, serverTransportTpl string) *Transport {
	return &Transport{
		Template:           template,
		packageName:        packageName,
		filePath:           filePath,
		imports:            imports,
		serverTransportTpl: serverTransportTpl,
	}
}
