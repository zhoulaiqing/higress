package rule_951

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule951200 struct {
	*Rule951
}

func (r *Rule951200) Id() string {
	return "951200"
}

func (r *Rule951200) Evaluate(tx *core.Transaction) int {
	if r.evaluate(tx, r.doEvaluate) {
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}

func (r *Rule951200) doEvaluate(tx *core.Transaction) bool {
	return rule_tasks.Re951200.MatchString(tx.Variables.ResponseBody)
}
