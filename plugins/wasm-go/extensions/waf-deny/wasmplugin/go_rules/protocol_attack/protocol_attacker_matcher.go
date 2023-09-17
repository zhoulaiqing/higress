package protocol_attack

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/wasilibs/go-re2"
	"strings"
)

var ptn = `^[^:\(\)\&\|\!\<\>\~]*\)\s*(?:\((?:[^,\(\)\=\&\|\!\<\>\~]+[><~]?=|\s*[&!|]\s*(?:\)|\()?\s*)|\)\s*\(\s*[\&\|\!]\s*|[&!|]\s*\([^\(\)\=\&\|\!\<\>\~]+[><~]?=[^:\(\)\&\|\!\<\>\~]*)`
var re921200 = re2.MustCompile(`^[^:\(\)\&\|\!\<\>\~]*\)\s*(?:\((?:[^,\(\)\=\&\|\!\<\>\~]+[><~]?=|\s*[&!|]\s*(?:\)|\()?\s*)|\)\s*\(\s*[\&\|\!]\s*|[&!|]\s*\([^\(\)\=\&\|\!\<\>\~]+[><~]?=[^:\(\)\&\|\!\<\>\~]*)`)

func matchDefaultRisk(value string) bool {

	data := []byte(value)
	if matchRisk(data) {
		return true
	}

	//if re921200.MatchString(value) {
	//	return true
	//}

	//if m, _ := hyperscan.MatchString(ptn, value); m {
	//	return true
	//}

	if matchLdapInjection(data) {
		return true
	}

	return false
}

func matchArgGet(value string) bool {
	data := []byte(value)
	if matchHeaderInjection(data) {
		return true
	}

	if re921200.MatchString(value) {
		return true
	}

	return false
}

func transformDefault(value string) string {
	v := value
	v, _, _ = core.UrlDecode(v)
	v, _, _ = core.HtmlEntityDecode(v)
	v = strings.ToLower(v)

	return v
}
