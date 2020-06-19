package processor

import (
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/vetcher/go-astra/types"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

// Services ...
type Services struct {
	tagMark    string
	processors map[string]Processor
}

// Process ...
func (s *Services) Process(info *api.GenerationInfo, astra *types.File, outPath string) (err error) {
	for _, i := range astra.Interfaces {
		for _, doc := range i.Docs {
			doc = strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(doc), "//"))
			if strings.HasPrefix(doc, s.tagMark) {
				iface := api.Interface{
					Iface: i,
				}
				iface.RelOutputPath = outPath
				if iface.AbsOutputPath, err = filepath.Abs(outPath); err != nil {
					err = errors.Wrap(err, "filepath.Abs error")
					return
				}
				info.Interfaces = append(info.Interfaces, &iface)
				words := strings.Split(strings.TrimSpace(doc[len(s.tagMark):]), " ")
				for _, word := range words {
					if processor, ok := s.processors[word]; ok {
						err = processor.Process(info, &iface)
						if err != nil {
							err = errors.Wrap(err, "processor.Process error")
							return
						}
					}
				}
			}
		}
	}
	return
}

// NewServices ...
func NewServices(tagMark string, processors map[string]Processor) *Services {
	return &Services{
		tagMark:    tagMark,
		processors: processors,
	}
}
