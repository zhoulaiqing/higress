
//line php_injection.rl:1
package php_attack

func matchPhpInjection(data []byte) bool {

//line php_injection.rl:5

//line php_injection.go:10
const phpInjection_start int = 66
const phpInjection_first_final int = 66
const phpInjection_error int = -1

const phpInjection_en_main int = 66


//line php_injection.rl:6
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    var ts, te, act int
    _, _, _ = ts, te, act

    
//line php_injection.go:26
	{
	cs = phpInjection_start
	ts = 0
	te = 0
	act = 0
	}

//line php_injection.go:34
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 66:
		goto st_case_66
	case 67:
		goto st_case_67
	case 68:
		goto st_case_68
	case 69:
		goto st_case_69
	case 70:
		goto st_case_70
	case 71:
		goto st_case_71
	case 0:
		goto st_case_0
	case 72:
		goto st_case_72
	case 1:
		goto st_case_1
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 73:
		goto st_case_73
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
	case 74:
		goto st_case_74
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 75:
		goto st_case_75
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 76:
		goto st_case_76
	case 17:
		goto st_case_17
	case 77:
		goto st_case_77
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
	case 47:
		goto st_case_47
	case 78:
		goto st_case_78
	case 79:
		goto st_case_79
	case 48:
		goto st_case_48
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 51:
		goto st_case_51
	case 52:
		goto st_case_52
	case 53:
		goto st_case_53
	case 54:
		goto st_case_54
	case 55:
		goto st_case_55
	case 56:
		goto st_case_56
	case 57:
		goto st_case_57
	case 58:
		goto st_case_58
	case 59:
		goto st_case_59
	case 60:
		goto st_case_60
	case 61:
		goto st_case_61
	case 62:
		goto st_case_62
	case 63:
		goto st_case_63
	case 64:
		goto st_case_64
	case 80:
		goto st_case_80
	case 65:
		goto st_case_65
	}
	goto st_out
tr0:
//line php_injection.rl:27
p = (te) - 1
{
                return true
            }
	goto st66
tr2:
//line php_injection.rl:39
p = (te) - 1

	goto st66
tr6:
//line php_injection.rl:27
te = p+1
{
                return true
            }
	goto st66
tr12:
//line php_injection.rl:35
te = p+1
{
                return true
            }
	goto st66
tr28:
//line php_injection.rl:31
te = p+1
{
                return true
            }
	goto st66
tr63:
//line php_injection.rl:39
te = p+1

	goto st66
tr74:
//line php_injection.rl:39
te = p
p--

	goto st66
tr76:
//line php_injection.rl:27
te = p
p--
{
                return true
            }
	goto st66
	st66:
//line NONE:1
ts = 0

		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
//line NONE:1
ts = p

//line php_injection.go:267
		switch data[p] {
		case 60:
			goto st67
		case 91:
			goto tr65
		case 98:
			goto tr66
		case 101:
			goto tr67
		case 103:
			goto tr68
		case 111:
			goto tr69
		case 112:
			goto tr70
		case 114:
			goto tr71
		case 115:
			goto tr72
		case 122:
			goto tr73
		}
		goto tr63
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
		if data[p] == 63 {
			goto st68
		}
		goto tr74
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
		if data[p] == 120 {
			goto st70
		}
		goto st69
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		goto st69
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
		if data[p] == 109 {
			goto tr78
		}
		goto st69
tr78:
//line NONE:1
te = p+1

	goto st71
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
//line php_injection.go:334
		if data[p] == 108 {
			goto st0
		}
		goto st69
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		goto st69
tr65:
//line NONE:1
te = p+1

	goto st72
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
//line php_injection.go:355
		switch data[p] {
		case 47:
			goto st1
		case 92:
			goto st1
		}
		goto tr74
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		if data[p] == 112 {
			goto st2
		}
		goto tr2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if data[p] == 104 {
			goto st3
		}
		goto tr2
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 112 {
			goto st4
		}
		goto tr2
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if data[p] == 93 {
			goto tr6
		}
		goto tr2
tr66:
//line NONE:1
te = p+1

	goto st73
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
//line php_injection.go:409
		if data[p] == 122 {
			goto st5
		}
		goto tr74
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		if data[p] == 105 {
			goto st6
		}
		goto tr2
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		if data[p] == 112 {
			goto st7
		}
		goto tr2
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		if data[p] == 50 {
			goto st8
		}
		goto tr2
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		if data[p] == 58 {
			goto st9
		}
		goto tr2
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		if data[p] == 47 {
			goto st10
		}
		goto tr2
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		if data[p] == 47 {
			goto tr12
		}
		goto tr2
