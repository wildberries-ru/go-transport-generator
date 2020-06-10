package response

import (
	"errors"
	"strings"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

var (
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
			if info.ResponseJsonTags == nil {
				info.ResponseJsonTags = make(map[string]string)
			}
			info.ResponseJsonTags[tags[0]] = tags[1]
			return
		}
		return errors.New(errHTTPJsonTagDidNotSet)
	}
	return t.next.Parse(info, firstTag, tags...)
}

// NewJsonTag ...
func NewJsonTag(prefix string, suffix string, next Parser) Parser {
	return &jsonTag{
		prefix: prefix,
		suffix: suffix,
		next:   next,
	}
}
