package request

import (
	"errors"
	"strings"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

const (
	errPlainObjectTagIncorrect        = "http plain object tag incorrect, should be one value with param name"
	errPlainObjectSingle              = "http plain object should be only one"
	errJSONTagPlainObjectIncompatible = "http json tag and plain object are incompatible, please use on of them exclusively"
)

type plainObjectTag struct {
	prefix string
	suffix string
	next   Parser
}

func (t *plainObjectTag) Parse(info *api.HTTPMethod, firstTag string, tags ...string) (err error) {
	if strings.HasPrefix(firstTag, t.prefix) && strings.HasSuffix(firstTag, t.suffix) {
		if len(tags) == 1 {
			if info.PlainObject != "" {
				return errors.New(errPlainObjectSingle)
			}
			if len(info.JSONTags) != 0 {
				return errors.New(errJSONTagPlainObjectIncompatible)
			}
			info.PlainObject = tags[0]
			return
		}
		return errors.New(errPlainObjectTagIncorrect)
	}
	return t.next.Parse(info, firstTag, tags...)
}

// NewPlainObjectTag ...
func NewPlainObjectTag(prefix string, suffix string, next Parser) Parser {
	return &plainObjectTag{
		prefix: prefix,
		suffix: suffix,
		next:   next,
	}
}
