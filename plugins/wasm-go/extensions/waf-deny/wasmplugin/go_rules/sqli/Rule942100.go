package sqli

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	libinjection2 "github.com/wasilibs/go-libinjection"
)

type Rule942100 struct {
	*Rule942
}

func (r *Rule942100) Id() string {
	return "942100"
}

func (r *Rule942100) Evaluate(tx *core.Transaction) int {
	return r.evaluateRawValue(tx, r.doEvaluate, additionWithHeader)
}

// This is multiMatch
func (r *Rule942100) doEvaluate(tx *core.Transaction, value *string) bool {
	// 初始化缓存
	emptyString := ""
	r.Rule942.doEvaluate(tx, &emptyString)

	m, _ := libinjection2.IsSQLi(*value)
	if m {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.SqlInjectionScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		return true
	}

	v, _, _ := core.Utf8ToUnicode(*value)
	m, _ = libinjection2.IsSQLi(v)
	if m {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.SqlInjectionScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		return true
	}

	v, _, _ = core.UrlDecodeUni(v)
	m, _ = libinjection2.IsSQLi(v)
	if m {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.SqlInjectionScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		return true
	}

	v, _, _ = core.RemoveNulls(v)
	m, _ = libinjection2.IsSQLi(v)
	if m {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.SqlInjectionScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		return true
	}

	return false
}