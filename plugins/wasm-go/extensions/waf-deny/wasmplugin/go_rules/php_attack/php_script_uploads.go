
//line php_script_uploads.rl:1
package php_attack

func matchPhpScriptUploads(data []byte) bool {

//line php_script_uploads.rl:5

//line php_script_uploads.go:10
const phpScriptUploads_start int = 4
const phpScriptUploads_first_final int = 4
const phpScriptUploads_error int = -1

const phpScriptUploads_en_main int = 4


//line php_script_uploads.rl:6
    cs, p, pe, eof := 0, 0, len(data), len(data)

    var ts, te, act int
    _, _, _ = ts, te, act

    
//line php_script_uploads.go:25
	{
	cs = phpScriptUploads_start
	ts = 0
	te = 0
	act = 0
	}

//line php_script_uploads.go:33
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
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 3:
		goto st_case_3
	}
	goto st_out
tr0:
//line php_script_uploads.rl:28
p = (te) - 1

	goto st4
tr6:
//line php_script_uploads.rl:24
p = (te) - 1
{
                return true
            }
	goto st4
tr7:
//line php_script_uploads.rl:28
te = p+1

	goto st4
tr9:
//line php_script_uploads.rl:28
te = p
p--

	goto st4
tr11:
//line php_script_uploads.rl:24
te = p
p--
{
                return true
            }
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

//line php_script_uploads.go:103
		if data[p] == 46 {
			goto tr8
		}
		goto tr7
tr8:
//line NONE:1
te = p+1

	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line php_script_uploads.go:118
		if data[p] == 112 {
			goto st0
		}
		goto tr9
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		if data[p] == 104 {
			goto st1
		}
		goto tr0
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch data[p] {
		case 97:
			goto st2
		case 112:
			goto st7
		case 116:
			goto tr4
		}
		goto tr0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if data[p] == 114 {
			goto st6
		}
		goto tr0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		if data[p] == 46 {
			goto st6
		}
		goto tr11
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		if data[p] == 46 {
			goto st6
		}
		switch {
		case data[p] > 57:
			if 115 <= data[p] && data[p] <= 116 {
				goto st6
			}
		case data[p] >= 48:
			goto st8
		}
		goto tr11
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		if data[p] == 46 {
			goto st6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st8
		}
		goto tr11
tr4:
//line NONE:1
te = p+1

	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line php_script_uploads.go:203
		switch data[p] {
		case 46:
			goto st6
		case 109:
			goto st3
		}
		goto tr11
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 108 {
			goto st6
		}
		goto tr6
	st_out:
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof0: cs = 0; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 5:
			goto tr9
		case 0:
			goto tr0
		case 1:
			goto tr0
		case 2:
			goto tr0
		case 6:
			goto tr11
		case 7:
			goto tr11
		case 8:
			goto tr11
		case 9:
			goto tr11
		case 3:
			goto tr6
		}
	}

	}

//line php_script_uploads.rl:33

        return false
}