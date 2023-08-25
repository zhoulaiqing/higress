package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"golang.org/x/exp/slices"
	"strings"
)

type Rule922100 struct {
}

func (r *Rule922100) Id() string {
	return "922100"
}

func (r *Rule922100) Phase() int {
	return 2
}

func (r *Rule922100) Evaluate(tx core.Transaction) bool {
	if len(tx.Variables.MultipartPartHeaders["_charset_"]) == 0 {
		return true
	}

	for _, argMap := range tx.Variables.Args {
		value, ok := (*argMap)["_charset_"]
		if ok && !slices.Contains(core.ALLOWED_REQUEST_CONTENT_TYPE_CHARSET, strings.ToLower(value)) {
			tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
			return true
		}
	}

	return true
}
