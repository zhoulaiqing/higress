package rule_951

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule951100 struct {
}

func (r *Rule951100) Id() string {
	return "951100"
}

func (r *Rule951100) Phase() int {
	return 4
}

func (r *Rule951100) Evaluate(tx *core.Transaction) int {
	m, _ := core.PmEvaluate(rule_tasks.Rule951100Matcher, tx.Variables.ResponseBody, false)
	if m {
		tx.Variables.Skip951 = true
	}
	return rule_tasks.PASS
}
