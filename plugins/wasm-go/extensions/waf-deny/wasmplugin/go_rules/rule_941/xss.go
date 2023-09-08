
//line xss.rl:1
package rule_941

func matchXSS(data []byte) bool {

//line xss.rl:5

//line xss.go:10
const xss_start int = 0
const xss_first_final int = 208
const xss_error int = -1

const xss_en_main int = 0


//line xss.rl:6
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof
    
//line xss.go:22
	{
	cs = xss_start
	}

//line xss.go:27
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
	case 208:
		goto st_case_208
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
	case 209:
		goto st_case_209
	case 210:
		goto st_case_210
	case 211:
		goto st_case_211
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
	case 48:
		goto st_case_48
	case 212:
		goto st_case_212
	case 49:
		goto st_case_49
	case 213:
		goto st_case_213
	case 214:
		goto st_case_214
	case 215:
		goto st_case_215
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
	case 216:
		goto st_case_216
	case 217:
		goto st_case_217
	case 218:
		goto st_case_218
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
	case 219:
		goto st_case_219
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
	case 220:
		goto st_case_220
	case 221:
		goto st_case_221
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
	case 222:
		goto st_case_222
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
	case 223:
		goto st_case_223
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
	case 224:
		goto st_case_224
	case 225:
		goto st_case_225
	case 226:
		goto st_case_226
	case 227:
		goto st_case_227
	case 228:
		goto st_case_228
	case 229:
		goto st_case_229
	case 230:
		goto st_case_230
	case 231:
		goto st_case_231
	case 232:
		goto st_case_232
	case 233:
		goto st_case_233
	case 234:
		goto st_case_234
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
	case 235:
		goto st_case_235
	case 236:
		goto st_case_236
	case 172:
		goto st_case_172
	case 237:
		goto st_case_237
	case 238:
		goto st_case_238
	case 173:
		goto st_case_173
	case 239:
		goto st_case_239
	case 240:
		goto st_case_240
	case 241:
		goto st_case_241
	case 174:
		goto st_case_174
	case 242:
		goto st_case_242
	case 243:
		goto st_case_243
	case 244:
		goto st_case_244
	case 245:
		goto st_case_245
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
	case 246:
		goto st_case_246
	case 247:
		goto st_case_247
	case 248:
		goto st_case_248
	case 249:
		goto st_case_249
	case 250:
		goto st_case_250
	case 181:
		goto st_case_181
	case 182:
		goto st_case_182
	case 183:
		goto st_case_183
	case 251:
		goto st_case_251
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
	case 189:
		goto st_case_189
	case 252:
		goto st_case_252
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
	case 253:
		goto st_case_253
	case 254:
		goto st_case_254
	case 196:
		goto st_case_196
	case 255:
		goto st_case_255
	case 256:
		goto st_case_256
	case 197:
		goto st_case_197
	case 198:
		goto st_case_198
	case 199:
		goto st_case_199
	case 200:
		goto st_case_200
	case 201:
		goto st_case_201
	case 202:
		goto st_case_202
	case 203:
		goto st_case_203
	case 204:
		goto st_case_204
	case 205:
		goto st_case_205
	case 206:
		goto st_case_206
	case 207:
		goto st_case_207
	}
	goto st_out
	st_case_0:
		switch data[p] {
		case 33:
			goto st10
		case 59:
			goto st11
		case 60:
			goto st12
		case 64:
			goto st13
		case 100:
			goto st152
		case 102:
			goto st156
		case 112:
			goto st165
		case 120:
			goto st197
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] > 13:
				if 32 <= data[p] && data[p] <= 47 {
					goto st9
				}
			case data[p] >= 9:
				goto st9
			}
		case data[p] > 63:
			switch {
			case data[p] < 97:
				if 91 <= data[p] && data[p] <= 96 {
					goto st9
				}
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 126 {
					goto st9
				}
			default:
				goto st94
			}
		default:
			goto st9
		}
		goto st1
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		if data[p] == 60 {
			goto st2
		}
		goto st1
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 60:
			goto st2
		case 115:
			goto st3
		}
		goto st1
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 60:
			goto st2
		case 99:
			goto st4
		}
		goto st1
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 60:
			goto st2
		case 114:
			goto st5
		}
		goto st1
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 60:
			goto st2
		case 105:
			goto st6
		}
		goto st1
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch data[p] {
		case 60:
			goto st2
		case 112:
			goto st7
		}
		goto st1
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 60:
			goto st2
		case 116:
			goto st8
		}
		goto st1
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		if data[p] == 62 {
			goto tr18
		}
		goto st8
tr18:
//line xss.rl:10

            return true
        
	goto st208
	st208:
		if p++; p == pe {
			goto _test_eof208
		}
	st_case_208:
//line xss.go:696
		if data[p] == 62 {
			goto tr18
		}
		goto st8
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		switch data[p] {
		case 33:
			goto st10
		case 59:
			goto st11
		case 60:
			goto st12
		case 64:
			goto st13
		case 100:
			goto st14
		case 102:
			goto st27
		case 112:
			goto st41
		case 120:
			goto st56
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] > 13:
				if 32 <= data[p] && data[p] <= 47 {
					goto st9
				}
			case data[p] >= 9:
				goto st9
			}
		case data[p] > 63:
			switch {
			case data[p] > 96:
				if 123 <= data[p] && data[p] <= 126 {
					goto st9
				}
			case data[p] >= 91:
				goto st9
			}
		default:
			goto st9
		}
		goto st1
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch data[p] {
		case 33:
			goto st10
		case 59:
			goto st11
		case 60:
			goto st12
		case 64:
			goto st13
		case 69:
			goto st73
		case 100:
			goto st14
		case 101:
			goto st73
		case 102:
			goto st27
		case 112:
			goto st41
		case 120:
			goto st56
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] > 13:
				if 32 <= data[p] && data[p] <= 47 {
					goto st9
				}
			case data[p] >= 9:
				goto st9
			}
		case data[p] > 63:
			switch {
			case data[p] > 96:
				if 123 <= data[p] && data[p] <= 126 {
					goto st9
				}
			case data[p] >= 91:
				goto st9
			}
		default:
			goto st9
		}
		goto st1
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch data[p] {
		case 33:
			goto st10
		case 59:
			goto st11
		case 60:
			goto st12
		case 64:
			goto st13
		case 98:
			goto st68
		case 100:
			goto st14
		case 102:
			goto st27
		case 112:
			goto st41
		case 120:
			goto st56
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] > 13:
				if 32 <= data[p] && data[p] <= 47 {
					goto st9
				}
			case data[p] >= 9:
				goto st9
			}
		case data[p] > 63:
			switch {
			case data[p] > 96:
				if 123 <= data[p] && data[p] <= 126 {
					goto st9
				}
			case data[p] >= 91:
				goto st9
			}
		default:
			goto st9
		}
		goto st1
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch data[p] {
		case 33:
			goto st10
		case 59:
			goto st11
		case 60:
			goto st12
		case 64:
			goto st13
		case 100:
			goto st14
		case 102:
			goto st27
		case 112:
			goto st41
		case 115:
			goto st3
		case 120:
			goto st56
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] > 13:
				if 32 <= data[p] && data[p] <= 47 {
					goto st9
				}
			case data[p] >= 9:
				goto st9
			}
		case data[p] > 63:
			switch {
			case data[p] > 96:
				if 123 <= data[p] && data[p] <= 126 {
					goto st9
				}
			case data[p] >= 91:
				goto st9
			}
		default:
			goto st9
		}
		goto st1
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		switch data[p] {
		case 33:
			goto st10
		case 59:
			goto st11
		case 60:
			goto st12
		case 64:
			goto st13
		case 100:
			goto st14
		case 102:
			goto st27
		case 105:
			goto st36
		case 112:
			goto st41
		case 120:
			goto st56
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] > 13:
				if 32 <= data[p] && data[p] <= 47 {
					goto st9
				}
			case data[p] >= 9:
				goto st9
			}
		case data[p] > 63:
			switch {
			case data[p] > 96:
				if 123 <= data[p] && data[p] <= 126 {
					goto st9
				}
			case data[p] >= 91:
				goto st9
			}
		default:
			goto st9
		}
		goto st1
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		switch data[p] {
		case 60:
			goto st2
		case 97:
			goto st15
		}
		goto st1
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		switch data[p] {
		case 60:
			goto st2
		case 116:
			goto st16
		}
		goto st1
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		switch data[p] {
		case 60:
			goto st2
		case 97:
			goto st17
		}
		goto st1
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		switch data[p] {
		case 58:
			goto st18
		case 60:
			goto st2
		}
		goto st1
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		switch data[p] {
		case 60:
			goto st2
		case 116:
			goto st19
		}
		goto st1
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		switch data[p] {
		case 60:
			goto st2
		case 101:
			goto st20
		}
		goto st1
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		switch data[p] {
		case 60:
			goto st2
		case 120:
			goto st21
		}
		goto st1
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		switch data[p] {
		case 60:
			goto st2
		case 116:
			goto st22
		}
		goto st1
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		switch data[p] {
		case 47:
			goto st23
		case 60:
			goto st2
		}
		goto st1
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		switch data[p] {
		case 60:
			goto st2
		case 104:
			goto st24
		}
		goto st1
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		switch data[p] {
		case 60:
			goto st2
		case 116:
			goto st25
		}
		goto st1
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		switch data[p] {
		case 60:
			goto st2
		case 109:
			goto st26
		}
		goto st1
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		switch data[p] {
		case 60:
			goto st2
		case 108:
			goto st209
		}
		goto st1
	st209:
		if p++; p == pe {
			goto _test_eof209
		}
	st_case_209:
		if data[p] == 60 {
			goto tr219
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] > 13:
				if 32 <= data[p] && data[p] <= 47 {
					goto tr218
				}
			case data[p] >= 9:
				goto tr218
			}
		case data[p] > 64:
			switch {
			case data[p] > 96:
				if 123 <= data[p] && data[p] <= 126 {
					goto tr218
				}
			case data[p] >= 91:
				goto tr218
			}
		default:
			goto tr218
		}
		goto st1
