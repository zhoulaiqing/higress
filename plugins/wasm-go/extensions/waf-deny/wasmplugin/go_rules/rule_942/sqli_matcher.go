package rule_942

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/libinjection-go"
)

func matchSqli(value string) bool {
	v, _, _ := core.UrlDecodeUni(value)

	if m, _ := libinjection.IsSQLi(v); m {
		return true
	}

	return false
}
