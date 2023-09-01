package rule_952

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule952100 struct {
}

func (r *Rule952100) Id() string {
	return "952100"
}

func (r *Rule952100) Phase() int {
	return 4
}

func (r *Rule952100) Evaluate(tx *core.Transaction) int {
	if m, _ := core.PmEvaluate(rule_tasks.Rule952100Matcher, tx.Variables.ResponseBody, false); m {
		tx.Variables.OutboundAnomalyScorePl1 += rule_tasks.ERROR_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}
