package rule_955

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule955140 struct {
}

func (r *Rule955140) Id() string {
	return "955140"
}

func (r *Rule955140) Phase() int {
	return 4
}

func (r *Rule955140) Evaluate(tx *core.Transaction) int {
	if rule_tasks.Re955140.MatchString(tx.Variables.ResponseBody) {
		tx.Variables.OutboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}