tr218:
//line xss.rl:10

            return true
        
	goto st210
	st210:
		if p++; p == pe {
			goto _test_eof210
		}
	st_case_210:
//line xss.go:1139
		if data[p] == 60 {
			goto tr219
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] > 13:
				if 32 <= data[p] && data[p] <= 47 {
					goto tr218
				}
			case data[p] >= 9:
				goto tr218
			}
		case data[p] > 64:
			switch {
			case data[p] > 96:
				if 123 <= data[p] && data[p] <= 126 {
					goto tr218
				}
			case data[p] >= 91:
				goto tr218
			}
		default:
			goto tr218
		}
		goto st1
tr219:
//line xss.rl:10

            return true
        
	goto st211
	st211:
		if p++; p == pe {
			goto _test_eof211
		}
	st_case_211:
//line xss.go:1177
		switch data[p] {
		case 60:
			goto tr219
		case 115:
			goto st3
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] > 13:
				if 32 <= data[p] && data[p] <= 47 {
					goto tr218
				}
			case data[p] >= 9:
				goto tr218
			}
		case data[p] > 64:
			switch {
			case data[p] > 96:
				if 123 <= data[p] && data[p] <= 126 {
					goto tr218
				}
			case data[p] >= 91:
				goto tr218
			}
		default:
			goto tr218
		}
		goto st1
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		switch data[p] {
		case 60:
			goto st2
		case 111:
			goto st28
		}
		goto st1
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		switch data[p] {
		case 60:
			goto st2
		case 114:
			goto st29
		}
		goto st1
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		switch data[p] {
		case 60:
			goto st2
		case 109:
			goto st30
		}
		goto st1
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		switch data[p] {
		case 60:
			goto st2
		case 97:
			goto st31
		}
		goto st1
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		switch data[p] {
		case 60:
			goto st2
		case 99:
			goto st32
		}
		goto st1
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		switch data[p] {
		case 60:
			goto st2
		case 116:
			goto st33
		}
		goto st1
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		switch data[p] {
		case 60:
			goto st2
		case 105:
			goto st34
		}
		goto st1
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		switch data[p] {
		case 60:
			goto st2
		case 111:
			goto st35
		}
		goto st1
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		switch data[p] {
		case 60:
			goto st2
		case 110:
			goto st209
		}
		goto st1
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		switch data[p] {
		case 60:
			goto st2
		case 109:
			goto st37
		}
		goto st1
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		switch data[p] {
		case 60:
			goto st2
		case 112:
			goto st38
		}
		goto st1
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		switch data[p] {
		case 60:
			goto st2
		case 111:
			goto st39
		}
		goto st1
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		switch data[p] {
		case 60:
			goto st2
		case 114:
			goto st40
		}
		goto st1
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		switch data[p] {
		case 60:
			goto st2
		case 116:
			goto st209
		}
		goto st1
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		switch data[p] {
		case 60:
			goto st2
		case 97:
			goto st42
		}
		goto st1
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		switch data[p] {
		case 60:
			goto st2
		case 116:
			goto st43
		}
		goto st1
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		switch data[p] {
		case 60:
			goto st2
		case 116:
			goto st44
		}
		goto st1
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
		switch data[p] {
		case 60:
			goto st2
		case 101:
			goto st45
		}
		goto st1
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		switch data[p] {
		case 60:
			goto st2
		case 114:
			goto st46
		}
		goto st1
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		switch data[p] {
		case 60:
			goto st2
		case 110:
			goto st47
		}
		goto st1
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		switch data[p] {
		case 60:
			goto st49
		case 61:
			goto st213
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto st47
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st48
				}
			default:
				goto st47
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st48
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st47
					}
				case data[p] >= 97:
					goto st48
				}
			default:
				goto st47
			}
		default:
			goto st47
		}
		goto st1
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st212
		case 95:
			goto st48
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st48
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st48
			}
		default:
			goto st48
		}
		goto st1
tr220:
//line xss.rl:10

            return true
        
	goto st212
	st212:
		if p++; p == pe {
			goto _test_eof212
		}
	st_case_212:
//line xss.go:1531
		switch data[p] {
		case 60:
			goto tr219
		case 95:
			goto tr220
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr218
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st212
				}
			default:
				goto tr218
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st212
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr218
					}
				case data[p] >= 97:
					goto st212
				}
			default:
				goto tr218
			}
		default:
			goto tr218
		}
		goto st1
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		switch data[p] {
		case 60:
			goto st49
		case 61:
			goto st213
		case 115:
			goto st50
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto st47
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st48
				}
			default:
				goto st47
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st48
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st47
					}
				case data[p] >= 97:
					goto st48
				}
			default:
				goto st47
			}
		default:
			goto st47
		}
		goto st1
tr224:
//line xss.rl:10

            return true
        
	goto st213
	st213:
		if p++; p == pe {
			goto _test_eof213
		}
	st_case_213:
//line xss.go:1634
		switch data[p] {
		case 60:
			goto tr223
		case 61:
			goto tr224
		case 95:
			goto tr224
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr221
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st218
				}
			default:
				goto tr221
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st218
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr221
					}
				case data[p] >= 97:
					goto st218
				}
			default:
				goto tr221
			}
		default:
			goto tr221
		}
		goto st1
tr221:
//line xss.rl:10

            return true
        
	goto st214
	st214:
		if p++; p == pe {
			goto _test_eof214
		}
	st_case_214:
//line xss.go:1690
		switch data[p] {
		case 60:
			goto tr223
		case 61:
			goto tr224
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr221
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st48
				}
			default:
				goto tr221
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st48
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr221
					}
				case data[p] >= 97:
					goto st48
				}
			default:
				goto tr221
			}
		default:
			goto tr221
		}
		goto st1
tr223:
//line xss.rl:10

            return true
        
	goto st215
	st215:
		if p++; p == pe {
			goto _test_eof215
		}
	st_case_215:
//line xss.go:1744
		switch data[p] {
		case 60:
			goto tr223
		case 61:
			goto tr224
		case 115:
			goto st50
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr221
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st48
				}
			default:
				goto tr221
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st48
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr221
					}
				case data[p] >= 97:
					goto st48
				}
			default:
				goto tr221
			}
		default:
			goto tr221
		}
		goto st1
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st212
		case 95:
			goto st48
		case 99:
			goto st51
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st48
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st48
			}
		default:
			goto st48
		}
		goto st1
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st212
		case 95:
			goto st48
		case 114:
			goto st52
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st48
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st48
			}
		default:
			goto st48
		}
		goto st1
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st212
		case 95:
			goto st48
		case 105:
			goto st53
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st48
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st48
			}
		default:
			goto st48
		}
		goto st1
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st212
		case 95:
			goto st48
		case 112:
			goto st54
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st48
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st48
			}
		default:
			goto st48
		}
		goto st1
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st212
		case 95:
			goto st48
		case 116:
			goto st55
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st48
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st48
			}
		default:
			goto st48
		}
		goto st1
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
		switch data[p] {
		case 61:
			goto st216
		case 62:
			goto tr18
		case 95:
			goto st55
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st55
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st55
			}
		default:
			goto st55
		}
		goto st8
tr226:
//line xss.rl:10

            return true
        
	goto st216
	st216:
		if p++; p == pe {
			goto _test_eof216
		}
	st_case_216:
//line xss.go:1966
		if data[p] == 95 {
			goto tr226
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr225
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st216
				}
			default:
				goto tr225
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st216
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr225
					}
				case data[p] >= 97:
					goto st216
				}
			default:
				goto tr225
			}
		default:
			goto tr225
		}
		goto st8
tr225:
//line xss.rl:10

            return true
        
	goto st217
	st217:
		if p++; p == pe {
			goto _test_eof217
		}
	st_case_217:
//line xss.go:2017
		switch {
		case data[p] < 58:
			switch {
			case data[p] > 13:
				if 32 <= data[p] && data[p] <= 47 {
					goto tr225
				}
			case data[p] >= 9:
				goto tr225
			}
		case data[p] > 64:
			switch {
			case data[p] > 96:
				if 123 <= data[p] && data[p] <= 126 {
					goto tr225
				}
			case data[p] >= 91:
				goto tr225
			}
		default:
			goto tr225
		}
		goto st8
tr227:
//line xss.rl:10

            return true
        
	goto st218
	st218:
		if p++; p == pe {
			goto _test_eof218
		}
	st_case_218:
