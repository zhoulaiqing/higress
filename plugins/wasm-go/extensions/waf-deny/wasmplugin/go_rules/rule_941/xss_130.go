
//line xss_130.rl:1
package rule_941

import (
    "fmt"
)

func matchXSS130(data []byte) bool {

//line xss_130.rl:9

//line xss_130.go:14
const xss_start int = 0
const xss_first_final int = 183
const xss_error int = -1

const xss_en_main int = 0


//line xss_130.rl:10
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    lastIsWordEle := false

    
//line xss_130.go:29
	{
	cs = xss_start
	}

//line xss_130.go:34
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
	case 183:
		goto st_case_183
	case 184:
		goto st_case_184
	case 185:
		goto st_case_185
	case 186:
		goto st_case_186
	case 187:
		goto st_case_187
	case 188:
		goto st_case_188
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
	case 189:
		goto st_case_189
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
	case 190:
		goto st_case_190
	case 191:
		goto st_case_191
	case 192:
		goto st_case_192
	case 193:
		goto st_case_193
	case 194:
		goto st_case_194
	case 195:
		goto st_case_195
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
	case 65:
		goto st_case_65
	case 66:
		goto st_case_66
	case 196:
		goto st_case_196
	case 197:
		goto st_case_197
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
	case 101:
		goto st_case_101
	case 102:
		goto st_case_102
	case 103:
		goto st_case_103
	case 104:
		goto st_case_104
	case 105:
		goto st_case_105
	case 106:
		goto st_case_106
	case 107:
		goto st_case_107
	case 108:
		goto st_case_108
	case 109:
		goto st_case_109
	case 110:
		goto st_case_110
	case 111:
		goto st_case_111
	case 112:
		goto st_case_112
	case 113:
		goto st_case_113
	case 114:
		goto st_case_114
	case 115:
		goto st_case_115
	case 116:
		goto st_case_116
	case 117:
		goto st_case_117
	case 118:
		goto st_case_118
	case 119:
		goto st_case_119
	case 120:
		goto st_case_120
	case 121:
		goto st_case_121
	case 122:
		goto st_case_122
	case 123:
		goto st_case_123
	case 124:
		goto st_case_124
	case 125:
		goto st_case_125
	case 126:
		goto st_case_126
	case 127:
		goto st_case_127
	case 128:
		goto st_case_128
	case 129:
		goto st_case_129
	case 130:
		goto st_case_130
	case 198:
		goto st_case_198
	case 199:
		goto st_case_199
	case 200:
		goto st_case_200
	case 131:
		goto st_case_131
	case 132:
		goto st_case_132
	case 133:
		goto st_case_133
	case 134:
		goto st_case_134
	case 135:
		goto st_case_135
	case 136:
		goto st_case_136
	case 137:
		goto st_case_137
	case 138:
		goto st_case_138
	case 139:
		goto st_case_139
	case 140:
		goto st_case_140
	case 141:
		goto st_case_141
	case 142:
		goto st_case_142
	case 143:
		goto st_case_143
	case 144:
		goto st_case_144
	case 145:
		goto st_case_145
	case 146:
		goto st_case_146
	case 147:
		goto st_case_147
	case 148:
		goto st_case_148
	case 149:
		goto st_case_149
	case 150:
		goto st_case_150
	case 151:
		goto st_case_151
	case 152:
		goto st_case_152
	case 153:
		goto st_case_153
	case 154:
		goto st_case_154
	case 155:
		goto st_case_155
	case 156:
		goto st_case_156
	case 157:
		goto st_case_157
	case 158:
		goto st_case_158
	case 159:
		goto st_case_159
	case 160:
		goto st_case_160
	case 161:
		goto st_case_161
	case 162:
		goto st_case_162
	case 163:
		goto st_case_163
	case 164:
		goto st_case_164
	case 165:
		goto st_case_165
	case 166:
		goto st_case_166
	case 167:
		goto st_case_167
	case 168:
		goto st_case_168
	case 169:
		goto st_case_169
	case 170:
		goto st_case_170
	case 171:
		goto st_case_171
	case 172:
		goto st_case_172
	case 173:
		goto st_case_173
	case 174:
		goto st_case_174
	case 175:
		goto st_case_175
	case 176:
		goto st_case_176
	case 177:
		goto st_case_177
	case 178:
		goto st_case_178
	case 179:
		goto st_case_179
	case 180:
		goto st_case_180
	case 181:
		goto st_case_181
	case 182:
		goto st_case_182
	}
	goto st_out
	st_case_0:
		goto st1
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		if data[p] == 95 {
			goto st1
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 33:
			goto tr2
		case 59:
			goto tr3
		case 64:
			goto tr4
		case 95:
			goto st1
		case 100:
			goto tr5
		case 102:
			goto tr6
		case 104:
			goto tr7
		case 112:
			goto tr8
		case 120:
			goto tr9
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
tr2:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line xss_130.go:520
		switch data[p] {
		case 33:
			goto tr2
		case 59:
			goto tr3
		case 64:
			goto tr4
		case 95:
			goto st1
		case 100:
			goto tr5
		case 101:
			goto st170
		case 102:
			goto tr6
		case 104:
			goto tr7
		case 112:
			goto tr8
		case 120:
			goto tr9
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
tr3:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
	goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line xss_130.go:571
		switch data[p] {
		case 33:
			goto tr2
		case 59:
			goto tr3
		case 64:
			goto tr4
		case 95:
			goto st1
		case 98:
			goto st165
		case 100:
			goto tr5
		case 102:
			goto tr6
		case 104:
			goto tr7
		case 112:
			goto tr8
		case 120:
			goto tr9
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
tr4:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line xss_130.go:622
		switch data[p] {
		case 33:
			goto tr2
		case 59:
			goto tr3
		case 64:
			goto tr4
		case 95:
			goto st1
		case 100:
			goto tr5
		case 102:
			goto tr6
		case 104:
			goto tr7
		case 105:
			goto st23
		case 112:
			goto tr8
		case 120:
			goto tr9
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
tr5:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
	goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line xss_130.go:673
		switch data[p] {
		case 95:
			goto st1
		case 97:
			goto st7
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 95:
			goto st1
		case 116:
			goto st8
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		switch data[p] {
		case 95:
			goto st1
		case 97:
			goto st9
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		switch data[p] {
		case 58:
			goto st10
		case 95:
			goto st1
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch data[p] {
		case 33:
			goto tr2
		case 59:
			goto tr3
		case 64:
			goto tr4
		case 95:
			goto st1
		case 100:
			goto tr5
		case 102:
			goto tr6
		case 104:
			goto tr7
		case 112:
			goto tr8
		case 116:
			goto st176
		case 120:
			goto tr9
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
tr6:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
	goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line xss_130.go:820
		switch data[p] {
		case 95:
			goto st1
		case 111:
			goto st12
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch data[p] {
		case 95:
			goto st1
		case 114:
			goto st13
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		switch data[p] {
		case 95:
			goto st1
		case 109:
			goto st14
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		switch data[p] {
		case 95:
			goto st1
		case 97:
			goto st15
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		switch data[p] {
		case 95:
			goto st1
		case 99:
			goto st16
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		switch data[p] {
		case 95:
			goto st1
		case 116:
			goto st17
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		switch data[p] {
		case 95:
			goto st1
		case 105:
			goto st18
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		switch data[p] {
		case 95:
			goto st1
		case 111:
			goto st19
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		switch data[p] {
		case 95:
			goto st1
		case 110:
			goto st183
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st183:
		if p++; p == pe {
			goto _test_eof183
		}
	st_case_183:
		if data[p] == 95 {
			goto tr188
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr187
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st1
				}
			default:
				goto tr187
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st1
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr187
					}
				case data[p] >= 97:
					goto st1
				}
			default:
				goto tr187
			}
		default:
			goto tr187
		}
		goto st2
tr187:
//line xss_130.rl:35

            return true
        
	goto st184
	st184:
		if p++; p == pe {
			goto _test_eof184
		}
	st_case_184:
//line xss_130.go:1087
		switch data[p] {
		case 33:
			goto tr189
		case 59:
			goto tr190
		case 64:
			goto tr191
		case 95:
			goto tr188
		case 100:
			goto tr5
		case 102:
			goto tr6
		case 104:
			goto tr7
		case 112:
			goto tr8
		case 120:
			goto tr9
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr187
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st1
				}
			default:
				goto tr187
			}
		case data[p] > 63:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st1
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr187
					}
				case data[p] >= 97:
					goto st1
				}
			default:
				goto tr187
			}
		default:
			goto tr187
		}
		goto st2
tr189:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
//line xss_130.rl:35

            return true
        
	goto st185
	st185:
		if p++; p == pe {
			goto _test_eof185
		}
	st_case_185:
//line xss_130.go:1163
		switch data[p] {
		case 33:
			goto tr189
		case 59:
			goto tr190
		case 64:
			goto tr191
		case 95:
			goto tr188
		case 100:
			goto tr5
		case 101:
			goto st170
		case 102:
			goto tr6
		case 104:
			goto tr7
		case 112:
			goto tr8
		case 120:
			goto tr9
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr187
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st1
				}
			default:
				goto tr187
			}
		case data[p] > 63:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st1
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr187
					}
				case data[p] >= 97:
					goto st1
				}
			default:
				goto tr187
			}
		default:
			goto tr187
		}
		goto st2
tr190:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
//line xss_130.rl:35

            return true
        
	goto st186
	st186:
		if p++; p == pe {
			goto _test_eof186
		}
	st_case_186:
//line xss_130.go:1241
		switch data[p] {
		case 33:
			goto tr189
		case 59:
			goto tr190
		case 64:
			goto tr191
		case 95:
			goto tr188
		case 98:
			goto st165
		case 100:
			goto tr5
		case 102:
			goto tr6
		case 104:
			goto tr7
		case 112:
			goto tr8
		case 120:
			goto tr9
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr187
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st1
				}
			default:
				goto tr187
			}
		case data[p] > 63:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st1
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr187
					}
				case data[p] >= 97:
					goto st1
				}
			default:
				goto tr187
			}
		default:
			goto tr187
		}
		goto st2
