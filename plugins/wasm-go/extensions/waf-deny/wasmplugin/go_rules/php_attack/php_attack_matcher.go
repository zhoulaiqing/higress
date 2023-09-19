package php_attack

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"strings"
)

func matchDefault(value string) bool {

	lowerV := transformDefault(value)

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
	if matchPhpObjectInjection(data) {
		return true
	}
	if matchPhpWrapper([]byte(transformForWrapper(value))) {
		return true
	}

	return false
}

func matchHeader(value string) bool {
	lowerV := transformDefault(value)
	data := []byte(lowerV)

	if matchPhpObjectInjection(data) {
		return true
	}

	return false
}

func matchRequestFileName(value string) bool {
	lowerV := transformDefault(value)
	data := []byte(lowerV)
	if matchPhpFunctions(data) {
		return true
	}
	if matchPhpVariableFunction(data) {
		return true
	}

	return false
}

func matchFile(value string) bool {
	lowerV := transformDefault(value)
	data := []byte(lowerV)

	if matchPhpScriptUploads(data) {
		return true
	}

	return false
}

func transformDefault(value string) string {
	v, _, _ := core.UrlDecodeUni(value)
	return strings.ToLower(v)
}

func transformForWrapper(value string) string {
	v, _, _ := core.Utf8ToUnicode(value)
	v, _, _ = core.UrlDecodeUni(v)
	v, _, _ = core.RemoveNulls(v)
	v, _, _ = core.CmdLine(v)

	return v
}
