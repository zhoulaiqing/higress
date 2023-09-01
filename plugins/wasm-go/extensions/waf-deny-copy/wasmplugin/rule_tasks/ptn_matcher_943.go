package rule_tasks

import "github.com/wasilibs/go-re2"

const (
	PTN_943100   = `(?i:\.cookie\b.*?;\W*?(?:expires|domain)\W*?=|\bhttp-equiv\W+set-cookie\b)`
	PTN_943110_1 = `^(?:jsessionid|aspsessionid|asp\.net_sessionid|phpsession|phpsessid|weblogicsession|session_id|session-id|cfid|cftoken|cfsid|jservsession|jwsession)$`
	PTN_943110_2 = `^(?:ht|f)tps?://(.*?)/`
)

var Re943100, Re9431101, Re9431102 *re2.Regexp

func init() {
	Re943100, _ = re2.Compile(PTN_943100)
	Re9431101, _ = re2.Compile(PTN_943110_1)
	Re9431102, _ = re2.Compile(PTN_943110_2)
}
