package rule_tasks

import ahocorasick "github.com/petar-dambovaliev/aho-corasick"

var JAVA_CODE_LEAKAGES = []string{
	"<jsp:",
	"javax.servlet",
	".addheader",
	".createtextfile",
	".getfile",
	".loadfromfile",
	"response.binarywrite",
	"response.write",
	"scripting.filesystemobject",
	"server.createobject",
	"server.execute",
	"server.htmlencode",
	"server.mappath",
	"server.urlencode",
	"vbscript.encode",
	"wscript.network",
	"wscript.shell",
}
var Rule952100Matcher ahocorasick.AhoCorasick

var JAVA_ERRORS = []string{
	"[java.lang.",
	"class java.lang.",
	"java.lang.NullPointerException",
	"java.rmi.ServerException",
	"at java.lang.",
	"onclick=\"toggle('full exception chain stacktrace')\"",
	"at org.apache.catalina",
	"at org.apache.coyote.",
	"at org.apache.tomcat.",
	"at org.apache.jasper.",
}
var Rule952110Matcher ahocorasick.AhoCorasick

func init() {
	Rule952100Matcher = AHO_CORASICK_BUILDER.Build(JAVA_CODE_LEAKAGES)
	Rule952110Matcher = AHO_CORASICK_BUILDER.Build(JAVA_ERRORS)
}