tr67:
//line NONE:1
te = p+1

	goto st74
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
//line php_injection.go:478
		if data[p] == 120 {
			goto st11
		}
		goto tr74
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		if data[p] == 112 {
			goto st12
		}
		goto tr2
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		if data[p] == 101 {
			goto st13
		}
		goto tr2
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		if data[p] == 99 {
			goto st14
		}
		goto tr2
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		if data[p] == 116 {
			goto st8
		}
		goto tr2
tr68:
//line NONE:1
te = p+1

	goto st75
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
//line php_injection.go:529
		if data[p] == 108 {
			goto st15
		}
		goto tr74
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		if data[p] == 111 {
			goto st16
		}
		goto tr2
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		if data[p] == 98 {
			goto st8
		}
		goto tr2
tr69:
//line NONE:1
te = p+1

	goto st76
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
//line php_injection.go:562
		if data[p] == 103 {
			goto st17
		}
		goto tr74
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		if data[p] == 103 {
			goto st8
		}
		goto tr2
tr70:
//line NONE:1
te = p+1

	goto st77
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
//line php_injection.go:586
		if data[p] == 104 {
			goto st18
		}
		goto tr74
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		switch data[p] {
		case 97:
			goto st19
		case 112:
			goto st20
		}
		goto tr2
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		if data[p] == 114 {
			goto st8
		}
		goto tr2
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		if data[p] == 58 {
			goto st21
		}
		goto tr2
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		if data[p] == 47 {
			goto st22
		}
		goto tr2
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		if data[p] == 47 {
			goto st23
		}
		goto tr2
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		switch data[p] {
		case 102:
			goto st24
		case 105:
			goto st29
		case 109:
			goto st33
		case 111:
			goto st38
		case 115:
			goto st40
		case 116:
			goto st45
		}
		goto tr2
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		switch data[p] {
		case 100:
			goto tr28
		case 105:
			goto st25
		}
		goto tr2
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		if data[p] == 108 {
			goto st26
		}
		goto tr2
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		if data[p] == 116 {
			goto st27
		}
		goto tr2
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		if data[p] == 101 {
			goto st28
		}
		goto tr2
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		if data[p] == 114 {
			goto tr28
		}
		goto tr2
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		if data[p] == 110 {
			goto st30
		}
		goto tr2
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		if data[p] == 112 {
			goto st31
		}
		goto tr2
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		if data[p] == 117 {
			goto st32
		}
		goto tr2
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		if data[p] == 116 {
			goto tr28
		}
		goto tr2
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		if data[p] == 101 {
			goto st34
		}
		goto tr2
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		if data[p] == 109 {
			goto st35
		}
		goto tr2
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		if data[p] == 111 {
			goto st36
		}
		goto tr2
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		if data[p] == 114 {
			goto st37
		}
		goto tr2
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		if data[p] == 121 {
			goto tr28
		}
		goto tr2
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		if data[p] == 117 {
			goto st39
		}
		goto tr2
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		if data[p] == 116 {
			goto st30
		}
		goto tr2
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		if data[p] == 116 {
			goto st41
		}
		goto tr2
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		if data[p] == 100 {
			goto st42
		}
		goto tr2
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		switch data[p] {
		case 101:
			goto st43
		case 105:
			goto st44
		case 111:
			goto st31
		}
		goto tr2
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		if data[p] == 114 {
			goto st28
		}
		goto tr2
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
		if data[p] == 110 {
			goto tr28
		}
		goto tr2
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		if data[p] == 101 {
			goto st46
		}
		goto tr2
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		if data[p] == 109 {
			goto st47
		}
		goto tr2
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		if data[p] == 112 {
			goto tr28
		}
		goto tr2
tr71:
//line NONE:1
te = p+1

	goto st78
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
//line php_injection.go:893
		if data[p] == 97 {
			goto st19
		}
		goto tr74
tr72:
//line NONE:1
te = p+1

	goto st79
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
//line php_injection.go:908
		if data[p] == 115 {
			goto st48
		}
		goto tr74
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		if data[p] == 104 {
			goto st49
		}
		goto tr2
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		if data[p] == 50 {
			goto st50
		}
		goto tr2
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		switch data[p] {
		case 46:
			goto st51
		case 58:
			goto st9
		}
		goto tr2
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
		switch data[p] {
		case 101:
			goto st52
		case 115:
			goto st55
		case 116:
			goto st61
		}
		goto tr2
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
		if data[p] == 120 {
			goto st53
		}
		goto tr2
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
		if data[p] == 101 {
			goto st54
		}
		goto tr2
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		if data[p] == 99 {
			goto st8
		}
		goto tr2
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
		switch data[p] {
		case 99:
			goto st56
		case 102:
			goto st57
		case 104:
			goto st58
		}
		goto tr2
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		if data[p] == 112 {
			goto st8
		}
		goto tr2
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		if data[p] == 116 {
			goto st56
		}
		goto tr2
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		if data[p] == 101 {
			goto st59
		}
		goto tr2
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		if data[p] == 108 {
			goto st60
		}
		goto tr2
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
		if data[p] == 108 {
			goto st8
		}
		goto tr2
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		if data[p] == 117 {
			goto st62
		}
		goto tr2
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		if data[p] == 110 {
			goto st63
		}
		goto tr2
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		if data[p] == 110 {
			goto st64
		}
		goto tr2
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		if data[p] == 101 {
			goto st60
		}
		goto tr2
