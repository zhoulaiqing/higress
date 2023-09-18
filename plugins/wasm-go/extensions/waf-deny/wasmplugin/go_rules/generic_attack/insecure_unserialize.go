
//line insecure_unserialize.rl:1
package generic_attack

import (
    "strings"
    "golang.org/x/exp/slices"
)

var fnsForProcess = []string{
	"access", "appendfile", "argv", "availability",
	"caveats", "chmod", "chown", "close", "copyfile", "cp", "createreadstream", "createwritestream",
	"exec", "execfile", "exists",
	"fchmod", "fchown", "fdata", "fdatasync", "fstat", "fsync", "futimes",
	"inodes",
	"lchmod", "link", "lstat", "lutimes",
	"mkdir", "mkdtemp",
	"open", "opendir",
	"read", "readdir", "readfile", "readlink", "readv", "rename", "rm",
	"spawn", "spawnfile", "stat", "symlink",
	"truncate",
	"unlink", "unwatchfile", "utimes",
	"watchfile", "write", "writefile", "writev",
}

var keywordsForProcess = []string{
	"binding", "constructor", "env", "global", "main", "mainmodule", "process", "require",
}

var fnsForConsole = []string{
	"debug", "error", "info", "trace", "warn",
}

var fnsForRequire = []string{
	"resolve", "main", "extensions", "cache",
}

func checkPattern1(value string) bool {
	snippets := strings.Split(value, ".")

	switch snippets[0] {
	case "this":
		return snippets[1] == "constructor"
	case "string":
	    return snippets[1] == "fromcharcode"
	case "module":
		return snippets[1] == "exports"
    case "constructor":
        return snippets[1] == "prototype"
	case "console":
		return slices.Contains(fnsForConsole, snippets[1])
	case "require":
		return slices.Contains(fnsForRequire, snippets[1])
	case "process":
		return slices.Contains(fnsForProcess, snippets[1]) || slices.Contains(keywordsForProcess, snippets[1])
	}

	return false
}

func checkPattern2(value string) bool {
	idx := strings.IndexByte(value, '[')
	key := value[:idx]

	// ["xxx"]
	attr := value[idx+2 : len(value)-2]

	switch key {
	case "console":
		return slices.Contains(fnsForConsole, attr)
	case "require":
		return slices.Contains(fnsForRequire, attr)
	case "process":
		return slices.Contains(fnsForProcess, attr) || slices.Contains(keywordsForProcess, attr)
	}

	return false
}



func matchInsecureUnserialize(data []byte) bool {

//line insecure_unserialize.rl:82

//line insecure_unserialize.go:87
const insecure_unserialize_start int = 105
const insecure_unserialize_first_final int = 105
const insecure_unserialize_error int = -1

const insecure_unserialize_en_main int = 105


//line insecure_unserialize.rl:83
    cs, p, pe, eof := 0, 0, len(data), len(data)

    var ts, te, act int
    _ = act

    
//line insecure_unserialize.go:102
	{
	cs = insecure_unserialize_start
	ts = 0
	te = 0
	act = 0
	}

//line insecure_unserialize.go:110
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 105:
		goto st_case_105
	case 106:
		goto st_case_106
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
	case 107:
		goto st_case_107
	case 15:
		goto st_case_15
	case 108:
		goto st_case_108
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 109:
		goto st_case_109
	case 110:
		goto st_case_110
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
	case 111:
		goto st_case_111
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
	case 112:
		goto st_case_112
	case 113:
		goto st_case_113
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
	case 114:
		goto st_case_114
	case 62:
		goto st_case_62
	case 63:
		goto st_case_63
	case 64:
		goto st_case_64
	case 65:
		goto st_case_65
	case 115:
		goto st_case_115
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
	case 72:
		goto st_case_72
	case 73:
		goto st_case_73
	case 74:
		goto st_case_74
	case 116:
		goto st_case_116
	case 75:
		goto st_case_75
	case 76:
		goto st_case_76
	case 77:
		goto st_case_77
	case 78:
		goto st_case_78
	case 117:
		goto st_case_117
	case 79:
		goto st_case_79
	case 80:
		goto st_case_80
	case 81:
		goto st_case_81
	case 82:
		goto st_case_82
	case 83:
		goto st_case_83
	case 84:
		goto st_case_84
	case 85:
		goto st_case_85
	case 86:
		goto st_case_86
	case 118:
		goto st_case_118
	case 87:
		goto st_case_87
	case 88:
		goto st_case_88
	case 89:
		goto st_case_89
	case 90:
		goto st_case_90
	case 91:
		goto st_case_91
	case 92:
		goto st_case_92
	case 93:
		goto st_case_93
	case 94:
		goto st_case_94
	case 95:
		goto st_case_95
	case 119:
		goto st_case_119
	case 96:
		goto st_case_96
	case 97:
		goto st_case_97
	case 98:
		goto st_case_98
	case 99:
		goto st_case_99
	case 100:
		goto st_case_100
	case 120:
		goto st_case_120
	case 101:
		goto st_case_101
	case 102:
		goto st_case_102
	case 103:
		goto st_case_103
	case 104:
		goto st_case_104
	}
	goto st_out
