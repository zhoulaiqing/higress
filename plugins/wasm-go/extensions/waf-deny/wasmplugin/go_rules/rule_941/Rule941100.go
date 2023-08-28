package rule_941

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule941100 struct {
	*Rule941
}

func (r *Rule941100) Id() string {
	return "941100"
}

func (r *Rule941100) doEvaluate(tx *core.Transaction, value string) bool {

	v, _, _ := core.Utf8ToUnicode(value)
	v, _, _ = core.JsDecode(v)
	v, _, _ = core.RemoveWhitespace(v)
	v, _, _ = core.Base64decode(v)

	m := rule_tasks.Re934100.MatchString(v)
	if m {
		tx.Variables.RceScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
