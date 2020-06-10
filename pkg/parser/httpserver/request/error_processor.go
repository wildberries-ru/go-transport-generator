package request

import (
	"errors"
	"strings"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

const (
	errHTTPErrorProcessorDidNotSet = "http errorProcessor did not set"
)

type errorProcessor struct {
	prefix string
	suffix string
	next   Parser
}

func (t *errorProcessor) Parse(info *api.HTTPMethod, firstTag string, tags ...string) (err error) {
	if strings.HasPrefix(firstTag, t.prefix) && strings.HasSuffix(firstTag, t.suffix) {
		if len(tags) == 1 {
			info.ErrorProcessor = tags[0]
			return
		}
		return errors.New(errHTTPErrorProcessorDidNotSet)
	}
	return t.next.Parse(info, firstTag, tags...)
}

// NewErrorProcessor ...
func NewErrorProcessor(prefix string, suffix string, next Parser) Parser {
	return &errorProcessor{
		prefix: prefix,
		suffix: suffix,
		next:   next,
	}
}