tr0:
//line insecure_unserialize.rl:126
p = (te) - 1

	goto st105
tr15:
//line insecure_unserialize.rl:110
te = p+1
{
                return true
            }
	goto st105
tr16:
//line NONE:1
	switch act {
	case 1:
	{p = (te) - 1

                return true
            }
	case 2:
	{p = (te) - 1

                if checkPattern1(string(data[ts:te])) {
                    return true
                }
            }
	default:
	{p = (te) - 1
}
	}
	
	goto st105
tr98:
//line insecure_unserialize.rl:126
te = p+1

	goto st105
tr111:
//line insecure_unserialize.rl:126
te = p
p--

	goto st105
tr113:
//line insecure_unserialize.rl:114
te = p
p--
{
                if checkPattern1(string(data[ts:te])) {
                    return true
                }
            }
	goto st105
tr114:
//line insecure_unserialize.rl:120
te = p
p--
{
                if checkPattern2(string(data[ts:te])) {
                    return true
                }
            }
	goto st105
tr115:
//line insecure_unserialize.rl:120
te = p+1
{
                if checkPattern2(string(data[ts:te])) {
                    return true
                }
            }
	goto st105
tr119:
//line insecure_unserialize.rl:110
te = p
p--
{
                return true
            }
	goto st105
	st105:
//line NONE:1
ts = 0

		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
//line NONE:1
ts = p

//line insecure_unserialize.go:452
		switch data[p] {
		case 40:
			goto tr99
		case 95:
			goto tr101
		case 98:
			goto tr102
		case 99:
			goto tr103
		case 101:
			goto tr104
		case 102:
			goto tr105
		case 103:
			goto tr106
		case 109:
			goto tr107
		case 110:
			goto tr108
		case 112:
			goto tr109
		case 114:
			goto tr110
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr100
			}
		case data[p] >= 65:
			goto tr100
		}
		goto tr98
tr99:
//line NONE:1
te = p+1

	goto st106
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
//line insecure_unserialize.go:496
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr111
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr111
			}
		default:
			goto tr111
		}
		goto st0
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		if data[p] == 99 {
			goto st1
		}
		goto tr0
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		if data[p] == 104 {
			goto st2
		}
		goto tr0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if data[p] == 105 {
			goto st3
		}
		goto tr0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 108 {
			goto st4
		}
		goto tr0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if data[p] == 100 {
			goto st5
		}
		goto tr0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		if data[p] == 95 {
			goto st6
		}
		goto tr0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		if data[p] == 112 {
			goto st7
		}
		goto tr0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		if data[p] == 114 {
			goto st8
		}
		goto tr0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		if data[p] == 111 {
			goto st9
		}
		goto tr0
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		if data[p] == 99 {
			goto st10
		}
		goto tr0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		if data[p] == 101 {
			goto st11
		}
		goto tr0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		if data[p] == 115 {
			goto st12
		}
		goto tr0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		if data[p] == 115 {
			goto st13
		}
		goto tr0
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr0
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr0
			}
		default:
			goto tr0
		}
		goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		if data[p] == 41 {
			goto tr15
		}
		goto tr0
tr100:
//line NONE:1
te = p+1

//line insecure_unserialize.rl:126
act = 4;
	goto st107
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
//line insecure_unserialize.go:666
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr111
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr17
			}
		case data[p] >= 65:
			goto tr17
		}
		goto tr16
tr17:
//line NONE:1
te = p+1

