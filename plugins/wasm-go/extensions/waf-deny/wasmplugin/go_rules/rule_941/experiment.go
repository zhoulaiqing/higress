
//line experiment.rl:1
package rule_941

func matchExp(data []byte) bool {

//line experiment.rl:5

//line experiment.go:10
const xss_start int = 1
const xss_first_final int = 4
const xss_error int = 0

const xss_en_main int = 1


//line experiment.rl:6
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof
    
//line experiment.go:22
	{
	cs = xss_start
	}

//line experiment.go:27
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 1:
		goto st_case_1
	case 0:
		goto st_case_0
	case 4:
		goto st_case_4
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 5:
		goto st_case_5
	}
	goto st_out
	st_case_1:
		switch data[p] {
		case 34:
			goto st4
		case 39:
			goto st4
		case 60:
			goto st3
		}
		goto st0
st_case_0:
	st0:
		cs = 0
		goto _out
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 32:
			goto st4
		case 47:
			goto st4
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st4
		}
		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 32:
			goto st4
		case 47:
			goto st4
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st4
		}
		goto st2
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 32:
			goto st5
		case 47:
			goto st5
		case 95:
			goto st3
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st5
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st3
				}
			case data[p] >= 65:
				goto st3
			}
		default:
			goto st3
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		goto st0
	st_out:
	_test_eof4: cs = 4; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof

	_test_eof: {}
	_out: {}
	}

//line experiment.rl:27

        return false
}
