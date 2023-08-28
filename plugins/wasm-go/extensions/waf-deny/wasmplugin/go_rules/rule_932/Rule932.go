package rule_932

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"strings"
)

type Rule932 struct {
}

func (r *Rule932) Id() string {
	return "932"
}

func (r *Rule932) Phase() int {
	return 2
}

func (r *Rule932) Evaluate(tx *core.Transaction) int {
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

	return rule_tasks.PASS
}

func (r *Rule932) doEvaluate(tx *core.Transaction, value string) bool {
	return false
}
