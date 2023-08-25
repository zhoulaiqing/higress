package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"unicode/utf8"
)

type Rule920250 struct {
}

func (r *Rule920250) Id() string {
	return "920250"
}

func (r *Rule920250) Phase() int {
	return 2
}

func (r *Rule920250) Evaluate(tx core.Transaction) bool {
	if core.CRS_VALIDATE_UTF8_ENCODING != 1 {
		return true
	}

	if !utf8.ValidString(tx.Variables.RequestFileName) {
		tx.Variables.InboundAnomalyScorePl1 += core.WARNING_ANOMALY_SCORE
		return true
	}

	for _, argMap := range tx.Variables.Args {
		for ak, av := range *argMap {
			if !utf8.ValidString(ak) {
				tx.Variables.InboundAnomalyScorePl1 += core.WARNING_ANOMALY_SCORE
				return true
			}

			if !utf8.ValidString(av) {
				tx.Variables.InboundAnomalyScorePl1 += core.WARNING_ANOMALY_SCORE
				return true
			}
		}
	}

	return true
}
