package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"github.com/wasilibs/go-re2"
)

const PTN_921421 = `^[^\s\v,;]+[\s\v,;].*?(?:application/(?:.+\+)?json|(?:application/(?:soap\+)?|text/)xml)`

var re921421 *re2.Regexp

type Rule921421 struct {
}

func (r *Rule921421) Id() string {
	return "921421"
}

func (r *Rule921421) Phase() int {
	return 1
}

func (r *Rule921421) Evaluate(tx *core.Transaction) int {
	if m := re921421.MatchString(tx.Variables.RequestHeaders["content-type"]); m {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}

func init() {
	re921421, _ = re2.Compile(PTN_921421)
}
