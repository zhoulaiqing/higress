package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/flier/gohs/hyperscan"
)

const PTN_931110 = `(?i)(?:\binclude\s*\([^)]*|mosConfig_absolute_path|_CONF\[path\]|_SERVER\[DOCUMENT_ROOT\]|GALLERY_BASEDIR|path\[docroot\]|appserv_root|config\[root_dir\])=(?:file|ftps?|https?)://`

type Rule931110 struct {
}

func (r *Rule931110) Id() string {
	return "931110"
}

func (r *Rule931110) Phase() int {
	return 2
}

func (r *Rule931110) Evaluate(tx core.Transaction) bool {
	if r.doEvaluate(tx, tx.Variables.QueryString) {
		return true
	}

	if r.doEvaluate(tx, tx.Variables.RequestBody) {
		return true
	}

	return true
}

func (r *Rule931110) doEvaluate(tx core.Transaction, value string) bool {
	v, _, _ := core.UrlDecodeUni(value)

	m, _ := hyperscan.MatchString(PTN_931110, v)
	if m {
		tx.Variables.RfiScore += core.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
	}

	return m
}
