package rule_955

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule955310 struct {
}

func (r *Rule955310) Id() string {
	return "955310"
}

func (r *Rule955310) Phase() int {
	return 4
}

func (r *Rule955310) Evaluate(tx *core.Transaction) int {
	if rule_tasks.Re955310.MatchString(tx.Variables.ResponseBody) {
		tx.Variables.OutboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}
