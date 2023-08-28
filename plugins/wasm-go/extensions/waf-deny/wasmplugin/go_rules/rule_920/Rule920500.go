package rule_920

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
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

func (r *Rule920500) Evaluate(tx *core.Transaction) int {

	fileName, _, _ := core.UrlDecodeUni(tx.Variables.RequestFileName)

	if m := re920500.MatchString(fileName); m {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}

func init() {
	re920500, _ = re2.Compile(PTN_920500)
}
