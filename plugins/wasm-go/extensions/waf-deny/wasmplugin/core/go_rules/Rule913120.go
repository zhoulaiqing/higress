package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
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

func (r *Rule913120) Evaluate(tx core.Transaction) bool {
	matched, _ := core.PmEvaluate(rule913120Matcher, tx.Variables.RequestFileName, false)
	if matched {
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
		return true
	}

	for _, argMap := range tx.Variables.Args {
		for k, _ := range *argMap {
			m, _ := core.PmEvaluate(rule913120Matcher, k, false)
			if m {
				tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
				return true
			}
		}
	}

	return true
}

func init() {
	rule913120Matcher = core.AHO_CORASICK_BUILDER.Build(SCANNERS_URLS)
}
