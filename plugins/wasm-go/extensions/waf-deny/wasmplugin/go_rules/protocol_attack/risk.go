
//line risk.rl:1
package protocol_attack

func matchRisk(data []byte) bool {

//line risk.rl:5

//line risk.go:10
const protocol_attack_risk_start int = 86
const protocol_attack_risk_first_final int = 86
const protocol_attack_risk_error int = -1

const protocol_attack_risk_en_main int = 86


//line risk.rl:6
    cs, p, pe, eof := 0, 0, len(data), len(data)

    var ts, te, act int
    _, _, _ = ts, te, act

    
//line risk.go:25
	{
	cs = protocol_attack_risk_start
	ts = 0
	te = 0
	act = 0
	}

//line risk.go:33
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 86:
		goto st_case_86
	case 87:
		goto st_case_87
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
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
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
	case 88:
		goto st_case_88
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 89:
		goto st_case_89
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 90:
		goto st_case_90
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
	case 91:
		goto st_case_91
	case 55:
		goto st_case_55
	case 56:
		goto st_case_56
	case 57:
		goto st_case_57
	case 58:
		goto st_case_58
	case 92:
		goto st_case_92
	case 93:
		goto st_case_93
	case 59:
		goto st_case_59
	case 60:
		goto st_case_60
	case 94:
		goto st_case_94
	case 61:
		goto st_case_61
	case 62:
		goto st_case_62
	case 95:
		goto st_case_95
	case 63:
		goto st_case_63
	case 64:
		goto st_case_64
	case 65:
		goto st_case_65
	case 66:
		goto st_case_66
	case 96:
		goto st_case_96
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
	case 97:
		goto st_case_97
	case 72:
		goto st_case_72
	case 73:
		goto st_case_73
	case 74:
		goto st_case_74
	case 75:
		goto st_case_75
	case 76:
		goto st_case_76
	case 77:
		goto st_case_77
	case 78:
		goto st_case_78
	case 79:
		goto st_case_79
	case 80:
		goto st_case_80
	case 98:
		goto st_case_98
	case 81:
		goto st_case_81
	case 82:
		goto st_case_82
	case 83:
		goto st_case_83
	case 99:
		goto st_case_99
	case 84:
		goto st_case_84
	case 85:
		goto st_case_85
	}
	goto st_out
tr0:
//line risk.rl:47
p = (te) - 1

	goto st86
tr20:
//line risk.rl:43
te = p+1
{
                return true
            }
	goto st86
tr55:
//line risk.rl:35
te = p+1
{
                return true
            }
	goto st86
tr78:
//line risk.rl:47
te = p+1

	goto st86
tr91:
//line risk.rl:47
te = p
p--

	goto st86
tr94:
//line risk.rl:39
te = p
p--
{
                return true
            }
	goto st86
tr95:
//line risk.rl:39
te = p+1
{
                return true
            }
	goto st86
	st86:
//line NONE:1
ts = 0

		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
//line NONE:1
ts = p

//line risk.go:297
		switch data[p] {
		case 10:
			goto tr79
		case 13:
			goto tr79
		case 60:
			goto tr80
		case 99:
			goto tr81
		case 100:
			goto tr82
		case 103:
			goto tr83
		case 104:
			goto tr84
		case 108:
			goto tr85
		case 109:
			goto tr86
		case 111:
			goto tr87
		case 112:
			goto tr88
		case 116:
			goto tr89
		case 117:
			goto tr90
		}
		goto tr78
