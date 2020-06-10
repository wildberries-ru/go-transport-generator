package response

import (
	"errors"
	"strings"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

const (
	errHTTPContentEncodingDidNotSet = "http content encoding did not set"
)

type contentEncoding struct {
	prefix string
	suffix string
	next   Parser
}

func (t *contentEncoding) Parse(info *api.HTTPMethod, firstTag string, tags ...string) (err error) {
	if strings.HasPrefix(firstTag, t.prefix) && strings.HasSuffix(firstTag, t.suffix) {
		if len(tags) == 1 {
			info.ResponseContentEncoding = tags[0]
			return
		}
		return errors.New(errHTTPContentEncodingDidNotSet)
	}
	return t.next.Parse(info, firstTag, tags...)
}

// NewEncodingType ...
func NewEncodingType(prefix string, suffix string, next Parser) Parser {
	return &contentEncoding{
		prefix: prefix,
		suffix: suffix,
		next:   next,
	}
}
