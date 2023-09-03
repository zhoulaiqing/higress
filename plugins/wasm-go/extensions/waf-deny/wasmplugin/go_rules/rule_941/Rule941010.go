package rule_941

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule941010 struct {
	*Rule941
}

func (r *Rule941010) Id() string {
	return "941010"
}

func (r *Rule941010) Phase() int {
	return 1
}

func (r *Rule941010) Evaluate(tx *core.Transaction) int {
	if !rule_tasks.ValidateByteRange941010.Validate(tx.Variables.RequestFileName) {
		tx.Variables.Skip941ForFileName = true
	}

	return rule_tasks.PASS
}