tr191:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
//line xss_130.rl:35

            return true
        
	goto st187
	st187:
		if p++; p == pe {
			goto _test_eof187
		}
	st_case_187:
//line xss_130.go:1319
		switch data[p] {
		case 33:
			goto tr189
		case 59:
			goto tr190
		case 64:
			goto tr191
		case 95:
			goto tr188
		case 100:
			goto tr5
		case 102:
			goto tr6
		case 104:
			goto tr7
		case 105:
			goto st23
		case 112:
			goto tr8
		case 120:
			goto tr9
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr187
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st1
				}
			default:
				goto tr187
			}
		case data[p] > 63:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st1
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr187
					}
				case data[p] >= 97:
					goto st1
				}
			default:
				goto tr187
			}
		default:
			goto tr187
		}
		goto st2
tr188:
//line xss_130.rl:35

            return true
        
	goto st188
	st188:
		if p++; p == pe {
			goto _test_eof188
		}
	st_case_188:
//line xss_130.go:1389
		if data[p] == 95 {
			goto tr188
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr187
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st1
				}
			default:
				goto tr187
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st1
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr187
					}
				case data[p] >= 97:
					goto st1
				}
			default:
				goto tr187
			}
		default:
			goto tr187
		}
		goto st2
tr7:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
	goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line xss_130.go:1444
		switch data[p] {
		case 95:
			goto st1
		case 116:
			goto st21
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		switch data[p] {
		case 95:
			goto st1
		case 116:
			goto st22
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		switch data[p] {
		case 95:
			goto st1
		case 112:
			goto st183
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		switch data[p] {
		case 95:
			goto st1
		case 109:
			goto st24
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		switch data[p] {
		case 95:
			goto st1
		case 112:
			goto st25
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		switch data[p] {
		case 95:
			goto st1
		case 111:
			goto st26
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		switch data[p] {
		case 95:
			goto st1
		case 114:
			goto st27
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		switch data[p] {
		case 95:
			goto st1
		case 116:
			goto st183
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
tr8:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
	goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line xss_130.go:1647
		switch data[p] {
		case 95:
			goto st1
		case 97:
			goto st29
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		switch data[p] {
		case 95:
			goto st1
		case 116:
			goto st30
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		switch data[p] {
		case 95:
			goto st1
		case 116:
			goto st31
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		switch data[p] {
		case 95:
			goto st1
		case 101:
			goto st32
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		switch data[p] {
		case 95:
			goto st1
		case 114:
			goto st33
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		switch data[p] {
		case 95:
			goto st1
		case 110:
			goto st34
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		switch data[p] {
		case 61:
			goto st190
		case 95:
			goto st34
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto st35
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st37
				}
			default:
				goto st35
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st37
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st35
					}
				case data[p] >= 97:
					goto st37
				}
			default:
				goto st35
			}
		default:
			goto st35
		}
		goto st2
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		switch data[p] {
		case 33:
			goto tr42
		case 59:
			goto tr43
		case 61:
			goto st190
		case 64:
			goto tr44
		case 95:
			goto st34
		case 100:
			goto tr45
		case 102:
			goto tr46
		case 104:
			goto tr47
		case 112:
			goto tr48
		case 120:
			goto tr49
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto st35
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st37
				}
			default:
				goto st35
			}
		case data[p] > 63:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st37
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st35
					}
				case data[p] >= 97:
					goto st37
				}
			default:
				goto st35
			}
		default:
			goto st35
		}
		goto st2
tr42:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
	goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
//line xss_130.go:1912
		switch data[p] {
		case 33:
			goto tr42
		case 59:
			goto tr43
		case 61:
			goto st190
		case 64:
			goto tr44
		case 95:
			goto st34
		case 100:
			goto tr45
		case 101:
			goto st97
		case 102:
			goto tr46
		case 104:
			goto tr47
		case 112:
			goto tr48
		case 120:
			goto tr49
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto st35
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st37
				}
			default:
				goto st35
			}
		case data[p] > 63:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st37
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st35
					}
				case data[p] >= 97:
					goto st37
				}
			default:
				goto st35
			}
		default:
			goto st35
		}
		goto st2
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
tr198:
//line xss_130.rl:35

            return true
        
	goto st189
	st189:
		if p++; p == pe {
			goto _test_eof189
		}
	st_case_189:
