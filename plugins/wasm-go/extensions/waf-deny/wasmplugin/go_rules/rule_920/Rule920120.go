package rule_920

import (
	"github.com/tianchi/waf/wasmplugin/core"
	"github.com/tianchi/waf/wasmplugin/rule_tasks"
	"github.com/wasilibs/go-re2"
)

type Rule920120 struct {
}

const PTN_920120 = "(?i)^(?:&(?:(?:[acegiln-or-suz]acut|[aeiou]grav|[ain-o]tild)e|[c-elnr-tz]caron|(?:[cgk-lnr-t]cedi|[aeiouy]um)l|[aceg-josuwy]circ|[au]ring|a(?:mp|pos)|nbsp|oslash);|[^\\\"';=])*$"

var re920120 *re2.Regexp

func (r *Rule920120) Id() string {
	return "920120"
}

func (r *Rule920120) Phase() int {
	return 2
}

func (r *Rule920120) Evaluate(tx *core.Transaction) int {

	if len(tx.Variables.Files) == 0 {
		return rule_tasks.PASS
	}

	for fk, values := range tx.Variables.Files {
		m := re920120.MatchString(fk)
		if m {
			return rule_tasks.PASS
		}

		for _, v := range values {
			m = re920120.MatchString(v)
			if m {
				return rule_tasks.PASS
			}
		}
	}

	tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	return rule_tasks.BLOCK
}

func init() {
	re920120, _ = re2.Compile(PTN_920120)
}