tr79:
//line NONE:1
te = p+1

	goto st87
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
//line risk.go:337
		switch data[p] {
		case 95:
			goto tr91
		case 99:
			goto st1
		case 108:
			goto st19
		case 115:
			goto st26
		}
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr91
			}
		case data[p] >= 48:
			goto tr91
		}
		goto st0
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		switch data[p] {
		case 95:
			goto tr0
		case 99:
			goto st1
		case 108:
			goto st19
		case 115:
			goto st26
		}
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr0
			}
		case data[p] >= 48:
			goto tr0
		}
		goto st0
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		if data[p] == 111 {
			goto st2
		}
		goto tr0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if data[p] == 110 {
			goto st3
		}
		goto tr0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 116 {
			goto st4
		}
		goto tr0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if data[p] == 101 {
			goto st5
		}
		goto tr0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		if data[p] == 110 {
			goto st6
		}
		goto tr0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		if data[p] == 116 {
			goto st7
		}
		goto tr0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		if data[p] == 45 {
			goto st8
		}
		goto tr0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		switch data[p] {
		case 108:
			goto st9
		case 116:
			goto st16
		}
		goto tr0
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		if data[p] == 101 {
			goto st10
		}
		goto tr0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		if data[p] == 110 {
			goto st11
		}
		goto tr0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		if data[p] == 103 {
			goto st12
		}
		goto tr0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		if data[p] == 116 {
			goto st13
		}
		goto tr0
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		if data[p] == 104 {
			goto st14
		}
		goto tr0
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		if data[p] == 58 {
			goto st15
		}
		goto tr0
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		switch data[p] {
		case 32:
			goto st15
		case 95:
			goto tr20
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st15
			}
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr20
			}
		default:
			goto tr20
		}
		goto tr0
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		if data[p] == 121 {
			goto st17
		}
		goto tr0
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		if data[p] == 112 {
			goto st18
		}
		goto tr0
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		if data[p] == 101 {
			goto st14
		}
		goto tr0
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		if data[p] == 111 {
			goto st20
		}
		goto tr0
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		if data[p] == 99 {
			goto st21
		}
		goto tr0
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		if data[p] == 97 {
			goto st22
		}
		goto tr0
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		if data[p] == 116 {
			goto st23
		}
		goto tr0
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		if data[p] == 105 {
			goto st24
		}
		goto tr0
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		if data[p] == 111 {
			goto st25
		}
		goto tr0
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		if data[p] == 110 {
			goto st14
		}
		goto tr0
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		if data[p] == 101 {
			goto st27
		}
		goto tr0
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		if data[p] == 116 {
			goto st28
		}
		goto tr0
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		if data[p] == 45 {
			goto st29
		}
		goto tr0
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		if data[p] == 99 {
			goto st30
		}
		goto tr0
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		if data[p] == 111 {
			goto st31
		}
		goto tr0
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		if data[p] == 111 {
			goto st32
		}
		goto tr0
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		if data[p] == 107 {
			goto st33
		}
		goto tr0
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		if data[p] == 105 {
			goto st18
		}
		goto tr0
tr80:
//line NONE:1
te = p+1

	goto st88
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
//line risk.go:706
		switch data[p] {
		case 104:
			goto st34
		case 109:
			goto st37
		}
		goto tr91
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		if data[p] == 116 {
			goto st35
		}
		goto tr0
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		if data[p] == 109 {
			goto st36
		}
		goto tr0
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		if data[p] == 108 {
			goto st89
		}
		goto tr0
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
		if data[p] == 95 {
			goto tr94
		}
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr94
			}
		case data[p] >= 48:
			goto tr94
		}
		goto tr95
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		if data[p] == 101 {
			goto st38
		}
		goto tr0
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		if data[p] == 116 {
			goto st39
		}
		goto tr0
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		if data[p] == 97 {
			goto st89
		}
		goto tr0
tr81:
//line NONE:1
te = p+1

	goto st90
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
//line risk.go:795
		if data[p] == 111 {
			goto st40
		}
		goto tr91
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		switch data[p] {
		case 110:
			goto st41
		case 112:
			goto st54
		}
		goto tr0
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		if data[p] == 110 {
			goto st42
		}
		goto tr0
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		if data[p] == 101 {
			goto st43
		}
		goto tr0
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		if data[p] == 99 {
			goto st44
		}
		goto tr0
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
		if data[p] == 116 {
			goto st45
		}
		goto tr0
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		if data[p] == 32 {
			goto st46
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st46
		}
		goto tr0
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		if data[p] == 32 {
			goto st46
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st46
		}
		goto st47
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		if data[p] == 32 {
			goto st48
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st48
		}
		goto st47
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		switch data[p] {
		case 32:
			goto st48
		case 104:
			goto st49
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st48
		}
		goto tr0
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		if data[p] == 116 {
			goto st50
		}
		goto tr0
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		if data[p] == 116 {
			goto st51
		}
		goto tr0
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
		if data[p] == 112 {
			goto st52
		}
		goto tr0
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
		if data[p] == 47 {
			goto st53
		}
		goto tr0
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr55
		}
		goto tr0
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		if data[p] == 121 {
			goto st45
		}
		goto tr0
tr82:
//line NONE:1
te = p+1

	goto st91
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
//line risk.go:963
		if data[p] == 101 {
			goto st55
		}
		goto tr91
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
		if data[p] == 108 {
			goto st56
		}
		goto tr0
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		if data[p] == 101 {
			goto st57
		}
		goto tr0
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		if data[p] == 116 {
			goto st58
		}
		goto tr0
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		if data[p] == 101 {
			goto st45
		}
		goto tr0
