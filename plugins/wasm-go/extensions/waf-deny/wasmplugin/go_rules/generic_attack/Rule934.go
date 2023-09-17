package generic_attack

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"strings"
)

type Rule934 struct {
}

func (r *Rule934) Id() string {
	return "934"
}

func (r *Rule934) Phase() int {
	return 2
}

func (r *Rule934) Evaluate(tx *core.Transaction) int {
	for k, v := range tx.Variables.RequestCookies {
		if strings.Contains(k, "__utm") {
			continue
		}

		if r.doEvaluate(tx, k) {
			return rule_tasks.BLOCK
		}

		if r.doEvaluate(tx, v) {
			return rule_tasks.BLOCK
		}
	}

	for _, argMap := range tx.Variables.Args {
		for k, v := range *argMap {
			if r.doEvaluate(tx, k) {
				return rule_tasks.BLOCK
			}

			if r.doEvaluate(tx, v) {
				return rule_tasks.BLOCK
			}
		}
	}

	for _, v := range tx.Variables.XML["/*"] {
		if r.doEvaluate(tx, v) {
			return rule_tasks.BLOCK
		}
	}

	if r.validateFileName() && r.doEvaluate(tx, tx.Variables.RequestFileName) {
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}

func (r *Rule934) validateFileName() bool {
	return false
}

func (r *Rule934) doEvaluate(tx *core.Transaction, value string) bool {
	return false
}
