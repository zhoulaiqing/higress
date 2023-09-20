package rce_attack

import "github.com/tianchi/waf/wasmplugin/core"

func matchDefault(value string) bool {
	//v := transformForCmd(value)
	//data := []byte(v)

	//if matchShortCommandBase(data) {
	//	return true
	//}

	return false
}

func transformForCmd(value string) string {
	v, _, _ := core.CmdLineRce(value)

	return v
}
