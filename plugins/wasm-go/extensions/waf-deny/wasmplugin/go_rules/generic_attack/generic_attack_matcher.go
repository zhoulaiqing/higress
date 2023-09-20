package generic_attack

import (
	"bytes"
	"github.com/tianchi/waf/wasmplugin/core"
	"github.com/tianchi/waf/wasmplugin/rule_tasks"
	"strings"
)

func matchDefault(value string) bool {

	if m, _ := core.PmEvaluate(rule_tasks.Rule934110Matcher, value, false); m {
		return true
	}

	tv := transformDefault(value)
	tvd := transformWithBase64Decode(tv)
	tv = strings.ToLower(tv)
	data := []byte(tv)

	//在调用 insecure_unserialize 之前需要去除空格
	vrw, _, _ := core.RemoveWhitespace(tv)
	data2 := []byte(vrw)

	var data2Decode []byte
	if tv != tvd {
		vrwdec, _, _ := core.RemoveWhitespace(tvd)
		data2Decode = []byte(vrwdec)
	}

	if multiMatch(matchInsecureUnserialize, data2, data2Decode) {
		return true
	}

	vrc, _, _ := core.ReplaceComments(tv)
	dataRc := []byte(vrc)

	var dataRcDecode []byte
	if tv != tvd {
		vrcdec, _, _ := core.ReplaceComments(tvd)
		dataRcDecode = []byte(vrcdec)
	}

	if multiMatch(matchNodejsDos, dataRc, dataRcDecode) {
		return true
	}

	if matchPhpDataScheme(data) {
		return true
	}

	return false
}

func transformDefault(value string) string {
	v := value
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
