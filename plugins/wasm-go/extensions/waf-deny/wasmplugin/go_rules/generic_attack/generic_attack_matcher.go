package generic_attack

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"strings"
)

func matchDefault(value string, valueDecode string) bool {

	//在调用 insecure_unserialize 之前需要去除空格
	v, _, _ := core.RemoveWhitespace(value)
	data2 := []byte(v)
	vd, _, _ := core.RemoveWhitespace(valueDecode)
	data2Decode := []byte(vd)

	if multiMatch(matchInsecureUnserialize, data2, data2Decode) {
		return true
	}

	return false
}

func transformDefault(value string) string {
	v := value
	v, _, _ = core.UrlDecodeUni(v)
	v, _, _ = core.JsDecode(v)

	return v
}

func transformWithBase64Decode(value string) string {
	v := value
	v, _, _ = core.Base64decode(v)
	v, _, _ = core.UrlDecodeUni(v)
	v, _, _ = core.JsDecode(v)

	return strings.ToLower(v)
}

type MatchFunc func([]byte) bool

func multiMatch(doMatchFunc MatchFunc, data []byte, dataDecode []byte) bool {
	if doMatchFunc(data) {
		return true
	}

	if doMatchFunc(dataDecode) {
		return true
	}

	return false
}
