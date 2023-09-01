package rule_952

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule952110 struct {
}

func (r *Rule952110) Id() string {
	return "952110"
}

func (r *Rule952110) Phase() int {
	return 4
}

func (r *Rule952110) Evaluate(tx *core.Transaction) int {
	if m, _ := core.PmEvaluate(rule_tasks.Rule952110Matcher, tx.Variables.ResponseBody, false); m {
		tx.Variables.OutboundAnomalyScorePl1 += rule_tasks.ERROR_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}
