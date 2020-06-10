package processor

import (
	"github.com/pkg/errors"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

type Errors struct {
	tagMark            string
	uiErrorsRender     serviceRender
	clientErrorsRender serviceRender
}

func (s *Errors) Process(_ *api.GenerationInfo, iface *api.Interface) (err error) {
	err = s.uiErrorsRender.Generate(*iface)
	if err != nil {
		err = errors.Wrap(err, "[Errors]s.uiErrorsRender.Generate error")
		return
	}
	err = s.clientErrorsRender.Generate(*iface)
	if err != nil {
		err = errors.Wrap(err, "[Errors]s.clientErrorsRender.Generate error")
	}
	return
}

// NewErrors ...
func NewErrors(
	tagMark string,
	uiErrorsRender serviceRender,
	clientErrorsRender serviceRender,
) Processor {
	return &Errors{
		tagMark:            tagMark,
		uiErrorsRender:     uiErrorsRender,
		clientErrorsRender: clientErrorsRender,
	}
}
