package request

import (
	"errors"
	"strings"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

const (
	errHTTPMultipartFileTagDidNotSet = "http multipart file tag did not set"
)

type multipartFileTag struct {
	prefix string
	suffix string
	next   Parser
}

func (t *multipartFileTag) Parse(info *api.HTTPMethod, firstTag string, tags ...string) (err error) {
	if strings.HasPrefix(firstTag, t.prefix) && strings.HasSuffix(firstTag, t.suffix) {
		if len(tags) == 2 {
			if info.MultipartFileTags == nil {
				info.MultipartFileTags = make(map[string]string)
			}
			info.MultipartFileTags[tags[0]] = tags[1]
			return
		}
		return errors.New(errHTTPMultipartFileTagDidNotSet)
	}
	return t.next.Parse(info, firstTag, tags...)
}

// NewMultipartFileTag ...
func NewMultipartFileTag(prefix string, suffix string, next Parser) Parser {
	return &multipartFileTag{
		prefix: prefix,
		suffix: suffix,
		next:   next,
	}
}
