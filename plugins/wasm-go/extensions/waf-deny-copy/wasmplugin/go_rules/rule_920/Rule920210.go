package rule_920

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"github.com/wasilibs/go-re2"
)

const PTN_920210 = `\b(?:keep-alive|close),\s?(?:keep-alive|close)\b`

var re920210 *re2.Regexp

type Rule920210 struct {
}

func (r *Rule920210) Id() string {
	return "920210"
}

func (r *Rule920210) Phase() int {
	return 1
}

func (r *Rule920210) Evaluate(tx *core.Transaction) int {

	matched := re920210.MatchString(tx.Variables.RequestHeaders["connection"])
	if matched {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.WARNING_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}
	return rule_tasks.PASS
}

func init() {
	re920210, _ = re2.Compile(PTN_920210)
}
