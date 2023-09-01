package rule_953

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule953100 struct {
}

func (r *Rule953100) Id() string {
	return "953100"
}

func (r *Rule953100) Phase() int {
	return 4
}

func (r *Rule953100) Evaluate(tx *core.Transaction) int {
	if m, _ := core.PmEvaluate(rule_tasks.Rule953100Matcher, tx.Variables.ResponseBody, false); m {
		tx.Variables.OutboundAnomalyScorePl1 += rule_tasks.ERROR_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}
