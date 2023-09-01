package rule_955

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule955150 struct {
}

func (r *Rule955150) Id() string {
	return "955150"
}

func (r *Rule955150) Phase() int {
	return 4
}

func (r *Rule955150) Evaluate(tx *core.Transaction) int {
	if rule_tasks.Re955150.MatchString(tx.Variables.ResponseBody) {
		tx.Variables.OutboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}
