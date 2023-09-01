package rule_954

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule954130 struct {
}

func (r *Rule954130) Id() string {
	return "954130"
}

func (r *Rule954130) Phase() int {
	return 4
}

func (r *Rule954130) Evaluate(tx *core.Transaction) int {

	if tx.Variables.ResponseStatus == 404 {
		return rule_tasks.PASS
	}

	if rule_tasks.Re954130.MatchString(tx.Variables.ResponseBody) {
		tx.Variables.OutboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}
	return rule_tasks.PASS
}
