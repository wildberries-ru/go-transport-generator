package request

import (
	"errors"
	"strings"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

const (
	errHTTPJsonTagDidNotSet = "http json tag did not set"
)

type jsonTag struct {
	prefix string
	suffix string
	next   Parser
}

func (t *jsonTag) Parse(info *api.HTTPMethod, firstTag string, tags ...string) (err error) {
	if strings.HasPrefix(firstTag, t.prefix) && strings.HasSuffix(firstTag, t.suffix) {
		if len(tags) == 2 {
			if info.PlainObject != "" {
				return errors.New(errJSONTagPlainObjectIncompatible)
			}
			if info.JSONTags == nil {
				info.JSONTags = make(map[string]string)
			}
			info.JSONTags[tags[0]] = tags[1]
			return
		}
		return errors.New(errHTTPJsonTagDidNotSet)
	}
	return t.next.Parse(info, firstTag, tags...)
}

// NewJSONTag ...
func NewJSONTag(prefix string, suffix string, next Parser) Parser {
	return &jsonTag{
		prefix: prefix,
		suffix: suffix,
		next:   next,
	}
}