//line xss_130.go:2008
		switch data[p] {
		case 33:
			goto tr189
		case 59:
			goto tr190
		case 64:
			goto tr191
		case 95:
			goto tr188
		case 100:
			goto tr5
		case 102:
			goto tr6
		case 104:
			goto tr7
		case 112:
			goto tr8
		case 120:
			goto tr9
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr187
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st1
				}
			default:
				goto tr187
			}
		case data[p] > 63:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st1
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr187
					}
				case data[p] >= 97:
					goto st1
				}
			default:
				goto tr187
			}
		default:
			goto tr187
		}
		goto st2
tr9:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
	goto st38
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
//line xss_130.go:2080
		switch data[p] {
		case 95:
			goto st1
		case 104:
			goto st39
		case 108:
			goto st42
		case 109:
			goto st50
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		switch data[p] {
		case 95:
			goto st1
		case 116:
			goto st40
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		switch data[p] {
		case 95:
			goto st1
		case 109:
			goto st41
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		switch data[p] {
		case 95:
			goto st1
		case 108:
			goto st183
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		switch data[p] {
		case 95:
			goto st1
		case 105:
			goto st43
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		switch data[p] {
		case 95:
			goto st1
		case 110:
			goto st44
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
		switch data[p] {
		case 95:
			goto st1
		case 107:
			goto st45
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		switch data[p] {
		case 58:
			goto st46
		case 95:
			goto st1
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		switch data[p] {
		case 33:
			goto tr2
		case 59:
			goto tr3
		case 64:
			goto tr4
		case 95:
			goto st1
		case 100:
			goto tr5
		case 102:
			goto tr6
		case 104:
			goto tr61
		case 112:
			goto tr8
		case 120:
			goto tr9
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
tr61:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
	goto st47
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
//line xss_130.go:2325
		switch data[p] {
		case 95:
			goto st1
		case 114:
			goto st48
		case 116:
			goto st21
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		switch data[p] {
		case 95:
			goto st1
		case 101:
			goto st49
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		switch data[p] {
		case 95:
			goto st1
		case 102:
			goto st183
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		switch data[p] {
		case 95:
			goto st1
		case 108:
			goto st51
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
		switch data[p] {
		case 95:
			goto st1
		case 110:
			goto st52
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
		switch data[p] {
		case 95:
			goto st1
		case 115:
			goto st183
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
tr43:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
	goto st53
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
//line xss_130.go:2482
		switch data[p] {
		case 33:
			goto tr42
		case 59:
			goto tr43
		case 61:
			goto st190
		case 64:
			goto tr44
		case 95:
			goto st34
		case 98:
			goto st92
		case 100:
			goto tr45
		case 102:
			goto tr46
		case 104:
			goto tr47
		case 112:
			goto tr48
		case 120:
			goto tr49
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto st35
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st37
				}
			default:
				goto st35
			}
		case data[p] > 63:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st37
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st35
					}
				case data[p] >= 97:
					goto st37
				}
			default:
				goto st35
			}
		default:
			goto st35
		}
		goto st2
tr195:
//line xss_130.rl:35

            return true
        
	goto st190
	st190:
		if p++; p == pe {
			goto _test_eof190
		}
	st_case_190:
//line xss_130.go:2554
		switch data[p] {
		case 33:
			goto tr193
		case 59:
			goto tr194
		case 61:
			goto tr195
		case 64:
			goto tr196
		case 95:
			goto tr197
		case 100:
			goto tr45
		case 102:
			goto tr46
		case 104:
			goto tr47
		case 112:
			goto tr48
		case 120:
			goto tr49
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr192
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st37
				}
			default:
				goto tr192
			}
		case data[p] > 63:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st37
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr192
					}
				case data[p] >= 97:
					goto st37
				}
			default:
				goto tr192
			}
		default:
			goto tr192
		}
		goto st2
tr192:
//line xss_130.rl:35

            return true
        
	goto st191
	st191:
		if p++; p == pe {
			goto _test_eof191
		}
	st_case_191:
//line xss_130.go:2624
		switch data[p] {
		case 33:
			goto tr193
		case 59:
			goto tr194
		case 61:
			goto tr195
		case 64:
			goto tr196
		case 95:
			goto tr197
		case 100:
			goto tr45
		case 102:
			goto tr46
		case 104:
			goto tr47
		case 112:
			goto tr48
		case 120:
			goto tr49
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr192
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st37
				}
			default:
				goto tr192
			}
		case data[p] > 63:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st37
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr192
					}
				case data[p] >= 97:
					goto st37
				}
			default:
				goto tr192
			}
		default:
			goto tr192
		}
		goto st2
