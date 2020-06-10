package processor

import (
	"github.com/pkg/errors"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

type instrumenting struct {
	instrumentingRender serviceRender
}

func (s *instrumenting) Process(_ *api.GenerationInfo, iface *api.Interface) (err error) {
	err = s.instrumentingRender.Generate(*iface)
	if err != nil {
		err = errors.Wrap(err, "[instrumenting]s.instrumentingRender.Generate error")
	}
	return
}

// NewInstrumenting ...
func NewInstrumenting(instrumentingRender serviceRender) Processor {
	return &instrumenting{
		instrumentingRender: instrumentingRender,
	}
}
