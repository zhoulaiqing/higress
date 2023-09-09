package rule_941

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExperiment(t *testing.T) {
	//var builder strings.Builder
	//builder.Reset()
	assert.True(t, matchExp([]byte("/char_test?mime=text/xml&body=<x onend=")))
	assert.True(t, matchExp([]byte("xyz=< a b c /> < s c r i p t >alert(1);</script>")))
}
