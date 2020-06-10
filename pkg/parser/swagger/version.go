package swagger

import (
	"errors"
	"strings"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

var (
	errSwaggerVersionDidNotSet = errors.New("swagger version did not set")
)

type version struct {
	prefix string
	suffix string
	next   Parser
}

func (t *version) Parse(info *api.SwaggerInfo, firstTag string, tags ...string) (err error) {
	if strings.HasPrefix(firstTag, t.prefix) && strings.HasSuffix(firstTag, t.suffix) {
		if len(tags) == 1 {
			info.Version = &tags[0]
			return
		}
		return errSwaggerVersionDidNotSet
	}
	return t.next.Parse(info, firstTag, tags...)
}

// NewVersion ...
func NewVersion(prefix string, suffix string, next Parser) Parser {
	return &version{
		prefix: prefix,
		suffix: suffix,
		next:   next,
	}
}
