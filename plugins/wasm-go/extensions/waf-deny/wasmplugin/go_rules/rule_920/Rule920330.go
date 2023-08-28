package rule_920

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule920330 struct {
}

func (r *Rule920330) Id() string {
	return "920330"
}

func (r *Rule920330) Phase() int {
	return 1
}

func (r *Rule920330) Evaluate(tx *core.Transaction) int {
	if len(tx.Variables.RequestHeaders["user-agent"]) == 0 {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.NOTICE_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}
