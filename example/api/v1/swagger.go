package v1

import (
	"encoding/json"
	"io/ioutil"
)

// Swagger ...
type Swagger struct {
	OpenAPI    string           `json:"openapi" yaml:"openapi"`
	Info       Info             `json:"info,omitempty" yaml:"info,omitempty"`
	Servers    []Server         `json:"servers,omitempty" yaml:"servers,omitempty"`
	Tags       []Tag            `json:"tags,omitempty" yaml:"tags,omitempty"`
	Schemes    []string         `json:"schemes,omitempty" yaml:"schemes,omitempty"`
	Paths      map[string]*Path `json:"paths" yaml:"paths"`
	Components Components       `json:"components,omitempty" yaml:"components,omitempty"`
}

// Contact ...
type Contact struct {
	Name  string `json:"name,omitempty" yaml:"name,omitempty"`
	URL   string `json:"url,omitempty" yaml:"url,omitempty"`
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
}

// License ...
type License struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	URL  string `json:"url,omitempty" yaml:"url,omitempty"`
}

// Info ...
type Info struct {
	Title          *string  `json:"title,omitempty" yaml:"title,omitempty"`
	Description    *string  `json:"description,omitempty" yaml:"description,omitempty"`
	TermsOfService *string  `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`
	Contact        *Contact `json:"contact,omitempty" yaml:"contact,omitempty"`
	License        *License `json:"license,omitempty" yaml:"license,omitempty"`
	Version        *string  `json:"version,omitempty" yaml:"version,omitempty"`
}

// ExternalDocs ...
type ExternalDocs struct {
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	URL         string `json:"url,omitempty" yaml:"url,omitempty"`
}

// Tag ...
type Tag struct {
	Name         string       `json:"name,omitempty" yaml:"name,omitempty"`
	Description  string       `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs ExternalDocs `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

// Properties ...
type Properties map[string]Schema

// Schema ...
type Schema struct {
	Ref         string      `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Type        string      `json:"type,omitempty" yaml:"type,omitempty"`
	Format      string      `json:"format,omitempty" yaml:"format,omitempty"`
	Properties  Properties  `json:"properties,omitempty" yaml:"properties,omitempty"`
	Items       *Schema     `json:"items,omitempty" yaml:"items,omitempty"`
	Enum        []string    `json:"enum,omitempty" yaml:"enum,omitempty"`
	Nullable    bool        `json:"nullable,omitempty" yaml:"nullable,omitempty"`
	Example     interface{} `json:"example,omitempty" yaml:"example,omitempty"`
	Description string      `json:"description,omitempty" yaml:"description,omitempty"`

	OneOf []Schema `json:"oneOf,omitempty" yaml:"oneOf,omitempty"`

	AdditionalProperties interface{} `json:"additionalProperties,omitempty" yaml:"additionalProperties,omitempty"`
}

// Parameter ...
type Parameter struct {
	Ref         string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	In          string `json:"in,omitempty" yaml:"in,omitempty"`
	Name        string `json:"name,omitempty" yaml:"name,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Required    bool   `json:"required,omitempty" yaml:"required,omitempty"`
	Schema      Schema `json:"schema,omitempty" yaml:"schema,omitempty"`
}

// Media ...
type Media struct {
	Schema Schema `json:"schema,omitempty" yaml:"schema,omitempty"`
}

// Content ...
type Content map[string]Media

// Response ...
type Response struct {
	Description string            `json:"description" yaml:"description"`
	Content     Content           `json:"content,omitempty" yaml:"content,omitempty"`
	Headers     map[string]Header `json:"headers,omitempty" yaml:"headers,omitempty"`
}

// Header ...
type Header struct {
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Schema      Schema `json:"schema,omitempty" yaml:"schema,omitempty"`
}

// Responses ...
type Responses map[string]Response

// RequestBody ...
type RequestBody struct {
	Description string  `json:"description,omitempty" yaml:"description,omitempty"`
	Content     Content `json:"content,omitempty" yaml:"content,omitempty"`
	Required    bool    `json:"required,omitempty" yaml:"required,omitempty"`
}

// Operation ...
type Operation struct {
	Tags        []string     `json:"tags,omitempty" yaml:"tags,omitempty"`
	Summary     *string      `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description *string      `json:"description,omitempty" yaml:"description,omitempty"`
	OperationID string       `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Consumes    []string     `json:"consumes,omitempty" yaml:"consumes,omitempty"`
	Produces    []string     `json:"produces,omitempty" yaml:"produces,omitempty"`
	Parameters  []Parameter  `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody *RequestBody `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Responses   Responses    `json:"responses,omitempty" yaml:"responses,omitempty"`
}

// Path ...
type Path struct {
	Ref         string     `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     string     `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string     `json:"description,omitempty" yaml:"description,omitempty"`
	Get         *Operation `json:"get,omitempty" yaml:"get,omitempty"`
	Post        *Operation `json:"post,omitempty" yaml:"post,omitempty"`
	Patch       *Operation `json:"patch,omitempty" yaml:"patch,omitempty"`
	Put         *Operation `json:"put,omitempty" yaml:"put,omitempty"`
	Delete      *Operation `json:"delete,omitempty" yaml:"delete,omitempty"`
}

// Variable ...
type Variable struct {
	Enum        []string `json:"enum,omitempty" yaml:"enum,omitempty"`
	Default     string   `json:"default,omitempty" yaml:"default,omitempty"`
	Description string   `json:"description,omitempty" yaml:"description,omitempty"`
}

// Server ...
type Server struct {
	URL         string              `json:"url,omitempty" yaml:"url,omitempty"`
	Description string              `json:"description,omitempty" yaml:"description,omitempty"`
	Variables   map[string]Variable `json:"variables,omitempty" yaml:"variables,omitempty"`
}

// Schemas ...
type Schemas map[string]Schema

// Components ...
type Components struct {
	Schemas Schemas `json:"schemas,omitempty" yaml:"schemas,omitempty"`
}

// SaveJSON ...
func (s *Swagger) SaveJSON(path string) (err error) {
	var data []byte
	if data, err = json.MarshalIndent(s, "", " "); err != nil {
		return
	}

	return ioutil.WriteFile(path, data, 0666)
}
