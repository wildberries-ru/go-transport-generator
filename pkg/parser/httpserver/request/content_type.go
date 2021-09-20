package request

import (
	"errors"
	"strings"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

const (
	errHTTPContentTypeDidNotSet = "http content type did not set"
)

type contentType struct {
	prefix string
	suffix string
	next   Parser
}

func (t *contentType) Parse(info *api.HTTPMethod, firstTag string, tags ...string) (err error) {
	if strings.HasPrefix(firstTag, t.prefix) && strings.HasSuffix(firstTag, t.suffix) {
		if len(tags) == 1 {
			info.ContentType = tags[0]
			return
		}
		return errors.New(errHTTPContentTypeDidNotSet)
	}
	return t.next.Parse(info, firstTag, tags...)
}

// NewContentType ...
func NewContentType(prefix string, suffix string, next Parser) Parser {
	return &contentType{
		prefix: prefix,
		suffix: suffix,
		next:   next,
	}
}