//line xss.go:2052
		switch data[p] {
		case 60:
			goto tr219
		case 61:
			goto tr220
		case 95:
			goto tr227
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr218
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st218
				}
			default:
				goto tr218
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st218
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr218
					}
				case data[p] >= 97:
					goto st218
				}
			default:
				goto tr218
			}
		default:
			goto tr218
		}
		goto st1
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		switch data[p] {
		case 60:
			goto st2
		case 104:
			goto st24
		case 108:
			goto st57
		case 109:
			goto st65
		}
		goto st1
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		switch data[p] {
		case 60:
			goto st2
		case 105:
			goto st58
		}
		goto st1
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		switch data[p] {
		case 60:
			goto st2
		case 110:
			goto st59
		}
		goto st1
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		switch data[p] {
		case 60:
			goto st2
		case 107:
			goto st60
		}
		goto st1
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
		switch data[p] {
		case 58:
			goto st61
		case 60:
			goto st2
		}
		goto st1
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		switch data[p] {
		case 60:
			goto st2
		case 104:
			goto st62
		}
		goto st1
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		switch data[p] {
		case 60:
			goto st2
		case 114:
			goto st63
		}
		goto st1
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		switch data[p] {
		case 60:
			goto st2
		case 101:
			goto st64
		}
		goto st1
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		switch data[p] {
		case 60:
			goto st2
		case 102:
			goto st209
		}
		goto st1
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		switch data[p] {
		case 60:
			goto st2
		case 108:
			goto st66
		}
		goto st1
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
		switch data[p] {
		case 60:
			goto st2
		case 110:
			goto st67
		}
		goto st1
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
		switch data[p] {
		case 60:
			goto st2
		case 115:
			goto st209
		}
		goto st1
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
		switch data[p] {
		case 60:
			goto st2
		case 97:
			goto st69
		}
		goto st1
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		switch data[p] {
		case 60:
			goto st2
		case 115:
			goto st70
		}
		goto st1
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
		switch data[p] {
		case 60:
			goto st2
		case 101:
			goto st71
		}
		goto st1
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
		switch data[p] {
		case 54:
			goto st72
		case 60:
			goto st2
		}
		goto st1
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
		switch data[p] {
		case 52:
			goto st209
		case 60:
			goto st2
		}
		goto st1
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
		switch data[p] {
		case 60:
			goto st2
		case 78:
			goto st74
		case 110:
			goto st74
		}
		goto st1
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
		switch data[p] {
		case 60:
			goto st2
		case 84:
			goto st75
		case 116:
			goto st75
		}
		goto st1
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
		switch data[p] {
		case 60:
			goto st2
		case 73:
			goto st76
		case 105:
			goto st76
		}
		goto st1
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
		switch data[p] {
		case 60:
			goto st2
		case 84:
			goto st77
		case 116:
			goto st77
		}
		goto st1
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
		switch data[p] {
		case 60:
			goto st2
		case 89:
			goto st78
		case 121:
			goto st78
		}
		goto st1
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
		switch data[p] {
		case 32:
			goto st79
		case 60:
			goto st2
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st79
		}
		goto st1
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
		switch data[p] {
		case 32:
			goto st79
		case 37:
			goto st80
		case 60:
			goto st2
		case 95:
			goto st82
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st79
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st82
			}
		default:
			goto st82
		}
		goto st1
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
		switch data[p] {
		case 32:
			goto st81
		case 60:
			goto st2
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st81
		}
		goto st1
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
		switch data[p] {
		case 32:
			goto st81
		case 60:
			goto st2
		case 95:
			goto st82
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st81
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st82
			}
		default:
			goto st82
		}
		goto st1
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
		switch data[p] {
		case 32:
			goto st83
		case 60:
			goto st2
		case 95:
			goto st82
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st83
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st82
				}
			case data[p] >= 65:
				goto st82
			}
		default:
			goto st82
		}
		goto st1
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
		switch data[p] {
		case 32:
			goto st83
		case 60:
			goto st2
		case 80:
			goto st84
		case 83:
			goto st89
		case 112:
			goto st84
		case 115:
			goto st89
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st83
		}
		goto st1
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
		switch data[p] {
		case 60:
			goto st2
		case 85:
			goto st85
		case 117:
			goto st85
		}
		goto st1
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
		switch data[p] {
		case 60:
			goto st2
		case 66:
			goto st86
		case 98:
			goto st86
		}
		goto st1
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
		switch data[p] {
		case 60:
			goto st2
		case 76:
			goto st87
		case 108:
			goto st87
		}
		goto st1
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
		switch data[p] {
		case 60:
			goto st2
		case 73:
			goto st88
		case 105:
			goto st88
		}
		goto st1
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
		switch data[p] {
		case 60:
			goto st2
		case 67:
			goto st209
		case 99:
			goto st209
		}
		goto st1
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
		switch data[p] {
		case 60:
			goto st2
		case 89:
			goto st90
		case 121:
			goto st90
		}
		goto st1
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
		switch data[p] {
		case 60:
			goto st2
		case 83:
			goto st91
		case 115:
			goto st91
		}
		goto st1
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
		switch data[p] {
		case 60:
			goto st2
		case 84:
			goto st92
		case 116:
			goto st92
		}
		goto st1
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
		switch data[p] {
		case 60:
			goto st2
		case 69:
			goto st93
		case 101:
			goto st93
		}
		goto st1
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
		switch data[p] {
		case 60:
			goto st2
		case 77:
			goto st209
		case 109:
			goto st209
		}
		goto st1
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st95
		}
		if 97 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
		switch data[p] {
		case 58:
			goto st1
		case 60:
			goto st144
		case 61:
			goto st1
		}
		goto st96
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
		switch data[p] {
		case 58:
			goto st97
		case 60:
			goto st144
		case 61:
			goto st1
		}
		goto st96
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
		switch data[p] {
		case 60:
			goto st102
		case 117:
			goto st125
		}
		goto st98
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
		switch data[p] {
		case 59:
			goto st99
		case 60:
			goto st102
		}
		goto st98
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
		switch data[p] {
		case 58:
			goto st98
		case 60:
			goto st138
		case 61:
			goto st98
		}
		goto st100
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
		switch data[p] {
		case 58:
			goto st101
		case 60:
			goto st138
		case 61:
			goto st98
		}
		goto st100
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
		switch data[p] {
		case 59:
			goto st99
		case 60:
			goto st102
		case 117:
			goto st125
		}
		goto st98
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
		switch data[p] {
		case 59:
			goto st99
		case 60:
			goto st102
		case 115:
			goto st103
		}
		goto st98
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
		switch data[p] {
		case 59:
			goto st99
		case 60:
			goto st102
		case 99:
			goto st104
		}
		goto st98
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
		switch data[p] {
		case 59:
			goto st99
		case 60:
			goto st102
		case 114:
			goto st105
		}
		goto st98
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
		switch data[p] {
		case 59:
			goto st99
		case 60:
			goto st102
		case 105:
			goto st106
		}
		goto st98
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
		switch data[p] {
		case 59:
			goto st99
		case 60:
			goto st102
		case 112:
			goto st107
		}
		goto st98
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
		switch data[p] {
		case 59:
			goto st99
		case 60:
			goto st102
		case 116:
			goto st108
		}
		goto st98
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
		switch data[p] {
		case 59:
			goto st109
		case 62:
			goto tr121
		}
		goto st108
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
		switch data[p] {
		case 58:
			goto st108
		case 61:
			goto st108
		case 62:
			goto tr123
		}
		goto st110
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
		switch data[p] {
		case 58:
			goto st111
		case 61:
			goto st108
		case 62:
			goto tr123
		}
		goto st110
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
		switch data[p] {
		case 59:
			goto st109
		case 62:
			goto tr121
		case 117:
			goto st112
		}
		goto st108
tr121:
//line xss.rl:10

            return true
        
	goto st219
	st219:
		if p++; p == pe {
			goto _test_eof219
		}
	st_case_219:
//line xss.go:2911
		switch data[p] {
		case 59:
			goto st109
		case 62:
			goto tr121
		}
		goto st108
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
		switch data[p] {
		case 59:
			goto st109
		case 62:
			goto tr121
		case 114:
			goto st113
		}
		goto st108
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
		switch data[p] {
		case 59:
			goto st109
		case 62:
			goto tr121
		case 108:
			goto st114
		}
		goto st108
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
		switch data[p] {
		case 40:
			goto st115
		case 59:
			goto st109
		case 62:
			goto tr121
		}
		goto st108
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
		switch data[p] {
		case 59:
			goto st109
		case 62:
			goto tr121
		case 106:
			goto st116
		}
		goto st108
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
		switch data[p] {
		case 59:
			goto st109
		case 62:
			goto tr121
		case 97:
			goto st117
		}
		goto st108
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
		switch data[p] {
		case 59:
			goto st109
		case 62:
			goto tr121
		case 118:
			goto st118
		}
		goto st108
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
		switch data[p] {
		case 59:
			goto st109
		case 62:
			goto tr121
		case 97:
			goto st119
		}
		goto st108
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
		switch data[p] {
		case 59:
			goto st109
		case 62:
			goto tr121
		case 115:
			goto st120
		}
		goto st108
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
		switch data[p] {
		case 59:
			goto st109
		case 62:
			goto tr121
		case 99:
			goto st121
		}
		goto st108
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
		switch data[p] {
		case 59:
			goto st109
		case 62:
			goto tr121
		case 114:
			goto st122
		}
		goto st108
	st122:
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
		switch data[p] {
		case 59:
			goto st109
		case 62:
			goto tr121
		case 105:
			goto st123
		}
		goto st108
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
		switch data[p] {
		case 59:
			goto st109
		case 62:
			goto tr121
		case 112:
			goto st124
		}
		goto st108
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
		switch data[p] {
		case 59:
			goto st109
		case 62:
			goto tr121
		case 116:
			goto st220
		}
		goto st108
	st220:
		if p++; p == pe {
			goto _test_eof220
		}
	st_case_220:
		switch data[p] {
		case 59:
			goto st109
		case 62:
			goto tr121
		}
		goto st108
