package rule_953

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule953110 struct {
}

func (r *Rule953110) Id() string {
	return "953110"
}

func (r *Rule953110) Phase() int {
	return 4
}

func (r *Rule953110) Evaluate(tx *core.Transaction) int {
	if rule_tasks.Re953110.MatchString(tx.Variables.ResponseBody) {
		tx.Variables.OutboundAnomalyScorePl1 += rule_tasks.ERROR_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}
