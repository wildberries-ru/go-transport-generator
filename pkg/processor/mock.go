package processor

import (
	"github.com/pkg/errors"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

type mock struct {
	mockRender serviceRender
}

func (s *mock) Process(_ *api.GenerationInfo, iface *api.Interface) (err error) {
	err = s.mockRender.Generate(*iface)
	if err != nil {
		err = errors.Wrap(err, "[mock]s.mockRender.Generate error")
	}
	return
}

// NewMock ...
func NewMock(mockRender serviceRender) Processor {
	return &mock{
		mockRender: mockRender,
	}
}
