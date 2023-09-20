package go_rules

import (
	"github.com/tianchi/waf/wasmplugin/core"
	"github.com/tianchi/waf/wasmplugin/rule_tasks"
	"github.com/wasilibs/go-re2"
)

const PTN_931110 = `(?i)(?:\binclude\s*\([^)]*|mosConfig_absolute_path|_CONF\[path\]|_SERVER\[DOCUMENT_ROOT\]|GALLERY_BASEDIR|path\[docroot\]|appserv_root|config\[root_dir\])=(?:file|ftps?|https?)://`

var re931110 *re2.Regexp

type Rule931110 struct {
}

func (r *Rule931110) Id() string {
	return "931110"
}

func (r *Rule931110) Phase() int {
	return 2
}

func (r *Rule931110) Evaluate(tx *core.Transaction) int {
	if r.doEvaluate(tx, tx.Variables.QueryString) {
		return rule_tasks.BLOCK
	}

	if r.doEvaluate(tx, tx.Variables.RequestBody) {
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}

func (r *Rule931110) doEvaluate(tx *core.Transaction, value string) bool {

	m := re931110.MatchString(value)
	if m {
		tx.Variables.RfiScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}

func init() {
	re931110, _ = re2.Compile(PTN_931110)
}
