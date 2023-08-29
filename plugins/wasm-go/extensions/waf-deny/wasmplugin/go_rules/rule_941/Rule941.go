package rule_941

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"strings"
)

type Rule941 struct {
}

type Rule941Addition struct {
	validateFileName  bool
	validateUserAgent bool
	validateReferer   bool
}

var noAddition941 = &Rule941Addition{
	validateFileName:  false,
	validateUserAgent: false,
	validateReferer:   false,
}

var additionWithoutReferer = &Rule941Addition{
	validateFileName:  true,
	validateUserAgent: true,
	validateReferer:   false,
}

var fullAddition = &Rule941Addition{
	validateFileName:  true,
	validateUserAgent: true,
	validateReferer:   true,
}

var fileNameAddition = &Rule941Addition{
	validateFileName:  true,
	validateUserAgent: false,
	validateReferer:   false,
}

func (r *Rule941) Id() string {
	return "941"
}

func (r *Rule941) Phase() int {
	return 2
}

func (r *Rule941) Evaluate(tx *core.Transaction) int {
	return r.evaluateByCache(tx)
}

func (r *Rule941) GetAddition() *Rule941Addition {
	return noAddition941
}

func (r *Rule941) evaluateByCache(tx *core.Transaction) int {
	cookieKeyCache := tx.Variables.TransMap["cookie_key_941"]
	cookieValueCache := tx.Variables.TransMap["cookie_value_941"]
	for i, _ := range cookieKeyCache {
		if r.doEvaluate(tx, &cookieKeyCache[i]) {
			return rule_tasks.BLOCK
		}

		if r.doEvaluate(tx, &cookieValueCache[i]) {
			return rule_tasks.BLOCK
		}
	}

	argsKeyCache := tx.Variables.TransMap["arg_key_941"]
	argsValueCache := tx.Variables.TransMap["arg_value_941"]
	for i, _ := range argsKeyCache {
		if r.doEvaluate(tx, &argsKeyCache[i]) {
			return rule_tasks.BLOCK
		}

		if r.doEvaluate(tx, &argsValueCache[i]) {
			return rule_tasks.BLOCK
		}
	}

	xmlCache := tx.Variables.TransMap["xml_941"]
	for i, _ := range xmlCache {
		if r.doEvaluate(tx, &xmlCache[i]) {
			return rule_tasks.BLOCK
		}
	}

	addition := r.GetAddition()
	if addition.validateFileName {
		fileNameCache := tx.Variables.TransMap["file_name_941"]
		if len(fileNameCache) > 0 && r.doEvaluate(tx, &fileNameCache[0]) {
			return rule_tasks.BLOCK
		}
	}

	if addition.validateReferer {
		refererCache := tx.Variables.TransMap["referer_941"]
		if len(refererCache) > 0 && r.doEvaluate(tx, &refererCache[0]) {
			return rule_tasks.BLOCK
		}
	}

	if addition.validateUserAgent {
		uaCache := tx.Variables.TransMap["ua_941"]
		if len(uaCache) > 0 && r.doEvaluate(tx, &uaCache[0]) {
			return rule_tasks.BLOCK
		}
	}

	return rule_tasks.PASS
}

func (r *Rule941) evaluateRawValue(tx *core.Transaction) int {
	for k, v := range tx.Variables.RequestCookies {
		if strings.Contains(k, "__utm") {
			continue
		}

		if r.doEvaluate(tx, &k) {
			return rule_tasks.BLOCK
		}

		if r.doEvaluate(tx, &v) {
			return rule_tasks.BLOCK
		}
	}

	for _, argMap := range tx.Variables.Args {
		for k, v := range *argMap {
			if r.doEvaluate(tx, &k) {
				return rule_tasks.BLOCK
			}

			if r.doEvaluate(tx, &v) {
				return rule_tasks.BLOCK
			}
		}
	}

	for _, v := range tx.Variables.XML["/*"] {
		if r.doEvaluate(tx, &v) {
			return rule_tasks.BLOCK
		}
	}

	addition := r.GetAddition()
	if addition.validateFileName {
		if r.doEvaluate(tx, &tx.Variables.RequestFileName) {
			return rule_tasks.BLOCK
		}
	}

	if addition.validateReferer {
		referer := tx.Variables.RequestHeaders["referer"]
		if r.doEvaluate(tx, &referer) {
			return rule_tasks.BLOCK
		}
	}

	if addition.validateUserAgent {
		ua := tx.Variables.RequestHeaders["user-agent"]
		if r.doEvaluate(tx, &ua) {
			return rule_tasks.BLOCK
		}
	}

	return rule_tasks.PASS
}

func (r *Rule941) doEvaluate(tx *core.Transaction, value *string) bool {

	// 一次加载全部，因此拿一个key来判断即可（这里以 cookie_key_941 为判断依据）
	_, ok := tx.Variables.TransMap["cookie_key_941"]
	if ok {
		return true
	}

	for k, v := range tx.Variables.RequestCookies {
		if strings.Contains(k, "__utm") {
			continue
		}
		tx.Variables.TransMap["cookie_key_941"] = append(tx.Variables.TransMap["cookie_key_941"], r.transform(&k))
		tx.Variables.TransMap["cookie_value_941"] = append(tx.Variables.TransMap["cookie_value_941"], r.transform(&v))
	}

	for _, argMap := range tx.Variables.Args {
		for k, v := range *argMap {
			tx.Variables.TransMap["arg_key_941"] = append(tx.Variables.TransMap["cookie_key_941"], r.transform(&k))
			tx.Variables.TransMap["arg_value_941"] = append(tx.Variables.TransMap["cookie_value_941"], r.transform(&v))
		}
	}

	for _, v := range tx.Variables.XML["/*"] {
		tx.Variables.TransMap["xml_941"] = append(tx.Variables.TransMap["xml_941"], r.transform(&v))
	}
	tx.Variables.TransMap["file_name_941"] = append(tx.Variables.TransMap["file_name_941"], r.transform(&tx.Variables.RequestFileName))
	ua, ok := tx.Variables.RequestHeaders["user-agent"]
	if ok {
		tx.Variables.TransMap["ua_941"] = append(tx.Variables.TransMap["ua_941"], r.transform(&ua))
	}
	referer, ok := tx.Variables.RequestHeaders["referer"]
	if ok {
		tx.Variables.TransMap["referer_941"] = append(tx.Variables.TransMap["referer_941"], r.transform(&referer))
	}

	return true
}

func (r *Rule941) transform(value *string) string {
	v, _, _ := core.Utf8ToUnicode(*value)
	v, _, _ = core.UrlDecodeUni(v)
	v, _, _ = core.HtmlEntityDecode(v)
	v, _, _ = core.JsDecode(v)
	v, _, _ = core.CssDecode(v)
	v, _, _ = core.RemoveNulls(v)

	return v
}