tr123:
//line xss.rl:10

            return true
        
	goto st221
	st221:
		if p++; p == pe {
			goto _test_eof221
		}
	st_case_221:
//line xss.go:3124
		switch data[p] {
		case 58:
			goto st111
		case 61:
			goto st108
		case 62:
			goto tr123
		}
		goto st110
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
		switch data[p] {
		case 59:
			goto st99
		case 60:
			goto st102
		case 114:
			goto st126
		}
		goto st98
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
		switch data[p] {
		case 59:
			goto st99
		case 60:
			goto st102
		case 108:
			goto st127
		}
		goto st98
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
		switch data[p] {
		case 40:
			goto st128
		case 59:
			goto st99
		case 60:
			goto st102
		}
		goto st98
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
		switch data[p] {
		case 59:
			goto st99
		case 60:
			goto st102
		case 106:
			goto st129
		}
		goto st98
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
		switch data[p] {
		case 59:
			goto st99
		case 60:
			goto st102
		case 97:
			goto st130
		}
		goto st98
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
		switch data[p] {
		case 59:
			goto st99
		case 60:
			goto st102
		case 118:
			goto st131
		}
		goto st98
	st131:
		if p++; p == pe {
			goto _test_eof131
		}
	st_case_131:
		switch data[p] {
		case 59:
			goto st99
		case 60:
			goto st102
		case 97:
			goto st132
		}
		goto st98
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
		switch data[p] {
		case 59:
			goto st99
		case 60:
			goto st102
		case 115:
			goto st133
		}
		goto st98
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
		switch data[p] {
		case 59:
			goto st99
		case 60:
			goto st102
		case 99:
			goto st134
		}
		goto st98
	st134:
		if p++; p == pe {
			goto _test_eof134
		}
	st_case_134:
		switch data[p] {
		case 59:
			goto st99
		case 60:
			goto st102
		case 114:
			goto st135
		}
		goto st98
	st135:
		if p++; p == pe {
			goto _test_eof135
		}
	st_case_135:
		switch data[p] {
		case 59:
			goto st99
		case 60:
			goto st102
		case 105:
			goto st136
		}
		goto st98
	st136:
		if p++; p == pe {
			goto _test_eof136
		}
	st_case_136:
		switch data[p] {
		case 59:
			goto st99
		case 60:
			goto st102
		case 112:
			goto st137
		}
		goto st98
	st137:
		if p++; p == pe {
			goto _test_eof137
		}
	st_case_137:
		switch data[p] {
		case 59:
			goto st99
		case 60:
			goto st102
		case 116:
			goto st222
		}
		goto st98
	st222:
		if p++; p == pe {
			goto _test_eof222
		}
	st_case_222:
		switch data[p] {
		case 59:
			goto st99
		case 60:
			goto st102
		}
		goto st98
	st138:
		if p++; p == pe {
			goto _test_eof138
		}
	st_case_138:
		switch data[p] {
		case 58:
			goto st101
		case 60:
			goto st138
		case 61:
			goto st98
		case 115:
			goto st139
		}
		goto st100
	st139:
		if p++; p == pe {
			goto _test_eof139
		}
	st_case_139:
		switch data[p] {
		case 58:
			goto st101
		case 60:
			goto st138
		case 61:
			goto st98
		case 99:
			goto st140
		}
		goto st100
	st140:
		if p++; p == pe {
			goto _test_eof140
		}
	st_case_140:
		switch data[p] {
		case 58:
			goto st101
		case 60:
			goto st138
		case 61:
			goto st98
		case 114:
			goto st141
		}
		goto st100
	st141:
		if p++; p == pe {
			goto _test_eof141
		}
	st_case_141:
		switch data[p] {
		case 58:
			goto st101
		case 60:
			goto st138
		case 61:
			goto st98
		case 105:
			goto st142
		}
		goto st100
	st142:
		if p++; p == pe {
			goto _test_eof142
		}
	st_case_142:
		switch data[p] {
		case 58:
			goto st101
		case 60:
			goto st138
		case 61:
			goto st98
		case 112:
			goto st143
		}
		goto st100
	st143:
		if p++; p == pe {
			goto _test_eof143
		}
	st_case_143:
		switch data[p] {
		case 58:
			goto st101
		case 60:
			goto st138
		case 61:
			goto st98
		case 116:
			goto st110
		}
		goto st100
	st144:
		if p++; p == pe {
			goto _test_eof144
		}
	st_case_144:
		switch data[p] {
		case 58:
			goto st97
		case 60:
			goto st144
		case 61:
			goto st1
		case 115:
			goto st145
		}
		goto st96
	st145:
		if p++; p == pe {
			goto _test_eof145
		}
	st_case_145:
		switch data[p] {
		case 58:
			goto st97
		case 60:
			goto st144
		case 61:
			goto st1
		case 99:
			goto st146
		}
		goto st96
	st146:
		if p++; p == pe {
			goto _test_eof146
		}
	st_case_146:
		switch data[p] {
		case 58:
			goto st97
		case 60:
			goto st144
		case 61:
			goto st1
		case 114:
			goto st147
		}
		goto st96
	st147:
		if p++; p == pe {
			goto _test_eof147
		}
	st_case_147:
		switch data[p] {
		case 58:
			goto st97
		case 60:
			goto st144
		case 61:
			goto st1
		case 105:
			goto st148
		}
		goto st96
	st148:
		if p++; p == pe {
			goto _test_eof148
		}
	st_case_148:
		switch data[p] {
		case 58:
			goto st97
		case 60:
			goto st144
		case 61:
			goto st1
		case 112:
			goto st149
		}
		goto st96
	st149:
		if p++; p == pe {
			goto _test_eof149
		}
	st_case_149:
		switch data[p] {
		case 58:
			goto st97
		case 60:
			goto st144
		case 61:
			goto st1
		case 116:
			goto st150
		}
		goto st96
	st150:
		if p++; p == pe {
			goto _test_eof150
		}
	st_case_150:
		switch data[p] {
		case 58:
			goto st151
		case 61:
			goto st8
		case 62:
			goto tr164
		}
		goto st150
	st151:
		if p++; p == pe {
			goto _test_eof151
		}
	st_case_151:
		switch data[p] {
		case 62:
			goto tr121
		case 117:
			goto st112
		}
		goto st108
tr164:
//line xss.rl:10

            return true
        
	goto st223
	st223:
		if p++; p == pe {
			goto _test_eof223
		}
	st_case_223:
//line xss.go:3557
		switch data[p] {
		case 58:
			goto st151
		case 61:
			goto st8
		case 62:
			goto tr164
		}
		goto st150
	st152:
		if p++; p == pe {
			goto _test_eof152
		}
	st_case_152:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st95
		case 97:
			goto st153
		}
		if 98 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st153:
		if p++; p == pe {
			goto _test_eof153
		}
	st_case_153:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st95
		case 116:
			goto st154
		}
		if 97 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st154:
		if p++; p == pe {
			goto _test_eof154
		}
	st_case_154:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st95
		case 97:
			goto st155
		}
		if 98 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st155:
		if p++; p == pe {
			goto _test_eof155
		}
	st_case_155:
		switch data[p] {
		case 58:
			goto st18
		case 60:
			goto st2
		case 61:
			goto st95
		}
		if 97 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st156:
		if p++; p == pe {
			goto _test_eof156
		}
	st_case_156:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st95
		case 111:
			goto st157
		}
		if 97 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st157:
		if p++; p == pe {
			goto _test_eof157
		}
	st_case_157:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st95
		case 114:
			goto st158
		}
		if 97 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st158:
		if p++; p == pe {
			goto _test_eof158
		}
	st_case_158:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st95
		case 109:
			goto st159
		}
		if 97 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st159:
		if p++; p == pe {
			goto _test_eof159
		}
	st_case_159:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st95
		case 97:
			goto st160
		}
		if 98 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st160:
		if p++; p == pe {
			goto _test_eof160
		}
	st_case_160:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st95
		case 99:
			goto st161
		}
		if 97 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st161:
		if p++; p == pe {
			goto _test_eof161
		}
	st_case_161:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st95
		case 116:
			goto st162
		}
		if 97 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st162:
		if p++; p == pe {
			goto _test_eof162
		}
	st_case_162:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st95
		case 105:
			goto st163
		}
		if 97 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st163:
		if p++; p == pe {
			goto _test_eof163
		}
	st_case_163:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st95
		case 111:
			goto st164
		}
		if 97 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st164:
		if p++; p == pe {
			goto _test_eof164
		}
	st_case_164:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st95
		case 110:
			goto st224
		}
		if 97 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st224:
		if p++; p == pe {
			goto _test_eof224
		}
	st_case_224:
		switch data[p] {
		case 60:
			goto tr219
		case 61:
			goto tr228
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] > 13:
				if 32 <= data[p] && data[p] <= 47 {
					goto tr218
				}
			case data[p] >= 9:
				goto tr218
			}
		case data[p] > 64:
			switch {
			case data[p] < 97:
				if 91 <= data[p] && data[p] <= 96 {
					goto tr218
				}
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 126 {
					goto tr218
				}
			default:
				goto st94
			}
		default:
			goto tr218
		}
		goto st1
