package processor

import (
	"github.com/pkg/errors"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

type httpClient struct {
	isTLS               bool
	clientRender        httpRender
	transportRender     httpRender
	builderRender       httpRender
	httpMethodProcessor HTTPMethod
}

func (s *httpClient) Process(_ *api.GenerationInfo, iface *api.Interface) (err error) {
	iface.IsTLSClient = s.isTLS
	iface.HTTPMethods = make(map[string]api.HTTPMethod)
	for _, method := range iface.Iface.Methods {
		httpMethod := api.HTTPMethod{}
		err = s.httpMethodProcessor.Process(&httpMethod, iface, method)
		if err != nil {
			err = errors.Wrap(err, "[httpClient]s.httpMethodProcessor.Process error")
			return
		}
		iface.HTTPMethods[method.Name] = httpMethod
	}
	err = s.builderRender.Generate(*iface)
	if err != nil {
		err = errors.Wrap(err, "[httpClient]s.builderRender.Generate error")
		return
	}
	err = s.clientRender.Generate(*iface)
	if err != nil {
		err = errors.Wrap(err, "[httpClient]s.clientRender.Generate error")
		return
	}
	err = s.transportRender.Generate(*iface)
	if err != nil {
		err = errors.Wrap(err, "[httpClient]s.transportRender.Generate error")
		return
	}
	return
}

// NewHTTPClient ...
func NewHTTPClient(
	isTLS bool,
	clientRender httpRender,
	transportRender httpRender,
	builderRender httpRender,
	httpMethodProcessor HTTPMethod,
) Processor {
	return &httpClient{
		isTLS:               isTLS,
		clientRender:        clientRender,
		transportRender:     transportRender,
		builderRender:       builderRender,
		httpMethodProcessor: httpMethodProcessor,
	}
}
