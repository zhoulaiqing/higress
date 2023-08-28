package rule_920

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"golang.org/x/exp/slices"
)

type Rule920430 struct {
}

func (r *Rule920430) Id() string {
	return "920430"
}

func (r *Rule920430) Phase() int {
	return 1
}

func (r *Rule920430) Evaluate(tx *core.Transaction) bool {
	if !slices.Contains(ALLOWED_HTTP_VERSIONS, tx.Variables.RequestProtocol) {
		tx.Variables.InboundAnomalyScorePl1 += go_rules.CRITICAL_ANOMALY_SCORE
	}

	return true
}
