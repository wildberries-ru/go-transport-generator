package log

import (
	"errors"
	"strings"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

const (
	errLogIgnoreFieldsDidNotSet = "log ignore fields did not set"
	logIgnoreSeparator          = ","
)

// Parser ...
type Parser interface {
	Parse(info *api.HTTPMethod, firstTag string, tags ...string) (err error)
}

type logIgnore struct {
	prefix string
	suffix string
	next   Parser
}

func (t *logIgnore) Parse(info *api.HTTPMethod, firstTag string, tags ...string) (err error) {
	if strings.HasPrefix(firstTag, t.prefix) && strings.HasSuffix(firstTag, t.suffix) {
		if len(tags) > 0 {
			s := strings.Split(strings.Join(tags, " "), logIgnoreSeparator)
			info.LogIgnores = make([]string, len(s))
			for k, v := range s {
				info.LogIgnores[k] = strings.TrimSpace(v)
			}
			return
		}
		return errors.New(errLogIgnoreFieldsDidNotSet)
	}
	return t.next.Parse(info, firstTag, tags...)
}

// NewLogIgnore ...
func NewLogIgnore(prefix string, suffix string, next Parser) Parser {
	return &logIgnore{
		prefix: prefix,
		suffix: suffix,
		next:   next,
	}
}