//line insecure_unserialize.rl:114
act = 2;
	goto st108
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
//line insecure_unserialize.go:708
		if data[p] == 46 {
			goto st15
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr17
			}
		case data[p] >= 65:
			goto tr17
		}
		goto tr113
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		switch data[p] {
		case 34:
			goto st18
		case 39:
			goto st18
		case 96:
			goto st18
		}
		goto tr0
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st19
			}
		case data[p] >= 65:
			goto st19
		}
		goto tr16
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		switch data[p] {
		case 34:
			goto st109
		case 39:
			goto st109
		case 96:
			goto st109
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st19
			}
		case data[p] >= 65:
			goto st19
		}
		goto tr16
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
		if data[p] == 93 {
			goto tr115
		}
		goto tr114
tr101:
//line NONE:1
te = p+1

	goto st110
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
//line insecure_unserialize.go:810
		switch data[p] {
		case 36:
			goto st20
		case 95:
			goto st31
		}
		goto tr111
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		if data[p] == 36 {
			goto st21
		}
		goto tr0
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		if data[p] == 110 {
			goto st22
		}
		goto tr0
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		if data[p] == 100 {
			goto st23
		}
		goto tr0
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		if data[p] == 95 {
			goto st24
		}
		goto tr0
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		if data[p] == 102 {
			goto st25
		}
		goto tr0
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		if data[p] == 117 {
			goto st26
		}
		goto tr0
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		if data[p] == 110 {
			goto st27
		}
		goto tr0
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		if data[p] == 99 {
			goto st28
		}
		goto tr0
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		if data[p] == 36 {
			goto st29
		}
		goto tr0
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		if data[p] == 36 {
			goto st30
		}
		goto tr0
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		if data[p] == 95 {
			goto tr15
		}
		goto tr0
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		switch data[p] {
		case 106:
			goto st32
		case 112:
			goto st42
		}
		goto tr0
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		if data[p] == 115 {
			goto st33
		}
		goto tr0
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		if data[p] == 95 {
			goto st34
		}
		goto tr0
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		if data[p] == 102 {
			goto st35
		}
		goto tr0
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		if data[p] == 117 {
			goto st36
		}
		goto tr0
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		if data[p] == 110 {
			goto st37
		}
		goto tr0
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		if data[p] == 99 {
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
		if data[p] == 105 {
			goto st40
		}
		goto tr0
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		if data[p] == 111 {
			goto st41
		}
		goto tr0
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		if data[p] == 110 {
			goto tr15
		}
		goto tr0
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		if data[p] == 114 {
			goto st43
		}
		goto tr0
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		if data[p] == 111 {
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
		if data[p] == 111 {
			goto st46
		}
		goto tr0
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		if data[p] == 95 {
			goto st30
		}
		goto tr0
tr102:
//line NONE:1
te = p+1

//line insecure_unserialize.rl:126
act = 4;
	goto st111
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
//line insecure_unserialize.go:1076
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 105:
			goto st47
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr111
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 110:
			goto st48
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 100:
			goto st49
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 105:
			goto st50
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 110:
			goto st51
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 103:
			goto st52
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto tr54
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
tr54:
//line NONE:1
te = p+1

//line insecure_unserialize.rl:110
act = 1;
	goto st112
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
//line insecure_unserialize.go:1236
		switch data[p] {
		case 34:
			goto st18
		case 39:
			goto st18
		case 96:
			goto st18
		}
		goto tr119
tr103:
//line NONE:1
te = p+1

//line insecure_unserialize.rl:126
act = 4;
	goto st113
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
//line insecure_unserialize.go:1258
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 111:
			goto st53
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr111
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 110:
			goto st54
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 115:
			goto st55
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 116:
			goto st56
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 114:
			goto st57
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 117:
			goto st58
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 99:
			goto st59
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 116:
			goto st60
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 111:
			goto st61
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 114:
			goto st52
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
tr104:
//line NONE:1
te = p+1

//line insecure_unserialize.rl:126
act = 4;
	goto st114
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
//line insecure_unserialize.go:1486
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 110:
			goto st62
		case 118:
			goto st63
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr111
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 118:
			goto st52
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 97:
			goto st64
		}
		switch {
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 108:
			goto st65
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		switch data[p] {
		case 40:
			goto tr15
		case 46:
			goto st15
		case 91:
			goto st17
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
tr105:
//line NONE:1
te = p+1

//line insecure_unserialize.rl:126
act = 4;
	goto st115
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
//line insecure_unserialize.go:1606
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 117:
			goto st66
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr111
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 110:
			goto st67
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 99:
			goto st68
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 116:
			goto st69
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 105:
			goto st70
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 111:
			goto st71
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 110:
			goto st72
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
		switch data[p] {
		case 40:
			goto st73
		case 46:
			goto st15
		case 91:
			goto st17
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
		if data[p] == 41 {
			goto st74
		}
		goto tr0
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
		if data[p] == 123 {
			goto tr15
		}
		goto tr0
tr106:
//line NONE:1
te = p+1

//line insecure_unserialize.rl:126
act = 4;
	goto st116
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
//line insecure_unserialize.go:1808
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 108:
			goto st75
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr111
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 111:
			goto st76
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 98:
			goto st77
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 97:
			goto st78
		}
		switch {
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 108:
			goto st52
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
tr107:
//line NONE:1
te = p+1

//line insecure_unserialize.rl:126
act = 4;
	goto st117
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
//line insecure_unserialize.go:1926
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 97:
			goto st79
		}
		switch {
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr111
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 105:
			goto st80
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 110:
			goto st81
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto tr54
		case 109:
			goto st82
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 111:
			goto st83
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 100:
			goto st84
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 117:
			goto st85
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 108:
			goto st86
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 101:
			goto st52
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
tr108:
//line NONE:1
te = p+1

//line insecure_unserialize.rl:126
act = 4;
	goto st118
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
//line insecure_unserialize.go:2132
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 101:
			goto st87
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr111
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 119:
			goto st88
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 102:
			goto st89
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 117:
			goto st90
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 110:
			goto st91
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 99:
			goto st92
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 116:
			goto st93
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 105:
			goto st94
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 111:
			goto st95
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 110:
			goto st65
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
tr109:
//line NONE:1
te = p+1

//line insecure_unserialize.rl:126
act = 4;
	goto st119
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
//line insecure_unserialize.go:2360
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 114:
			goto st96
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr111
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 111:
			goto st97
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 99:
			goto st98
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 101:
			goto st99
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 115:
			goto st100
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 115:
			goto st52
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
tr110:
//line NONE:1
te = p+1

//line insecure_unserialize.rl:126
act = 4;
	goto st120
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
//line insecure_unserialize.go:2500
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 101:
			goto st101
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr111
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 113:
			goto st102
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 117:
			goto st103
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 105:
			goto st104
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
		switch data[p] {
		case 46:
			goto st15
		case 91:
			goto st17
		case 114:
			goto st86
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		case data[p] >= 65:
			goto st16
		}
		goto tr0
	st_out:
	_test_eof105: cs = 105; goto _test_eof
	_test_eof106: cs = 106; goto _test_eof
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
	_test_eof107: cs = 107; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof108: cs = 108; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof109: cs = 109; goto _test_eof
	_test_eof110: cs = 110; goto _test_eof
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
	_test_eof111: cs = 111; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof
	_test_eof112: cs = 112; goto _test_eof
	_test_eof113: cs = 113; goto _test_eof
	_test_eof53: cs = 53; goto _test_eof
	_test_eof54: cs = 54; goto _test_eof
	_test_eof55: cs = 55; goto _test_eof
	_test_eof56: cs = 56; goto _test_eof
	_test_eof57: cs = 57; goto _test_eof
	_test_eof58: cs = 58; goto _test_eof
	_test_eof59: cs = 59; goto _test_eof
	_test_eof60: cs = 60; goto _test_eof
	_test_eof61: cs = 61; goto _test_eof
	_test_eof114: cs = 114; goto _test_eof
	_test_eof62: cs = 62; goto _test_eof
	_test_eof63: cs = 63; goto _test_eof
	_test_eof64: cs = 64; goto _test_eof
	_test_eof65: cs = 65; goto _test_eof
	_test_eof115: cs = 115; goto _test_eof
	_test_eof66: cs = 66; goto _test_eof
	_test_eof67: cs = 67; goto _test_eof
	_test_eof68: cs = 68; goto _test_eof
	_test_eof69: cs = 69; goto _test_eof
	_test_eof70: cs = 70; goto _test_eof
	_test_eof71: cs = 71; goto _test_eof
	_test_eof72: cs = 72; goto _test_eof
	_test_eof73: cs = 73; goto _test_eof
	_test_eof74: cs = 74; goto _test_eof
	_test_eof116: cs = 116; goto _test_eof
	_test_eof75: cs = 75; goto _test_eof
	_test_eof76: cs = 76; goto _test_eof
	_test_eof77: cs = 77; goto _test_eof
	_test_eof78: cs = 78; goto _test_eof
	_test_eof117: cs = 117; goto _test_eof
	_test_eof79: cs = 79; goto _test_eof
	_test_eof80: cs = 80; goto _test_eof
	_test_eof81: cs = 81; goto _test_eof
	_test_eof82: cs = 82; goto _test_eof
	_test_eof83: cs = 83; goto _test_eof
	_test_eof84: cs = 84; goto _test_eof
	_test_eof85: cs = 85; goto _test_eof
	_test_eof86: cs = 86; goto _test_eof
	_test_eof118: cs = 118; goto _test_eof
	_test_eof87: cs = 87; goto _test_eof
	_test_eof88: cs = 88; goto _test_eof
	_test_eof89: cs = 89; goto _test_eof
	_test_eof90: cs = 90; goto _test_eof
	_test_eof91: cs = 91; goto _test_eof
	_test_eof92: cs = 92; goto _test_eof
	_test_eof93: cs = 93; goto _test_eof
	_test_eof94: cs = 94; goto _test_eof
	_test_eof95: cs = 95; goto _test_eof
	_test_eof119: cs = 119; goto _test_eof
	_test_eof96: cs = 96; goto _test_eof
	_test_eof97: cs = 97; goto _test_eof
	_test_eof98: cs = 98; goto _test_eof
	_test_eof99: cs = 99; goto _test_eof
	_test_eof100: cs = 100; goto _test_eof
	_test_eof120: cs = 120; goto _test_eof
	_test_eof101: cs = 101; goto _test_eof
	_test_eof102: cs = 102; goto _test_eof
	_test_eof103: cs = 103; goto _test_eof
	_test_eof104: cs = 104; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 106:
			goto tr111
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
		case 107:
			goto tr111
		case 15:
			goto tr16
		case 108:
			goto tr113
		case 16:
			goto tr0
		case 17:
			goto tr0
		case 18:
			goto tr16
		case 19:
			goto tr16
		case 109:
			goto tr114
		case 110:
			goto tr111
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
		case 34:
			goto tr0
		case 35:
			goto tr0
		case 36:
			goto tr0
		case 37:
			goto tr0
		case 38:
			goto tr0
		case 39:
			goto tr0
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
		case 111:
			goto tr111
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
		case 112:
			goto tr119
		case 113:
			goto tr111
		case 53:
			goto tr0
		case 54:
			goto tr0
		case 55:
			goto tr0
		case 56:
			goto tr0
		case 57:
			goto tr0
		case 58:
			goto tr0
		case 59:
			goto tr0
		case 60:
			goto tr0
		case 61:
			goto tr0
		case 114:
			goto tr111
		case 62:
			goto tr0
		case 63:
			goto tr0
		case 64:
			goto tr0
		case 65:
			goto tr0
		case 115:
			goto tr111
		case 66:
			goto tr0
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
		case 72:
			goto tr0
		case 73:
			goto tr0
		case 74:
			goto tr0
		case 116:
			goto tr111
		case 75:
			goto tr0
		case 76:
			goto tr0
		case 77:
			goto tr0
		case 78:
			goto tr0
		case 117:
			goto tr111
		case 79:
			goto tr0
		case 80:
			goto tr0
		case 81:
			goto tr0
		case 82:
			goto tr0
		case 83:
			goto tr0
		case 84:
			goto tr0
		case 85:
			goto tr0
		case 86:
			goto tr0
		case 118:
			goto tr111
		case 87:
			goto tr0
		case 88:
			goto tr0
		case 89:
			goto tr0
		case 90:
			goto tr0
		case 91:
			goto tr0
		case 92:
			goto tr0
		case 93:
			goto tr0
		case 94:
			goto tr0
		case 95:
			goto tr0
		case 119:
			goto tr111
		case 96:
			goto tr0
		case 97:
			goto tr0
		case 98:
			goto tr0
		case 99:
			goto tr0
		case 100:
			goto tr0
		case 120:
			goto tr111
		case 101:
			goto tr0
		case 102:
			goto tr0
		case 103:
			goto tr0
		case 104:
			goto tr0
		}
	}

	}

//line insecure_unserialize.rl:133

        return false
}