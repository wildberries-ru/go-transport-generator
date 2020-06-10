package processor

import (
	"github.com/pkg/errors"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

type httpRender interface {
	Generate(info api.Interface) (err error)
}

type httpServerTagsParser interface {
	Parse(info *api.HTTPMethod, firstTag string, tags ...string) (err error)
}

// Processor ...
type Processor interface {
	Process(info *api.GenerationInfo, iface *api.Interface) (err error)
}

type httpServer struct {
	serverRender        httpRender
	transportRender     httpRender
	builderRender       httpRender
	httpMethodProcessor HTTPMethod
}

func (s *httpServer) Process(_ *api.GenerationInfo, iface *api.Interface) (err error) {
	iface.HTTPMethods = make(map[string]api.HTTPMethod)
	for _, method := range iface.Iface.Methods {
		httpMethod := api.HTTPMethod{}
		err = s.httpMethodProcessor.Process(&httpMethod, iface, method)
		if err != nil {
			err = errors.Wrap(err, "[httpServer]s.httpMethodProcessor.Process error")
			return
		}
		iface.HTTPMethods[method.Name] = httpMethod
	}
	err = s.builderRender.Generate(*iface)
	if err != nil {
		err = errors.Wrap(err, "[httpServer]s.builderRender.Generate error")
		return
	}
	err = s.serverRender.Generate(*iface)
	if err != nil {
		err = errors.Wrap(err, "[httpServer]s.serverRender.Generate error")
		return
	}
	err = s.transportRender.Generate(*iface)
	if err != nil {
		err = errors.Wrap(err, "[httpServer]s.transportRender.Generate error")
		return
	}
	return
}

// NewHTTPServer ...
func NewHTTPServer(serverRender httpRender, transportRender httpRender, builderRender httpRender, httpMethodProcessor HTTPMethod) Processor {
	return &httpServer{
		serverRender:        serverRender,
		transportRender:     transportRender,
		builderRender:       builderRender,
		httpMethodProcessor: httpMethodProcessor,
	}
}
