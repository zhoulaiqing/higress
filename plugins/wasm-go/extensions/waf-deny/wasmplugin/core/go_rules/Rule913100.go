package core

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/flier/gohs/hyperscan"
	ahocorasick "github.com/petar-dambovaliev/aho-corasick"
)

var SCANNERS_USER_AGENTS = []string{
	"(hydra)",
	".nasl",
	"absinthe",
	"advanced email extractor",
	"arachni/",
	"autogetcontent",
	"bilbo",
	"BFAC",
	"brutus",
	"brutus/aet",
	"bsqlbf",
	"burpcollaborator",
	"cgichk",
	"cisco-torch",
	"commix",
	"core-project/1.0",
	"crimscanner/",
	"datacha0s",
	"Detectify",
	"dirbuster",
	"domino hunter",
	"dotdotpwn",
	"ecairn-grabber",
	"email extractor",
	"fhscan core 1.",
	"floodgate",
	"Fuzz Faster U Fool",
	"F-Secure Radar",
	"get-minimal",
	"gobuster",
	"gootkit auto-rooter scanner",
	"grabber",
	"grendel-scan",
	"havij",
	"httpx - Open-source project",
	"inspath",
	"internet ninja",
	"jaascois",
	"Jorgee",
	"masscan",
	"metis",
	"morfeus fucking scanner",
	"mysqloit",
	"n-stealth",
	"nessus",
	"netsparker",
	"nikto",
	"nmap nse",
	"nmap scripting engine",
	"nmap-nse",
	"nsauditor",
	"Nuclei",
	"openvas",
	"pangolin",
	"paros",
	"pmafind",
	"prog.customcrawler",
	"QQGameHall",
	"qualys was",
	"s.t.a.l.k.e.r.",
	"security scan",
	"springenwerk",
	"sql power injector",
	"sqlmap",
	"sqlninja",
	"struts-pwn",
	"sysscan",
	"TBI-WebScanner",
	"teh forest lobster",
	"this is an exploit",
	"toata dragostea",
	"toata dragostea mea pentru diavola",
	"uil2pn",
	"user-agent:",
	"vega/",
	"voideye",
	"w3af.sf.net",
	"w3af.sourceforge.net",
	"w3af.org",
	"webbandit",
	"webinspect",
	"webshag",
	"webtrends security analyzer",
	"webvulnscan",
	"Wfuzz",
	"whatweb",
	"whcc/",
	"wordpress hash grabber",
	"WPScan",
	"xmlrpc exploit",
	"zgrab",
	"zmeu",
}
var rule913100Matcher ahocorasick.AhoCorasick

type Rule913100 struct {
}

func (r *Rule913100) Id() string {
	return "913100"
}

func (r *Rule913100) Phase() int {
	return 1
}

func (r *Rule913100) Evaluate(tx core.Transaction) bool {
	matched, matchedValues := core.PmEvaluate(rule913100Matcher, tx.Variables.RequestHeaders["User-Agent"], true)
	if !matched {
		return true
	}

	existMatched := false
	for _, value := range matchedValues {
		m, _ := hyperscan.MatchString(`^(?:urlgrabber/[0-9\.]+ yum/[0-9\.]+|mozilla/[0-9\.]+ ecairn-grabber/[0-9\.]+ \(\+http://ecairn.com/grabber\))$`, value)
		if m {
			existMatched = true
			break
		}
	}
	if existMatched {
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
	}
	return true
}

func init() {
	//builder := ahocorasick.NewAhoCorasickBuilder(ahocorasick.Opts{
	//	AsciiCaseInsensitive: true,
	//	MatchOnlyWholeWords:  false,
	//	MatchKind:            ahocorasick.LeftMostLongestMatch,
	//	DFA:                  true,
	//})

	rule913100Matcher = core.AHO_CORASICK_BUILDER.Build(SCANNERS_USER_AGENTS)
}