tr73:
//line NONE:1
te = p+1

	goto st80
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
//line php_injection.go:1089
		switch data[p] {
		case 105:
			goto st56
		case 108:
			goto st65
		}
		goto tr74
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		if data[p] == 105 {
			goto st16
		}
		goto tr2
	st_out:
	_test_eof66: cs = 66; goto _test_eof
	_test_eof67: cs = 67; goto _test_eof
	_test_eof68: cs = 68; goto _test_eof
	_test_eof69: cs = 69; goto _test_eof
	_test_eof70: cs = 70; goto _test_eof
	_test_eof71: cs = 71; goto _test_eof
	_test_eof0: cs = 0; goto _test_eof
	_test_eof72: cs = 72; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof73: cs = 73; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof74: cs = 74; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof75: cs = 75; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof76: cs = 76; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof77: cs = 77; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof45: cs = 45; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof78: cs = 78; goto _test_eof
	_test_eof79: cs = 79; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof
	_test_eof53: cs = 53; goto _test_eof
	_test_eof54: cs = 54; goto _test_eof
	_test_eof55: cs = 55; goto _test_eof
	_test_eof56: cs = 56; goto _test_eof
	_test_eof57: cs = 57; goto _test_eof
	_test_eof58: cs = 58; goto _test_eof
	_test_eof59: cs = 59; goto _test_eof
	_test_eof60: cs = 60; goto _test_eof
	_test_eof61: cs = 61; goto _test_eof
	_test_eof62: cs = 62; goto _test_eof
	_test_eof63: cs = 63; goto _test_eof
	_test_eof64: cs = 64; goto _test_eof
	_test_eof80: cs = 80; goto _test_eof
	_test_eof65: cs = 65; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 67:
			goto tr74
		case 68:
			goto tr76
		case 69:
			goto tr76
		case 70:
			goto tr76
		case 71:
			goto tr76
		case 0:
			goto tr0
		case 72:
			goto tr74
		case 1:
			goto tr2
		case 2:
			goto tr2
		case 3:
			goto tr2
		case 4:
			goto tr2
		case 73:
			goto tr74
		case 5:
			goto tr2
		case 6:
			goto tr2
		case 7:
			goto tr2
		case 8:
			goto tr2
		case 9:
			goto tr2
		case 10:
			goto tr2
		case 74:
			goto tr74
		case 11:
			goto tr2
		case 12:
			goto tr2
		case 13:
			goto tr2
		case 14:
			goto tr2
		case 75:
			goto tr74
		case 15:
			goto tr2
		case 16:
			goto tr2
		case 76:
			goto tr74
		case 17:
			goto tr2
		case 77:
			goto tr74
		case 18:
			goto tr2
		case 19:
			goto tr2
		case 20:
			goto tr2
		case 21:
			goto tr2
		case 22:
			goto tr2
		case 23:
			goto tr2
		case 24:
			goto tr2
		case 25:
			goto tr2
		case 26:
			goto tr2
		case 27:
			goto tr2
		case 28:
			goto tr2
		case 29:
			goto tr2
		case 30:
			goto tr2
		case 31:
			goto tr2
		case 32:
			goto tr2
		case 33:
			goto tr2
		case 34:
			goto tr2
		case 35:
			goto tr2
		case 36:
			goto tr2
		case 37:
			goto tr2
		case 38:
			goto tr2
		case 39:
			goto tr2
		case 40:
			goto tr2
		case 41:
			goto tr2
		case 42:
			goto tr2
		case 43:
			goto tr2
		case 44:
			goto tr2
		case 45:
			goto tr2
		case 46:
			goto tr2
		case 47:
			goto tr2
		case 78:
			goto tr74
		case 79:
			goto tr74
		case 48:
			goto tr2
		case 49:
			goto tr2
		case 50:
			goto tr2
		case 51:
			goto tr2
		case 52:
			goto tr2
		case 53:
			goto tr2
		case 54:
			goto tr2
		case 55:
			goto tr2
		case 56:
			goto tr2
		case 57:
			goto tr2
		case 58:
			goto tr2
		case 59:
			goto tr2
		case 60:
			goto tr2
		case 61:
			goto tr2
		case 62:
			goto tr2
		case 63:
			goto tr2
		case 64:
			goto tr2
		case 80:
			goto tr74
		case 65:
			goto tr2
		}
	}

	}

//line php_injection.rl:44

        return false
}