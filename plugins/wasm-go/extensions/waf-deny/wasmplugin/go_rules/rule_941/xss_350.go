
//line xss_350.rl:1
package rule_941

func matchXSS350(data []byte) bool {

//line xss_350.rl:5

//line xss_350.go:10
const xss_start int = 4
const xss_first_final int = 4
const xss_error int = -1

const xss_en_main int = 4


//line xss_350.rl:6
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    // -1. No start; 0. Normal start; 1. Dangerous start
    startTag := -1

    var ts, te, act int
    _, _, _ = ts, te, act

    
//line xss_350.go:29
	{
	cs = xss_start
	ts = 0
	te = 0
	act = 0
	}

//line xss_350.go:37
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 0:
		goto st_case_0
	case 1:
		goto st_case_1
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	}
	goto st_out
tr0:
//line xss_350.rl:37
p = (te) - 1

	goto st4
tr4:
//line xss_350.rl:25
te = p+1
{
                if startTag >= 0 {
                    return true
                }
            }
	goto st4
tr5:
//line xss_350.rl:17
te = p+1
{
                startTag = 1
            }
	goto st4
tr6:
//line xss_350.rl:37
te = p+1

	goto st4
tr8:
//line xss_350.rl:21
te = p+1
{
                startTag = 0
            }
	goto st4
tr9:
//line xss_350.rl:31
te = p+1
{
                if startTag > 0 {
                    return true
                }
            }
	goto st4
tr10:
//line xss_350.rl:37
te = p
p--

	goto st4
	st4:
//line NONE:1
ts = 0

		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line NONE:1
ts = p

//line xss_350.go:116
		switch data[p] {
		case 43:
			goto tr7
		case 60:
			goto tr8
		case 62:
			goto tr9
		}
		goto tr6
tr7:
//line NONE:1
te = p+1

	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line xss_350.go:136
		if data[p] == 97 {
			goto st0
		}
		goto tr10
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		if data[p] == 100 {
			goto st1
		}
		goto tr0
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch data[p] {
		case 52:
			goto st2
		case 119:
			goto st3
		}
		goto tr0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if data[p] == 45 {
			goto tr4
		}
		goto tr0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 45 {
			goto tr5
		}
		goto tr0
	st_out:
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof0: cs = 0; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 5:
			goto tr10
		case 0:
			goto tr0
		case 1:
			goto tr0
		case 2:
			goto tr0
		case 3:
			goto tr0
		}
	}

	}

//line xss_350.rl:43


    return false
}
