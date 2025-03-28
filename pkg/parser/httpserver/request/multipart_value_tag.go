package request

import (
	"errors"
	"strings"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

const (
	errHTTPMultipartValueTagDidNotSet = "http multipart value tag did not set"
)

type multipartValueTag struct {
	prefix string
	suffix string
	next   Parser
}

func (t *multipartValueTag) Parse(info *api.HTTPMethod, firstTag string, tags ...string) (err error) {
	if strings.HasPrefix(firstTag, t.prefix) && strings.HasSuffix(firstTag, t.suffix) {
		if len(tags) == 2 {
			if info.MultipartValueTags == nil {
				info.MultipartValueTags = make(map[string]string)
			}
			info.MultipartValueTags[tags[0]] = tags[1]
			return
		}
		return errors.New(errHTTPMultipartValueTagDidNotSet)
	}
	return t.next.Parse(info, firstTag, tags...)
}

// NewMultipartValueTag ...
func NewMultipartValueTag(prefix string, suffix string, next Parser) Parser {
	return &multipartValueTag{
		prefix: prefix,
		suffix: suffix,
		next:   next,
	}
}
