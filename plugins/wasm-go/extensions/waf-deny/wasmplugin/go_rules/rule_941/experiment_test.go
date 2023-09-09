package rule_941

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExperiment(t *testing.T) {
	assert.True(t, matchExp([]byte("<x s ")))
}
