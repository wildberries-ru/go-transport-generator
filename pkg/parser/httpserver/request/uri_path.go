package request

import (
	"errors"
	"strings"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

const (
	errHTTPURIPathDidNotSet = "http uri path did not set"
)

type uriPath struct {
	prefix string
	suffix string
	next   Parser
}

func (t *uriPath) Parse(info *api.HTTPMethod, firstTag string, tags ...string) (err error) {
	if strings.HasPrefix(firstTag, t.prefix) && strings.HasSuffix(firstTag, t.suffix) {
		if len(tags) == 1 {
			p := strings.Split(tags[0], "{")
			if len(p) > 0 {
				pn := make([]string, len(p))
				pc := make([]string, len(p))
				pn[0] = p[0]
				pc[0] = p[0]
				for i := 1; i < len(p); i++ {
					ph := strings.Split(p[i], "}")
					if len(ph) > 0 {
						info.URIPathPlaceholders = append(info.URIPathPlaceholders, ph[0])
					}
					pn[i] = strings.Join(ph, "")
					ph[0] = "s"
					pc[i] = strings.Join(ph, "")
				}
				info.URIPath = strings.Join(pn, ":")
				info.ClientURIPath = strings.Join(pc, "%")
			}
			info.RawURIPath = tags[0]
			return
		}
		return errors.New(errHTTPURIPathDidNotSet)
	}
	return t.next.Parse(info, firstTag, tags...)
}

// NewURIPath ...
func NewURIPath(prefix string, suffix string, next Parser) Parser {
	return &uriPath{
		prefix: prefix,
		suffix: suffix,
		next:   next,
	}
}
