package httpserver

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v2"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

const (
	jsonExt = ".json"
	yamlExt = ".yaml"
)

// Swagger ...
type Swagger struct {
	fileName string
}

// Generate ...
func (s *Swagger) Generate(info api.GenerationInfo) (err error) {
	var data []byte
	err = os.MkdirAll(info.SwaggerAbsOutputPath, 0750)
	if err != nil {
		return
	}
	fileName := path.Join(info.SwaggerAbsOutputPath, s.fileName)
	if info.SwaggerToJSON != nil && *info.SwaggerToJSON {
		fileName += jsonExt
		if data, err = json.Marshal(info.Swagger); err != nil {
			return
		}
	}
	if info.SwaggerToYaml != nil && *info.SwaggerToYaml {
		fileName += yamlExt
		if data, err = yaml.Marshal(info.Swagger); err != nil {
			return
		}
	}
	return ioutil.WriteFile(fileName, data, 0750)
}

// NewSwagger ...
func NewSwagger(fileName string) *Swagger {
	return &Swagger{
		fileName: fileName,
	}
}
