package go_rules

import "github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"

type Rule interface {
	Id() string
	Phase() int
	Evaluate(tx *core.Transaction) bool
}

const (
	PASS  = 1
	BLOCK = 0
	DENY  = -1
)
