
//line xss_200.rl:1
package rule_941


func matchXSS200(data []byte) bool {

//line xss_200.rl:6

//line xss_200.go:11
const xss_start int = 10
const xss_first_final int = 10
const xss_error int = -1

const xss_en_main int = 10


//line xss_200.rl:7
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    var ts, te, act int
        _, _, _ = ts, te, act

    step := 0

    
//line xss_200.go:29
	{
	cs = xss_start
	ts = 0
	te = 0
	act = 0
	}

//line xss_200.go:37
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
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
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 12:
		goto st_case_12
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	}
	goto st_out
tr0:
//line NONE:1
	switch act {
	case 1:
	{p = (te) - 1

                step = 1
            }
	default:
	{p = (te) - 1
}
	}
	
	goto st10
tr10:
//line xss_200.rl:33
p = (te) - 1

	goto st10
tr12:
//line xss_200.rl:25
te = p+1
{
                if step == 1 {
                    return true
                } else {
                    step = 0
                }
            }
	goto st10
tr13:
//line xss_200.rl:33
te = p+1

	goto st10
tr16:
//line xss_200.rl:33
te = p
p--

	goto st10
	st10:
//line NONE:1
ts = 0

		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line NONE:1
ts = p

//line xss_200.go:123
		switch data[p] {
		case 60:
			goto tr14
		case 115:
			goto tr15
		}
		goto tr13
tr9:
//line NONE:1
te = p+1

//line xss_200.rl:21
act = 1;
	goto st11
tr14:
//line NONE:1
te = p+1

//line xss_200.rl:33
act = 3;
	goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line xss_200.go:150
		if data[p] == 118 {
			goto st1
		}
		goto st0
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		if data[p] == 118 {
			goto st1
		}
		goto st0
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch data[p] {
		case 109:
			goto st2
		case 118:
			goto st1
		}
		goto st0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 108:
			goto st3
		case 118:
			goto st1
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 102:
			goto st4
		case 118:
			goto st1
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 114:
			goto st5
		case 118:
			goto st1
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 97:
			goto st6
		case 118:
			goto st1
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch data[p] {
		case 109:
			goto st7
		case 118:
			goto st1
		}
		goto st0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 101:
			goto tr9
		case 118:
			goto st1
		}
		goto st0
tr15:
//line NONE:1
te = p+1

	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line xss_200.go:258
		if data[p] == 114 {
			goto st8
		}
		goto tr16
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		if data[p] == 99 {
			goto st9
		}
		goto tr10
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		switch data[p] {
		case 32:
			goto st9
		case 43:
			goto st9
		case 47:
			goto st9
		case 61:
			goto tr12
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st9
		}
		goto tr10
	st_out:
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof0: cs = 0; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 11:
			goto tr0
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
		case 5:
			goto tr0
		case 6:
			goto tr0
		case 7:
			goto tr0
		case 12:
			goto tr16
		case 8:
			goto tr10
		case 9:
			goto tr10
		}
	}

	}

//line xss_200.rl:38


    return false
}
