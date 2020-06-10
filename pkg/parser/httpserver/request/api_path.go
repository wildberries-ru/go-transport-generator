package request

import (
	"errors"
	"strings"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

const (
	errHTTPAPIPathDidNotSet = "http api path did not set"
)

type apiPath struct {
	prefix string
	suffix string
	next   Parser
}

func (t *apiPath) Parse(info *api.HTTPMethod, firstTag string, tags ...string) (err error) {
	if strings.HasPrefix(firstTag, t.prefix) && strings.HasSuffix(firstTag, t.suffix) {
		if len(tags) == 1 {
			info.APIPath = tags[0]
			return
		}
		return errors.New(errHTTPAPIPathDidNotSet)
	}
	return t.next.Parse(info, firstTag, tags...)
}

// NewAPIPath ...
func NewAPIPath(prefix string, suffix string, next Parser) Parser {
	return &apiPath{
		prefix: prefix,
		suffix: suffix,
		next:   next,
	}
}
