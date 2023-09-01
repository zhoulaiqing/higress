package rule_954

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule954120 struct {
}

func (r *Rule954120) Id() string {
	return "954120"
}

func (r *Rule954120) Phase() int {
	return 4
}

func (r *Rule954120) Evaluate(tx *core.Transaction) int {
	if m, _ := core.PmEvaluate(rule_tasks.Rule954120Matcher, tx.Variables.ResponseBody, false); m {
		tx.Variables.OutboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}
	return rule_tasks.PASS
}
