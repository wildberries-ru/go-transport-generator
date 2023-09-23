package swagger

import (
	"errors"
	"strings"

	v1 "github.com/wildberries-ru/go-transport-generator/example/api/v1"
	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

const (
	emptyString = ""
	spaceString = " "
)

var (
	errSwaggerServersDidNotSet = errors.New("swagger servers did not set")
)

type servers struct {
	prefix string
	suffix string
	next   Parser
}

func (t *servers) Parse(info *api.SwaggerInfo, firstTag string, tags ...string) (err error) {
	if strings.HasPrefix(firstTag, t.prefix) && strings.HasSuffix(firstTag, t.suffix) {
		if len(tags) == 1 {
			info.Servers = append(info.Servers, v1.Server{
				URL:         tags[0],
				Description: emptyString,
			})
			return
		}
		if len(tags) > 1 {
			info.Servers = append(info.Servers, v1.Server{
				URL:         tags[0],
				Description: strings.Join(tags[1:], spaceString),
			})
			return
		}
		return errSwaggerServersDidNotSet
	}
	return t.next.Parse(info, firstTag, tags...)
}

// NewServers ...
func NewServers(prefix string, suffix string, next Parser) Parser {
	return &servers{
		prefix: prefix,
		suffix: suffix,
		next:   next,
	}
}
