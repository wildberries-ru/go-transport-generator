package swagger

import (
	"errors"
	"strings"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

var (
	errSwaggerTitleDidNotSet = errors.New("swagger title did not set")
)

// Parser ...
type Parser interface {
	Parse(info *api.SwaggerInfo, firstTag string, tags ...string) (err error)
}

type title struct {
	prefix string
	suffix string
	next   Parser
}

func (t *title) Parse(info *api.SwaggerInfo, firstTag string, tags ...string) (err error) {
	if strings.HasPrefix(firstTag, t.prefix) && strings.HasSuffix(firstTag, t.suffix) {
		if len(tags) > 0 {
			var b strings.Builder
			b.Grow(len(tags))
			for _, tag := range tags {
				b.WriteString(tag)
				b.WriteString(" ")
			}
			title := b.String()
			info.Title = &title
			return
		}
		return errSwaggerTitleDidNotSet
	}
	return t.next.Parse(info, firstTag, tags...)
}

// NewTitle ...
func NewTitle(prefix string, suffix string, next Parser) Parser {
	return &title{
		prefix: prefix,
		suffix: suffix,
		next:   next,
	}
}
