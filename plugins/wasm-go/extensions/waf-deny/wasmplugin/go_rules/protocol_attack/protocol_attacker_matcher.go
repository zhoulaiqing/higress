package protocol_attack

import (
	"github.com/tianchi/waf/wasmplugin/core"
	"strings"
)

func matchDefaultRisk(value string) bool {

	data := []byte(value)
	if matchRisk(data) {
		return true
	}

	//if re921200.MatchString(value) {
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

	if matchLdapInjection(data) {
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
