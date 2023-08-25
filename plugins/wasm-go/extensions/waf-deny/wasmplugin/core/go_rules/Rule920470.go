package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/flier/gohs/hyperscan"
	"strings"
)

const PTN_920470 = `^[\w/.+*-]+(?:\s?;\s?(?:action|boundary|charset|component|start(?:-info)?|type|version)\s?=\s?['\"\w.()+,/:=?<>@#*-]+)*$`

type Rule920470 struct {
}

func (r *Rule920470) Id() string {
	return "920470"
}

func (r *Rule920470) Phase() int {
	return 1
}

func (r *Rule920470) Evaluate(tx core.Transaction) bool {
	contentType := strings.ToLower(tx.Variables.RequestHeaders["Content-Type"])

	if m, _ := hyperscan.MatchString(PTN_920470, contentType); !m {
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
	}

	return true
}
