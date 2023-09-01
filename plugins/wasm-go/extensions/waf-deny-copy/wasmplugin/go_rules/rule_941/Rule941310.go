package rule_941

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"strings"
)

type Rule941310 struct {
	*Rule941
}

func (r *Rule941310) Id() string {
	return "941310"
}

func (r *Rule941310) Evaluate(tx *core.Transaction) int {
	return r.evaluateByCache(tx, r.doEvaluate, fileNameAddition)
}

// todo: 与原逻辑不是很一致，回头 check
func (r *Rule941310) doEvaluate(tx *core.Transaction, value *string) bool {

	v := strings.ToLower(*value)
	//proxywasm.LogInfof("evaluate value: %s, v: %s", *value, v)
	m := rule_tasks.Re9413101.MatchString(v)

	//proxywasm.LogInfof("9413101 match: %q ", m)
	if !m {
		return false
	}

	m = rule_tasks.Re9413102.MatchString(v)
	//proxywasm.LogInfof("9413102 match: %q ", m)
	if m {
		tx.Variables.XssScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
