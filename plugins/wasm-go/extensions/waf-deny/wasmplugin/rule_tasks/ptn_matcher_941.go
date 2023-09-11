package rule_tasks

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	ahocorasick "github.com/petar-dambovaliev/aho-corasick"
	"github.com/wasilibs/go-re2"
)

const (
	PTN_941110 = `(?i)<script[^>]*>[\s\S]*?`
	PTN_941130 = `(?i).(?:\b(?:x(?:link:href|html|mlns)|data:text/html|formaction|pattern\b.*?=)|!ENTITY[\s\v]+(?:%[\s\v]+)?[^\s\v]+[\s\v]+(?:SYSTEM|PUBLIC)|@import|;base64)\b`
	PTN_941140 = `(?i)[a-z]+=(?:[^:=]+:.+;)*?[^:=]+:url\(javascript`
	PTN_941160 = `(?i)<[^0-9<>A-Z_a-z]*(?:[^\s\v\"'<>]*:)?[^0-9<>A-Z_a-z]*[^0-9A-Z_a-z]*?(?:s[^0-9A-Z_a-z]*?(?:c[^0-9A-Z_a-z]*?r[^0-9A-Z_a-z]*?i[^0-9A-Z_a-z]*?p[^0-9A-Z_a-z]*?t|t[^0-9A-Z_a-z]*?y[^0-9A-Z_a-z]*?l[^0-9A-Z_a-z]*?e|v[^0-9A-Z_a-z]*?g|e[^0-9A-Z_a-z]*?t[^0-9>A-Z_a-z])|f[^0-9A-Z_a-z]*?o[^0-9A-Z_a-z]*?r[^0-9A-Z_a-z]*?m|m[^0-9A-Z_a-z]*?(?:a[^0-9A-Z_a-z]*?r[^0-9A-Z_a-z]*?q[^0-9A-Z_a-z]*?u[^0-9A-Z_a-z]*?e[^0-9A-Z_a-z]*?e|e[^0-9A-Z_a-z]*?t[^0-9A-Z_a-z]*?a[^0-9>A-Z_a-z])|(?:l[^0-9A-Z_a-z]*?i[^0-9A-Z_a-z]*?n[^0-9A-Z_a-z]*?k|o[^0-9A-Z_a-z]*?b[^0-9A-Z_a-z]*?j[^0-9A-Z_a-z]*?e[^0-9A-Z_a-z]*?c[^0-9A-Z_a-z]*?t|e[^0-9A-Z_a-z]*?m[^0-9A-Z_a-z]*?b[^0-9A-Z_a-z]*?e[^0-9A-Z_a-z]*?d|a[^0-9A-Z_a-z]*?(?:p[^0-9A-Z_a-z]*?p[^0-9A-Z_a-z]*?l[^0-9A-Z_a-z]*?e[^0-9A-Z_a-z]*?t|u[^0-9A-Z_a-z]*?d[^0-9A-Z_a-z]*?i[^0-9A-Z_a-z]*?o|n[^0-9A-Z_a-z]*?i[^0-9A-Z_a-z]*?m[^0-9A-Z_a-z]*?a[^0-9A-Z_a-z]*?t[^0-9A-Z_a-z]*?e)|p[^0-9A-Z_a-z]*?a[^0-9A-Z_a-z]*?r[^0-9A-Z_a-z]*?a[^0-9A-Z_a-z]*?m|i?[^0-9A-Z_a-z]*?f[^0-9A-Z_a-z]*?r[^0-9A-Z_a-z]*?a[^0-9A-Z_a-z]*?m[^0-9A-Z_a-z]*?e|b[^0-9A-Z_a-z]*?(?:a[^0-9A-Z_a-z]*?s[^0-9A-Z_a-z]*?e|o[^0-9A-Z_a-z]*?d[^0-9A-Z_a-z]*?y|i[^0-9A-Z_a-z]*?n[^0-9A-Z_a-z]*?d[^0-9A-Z_a-z]*?i[^0-9A-Z_a-z]*?n[^0-9A-Z_a-z]*?g[^0-9A-Z_a-z]*?s)|i[^0-9A-Z_a-z]*?m[^0-9A-Z_a-z]*?a?[^0-9A-Z_a-z]*?g[^0-9A-Z_a-z]*?e?|v[^0-9A-Z_a-z]*?i[^0-9A-Z_a-z]*?d[^0-9A-Z_a-z]*?e[^0-9A-Z_a-z]*?o)[^0-9>A-Z_a-z])|(?:<[0-9A-Z_a-z].*[\s\v/]|[\"'](?:.*[\s\v/])?)(?:background|formaction|lowsrc|on(?:a(?:bort|ctivate|d(?:apteradded|dtrack)|fter(?:print|(?:scriptexecu|upda)te)|lerting|n(?:imation(?:end|iteration|start)|tennastatechange)|ppcommand|udio(?:end|process|start))|b(?:e(?:fore(?:(?:(?:de)?activa|scriptexecu)te|c(?:opy|ut)|editfocus|p(?:aste|rint)|u(?:nload|pdate))|gin(?:Event)?)|l(?:ocked|ur)|oun(?:ce|dary)|roadcast|usy)|c(?:a(?:(?:ch|llschang)ed|nplay(?:through)?|rdstatechange)|(?:ell|fstate)change|h(?:a(?:rging(?:time)?cha)?nge|ecking)|l(?:ick|ose)|o(?:m(?:mand(?:update)?|p(?:lete|osition(?:end|start|update)))|n(?:nect(?:ed|ing)|t(?:extmenu|rolselect))|py)|u(?:echange|t))|d(?:ata(?:(?:availabl|chang)e|error|setc(?:hanged|omplete))|blclick|e(?:activate|livery(?:error|success)|vice(?:found|light|(?:mo|orienta)tion|proximity))|i(?:aling|s(?:abled|c(?:hargingtimechange|onnect(?:ed|ing))))|o(?:m(?:a(?:ctivate|ttrmodified)|(?:characterdata|subtree)modified|focus(?:in|out)|mousescroll|node(?:inserted(?:intodocument)?|removed(?:fromdocument)?))|wnloading)|r(?:ag(?:drop|e(?:n(?:d|ter)|xit)|(?:gestur|leav)e|over|start)|op)|urationchange)|e(?:mptied|n(?:abled|d(?:ed|Event)?|ter)|rror(?:update)?|xit)|f(?:ailed|i(?:lterchange|nish)|o(?:cus(?:in|out)?|rm(?:change|input)))|g(?:amepad(?:axismove|button(?:down|up)|(?:dis)?connected)|et)|h(?:ashchange|e(?:adphoneschange|l[dp])|olding)|i(?:cc(?:cardlockerror|infochange)|n(?:coming|put|valid))|key(?:down|press|up)|l(?:evelchange|o(?:ad(?:e(?:d(?:meta)?data|nd)|start)?|secapture)|y)|m(?:ark|essage|o(?:use(?:down|enter|(?:lea|mo)ve|o(?:ut|ver)|up|wheel)|ve(?:end|start)?|z(?:a(?:fterpaint|udioavailable)|(?:beforeresiz|orientationchang|t(?:apgestur|imechang))e|(?:edgeui(?:c(?:ancel|omplet)|start)e|network(?:down|up)loa)d|fullscreen(?:change|error)|m(?:agnifygesture(?:start|update)?|ouse(?:hittest|pixelscroll))|p(?:ointerlock(?:change|error)|resstapgesture)|rotategesture(?:start|update)?|s(?:crolledareachanged|wipegesture(?:end|start|update)?))))|no(?:match|update)|o(?:(?:bsolet|(?:ff|n)lin)e|pen|verflow(?:changed)?)|p(?:a(?:ge(?:hide|show)|int|(?:st|us)e)|lay(?:ing)?|op(?:state|up(?:hid(?:den|ing)|show(?:ing|n)))|ro(?:gress|pertychange))|r(?:atechange|e(?:adystatechange|ceived|movetrack|peat(?:Event)?|quest|s(?:et|ize|u(?:lt|m(?:e|ing)))|trieving)|ow(?:e(?:nter|xit)|s(?:delete|inserted)))|s(?:croll|e(?:ek(?:complete|ed|ing)|lect(?:start)?|n(?:ding|t)|t)|how|(?:ound|peech)(?:end|start)|t(?:a(?:lled|rt|t(?:echange|uschanged))|k(?:comma|sessione)nd|op)|u(?:bmit|ccess|spend)|vg(?:abort|error|(?:un)?load|resize|scroll|zoom))|t(?:ext|ime(?:out|update)|ouch(?:cancel|en(?:d|ter)|(?:lea|mo)ve|start)|ransition(?:cancel|end|run))|u(?:n(?:derflow|load)|p(?:dateready|gradeneeded)|s(?:erproximity|sdreceived))|v(?:ersion|o(?:ic|lum)e)change|w(?:a(?:it|rn)ing|heel)|zoom)|ping|s(?:rc|tyle))[\x08-\n\f-\r ]*?=`
	//PTN_941160   = `(?i)<[^0-9<>A-Z_a-z]*(?:[^\s\v\"'<>]*:)?[^0-9<>A-Z_a-z]*[^0-9A-Z_a-z]*?(?:s[^0-9A-Z_a-z]*?(?:c[^0-9A-Z_a-z]*?r[^0-9A-Z_a-z]*?i[^0-9A-Z_a-z]*?p[^0-9A-Z_a-z]*?t|t[^0-9A-Z_a-z]*?y[^0-9A-Z_a-z]*?l[^0-9A-Z_a-z]*?e|v[^0-9A-Z_a-z]*?g|e[^0-9A-Z_a-z]*?t[^0-9>A-Z_a-z])|f[^0-9A-Z_a-z]*?o[^0-9A-Z_a-z]*?r[^0-9A-Z_a-z]*?m|m[^0-9A-Z_a-z]*?(?:a[^0-9A-Z_a-z]*?r[^0-9A-Z_a-z]*?q[^0-9A-Z_a-z]*?u[^0-9A-Z_a-z]*?e[^0-9A-Z_a-z]*?e|e[^0-9A-Z_a-z]*?t[^0-9A-Z_a-z]*?a[^0-9>A-Z_a-z])|(?:l[^0-9A-Z_a-z]*?i[^0-9A-Z_a-z]*?n[^0-9A-Z_a-z]*?k|o[^0-9A-Z_a-z]*?b[^0-9A-Z_a-z]*?j[^0-9A-Z_a-z]*?e[^0-9A-Z_a-z]*?c[^0-9A-Z_a-z]*?t|e[^0-9A-Z_a-z]*?m[^0-9A-Z_a-z]*?b[^0-9A-Z_a-z]*?e[^0-9A-Z_a-z]*?d|a[^0-9A-Z_a-z]*?(?:p[^0-9A-Z_a-z]*?p[^0-9A-Z_a-z]*?l[^0-9A-Z_a-z]*?e[^0-9A-Z_a-z]*?t|u[^0-9A-Z_a-z]*?d[^0-9A-Z_a-z]*?i[^0-9A-Z_a-z]*?o|n[^0-9A-Z_a-z]*?i[^0-9A-Z_a-z]*?m[^0-9A-Z_a-z]*?a[^0-9A-Z_a-z]*?t[^0-9A-Z_a-z]*?e)|p[^0-9A-Z_a-z]*?a[^0-9A-Z_a-z]*?r[^0-9A-Z_a-z]*?a[^0-9A-Z_a-z]*?m|i?[^0-9A-Z_a-z]*?f[^0-9A-Z_a-z]*?r[^0-9A-Z_a-z]*?a[^0-9A-Z_a-z]*?m[^0-9A-Z_a-z]*?e|b[^0-9A-Z_a-z]*?(?:a[^0-9A-Z_a-z]*?s[^0-9A-Z_a-z]*?e|o[^0-9A-Z_a-z]*?d[^0-9A-Z_a-z]*?y|i[^0-9A-Z_a-z]*?n[^0-9A-Z_a-z]*?d[^0-9A-Z_a-z]*?i[^0-9A-Z_a-z]*?n[^0-9A-Z_a-z]*?g[^0-9A-Z_a-z]*?s)|i[^0-9A-Z_a-z]*?m[^0-9A-Z_a-z]*?a?[^0-9A-Z_a-z]*?g[^0-9A-Z_a-z]*?e?|v[^0-9A-Z_a-z]*?i[^0-9A-Z_a-z]*?d[^0-9A-Z_a-z]*?e[^0-9A-Z_a-z]*?o)[^0-9>A-Z_a-z])`
	PTN_941170   = `(?i)(?:\W|^)(?:javascript:(?:[\s\S]+[=\x5c\(\[\.<]|[\s\S]*?(?:\bname\b|\x5c[ux]\d))|data:(?:(?:[a-z]\w+/\w[\w+-]+\w)?[;,]|[\s\S]*?;[\s\S]*?\b(?:base64|charset=)|[\s\S]*?,[\s\S]*?<[\s\S]*?\w[\s\S]*?>))|@\W*?i\W*?m\W*?p\W*?o\W*?r\W*?t\W*?(?:/\*[\s\S]*?)?(?:[\"']|\W*?u\W*?r\W*?l[\s\S]*?\()|[^-]*?-\W*?m\W*?o\W*?z\W*?-\W*?b\W*?i\W*?n\W*?d\W*?i\W*?n\W*?g[^:]*?:\W*?u\W*?r\W*?l[\s\S]*?\(`
	PTN_941190   = `(?i:<style.*?>.*?(?:@[i\x5c]|(?:[:=]|&#x?0*(?:58|3A|61|3D);?).*?(?:[(\x5c]|&#x?0*(?:40|28|92|5C);?)))`
	PTN_941200   = `(?i:<.*[:]?vmlframe.*?[\s/+]*?src[\s/+]*=)`
	PTN_941210   = `(?i:(?:j|&#x?0*(?:74|4A|106|6A);?)(?:\t|\n|\r|&(?:#x?0*(?:9|13|10|A|D);?|tab;|newline;))*(?:a|&#x?0*(?:65|41|97|61);?)(?:\t|\n|\r|&(?:#x?0*(?:9|13|10|A|D);?|tab;|newline;))*(?:v|&#x?0*(?:86|56|118|76);?)(?:\t|\n|\r|&(?:#x?0*(?:9|13|10|A|D);?|tab;|newline;))*(?:a|&#x?0*(?:65|41|97|61);?)(?:\t|\n|\r|&(?:#x?0*(?:9|13|10|A|D);?|tab;|newline;))*(?:s|&#x?0*(?:83|53|115|73);?)(?:\t|\n|\r|&(?:#x?0*(?:9|13|10|A|D);?|tab;|newline;))*(?:c|&#x?0*(?:67|43|99|63);?)(?:\t|\n|\r|&(?:#x?0*(?:9|13|10|A|D);?|tab;|newline;))*(?:r|&#x?0*(?:82|52|114|72);?)(?:\t|\n|\r|&(?:#x?0*(?:9|13|10|A|D);?|tab;|newline;))*(?:i|&#x?0*(?:73|49|105|69);?)(?:\t|\n|\r|&(?:#x?0*(?:9|13|10|A|D);?|tab;|newline;))*(?:p|&#x?0*(?:80|50|112|70);?)(?:\t|\n|\r|&(?:#x?0*(?:9|13|10|A|D);?|tab;|newline;))*(?:t|&#x?0*(?:84|54|116|74);?)(?:\t|\n|\r|&(?:#x?0*(?:9|13|10|A|D);?|tab;|newline;))*(?::|&(?:#x?0*(?:58|3A);?|colon;)).)`
	PTN_941220   = `(?i:(?:v|&#x?0*(?:86|56|118|76);?)(?:\t|&(?:#x?0*(?:9|13|10|A|D);?|tab;|newline;))*(?:b|&#x?0*(?:66|42|98|62);?)(?:\t|&(?:#x?0*(?:9|13|10|A|D);?|tab;|newline;))*(?:s|&#x?0*(?:83|53|115|73);?)(?:\t|&(?:#x?0*(?:9|13|10|A|D);?|tab;|newline;))*(?:c|&#x?0*(?:67|43|99|63);?)(?:\t|&(?:#x?0*(?:9|13|10|A|D);?|tab;|newline;))*(?:r|&#x?0*(?:82|52|114|72);?)(?:\t|&(?:#x?0*(?:9|13|10|A|D);?|tab;|newline;))*(?:i|&#x?0*(?:73|49|105|69);?)(?:\t|&(?:#x?0*(?:9|13|10|A|D);?|tab;|newline;))*(?:p|&#x?0*(?:80|50|112|70);?)(?:\t|&(?:#x?0*(?:9|13|10|A|D);?|tab;|newline;))*(?:t|&#x?0*(?:84|54|116|74);?)(?:\t|&(?:#x?0*(?:9|13|10|A|D);?|tab;|newline;))*(?::|&(?:#x?0*(?:58|3A);?|colon;)).)`
	PTN_941230   = `(?i)<EMBED[\s/+].*?(?:src|type).*?=`
	PTN_941240   = `<[?]?import[\s/+\S]*?implementation[\s/+]*?=`
	PTN_941250   = "(?i:<META[\\s/+].*?http-equiv[\\s/+]*=[\\s/+]*[\\\"'`]?(?:(?:c|&#x?0*(?:67|43|99|63);?)|(?:r|&#x?0*(?:82|52|114|72);?)|(?:s|&#x?0*(?:83|53|115|73);?)))"
	PTN_941260   = `(?i:<META[\s/+].*?charset[\s/+]*=)`
	PTN_941270   = `(?i)<LINK[\s/+].*?href[\s/+]*=`
	PTN_941280   = `(?i)<BASE[\s/+].*?href[\s/+]*=`
	PTN_941290   = `(?i)<APPLET[\s/+>]`
	PTN_941300   = `(?i)<OBJECT[\s/+].*?(?:type|codetype|classid|code|data)[\s/+]*=`
	PTN_941310_1 = `\xbc[^\xbe>]*[\xbe>]|<[^\xbe]*\xbe`
	PTN_941310_2 = `(?:\xbc\s*/\s*[^\xbe>]*[\xbe>])|(?:<\s*/\s*[^\xbe]*\xbe)`
	PTN_941350   = `\+ADw-.*(?:\+AD4-|>)|<.*\+AD4-`
	PTN_941360   = `![!+ ]\[\]`
	PTN_941370   = `(?:self|document|this|top|window)\s*(?:/\*|[\[)]).+?(?:\]|\*/)`
	PTN_941390   = `(?i)\b(?:eval|set(?:timeout|interval)|new[\s\v]+Function|a(?:lert|tob)|btoa)[\s\v]*\(`
	PTN_941400   = "((?:\\[[^\\]]*\\][^.]*\\.)|Reflect[^.]*\\.).*(?:map|sort|apply)[^.]*\\..*call[^`]*`.*`"
)

