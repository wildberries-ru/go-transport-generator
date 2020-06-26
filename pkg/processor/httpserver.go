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
	serverRender    httpRender
	transportRender httpRender
	builderRender   httpRender
}

func (s *httpServer) Process(_ *api.GenerationInfo, iface *api.Interface) (err error) {
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
func NewHTTPServer(serverRender httpRender, transportRender httpRender, builderRender httpRender) Processor {
	return &httpServer{
		serverRender:    serverRender,
		transportRender: transportRender,
		builderRender:   builderRender,
	}
}