tr193:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
//line xss_130.rl:35

            return true
        
	goto st192
	st192:
		if p++; p == pe {
			goto _test_eof192
		}
	st_case_192:
//line xss_130.go:2702
		switch data[p] {
		case 33:
			goto tr193
		case 59:
			goto tr194
		case 61:
			goto tr195
		case 64:
			goto tr196
		case 95:
			goto tr197
		case 100:
			goto tr45
		case 101:
			goto st97
		case 102:
			goto tr46
		case 104:
			goto tr47
		case 112:
			goto tr48
		case 120:
			goto tr49
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr192
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st37
				}
			default:
				goto tr192
			}
		case data[p] > 63:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st37
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr192
					}
				case data[p] >= 97:
					goto st37
				}
			default:
				goto tr192
			}
		default:
			goto tr192
		}
		goto st2
tr194:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
//line xss_130.rl:35

            return true
        
	goto st193
	st193:
		if p++; p == pe {
			goto _test_eof193
		}
	st_case_193:
//line xss_130.go:2782
		switch data[p] {
		case 33:
			goto tr193
		case 59:
			goto tr194
		case 61:
			goto tr195
		case 64:
			goto tr196
		case 95:
			goto tr197
		case 98:
			goto st92
		case 100:
			goto tr45
		case 102:
			goto tr46
		case 104:
			goto tr47
		case 112:
			goto tr48
		case 120:
			goto tr49
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr192
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st37
				}
			default:
				goto tr192
			}
		case data[p] > 63:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st37
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr192
					}
				case data[p] >= 97:
					goto st37
				}
			default:
				goto tr192
			}
		default:
			goto tr192
		}
		goto st2
tr196:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
//line xss_130.rl:35

            return true
        
	goto st194
	st194:
		if p++; p == pe {
			goto _test_eof194
		}
	st_case_194:
//line xss_130.go:2862
		switch data[p] {
		case 33:
			goto tr193
		case 59:
			goto tr194
		case 61:
			goto tr195
		case 64:
			goto tr196
		case 95:
			goto tr197
		case 100:
			goto tr45
		case 102:
			goto tr46
		case 104:
			goto tr47
		case 105:
			goto st70
		case 112:
			goto tr48
		case 120:
			goto tr49
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr192
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st37
				}
			default:
				goto tr192
			}
		case data[p] > 63:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st37
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr192
					}
				case data[p] >= 97:
					goto st37
				}
			default:
				goto tr192
			}
		default:
			goto tr192
		}
		goto st2
tr197:
//line xss_130.rl:35

            return true
        
	goto st195
	st195:
		if p++; p == pe {
			goto _test_eof195
		}
	st_case_195:
//line xss_130.go:2934
		switch data[p] {
		case 61:
			goto tr195
		case 95:
			goto tr197
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr192
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st37
				}
			default:
				goto tr192
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st37
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr192
					}
				case data[p] >= 97:
					goto st37
				}
			default:
				goto tr192
			}
		default:
			goto tr192
		}
		goto st2
tr45:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
	goto st54
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
//line xss_130.go:2992
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 97:
			goto st55
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 116:
			goto st56
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 97:
			goto st57
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		switch data[p] {
		case 58:
			goto st10
		case 61:
			goto st189
		case 95:
			goto st37
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
tr46:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
	goto st58
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
//line xss_130.go:3107
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 111:
			goto st59
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 114:
			goto st60
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 109:
			goto st61
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 97:
			goto st62
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 99:
			goto st63
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 116:
			goto st64
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 105:
			goto st65
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 111:
			goto st66
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 110:
			goto st196
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st196:
		if p++; p == pe {
			goto _test_eof196
		}
	st_case_196:
		switch data[p] {
		case 61:
			goto tr198
		case 95:
			goto tr199
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr187
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st37
				}
			default:
				goto tr187
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st37
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr187
					}
				case data[p] >= 97:
					goto st37
				}
			default:
				goto tr187
			}
		default:
			goto tr187
		}
		goto st2
tr199:
//line xss_130.rl:35

            return true
        
	goto st197
	st197:
		if p++; p == pe {
			goto _test_eof197
		}
	st_case_197:
//line xss_130.go:3395
		switch data[p] {
		case 61:
			goto tr198
		case 95:
			goto tr199
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr187
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st37
				}
			default:
				goto tr187
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st37
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr187
					}
				case data[p] >= 97:
					goto st37
				}
			default:
				goto tr187
			}
		default:
			goto tr187
		}
		goto st2