var Re941110, Re941130, Re941140, Re941160, Re941170, Re941190, Re941200, Re941210, Re941220, Re941230, Re941240, Re941250,
	Re941260, Re941270, Re941280, Re941290, Re941300, Re9413101, Re9413102, Re941350, Re941360, Re941370, Re941390, Re941400 *re2.Regexp

var ValidateByteRange941010 *core.ValidateByteRange

var DENY_KEYWORDS_941180 = []string{
	"document.cookie", "document.domain", "document.write", ".parentnode", ".innerhtml", "window.location", "-moz-binding", "<!--", "<![cdata[",
}
var Rule941180Matcher ahocorasick.AhoCorasick

//riskTags := []string {
//	"script", "style", "svg", "set", "form", "marquee", "meta", "link", "object", "embed", "applet", "audio", "animate",
//	"param", "iframe", "frame", "base", "body", "bindings", "image", "img", "video"
//}

//var riskAttrs = []string{
//	"background", "formaction", "lowsrc", "ping", "src", "style",
//}

func init() {
	ValidateByteRange941010, _ = core.NewValidateByRange("20, 45-47, 48-57, 65-90, 95, 97-122")
	Rule941180Matcher = AHO_CORASICK_BUILDER.Build(DENY_KEYWORDS_941180)

	Re941110 = re2.MustCompile(PTN_941110)
	Re941130 = re2.MustCompile(PTN_941130)
	Re941140 = re2.MustCompile(PTN_941140)
	Re941160 = re2.MustCompile(PTN_941160)
	Re941170 = re2.MustCompile(PTN_941170)
	Re941190 = re2.MustCompile(PTN_941190)
	Re941200 = re2.MustCompile(PTN_941200)
	Re941210 = re2.MustCompile(PTN_941210)
	Re941220 = re2.MustCompile(PTN_941220)
	Re941230 = re2.MustCompile(PTN_941230)
	Re941240 = re2.MustCompile(PTN_941240)
	Re941250 = re2.MustCompile(PTN_941250)
	Re941260 = re2.MustCompile(PTN_941260)
	Re941270 = re2.MustCompile(PTN_941270)
	Re941280 = re2.MustCompile(PTN_941280)
	Re941290 = re2.MustCompile(PTN_941290)
	Re941300 = re2.MustCompile(PTN_941300)
	Re9413101 = re2.MustCompile(PTN_941310_1)
	Re9413102 = re2.MustCompile(PTN_941310_2)
	Re941350 = re2.MustCompile(PTN_941350)
	Re941360 = re2.MustCompile(PTN_941360)
	Re941370 = re2.MustCompile(PTN_941370)
	Re941390 = re2.MustCompile(PTN_941390)
	Re941400 = re2.MustCompile(PTN_941400)
}
