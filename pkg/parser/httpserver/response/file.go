package response

import (
	"errors"
	"strings"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

var (
	errHTTPByteDataSet = "http byte data did not set"
)

type byteData struct {
	prefix string
	suffix string
	next   Parser
}

func (t *byteData) Parse(info *api.HTTPMethod, firstTag string, tags ...string) (err error) {
	if strings.HasPrefix(firstTag, t.prefix) && strings.HasSuffix(firstTag, t.suffix) {
		if len(tags) == 1 {
			info.ResponseFile = tags[0]
			return
		}
		if len(tags) == 2 {
			info.ResponseFile = tags[0]
			info.ResponseFileName = tags[1]
			return
		}
		return errors.New(errHTTPByteDataSet)
	}
	return t.next.Parse(info, firstTag, tags...)
}

// NewFile ...
func NewFile(prefix string, suffix string, next Parser) Parser {
	return &byteData{
		prefix: prefix,
		suffix: suffix,
		next:   next,
	}
}
