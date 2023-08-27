package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/wasilibs/go-re2"
)

const PTN_920500 = `\.[^.~]+~(?:/.*|)$`

var re920500 *re2.Regexp

type Rule920500 struct {
}

func (r *Rule920500) Id() string {
	return "920500"
}

func (r *Rule920500) Phase() int {
	return 1
}

func (r *Rule920500) Evaluate(tx *core.Transaction) bool {

	fileName, _, _ := core.UrlDecodeUni(tx.Variables.RequestFileName)

	if m := re920500.MatchString(fileName); m {
		tx.Variables.InboundAnomalyScorePl1 += CRITICAL_ANOMALY_SCORE
	}

	return true
}

func init() {
	re920500, _ = re2.Compile(PTN_920500)
}
