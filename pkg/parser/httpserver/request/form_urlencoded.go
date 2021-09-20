package request

import (
	"errors"
	"strings"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

const (
	errHTTPFormUrlencodedTagDidNotSet = "http xxx-form-urlencoded value tag did not set"
)

type formUrlencodedTag struct {
	prefix string
	suffix string
	next   Parser
}

func (t *formUrlencodedTag) Parse(info *api.HTTPMethod, firstTag string, tags ...string) (err error) {
	if strings.HasPrefix(firstTag, t.prefix) && strings.HasSuffix(firstTag, t.suffix) {
		if len(tags) == 2 {
			if info.FormUrlencodedTags == nil {
				info.FormUrlencodedTags = make(map[string]string)
			}
			info.FormUrlencodedTags[tags[0]] = tags[1]
			return
		}
		return errors.New(errHTTPFormUrlencodedTagDidNotSet)
	}
	return t.next.Parse(info, firstTag, tags...)
}

// NewFormUrlencodedTag ...
func NewFormUrlencodedTag(prefix string, suffix string, next Parser) Parser {
	return &formUrlencodedTag{
		prefix: prefix,
		suffix: suffix,
		next:   next,
	}
}
