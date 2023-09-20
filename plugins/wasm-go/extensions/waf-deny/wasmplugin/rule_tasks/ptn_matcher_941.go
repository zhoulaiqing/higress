package rule_tasks

import (
	"github.com/tianchi/waf/wasmplugin/core"
	ahocorasick "github.com/wasilibs/go-aho-corasick"
)

var ValidateByteRange941010 *core.ValidateByteRange

var DENY_KEYWORDS_941180 = []string{
	"document.cookie", "document.domain", "document.write", ".parentnode", ".innerhtml", "window.location", "-moz-binding", "<!--", "<![cdata[",
}
var Rule941180Matcher ahocorasick.AhoCorasick

func init() {
	ValidateByteRange941010, _ = core.NewValidateByRange("20, 45-47, 48-57, 65-90, 95, 97-122")
	Rule941180Matcher = AHO_CORASICK_BUILDER.Build(DENY_KEYWORDS_941180)
}
