package processor

import (
	"github.com/pkg/errors"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

type httpClientBuilderTests struct {
	testsRender httpRender
}

func (s *httpClientBuilderTests) Process(_ *api.GenerationInfo, iface *api.Interface) (err error) {
	err = s.testsRender.Generate(*iface)
	if err != nil {
		err = errors.Wrap(err, "[httpClient]s.builderTestsRender.Generate error")
		return
	}
	return
}

// NewHTTPClientBuilderTests ...
func NewHTTPClientBuilderTests(
	testsRender httpRender,
) Processor {
	return &httpClientBuilderTests{
		testsRender: testsRender,
	}
}
