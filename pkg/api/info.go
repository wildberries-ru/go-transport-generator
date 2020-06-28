package api

import (
	"github.com/vetcher/go-astra/types"

	v1 "github.com/wildberries-ru/go-transport-generator/example/api/v1"
)

// GenerationInfo present generation info
type GenerationInfo struct {
	Interfaces []*Interface

	SwaggerInfo
	Swagger              *v1.Swagger
	SwaggerToJson        *bool
	SwaggerToYaml        *bool
	SwaggerAbsOutputPath string
}

// Interface ...
type Interface struct {
	PkgName       string
	AbsOutputPath string
	RelOutputPath string
	Iface         types.Interface
	IsTLSClient   bool
	HTTPMethods   map[string]HTTPMethod

	SwaggerInfo
}

// HTTPMethod ...
type HTTPMethod struct {
	Method                  string
	APIPath                 string
	RawURIPath              string
	URIPath                 string
	ClientURIPath           string
	ErrorProcessor          string
	URIPathPlaceholders     []string
	QueryPlaceholders       map[string]*Placeholder
	IsIntQueryPlaceholders  bool
	HeaderPlaceholders      map[string]string
	ContentType             string
	JsonTags                map[string]string
	MultipartValueTags      map[string]string
	MultipartFileTags       map[string]string
	Body                    map[string]string
	BodyPlaceholders        map[string]*Placeholder
	IsIntBodyPlaceholders   bool
	ResponseHeaders         map[string]string
	ResponseStatus          string
	ResponseContentType     string
	ResponseContentEncoding string
	ResponseJsonTags        map[string]string
	ResponseBody            map[string]string
	ResponseFile            string
	ResponseFileName        string
	ResponseBodyField       string
	LogIgnores              []string

	SwaggerInfo
}

// Placeholder ...
type Placeholder struct {
	IsPointer bool
	IsString  bool
	IsInt     bool
	Type      string
	Name      string
}

// SwaggerInfo ...
type SwaggerInfo struct {
	Description *string
	Summary     *string
	Title       *string
	Version     *string
	Servers     []v1.Server
}
