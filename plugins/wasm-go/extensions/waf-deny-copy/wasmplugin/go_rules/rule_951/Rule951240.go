package rule_951

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule951240 struct {
	*Rule951
}

func (r *Rule951240) Id() string {
	return "951240"
}

func (r *Rule951240) Evaluate(tx *core.Transaction) int {
	if r.evaluate(tx, r.doEvaluate) {
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}

func (r *Rule951240) doEvaluate(tx *core.Transaction) bool {
	return rule_tasks.Re951240.MatchString(tx.Variables.ResponseBody)
}
