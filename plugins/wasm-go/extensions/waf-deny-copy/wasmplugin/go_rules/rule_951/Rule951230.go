package rule_951

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule951230 struct {
	*Rule951
}

func (r *Rule951230) Id() string {
	return "951230"
}

func (r *Rule951230) Evaluate(tx *core.Transaction) int {
	if r.evaluate(tx, r.doEvaluate) {
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}

func (r *Rule951230) doEvaluate(tx *core.Transaction) bool {
	return rule_tasks.Re951230.MatchString(tx.Variables.ResponseBody)
}
