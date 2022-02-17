package cor

import "strings"

type Handler interface {
	Handle(filename string) string
}

type JSONHandler struct {
	next Handler
}

func NewChain() Handler {
	return &JSONHandler{
		next: &INIHandler{
			&XMLHandler{
				next: &FallbackHandler{},
			},
		},
	}
}

func (h *JSONHandler) Handle(filename string) string {
	if !strings.Contains(filename, ".json") {
		return h.next.Handle(filename)
	}

	return "changed JSON file"
}

type XMLHandler struct {
	next Handler
}

func (h *XMLHandler) Handle(filename string) string {
	if !strings.Contains(filename, ".xml") {
		return h.next.Handle(filename)
	}

	return "changed XML file"
}

type INIHandler struct {
	next Handler
}

func (h *INIHandler) Handle(filename string) string {
	if !strings.Contains(filename, ".ini") {
		return h.next.Handle(filename)
	}

	return "changed INI file"
}

type FallbackHandler struct {
	next Handler
}

func (h *FallbackHandler) Handle(filename string) string {
	return "no file handler found!"
}
