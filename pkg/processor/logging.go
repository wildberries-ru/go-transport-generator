package processor

import (
	"github.com/pkg/errors"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

type serviceRender interface {
	Generate(info api.Interface) (err error)
}

type logging struct {
	loggingRender serviceRender
}

func (s *logging) Process(_ *api.GenerationInfo, iface *api.Interface) (err error) {
	err = s.loggingRender.Generate(*iface)
	if err != nil {
		err = errors.Wrap(err, "[logging]s.loggingRender.Generate error")
	}
	return
}

// NewLogging ...
func NewLogging(loggingRender serviceRender) Processor {
	return &logging{
		loggingRender: loggingRender,
	}
}