tr47:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
	goto st67
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
//line xss_130.go:3453
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 116:
			goto st68
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 116:
			goto st69
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 112:
			goto st196
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 109:
			goto st71
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 112:
			goto st72
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 111:
			goto st73
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 114:
			goto st74
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 116:
			goto st196
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
tr48:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
	goto st75
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
//line xss_130.go:3672
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 97:
			goto st76
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 116:
			goto st77
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 116:
			goto st78
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 101:
			goto st79
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 114:
			goto st80
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 110:
			goto st34
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
tr49:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
	goto st81
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
//line xss_130.go:3839
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 104:
			goto st82
		case 108:
			goto st85
		case 109:
			goto st89
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 116:
			goto st83
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 109:
			goto st84
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 108:
			goto st196
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 105:
			goto st86
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 110:
			goto st87
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 107:
			goto st88
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
		switch data[p] {
		case 58:
			goto st46
		case 61:
			goto st189
		case 95:
			goto st37
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 108:
			goto st90
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 110:
			goto st91
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 115:
			goto st196
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 97:
			goto st93
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 115:
			goto st94
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 101:
			goto st95
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
		switch data[p] {
		case 54:
			goto st96
		case 61:
			goto st189
		case 95:
			goto st37
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
		switch data[p] {
		case 52:
			goto st196
		case 61:
			goto st189
		case 95:
			goto st37
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 110:
			goto st98
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 116:
			goto st99
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 105:
			goto st100
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 116:
			goto st101
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 121:
			goto st102
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
		switch data[p] {
		case 32:
			goto st103
		case 61:
			goto st189
		case 95:
			goto st37
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st103
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st37
				}
			case data[p] >= 65:
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
		switch data[p] {
		case 32:
			goto st103
		case 33:
			goto tr2
		case 37:
			goto st104
		case 59:
			goto tr3
		case 64:
			goto tr4
		case 95:
			goto st106
		case 100:
			goto tr112
		case 102:
			goto tr113
		case 104:
			goto tr114
		case 112:
			goto tr115
		case 120:
			goto tr116
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st103
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st1
		}
		goto st2
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
		switch data[p] {
		case 32:
			goto st105
		case 33:
			goto tr2
		case 59:
			goto tr3
		case 64:
			goto tr4
		case 95:
			goto st1
		case 100:
			goto tr5
		case 102:
			goto tr6
		case 104:
			goto tr7
		case 112:
			goto tr8
		case 120:
			goto tr9
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st105
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st1
				}
			case data[p] >= 65:
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
		switch data[p] {
		case 32:
			goto st105
		case 33:
			goto tr2
		case 59:
			goto tr3
		case 64:
			goto tr4
		case 95:
			goto st106
		case 100:
			goto tr112
		case 102:
			goto tr113
		case 104:
			goto tr114
		case 112:
			goto tr115
		case 120:
			goto tr116
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st105
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st1
		}
		goto st2
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
		switch data[p] {
		case 32:
			goto st107
		case 33:
			goto tr2
		case 59:
			goto tr3
		case 64:
			goto tr4
		case 95:
			goto st1
		case 100:
			goto tr5
		case 102:
			goto tr6
		case 104:
			goto tr7
		case 112:
			goto tr119
		case 115:
			goto st113
		case 120:
			goto tr9
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st1
				}
			case data[p] >= 65:
				goto st1
			}
		default:
			goto st1
		}
		goto st2
tr119:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
	goto st108
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
//line xss_130.go:4644
		switch data[p] {
		case 95:
			goto st1
		case 97:
			goto st29
		case 117:
			goto st109
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
		switch data[p] {
		case 95:
			goto st1
		case 98:
			goto st110
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
		switch data[p] {
		case 95:
			goto st1
		case 108:
			goto st111
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
		switch data[p] {
		case 95:
			goto st1
		case 105:
			goto st112
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
		switch data[p] {
		case 95:
			goto st1
		case 99:
			goto st183
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
		switch data[p] {
		case 95:
			goto st1
		case 121:
			goto st114
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
		switch data[p] {
		case 95:
			goto st1
		case 115:
			goto st115
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
		switch data[p] {
		case 95:
			goto st1
		case 116:
			goto st116
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
		switch data[p] {
		case 95:
			goto st1
		case 101:
			goto st117
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
		switch data[p] {
		case 95:
			goto st1
		case 109:
			goto st183
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
tr112:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
	goto st118
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
//line xss_130.go:4897
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 97:
			goto st119
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 116:
			goto st120
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 97:
			goto st121
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
		switch data[p] {
		case 32:
			goto st107
		case 58:
			goto st10
		case 95:
			goto st106
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
tr113:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
	goto st122
	st122:
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
//line xss_130.go:5032
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 111:
			goto st123
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 114:
			goto st124
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 109:
			goto st125
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 97:
			goto st126
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 99:
			goto st127
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 116:
			goto st128
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 105:
			goto st129
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 111:
			goto st130
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 110:
			goto st198
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st198:
		if p++; p == pe {
			goto _test_eof198
		}
	st_case_198:
		switch data[p] {
		case 32:
			goto tr200
		case 95:
			goto tr201
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 33:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr200
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st106
				}
			default:
				goto tr187
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st106
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr187
					}
				case data[p] >= 97:
					goto st106
				}
			default:
				goto tr187
			}
		default:
			goto tr187
		}
		goto st2
tr200:
//line xss_130.rl:35

            return true
        
	goto st199
	st199:
		if p++; p == pe {
			goto _test_eof199
		}
	st_case_199:
//line xss_130.go:5365
		switch data[p] {
		case 32:
			goto tr200
		case 33:
			goto tr189
		case 59:
			goto tr190
		case 64:
			goto tr191
		case 95:
			goto tr188
		case 100:
			goto tr5
		case 102:
			goto tr6
		case 104:
			goto tr7
		case 112:
			goto tr119
		case 115:
			goto st113
		case 120:
			goto tr9
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 34:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr200
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st1
				}
			default:
				goto tr187
			}
		case data[p] > 63:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st1
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr187
					}
				case data[p] >= 97:
					goto st1
				}
			default:
				goto tr187
			}
		default:
			goto tr187
		}
		goto st2
tr201:
//line xss_130.rl:35

            return true
        
	goto st200
	st200:
		if p++; p == pe {
			goto _test_eof200
		}
	st_case_200:
//line xss_130.go:5437
		switch data[p] {
		case 32:
			goto tr200
		case 95:
			goto tr201
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 33:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr200
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st106
				}
			default:
				goto tr187
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st106
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr187
					}
				case data[p] >= 97:
					goto st106
				}
			default:
				goto tr187
			}
		default:
			goto tr187
		}
		goto st2
