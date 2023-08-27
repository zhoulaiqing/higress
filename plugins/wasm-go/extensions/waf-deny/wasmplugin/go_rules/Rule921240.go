package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/wasilibs/go-re2"
)

const PTN_921240 = `unix:[^|]*\|`

var re921240 *re2.Regexp

type Rule921240 struct {
}

func (r *Rule921240) Id() string {
	return "921240"
}

func (r *Rule921240) Phase() int {
	return 1
}

func (r *Rule921240) Evaluate(tx *core.Transaction) bool {
	if m := re921240.MatchString(tx.Variables.RequestUri); m {
		tx.Variables.InboundAnomalyScorePl1 += CRITICAL_ANOMALY_SCORE
	}

	return true
}

func init() {
	re921240, _ = re2.Compile(PTN_921240)
}
