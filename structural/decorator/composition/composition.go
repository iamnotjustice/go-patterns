package composition

import "strings"

type Processor interface {
	Process([]string) string
}

type OriginalProcessor struct {
}

func (OriginalProcessor) Process(in []string) string {
	buf := "|"
	for _, v := range in {
		buf += v + "|"
	}

	return buf
}

type DecoratedProcessor struct {
	orig Processor
}

func NewDecoratedProcessor(o Processor) Processor {
	return &DecoratedProcessor{orig: o}
}

func (d *DecoratedProcessor) Process(in []string) string {
	res := d.orig.Process(in)

	return strings.ReplaceAll(res, "|", ":")
}
