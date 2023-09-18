package php_attack

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"strings"
)

type Rule933110 struct {
}

func (r *Rule933110) Id() string {
	return "933110"
}

func (r *Rule933110) Phase() int {
	return 2
}

func (r *Rule933110) Evaluate(tx *core.Transaction) int {
	for k, values := range tx.Variables.Files {
		if r.doEvaluate(tx, strings.ToLower(k)) {
			return rule_tasks.BLOCK
		}

		for _, v := range values {
			if r.doEvaluate(tx, strings.ToLower(v)) {
				return rule_tasks.BLOCK
			}
		}
	}

	for _, fh := range rule_tasks.FILE_NAME_HEADERS {
		fn, ok := tx.Variables.RequestHeaders[fh]
		if !ok {
			continue
		}
		if r.doEvaluate(tx, strings.ToLower(fn)) {
			return rule_tasks.BLOCK
		}
	}

	return rule_tasks.PASS
}

func (r *Rule933110) doEvaluate(tx *core.Transaction, value string) bool {

	m := matchFile(value)
	if m {
		tx.Variables.PhpInjectionScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
