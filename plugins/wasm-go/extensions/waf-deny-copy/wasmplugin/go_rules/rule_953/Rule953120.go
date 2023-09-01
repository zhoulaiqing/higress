package rule_953

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule953120 struct {
}

func (r *Rule953120) Id() string {
	return "953120"
}

func (r *Rule953120) Phase() int {
	return 4
}

func (r *Rule953120) Evaluate(tx *core.Transaction) int {
	if rule_tasks.Re953120.MatchString(tx.Variables.ResponseBody) {
		tx.Variables.OutboundAnomalyScorePl1 += rule_tasks.ERROR_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}
