package swagger

import (
	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

type Term struct {
}

func (t *Term) Parse(info *api.SwaggerInfo, firstTag string, tags ...string) (err error) {
	return
}
