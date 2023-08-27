package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"strings"
)

type Rule920170 struct {
}

func (r *Rule920170) Id() string {
	return "920170"
}

func (r *Rule920170) Phase() int {
	return 1
}

func (r *Rule920170) Evaluate(tx *core.Transaction) bool {
	if tx.Variables.RequestMethod != "GET" && tx.Variables.RequestMethod != "HEAD" {
		return true
	}

	contentLength := tx.Variables.RequestHeaders["content-length"]

	if len(contentLength) > 0 && strings.TrimSpace(contentLength) != "0" {
		tx.Variables.InboundAnomalyScorePl1 += CRITICAL_ANOMALY_SCORE
	}
	if len(tx.Variables.RequestHeaders["transfer-encoding"]) > 0 {
		tx.Variables.InboundAnomalyScorePl1 += CRITICAL_ANOMALY_SCORE
	}

	return true
}