tr228:
//line xss.rl:10

            return true
        
	goto st225
	st225:
		if p++; p == pe {
			goto _test_eof225
		}
	st_case_225:
//line xss.go:3837
		switch data[p] {
		case 59:
			goto tr229
		case 60:
			goto tr230
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] > 13:
				if 32 <= data[p] && data[p] <= 47 {
					goto tr229
				}
			case data[p] >= 9:
				goto tr229
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 62 <= data[p] && data[p] <= 64 {
					goto tr229
				}
			case data[p] > 96:
				if 123 <= data[p] && data[p] <= 126 {
					goto tr229
				}
			default:
				goto tr229
			}
		default:
			goto tr218
		}
		goto st96
tr229:
//line xss.rl:10

            return true
        
	goto st226
	st226:
		if p++; p == pe {
			goto _test_eof226
		}
	st_case_226:
//line xss.go:3882
		switch data[p] {
		case 58:
			goto tr231
		case 60:
			goto tr230
		case 61:
			goto tr218
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] > 13:
				if 32 <= data[p] && data[p] <= 47 {
					goto tr229
				}
			case data[p] >= 9:
				goto tr229
			}
		case data[p] > 64:
			switch {
			case data[p] > 96:
				if 123 <= data[p] && data[p] <= 126 {
					goto tr229
				}
			case data[p] >= 91:
				goto tr229
			}
		default:
			goto tr229
		}
		goto st96
tr231:
//line xss.rl:10

            return true
        
	goto st227
	st227:
		if p++; p == pe {
			goto _test_eof227
		}
	st_case_227:
//line xss.go:3925
		switch data[p] {
		case 60:
			goto tr233
		case 117:
			goto st125
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] > 13:
				if 32 <= data[p] && data[p] <= 47 {
					goto tr232
				}
			case data[p] >= 9:
				goto tr232
			}
		case data[p] > 64:
			switch {
			case data[p] > 96:
				if 123 <= data[p] && data[p] <= 126 {
					goto tr232
				}
			case data[p] >= 91:
				goto tr232
			}
		default:
			goto tr232
		}
		goto st98
tr232:
//line xss.rl:10

            return true
        
	goto st228
	st228:
		if p++; p == pe {
			goto _test_eof228
		}
	st_case_228:
//line xss.go:3966
		switch data[p] {
		case 59:
			goto tr234
		case 60:
			goto tr233
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] > 13:
				if 32 <= data[p] && data[p] <= 47 {
					goto tr232
				}
			case data[p] >= 9:
				goto tr232
			}
		case data[p] > 64:
			switch {
			case data[p] > 96:
				if 123 <= data[p] && data[p] <= 126 {
					goto tr232
				}
			case data[p] >= 91:
				goto tr232
			}
		default:
			goto tr232
		}
		goto st98
tr234:
//line xss.rl:10

            return true
        
	goto st229
	st229:
		if p++; p == pe {
			goto _test_eof229
		}
	st_case_229:
//line xss.go:4007
		switch data[p] {
		case 59:
			goto tr235
		case 60:
			goto tr236
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] > 13:
				if 32 <= data[p] && data[p] <= 47 {
					goto tr235
				}
			case data[p] >= 9:
				goto tr235
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 62 <= data[p] && data[p] <= 64 {
					goto tr235
				}
			case data[p] > 96:
				if 123 <= data[p] && data[p] <= 126 {
					goto tr235
				}
			default:
				goto tr235
			}
		default:
			goto tr232
		}
		goto st100
tr235:
//line xss.rl:10

            return true
        
	goto st230
	st230:
		if p++; p == pe {
			goto _test_eof230
		}
	st_case_230:
//line xss.go:4052
		switch data[p] {
		case 58:
			goto tr237
		case 60:
			goto tr236
		case 61:
			goto tr232
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] > 13:
				if 32 <= data[p] && data[p] <= 47 {
					goto tr235
				}
			case data[p] >= 9:
				goto tr235
			}
		case data[p] > 64:
			switch {
			case data[p] > 96:
				if 123 <= data[p] && data[p] <= 126 {
					goto tr235
				}
			case data[p] >= 91:
				goto tr235
			}
		default:
			goto tr235
		}
		goto st100
tr237:
//line xss.rl:10

            return true
        
	goto st231
	st231:
		if p++; p == pe {
			goto _test_eof231
		}
	st_case_231:
//line xss.go:4095
		switch data[p] {
		case 59:
			goto tr234
		case 60:
			goto tr233
		case 117:
			goto st125
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] > 13:
				if 32 <= data[p] && data[p] <= 47 {
					goto tr232
				}
			case data[p] >= 9:
				goto tr232
			}
		case data[p] > 64:
			switch {
			case data[p] > 96:
				if 123 <= data[p] && data[p] <= 126 {
					goto tr232
				}
			case data[p] >= 91:
				goto tr232
			}
		default:
			goto tr232
		}
		goto st98
tr233:
//line xss.rl:10

            return true
        
	goto st232
	st232:
		if p++; p == pe {
			goto _test_eof232
		}
	st_case_232:
//line xss.go:4138
		switch data[p] {
		case 59:
			goto tr234
		case 60:
			goto tr233
		case 115:
			goto st103
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] > 13:
				if 32 <= data[p] && data[p] <= 47 {
					goto tr232
				}
			case data[p] >= 9:
				goto tr232
			}
		case data[p] > 64:
			switch {
			case data[p] > 96:
				if 123 <= data[p] && data[p] <= 126 {
					goto tr232
				}
			case data[p] >= 91:
				goto tr232
			}
		default:
			goto tr232
		}
		goto st98
tr236:
//line xss.rl:10

            return true
        
	goto st233
	st233:
		if p++; p == pe {
			goto _test_eof233
		}
	st_case_233:
//line xss.go:4181
		switch data[p] {
		case 58:
			goto tr237
		case 60:
			goto tr236
		case 61:
			goto tr232
		case 115:
			goto st139
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] > 13:
				if 32 <= data[p] && data[p] <= 47 {
					goto tr235
				}
			case data[p] >= 9:
				goto tr235
			}
		case data[p] > 64:
			switch {
			case data[p] > 96:
				if 123 <= data[p] && data[p] <= 126 {
					goto tr235
				}
			case data[p] >= 91:
				goto tr235
			}
		default:
			goto tr235
		}
		goto st100
tr230:
//line xss.rl:10

            return true
        
	goto st234
	st234:
		if p++; p == pe {
			goto _test_eof234
		}
	st_case_234:
//line xss.go:4226
		switch data[p] {
		case 58:
			goto tr231
		case 60:
			goto tr230
		case 61:
			goto tr218
		case 115:
			goto st145
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] > 13:
				if 32 <= data[p] && data[p] <= 47 {
					goto tr229
				}
			case data[p] >= 9:
				goto tr229
			}
		case data[p] > 64:
			switch {
			case data[p] > 96:
				if 123 <= data[p] && data[p] <= 126 {
					goto tr229
				}
			case data[p] >= 91:
				goto tr229
			}
		default:
			goto tr229
		}
		goto st96
	st165:
		if p++; p == pe {
			goto _test_eof165
		}
	st_case_165:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st95
		case 97:
			goto st166
		}
		if 98 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st166:
		if p++; p == pe {
			goto _test_eof166
		}
	st_case_166:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st95
		case 116:
			goto st167
		}
		if 97 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st167:
		if p++; p == pe {
			goto _test_eof167
		}
	st_case_167:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st95
		case 116:
			goto st168
		}
		if 97 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st168:
		if p++; p == pe {
			goto _test_eof168
		}
	st_case_168:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st95
		case 101:
			goto st169
		}
		if 97 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st169:
		if p++; p == pe {
			goto _test_eof169
		}
	st_case_169:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st95
		case 114:
			goto st170
		}
		if 97 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st170:
		if p++; p == pe {
			goto _test_eof170
		}
	st_case_170:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st95
		case 110:
			goto st171
		}
		if 97 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st171:
		if p++; p == pe {
			goto _test_eof171
		}
	st_case_171:
		switch data[p] {
		case 60:
			goto st49
		case 61:
			goto st235
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto st47
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st48
				}
			default:
				goto st47
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st48
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st47
					}
				case data[p] >= 97:
					goto st196
				}
			default:
				goto st47
			}
		default:
			goto st47
		}
		goto st1
	st235:
		if p++; p == pe {
			goto _test_eof235
		}
	st_case_235:
		switch data[p] {
		case 58:
			goto tr221
		case 60:
			goto tr240
		case 61:
			goto tr224
		case 95:
			goto tr241
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr238
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st253
				}
			default:
				goto tr238
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st253
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr238
					}
				case data[p] >= 97:
					goto st253
				}
			default:
				goto tr238
			}
		default:
			goto tr238
		}
		goto st96
tr238:
//line xss.rl:10

            return true
        
	goto st236
	st236:
		if p++; p == pe {
			goto _test_eof236
		}
	st_case_236:
