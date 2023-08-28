package rule_920

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules"
)

type Rule920180 struct {
}

func (r *Rule920180) Id() string {
	return "920180"
}

func (r *Rule920180) Phase() int {
	return 1
}

func (r *Rule920180) Evaluate(tx *core.Transaction) bool {
	if tx.Variables.RequestProtocol == "HTTP/2" || tx.Variables.RequestProtocol == "HTTP/2.0" {
		return true
	}

	if tx.Variables.RequestMethod == "POST" && len(tx.Variables.RequestHeaders["content-length"]) == 0 &&
		len(tx.Variables.RequestHeaders["transfer-encoding"]) == 0 {
		tx.Variables.InboundAnomalyScorePl1 += go_rules.WARNING_ANOMALY_SCORE
	}

	return true
}
