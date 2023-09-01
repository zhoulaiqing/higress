package go_rules

import "github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"

type Rule interface {
	Id() string
	Phase() int
	Evaluate(tx *core.Transaction) int
}
