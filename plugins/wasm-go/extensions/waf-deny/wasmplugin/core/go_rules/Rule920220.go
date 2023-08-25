package core

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"strings"
)

type Rule920220 struct {
}

func (r *Rule920220) Id() string {
	return "920220"
}

func (r *Rule920220) Phase() int {
	return 1
}

func (r *Rule920220) Evaluate(tx core.Transaction) bool {
	requestUri := tx.Variables.RequestUri

	if strings.Contains(requestUri, "%") && core.ValidateURLEncoding(requestUri) {
		tx.Variables.InboundAnomalyScorePl1 += core.WARNING_ANOMALY_SCORE
	}

	return true
}