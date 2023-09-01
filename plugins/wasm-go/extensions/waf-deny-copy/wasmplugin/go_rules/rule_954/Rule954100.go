package rule_954

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"strings"
)

type Rule954100 struct {
}

func (r *Rule954100) Id() string {
	return "954100"
}

func (r *Rule954100) Phase() int {
	return 4
}

func (r *Rule954100) Evaluate(tx *core.Transaction) int {
	if rule_tasks.Re954100.MatchString(strings.ToLower(tx.Variables.ResponseBody)) {
		tx.Variables.OutboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}
	return rule_tasks.PASS
}
