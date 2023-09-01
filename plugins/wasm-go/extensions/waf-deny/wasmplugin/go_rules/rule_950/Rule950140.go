package rule_950

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule950140 struct {
}

func (r *Rule950140) Id() string {
	return "950140"
}

func (r *Rule950140) Phase() int {
	return 4
}

func (r *Rule950140) Evaluate(tx *core.Transaction) int {
	if rule_tasks.Re950140.MatchString(tx.Variables.ResponseBody) {
		tx.Variables.OutboundAnomalyScorePl1 += rule_tasks.ERROR_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}
