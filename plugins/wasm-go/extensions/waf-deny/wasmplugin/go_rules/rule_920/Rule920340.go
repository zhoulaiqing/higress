package rule_920

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules"
)

type Rule920340 struct {
}

func (r *Rule920340) Id() string {
	return "920340"
}

func (r *Rule920340) Phase() int {
	return 1
}

func (r *Rule920340) Evaluate(tx *core.Transaction) bool {
	contentLength, ok := tx.Variables.RequestHeaders["content-length"]
	if !ok {
		return true
	}
	if contentLength == "0" {
		return true
	}

	if len(tx.Variables.RequestHeaders["content-type"]) == 0 {
		tx.Variables.InboundAnomalyScorePl1 += go_rules.NOTICE_ANOMALY_SCORE
	}

	return true
}
