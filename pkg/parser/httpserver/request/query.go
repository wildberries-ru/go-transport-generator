package request

import (
	"errors"
	"strings"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

const (
	errHTTPQueryDidNotSet = "http query did not set"
)

type query struct {
	prefix string
	suffix string
	next   Parser
}

func (t *query) Parse(info *api.HTTPMethod, firstTag string, tags ...string) (err error) {
	if strings.HasPrefix(firstTag, t.prefix) && strings.HasSuffix(firstTag, t.suffix) {
		if len(tags) == 1 {
			qs := strings.Split(tags[0], "&")
			if len(qs) > 0 {
				info.RawQueryPlaceholders = make(map[string]*api.Placeholder, len(qs))
				for _, q := range qs {
					qq := strings.Split(q, "=")
					if len(qq) != 2 {
						continue
					}
					info.RawQueryPlaceholders[qq[0]] = &api.Placeholder{
						Name: strings.TrimRight(strings.TrimLeft(qq[1], "{"), "}"),
					}
				}
			}
			return
		}
		return errors.New(errHTTPQueryDidNotSet)
	}
	return t.next.Parse(info, firstTag, tags...)
}

// NewQuery ...
func NewQuery(prefix string, suffix string, next Parser) Parser {
	return &query{
		prefix: prefix,
		suffix: suffix,
		next:   next,
	}
}