//line xss.go:4471
		switch data[p] {
		case 58:
			goto tr242
		case 60:
			goto tr240
		case 61:
			goto tr224
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr238
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st172
				}
			default:
				goto tr238
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st172
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr238
					}
				case data[p] >= 97:
					goto st172
				}
			default:
				goto tr238
			}
		default:
			goto tr238
		}
		goto st96
	st172:
		if p++; p == pe {
			goto _test_eof172
		}
	st_case_172:
		switch data[p] {
		case 58:
			goto st97
		case 60:
			goto st144
		case 61:
			goto st212
		case 95:
			goto st172
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st172
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st172
			}
		default:
			goto st172
		}
		goto st96
tr242:
//line xss.rl:10

            return true
        
	goto st237
	st237:
		if p++; p == pe {
			goto _test_eof237
		}
	st_case_237:
//line xss.go:4555
		switch data[p] {
		case 60:
			goto tr244
		case 61:
			goto tr245
		case 117:
			goto st181
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr243
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st173
				}
			default:
				goto tr243
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st173
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr243
					}
				case data[p] >= 97:
					goto st173
				}
			default:
				goto tr243
			}
		default:
			goto tr243
		}
		goto st98
tr243:
//line xss.rl:10

            return true
        
	goto st238
	st238:
		if p++; p == pe {
			goto _test_eof238
		}
	st_case_238:
//line xss.go:4611
		switch data[p] {
		case 59:
			goto tr247
		case 60:
			goto tr244
		case 61:
			goto tr245
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr243
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st173
				}
			default:
				goto tr243
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st173
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr243
					}
				case data[p] >= 97:
					goto st173
				}
			default:
				goto tr243
			}
		default:
			goto tr243
		}
		goto st98
	st173:
		if p++; p == pe {
			goto _test_eof173
		}
	st_case_173:
		switch data[p] {
		case 59:
			goto st99
		case 60:
			goto st102
		case 61:
			goto st239
		case 95:
			goto st173
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st173
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st173
			}
		default:
			goto st173
		}
		goto st98
tr248:
//line xss.rl:10

            return true
        
	goto st239
	st239:
		if p++; p == pe {
			goto _test_eof239
		}
	st_case_239:
//line xss.go:4695
		switch data[p] {
		case 59:
			goto tr234
		case 60:
			goto tr233
		case 95:
			goto tr248
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr232
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st239
				}
			default:
				goto tr232
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st239
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr232
					}
				case data[p] >= 97:
					goto st239
				}
			default:
				goto tr232
			}
		default:
			goto tr232
		}
		goto st98
tr247:
//line xss.rl:10

            return true
        
	goto st240
	st240:
		if p++; p == pe {
			goto _test_eof240
		}
	st_case_240:
//line xss.go:4751
		switch data[p] {
		case 58:
			goto tr243
		case 60:
			goto tr250
		case 61:
			goto tr245
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr249
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st174
				}
			default:
				goto tr249
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st174
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr249
					}
				case data[p] >= 97:
					goto st174
				}
			default:
				goto tr249
			}
		default:
			goto tr249
		}
		goto st100
tr249:
//line xss.rl:10

            return true
        
	goto st241
	st241:
		if p++; p == pe {
			goto _test_eof241
		}
	st_case_241:
//line xss.go:4807
		switch data[p] {
		case 58:
			goto tr251
		case 60:
			goto tr250
		case 61:
			goto tr245
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr249
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st174
				}
			default:
				goto tr249
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st174
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr249
					}
				case data[p] >= 97:
					goto st174
				}
			default:
				goto tr249
			}
		default:
			goto tr249
		}
		goto st100
	st174:
		if p++; p == pe {
			goto _test_eof174
		}
	st_case_174:
		switch data[p] {
		case 58:
			goto st101
		case 60:
			goto st138
		case 61:
			goto st239
		case 95:
			goto st174
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st174
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st174
			}
		default:
			goto st174
		}
		goto st100
tr251:
//line xss.rl:10

            return true
        
	goto st242
	st242:
		if p++; p == pe {
			goto _test_eof242
		}
	st_case_242:
//line xss.go:4891
		switch data[p] {
		case 59:
			goto tr247
		case 60:
			goto tr244
		case 61:
			goto tr245
		case 117:
			goto st181
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr243
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st173
				}
			default:
				goto tr243
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st173
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr243
					}
				case data[p] >= 97:
					goto st173
				}
			default:
				goto tr243
			}
		default:
			goto tr243
		}
		goto st98
tr244:
//line xss.rl:10

            return true
        
	goto st243
	st243:
		if p++; p == pe {
			goto _test_eof243
		}
	st_case_243:
//line xss.go:4949
		switch data[p] {
		case 59:
			goto tr247
		case 60:
			goto tr244
		case 61:
			goto tr245
		case 115:
			goto st175
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr243
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st173
				}
			default:
				goto tr243
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st173
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr243
					}
				case data[p] >= 97:
					goto st173
				}
			default:
				goto tr243
			}
		default:
			goto tr243
		}
		goto st98
tr245:
//line xss.rl:10

            return true
        
	goto st244
	st244:
		if p++; p == pe {
			goto _test_eof244
		}
	st_case_244:
//line xss.go:5007
		switch data[p] {
		case 59:
			goto tr247
		case 60:
			goto tr244
		case 61:
			goto tr245
		case 95:
			goto tr245
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr243
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st245
				}
			default:
				goto tr243
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st245
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr243
					}
				case data[p] >= 97:
					goto st245
				}
			default:
				goto tr243
			}
		default:
			goto tr243
		}
		goto st98
tr254:
//line xss.rl:10

            return true
        
	goto st245
	st245:
		if p++; p == pe {
			goto _test_eof245
		}
	st_case_245:
//line xss.go:5065
		switch data[p] {
		case 59:
			goto tr234
		case 60:
			goto tr233
		case 61:
			goto tr248
		case 95:
			goto tr254
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr232
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st245
				}
			default:
				goto tr232
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st245
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr232
					}
				case data[p] >= 97:
					goto st245
				}
			default:
				goto tr232
			}
		default:
			goto tr232
		}
		goto st98
	st175:
		if p++; p == pe {
			goto _test_eof175
		}
	st_case_175:
		switch data[p] {
		case 59:
			goto st99
		case 60:
			goto st102
		case 61:
			goto st239
		case 95:
			goto st173
		case 99:
			goto st176
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st173
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st173
			}
		default:
			goto st173
		}
		goto st98
	st176:
		if p++; p == pe {
			goto _test_eof176
		}
	st_case_176:
		switch data[p] {
		case 59:
			goto st99
		case 60:
			goto st102
		case 61:
			goto st239
		case 95:
			goto st173
		case 114:
			goto st177
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st173
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st173
			}
		default:
			goto st173
		}
		goto st98
	st177:
		if p++; p == pe {
			goto _test_eof177
		}
	st_case_177:
		switch data[p] {
		case 59:
			goto st99
		case 60:
			goto st102
		case 61:
			goto st239
		case 95:
			goto st173
		case 105:
			goto st178
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st173
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st173
			}
		default:
			goto st173
		}
		goto st98
	st178:
		if p++; p == pe {
			goto _test_eof178
		}
	st_case_178:
		switch data[p] {
		case 59:
			goto st99
		case 60:
			goto st102
		case 61:
			goto st239
		case 95:
			goto st173
		case 112:
			goto st179
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st173
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st173
			}
		default:
			goto st173
		}
		goto st98
	st179:
		if p++; p == pe {
			goto _test_eof179
		}
	st_case_179:
		switch data[p] {
		case 59:
			goto st99
		case 60:
			goto st102
		case 61:
			goto st239
		case 95:
			goto st173
		case 116:
			goto st180
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st173
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st173
			}
		default:
			goto st173
		}
		goto st98
	st180:
		if p++; p == pe {
			goto _test_eof180
		}
	st_case_180:
		switch data[p] {
		case 59:
			goto st109
		case 61:
			goto st246
		case 62:
			goto tr121
		case 95:
			goto st180
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st180
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st180
			}
		default:
			goto st180
		}
		goto st108
tr257:
//line xss.rl:10

            return true
        
	goto st246
	st246:
		if p++; p == pe {
			goto _test_eof246
		}
	st_case_246:
//line xss.go:5301
		switch data[p] {
		case 59:
			goto tr256
		case 95:
			goto tr257
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr255
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st246
				}
			default:
				goto tr255
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st246
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr255
					}
				case data[p] >= 97:
					goto st246
				}
			default:
				goto tr255
			}
		default:
			goto tr255
		}
		goto st108
tr255:
//line xss.rl:10

            return true
        
	goto st247
	st247:
		if p++; p == pe {
			goto _test_eof247
		}
	st_case_247:
//line xss.go:5355
		if data[p] == 59 {
			goto tr256
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] > 13:
				if 32 <= data[p] && data[p] <= 47 {
					goto tr255
				}
			case data[p] >= 9:
				goto tr255
			}
		case data[p] > 64:
			switch {
			case data[p] > 96:
				if 123 <= data[p] && data[p] <= 126 {
					goto tr255
				}
			case data[p] >= 91:
				goto tr255
			}
		default:
			goto tr255
		}
		goto st108
tr256:
//line xss.rl:10

            return true
        
	goto st248
	st248:
		if p++; p == pe {
			goto _test_eof248
		}
	st_case_248:
