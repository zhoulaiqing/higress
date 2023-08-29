package rule_942

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"github.com/corazawaf/libinjection-go"
)

type Rule942100 struct {
	*Rule942
}

func (r *Rule942100) Id() string {
	return "942100"
}

func (r *Rule942100) Evaluate(tx *core.Transaction) int {
	return r.evaluateRawValue(tx)
}

func (r *Rule942100) GetAddition() *Rule942Addition {
	return additionWithHeader
}

// This is multiMatch
func (r *Rule942100) doEvaluate(tx *core.Transaction, value *string) bool {

	m, _ := libinjection.IsSQLi(*value)
	if m {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.SqlInjectionScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		return true
	}

	v, _, _ := core.Utf8ToUnicode(*value)
	m, _ = libinjection.IsSQLi(v)
	if m {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.SqlInjectionScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		return true
	}

	v, _, _ = core.UrlDecodeUni(v)
	m, _ = libinjection.IsSQLi(v)
	if m {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.SqlInjectionScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		return true
	}

	v, _, _ = core.RemoveNulls(v)
	m, _ = libinjection.IsSQLi(v)
	if m {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.SqlInjectionScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		return true
	}

	return false
}
