package rce_attack

import "github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"

func matchDefault(value string) bool {
	v := transformForCmd(value)
	data := []byte(v)

	if matchShortCommandBase(data) {
		return true
	}

	return false
}

func transformForCmd(value string) string {
	v, _, _ := core.UrlDecodeUni(value)
	v, _, _ = core.CmdLineRce(v)

	return v
}
