package protocol_attack

import (
	"reflect"
	"testing"
)

func TestProtocolAttackDefaultRisk(t *testing.T) {
	type testCase struct {
		value string
		want  bool //期望的计算结果
	}

	testGroup := map[string]testCase{
		"case921110-1":  {"%0aPOST / HTTP/1.0", true},
		"case921110-2":  {"aaa%0aGET+/+HTTP/1.1", true},
		"case921110-3":  {"aaa%0dHEAD+http://example.com/+HTTP/1.1", true},
		"case921110-4":  {"aaa%0d%0aGet+/foo%0d", false},
		"case921110-5":  {"aaa%0d%0aGet+foo+bar", false},
		"case921110-6":  {"barGET /a.html HTTP/1.1\\r\\nSomething: GET /b.html HTTP/1.1\\r\\nHost: foo.com\\r\\nUser-Agent: foo\\r\\nAccept: */*\\r\\n\\r\\n", true},
		"case921110-7":  {"GET%20http%3A%2F%2Fwww.foo.bar%20HTTP%2F1.2", true},
		"case921110-8":  {"GET%20http%3A%2F%2Fwww.foo.bar%20HTTP%2F3.2", true},
		"case921110-9":  {"soundtrack Gympl\\r\\nanything", false},
		"case921110-10": {"budget foo)</bar>\\n", false},
		"case921110-11": {"get it\\n", false},

		"case921120-1": {"foobar%0d%0aContent-Length:%200%0d%0a%0d%0aHTTP/1.1%20200%20OK%0d%0aContent-Type:%20text/html%0d%0aContent-Length:%2019%0d%0a%0d%0a<html>Shazam</html>", true},
		"case921120-2": {"foobar%0d%0aContent-Length:%2002343432423<html>ftw</html>", true},
		"case921120-3": {"%0A%0Dlocation:%0A%0D", false},
		"case921120-4": {"%0d%0aContent-Length: 0", true},

		"case921130-1": {"/?lang=foobar%3Cmeta%20http-equiv%3D%22Refresh%22%20content%3D%220%3B%20url%3Dhttp%3A%2F%2Fwww.hacker.com%2F%22%3E", true},
		"case921130-2": {"oreo=munchmuch%0d%0a%0d%0a<HTML><title></title></HTML>", true},
		"case921130-3": {"/?arg1=GET%20http%3A%2F%2Fwww.foo.bar%20HTTP%2F1.2", true},
		"case921130-4": {"/?arg1=GET%20http%3A%2F%2Fwww.foo.bar%20HTTP%2F3.2", true},

		"case921200-1":  {"(%26(objectCategory=computer)  (userAccountControl:1.2.840.113556.1.4.803:=8192))", false},
		"case921200-2":  {"(objectSID=S-1-5-21-73586283-152049171-839522115-1111)", false},
		"case921200-3":  {"(userAccountControl:1.2.840.113556.1.4.803:=67108864)(%26(objectCategory=group)(groupType:1.2.840.113556.1.4.803:=2147483648))", false},
		"case921200-4":  {"bar)(%26)", true},
		"case921120-5":  {"printer)(uid=*)", true},
		"case921120-6":  {"void)(objectClass=users))(%26(objectClass=void)", true},
		"case921120-7":  {"eb9adbd87d)!(sn=*", true},
		"case921120-8":  {"*)!(sn=*", true},
		"case921120-9":  {"*)(uid=*))(|(uid=*", true},
		"case921120-10": {"aaa*aaa)(cn>=bob)", true},
	}

	for key, v := range testGroup { //遍历
		t.Run(key, func(t *testing.T) {
			tv := transformDefault(v.value)
			got := matchDefaultRisk(tv)
			want := v.want
			if !reflect.DeepEqual(want, got) { //比较
				t.Errorf("excepted:%v, got:%v", want, got)
			}
		})
	}
}

