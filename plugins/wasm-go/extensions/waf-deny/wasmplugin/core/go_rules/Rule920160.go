package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"strconv"
)

type Rule920160 struct {
}

func (r *Rule920160) Id() string {
	return "920160"
}

func (r *Rule920160) Phase() int {
	return 1
}

func (r *Rule920160) Evaluate(tx core.Transaction) bool {

	contentLength := tx.Variables.RequestHeaders["Content-Length"]
	if len(contentLength) == 0 {
		return true
	}
	if _, err := strconv.Atoi(contentLength); err != nil {
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
	}

	return true
}
