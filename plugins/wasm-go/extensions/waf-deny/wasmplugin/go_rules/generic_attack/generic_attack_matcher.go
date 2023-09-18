package generic_attack

import (
	"bytes"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"strings"
)

func matchDefault(value string, valueDecode string) bool {

	//在调用 insecure_unserialize 之前需要去除空格
	vrw, _, _ := core.RemoveWhitespace(value)
	data2 := []byte(vrw)

	var data2Decode []byte
	if value != valueDecode {
		vrwdec, _, _ := core.RemoveWhitespace(valueDecode)
		data2Decode = []byte(vrwdec)
	}

	if multiMatch(matchInsecureUnserialize, data2, data2Decode) {
		return true
	}

	vrc, _, _ := core.ReplaceComments(value)
	dataRc := []byte(vrc)

	var dataRcDecode []byte
	if value != valueDecode {
		vrcdec, _, _ := core.ReplaceComments(valueDecode)
		dataRcDecode = []byte(vrcdec)
	}

	if multiMatch(matchNodejsDos, dataRc, dataRcDecode) {
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

	if len(dataDecode) > 0 && !bytes.Equal(data, dataDecode) {
		if doMatchFunc(dataDecode) {
			return true
		}
	}

	return false
}
