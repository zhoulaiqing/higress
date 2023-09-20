package go_rules

import "github.com/tianchi/waf/wasmplugin/core"

type Rule interface {
	Id() string
	Phase() int
	Evaluate(tx *core.Transaction) int
}
