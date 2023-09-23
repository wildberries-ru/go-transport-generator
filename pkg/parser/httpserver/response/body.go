package response

import (
	"errors"
	"strings"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

var (
	errHTTPBodyDidNotSet = "http body did not set"
)

type embed struct {
	prefix string
	suffix string
	next   Parser
}

func (t *embed) Parse(info *api.HTTPMethod, firstTag string, tags ...string) (err error) {
	if strings.HasPrefix(firstTag, t.prefix) && strings.HasSuffix(firstTag, t.suffix) {
		if len(tags) == 1 {
			info.ResponseBodyField = tags[0]
			return
		}
		return errors.New(errHTTPBodyDidNotSet)
	}
	return t.next.Parse(info, firstTag, tags...)
}

// NewBody ...
func NewBody(prefix string, suffix string, next Parser) Parser {
	return &embed{
		prefix: prefix,
		suffix: suffix,
		next:   next,
	}
}
