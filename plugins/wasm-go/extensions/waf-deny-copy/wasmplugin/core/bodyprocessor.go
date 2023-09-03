package core

import (
	"fmt"
	"io"
	"strings"
)

type BodyProcessor interface {
	ProcessRequest(reader io.Reader, tx *Transaction, mime string) error
	ProcessResponse(reader io.Reader, tx *Transaction, mime string) error
}

type bodyProcessorWrapper = func() BodyProcessor

var processors = map[string]bodyProcessorWrapper{}

// RegisterBodyProcessor registers a body processor
// by name. If the body processor is already registered,
// it will be overwritten
func RegisterBodyProcessor(name string, fn func() BodyProcessor) {
	processors[name] = fn
}

// GetBodyProcessor returns a body processor by name
// If the body processor is not found, it returns an error
func GetBodyProcessor(name string) (BodyProcessor, error) {
	if fn, ok := processors[strings.ToLower(name)]; ok {
		return fn(), nil
	}
	return nil, fmt.Errorf("invalid bodyprocessor %q", name)
}