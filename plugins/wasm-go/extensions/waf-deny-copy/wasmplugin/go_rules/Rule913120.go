package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	ahocorasick "github.com/petar-dambovaliev/aho-corasick"
)

var SCANNERS_URLS = []string{
	"/.adSensepostnottherenonobook",
	"/<invalid>hello.html",
	"/actSensepostnottherenonotive",
	"/acunetix-wvs-test-for-some-inexistent-file",
	"/antidisestablishmentarianism",
	"/appscan_fingerprint/mac_address",
	"/arachni-",
	"/cybercop",
	"/nessus_is_probing_you_",
	"/nessustest",
	"/netsparker-",
	"/rfiinc.txt",
	"/thereisnowaythat-you-canbethere",
	"/w3af/remotefileinclude.html",
	"appscan_fingerprint",
	"w00tw00t.at.ISC.SANS.DFind",
	"w00tw00t.at.blackhats.romanian.anti-sec",
}
var rule913120Matcher ahocorasick.AhoCorasick

type Rule913120 struct {
}

func (r *Rule913120) Id() string {
	return "913120"
}

func (r *Rule913120) Phase() int {
	return 2
}

func (r *Rule913120) Evaluate(tx *core.Transaction) int {
	matched, _ := core.PmEvaluate(rule913120Matcher, tx.Variables.RequestFileName, false)
	if matched {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	for _, argMap := range tx.Variables.Args {
		for k, _ := range *argMap {
			m, _ := core.PmEvaluate(rule913120Matcher, k, false)
			if m {
				tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
				return rule_tasks.BLOCK
			}
		}
	}

	return rule_tasks.PASS
}

func init() {
	rule913120Matcher = rule_tasks.AHO_CORASICK_BUILDER.Build(SCANNERS_URLS)
}