//line xss.go:5393
		switch data[p] {
		case 58:
			goto tr255
		case 61:
			goto tr255
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] > 13:
				if 32 <= data[p] && data[p] <= 47 {
					goto tr258
				}
			case data[p] >= 9:
				goto tr258
			}
		case data[p] > 64:
			switch {
			case data[p] > 96:
				if 123 <= data[p] && data[p] <= 126 {
					goto tr258
				}
			case data[p] >= 91:
				goto tr258
			}
		default:
			goto tr258
		}
		goto st110
tr258:
//line xss.rl:10

            return true
        
	goto st249
	st249:
		if p++; p == pe {
			goto _test_eof249
		}
	st_case_249:
//line xss.go:5434
		switch data[p] {
		case 58:
			goto tr259
		case 61:
			goto tr255
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] > 13:
				if 32 <= data[p] && data[p] <= 47 {
					goto tr258
				}
			case data[p] >= 9:
				goto tr258
			}
		case data[p] > 64:
			switch {
			case data[p] > 96:
				if 123 <= data[p] && data[p] <= 126 {
					goto tr258
				}
			case data[p] >= 91:
				goto tr258
			}
		default:
			goto tr258
		}
		goto st110
tr259:
//line xss.rl:10

            return true
        
	goto st250
	st250:
		if p++; p == pe {
			goto _test_eof250
		}
	st_case_250:
//line xss.go:5475
		switch data[p] {
		case 59:
			goto tr256
		case 117:
			goto st112
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] > 13:
				if 32 <= data[p] && data[p] <= 47 {
					goto tr255
				}
			case data[p] >= 9:
				goto tr255
			}
		case data[p] > 64:
			switch {
			case data[p] > 96:
				if 123 <= data[p] && data[p] <= 126 {
					goto tr255
				}
			case data[p] >= 91:
				goto tr255
			}
		default:
			goto tr255
		}
		goto st108
	st181:
		if p++; p == pe {
			goto _test_eof181
		}
	st_case_181:
		switch data[p] {
		case 59:
			goto st99
		case 60:
			goto st102
		case 61:
			goto st239
		case 95:
			goto st173
		case 114:
			goto st182
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st173
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st173
			}
		default:
			goto st173
		}
		goto st98
	st182:
		if p++; p == pe {
			goto _test_eof182
		}
	st_case_182:
		switch data[p] {
		case 59:
			goto st99
		case 60:
			goto st102
		case 61:
			goto st239
		case 95:
			goto st173
		case 108:
			goto st183
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st173
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st173
			}
		default:
			goto st173
		}
		goto st98
	st183:
		if p++; p == pe {
			goto _test_eof183
		}
	st_case_183:
		switch data[p] {
		case 40:
			goto st128
		case 59:
			goto st99
		case 60:
			goto st102
		case 61:
			goto st239
		case 95:
			goto st173
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st173
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st173
			}
		default:
			goto st173
		}
		goto st98
tr250:
//line xss.rl:10

            return true
        
	goto st251
	st251:
		if p++; p == pe {
			goto _test_eof251
		}
	st_case_251:
//line xss.go:5606
		switch data[p] {
		case 58:
			goto tr251
		case 60:
			goto tr250
		case 61:
			goto tr245
		case 115:
			goto st184
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr249
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st174
				}
			default:
				goto tr249
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st174
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr249
					}
				case data[p] >= 97:
					goto st174
				}
			default:
				goto tr249
			}
		default:
			goto tr249
		}
		goto st100
	st184:
		if p++; p == pe {
			goto _test_eof184
		}
	st_case_184:
		switch data[p] {
		case 58:
			goto st101
		case 60:
			goto st138
		case 61:
			goto st239
		case 95:
			goto st174
		case 99:
			goto st185
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st174
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st174
			}
		default:
			goto st174
		}
		goto st100
	st185:
		if p++; p == pe {
			goto _test_eof185
		}
	st_case_185:
		switch data[p] {
		case 58:
			goto st101
		case 60:
			goto st138
		case 61:
			goto st239
		case 95:
			goto st174
		case 114:
			goto st186
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st174
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st174
			}
		default:
			goto st174
		}
		goto st100
	st186:
		if p++; p == pe {
			goto _test_eof186
		}
	st_case_186:
		switch data[p] {
		case 58:
			goto st101
		case 60:
			goto st138
		case 61:
			goto st239
		case 95:
			goto st174
		case 105:
			goto st187
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st174
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st174
			}
		default:
			goto st174
		}
		goto st100
	st187:
		if p++; p == pe {
			goto _test_eof187
		}
	st_case_187:
		switch data[p] {
		case 58:
			goto st101
		case 60:
			goto st138
		case 61:
			goto st239
		case 95:
			goto st174
		case 112:
			goto st188
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st174
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st174
			}
		default:
			goto st174
		}
		goto st100
	st188:
		if p++; p == pe {
			goto _test_eof188
		}
	st_case_188:
		switch data[p] {
		case 58:
			goto st101
		case 60:
			goto st138
		case 61:
			goto st239
		case 95:
			goto st174
		case 116:
			goto st189
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st174
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st174
			}
		default:
			goto st174
		}
		goto st100
	st189:
		if p++; p == pe {
			goto _test_eof189
		}
	st_case_189:
		switch data[p] {
		case 58:
			goto st111
		case 61:
			goto st246
		case 62:
			goto tr123
		case 95:
			goto st189
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st189
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st189
			}
		default:
			goto st189
		}
		goto st110
tr240:
//line xss.rl:10

            return true
        
	goto st252
	st252:
		if p++; p == pe {
			goto _test_eof252
		}
	st_case_252:
//line xss.go:5842
		switch data[p] {
		case 58:
			goto tr242
		case 60:
			goto tr240
		case 61:
			goto tr224
		case 115:
			goto st190
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr238
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st172
				}
			default:
				goto tr238
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st172
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr238
					}
				case data[p] >= 97:
					goto st172
				}
			default:
				goto tr238
			}
		default:
			goto tr238
		}
		goto st96
	st190:
		if p++; p == pe {
			goto _test_eof190
		}
	st_case_190:
		switch data[p] {
		case 58:
			goto st97
		case 60:
			goto st144
		case 61:
			goto st212
		case 95:
			goto st172
		case 99:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st172
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st172
			}
		default:
			goto st172
		}
		goto st96
	st191:
		if p++; p == pe {
			goto _test_eof191
		}
	st_case_191:
		switch data[p] {
		case 58:
			goto st97
		case 60:
			goto st144
		case 61:
			goto st212
		case 95:
			goto st172
		case 114:
			goto st192
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st172
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st172
			}
		default:
			goto st172
		}
		goto st96
	st192:
		if p++; p == pe {
			goto _test_eof192
		}
	st_case_192:
		switch data[p] {
		case 58:
			goto st97
		case 60:
			goto st144
		case 61:
			goto st212
		case 95:
			goto st172
		case 105:
			goto st193
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st172
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st172
			}
		default:
			goto st172
		}
		goto st96
	st193:
		if p++; p == pe {
			goto _test_eof193
		}
	st_case_193:
		switch data[p] {
		case 58:
			goto st97
		case 60:
			goto st144
		case 61:
			goto st212
		case 95:
			goto st172
		case 112:
			goto st194
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st172
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st172
			}
		default:
			goto st172
		}
		goto st96
	st194:
		if p++; p == pe {
			goto _test_eof194
		}
	st_case_194:
		switch data[p] {
		case 58:
			goto st97
		case 60:
			goto st144
		case 61:
			goto st212
		case 95:
			goto st172
		case 116:
			goto st195
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st172
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st172
			}
		default:
			goto st172
		}
		goto st96
	st195:
		if p++; p == pe {
			goto _test_eof195
		}
	st_case_195:
		switch data[p] {
		case 58:
			goto st151
		case 61:
			goto st216
		case 62:
			goto tr164
		case 95:
			goto st195
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st195
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st195
			}
		default:
			goto st195
		}
		goto st150
tr262:
//line xss.rl:10

            return true
        
	goto st253
	st253:
		if p++; p == pe {
			goto _test_eof253
		}
	st_case_253:
//line xss.go:6078
		switch data[p] {
		case 58:
			goto tr231
		case 60:
			goto tr230
		case 61:
			goto tr220
		case 95:
			goto tr262
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr229
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st253
				}
			default:
				goto tr229
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st253
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr229
					}
				case data[p] >= 97:
					goto st253
				}
			default:
				goto tr229
			}
		default:
			goto tr229
		}
		goto st96
tr241:
//line xss.rl:10

            return true
        
	goto st254
	st254:
		if p++; p == pe {
			goto _test_eof254
		}
	st_case_254:
//line xss.go:6136
		switch data[p] {
		case 58:
			goto tr242
		case 60:
			goto tr240
		case 61:
			goto tr224
		case 95:
			goto tr241
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr238
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st253
				}
			default:
				goto tr238
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st253
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr238
					}
				case data[p] >= 97:
					goto st253
				}
			default:
				goto tr238
			}
		default:
			goto tr238
		}
		goto st96
	st196:
		if p++; p == pe {
			goto _test_eof196
		}
	st_case_196:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st255
		case 95:
			goto st48
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st48
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st196
			}
		default:
			goto st48
		}
		goto st1
	st255:
		if p++; p == pe {
			goto _test_eof255
		}
	st_case_255:
		switch data[p] {
		case 59:
			goto tr229
		case 60:
			goto tr230
		case 95:
			goto tr264
		}
		switch {
		case data[p] < 62:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr229
				}
			case data[p] > 47:
				switch {
				case data[p] > 57:
					if 58 <= data[p] && data[p] <= 61 {
						goto tr218
					}
				case data[p] >= 48:
					goto st256
				}
			default:
				goto tr229
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st256
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr229
					}
				case data[p] >= 97:
					goto st256
				}
			default:
				goto tr229
			}
		default:
			goto tr229
		}
		goto st96
