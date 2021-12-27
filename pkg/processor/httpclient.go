package processor

import (
	"github.com/pkg/errors"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

type httpClient struct {
	isTLS           bool
	isInsecureTLS   bool
	clientRender    httpRender
	transportRender httpRender
	builderRender   httpRender
}

func (s *httpClient) Process(_ *api.GenerationInfo, iface *api.Interface) (err error) {
	iface.IsTLSClient = s.isTLS
	iface.IsInsecureTLS = s.isInsecureTLS
	err = s.builderRender.Generate(*iface)
	err = s.transportRender.Generate(*iface)
	if err != nil {
		err = errors.Wrap(err, "[httpClient]s.transportRender.Generate error")
		return
	}
	err = s.clientRender.Generate(*iface)
	if err != nil {
		err = errors.Wrap(err, "[httpClient]s.clientRender.Generate error")
		return
	}
	if err != nil {
		err = errors.Wrap(err, "[httpClient]s.builderRender.Generate error")
		return
	}
	return
}

// NewHTTPClient ...
func NewHTTPClient(
	isTLS bool,
	isInsecureTLS bool,
	clientRender httpRender,
	transportRender httpRender,
	builderRender httpRender,
) Processor {
	return &httpClient{
		isTLS:           isTLS,
		isInsecureTLS:   isInsecureTLS,
		clientRender:    clientRender,
		transportRender: transportRender,
		builderRender:   builderRender,
	}
}