func TestArgGet(t *testing.T) {
	type testCase struct {
		value string
		want  bool //期望的计算结果
	}

	testGroup := map[string]testCase{
		"case921160-1": {"Y&%0d%0Remote-addr%0d%0d%0d:%20foo.bar.com", true},
		"case921160-2": {"Y&%0d%0aRemote-addr%0d%0d%0d:%20foo.bar.com", true},
		"case921160-3": {"Y&%0d%0aRemote-addr:%20foo.bar.com", true},
		"case921160-4": {"%0d%0aRemote-addr:%20foo.bar.com", true},
		"case921160-5": {"bar&%0d%0aRemote-addr:%20foo.bar.com=Y", true},
	}

	for key, v := range testGroup { //遍历
		t.Run(key, func(t *testing.T) {
			tv := transformDefault(v.value)
			got := matchArgGet(tv)
			want := v.want
			if !reflect.DeepEqual(want, got) { //比较
				t.Errorf("excepted:%v, got:%v", want, got)
			}
		})
	}
}

func BenchmarkMatchLdapRagel_test(b *testing.B) {
	testCases := [][]byte{
		[]byte("(%26(objectCategory=computer)  (userAccountControl:1.2.840.113556.1.4.803:=8192))"),
		[]byte("(objectSID=S-1-5-21-73586283-152049171-839522115-1111)"),
		[]byte("(userAccountControl:1.2.840.113556.1.4.803:=67108864)(%26(objectCategory=group)(groupType:1.2.840.113556.1.4.803:=2147483648))"),
		[]byte("bar)(%26)"),
		[]byte("printer)(uid=*)"),
		[]byte("void)(objectClass=users))(%26(objectClass=void)"),
		[]byte("eb9adbd87d)!(sn=*"),
		[]byte("*)!(sn=*"),
		[]byte("*)(uid=*))(|(uid=*"),
		[]byte("aaa*aaa)(cn>=bob)"),
	}

	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			matchLdapInjection(testCase)
		}
	}
}

func BenchmarkMatchLdapRegex_test(b *testing.B) {
	testCases := []string{
		"(%26(objectCategory=computer)  (userAccountControl:1.2.840.113556.1.4.803:=8192))",
		"(objectSID=S-1-5-21-73586283-152049171-839522115-1111)",
		"(userAccountControl:1.2.840.113556.1.4.803:=67108864)(%26(objectCategory=group)(groupType:1.2.840.113556.1.4.803:=2147483648))",
		"bar)(%26)",
		"printer)(uid=*)",
		"void)(objectClass=users))(%26(objectClass=void)",
		"eb9adbd87d)!(sn=*",
		"*)!(sn=*",
		"*)(uid=*))(|(uid=*",
		"aaa*aaa)(cn>=bob)",
	}

	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			re921200.MatchString(testCase)
		}
	}
}

//
//func BenchmarkMatchLdapRegex2_test(b *testing.B) {
//	testCases := []string{
//		"(%26(objectCategory=computer)  (userAccountControl:1.2.840.113556.1.4.803:=8192))",
//		"(objectSID=S-1-5-21-73586283-152049171-839522115-1111)",
//		"(userAccountControl:1.2.840.113556.1.4.803:=67108864)(%26(objectCategory=group)(groupType:1.2.840.113556.1.4.803:=2147483648))",
//		"bar)(%26)",
//		"printer)(uid=*)",
//		"void)(objectClass=users))(%26(objectClass=void)",
//		"eb9adbd87d)!(sn=*",
//		"*)!(sn=*",
//		"*)(uid=*))(|(uid=*",
//		"aaa*aaa)(cn>=bob)",
//	}
//
//	for i := 0; i < b.N; i++ {
//		for _, testCase := range testCases {
//			hyperscan.MatchString(`^[^:\(\)\&\|\!\<\>\~]*\)\s*(?:\((?:[^,\(\)\=\&\|\!\<\>\~]+[><~]?=|\s*[&!|]\s*(?:\)|\()?\s*)|\)\s*\(\s*[\&\|\!]\s*|[&!|]\s*\([^\(\)\=\&\|\!\<\>\~]+[><~]?=[^:\(\)\&\|\!\<\>\~]*)`, testCase)
//		}
//	}
//}
