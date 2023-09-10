package rule_941

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestXss(t *testing.T) {
	//assert.True(t, matchXSS([]byte("/char_test?mime=text/xml&body=<x onend=")))
	matches := rule_tasks.Re941160.FindStringSubmatch("/char_test?mime=text/xml&body=$<f o r m ")
	assert.True(t, len(matches) > 0)
}
