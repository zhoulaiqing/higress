package rule_tasks

import (
	"github.com/wasilibs/go-re2"
)

const (
	PTN_942140 = `(?i)\b(?:d(?:atabas|b_nam)e[^0-9A-Z_a-z]*\(|(?:information_schema|m(?:aster\.\.sysdatabases|s(?:db|ys(?:ac(?:cess(?:objects|storage|xml)|es)|modules2?|(?:object|querie|relationship)s))|ysql\.db)|northwind|pg_(?:catalog|toast)|tempdb)\b|s(?:chema(?:_name\b|[^0-9A-Z_a-z]*\()|(?:qlite_(?:temp_)?master|ys(?:aux|\.database_name))\b))`
	PTN_942151 = `(?i)\b(?:a(?:dd(?:dat|tim)e|es_(?:de|en)crypt|s(?:cii(?:str)?|in)|tan2?)|b(?:enchmark|i(?:n_to_num|t_(?:and|count|length|x?or)))|c(?:har(?:acter)?_length|iel(?:ing)?|o(?:alesce|ercibility|llation|(?:mpres)?s|n(?:cat(?:_ws)?|nection_id|v(?:ert(?:_tz)?)?)|t)|r32|ur(?:(?:dat|tim)e|rent_(?:date|time(?:stamp)?|user)))|d(?:a(?:t(?:abase|e(?:_(?:add|format|sub)|diff))|y(?:name|of(?:month|week|year)))|count|e(?:code|grees|s_(?:de|en)crypt)|ump)|e(?:lt|n(?:c(?:ode|rypt)|ds_?with)|x(?:p(?:ort_set)?|tract(?:value)?))|f(?:i(?:el|n)d_in_set|ound_rows|rom_(?:base64|days|unixtime))|g(?:e(?:ometrycollection|t_(?:format|lock))|(?:r(?:eates|oup_conca)|tid_subse)t)|hex(?:toraw)?|i(?:fnull|n(?:et6?_(?:aton|ntoa)|s(?:ert|tr)|terval)|s(?:_(?:(?:free|used)_lock|ipv(?:4(?:_(?:compat|mapped))?|6)|n(?:ot(?:_null)?|ull))|null))|json(?:_(?:a(?:gg|rray(?:_(?:elements(?:_text)?|length))?)|build_(?:array|object)|e(?:ac|xtract_pat)h(?:_text)?|object(?:_(?:agg|keys))?|populate_record(?:set)?|strip_nulls|t(?:o_record(?:set)?|ypeof))|b(?:_(?:array(?:_(?:elements(?:_text)?|length))?|build_(?:array|object)|object(?:_(?:agg|keys))?|e(?:ac|xtract_pat)h(?:_text)?|insert|p(?:ath_(?:(?:exists|match)(?:_tz)?|query(?:_(?:(?:array|first)(?:_tz)?|tz))?)|opulate_record(?:set)?|retty)|s(?:et(?:_lax)?|trip_nulls)|t(?:o_record(?:set)?|ypeof)))?|path)?|l(?:ast_(?:day|inser_id)|case|e(?:as|f)t|i(?:kel(?:ihood|y)|nestring)|o(?:ad_file|ca(?:ltimestamp|te)|g(?:10|2)|wer)|pad|trim)|m(?:a(?:ke(?:_set|date)|ster_pos_wait)|d5|i(?:crosecon)?d|onthname|ulti(?:linestring|po(?:int|lygon)))|n(?:ame_const|ot_in|ullif)|o(?:ct(?:et_length)?|(?:ld_passwo)?rd)|p(?:eriod_(?:add|diff)|g_(?:client_encoding|sleep)|o(?:(?:lyg|siti)on|w)|rocedure_analyse)|qu(?:arter|ote)|r(?:a(?:dians|nd|wtohex)|elease_lock|ow_(?:count|to_json)|pad|trim)|s(?:chema|e(?:c_to_time|ssion_user)|ha[1-2]?|in|oundex|pace|q(?:lite_(?:compileoption_(?:get|used)|source_id)|rt)|t(?:arts_?with|d(?:dev_(?:po|sam)p)?|r(?:_to_date|cmp))|ub(?:(?:dat|tim)e|str(?:ing(?:_index)?)?)|ys(?:date|tem_user))|t(?:ime(?:_(?:format|to_sec)|diff|stamp(?:add|diff)?)|o(?:_(?:base64|jsonb?)|n?char|(?:day|second)s)|r(?:im|uncate))|u(?:case|n(?:compress(?:ed_length)?|hex|i(?:str|x_timestamp)|likely)|(?:pdatexm|se_json_nul)l|tc_(?:date|time(?:stamp)?)|uid(?:_short)?)|var(?:_(?:po|sam)p|iance)|we(?:ek(?:day|ofyear)|ight_string)|xmltype|yearweek)[^0-9A-Z_a-z]*\(`
	PTN_942160 = `(?i:sleep\(\s*?\d*?\s*?\)|benchmark\(.*?\,.*?\))`
	PTN_942170 = `(?i)(?:select|;)[\s\v]+(?:benchmark|if|sleep)[\s\v]*?\([\s\v]*?\(?[\s\v]*?[0-9A-Z_a-z]+`
	PTN_942190 = "(?i)[\\\"'`](?:[\\s\\v]*![\\s\\v]*[\\\"'0-9A-Z_-z]|;?[\\s\\v]*(?:having|select|union\\b[\\s\\v]*(?:all|(?:distin|sele)ct))\\b[\\s\\v]*[^\\s\\v])|\\b(?:(?:(?:c(?:onnection_id|urrent_user)|database|schema|user)[\\s\\v]*?|select.*?[0-9A-Z_a-z]?user)\\(|exec(?:ute)?[\\s\\v]+master\\.|from[^0-9A-Z_a-z]+information_schema[^0-9A-Z_a-z]|into[\\s\\v\\+]+(?:dump|out)file[\\s\\v]*?[\\\"'`]|union(?:[\\s\\v]select[\\s\\v]@|[\\s\\v\\(0-9A-Z_a-z]*?select))|[\\s\\v]*?exec(?:ute)?.*?[^0-9A-Z_a-z]xp_cmdshell|[^0-9A-Z_a-z]iif[\\s\\v]*?\\("
	PTN_942220 = `^(?i:-0000023456|4294967295|4294967296|2147483648|2147483647|0000012345|-2147483648|-2147483649|0000023456|2.2250738585072007e-308|2.2250738585072011e-308|1e309)$`
	PTN_942230 = `(?i)[\s\v\(-\)]case[\s\v]+when.*?then|\)[\s\v]*?like[\s\v]*?\(|select.*?having[\s\v]*?[^\s\v]+[\s\v]*?[^\s\v0-9A-Z_a-z]|if[\s\v]?\([0-9A-Z_a-z]+[\s\v]*?[<->~]`
	PTN_942240 = "(?i)alter[\\s\\v]*?[0-9A-Z_a-z]+.*?char(?:acter)?[\\s\\v]+set[\\s\\v]+[0-9A-Z_a-z]+|[\\\"'`](?:;*?[\\s\\v]*?waitfor[\\s\\v]+(?:time|delay)[\\s\\v]+[\\\"'`]|;.*?:[\\s\\v]*?goto)"
	PTN_942250 = "(?i:merge.*?using\\s*?\\(|execute\\s*?immediate\\s*?[\\\"'`]|match\\s*?[\\w(),+-]+\\s*?against\\s*?\\()"
	PTN_942270 = `(?i)union.*?select.*?from`
	PTN_942280 = "(?i)select[\\s\\v]*?pg_sleep|waitfor[\\s\\v]*?delay[\\s\\v]?[\\\"'`]+[\\s\\v]?[0-9]|;[\\s\\v]*?shutdown[\\s\\v]*?(?:[#;\\{]|/\\*|--)"
	PTN_942290 = `(?i)\[?\$(?:n(?:e|in?|o[rt])|e(?:q|xists|lemMatch)|l(?:te?|ike)|mod|a(?:ll|nd)|(?:s(?:iz|lic)|wher)e|t(?:ype|ext)|x?or|div|between|regex|jsonSchema)\]?`
	PTN_942320 = `(?i)create[\s\v]+(?:function|procedure)[\s\v]*?[0-9A-Z_a-z]+[\s\v]*?\([\s\v]*?\)[\s\v]*?-|d(?:eclare[^0-9A-Z_a-z]+[#@][\s\v]*?[0-9A-Z_a-z]+|iv[\s\v]*?\([\+\-]*[\s\v\.0-9]+,[\+\-]*[\s\v\.0-9]+\))|exec[\s\v]*?\([\s\v]*?@|(?:lo_(?:impor|ge)t|procedure[\s\v]+analyse)[\s\v]*?\(|;[\s\v]*?(?:declare|open)[\s\v]+[\-0-9A-Z_a-z]+|::(?:b(?:igint|ool)|double[\s\v]+precision|int(?:eger)?|numeric|oid|real|(?:tex|smallin)t)`
	PTN_942350 = `(?i)create[\s\v]+function[\s\v].+[\s\v]returns|;[\s\v]*?(?:alter|(?:(?:cre|trunc|upd)at|renam)e|d(?:e(?:lete|sc)|rop)|(?:inser|selec)t|load)\b[\s\v]*?[\(\[]?[0-9A-Z_a-z]{2,}`
	PTN_942360 = "(?i)\\b(?:(?:alter|(?:(?:cre|trunc|upd)at|renam)e|de(?:lete|sc)|(?:inser|selec)t|load)[\\s\\v]+(?:char|group_concat|load_file)\\b[\\s\\v]*\\(?|end[\\s\\v]*?\\);)|[\\s\\v\\(]load_file[\\s\\v]*?\\(|[\\\"'`][\\s\\v]+regexp[^0-9A-Z_a-z]|[\\\"'0-9A-Z_-z][\\s\\v]+as\\b[\\s\\v]*[\\\"'0-9A-Z_-z]+[\\s\\v]*\\bfrom|^[^A-Z_a-z]+[\\s\\v]*?(?:(?:(?:(?:cre|trunc)at|renam)e|d(?:e(?:lete|sc)|rop)|(?:inser|selec)t|load)[\\s\\v]+[0-9A-Z_a-z]+|u(?:pdate[\\s\\v]+[0-9A-Z_a-z]+|nion[\\s\\v]*(?:all|(?:sele|distin)ct)\\b)|alter[\\s\\v]*(?:a(?:(?:ggregat|pplication[\\s\\v]*rol)e|s(?:sembl|ymmetric[\\s\\v]*ke)y|u(?:dit|thorization)|vailability[\\s\\v]*group)|b(?:roker[\\s\\v]*priority|ufferpool)|c(?:ertificate|luster|o(?:l(?:latio|um)|nversio)n|r(?:edential|yptographic[\\s\\v]*provider))|d(?:atabase|efault|i(?:mension|skgroup)|omain)|e(?:(?:ndpoi|ve)nt|xte(?:nsion|rnal))|f(?:lashback|oreign|u(?:lltext|nction))|hi(?:erarchy|stogram)|group|in(?:dex(?:type)?|memory|stance)|java|l(?:a(?:ngua|r)ge|ibrary|o(?:ckdown|g(?:file[\\s\\v]*group|in)))|m(?:a(?:s(?:k|ter[\\s\\v]*key)|terialized)|e(?:ssage[\\s\\v]*type|thod)|odule)|(?:nicknam|queu)e|o(?:perator|utline)|p(?:a(?:ckage|rtition)|ermission|ro(?:cedur|fil)e)|r(?:e(?:mot|sourc)e|o(?:l(?:e|lback)|ute))|s(?:chema|e(?:arch|curity|rv(?:er|ice)|quence|ssion)|y(?:mmetric[\\s\\v]*key|nonym)|togroup)|t(?:able(?:space)?|ext|hreshold|r(?:igger|usted)|ype)|us(?:age|er)|view|w(?:ork(?:load)?|rapper)|x(?:ml[\\s\\v]*schema|srobject))\\b)"
	PTN_942500 = `(?i:/\*[!+](?:[\w\s=_\-()]+)?\*/)`
	PTN_942540 = "^(?:[^']*'|[^\\\"]*\\\"|[^`]*`)[\\s\\v]*;"
	PTN_942550 = "[\\\"'`][\\[\\{].*[\\]\\}][\\\"'`].*(::.*jsonb?)?.*(?:(?:@|->?)>|<@|\\?[&\\|]?|#>>?|[<>]|<-)|(?:(?:@|->?)>|<@|\\?[&\\|]?|#>>?|[<>]|<-)[\\\"'`][\\[\\{].*[\\]\\}][\\\"'`]|json_extract.*\\(.*\\)"
)

