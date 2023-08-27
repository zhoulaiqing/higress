package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"strings"
)

type Rule920610 struct {
}

func (r *Rule920610) Id() string {
	return "920610"
}

func (r *Rule920610) Phase() int {
	return 1
}

func (r *Rule920610) Evaluate(tx *core.Transaction) bool {

	if strings.Contains(tx.Variables.RequestUriRaw, "#") {
		tx.Variables.InboundAnomalyScorePl1 += CRITICAL_ANOMALY_SCORE
	}

	return true
}
