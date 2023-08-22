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

func (r *Rule920170) Evaluate(tx core.Transaction) bool {
	if tx.Variables.RequestMethod != "GET" && tx.Variables.RequestMethod != "HEAD" {
		return true
	}

	contentLength := tx.Variables.RequestHeaders["Content-Length"]

	if len(contentLength) > 0 && strings.Trim(contentLength, " ") != "0" {
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
	}
	if len(tx.Variables.RequestHeaders["Transfer-Encoding"]) > 0 {
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
	}

	return true
}
