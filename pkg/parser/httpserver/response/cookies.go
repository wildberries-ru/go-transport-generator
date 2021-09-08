package response

import (
	"errors"
	"strings"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

var (
	errHTTPCookiesDidNotSet = "http cookies did not set"
)

type cookies struct {
	prefix string
	suffix string
	next   Parser
}

func (t *cookies) Parse(info *api.HTTPMethod, firstTag string, tags ...string) (err error) {
	if strings.HasPrefix(firstTag, t.prefix) && strings.HasSuffix(firstTag, t.suffix) {
		if len(tags) == 2 {
			if info.ResponseCookies == nil {
				info.ResponseCookies = make(map[string]string)
			}
			info.ResponseCookies[tags[0]] = strings.TrimRight(strings.TrimLeft(strings.TrimSpace(tags[1]), "{"), "}")
			return
		}
		return errors.New(errHTTPCookiesDidNotSet)
	}
	return t.next.Parse(info, firstTag, tags...)
}

// NewCookies ...
func NewCookies(prefix string, suffix string, next Parser) Parser {
	return &cookies{
		prefix: prefix,
		suffix: suffix,
		next:   next,
	}
}
