package request

import (
	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

// Term ...
type Term struct {
}

func (t *Term) Parse(info *api.HTTPMethod, firstTag string, tags ...string) (err error) {
	return
}
