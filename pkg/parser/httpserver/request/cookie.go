package request

import (
	"errors"
	"strings"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

const (
	errHTTPCookieDidNotSet = "http cookie did not set"
)

type cookie struct {
	prefix string
	suffix string
	next   Parser
}

func (t *cookie) Parse(info *api.HTTPMethod, firstTag string, tags ...string) (err error) {
	if strings.HasPrefix(firstTag, t.prefix) && strings.HasSuffix(firstTag, t.suffix) {
		if len(tags) == 2 {
			if info.CookiePlaceholders == nil {
				info.CookiePlaceholders = make(map[string]string)
			}
			info.CookiePlaceholders[tags[0]] = strings.TrimRight(strings.TrimLeft(strings.TrimSpace(tags[1]), "{"), "}")
			return
		}
		return errors.New(errHTTPCookieDidNotSet)
	}
	return t.next.Parse(info, firstTag, tags...)
}

// NewCookie ...
func NewCookie(prefix string, suffix string, next Parser) Parser {
	return &cookie{
		prefix: prefix,
		suffix: suffix,
		next:   next,
	}
}
