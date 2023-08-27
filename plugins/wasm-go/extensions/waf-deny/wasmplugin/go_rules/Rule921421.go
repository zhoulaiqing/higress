package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
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

func (r *Rule921421) Evaluate(tx *core.Transaction) bool {
	if m := re921421.MatchString(tx.Variables.RequestHeaders["content-type"]); m {
		tx.Variables.InboundAnomalyScorePl1 += CRITICAL_ANOMALY_SCORE
	}

	return true
}

func init() {
	re921421, _ = re2.Compile(PTN_921421)
}
