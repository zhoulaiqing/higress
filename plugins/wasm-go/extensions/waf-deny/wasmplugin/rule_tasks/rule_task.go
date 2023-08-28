package rule_tasks

import "github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"

const (
	PASS  = 1
	BLOCK = 0
	DENY  = -1
)

type RuleTask interface {
	AssociatedVar() string
	Evaluate(tx *core.Transaction) int
}
