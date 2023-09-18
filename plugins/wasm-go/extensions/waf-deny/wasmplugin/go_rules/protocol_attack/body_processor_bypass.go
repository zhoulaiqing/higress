
//line body_processor_bypass.rl:1
package protocol_attack

func matchBodyProcessorBypass(data []byte) bool {

//line body_processor_bypass.rl:5

//line body_processor_bypass.go:10
const bodyProcessorBypass_start int = 1
const bodyProcessorBypass_first_final int = 19
const bodyProcessorBypass_error int = 0

const bodyProcessorBypass_en_main int = 1


//line body_processor_bypass.rl:6
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    
//line body_processor_bypass.go:23
	{
	cs = bodyProcessorBypass_start
	}

//line body_processor_bypass.go:28
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
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
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 19:
		goto st_case_19
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 0:
		goto st_case_0
	}
	goto st_out
	st_case_1:
		switch data[p] {
		case 32:
			goto st0
		case 44:
			goto st0
		case 59:
			goto st0
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st0
		}
		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 32:
			goto st3
		case 44:
			goto st3
		case 59:
			goto st3
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st3
		}
		goto st2
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 97:
			goto st4
		case 116:
			goto st13
		}
		goto st3
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 97:
			goto st4
		case 112:
			goto st5
		case 116:
			goto st13
		}
		goto st3
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 97:
			goto st4
		case 112:
			goto st6
		case 116:
			goto st13
		}
		goto st3
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch data[p] {
		case 97:
			goto st4
		case 108:
			goto st7
		case 116:
			goto st13
		}
		goto st3
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 97:
			goto st4
		case 105:
			goto st8
		case 116:
			goto st13
		}
		goto st3
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		switch data[p] {
		case 97:
			goto st4
		case 99:
			goto st9
		case 116:
			goto st13
		}
		goto st3
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		switch data[p] {
		case 97:
			goto st10
		case 116:
			goto st13
		}
		goto st3
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch data[p] {
		case 97:
			goto st4
		case 112:
			goto st5
		case 116:
			goto st11
		}
		goto st3
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch data[p] {
		case 97:
			goto st4
		case 101:
			goto st12
		case 105:
			goto st16
		case 116:
			goto st13
		}
		goto st3
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch data[p] {
		case 97:
			goto st4
		case 116:
			goto st13
		case 120:
			goto st14
		}
		goto st3
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		switch data[p] {
		case 97:
			goto st4
		case 101:
			goto st12
		case 116:
			goto st13
		}
		goto st3
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		switch data[p] {
		case 97:
			goto st4
		case 116:
			goto st15
		}
		goto st3
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		switch data[p] {
		case 47:
			goto tr16
		case 97:
			goto st4
		case 101:
			goto st12
		case 116:
			goto st13
		}
		goto st3
tr16:
//line body_processor_bypass.rl:11

            return true
        
	goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line body_processor_bypass.go:297
		switch data[p] {
		case 97:
			goto st4
		case 116:
			goto st13
		}
		goto st3
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		switch data[p] {
		case 97:
			goto st4
		case 111:
			goto st17
		case 116:
			goto st13
		}
		goto st3
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		switch data[p] {
		case 97:
			goto st4
		case 110:
			goto st18
		case 116:
			goto st13
		}
		goto st3
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		switch data[p] {
		case 47:
			goto tr16
		case 97:
			goto st4
		case 116:
			goto st13
		}
		goto st3
st_case_0:
	st0:
		cs = 0
		goto _out
	st_out:
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof

	_test_eof: {}
	_out: {}
	}

//line body_processor_bypass.rl:22

        return false
}

