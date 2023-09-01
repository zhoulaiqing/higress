package rule_942

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"strings"
)

type Rule942 struct {
}

type Rule942Addition struct {
	validateFileName  bool
	validateUserAgent bool
	validateReferer   bool
	validateBaseName  bool
}

var noAddition942 = &Rule942Addition{
	validateFileName:  false,
	validateUserAgent: false,
	validateReferer:   false,
}

var additionWithHeader = &Rule942Addition{
	validateFileName:  false,
	validateBaseName:  false,
	validateUserAgent: true,
	validateReferer:   true,
}

var fileNameAddition = &Rule942Addition{
	validateFileName:  true,
	validateUserAgent: false,
	validateReferer:   false,
}

var baseNameAddition = &Rule942Addition{
	validateBaseName: true,
}

func (r *Rule942) Id() string {
	return "942"
}

func (r *Rule942) Phase() int {
	return 2
}

func (r *Rule942) Evaluate(tx *core.Transaction) int {
	//emptyString := ""
	//r.doEvaluate(tx, &emptyString)

	return rule_tasks.PASS
}

func (r *Rule942) GetAddition() *Rule942Addition {
	return noAddition942
}

type doEvaluateFunc func(*core.Transaction, *string) bool

func (r *Rule942) evaluateByCache(tx *core.Transaction, evaluateFunc doEvaluateFunc, addition *Rule942Addition) int {
	cookieKeyCache := tx.Variables.TransMap["cookie_key_942"]
	cookieValueCache := tx.Variables.TransMap["cookie_value_942"]
	for i, _ := range cookieKeyCache {
		if evaluateFunc(tx, &cookieKeyCache[i]) {
			return rule_tasks.BLOCK
		}

		if evaluateFunc(tx, &cookieValueCache[i]) {
			return rule_tasks.BLOCK
		}
	}

	argsKeyCache := tx.Variables.TransMap["arg_key_942"]
	argsValueCache := tx.Variables.TransMap["arg_value_942"]
	for i, _ := range argsKeyCache {
		if evaluateFunc(tx, &argsKeyCache[i]) {
			return rule_tasks.BLOCK
		}

		if evaluateFunc(tx, &argsValueCache[i]) {
			return rule_tasks.BLOCK
		}
	}

	xmlCache := tx.Variables.TransMap["xml_942"]
	for i, _ := range xmlCache {
		if evaluateFunc(tx, &xmlCache[i]) {
			return rule_tasks.BLOCK
		}
	}

	if addition.validateFileName {
		fileNameCache := tx.Variables.TransMap["file_name_942"]
		if len(fileNameCache) > 0 && evaluateFunc(tx, &fileNameCache[0]) {
			return rule_tasks.BLOCK
		}
	}

	if addition.validateReferer {
		refererCache := tx.Variables.TransMap["referer_942"]
		if len(refererCache) > 0 && evaluateFunc(tx, &refererCache[0]) {
			return rule_tasks.BLOCK
		}
	}

	if addition.validateUserAgent {
		uaCache := tx.Variables.TransMap["ua_942"]
		if len(uaCache) > 0 && evaluateFunc(tx, &uaCache[0]) {
			return rule_tasks.BLOCK
		}
	}

	if addition.validateBaseName {
		bnCache := tx.Variables.TransMap["base_name_942"]
		if len(bnCache) > 0 && evaluateFunc(tx, &bnCache[0]) {
			return rule_tasks.BLOCK
		}
	}

	return rule_tasks.PASS
}

func (r *Rule942) evaluateRawValue(tx *core.Transaction, evaluateFunc doEvaluateFunc, addition *Rule942Addition) int {
	for k, v := range tx.Variables.RequestCookies {
		if strings.Contains(k, "__utm") {
			continue
		}

		if evaluateFunc(tx, &k) {
			return rule_tasks.BLOCK
		}

		if evaluateFunc(tx, &v) {
			return rule_tasks.BLOCK
		}
	}

	for _, argMap := range tx.Variables.Args {
		for k, v := range *argMap {
			if evaluateFunc(tx, &k) {
				return rule_tasks.BLOCK
			}

			if evaluateFunc(tx, &v) {
				return rule_tasks.BLOCK
			}
		}
	}

	for _, v := range tx.Variables.XML["/*"] {
		if evaluateFunc(tx, &v) {
			return rule_tasks.BLOCK
		}
	}

	if addition.validateFileName {
		if evaluateFunc(tx, &tx.Variables.RequestFileName) {
			return rule_tasks.BLOCK
		}
	}

	if addition.validateReferer {
		referer := tx.Variables.RequestHeaders["referer"]
		if evaluateFunc(tx, &referer) {
			return rule_tasks.BLOCK
		}
	}

	if addition.validateUserAgent {
		ua := tx.Variables.RequestHeaders["user-agent"]
		if evaluateFunc(tx, &ua) {
			return rule_tasks.BLOCK
		}
	}

	if addition.validateBaseName {
		if evaluateFunc(tx, &tx.Variables.RequestBaseName) {
			return rule_tasks.BLOCK
		}
	}

	return rule_tasks.PASS
}

func (r *Rule942) doEvaluate(tx *core.Transaction, value *string) bool {

	// 一次加载全部，因此拿一个key来判断即可（这里以 cookie_key_942 为判断依据）
	_, ok := tx.Variables.TransMap["cookie_key_942"]
	if ok {
		return true
	}

	for k, v := range tx.Variables.RequestCookies {
		if strings.Contains(k, "__utm") {
			continue
		}
		tx.Variables.TransMap["cookie_key_942"] = append(tx.Variables.TransMap["cookie_key_942"], r.transform(&k))
		tx.Variables.TransMap["cookie_value_942"] = append(tx.Variables.TransMap["cookie_value_942"], r.transform(&v))
	}

	for _, argMap := range tx.Variables.Args {
		for k, v := range *argMap {
			tx.Variables.TransMap["arg_key_942"] = append(tx.Variables.TransMap["cookie_key_942"], r.transform(&k))
			tx.Variables.TransMap["arg_value_942"] = append(tx.Variables.TransMap["cookie_value_942"], r.transform(&v))
		}
	}

	for _, v := range tx.Variables.XML["/*"] {
		tx.Variables.TransMap["xml_942"] = append(tx.Variables.TransMap["xml_942"], r.transform(&v))
	}
	tx.Variables.TransMap["file_name_942"] = append(tx.Variables.TransMap["file_name_942"], r.transform(&tx.Variables.RequestFileName))
	ua, ok := tx.Variables.RequestHeaders["user-agent"]
	if ok {
		tx.Variables.TransMap["ua_942"] = append(tx.Variables.TransMap["ua_942"], r.transform(&ua))
	}
	referer, ok := tx.Variables.RequestHeaders["referer"]
	if ok {
		tx.Variables.TransMap["referer_942"] = append(tx.Variables.TransMap["referer_942"], r.transform(&referer))
	}

	tx.Variables.TransMap["base_name_942"] = append(tx.Variables.TransMap["base_name_942"], r.transform(&tx.Variables.RequestBaseName))

	return true
}

func (r *Rule942) transform(value *string) string {
	v, _, _ := core.UrlDecodeUni(*value)

	return v
}
