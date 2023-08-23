package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"strings"
)

type Rule920240 struct {
}

func (r *Rule920240) Id() string {
	return "920240"
}

func (r *Rule920240) Evaluate(tx core.Transaction) bool {
	if !strings.HasPrefix(tx.Variables.RequestHeaders["Content-Type"], "application/x-www-form-urlencoded") {
		return true
	}

	requestBody := tx.Variables.RequestBody
	if strings.Contains(requestBody, "%") && core.ValidateURLEncoding(requestBody) {
		tx.Variables.InboundAnomalyScorePl1 += core.WARNING_ANOMALY_SCORE
	}

	return true
}
