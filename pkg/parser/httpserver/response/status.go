package response

import (
	"errors"
	"strings"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

var (
	errHTTPResponseStatusDidNotSet = "http response status did not set"
)

type status struct {
	prefix string
	suffix string
	next   Parser
}

func (t *status) Parse(info *api.HTTPMethod, firstTag string, tags ...string) (err error) {
	if strings.HasPrefix(firstTag, t.prefix) && strings.HasSuffix(firstTag, t.suffix) {
		if len(tags) == 1 {
			info.ResponseStatus = tags[0]
			return
		}
		return errors.New(errHTTPResponseStatusDidNotSet)
	}
	return t.next.Parse(info, firstTag, tags...)
}

// NewStatus ...
func NewStatus(prefix string, suffix string, next Parser) Parser {
	return &status{
		prefix: prefix,
		suffix: suffix,
		next:   next,
	}
}
