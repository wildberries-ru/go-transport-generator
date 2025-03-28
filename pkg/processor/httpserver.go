package processor

import (
	"github.com/pkg/errors"
	"reflect"

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
	if s.builderRender != nil && reflect.ValueOf(s.builderRender).Kind() == reflect.Ptr && !reflect.ValueOf(s.builderRender).IsNil() {
		err = s.builderRender.Generate(*iface)
		if err != nil {
			err = errors.Wrap(err, "[httpServer]s.builderRender.Generate error")
			return
		}
	}
	if s.serverRender != nil && reflect.ValueOf(s.serverRender).Kind() == reflect.Ptr && !reflect.ValueOf(s.serverRender).IsNil() {
		err = s.serverRender.Generate(*iface)
		if err != nil {
			err = errors.Wrap(err, "[httpServer]s.serverRender.Generate error")
			return
		}
	}
	if s.transportRender != nil && reflect.ValueOf(s.transportRender).Kind() == reflect.Ptr && !reflect.ValueOf(s.transportRender).IsNil() {
		err = s.transportRender.Generate(*iface)
		if err != nil {
			err = errors.Wrap(err, "[httpServer]s.transportRender.Generate error")
			return
		}
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
