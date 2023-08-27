package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	ahocorasick "github.com/petar-dambovaliev/aho-corasick"
)

const (
	INBOUND_ANOMALY_SCORE_THRESHOLD  = 5
	OUTBOUND_ANOMALY_SCORE_THRESHOLD = 4
	REPORTING_LEVEL                  = 4
	EARLY_BLOCKING                   = 0
	BLOCKING_PARANONIA_LEVEL         = 1
	CRITICAL_ANOMALY_SCORE           = 5
	WARNING_ANOMALY_SCORE            = 3
	NOTICE_ANOMALY_SCORE             = 2
	ENFORCE_BODYPROC_URLENCODED      = 0
	CRS_VALIDATE_UTF8_ENCODING       = 0
)

var (
	ALLOWED_METHODS              = []string{"GET", "HEAD", "OPTIONS", "POST"}
	ALLOWED_REQUEST_CONTENT_TYPE = []string{
		"application/x-www-form-urlencoded",
		"multipart/form-data",
		"multipart/related",
		"text/xml",
		"application/xml",
		"application/soap+xml",
		"application/json",
		"application/cloudevents+json",
		"application/cloudevents-batch+json",
	}
	ALLOWED_REQUEST_CONTENT_TYPE_CHARSET = []string{"utf-8", "iso-8859", "iso-8859-15", "windows-1252"}
	ALLOWED_HTTP_VERSIONS                = []string{"HTTP/1.0", "HTTP/1.1", "HTTP/2", "HTTP/2.0"}
	RESTRICTED_EXTENSIONS                = []string{
		".asa", ".asax", ".ascx",
		".backup", ".bak", ".bat",
		".cdx", ".cer", ".cfg", ".cmd", ".com", ".config", ".conf", ".cs", ".csproj", ".csr",
		".dat", ".db", ".dbf", ".dll", ".dos",
		".htr", ".htw",
		".ida", ".idc", ".idq", ".inc", ".ini",
		".key",
		".licx", ".lnk", ".log",
		".mdb",
		".old",
		".pass", ".pdb", ".pol", ".printer", ".pwd", ".rdb",
		".resources", ".resx",
		".sql", ".swp", ".sys", ".vb", ".vbs", ".vbproj", ".vsdisco",
		".xsd", ".xsx",
	}
	RESTRICTED_HEADERS = []string{"accept-charset", "content-encoding", "proxy", "lock-token", "content-range", "if",
		"x-http-method-override", "x-http-method", "x-method-override"}

	AHO_CORASICK_BUILDER = ahocorasick.NewAhoCorasickBuilder(ahocorasick.Opts{
		AsciiCaseInsensitive: true,
		MatchOnlyWholeWords:  false,
		MatchKind:            ahocorasick.LeftMostLongestMatch,
		DFA:                  true,
	})

	RULES = []Rule{
		&Rule911100{},
		&Rule913100{}, &Rule913120{},
		&Rule920100{}, &Rule920120{}, &Rule920160{}, &Rule920170{}, &Rule920180{},
		&Rule920181{}, &Rule920190{}, &Rule920210{}, &Rule920220{}, &Rule920240{},
		&Rule920250{}, &Rule920260{}, &Rule920270{}, &Rule920280{}, &Rule920310{},
		&Rule920330{}, &Rule920340{}, &Rule920350{}, &Rule920420{}, &Rule920430{},
		&Rule920440{}, &Rule920450{}, &Rule920470{}, &Rule920480{}, &Rule920500{},
		&Rule920520{}, &Rule920530{}, &Rule920540{}, &Rule920600{}, &Rule920610{},
		&Rule921110{}, &Rule921120{}, &Rule921130{}, &Rule921140{}, &Rule921150{},
		&Rule921160{}, &Rule921190{}, &Rule921200{}, &Rule921240{}, &Rule921421{},
		&Rule922100{}, &Rule922110_120{},
		&Rule930100{}, &Rule930110{}, &Rule930120{}, &Rule930130{},
		&Rule931100{}, &Rule931110{}, &Rule931120{},
		&Rule932230{}, &RuleFinal{},
	}
)

type RuleEngine struct {
}

func ProcessRequestHeaderRules(tx *core.Transaction) bool {
	for _, rule := range RULES {
		if rule.Phase() == 1 || rule.Phase() == 100 {
			r := rule.Evaluate(tx)
			if !r {
				return false
			}
		}
	}

	return true
}
