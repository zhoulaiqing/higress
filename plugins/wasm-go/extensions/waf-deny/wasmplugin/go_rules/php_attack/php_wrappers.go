
//line php_wrappers.rl:1
package php_attack

import (
    "fmt"
    "golang.org/x/exp/slices"
)

var wrappers = []string{
"bzip2", "expect", "glob", "ogg", "phar", "rar", "zip", "zlib",
"ssh2", "ssh2.shell", "ssh2.sftp", "ssh2.scp", "ssh2.exec", "ssh2.tunnel",
}

func checkWrapper(value string) bool {
    v := value[:len(value)-3]
    fmt.Printf("May be wrapper: %s \n", v)
    return slices.Contains(wrappers, v)
}

func matchPhpWrapper(data []byte) bool {

//line php_wrappers.rl:21

//line php_wrappers.go:26
const phpWrapper_start int = 5
const phpWrapper_first_final int = 5
const phpWrapper_error int = -1

const phpWrapper_en_main int = 5


//line php_wrappers.rl:22
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    var ts, te, act int
    _ = act

    
//line php_wrappers.go:42
	{
	cs = phpWrapper_start
	ts = 0
	te = 0
	act = 0
	}

//line php_wrappers.go:50
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 0:
		goto st_case_0
	case 1:
		goto st_case_1
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	}
	goto st_out
tr0:
//line php_wrappers.rl:40
p = (te) - 1

	goto st5
tr4:
//line php_wrappers.rl:34
te = p+1
{
                if checkWrapper(string(data[ts:te])) {
                    return true
                }
            }
	goto st5
tr7:
//line php_wrappers.rl:40
te = p+1

	goto st5
tr9:
//line php_wrappers.rl:40
te = p
p--

	goto st5
	st5:
//line NONE:1
ts = 0

		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line NONE:1
ts = p

//line php_wrappers.go:108
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr8
			}
		case data[p] >= 65:
			goto tr8
		}
		goto tr7
tr8:
//line NONE:1
te = p+1

	goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line php_wrappers.go:128
		switch data[p] {
		case 46:
			goto st0
		case 58:
			goto st2
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st4
			}
		case data[p] >= 65:
			goto st4
		}
		goto tr9
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		case data[p] >= 65:
			goto st1
		}
		goto tr0
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		if data[p] == 58 {
			goto st2
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		case data[p] >= 65:
			goto st1
		}
		goto tr0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if data[p] == 47 {
			goto st3
		}
		goto tr0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 47 {
			goto tr4
		}
		goto tr0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 46:
			goto st0
		case 58:
			goto st2
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st4
			}
		case data[p] >= 65:
			goto st4
		}
		goto tr0
	st_out:
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof0: cs = 0; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 6:
			goto tr9
		case 0:
			goto tr0
		case 1:
			goto tr0
		case 2:
			goto tr0
		case 3:
			goto tr0
		case 4:
			goto tr0
		}
	}

	}

//line php_wrappers.rl:45

        return false
}