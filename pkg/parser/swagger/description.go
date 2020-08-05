package swagger

import (
	"errors"
	"strings"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

var (
	errSwaggerDescriptionDidNotSet = errors.New("swagger description did not set")
)

type description struct {
	prefix string
	suffix string
	next   Parser
}

func (t *description) Parse(info *api.SwaggerInfo, firstTag string, tags ...string) (err error) {
	if strings.HasPrefix(firstTag, t.prefix) && strings.HasSuffix(firstTag, t.suffix) {
		if len(tags) > 0 {
			var b strings.Builder
			b.Grow(len(tags))
			for _, tag := range tags {
				b.WriteString(tag)
				b.WriteString(" ")
			}
			desc := b.String()
			info.Description = &desc
			return
		}
		return errSwaggerDescriptionDidNotSet
	}
	return t.next.Parse(info, firstTag, tags...)
}

// NewDescription ...
func NewDescription(prefix string, suffix string, next Parser) Parser {
	return &description{
		prefix: prefix,
		suffix: suffix,
		next:   next,
	}
}