tr114:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
	goto st131
	st131:
		if p++; p == pe {
			goto _test_eof131
		}
	st_case_131:
//line xss_130.go:5495
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 116:
			goto st132
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 116:
			goto st133
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 112:
			goto st198
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
tr115:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
	goto st134
	st134:
		if p++; p == pe {
			goto _test_eof134
		}
	st_case_134:
//line xss_130.go:5599
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 97:
			goto st135
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st135:
		if p++; p == pe {
			goto _test_eof135
		}
	st_case_135:
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 116:
			goto st136
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st136:
		if p++; p == pe {
			goto _test_eof136
		}
	st_case_136:
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 116:
			goto st137
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st137:
		if p++; p == pe {
			goto _test_eof137
		}
	st_case_137:
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 101:
			goto st138
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st138:
		if p++; p == pe {
			goto _test_eof138
		}
	st_case_138:
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 114:
			goto st139
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st139:
		if p++; p == pe {
			goto _test_eof139
		}
	st_case_139:
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 110:
			goto st140
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st140:
		if p++; p == pe {
			goto _test_eof140
		}
	st_case_140:
		switch data[p] {
		case 32:
			goto st141
		case 61:
			goto st190
		case 95:
			goto st140
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 33:
				if 9 <= data[p] && data[p] <= 13 {
					goto st141
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st153
				}
			default:
				goto st35
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st153
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st35
					}
				case data[p] >= 97:
					goto st153
				}
			default:
				goto st35
			}
		default:
			goto st35
		}
		goto st2
	st141:
		if p++; p == pe {
			goto _test_eof141
		}
	st_case_141:
		switch data[p] {
		case 32:
			goto st141
		case 33:
			goto tr42
		case 59:
			goto tr43
		case 61:
			goto st190
		case 64:
			goto tr44
		case 95:
			goto st34
		case 100:
			goto tr45
		case 102:
			goto tr46
		case 104:
			goto tr47
		case 112:
			goto tr151
		case 115:
			goto st148
		case 120:
			goto tr49
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 34:
				if 9 <= data[p] && data[p] <= 13 {
					goto st141
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st37
				}
			default:
				goto st35
			}
		case data[p] > 63:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st37
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st35
					}
				case data[p] >= 97:
					goto st37
				}
			default:
				goto st35
			}
		default:
			goto st35
		}
		goto st2
tr44:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
	goto st142
	st142:
		if p++; p == pe {
			goto _test_eof142
		}
	st_case_142:
//line xss_130.go:5912
		switch data[p] {
		case 33:
			goto tr42
		case 59:
			goto tr43
		case 61:
			goto st190
		case 64:
			goto tr44
		case 95:
			goto st34
		case 100:
			goto tr45
		case 102:
			goto tr46
		case 104:
			goto tr47
		case 105:
			goto st70
		case 112:
			goto tr48
		case 120:
			goto tr49
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto st35
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st37
				}
			default:
				goto st35
			}
		case data[p] > 63:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st37
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st35
					}
				case data[p] >= 97:
					goto st37
				}
			default:
				goto st35
			}
		default:
			goto st35
		}
		goto st2
tr151:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
	goto st143
	st143:
		if p++; p == pe {
			goto _test_eof143
		}
	st_case_143:
//line xss_130.go:5988
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 97:
			goto st76
		case 117:
			goto st144
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st144:
		if p++; p == pe {
			goto _test_eof144
		}
	st_case_144:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 98:
			goto st145
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st145:
		if p++; p == pe {
			goto _test_eof145
		}
	st_case_145:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 108:
			goto st146
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st146:
		if p++; p == pe {
			goto _test_eof146
		}
	st_case_146:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 105:
			goto st147
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st147:
		if p++; p == pe {
			goto _test_eof147
		}
	st_case_147:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 99:
			goto st196
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st148:
		if p++; p == pe {
			goto _test_eof148
		}
	st_case_148:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 121:
			goto st149
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st149:
		if p++; p == pe {
			goto _test_eof149
		}
	st_case_149:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 115:
			goto st150
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st150:
		if p++; p == pe {
			goto _test_eof150
		}
	st_case_150:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 116:
			goto st151
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st151:
		if p++; p == pe {
			goto _test_eof151
		}
	st_case_151:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 101:
			goto st152
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st152:
		if p++; p == pe {
			goto _test_eof152
		}
	st_case_152:
		switch data[p] {
		case 61:
			goto st189
		case 95:
			goto st37
		case 109:
			goto st196
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		goto st2
	st153:
		if p++; p == pe {
			goto _test_eof153
		}
	st_case_153:
		switch data[p] {
		case 32:
			goto st107
		case 61:
			goto st189
		case 95:
			goto st153
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st153
				}
			case data[p] >= 65:
				goto st153
			}
		default:
			goto st153
		}
		goto st2
tr116:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
	goto st154
	st154:
		if p++; p == pe {
			goto _test_eof154
		}
	st_case_154:
