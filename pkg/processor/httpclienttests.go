package processor

import (
	"github.com/pkg/errors"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

type httpClientTests struct {
	testsRender httpRender
}

func (s *httpClientTests) Process(_ *api.GenerationInfo, iface *api.Interface) (err error) {
	err = s.testsRender.Generate(*iface)
	if err != nil {
		err = errors.Wrap(err, "[httpClient]s.testsRender.Generate error")
		return
	}
	return
}

// NewHTTPClientTests ...
func NewHTTPClientTests(
	testsRender httpRender,
) Processor {
	return &httpClientTests{
		testsRender: testsRender,
	}
}
