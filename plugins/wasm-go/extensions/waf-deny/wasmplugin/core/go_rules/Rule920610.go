package go_rules

import "github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"

type Rule920610 struct {
}

func (r *Rule920610) Id() string {
	return "920610"
}

func (r *Rule920610) Phase() int {
	return 1
}

func (r *Rule920610) Evaluate(tx core.Transaction) bool {
	// todo should use REQUEST_URI_RAW
	return true
}
