package metrics

import (
	"errors"
	"strings"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

const (
	errAdditionalMetricsLabelsDidNotSet = "additional metrics labels did not set"
	logIgnoreSeparator                  = ","
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
			info.AdditionalMetricsLabels = make(map[string]*api.MetricsPlaceholder, len(s))
			for _, v := range s {
				info.AdditionalMetricsLabels[v] = &api.MetricsPlaceholder{
					Name: strings.TrimSpace(v),
				}
			}
			return
		}
		return errors.New(errAdditionalMetricsLabelsDidNotSet)
	}
	return t.next.Parse(info, firstTag, tags...)
}

// NewAdditionalMetricsLabels ...
func NewAdditionalMetricsLabels(prefix string, suffix string, next Parser) Parser {
	return &logIgnore{
		prefix: prefix,
		suffix: suffix,
		next:   next,
	}
}
