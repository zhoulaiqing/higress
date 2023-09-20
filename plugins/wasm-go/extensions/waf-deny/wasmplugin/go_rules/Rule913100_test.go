package go_rules

import (
	"github.com/stretchr/testify/assert"
	"github.com/tianchi/waf/wasmplugin/core"
	"testing"
)

func TestRule913100(t *testing.T) {
	rule := Rule913100{}
	tx := core.NewTransaction()

	tx.AddRequestHeader("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; SV1; .NET CLR 2.0.50727) Havij")

	_ = rule.Evaluate(&tx)
	assert.True(t, tx.Variables.InboundAnomalyScorePl1 > 0)
}