var Re942140, Re942151, Re942160, Re942170, Re942190, Re942220, Re942230, Re942240, Re942250, Re942270,
	Re942280, Re942290, Re942320, Re942350, Re942360, Re942500, Re942540, Re942550 *re2.Regexp

func init() {
	Re942140, _ = re2.Compile(PTN_942140)
	Re942151, _ = re2.Compile(PTN_942151)
	Re942160, _ = re2.Compile(PTN_942160)
	Re942170, _ = re2.Compile(PTN_942170)
	Re942190, _ = re2.Compile(PTN_942190)
	Re942220, _ = re2.Compile(PTN_942220)
	Re942230, _ = re2.Compile(PTN_942230)
	Re942240, _ = re2.Compile(PTN_942240)
	Re942250, _ = re2.Compile(PTN_942250)
	Re942270, _ = re2.Compile(PTN_942270)
	Re942280, _ = re2.Compile(PTN_942280)
	Re942290, _ = re2.Compile(PTN_942290)
	Re942320, _ = re2.Compile(PTN_942320)
	Re942350, _ = re2.Compile(PTN_942350)
	Re942360, _ = re2.Compile(PTN_942360)
	Re942500, _ = re2.Compile(PTN_942500)
	Re942540, _ = re2.Compile(PTN_942540)
	Re942550, _ = re2.Compile(PTN_942550)
}
