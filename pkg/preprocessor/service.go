package preprocessor

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
	"github.com/vetcher/go-astra"
	"github.com/vetcher/go-astra/types"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

const openAPIVersion = "3.0.0"

type swaggerRender interface {
	Generate(info api.GenerationInfo) (err error)
}

type servicesProcessor interface {
	Process(info *api.GenerationInfo, astra *types.File, outPath string) (err error)
}

// Service ...
type Service struct {
	servicesProcessor              servicesProcessor
	goGeneratedAutomaticallyPrefix []byte
	swaggerRender                  swaggerRender
}

func (s *Service) Process(serviceDirectory, outPath string, info *api.GenerationInfo) (err error) {
	var (
		filePath   string
		body       []byte
		files      []os.FileInfo
		serviceAst *types.File
	)

	if files, err = ioutil.ReadDir(serviceDirectory); err != nil {
		err = errors.Wrap(err, "ioutil.ReadDir error")
		return
	}

	for _, file := range files {
		filePath = path.Join(serviceDirectory, file.Name())
		if file.IsDir() {
			err = s.Process(filePath, path.Join(outPath, file.Name()), info)
			if err != nil {
				err = errors.Wrap(err, "s.Process error")
				return
			}
			continue
		}
		if !strings.HasSuffix(file.Name(), ".go") {
			continue
		}
		body, err = ioutil.ReadFile(filePath)
		if err != nil {
			err = errors.Wrap(err, "ioutil.ReadFile error")
			return
		}
		if bytes.Contains(body, s.goGeneratedAutomaticallyPrefix) {
			continue
		}
		if serviceAst, err = astra.ParseFile(filePath); err != nil {
			return
		}
		err = s.servicesProcessor.Process(info, serviceAst, outPath)
		if err != nil {
			err = errors.Wrap(err, "s.servicesProcessor.Process error")
			return
		}
	}

	if info.Swagger != nil {
		info.Swagger.OpenAPI = openAPIVersion
		err = s.swaggerRender.Generate(*info)
	}

	return
}

// NewService ...
func NewService(servicesProcessor servicesProcessor, goGeneratedAutomaticallyPrefix []byte, swaggerRender swaggerRender) *Service {
	return &Service{
		servicesProcessor:              servicesProcessor,
		goGeneratedAutomaticallyPrefix: goGeneratedAutomaticallyPrefix,
		swaggerRender:                  swaggerRender,
	}
}
