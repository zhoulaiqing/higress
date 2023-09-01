package rule_954

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule954110 struct {
}

func (r *Rule954110) Id() string {
	return "954110"
}

func (r *Rule954110) Phase() int {
	return 4
}

func (r *Rule954110) Evaluate(tx *core.Transaction) int {
	if rule_tasks.Re954110.MatchString(tx.Variables.ResponseBody) {
		tx.Variables.OutboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}
	return rule_tasks.PASS
}
