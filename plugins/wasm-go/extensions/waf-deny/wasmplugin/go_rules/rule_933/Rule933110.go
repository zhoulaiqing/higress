package rule_933

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules"
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

func (r *Rule933110) Evaluate(tx *core.Transaction) bool {
	for k, values := range tx.Variables.Files {
		if r.doEvaluate(tx, strings.ToLower(k)) {
			return true
		}

		for _, v := range values {
			if r.doEvaluate(tx, strings.ToLower(v)) {
				return true
			}
		}
	}

	for _, fh := range rule_tasks.FILE_NAME_HEADERS {
		fn, ok := tx.Variables.RequestHeaders[fh]
		if !ok {
			continue
		}
		if r.doEvaluate(tx, strings.ToLower(fn)) {
			return true
		}
	}

	return true
}

func (r *Rule933110) doEvaluate(tx *core.Transaction, value string) bool {

	m := rule_tasks.Re933110.MatchString(strings.ToLower(value))
	if m {
		tx.Variables.PhpInjectionScore += go_rules.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += go_rules.CRITICAL_ANOMALY_SCORE
	}

	return m
}
