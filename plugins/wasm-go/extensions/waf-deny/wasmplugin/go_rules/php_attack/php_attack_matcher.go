package php_attack

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"strings"
)

func matchDefault(value string) bool {

	lowerV := strings.ToLower(value)

	data := []byte(lowerV)

	if matchPhpInjection(data) {
		return true
	}

	if matchPhpFunctions(data) {
		return true
	}
	if matchPhpVariableFunction(data) {
		return true
	}

	if matchPhpWrapper([]byte(transformForWrapper(value))) {
		return true
	}

	return false
}

func matchFile(value string) bool {
	lowerV := strings.ToLower(value)
	data := []byte(lowerV)

	if matchPhpScriptUploads(data) {
		return true
	}

	return false
}

func transformForWrapper(value string) string {
	v, _, _ := core.Utf8ToUnicode(value)
	v, _, _ = core.UrlDecodeUni(v)
	v, _, _ = core.RemoveNulls(v)
	v, _, _ = core.CmdLine(v)

	return v
}