tr83:
//line NONE:1
te = p+1

	goto st92
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
//line risk.go:1014
		if data[p] == 101 {
			goto st44
		}
		goto tr91
tr84:
//line NONE:1
te = p+1

	goto st93
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
//line risk.go:1029
		if data[p] == 101 {
			goto st59
		}
		goto tr91
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		if data[p] == 97 {
			goto st60
		}
		goto tr0
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
		if data[p] == 100 {
			goto st45
		}
		goto tr0
tr85:
//line NONE:1
te = p+1

	goto st94
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
//line risk.go:1062
		if data[p] == 111 {
			goto st61
		}
		goto tr91
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		if data[p] == 99 {
			goto st62
		}
		goto tr0
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		if data[p] == 107 {
			goto st45
		}
		goto tr0
tr86:
//line NONE:1
te = p+1

	goto st95
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
//line risk.go:1095
		switch data[p] {
		case 107:
			goto st63
		case 111:
			goto st66
		}
		goto tr91
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		if data[p] == 99 {
			goto st64
		}
		goto tr0
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		if data[p] == 111 {
			goto st65
		}
		goto tr0
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		if data[p] == 108 {
			goto st45
		}
		goto tr0
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
		if data[p] == 118 {
			goto st58
		}
		goto tr0
tr87:
//line NONE:1
te = p+1

	goto st96
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
//line risk.go:1149
		if data[p] == 112 {
			goto st67
		}
		goto tr91
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
		if data[p] == 116 {
			goto st68
		}
		goto tr0
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
		if data[p] == 105 {
			goto st69
		}
		goto tr0
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		if data[p] == 111 {
			goto st70
		}
		goto tr0
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
		if data[p] == 110 {
			goto st71
		}
		goto tr0
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
		if data[p] == 115 {
			goto st45
		}
		goto tr0
tr88:
//line NONE:1
te = p+1

	goto st97
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
//line risk.go:1209
		switch data[p] {
		case 97:
			goto st72
		case 111:
			goto st75
		case 114:
			goto st76
		case 117:
			goto st44
		}
		goto tr91
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
		if data[p] == 116 {
			goto st73
		}
		goto tr0
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
		if data[p] == 99 {
			goto st74
		}
		goto tr0
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
		if data[p] == 104 {
			goto st45
		}
		goto tr0
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
		if data[p] == 115 {
			goto st44
		}
		goto tr0
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
		if data[p] == 111 {
			goto st77
		}
		goto tr0
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
		if data[p] == 112 {
			goto st78
		}
		goto tr0
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
		switch data[p] {
		case 97:
			goto st72
		case 102:
			goto st79
		}
		goto tr0
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
		if data[p] == 105 {
			goto st80
		}
		goto tr0
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
		if data[p] == 110 {
			goto st60
		}
		goto tr0
tr89:
//line NONE:1
te = p+1

	goto st98
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
//line risk.go:1315
		if data[p] == 114 {
			goto st81
		}
		goto tr91
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
		if data[p] == 97 {
			goto st82
		}
		goto tr0
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
		if data[p] == 99 {
			goto st83
		}
		goto tr0
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
		switch data[p] {
		case 101:
			goto st45
		case 107:
			goto st45
		}
		goto tr0
