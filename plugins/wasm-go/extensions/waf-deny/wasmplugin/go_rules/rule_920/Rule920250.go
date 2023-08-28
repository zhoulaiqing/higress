package rule_920

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
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

func (r *Rule920250) Evaluate(tx *core.Transaction) int {
	if rule_tasks.CRS_VALIDATE_UTF8_ENCODING != 1 {
		return rule_tasks.PASS
	}

	if !utf8.ValidString(tx.Variables.RequestFileName) {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.WARNING_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	for _, argMap := range tx.Variables.Args {
		for ak, av := range *argMap {
			if !utf8.ValidString(ak) {
				tx.Variables.InboundAnomalyScorePl1 += rule_tasks.WARNING_ANOMALY_SCORE
				return rule_tasks.BLOCK
			}

			if !utf8.ValidString(av) {
				tx.Variables.InboundAnomalyScorePl1 += rule_tasks.WARNING_ANOMALY_SCORE
				return rule_tasks.BLOCK
			}
		}
	}

	return rule_tasks.PASS
}
