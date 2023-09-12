
//line xss_360.rl:1
package rule_941

func matchXSS360(data []byte) bool {

//line xss_360.rl:5

//line xss_360.go:10
const xss_start int = 0
const xss_first_final int = 5
const xss_error int = -1

const xss_en_main int = 0


//line xss_360.rl:6
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof
    
//line xss_360.go:22
	{
	cs = xss_start
	}

//line xss_360.go:27
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 0:
		goto st_case_0
	case 1:
		goto st_case_1
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 4:
		goto st_case_4
	}
	goto st_out
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		if data[p] == 33 {
			goto st1
		}
		goto st0
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch data[p] {
		case 32:
			goto st2
		case 33:
			goto st4
		case 43:
			goto st2
		}
		goto st0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 33:
			goto st1
		case 91:
			goto st3
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 33:
			goto st1
		case 93:
			goto st5
		}
		goto st0
tr11:
//line xss_360.rl:9

            return true
        
	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line xss_360.go:115
		if data[p] == 33 {
			goto tr7
		}
		goto tr6
tr6:
//line xss_360.rl:9

            return true
        
	goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line xss_360.go:131
		if data[p] == 33 {
			goto tr7
		}
		goto tr6
tr7:
//line xss_360.rl:9

            return true
        
	goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line xss_360.go:147
		switch data[p] {
		case 32:
			goto tr8
		case 33:
			goto tr9
		case 43:
			goto tr8
		}
		goto tr6
tr8:
//line xss_360.rl:9

            return true
        
	goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line xss_360.go:168
		switch data[p] {
		case 33:
			goto tr7
		case 91:
			goto tr10
		}
		goto tr6
tr10:
//line xss_360.rl:9

            return true
        
	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line xss_360.go:187
		switch data[p] {
		case 33:
			goto tr7
		case 93:
			goto tr11
		}
		goto tr6
tr9:
//line xss_360.rl:9

            return true
        
	goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line xss_360.go:206
		switch data[p] {
		case 32:
			goto tr8
		case 33:
			goto tr9
		case 43:
			goto tr8
		case 91:
			goto tr10
		}
		goto tr6
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 32:
			goto st2
		case 33:
			goto st4
		case 43:
			goto st2
		case 91:
			goto st3
		}
		goto st0
	st_out:
	_test_eof0: cs = 0; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 5:
//line xss_360.rl:9

            return true
        
//line xss_360.go:255
		}
	}

	}

//line xss_360.rl:18


    return false
}