tr90:
//line NONE:1
te = p+1

	goto st99
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
//line risk.go:1360
		if data[p] == 110 {
			goto st84
		}
		goto tr91
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
		if data[p] == 108 {
			goto st85
		}
		goto tr0
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
		if data[p] == 111 {
			goto st61
		}
		goto tr0
	st_out:
	_test_eof86: cs = 86; goto _test_eof
	_test_eof87: cs = 87; goto _test_eof
	_test_eof0: cs = 0; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
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
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
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
	_test_eof88: cs = 88; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof89: cs = 89; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof90: cs = 90; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof45: cs = 45; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof
	_test_eof53: cs = 53; goto _test_eof
	_test_eof54: cs = 54; goto _test_eof
	_test_eof91: cs = 91; goto _test_eof
	_test_eof55: cs = 55; goto _test_eof
	_test_eof56: cs = 56; goto _test_eof
	_test_eof57: cs = 57; goto _test_eof
	_test_eof58: cs = 58; goto _test_eof
	_test_eof92: cs = 92; goto _test_eof
	_test_eof93: cs = 93; goto _test_eof
	_test_eof59: cs = 59; goto _test_eof
	_test_eof60: cs = 60; goto _test_eof
	_test_eof94: cs = 94; goto _test_eof
	_test_eof61: cs = 61; goto _test_eof
	_test_eof62: cs = 62; goto _test_eof
	_test_eof95: cs = 95; goto _test_eof
	_test_eof63: cs = 63; goto _test_eof
	_test_eof64: cs = 64; goto _test_eof
	_test_eof65: cs = 65; goto _test_eof
	_test_eof66: cs = 66; goto _test_eof
	_test_eof96: cs = 96; goto _test_eof
	_test_eof67: cs = 67; goto _test_eof
	_test_eof68: cs = 68; goto _test_eof
	_test_eof69: cs = 69; goto _test_eof
	_test_eof70: cs = 70; goto _test_eof
	_test_eof71: cs = 71; goto _test_eof
	_test_eof97: cs = 97; goto _test_eof
	_test_eof72: cs = 72; goto _test_eof
	_test_eof73: cs = 73; goto _test_eof
	_test_eof74: cs = 74; goto _test_eof
	_test_eof75: cs = 75; goto _test_eof
	_test_eof76: cs = 76; goto _test_eof
	_test_eof77: cs = 77; goto _test_eof
	_test_eof78: cs = 78; goto _test_eof
	_test_eof79: cs = 79; goto _test_eof
	_test_eof80: cs = 80; goto _test_eof
	_test_eof98: cs = 98; goto _test_eof
	_test_eof81: cs = 81; goto _test_eof
	_test_eof82: cs = 82; goto _test_eof
	_test_eof83: cs = 83; goto _test_eof
	_test_eof99: cs = 99; goto _test_eof
	_test_eof84: cs = 84; goto _test_eof
	_test_eof85: cs = 85; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 87:
			goto tr91
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
		case 8:
			goto tr0
		case 9:
			goto tr0
		case 10:
			goto tr0
		case 11:
			goto tr0
		case 12:
			goto tr0
		case 13:
			goto tr0
		case 14:
			goto tr0
		case 15:
			goto tr0
		case 16:
			goto tr0
		case 17:
			goto tr0
		case 18:
			goto tr0
		case 19:
			goto tr0
		case 20:
			goto tr0
		case 21:
			goto tr0
		case 22:
			goto tr0
		case 23:
			goto tr0
		case 24:
			goto tr0
		case 25:
			goto tr0
		case 26:
			goto tr0
		case 27:
			goto tr0
		case 28:
			goto tr0
		case 29:
			goto tr0
		case 30:
			goto tr0
		case 31:
			goto tr0
		case 32:
			goto tr0
		case 33:
			goto tr0
		case 88:
			goto tr91
		case 34:
			goto tr0
		case 35:
			goto tr0
		case 36:
			goto tr0
		case 89:
			goto tr94
		case 37:
			goto tr0
		case 38:
			goto tr0
		case 39:
			goto tr0
		case 90:
			goto tr91
		case 40:
			goto tr0
		case 41:
			goto tr0
		case 42:
			goto tr0
		case 43:
			goto tr0
		case 44:
			goto tr0
		case 45:
			goto tr0
		case 46:
			goto tr0
		case 47:
			goto tr0
		case 48:
			goto tr0
		case 49:
			goto tr0
		case 50:
			goto tr0
		case 51:
			goto tr0
		case 52:
			goto tr0
		case 53:
			goto tr0
		case 54:
			goto tr0
		case 91:
			goto tr91
		case 55:
			goto tr0
		case 56:
			goto tr0
		case 57:
			goto tr0
		case 58:
			goto tr0
		case 92:
			goto tr91
		case 93:
			goto tr91
		case 59:
			goto tr0
		case 60:
			goto tr0
		case 94:
			goto tr91
		case 61:
			goto tr0
		case 62:
			goto tr0
		case 95:
			goto tr91
		case 63:
			goto tr0
		case 64:
			goto tr0
		case 65:
			goto tr0
		case 66:
			goto tr0
		case 96:
			goto tr91
		case 67:
			goto tr0
		case 68:
			goto tr0
		case 69:
			goto tr0
		case 70:
			goto tr0
		case 71:
			goto tr0
		case 97:
			goto tr91
		case 72:
			goto tr0
		case 73:
			goto tr0
		case 74:
			goto tr0
		case 75:
			goto tr0
		case 76:
			goto tr0
		case 77:
			goto tr0
		case 78:
			goto tr0
		case 79:
			goto tr0
		case 80:
			goto tr0
		case 98:
			goto tr91
		case 81:
			goto tr0
		case 82:
			goto tr0
		case 83:
			goto tr0
		case 99:
			goto tr91
		case 84:
			goto tr0
		case 85:
			goto tr0
		}
	}

	}

//line risk.rl:52

        return false
}
