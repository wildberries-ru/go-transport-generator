package request

import (
	"errors"
	"strings"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

const (
	errHTTPMethodDidNotSet = "http errorProcessor did not set"
)

// Parser ...
type Parser interface {
	Parse(info *api.HTTPMethod, firstTag string, tags ...string) (err error)
}

type method struct {
	prefix string
	suffix string
	next   Parser
}

func (t *method) Parse(info *api.HTTPMethod, firstTag string, tags ...string) (err error) {
	if strings.HasPrefix(firstTag, t.prefix) && strings.HasSuffix(firstTag, t.suffix) {
		if len(tags) == 1 {
			info.Method = tags[0]
			return
		}
		return errors.New(errHTTPMethodDidNotSet)
	}
	return t.next.Parse(info, firstTag, tags...)
}

// NewMethod ...
func NewMethod(prefix string, suffix string, next Parser) Parser {
	return &method{
		prefix: prefix,
		suffix: suffix,
		next:   next,
	}
}
