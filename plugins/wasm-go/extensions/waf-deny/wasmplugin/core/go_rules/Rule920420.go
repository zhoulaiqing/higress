package core

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/flier/gohs/hyperscan"
	"strings"
)

const PTN_920420 = `^[^;\s]+`

type Rule920420 struct {
}

func (r *Rule920420) Id() string {
	return "920420"
}

func (r *Rule920420) Phase() int {
	return 1
}

func (r *Rule920420) Evaluate(tx core.Transaction) bool {
	m, _ := hyperscan.MatchString(PTN_920420, tx.Variables.RequestHeaders["Content-Type"])
	if !m {
		return true
	}

	isAllowed := false
	for _, allowed := range core.ALLOWED_REQUEST_CONTENT_TYPE {
		if strings.Contains(tx.Variables.RequestHeaders["Content-Type"], allowed) {
			isAllowed = true
			break
		}
	}

	if !isAllowed {
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
	}

	return true
}
