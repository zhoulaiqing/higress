package rule_933

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule933200 struct {
	*Rule933
}

func (r *Rule933200) Id() string {
	return "933200"
}

func (r *Rule933200) doEvaluate(tx *core.Transaction, value string) bool {
	v, _, _ := core.Utf8ToUnicode(value)
	v, _, _ = core.UrlDecodeUni(v)
	v, _, _ = core.RemoveNulls(v)
	v, _, _ = core.CmdLine(v)

	m := rule_tasks.Re933200.MatchString(v)
	if !m {
		tx.Variables.PhpInjectionScore += go_rules.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += go_rules.CRITICAL_ANOMALY_SCORE
	}

	return m
}
