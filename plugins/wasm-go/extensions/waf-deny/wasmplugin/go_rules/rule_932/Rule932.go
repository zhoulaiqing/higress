package rule_932

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
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

func (r *Rule932) Evaluate(tx *core.Transaction) bool {
	for k, v := range tx.Variables.RequestCookies {
		if strings.Contains(k, "__utm") {
			continue
		}

		if r.doEvaluate(tx, k) {
			return true
		}

		if r.doEvaluate(tx, v) {
			return true
		}
	}

	for _, argMap := range tx.Variables.Args {
		for k, v := range *argMap {
			if r.doEvaluate(tx, k) {
				return true
			}

			if r.doEvaluate(tx, v) {
				return true
			}
		}
	}

	for _, v := range tx.Variables.XML["/*"] {
		if r.doEvaluate(tx, v) {
			return true
		}
	}

	return true
}

func (r *Rule932) doEvaluate(tx *core.Transaction, value string) bool {
	return false
}
