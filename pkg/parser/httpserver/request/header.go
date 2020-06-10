package request

import (
	"errors"
	"strings"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

const (
	errHTTPHeaderDidNotSet = "http header did not set"
)

type header struct {
	prefix string
	suffix string
	next   Parser
}

func (t *header) Parse(info *api.HTTPMethod, firstTag string, tags ...string) (err error) {
	if strings.HasPrefix(firstTag, t.prefix) && strings.HasSuffix(firstTag, t.suffix) {
		if len(tags) == 2 {
			if info.HeaderPlaceholders == nil {
				info.HeaderPlaceholders = make(map[string]string)
			}
			info.HeaderPlaceholders[tags[0]] = strings.TrimRight(strings.TrimLeft(strings.TrimSpace(tags[1]), "{"), "}")
			return
		}
		return errors.New(errHTTPHeaderDidNotSet)
	}
	return t.next.Parse(info, firstTag, tags...)
}

// NewHeader ...
func NewHeader(prefix string, suffix string, next Parser) Parser {
	return &header{
		prefix: prefix,
		suffix: suffix,
		next:   next,
	}
}
