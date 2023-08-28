package rule_920

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule920340 struct {
}

func (r *Rule920340) Id() string {
	return "920340"
}

func (r *Rule920340) Phase() int {
	return 1
}

func (r *Rule920340) Evaluate(tx *core.Transaction) int {
	contentLength, ok := tx.Variables.RequestHeaders["content-length"]
	if !ok {
		return rule_tasks.PASS
	}
	if contentLength == "0" {
		return rule_tasks.PASS
	}

	if len(tx.Variables.RequestHeaders["content-type"]) == 0 {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.NOTICE_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}
