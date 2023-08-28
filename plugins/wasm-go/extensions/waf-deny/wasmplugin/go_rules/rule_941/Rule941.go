package rule_941

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"strings"
)

type Rule941 struct {
}

func (r *Rule941) Id() string {
	return "941"
}

func (r *Rule941) Phase() int {
	return 2
}

func (r *Rule941) Evaluate(tx *core.Transaction) int {
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