tr264:
//line xss.rl:10

            return true
        
	goto st256
	st256:
		if p++; p == pe {
			goto _test_eof256
		}
	st_case_256:
//line xss.go:6274
		switch data[p] {
		case 58:
			goto tr231
		case 60:
			goto tr230
		case 61:
			goto tr218
		case 95:
			goto tr264
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr229
				}
			case data[p] > 47:
				if 48 <= data[p] && data[p] <= 57 {
					goto st256
				}
			default:
				goto tr229
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st256
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto tr229
					}
				case data[p] >= 97:
					goto st256
				}
			default:
				goto tr229
			}
		default:
			goto tr229
		}
		goto st96
	st197:
		if p++; p == pe {
			goto _test_eof197
		}
	st_case_197:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st95
		case 104:
			goto st198
		case 108:
			goto st201
		case 109:
			goto st205
		}
		if 97 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st198:
		if p++; p == pe {
			goto _test_eof198
		}
	st_case_198:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st95
		case 116:
			goto st199
		}
		if 97 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st199:
		if p++; p == pe {
			goto _test_eof199
		}
	st_case_199:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st95
		case 109:
			goto st200
		}
		if 97 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st200:
		if p++; p == pe {
			goto _test_eof200
		}
	st_case_200:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st95
		case 108:
			goto st224
		}
		if 97 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st201:
		if p++; p == pe {
			goto _test_eof201
		}
	st_case_201:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st95
		case 105:
			goto st202
		}
		if 97 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st202:
		if p++; p == pe {
			goto _test_eof202
		}
	st_case_202:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st95
		case 110:
			goto st203
		}
		if 97 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st203:
		if p++; p == pe {
			goto _test_eof203
		}
	st_case_203:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st95
		case 107:
			goto st204
		}
		if 97 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st204:
		if p++; p == pe {
			goto _test_eof204
		}
	st_case_204:
		switch data[p] {
		case 58:
			goto st61
		case 60:
			goto st2
		case 61:
			goto st95
		}
		if 97 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st205:
		if p++; p == pe {
			goto _test_eof205
		}
	st_case_205:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st95
		case 108:
			goto st206
		}
		if 97 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st206:
		if p++; p == pe {
			goto _test_eof206
		}
	st_case_206:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st95
		case 110:
			goto st207
		}
		if 97 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st207:
		if p++; p == pe {
			goto _test_eof207
		}
	st_case_207:
		switch data[p] {
		case 60:
			goto st2
		case 61:
			goto st95
		case 115:
			goto st224
		}
		if 97 <= data[p] && data[p] <= 122 {
			goto st94
		}
		goto st1
	st_out:
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof208: cs = 208; goto _test_eof
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
	_test_eof209: cs = 209; goto _test_eof
	_test_eof210: cs = 210; goto _test_eof
	_test_eof211: cs = 211; goto _test_eof
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
	_test_eof48: cs = 48; goto _test_eof
	_test_eof212: cs = 212; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof213: cs = 213; goto _test_eof
	_test_eof214: cs = 214; goto _test_eof
	_test_eof215: cs = 215; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof
	_test_eof53: cs = 53; goto _test_eof
	_test_eof54: cs = 54; goto _test_eof
	_test_eof55: cs = 55; goto _test_eof
	_test_eof216: cs = 216; goto _test_eof
	_test_eof217: cs = 217; goto _test_eof
	_test_eof218: cs = 218; goto _test_eof
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
	_test_eof219: cs = 219; goto _test_eof
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
	_test_eof220: cs = 220; goto _test_eof
	_test_eof221: cs = 221; goto _test_eof
	_test_eof125: cs = 125; goto _test_eof
	_test_eof126: cs = 126; goto _test_eof
	_test_eof127: cs = 127; goto _test_eof
	_test_eof128: cs = 128; goto _test_eof
	_test_eof129: cs = 129; goto _test_eof
	_test_eof130: cs = 130; goto _test_eof
	_test_eof131: cs = 131; goto _test_eof
	_test_eof132: cs = 132; goto _test_eof
	_test_eof133: cs = 133; goto _test_eof
	_test_eof134: cs = 134; goto _test_eof
	_test_eof135: cs = 135; goto _test_eof
	_test_eof136: cs = 136; goto _test_eof
	_test_eof137: cs = 137; goto _test_eof
	_test_eof222: cs = 222; goto _test_eof
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
	_test_eof223: cs = 223; goto _test_eof
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
	_test_eof224: cs = 224; goto _test_eof
	_test_eof225: cs = 225; goto _test_eof
	_test_eof226: cs = 226; goto _test_eof
	_test_eof227: cs = 227; goto _test_eof
	_test_eof228: cs = 228; goto _test_eof
	_test_eof229: cs = 229; goto _test_eof
	_test_eof230: cs = 230; goto _test_eof
	_test_eof231: cs = 231; goto _test_eof
	_test_eof232: cs = 232; goto _test_eof
	_test_eof233: cs = 233; goto _test_eof
	_test_eof234: cs = 234; goto _test_eof
	_test_eof165: cs = 165; goto _test_eof
	_test_eof166: cs = 166; goto _test_eof
	_test_eof167: cs = 167; goto _test_eof
	_test_eof168: cs = 168; goto _test_eof
	_test_eof169: cs = 169; goto _test_eof
	_test_eof170: cs = 170; goto _test_eof
	_test_eof171: cs = 171; goto _test_eof
	_test_eof235: cs = 235; goto _test_eof
	_test_eof236: cs = 236; goto _test_eof
	_test_eof172: cs = 172; goto _test_eof
	_test_eof237: cs = 237; goto _test_eof
	_test_eof238: cs = 238; goto _test_eof
	_test_eof173: cs = 173; goto _test_eof
	_test_eof239: cs = 239; goto _test_eof
	_test_eof240: cs = 240; goto _test_eof
	_test_eof241: cs = 241; goto _test_eof
	_test_eof174: cs = 174; goto _test_eof
	_test_eof242: cs = 242; goto _test_eof
	_test_eof243: cs = 243; goto _test_eof
	_test_eof244: cs = 244; goto _test_eof
	_test_eof245: cs = 245; goto _test_eof
	_test_eof175: cs = 175; goto _test_eof
	_test_eof176: cs = 176; goto _test_eof
	_test_eof177: cs = 177; goto _test_eof
	_test_eof178: cs = 178; goto _test_eof
	_test_eof179: cs = 179; goto _test_eof
	_test_eof180: cs = 180; goto _test_eof
	_test_eof246: cs = 246; goto _test_eof
	_test_eof247: cs = 247; goto _test_eof
	_test_eof248: cs = 248; goto _test_eof
	_test_eof249: cs = 249; goto _test_eof
	_test_eof250: cs = 250; goto _test_eof
	_test_eof181: cs = 181; goto _test_eof
	_test_eof182: cs = 182; goto _test_eof
	_test_eof183: cs = 183; goto _test_eof
	_test_eof251: cs = 251; goto _test_eof
	_test_eof184: cs = 184; goto _test_eof
	_test_eof185: cs = 185; goto _test_eof
	_test_eof186: cs = 186; goto _test_eof
	_test_eof187: cs = 187; goto _test_eof
	_test_eof188: cs = 188; goto _test_eof
	_test_eof189: cs = 189; goto _test_eof
	_test_eof252: cs = 252; goto _test_eof
	_test_eof190: cs = 190; goto _test_eof
	_test_eof191: cs = 191; goto _test_eof
	_test_eof192: cs = 192; goto _test_eof
	_test_eof193: cs = 193; goto _test_eof
	_test_eof194: cs = 194; goto _test_eof
	_test_eof195: cs = 195; goto _test_eof
	_test_eof253: cs = 253; goto _test_eof
	_test_eof254: cs = 254; goto _test_eof
	_test_eof196: cs = 196; goto _test_eof
	_test_eof255: cs = 255; goto _test_eof
	_test_eof256: cs = 256; goto _test_eof
	_test_eof197: cs = 197; goto _test_eof
	_test_eof198: cs = 198; goto _test_eof
	_test_eof199: cs = 199; goto _test_eof
	_test_eof200: cs = 200; goto _test_eof
	_test_eof201: cs = 201; goto _test_eof
	_test_eof202: cs = 202; goto _test_eof
	_test_eof203: cs = 203; goto _test_eof
	_test_eof204: cs = 204; goto _test_eof
	_test_eof205: cs = 205; goto _test_eof
	_test_eof206: cs = 206; goto _test_eof
	_test_eof207: cs = 207; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 209, 212, 213, 216, 218, 220, 222, 224, 235, 239, 244, 245, 246, 253, 254, 255, 256:
//line xss.rl:10

            return true
        
//line xss.go:6778
		}
	}

	}

//line xss.rl:92

        return false
}
