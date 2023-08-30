package bodyprocessors

import (
	"fmt"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"strings"
)

type bodyProcessorWrapper = func() core.BodyProcessor

var processors = map[string]bodyProcessorWrapper{}

// RegisterBodyProcessor registers a body processor
// by name. If the body processor is already registered,
// it will be overwritten
func RegisterBodyProcessor(name string, fn func() core.BodyProcessor) {
	processors[name] = fn
}

// GetBodyProcessor returns a body processor by name
// If the body processor is not found, it returns an error
func GetBodyProcessor(name string) (core.BodyProcessor, error) {
	if fn, ok := processors[strings.ToLower(name)]; ok {
		return fn(), nil
	}
	return nil, fmt.Errorf("invalid bodyprocessor %q", name)
}