//line xss_130.go:6292
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 104:
			goto st155
		case 108:
			goto st158
		case 109:
			goto st162
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st155:
		if p++; p == pe {
			goto _test_eof155
		}
	st_case_155:
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 116:
			goto st156
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st156:
		if p++; p == pe {
			goto _test_eof156
		}
	st_case_156:
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 109:
			goto st157
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st157:
		if p++; p == pe {
			goto _test_eof157
		}
	st_case_157:
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 108:
			goto st198
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st158:
		if p++; p == pe {
			goto _test_eof158
		}
	st_case_158:
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 105:
			goto st159
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st159:
		if p++; p == pe {
			goto _test_eof159
		}
	st_case_159:
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 110:
			goto st160
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st160:
		if p++; p == pe {
			goto _test_eof160
		}
	st_case_160:
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 107:
			goto st161
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st161:
		if p++; p == pe {
			goto _test_eof161
		}
	st_case_161:
		switch data[p] {
		case 32:
			goto st107
		case 58:
			goto st46
		case 95:
			goto st106
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st162:
		if p++; p == pe {
			goto _test_eof162
		}
	st_case_162:
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 108:
			goto st163
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st163:
		if p++; p == pe {
			goto _test_eof163
		}
	st_case_163:
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 110:
			goto st164
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st164:
		if p++; p == pe {
			goto _test_eof164
		}
	st_case_164:
		switch data[p] {
		case 32:
			goto st107
		case 95:
			goto st106
		case 115:
			goto st198
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st107
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st106
				}
			case data[p] >= 65:
				goto st106
			}
		default:
			goto st106
		}
		goto st2
	st165:
		if p++; p == pe {
			goto _test_eof165
		}
	st_case_165:
		switch data[p] {
		case 95:
			goto st1
		case 97:
			goto st166
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st166:
		if p++; p == pe {
			goto _test_eof166
		}
	st_case_166:
		switch data[p] {
		case 95:
			goto st1
		case 115:
			goto st167
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st167:
		if p++; p == pe {
			goto _test_eof167
		}
	st_case_167:
		switch data[p] {
		case 95:
			goto st1
		case 101:
			goto st168
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st168:
		if p++; p == pe {
			goto _test_eof168
		}
	st_case_168:
		switch data[p] {
		case 54:
			goto st169
		case 95:
			goto st1
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st169:
		if p++; p == pe {
			goto _test_eof169
		}
	st_case_169:
		switch data[p] {
		case 52:
			goto st183
		case 95:
			goto st1
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st170:
		if p++; p == pe {
			goto _test_eof170
		}
	st_case_170:
		switch data[p] {
		case 95:
			goto st1
		case 110:
			goto st171
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st171:
		if p++; p == pe {
			goto _test_eof171
		}
	st_case_171:
		switch data[p] {
		case 95:
			goto st1
		case 116:
			goto st172
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st172:
		if p++; p == pe {
			goto _test_eof172
		}
	st_case_172:
		switch data[p] {
		case 95:
			goto st1
		case 105:
			goto st173
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st173:
		if p++; p == pe {
			goto _test_eof173
		}
	st_case_173:
		switch data[p] {
		case 95:
			goto st1
		case 116:
			goto st174
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st174:
		if p++; p == pe {
			goto _test_eof174
		}
	st_case_174:
		switch data[p] {
		case 95:
			goto st1
		case 121:
			goto st175
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st175:
		if p++; p == pe {
			goto _test_eof175
		}
	st_case_175:
		switch data[p] {
		case 32:
			goto st103
		case 95:
			goto st1
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st103
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st1
				}
			case data[p] >= 65:
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st176:
		if p++; p == pe {
			goto _test_eof176
		}
	st_case_176:
		switch data[p] {
		case 95:
			goto st1
		case 101:
			goto st177
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st177:
		if p++; p == pe {
			goto _test_eof177
		}
	st_case_177:
		switch data[p] {
		case 95:
			goto st1
		case 120:
			goto st178
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st178:
		if p++; p == pe {
			goto _test_eof178
		}
	st_case_178:
		switch data[p] {
		case 95:
			goto st1
		case 116:
			goto st179
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st179:
		if p++; p == pe {
			goto _test_eof179
		}
	st_case_179:
		switch data[p] {
		case 47:
			goto st180
		case 95:
			goto st1
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st180:
		if p++; p == pe {
			goto _test_eof180
		}
	st_case_180:
		switch data[p] {
		case 33:
			goto tr2
		case 59:
			goto tr3
		case 64:
			goto tr4
		case 95:
			goto st1
		case 100:
			goto tr5
		case 102:
			goto tr6
		case 104:
			goto tr185
		case 112:
			goto tr8
		case 120:
			goto tr9
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
tr185:
//line xss_130.rl:27

            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        
	goto st181
	st181:
		if p++; p == pe {
			goto _test_eof181
		}
	st_case_181:
//line xss_130.go:7051
		switch data[p] {
		case 95:
			goto st1
		case 116:
			goto st182
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st182:
		if p++; p == pe {
			goto _test_eof182
		}
	st_case_182:
		switch data[p] {
		case 95:
			goto st1
		case 109:
			goto st41
		case 116:
			goto st22
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st2
	st_out:
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
	_test_eof183: cs = 183; goto _test_eof
	_test_eof184: cs = 184; goto _test_eof
	_test_eof185: cs = 185; goto _test_eof
	_test_eof186: cs = 186; goto _test_eof
	_test_eof187: cs = 187; goto _test_eof
	_test_eof188: cs = 188; goto _test_eof
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
	_test_eof189: cs = 189; goto _test_eof
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
	_test_eof48: cs = 48; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof
	_test_eof53: cs = 53; goto _test_eof
	_test_eof190: cs = 190; goto _test_eof
	_test_eof191: cs = 191; goto _test_eof
	_test_eof192: cs = 192; goto _test_eof
	_test_eof193: cs = 193; goto _test_eof
	_test_eof194: cs = 194; goto _test_eof
	_test_eof195: cs = 195; goto _test_eof
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
	_test_eof65: cs = 65; goto _test_eof
	_test_eof66: cs = 66; goto _test_eof
	_test_eof196: cs = 196; goto _test_eof
	_test_eof197: cs = 197; goto _test_eof
	_test_eof67: cs = 67; goto _test_eof
	_test_eof68: cs = 68; goto _test_eof
	_test_eof69: cs = 69; goto _test_eof
	_test_eof70: cs = 70; goto _test_eof
	_test_eof71: cs = 71; goto _test_eof
	_test_eof72: cs = 72; goto _test_eof
	_test_eof73: cs = 73; goto _test_eof
	_test_eof74: cs = 74; goto _test_eof
	_test_eof75: cs = 75; goto _test_eof
	_test_eof76: cs = 76; goto _test_eof
	_test_eof77: cs = 77; goto _test_eof
	_test_eof78: cs = 78; goto _test_eof
	_test_eof79: cs = 79; goto _test_eof
	_test_eof80: cs = 80; goto _test_eof
	_test_eof81: cs = 81; goto _test_eof
	_test_eof82: cs = 82; goto _test_eof
	_test_eof83: cs = 83; goto _test_eof
	_test_eof84: cs = 84; goto _test_eof
	_test_eof85: cs = 85; goto _test_eof
	_test_eof86: cs = 86; goto _test_eof
	_test_eof87: cs = 87; goto _test_eof
	_test_eof88: cs = 88; goto _test_eof
	_test_eof89: cs = 89; goto _test_eof
	_test_eof90: cs = 90; goto _test_eof
	_test_eof91: cs = 91; goto _test_eof
	_test_eof92: cs = 92; goto _test_eof
	_test_eof93: cs = 93; goto _test_eof
	_test_eof94: cs = 94; goto _test_eof
	_test_eof95: cs = 95; goto _test_eof
	_test_eof96: cs = 96; goto _test_eof
	_test_eof97: cs = 97; goto _test_eof
	_test_eof98: cs = 98; goto _test_eof
	_test_eof99: cs = 99; goto _test_eof
	_test_eof100: cs = 100; goto _test_eof
	_test_eof101: cs = 101; goto _test_eof
	_test_eof102: cs = 102; goto _test_eof
	_test_eof103: cs = 103; goto _test_eof
	_test_eof104: cs = 104; goto _test_eof
	_test_eof105: cs = 105; goto _test_eof
	_test_eof106: cs = 106; goto _test_eof
	_test_eof107: cs = 107; goto _test_eof
	_test_eof108: cs = 108; goto _test_eof
	_test_eof109: cs = 109; goto _test_eof
	_test_eof110: cs = 110; goto _test_eof
	_test_eof111: cs = 111; goto _test_eof
	_test_eof112: cs = 112; goto _test_eof
	_test_eof113: cs = 113; goto _test_eof
	_test_eof114: cs = 114; goto _test_eof
	_test_eof115: cs = 115; goto _test_eof
	_test_eof116: cs = 116; goto _test_eof
	_test_eof117: cs = 117; goto _test_eof
	_test_eof118: cs = 118; goto _test_eof
	_test_eof119: cs = 119; goto _test_eof
	_test_eof120: cs = 120; goto _test_eof
	_test_eof121: cs = 121; goto _test_eof
	_test_eof122: cs = 122; goto _test_eof
	_test_eof123: cs = 123; goto _test_eof
	_test_eof124: cs = 124; goto _test_eof
	_test_eof125: cs = 125; goto _test_eof
	_test_eof126: cs = 126; goto _test_eof
	_test_eof127: cs = 127; goto _test_eof
	_test_eof128: cs = 128; goto _test_eof
	_test_eof129: cs = 129; goto _test_eof
	_test_eof130: cs = 130; goto _test_eof
	_test_eof198: cs = 198; goto _test_eof
	_test_eof199: cs = 199; goto _test_eof
	_test_eof200: cs = 200; goto _test_eof
	_test_eof131: cs = 131; goto _test_eof
	_test_eof132: cs = 132; goto _test_eof
	_test_eof133: cs = 133; goto _test_eof
	_test_eof134: cs = 134; goto _test_eof
	_test_eof135: cs = 135; goto _test_eof
	_test_eof136: cs = 136; goto _test_eof
	_test_eof137: cs = 137; goto _test_eof
	_test_eof138: cs = 138; goto _test_eof
	_test_eof139: cs = 139; goto _test_eof
	_test_eof140: cs = 140; goto _test_eof
	_test_eof141: cs = 141; goto _test_eof
	_test_eof142: cs = 142; goto _test_eof
	_test_eof143: cs = 143; goto _test_eof
	_test_eof144: cs = 144; goto _test_eof
	_test_eof145: cs = 145; goto _test_eof
	_test_eof146: cs = 146; goto _test_eof
	_test_eof147: cs = 147; goto _test_eof
	_test_eof148: cs = 148; goto _test_eof
	_test_eof149: cs = 149; goto _test_eof
	_test_eof150: cs = 150; goto _test_eof
	_test_eof151: cs = 151; goto _test_eof
	_test_eof152: cs = 152; goto _test_eof
	_test_eof153: cs = 153; goto _test_eof
	_test_eof154: cs = 154; goto _test_eof
	_test_eof155: cs = 155; goto _test_eof
	_test_eof156: cs = 156; goto _test_eof
	_test_eof157: cs = 157; goto _test_eof
	_test_eof158: cs = 158; goto _test_eof
	_test_eof159: cs = 159; goto _test_eof
	_test_eof160: cs = 160; goto _test_eof
	_test_eof161: cs = 161; goto _test_eof
	_test_eof162: cs = 162; goto _test_eof
	_test_eof163: cs = 163; goto _test_eof
	_test_eof164: cs = 164; goto _test_eof
	_test_eof165: cs = 165; goto _test_eof
	_test_eof166: cs = 166; goto _test_eof
	_test_eof167: cs = 167; goto _test_eof
	_test_eof168: cs = 168; goto _test_eof
	_test_eof169: cs = 169; goto _test_eof
	_test_eof170: cs = 170; goto _test_eof
	_test_eof171: cs = 171; goto _test_eof
	_test_eof172: cs = 172; goto _test_eof
	_test_eof173: cs = 173; goto _test_eof
	_test_eof174: cs = 174; goto _test_eof
	_test_eof175: cs = 175; goto _test_eof
	_test_eof176: cs = 176; goto _test_eof
	_test_eof177: cs = 177; goto _test_eof
	_test_eof178: cs = 178; goto _test_eof
	_test_eof179: cs = 179; goto _test_eof
	_test_eof180: cs = 180; goto _test_eof
	_test_eof181: cs = 181; goto _test_eof
	_test_eof182: cs = 182; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 183, 189, 190, 196, 198:
//line xss_130.rl:35

            return true
        
//line xss_130.go:7307
		}
	}

	}

//line xss_130.rl:53

        return false
}
