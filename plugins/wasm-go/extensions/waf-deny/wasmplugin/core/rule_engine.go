package core

import (
	core "github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core/go_rules"
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
		&core.Rule911100{},
		&core.Rule913100{}, &core.Rule913120{},
		&core.Rule920100{}, &core.Rule920120{}, &core.Rule920160{}, &core.Rule920170{}, &core.Rule920180{},
		&core.Rule920181{}, &core.Rule920190{}, &core.Rule920210{}, &core.Rule920220{}, &core.Rule920240{},
		&core.Rule920250{}, &core.Rule920260{}, &core.Rule920270{}, &core.Rule920280{}, &core.Rule920310{},
		&core.Rule920330{}, &core.Rule920340{}, &core.Rule920350{}, &core.Rule920420{}, &core.Rule920430{},
		&core.Rule920440{}, &core.Rule920450{}, &core.Rule920470{}, &core.Rule920480{}, &core.Rule920500{},
		&core.Rule920520{}, &core.Rule920530{}, &core.Rule920540{}, &core.Rule920600{}, &core.Rule920610{},
		&core.Rule921110{}, &core.Rule921120{}, &core.Rule921130{}, &core.Rule921140{}, &core.Rule921150{},
		&core.Rule921160{}, &core.Rule921190{}, &core.Rule921200{}, &core.Rule921240{}, &core.Rule921421{},
		&core.Rule922100{}, &core.Rule922110_120{},
		&core.Rule930100{}, &core.Rule930110{}, &core.Rule930120{}, &core.Rule930130{},
		&core.Rule931100{}, &core.Rule931110{}, &core.Rule931120{},
		&core.Rule932230{}, &core.RuleFinal{},
	}
)
