
//line xss_170.rl:1
package rule_941


func matchXSS170(data []byte) bool {

//line xss_170.rl:6

//line xss_170.go:11
const xss_start int = 0
const xss_first_final int = 286
const xss_error int = -1

const xss_en_main int = 0


//line xss_170.rl:7
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof


    
//line xss_170.go:25
	{
	cs = xss_start
	}

//line xss_170.go:30
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
	case 286:
		goto st_case_286
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
	case 287:
		goto st_case_287
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
	case 288:
		goto st_case_288
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
	case 67:
		goto st_case_67
	case 68:
		goto st_case_68
	case 69:
		goto st_case_69
	case 289:
		goto st_case_289
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
	case 290:
		goto st_case_290
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
	case 291:
		goto st_case_291
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
	case 292:
		goto st_case_292
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
	case 293:
		goto st_case_293
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
	case 294:
		goto st_case_294
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
	case 295:
		goto st_case_295
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
	case 296:
		goto st_case_296
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
	case 297:
		goto st_case_297
	case 298:
		goto st_case_298
	case 179:
		goto st_case_179
	case 180:
		goto st_case_180
	case 181:
		goto st_case_181
	case 182:
		goto st_case_182
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
	case 189:
		goto st_case_189
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
	case 196:
		goto st_case_196
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
	case 299:
		goto st_case_299
	case 205:
		goto st_case_205
	case 206:
		goto st_case_206
	case 207:
		goto st_case_207
	case 208:
		goto st_case_208
	case 209:
		goto st_case_209
	case 210:
		goto st_case_210
	case 211:
		goto st_case_211
	case 212:
		goto st_case_212
	case 213:
		goto st_case_213
	case 214:
		goto st_case_214
	case 215:
		goto st_case_215
	case 216:
		goto st_case_216
	case 217:
		goto st_case_217
	case 218:
		goto st_case_218
	case 219:
		goto st_case_219
	case 220:
		goto st_case_220
	case 221:
		goto st_case_221
	case 222:
		goto st_case_222
	case 223:
		goto st_case_223
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
	case 300:
		goto st_case_300
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
	case 235:
		goto st_case_235
	case 236:
		goto st_case_236
	case 237:
		goto st_case_237
	case 238:
		goto st_case_238
	case 239:
		goto st_case_239
	case 240:
		goto st_case_240
	case 301:
		goto st_case_301
	case 241:
		goto st_case_241
	case 242:
		goto st_case_242
	case 243:
		goto st_case_243
	case 244:
		goto st_case_244
	case 245:
		goto st_case_245
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
	case 251:
		goto st_case_251
	case 252:
		goto st_case_252
	case 253:
		goto st_case_253
	case 302:
		goto st_case_302
	case 254:
		goto st_case_254
	case 255:
		goto st_case_255
	case 256:
		goto st_case_256
	case 257:
		goto st_case_257
	case 258:
		goto st_case_258
	case 259:
		goto st_case_259
	case 260:
		goto st_case_260
	case 261:
		goto st_case_261
	case 262:
		goto st_case_262
	case 263:
		goto st_case_263
	case 264:
		goto st_case_264
	case 265:
		goto st_case_265
	case 266:
		goto st_case_266
	case 267:
		goto st_case_267
	case 268:
		goto st_case_268
	case 269:
		goto st_case_269
	case 270:
		goto st_case_270
	case 271:
		goto st_case_271
	case 272:
		goto st_case_272
	case 273:
		goto st_case_273
	case 274:
		goto st_case_274
	case 275:
		goto st_case_275
	case 276:
		goto st_case_276
	case 277:
		goto st_case_277
	case 278:
		goto st_case_278
	case 279:
		goto st_case_279
	case 280:
		goto st_case_280
	case 281:
		goto st_case_281
	case 282:
		goto st_case_282
	case 283:
		goto st_case_283
	case 284:
		goto st_case_284
	case 285:
		goto st_case_285
	}
	goto st_out
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		switch data[p] {
		case 95:
			goto st1
		case 100:
			goto st2
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
		goto st0
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
		goto st0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 95:
			goto st1
		case 97:
			goto st3
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
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 95:
			goto st1
		case 116:
			goto st4
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
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 95:
			goto st1
		case 97:
			goto st5
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
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 58:
			goto st6
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
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch data[p] {
		case 44:
			goto st178
		case 59:
			goto st212
		case 95:
			goto st269
		case 100:
			goto st282
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st196
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st269
			}
		default:
			goto st196
		}
		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st197
		case 95:
			goto st196
		case 100:
			goto st265
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st196
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st196
			}
		default:
			goto st196
		}
		goto st7
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		switch data[p] {
		case 59:
			goto st10
		case 60:
			goto st147
		case 95:
			goto st9
		case 100:
			goto st173
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st9
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st9
			}
		default:
			goto st9
		}
		goto st8
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		switch data[p] {
		case 59:
			goto st10
		case 60:
			goto st147
		case 95:
			goto st9
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st9
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st9
			}
		default:
			goto st9
		}
		goto st8
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 98:
			goto st81
		case 99:
			goto st87
		case 100:
			goto st95
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto st10
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto st10
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch data[p] {
		case 95:
			goto st13
		case 98:
			goto st16
		case 99:
			goto st22
		case 100:
			goto st29
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		goto st12
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st13
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st13
		case 98:
			goto st16
		case 99:
			goto st22
		case 100:
			goto st29
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		goto st14
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		switch data[p] {
		case 62:
			goto tr32
		case 95:
			goto st13
		case 98:
			goto st16
		case 99:
			goto st22
		case 100:
			goto st29
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		goto tr31
tr31:
//line xss_170.rl:13

            return true
        
	goto st286
	st286:
		if p++; p == pe {
			goto _test_eof286
		}
	st_case_286:
//line xss_170.go:1072
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st13
		case 98:
			goto st16
		case 99:
			goto st22
		case 100:
			goto st29
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		goto st14
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st13
		case 97:
			goto st17
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		goto st14
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st13
		case 115:
			goto st18
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		goto st14
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st13
		case 101:
			goto st19
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		goto st14
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		switch data[p] {
		case 54:
			goto st20
		case 62:
			goto st15
		case 95:
			goto st13
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		goto st14
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		switch data[p] {
		case 52:
			goto st21
		case 62:
			goto st15
		case 95:
			goto st13
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		goto st14
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		switch data[p] {
		case 62:
			goto tr32
		case 95:
			goto st13
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		goto tr31
tr32:
//line xss_170.rl:13

            return true
        
	goto st287
	st287:
		if p++; p == pe {
			goto _test_eof287
		}
	st_case_287:
//line xss_170.go:1263
		switch data[p] {
		case 62:
			goto tr32
		case 95:
			goto st13
		case 98:
			goto st16
		case 99:
			goto st22
		case 100:
			goto st29
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		goto tr31
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st13
		case 104:
			goto st23
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		goto st14
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st13
		case 97:
			goto st24
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		goto st14
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st13
		case 114:
			goto st25
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		goto st14
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st13
		case 115:
			goto st26
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		goto st14
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st13
		case 101:
			goto st27
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		goto st14
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st13
		case 116:
			goto st28
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		goto st14
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		if data[p] == 95 {
			goto st13
		}
		switch {
		case data[p] < 61:
			if 48 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 62:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st13
				}
			case data[p] >= 65:
				goto st13
			}
		default:
			goto st15
		}
		goto st14
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st13
		case 97:
			goto st30
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		goto st14
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st13
		case 116:
			goto st31
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		goto st14
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st13
		case 97:
			goto st32
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		goto st14
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		switch data[p] {
		case 58:
			goto st33
		case 62:
			goto st15
		case 95:
			goto st13
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		goto st14
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		switch data[p] {
		case 44:
			goto st15
		case 59:
			goto st15
		case 62:
			goto st15
		case 95:
			goto st34
		case 98:
			goto st64
		case 99:
			goto st70
		case 100:
			goto st77
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st34
			}
		default:
			goto st13
		}
		goto st14
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st35
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st35
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st35
			}
		default:
			goto st13
		}
		goto st14
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		switch data[p] {
		case 47:
			goto st36
		case 62:
			goto st15
		case 95:
			goto st35
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st35
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st35
			}
		default:
			goto st13
		}
		goto st14
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st37
		case 98:
			goto st58
		case 99:
			goto st60
		case 100:
			goto st62
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
		goto st14
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		switch data[p] {
		case 43:
			goto st38
		case 45:
			goto st38
		case 62:
			goto st15
		case 95:
			goto st57
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st57
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st57
			}
		default:
			goto st57
		}
		goto st14
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		switch data[p] {
		case 43:
			goto st38
		case 45:
			goto st38
		case 62:
			goto st15
		case 95:
			goto st39
		case 98:
			goto st40
		case 99:
			goto st46
		case 100:
			goto st53
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st39
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st39
			}
		default:
			goto st39
		}
		goto st14
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		switch data[p] {
		case 44:
			goto st15
		case 59:
			goto st15
		case 62:
			goto st15
		case 95:
			goto st39
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st38
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st39
				}
			case data[p] >= 65:
				goto st39
			}
		default:
			goto st39
		}
		goto st14
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		switch data[p] {
		case 44:
			goto st15
		case 59:
			goto st15
		case 62:
			goto st15
		case 95:
			goto st39
		case 97:
			goto st41
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st38
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st39
				}
			case data[p] >= 65:
				goto st39
			}
		default:
			goto st39
		}
		goto st14
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		switch data[p] {
		case 44:
			goto st15
		case 59:
			goto st15
		case 62:
			goto st15
		case 95:
			goto st39
		case 115:
			goto st42
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st38
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st39
				}
			case data[p] >= 65:
				goto st39
			}
		default:
			goto st39
		}
		goto st14
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		switch data[p] {
		case 44:
			goto st15
		case 59:
			goto st15
		case 62:
			goto st15
		case 95:
			goto st39
		case 101:
			goto st43
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st38
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st39
				}
			case data[p] >= 65:
				goto st39
			}
		default:
			goto st39
		}
		goto st14
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		switch data[p] {
		case 44:
			goto st15
		case 54:
			goto st44
		case 59:
			goto st15
		case 62:
			goto st15
		case 95:
			goto st39
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st38
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st39
				}
			case data[p] >= 65:
				goto st39
			}
		default:
			goto st39
		}
		goto st14
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
		switch data[p] {
		case 44:
			goto st15
		case 52:
			goto st45
		case 59:
			goto st15
		case 62:
			goto st15
		case 95:
			goto st39
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st38
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st39
				}
			case data[p] >= 65:
				goto st39
			}
		default:
			goto st39
		}
		goto st14
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		switch data[p] {
		case 44:
			goto tr32
		case 59:
			goto tr32
		case 62:
			goto tr32
		case 95:
			goto st39
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto tr69
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st39
				}
			case data[p] >= 65:
				goto st39
			}
		default:
			goto st39
		}
		goto tr31
tr69:
//line xss_170.rl:13

            return true
        
	goto st288
	st288:
		if p++; p == pe {
			goto _test_eof288
		}
	st_case_288:
//line xss_170.go:2003
		switch data[p] {
		case 43:
			goto st38
		case 45:
			goto st38
		case 62:
			goto st15
		case 95:
			goto st39
		case 98:
			goto st40
		case 99:
			goto st46
		case 100:
			goto st53
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st39
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st39
			}
		default:
			goto st39
		}
		goto st14
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		switch data[p] {
		case 44:
			goto st15
		case 59:
			goto st15
		case 62:
			goto st15
		case 95:
			goto st39
		case 104:
			goto st47
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st38
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st39
				}
			case data[p] >= 65:
				goto st39
			}
		default:
			goto st39
		}
		goto st14
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		switch data[p] {
		case 44:
			goto st15
		case 59:
			goto st15
		case 62:
			goto st15
		case 95:
			goto st39
		case 97:
			goto st48
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st38
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st39
				}
			case data[p] >= 65:
				goto st39
			}
		default:
			goto st39
		}
		goto st14
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		switch data[p] {
		case 44:
			goto st15
		case 59:
			goto st15
		case 62:
			goto st15
		case 95:
			goto st39
		case 114:
			goto st49
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st38
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st39
				}
			case data[p] >= 65:
				goto st39
			}
		default:
			goto st39
		}
		goto st14
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		switch data[p] {
		case 44:
			goto st15
		case 59:
			goto st15
		case 62:
			goto st15
		case 95:
			goto st39
		case 115:
			goto st50
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st38
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st39
				}
			case data[p] >= 65:
				goto st39
			}
		default:
			goto st39
		}
		goto st14
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		switch data[p] {
		case 44:
			goto st15
		case 59:
			goto st15
		case 62:
			goto st15
		case 95:
			goto st39
		case 101:
			goto st51
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st38
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st39
				}
			case data[p] >= 65:
				goto st39
			}
		default:
			goto st39
		}
		goto st14
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
		switch data[p] {
		case 44:
			goto st15
		case 59:
			goto st15
		case 62:
			goto st15
		case 95:
			goto st39
		case 116:
			goto st52
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st38
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st39
				}
			case data[p] >= 65:
				goto st39
			}
		default:
			goto st39
		}
		goto st14
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
		switch data[p] {
		case 44:
			goto st15
		case 59:
			goto st15
		case 95:
			goto st39
		}
		switch {
		case data[p] < 61:
			switch {
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st39
				}
			case data[p] >= 43:
				goto st38
			}
		case data[p] > 62:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st39
				}
			case data[p] >= 65:
				goto st39
			}
		default:
			goto st15
		}
		goto st14
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
		switch data[p] {
		case 44:
			goto st15
		case 59:
			goto st15
		case 62:
			goto st15
		case 95:
			goto st39
		case 97:
			goto st54
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st38
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st39
				}
			case data[p] >= 65:
				goto st39
			}
		default:
			goto st39
		}
		goto st14
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		switch data[p] {
		case 44:
			goto st15
		case 59:
			goto st15
		case 62:
			goto st15
		case 95:
			goto st39
		case 116:
			goto st55
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st38
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st39
				}
			case data[p] >= 65:
				goto st39
			}
		default:
			goto st39
		}
		goto st14
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
		switch data[p] {
		case 44:
			goto st15
		case 59:
			goto st15
		case 62:
			goto st15
		case 95:
			goto st39
		case 97:
			goto st56
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st38
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st39
				}
			case data[p] >= 65:
				goto st39
			}
		default:
			goto st39
		}
		goto st14
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		switch data[p] {
		case 44:
			goto st15
		case 58:
			goto st33
		case 59:
			goto st15
		case 62:
			goto st15
		case 95:
			goto st39
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st38
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st39
				}
			case data[p] >= 65:
				goto st39
			}
		default:
			goto st39
		}
		goto st14
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		switch data[p] {
		case 43:
			goto st38
		case 45:
			goto st38
		case 62:
			goto st15
		case 95:
			goto st39
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st39
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st39
			}
		default:
			goto st39
		}
		goto st14
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		switch data[p] {
		case 43:
			goto st38
		case 45:
			goto st38
		case 62:
			goto st15
		case 95:
			goto st57
		case 97:
			goto st59
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st57
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st57
			}
		default:
			goto st57
		}
		goto st14
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		switch data[p] {
		case 43:
			goto st38
		case 45:
			goto st38
		case 62:
			goto st15
		case 95:
			goto st39
		case 115:
			goto st42
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st39
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st39
			}
		default:
			goto st39
		}
		goto st14
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
		switch data[p] {
		case 43:
			goto st38
		case 45:
			goto st38
		case 62:
			goto st15
		case 95:
			goto st57
		case 104:
			goto st61
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st57
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st57
			}
		default:
			goto st57
		}
		goto st14
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		switch data[p] {
		case 43:
			goto st38
		case 45:
			goto st38
		case 62:
			goto st15
		case 95:
			goto st39
		case 97:
			goto st48
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st39
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st39
			}
		default:
			goto st39
		}
		goto st14
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		switch data[p] {
		case 43:
			goto st38
		case 45:
			goto st38
		case 62:
			goto st15
		case 95:
			goto st57
		case 97:
			goto st63
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st57
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st57
			}
		default:
			goto st57
		}
		goto st14
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		switch data[p] {
		case 43:
			goto st38
		case 45:
			goto st38
		case 62:
			goto st15
		case 95:
			goto st39
		case 116:
			goto st55
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st39
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st39
			}
		default:
			goto st39
		}
		goto st14
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st35
		case 97:
			goto st65
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st35
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st35
			}
		default:
			goto st13
		}
		goto st14
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		switch data[p] {
		case 47:
			goto st36
		case 62:
			goto st15
		case 95:
			goto st35
		case 115:
			goto st66
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st35
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st35
			}
		default:
			goto st13
		}
		goto st14
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
		switch data[p] {
		case 47:
			goto st36
		case 62:
			goto st15
		case 95:
			goto st35
		case 101:
			goto st67
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st35
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st35
			}
		default:
			goto st13
		}
		goto st14
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
		switch data[p] {
		case 47:
			goto st36
		case 54:
			goto st68
		case 62:
			goto st15
		case 95:
			goto st35
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st35
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st35
			}
		default:
			goto st13
		}
		goto st14
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
		switch data[p] {
		case 47:
			goto st36
		case 52:
			goto st69
		case 62:
			goto st15
		case 95:
			goto st35
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st35
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st35
			}
		default:
			goto st13
		}
		goto st14
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		switch data[p] {
		case 47:
			goto tr87
		case 62:
			goto tr32
		case 95:
			goto st35
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st35
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st35
			}
		default:
			goto st13
		}
		goto tr31
tr87:
//line xss_170.rl:13

            return true
        
	goto st289
	st289:
		if p++; p == pe {
			goto _test_eof289
		}
	st_case_289:
//line xss_170.go:2802
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st37
		case 98:
			goto st58
		case 99:
			goto st60
		case 100:
			goto st62
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
		goto st14
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st35
		case 104:
			goto st71
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st35
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st35
			}
		default:
			goto st13
		}
		goto st14
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
		switch data[p] {
		case 47:
			goto st36
		case 62:
			goto st15
		case 95:
			goto st35
		case 97:
			goto st72
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st35
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st35
			}
		default:
			goto st13
		}
		goto st14
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
		switch data[p] {
		case 47:
			goto st36
		case 62:
			goto st15
		case 95:
			goto st35
		case 114:
			goto st73
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st35
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st35
			}
		default:
			goto st13
		}
		goto st14
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
		switch data[p] {
		case 47:
			goto st36
		case 62:
			goto st15
		case 95:
			goto st35
		case 115:
			goto st74
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st35
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st35
			}
		default:
			goto st13
		}
		goto st14
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
		switch data[p] {
		case 47:
			goto st36
		case 62:
			goto st15
		case 95:
			goto st35
		case 101:
			goto st75
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st35
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st35
			}
		default:
			goto st13
		}
		goto st14
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
		switch data[p] {
		case 47:
			goto st36
		case 62:
			goto st15
		case 95:
			goto st35
		case 116:
			goto st76
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st35
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st35
			}
		default:
			goto st13
		}
		goto st14
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
		switch data[p] {
		case 47:
			goto st36
		case 95:
			goto st35
		}
		switch {
		case data[p] < 61:
			if 48 <= data[p] && data[p] <= 57 {
				goto st35
			}
		case data[p] > 62:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st35
				}
			case data[p] >= 65:
				goto st13
			}
		default:
			goto st15
		}
		goto st14
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st35
		case 97:
			goto st78
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st35
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st35
			}
		default:
			goto st13
		}
		goto st14
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
		switch data[p] {
		case 47:
			goto st36
		case 62:
			goto st15
		case 95:
			goto st35
		case 116:
			goto st79
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st35
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st35
			}
		default:
			goto st13
		}
		goto st14
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
		switch data[p] {
		case 47:
			goto st36
		case 62:
			goto st15
		case 95:
			goto st35
		case 97:
			goto st80
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st35
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st35
			}
		default:
			goto st13
		}
		goto st14
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
		switch data[p] {
		case 47:
			goto st36
		case 58:
			goto st33
		case 62:
			goto st15
		case 95:
			goto st35
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st35
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st35
			}
		default:
			goto st13
		}
		goto st14
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 97:
			goto st82
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto st10
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 115:
			goto st83
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto st10
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 101:
			goto st84
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto st10
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
		switch data[p] {
		case 54:
			goto st85
		case 60:
			goto st12
		case 95:
			goto st11
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto st10
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
		switch data[p] {
		case 52:
			goto st86
		case 60:
			goto st12
		case 95:
			goto st11
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto st10
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
		switch data[p] {
		case 60:
			goto tr103
		case 95:
			goto st11
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto tr102
tr102:
//line xss_170.rl:13

            return true
        
	goto st290
	st290:
		if p++; p == pe {
			goto _test_eof290
		}
	st_case_290:
//line xss_170.go:3298
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 98:
			goto st81
		case 99:
			goto st87
		case 100:
			goto st95
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto st10
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 104:
			goto st88
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto st10
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 97:
			goto st89
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto st10
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 114:
			goto st90
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto st10
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 115:
			goto st91
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto st10
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 101:
			goto st92
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto st10
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 116:
			goto st93
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto st10
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
		switch data[p] {
		case 60:
			goto st12
		case 61:
			goto st94
		case 95:
			goto st11
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto st10
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
		switch data[p] {
		case 60:
			goto tr103
		case 95:
			goto st11
		case 98:
			goto st81
		case 99:
			goto st87
		case 100:
			goto st95
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto tr102
tr103:
//line xss_170.rl:13

            return true
        
	goto st291
	st291:
		if p++; p == pe {
			goto _test_eof291
		}
	st_case_291:
//line xss_170.go:3547
		switch data[p] {
		case 95:
			goto st13
		case 98:
			goto st16
		case 99:
			goto st22
		case 100:
			goto st29
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		goto st12
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 97:
			goto st96
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto st10
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 116:
			goto st97
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto st10
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 97:
			goto st98
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto st10
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
		switch data[p] {
		case 58:
			goto st99
		case 60:
			goto st12
		case 95:
			goto st11
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto st10
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
		switch data[p] {
		case 44:
			goto st94
		case 59:
			goto st94
		case 60:
			goto st12
		case 95:
			goto st100
		case 98:
			goto st130
		case 99:
			goto st136
		case 100:
			goto st143
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st100
			}
		default:
			goto st11
		}
		goto st10
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st101
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st101
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st101
			}
		default:
			goto st11
		}
		goto st10
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
		switch data[p] {
		case 47:
			goto st102
		case 60:
			goto st12
		case 95:
			goto st101
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st101
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st101
			}
		default:
			goto st11
		}
		goto st10
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st103
		case 98:
			goto st124
		case 99:
			goto st126
		case 100:
			goto st128
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st103
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st103
			}
		default:
			goto st103
		}
		goto st10
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
		switch data[p] {
		case 43:
			goto st104
		case 45:
			goto st104
		case 60:
			goto st12
		case 95:
			goto st123
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st123
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st123
			}
		default:
			goto st123
		}
		goto st10
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
		switch data[p] {
		case 43:
			goto st104
		case 45:
			goto st104
		case 60:
			goto st12
		case 95:
			goto st105
		case 98:
			goto st106
		case 99:
			goto st112
		case 100:
			goto st119
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st105
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st105
			}
		default:
			goto st105
		}
		goto st10
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
		switch data[p] {
		case 44:
			goto st94
		case 59:
			goto st94
		case 60:
			goto st12
		case 95:
			goto st105
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st104
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st105
				}
			case data[p] >= 65:
				goto st105
			}
		default:
			goto st105
		}
		goto st10
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
		switch data[p] {
		case 44:
			goto st94
		case 59:
			goto st94
		case 60:
			goto st12
		case 95:
			goto st105
		case 97:
			goto st107
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st104
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st105
				}
			case data[p] >= 65:
				goto st105
			}
		default:
			goto st105
		}
		goto st10
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
		switch data[p] {
		case 44:
			goto st94
		case 59:
			goto st94
		case 60:
			goto st12
		case 95:
			goto st105
		case 115:
			goto st108
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st104
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st105
				}
			case data[p] >= 65:
				goto st105
			}
		default:
			goto st105
		}
		goto st10
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
		switch data[p] {
		case 44:
			goto st94
		case 59:
			goto st94
		case 60:
			goto st12
		case 95:
			goto st105
		case 101:
			goto st109
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st104
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st105
				}
			case data[p] >= 65:
				goto st105
			}
		default:
			goto st105
		}
		goto st10
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
		switch data[p] {
		case 44:
			goto st94
		case 54:
			goto st110
		case 59:
			goto st94
		case 60:
			goto st12
		case 95:
			goto st105
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st104
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st105
				}
			case data[p] >= 65:
				goto st105
			}
		default:
			goto st105
		}
		goto st10
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
		switch data[p] {
		case 44:
			goto st94
		case 52:
			goto st111
		case 59:
			goto st94
		case 60:
			goto st12
		case 95:
			goto st105
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st104
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st105
				}
			case data[p] >= 65:
				goto st105
			}
		default:
			goto st105
		}
		goto st10
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
		switch data[p] {
		case 44:
			goto tr137
		case 59:
			goto tr137
		case 60:
			goto tr103
		case 95:
			goto st105
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto tr136
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st105
				}
			case data[p] >= 65:
				goto st105
			}
		default:
			goto st105
		}
		goto tr102
tr136:
//line xss_170.rl:13

            return true
        
	goto st292
	st292:
		if p++; p == pe {
			goto _test_eof292
		}
	st_case_292:
//line xss_170.go:4103
		switch data[p] {
		case 43:
			goto st104
		case 45:
			goto st104
		case 60:
			goto st12
		case 95:
			goto st105
		case 98:
			goto st106
		case 99:
			goto st112
		case 100:
			goto st119
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st105
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st105
			}
		default:
			goto st105
		}
		goto st10
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
		switch data[p] {
		case 44:
			goto st94
		case 59:
			goto st94
		case 60:
			goto st12
		case 95:
			goto st105
		case 104:
			goto st113
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st104
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st105
				}
			case data[p] >= 65:
				goto st105
			}
		default:
			goto st105
		}
		goto st10
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
		switch data[p] {
		case 44:
			goto st94
		case 59:
			goto st94
		case 60:
			goto st12
		case 95:
			goto st105
		case 97:
			goto st114
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st104
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st105
				}
			case data[p] >= 65:
				goto st105
			}
		default:
			goto st105
		}
		goto st10
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
		switch data[p] {
		case 44:
			goto st94
		case 59:
			goto st94
		case 60:
			goto st12
		case 95:
			goto st105
		case 114:
			goto st115
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st104
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st105
				}
			case data[p] >= 65:
				goto st105
			}
		default:
			goto st105
		}
		goto st10
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
		switch data[p] {
		case 44:
			goto st94
		case 59:
			goto st94
		case 60:
			goto st12
		case 95:
			goto st105
		case 115:
			goto st116
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st104
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st105
				}
			case data[p] >= 65:
				goto st105
			}
		default:
			goto st105
		}
		goto st10
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
		switch data[p] {
		case 44:
			goto st94
		case 59:
			goto st94
		case 60:
			goto st12
		case 95:
			goto st105
		case 101:
			goto st117
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st104
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st105
				}
			case data[p] >= 65:
				goto st105
			}
		default:
			goto st105
		}
		goto st10
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
		switch data[p] {
		case 44:
			goto st94
		case 59:
			goto st94
		case 60:
			goto st12
		case 95:
			goto st105
		case 116:
			goto st118
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st104
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st105
				}
			case data[p] >= 65:
				goto st105
			}
		default:
			goto st105
		}
		goto st10
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
		switch data[p] {
		case 44:
			goto st94
		case 60:
			goto st12
		case 95:
			goto st105
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st105
				}
			case data[p] >= 43:
				goto st104
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st105
				}
			case data[p] >= 65:
				goto st105
			}
		default:
			goto st94
		}
		goto st10
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
		switch data[p] {
		case 44:
			goto st94
		case 59:
			goto st94
		case 60:
			goto st12
		case 95:
			goto st105
		case 97:
			goto st120
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st104
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st105
				}
			case data[p] >= 65:
				goto st105
			}
		default:
			goto st105
		}
		goto st10
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
		switch data[p] {
		case 44:
			goto st94
		case 59:
			goto st94
		case 60:
			goto st12
		case 95:
			goto st105
		case 116:
			goto st121
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st104
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st105
				}
			case data[p] >= 65:
				goto st105
			}
		default:
			goto st105
		}
		goto st10
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
		switch data[p] {
		case 44:
			goto st94
		case 59:
			goto st94
		case 60:
			goto st12
		case 95:
			goto st105
		case 97:
			goto st122
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st104
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st105
				}
			case data[p] >= 65:
				goto st105
			}
		default:
			goto st105
		}
		goto st10
	st122:
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
		switch data[p] {
		case 44:
			goto st94
		case 58:
			goto st99
		case 59:
			goto st94
		case 60:
			goto st12
		case 95:
			goto st105
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st104
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st105
				}
			case data[p] >= 65:
				goto st105
			}
		default:
			goto st105
		}
		goto st10
tr137:
//line xss_170.rl:13

            return true
        
	goto st293
	st293:
		if p++; p == pe {
			goto _test_eof293
		}
	st_case_293:
//line xss_170.go:4530
		switch data[p] {
		case 60:
			goto tr103
		case 95:
			goto st11
		case 98:
			goto st81
		case 99:
			goto st87
		case 100:
			goto st95
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto tr102
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
		switch data[p] {
		case 43:
			goto st104
		case 45:
			goto st104
		case 60:
			goto st12
		case 95:
			goto st105
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st105
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st105
			}
		default:
			goto st105
		}
		goto st10
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
		switch data[p] {
		case 43:
			goto st104
		case 45:
			goto st104
		case 60:
			goto st12
		case 95:
			goto st123
		case 97:
			goto st125
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st123
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st123
			}
		default:
			goto st123
		}
		goto st10
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
		switch data[p] {
		case 43:
			goto st104
		case 45:
			goto st104
		case 60:
			goto st12
		case 95:
			goto st105
		case 115:
			goto st108
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st105
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st105
			}
		default:
			goto st105
		}
		goto st10
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
		switch data[p] {
		case 43:
			goto st104
		case 45:
			goto st104
		case 60:
			goto st12
		case 95:
			goto st123
		case 104:
			goto st127
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st123
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st123
			}
		default:
			goto st123
		}
		goto st10
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
		switch data[p] {
		case 43:
			goto st104
		case 45:
			goto st104
		case 60:
			goto st12
		case 95:
			goto st105
		case 97:
			goto st114
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st105
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st105
			}
		default:
			goto st105
		}
		goto st10
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
		switch data[p] {
		case 43:
			goto st104
		case 45:
			goto st104
		case 60:
			goto st12
		case 95:
			goto st123
		case 97:
			goto st129
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st123
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st123
			}
		default:
			goto st123
		}
		goto st10
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
		switch data[p] {
		case 43:
			goto st104
		case 45:
			goto st104
		case 60:
			goto st12
		case 95:
			goto st105
		case 116:
			goto st121
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st105
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st105
			}
		default:
			goto st105
		}
		goto st10
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st101
		case 97:
			goto st131
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st101
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st101
			}
		default:
			goto st11
		}
		goto st10
	st131:
		if p++; p == pe {
			goto _test_eof131
		}
	st_case_131:
		switch data[p] {
		case 47:
			goto st102
		case 60:
			goto st12
		case 95:
			goto st101
		case 115:
			goto st132
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st101
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st101
			}
		default:
			goto st11
		}
		goto st10
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
		switch data[p] {
		case 47:
			goto st102
		case 60:
			goto st12
		case 95:
			goto st101
		case 101:
			goto st133
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st101
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st101
			}
		default:
			goto st11
		}
		goto st10
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
		switch data[p] {
		case 47:
			goto st102
		case 54:
			goto st134
		case 60:
			goto st12
		case 95:
			goto st101
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st101
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st101
			}
		default:
			goto st11
		}
		goto st10
	st134:
		if p++; p == pe {
			goto _test_eof134
		}
	st_case_134:
		switch data[p] {
		case 47:
			goto st102
		case 52:
			goto st135
		case 60:
			goto st12
		case 95:
			goto st101
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st101
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st101
			}
		default:
			goto st11
		}
		goto st10
	st135:
		if p++; p == pe {
			goto _test_eof135
		}
	st_case_135:
		switch data[p] {
		case 47:
			goto tr155
		case 60:
			goto tr103
		case 95:
			goto st101
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st101
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st101
			}
		default:
			goto st11
		}
		goto tr102
tr155:
//line xss_170.rl:13

            return true
        
	goto st294
	st294:
		if p++; p == pe {
			goto _test_eof294
		}
	st_case_294:
//line xss_170.go:4939
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st103
		case 98:
			goto st124
		case 99:
			goto st126
		case 100:
			goto st128
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st103
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st103
			}
		default:
			goto st103
		}
		goto st10
	st136:
		if p++; p == pe {
			goto _test_eof136
		}
	st_case_136:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st101
		case 104:
			goto st137
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st101
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st101
			}
		default:
			goto st11
		}
		goto st10
	st137:
		if p++; p == pe {
			goto _test_eof137
		}
	st_case_137:
		switch data[p] {
		case 47:
			goto st102
		case 60:
			goto st12
		case 95:
			goto st101
		case 97:
			goto st138
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st101
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st101
			}
		default:
			goto st11
		}
		goto st10
	st138:
		if p++; p == pe {
			goto _test_eof138
		}
	st_case_138:
		switch data[p] {
		case 47:
			goto st102
		case 60:
			goto st12
		case 95:
			goto st101
		case 114:
			goto st139
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st101
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st101
			}
		default:
			goto st11
		}
		goto st10
	st139:
		if p++; p == pe {
			goto _test_eof139
		}
	st_case_139:
		switch data[p] {
		case 47:
			goto st102
		case 60:
			goto st12
		case 95:
			goto st101
		case 115:
			goto st140
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st101
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st101
			}
		default:
			goto st11
		}
		goto st10
	st140:
		if p++; p == pe {
			goto _test_eof140
		}
	st_case_140:
		switch data[p] {
		case 47:
			goto st102
		case 60:
			goto st12
		case 95:
			goto st101
		case 101:
			goto st141
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st101
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st101
			}
		default:
			goto st11
		}
		goto st10
	st141:
		if p++; p == pe {
			goto _test_eof141
		}
	st_case_141:
		switch data[p] {
		case 47:
			goto st102
		case 60:
			goto st12
		case 95:
			goto st101
		case 116:
			goto st142
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st101
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st101
			}
		default:
			goto st11
		}
		goto st10
	st142:
		if p++; p == pe {
			goto _test_eof142
		}
	st_case_142:
		switch data[p] {
		case 47:
			goto st102
		case 60:
			goto st12
		case 61:
			goto st94
		case 95:
			goto st101
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st101
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st101
			}
		default:
			goto st11
		}
		goto st10
	st143:
		if p++; p == pe {
			goto _test_eof143
		}
	st_case_143:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st101
		case 97:
			goto st144
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st101
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st101
			}
		default:
			goto st11
		}
		goto st10
	st144:
		if p++; p == pe {
			goto _test_eof144
		}
	st_case_144:
		switch data[p] {
		case 47:
			goto st102
		case 60:
			goto st12
		case 95:
			goto st101
		case 116:
			goto st145
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st101
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st101
			}
		default:
			goto st11
		}
		goto st10
	st145:
		if p++; p == pe {
			goto _test_eof145
		}
	st_case_145:
		switch data[p] {
		case 47:
			goto st102
		case 60:
			goto st12
		case 95:
			goto st101
		case 97:
			goto st146
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st101
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st101
			}
		default:
			goto st11
		}
		goto st10
	st146:
		if p++; p == pe {
			goto _test_eof146
		}
	st_case_146:
		switch data[p] {
		case 47:
			goto st102
		case 58:
			goto st99
		case 60:
			goto st12
		case 95:
			goto st101
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st101
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st101
			}
		default:
			goto st11
		}
		goto st10
	st147:
		if p++; p == pe {
			goto _test_eof147
		}
	st_case_147:
		switch data[p] {
		case 59:
			goto st12
		case 95:
			goto st148
		case 100:
			goto st151
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st148
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st148
			}
		default:
			goto st148
		}
		goto st147
	st148:
		if p++; p == pe {
			goto _test_eof148
		}
	st_case_148:
		switch data[p] {
		case 59:
			goto st14
		case 62:
			goto st150
		case 95:
			goto st148
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st148
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st148
			}
		default:
			goto st148
		}
		goto st149
	st149:
		if p++; p == pe {
			goto _test_eof149
		}
	st_case_149:
		switch data[p] {
		case 59:
			goto st14
		case 62:
			goto st150
		case 95:
			goto st148
		case 100:
			goto st151
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st148
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st148
			}
		default:
			goto st148
		}
		goto st149
	st150:
		if p++; p == pe {
			goto _test_eof150
		}
	st_case_150:
		switch data[p] {
		case 59:
			goto tr31
		case 62:
			goto tr170
		case 95:
			goto st148
		case 100:
			goto st151
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st148
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st148
			}
		default:
			goto st148
		}
		goto tr169
tr169:
//line xss_170.rl:13

            return true
        
	goto st295
	st295:
		if p++; p == pe {
			goto _test_eof295
		}
	st_case_295:
//line xss_170.go:5388
		switch data[p] {
		case 59:
			goto st14
		case 62:
			goto st150
		case 95:
			goto st148
		case 100:
			goto st151
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st148
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st148
			}
		default:
			goto st148
		}
		goto st149
	st151:
		if p++; p == pe {
			goto _test_eof151
		}
	st_case_151:
		switch data[p] {
		case 59:
			goto st14
		case 62:
			goto st150
		case 95:
			goto st148
		case 97:
			goto st152
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st148
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st148
			}
		default:
			goto st148
		}
		goto st149
	st152:
		if p++; p == pe {
			goto _test_eof152
		}
	st_case_152:
		switch data[p] {
		case 59:
			goto st14
		case 62:
			goto st150
		case 95:
			goto st148
		case 116:
			goto st153
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st148
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st148
			}
		default:
			goto st148
		}
		goto st149
	st153:
		if p++; p == pe {
			goto _test_eof153
		}
	st_case_153:
		switch data[p] {
		case 59:
			goto st14
		case 62:
			goto st150
		case 95:
			goto st148
		case 97:
			goto st154
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st148
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st148
			}
		default:
			goto st148
		}
		goto st149
	st154:
		if p++; p == pe {
			goto _test_eof154
		}
	st_case_154:
		switch data[p] {
		case 58:
			goto st155
		case 59:
			goto st14
		case 62:
			goto st150
		case 95:
			goto st148
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st148
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st148
			}
		default:
			goto st148
		}
		goto st149
	st155:
		if p++; p == pe {
			goto _test_eof155
		}
	st_case_155:
		switch data[p] {
		case 44:
			goto st150
		case 59:
			goto st15
		case 62:
			goto st150
		case 95:
			goto st156
		case 100:
			goto st169
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st148
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st156
			}
		default:
			goto st148
		}
		goto st149
	st156:
		if p++; p == pe {
			goto _test_eof156
		}
	st_case_156:
		switch data[p] {
		case 59:
			goto st14
		case 62:
			goto st150
		case 95:
			goto st157
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st157
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st157
			}
		default:
			goto st148
		}
		goto st149
	st157:
		if p++; p == pe {
			goto _test_eof157
		}
	st_case_157:
		switch data[p] {
		case 47:
			goto st158
		case 59:
			goto st14
		case 62:
			goto st150
		case 95:
			goto st157
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st157
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st157
			}
		default:
			goto st148
		}
		goto st149
	st158:
		if p++; p == pe {
			goto _test_eof158
		}
	st_case_158:
		switch data[p] {
		case 59:
			goto st14
		case 62:
			goto st150
		case 95:
			goto st159
		case 100:
			goto st167
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st159
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st159
			}
		default:
			goto st159
		}
		goto st149
	st159:
		if p++; p == pe {
			goto _test_eof159
		}
	st_case_159:
		switch data[p] {
		case 43:
			goto st160
		case 45:
			goto st160
		case 59:
			goto st14
		case 62:
			goto st150
		case 95:
			goto st166
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st166
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st166
			}
		default:
			goto st166
		}
		goto st149
	st160:
		if p++; p == pe {
			goto _test_eof160
		}
	st_case_160:
		switch data[p] {
		case 43:
			goto st160
		case 45:
			goto st160
		case 59:
			goto st14
		case 62:
			goto st150
		case 95:
			goto st161
		case 100:
			goto st162
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st161
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st161
			}
		default:
			goto st161
		}
		goto st149
	st161:
		if p++; p == pe {
			goto _test_eof161
		}
	st_case_161:
		switch data[p] {
		case 44:
			goto st150
		case 59:
			goto st15
		case 62:
			goto st150
		case 95:
			goto st161
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st160
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st161
				}
			case data[p] >= 65:
				goto st161
			}
		default:
			goto st161
		}
		goto st149
	st162:
		if p++; p == pe {
			goto _test_eof162
		}
	st_case_162:
		switch data[p] {
		case 44:
			goto st150
		case 59:
			goto st15
		case 62:
			goto st150
		case 95:
			goto st161
		case 97:
			goto st163
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st160
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st161
				}
			case data[p] >= 65:
				goto st161
			}
		default:
			goto st161
		}
		goto st149
	st163:
		if p++; p == pe {
			goto _test_eof163
		}
	st_case_163:
		switch data[p] {
		case 44:
			goto st150
		case 59:
			goto st15
		case 62:
			goto st150
		case 95:
			goto st161
		case 116:
			goto st164
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st160
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st161
				}
			case data[p] >= 65:
				goto st161
			}
		default:
			goto st161
		}
		goto st149
	st164:
		if p++; p == pe {
			goto _test_eof164
		}
	st_case_164:
		switch data[p] {
		case 44:
			goto st150
		case 59:
			goto st15
		case 62:
			goto st150
		case 95:
			goto st161
		case 97:
			goto st165
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st160
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st161
				}
			case data[p] >= 65:
				goto st161
			}
		default:
			goto st161
		}
		goto st149
	st165:
		if p++; p == pe {
			goto _test_eof165
		}
	st_case_165:
		switch data[p] {
		case 44:
			goto st150
		case 58:
			goto st155
		case 59:
			goto st15
		case 62:
			goto st150
		case 95:
			goto st161
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st160
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st161
				}
			case data[p] >= 65:
				goto st161
			}
		default:
			goto st161
		}
		goto st149
	st166:
		if p++; p == pe {
			goto _test_eof166
		}
	st_case_166:
		switch data[p] {
		case 43:
			goto st160
		case 45:
			goto st160
		case 59:
			goto st14
		case 62:
			goto st150
		case 95:
			goto st161
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st161
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st161
			}
		default:
			goto st161
		}
		goto st149
	st167:
		if p++; p == pe {
			goto _test_eof167
		}
	st_case_167:
		switch data[p] {
		case 43:
			goto st160
		case 45:
			goto st160
		case 59:
			goto st14
		case 62:
			goto st150
		case 95:
			goto st166
		case 97:
			goto st168
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st166
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st166
			}
		default:
			goto st166
		}
		goto st149
	st168:
		if p++; p == pe {
			goto _test_eof168
		}
	st_case_168:
		switch data[p] {
		case 43:
			goto st160
		case 45:
			goto st160
		case 59:
			goto st14
		case 62:
			goto st150
		case 95:
			goto st161
		case 116:
			goto st164
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st161
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st161
			}
		default:
			goto st161
		}
		goto st149
	st169:
		if p++; p == pe {
			goto _test_eof169
		}
	st_case_169:
		switch data[p] {
		case 59:
			goto st14
		case 62:
			goto st150
		case 95:
			goto st157
		case 97:
			goto st170
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st157
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st157
			}
		default:
			goto st148
		}
		goto st149
	st170:
		if p++; p == pe {
			goto _test_eof170
		}
	st_case_170:
		switch data[p] {
		case 47:
			goto st158
		case 59:
			goto st14
		case 62:
			goto st150
		case 95:
			goto st157
		case 116:
			goto st171
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st157
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st157
			}
		default:
			goto st148
		}
		goto st149
	st171:
		if p++; p == pe {
			goto _test_eof171
		}
	st_case_171:
		switch data[p] {
		case 47:
			goto st158
		case 59:
			goto st14
		case 62:
			goto st150
		case 95:
			goto st157
		case 97:
			goto st172
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st157
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st157
			}
		default:
			goto st148
		}
		goto st149
	st172:
		if p++; p == pe {
			goto _test_eof172
		}
	st_case_172:
		switch data[p] {
		case 47:
			goto st158
		case 58:
			goto st155
		case 59:
			goto st14
		case 62:
			goto st150
		case 95:
			goto st157
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st157
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st157
			}
		default:
			goto st148
		}
		goto st149
tr170:
//line xss_170.rl:13

            return true
        
	goto st296
	st296:
		if p++; p == pe {
			goto _test_eof296
		}
	st_case_296:
//line xss_170.go:6094
		switch data[p] {
		case 59:
			goto tr31
		case 62:
			goto tr170
		case 95:
			goto st148
		case 100:
			goto st151
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st148
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st148
			}
		default:
			goto st148
		}
		goto tr169
	st173:
		if p++; p == pe {
			goto _test_eof173
		}
	st_case_173:
		switch data[p] {
		case 59:
			goto st10
		case 60:
			goto st147
		case 95:
			goto st9
		case 97:
			goto st174
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st9
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st9
			}
		default:
			goto st9
		}
		goto st8
	st174:
		if p++; p == pe {
			goto _test_eof174
		}
	st_case_174:
		switch data[p] {
		case 59:
			goto st10
		case 60:
			goto st147
		case 95:
			goto st9
		case 116:
			goto st175
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st9
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st9
			}
		default:
			goto st9
		}
		goto st8
	st175:
		if p++; p == pe {
			goto _test_eof175
		}
	st_case_175:
		switch data[p] {
		case 59:
			goto st10
		case 60:
			goto st147
		case 95:
			goto st9
		case 97:
			goto st176
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st9
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st9
			}
		default:
			goto st9
		}
		goto st8
	st176:
		if p++; p == pe {
			goto _test_eof176
		}
	st_case_176:
		switch data[p] {
		case 58:
			goto st177
		case 59:
			goto st10
		case 60:
			goto st147
		case 95:
			goto st9
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st9
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st9
			}
		default:
			goto st9
		}
		goto st8
	st177:
		if p++; p == pe {
			goto _test_eof177
		}
	st_case_177:
		switch data[p] {
		case 44:
			goto st178
		case 59:
			goto st94
		case 60:
			goto st147
		case 95:
			goto st179
		case 100:
			goto st192
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st9
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st179
			}
		default:
			goto st9
		}
		goto st8
	st178:
		if p++; p == pe {
			goto _test_eof178
		}
	st_case_178:
		switch data[p] {
		case 59:
			goto tr102
		case 60:
			goto tr199
		case 95:
			goto st9
		case 100:
			goto st173
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st9
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st9
			}
		default:
			goto st9
		}
		goto tr198
tr198:
//line xss_170.rl:13

            return true
        
	goto st297
	st297:
		if p++; p == pe {
			goto _test_eof297
		}
	st_case_297:
//line xss_170.go:6299
		switch data[p] {
		case 59:
			goto st10
		case 60:
			goto st147
		case 95:
			goto st9
		case 100:
			goto st173
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st9
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st9
			}
		default:
			goto st9
		}
		goto st8
tr199:
//line xss_170.rl:13

            return true
        
	goto st298
	st298:
		if p++; p == pe {
			goto _test_eof298
		}
	st_case_298:
//line xss_170.go:6334
		switch data[p] {
		case 59:
			goto st12
		case 95:
			goto st148
		case 100:
			goto st151
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st148
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st148
			}
		default:
			goto st148
		}
		goto st147
	st179:
		if p++; p == pe {
			goto _test_eof179
		}
	st_case_179:
		switch data[p] {
		case 59:
			goto st10
		case 60:
			goto st147
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
			goto st9
		}
		goto st8
	st180:
		if p++; p == pe {
			goto _test_eof180
		}
	st_case_180:
		switch data[p] {
		case 47:
			goto st181
		case 59:
			goto st10
		case 60:
			goto st147
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
			goto st9
		}
		goto st8
	st181:
		if p++; p == pe {
			goto _test_eof181
		}
	st_case_181:
		switch data[p] {
		case 59:
			goto st10
		case 60:
			goto st147
		case 95:
			goto st182
		case 100:
			goto st190
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st182
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st182
			}
		default:
			goto st182
		}
		goto st8
	st182:
		if p++; p == pe {
			goto _test_eof182
		}
	st_case_182:
		switch data[p] {
		case 43:
			goto st183
		case 45:
			goto st183
		case 59:
			goto st10
		case 60:
			goto st147
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
		goto st8
	st183:
		if p++; p == pe {
			goto _test_eof183
		}
	st_case_183:
		switch data[p] {
		case 43:
			goto st183
		case 45:
			goto st183
		case 59:
			goto st10
		case 60:
			goto st147
		case 95:
			goto st184
		case 100:
			goto st185
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st184
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st184
			}
		default:
			goto st184
		}
		goto st8
	st184:
		if p++; p == pe {
			goto _test_eof184
		}
	st_case_184:
		switch data[p] {
		case 44:
			goto st178
		case 59:
			goto st94
		case 60:
			goto st147
		case 95:
			goto st184
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st183
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st184
				}
			case data[p] >= 65:
				goto st184
			}
		default:
			goto st184
		}
		goto st8
	st185:
		if p++; p == pe {
			goto _test_eof185
		}
	st_case_185:
		switch data[p] {
		case 44:
			goto st178
		case 59:
			goto st94
		case 60:
			goto st147
		case 95:
			goto st184
		case 97:
			goto st186
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st183
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st184
				}
			case data[p] >= 65:
				goto st184
			}
		default:
			goto st184
		}
		goto st8
	st186:
		if p++; p == pe {
			goto _test_eof186
		}
	st_case_186:
		switch data[p] {
		case 44:
			goto st178
		case 59:
			goto st94
		case 60:
			goto st147
		case 95:
			goto st184
		case 116:
			goto st187
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st183
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st184
				}
			case data[p] >= 65:
				goto st184
			}
		default:
			goto st184
		}
		goto st8
	st187:
		if p++; p == pe {
			goto _test_eof187
		}
	st_case_187:
		switch data[p] {
		case 44:
			goto st178
		case 59:
			goto st94
		case 60:
			goto st147
		case 95:
			goto st184
		case 97:
			goto st188
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st183
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st184
				}
			case data[p] >= 65:
				goto st184
			}
		default:
			goto st184
		}
		goto st8
	st188:
		if p++; p == pe {
			goto _test_eof188
		}
	st_case_188:
		switch data[p] {
		case 44:
			goto st178
		case 58:
			goto st177
		case 59:
			goto st94
		case 60:
			goto st147
		case 95:
			goto st184
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st183
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st184
				}
			case data[p] >= 65:
				goto st184
			}
		default:
			goto st184
		}
		goto st8
	st189:
		if p++; p == pe {
			goto _test_eof189
		}
	st_case_189:
		switch data[p] {
		case 43:
			goto st183
		case 45:
			goto st183
		case 59:
			goto st10
		case 60:
			goto st147
		case 95:
			goto st184
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st184
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st184
			}
		default:
			goto st184
		}
		goto st8
	st190:
		if p++; p == pe {
			goto _test_eof190
		}
	st_case_190:
		switch data[p] {
		case 43:
			goto st183
		case 45:
			goto st183
		case 59:
			goto st10
		case 60:
			goto st147
		case 95:
			goto st189
		case 97:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st189
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st189
			}
		default:
			goto st189
		}
		goto st8
	st191:
		if p++; p == pe {
			goto _test_eof191
		}
	st_case_191:
		switch data[p] {
		case 43:
			goto st183
		case 45:
			goto st183
		case 59:
			goto st10
		case 60:
			goto st147
		case 95:
			goto st184
		case 116:
			goto st187
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st184
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st184
			}
		default:
			goto st184
		}
		goto st8
	st192:
		if p++; p == pe {
			goto _test_eof192
		}
	st_case_192:
		switch data[p] {
		case 59:
			goto st10
		case 60:
			goto st147
		case 95:
			goto st180
		case 97:
			goto st193
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st180
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st180
			}
		default:
			goto st9
		}
		goto st8
	st193:
		if p++; p == pe {
			goto _test_eof193
		}
	st_case_193:
		switch data[p] {
		case 47:
			goto st181
		case 59:
			goto st10
		case 60:
			goto st147
		case 95:
			goto st180
		case 116:
			goto st194
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
			goto st9
		}
		goto st8
	st194:
		if p++; p == pe {
			goto _test_eof194
		}
	st_case_194:
		switch data[p] {
		case 47:
			goto st181
		case 59:
			goto st10
		case 60:
			goto st147
		case 95:
			goto st180
		case 97:
			goto st195
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st180
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st180
			}
		default:
			goto st9
		}
		goto st8
	st195:
		if p++; p == pe {
			goto _test_eof195
		}
	st_case_195:
		switch data[p] {
		case 47:
			goto st181
		case 58:
			goto st177
		case 59:
			goto st10
		case 60:
			goto st147
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
			goto st9
		}
		goto st8
	st196:
		if p++; p == pe {
			goto _test_eof196
		}
	st_case_196:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st197
		case 95:
			goto st196
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st196
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st196
			}
		default:
			goto st196
		}
		goto st7
	st197:
		if p++; p == pe {
			goto _test_eof197
		}
	st_case_197:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st198
		case 98:
			goto st199
		case 99:
			goto st205
		case 100:
			goto st213
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st198
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st198
			}
		default:
			goto st198
		}
		goto st197
	st198:
		if p++; p == pe {
			goto _test_eof198
		}
	st_case_198:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st198
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st198
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st198
			}
		default:
			goto st198
		}
		goto st197
	st199:
		if p++; p == pe {
			goto _test_eof199
		}
	st_case_199:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st198
		case 97:
			goto st200
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st198
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st198
			}
		default:
			goto st198
		}
		goto st197
	st200:
		if p++; p == pe {
			goto _test_eof200
		}
	st_case_200:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st198
		case 115:
			goto st201
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st198
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st198
			}
		default:
			goto st198
		}
		goto st197
	st201:
		if p++; p == pe {
			goto _test_eof201
		}
	st_case_201:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st198
		case 101:
			goto st202
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st198
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st198
			}
		default:
			goto st198
		}
		goto st197
	st202:
		if p++; p == pe {
			goto _test_eof202
		}
	st_case_202:
		switch data[p] {
		case 44:
			goto st10
		case 54:
			goto st203
		case 95:
			goto st198
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st198
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st198
			}
		default:
			goto st198
		}
		goto st197
	st203:
		if p++; p == pe {
			goto _test_eof203
		}
	st_case_203:
		switch data[p] {
		case 44:
			goto st10
		case 52:
			goto st204
		case 95:
			goto st198
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st198
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st198
			}
		default:
			goto st198
		}
		goto st197
	st204:
		if p++; p == pe {
			goto _test_eof204
		}
	st_case_204:
		switch data[p] {
		case 44:
			goto tr102
		case 95:
			goto st198
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st198
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st198
			}
		default:
			goto st198
		}
		goto tr224
tr224:
//line xss_170.rl:13

            return true
        
	goto st299
	st299:
		if p++; p == pe {
			goto _test_eof299
		}
	st_case_299:
//line xss_170.go:7130
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st198
		case 98:
			goto st199
		case 99:
			goto st205
		case 100:
			goto st213
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st198
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st198
			}
		default:
			goto st198
		}
		goto st197
	st205:
		if p++; p == pe {
			goto _test_eof205
		}
	st_case_205:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st198
		case 104:
			goto st206
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st198
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st198
			}
		default:
			goto st198
		}
		goto st197
	st206:
		if p++; p == pe {
			goto _test_eof206
		}
	st_case_206:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st198
		case 97:
			goto st207
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st198
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st198
			}
		default:
			goto st198
		}
		goto st197
	st207:
		if p++; p == pe {
			goto _test_eof207
		}
	st_case_207:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st198
		case 114:
			goto st208
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st198
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st198
			}
		default:
			goto st198
		}
		goto st197
	st208:
		if p++; p == pe {
			goto _test_eof208
		}
	st_case_208:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st198
		case 115:
			goto st209
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st198
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st198
			}
		default:
			goto st198
		}
		goto st197
	st209:
		if p++; p == pe {
			goto _test_eof209
		}
	st_case_209:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st198
		case 101:
			goto st210
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st198
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st198
			}
		default:
			goto st198
		}
		goto st197
	st210:
		if p++; p == pe {
			goto _test_eof210
		}
	st_case_210:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st198
		case 116:
			goto st211
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st198
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st198
			}
		default:
			goto st198
		}
		goto st197
	st211:
		if p++; p == pe {
			goto _test_eof211
		}
	st_case_211:
		switch data[p] {
		case 44:
			goto st10
		case 61:
			goto st212
		case 95:
			goto st198
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st198
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st198
			}
		default:
			goto st198
		}
		goto st197
	st212:
		if p++; p == pe {
			goto _test_eof212
		}
	st_case_212:
		switch data[p] {
		case 44:
			goto tr102
		case 95:
			goto st198
		case 98:
			goto st199
		case 99:
			goto st205
		case 100:
			goto st213
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st198
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st198
			}
		default:
			goto st198
		}
		goto tr224
	st213:
		if p++; p == pe {
			goto _test_eof213
		}
	st_case_213:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st198
		case 97:
			goto st214
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st198
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st198
			}
		default:
			goto st198
		}
		goto st197
	st214:
		if p++; p == pe {
			goto _test_eof214
		}
	st_case_214:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st198
		case 116:
			goto st215
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st198
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st198
			}
		default:
			goto st198
		}
		goto st197
	st215:
		if p++; p == pe {
			goto _test_eof215
		}
	st_case_215:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st198
		case 97:
			goto st216
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st198
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st198
			}
		default:
			goto st198
		}
		goto st197
	st216:
		if p++; p == pe {
			goto _test_eof216
		}
	st_case_216:
		switch data[p] {
		case 44:
			goto st10
		case 58:
			goto st217
		case 95:
			goto st198
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st198
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st198
			}
		default:
			goto st198
		}
		goto st197
	st217:
		if p++; p == pe {
			goto _test_eof217
		}
	st_case_217:
		switch data[p] {
		case 44:
			goto st94
		case 59:
			goto st212
		case 95:
			goto st218
		case 98:
			goto st248
		case 99:
			goto st254
		case 100:
			goto st261
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st198
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st218
			}
		default:
			goto st198
		}
		goto st197
	st218:
		if p++; p == pe {
			goto _test_eof218
		}
	st_case_218:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st219
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st219
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st219
			}
		default:
			goto st198
		}
		goto st197
	st219:
		if p++; p == pe {
			goto _test_eof219
		}
	st_case_219:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st220
		case 95:
			goto st219
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st219
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st219
			}
		default:
			goto st198
		}
		goto st197
	st220:
		if p++; p == pe {
			goto _test_eof220
		}
	st_case_220:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st221
		case 98:
			goto st242
		case 99:
			goto st244
		case 100:
			goto st246
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st221
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st221
			}
		default:
			goto st221
		}
		goto st197
	st221:
		if p++; p == pe {
			goto _test_eof221
		}
	st_case_221:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st241
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st222
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st241
				}
			case data[p] >= 65:
				goto st241
			}
		default:
			goto st241
		}
		goto st197
	st222:
		if p++; p == pe {
			goto _test_eof222
		}
	st_case_222:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st223
		case 98:
			goto st224
		case 99:
			goto st230
		case 100:
			goto st237
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st222
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st223
				}
			case data[p] >= 65:
				goto st223
			}
		default:
			goto st223
		}
		goto st197
	st223:
		if p++; p == pe {
			goto _test_eof223
		}
	st_case_223:
		switch data[p] {
		case 44:
			goto st94
		case 59:
			goto st212
		case 95:
			goto st223
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st222
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st223
				}
			case data[p] >= 65:
				goto st223
			}
		default:
			goto st223
		}
		goto st197
	st224:
		if p++; p == pe {
			goto _test_eof224
		}
	st_case_224:
		switch data[p] {
		case 44:
			goto st94
		case 59:
			goto st212
		case 95:
			goto st223
		case 97:
			goto st225
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st222
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st223
				}
			case data[p] >= 65:
				goto st223
			}
		default:
			goto st223
		}
		goto st197
	st225:
		if p++; p == pe {
			goto _test_eof225
		}
	st_case_225:
		switch data[p] {
		case 44:
			goto st94
		case 59:
			goto st212
		case 95:
			goto st223
		case 115:
			goto st226
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st222
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st223
				}
			case data[p] >= 65:
				goto st223
			}
		default:
			goto st223
		}
		goto st197
	st226:
		if p++; p == pe {
			goto _test_eof226
		}
	st_case_226:
		switch data[p] {
		case 44:
			goto st94
		case 59:
			goto st212
		case 95:
			goto st223
		case 101:
			goto st227
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st222
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st223
				}
			case data[p] >= 65:
				goto st223
			}
		default:
			goto st223
		}
		goto st197
	st227:
		if p++; p == pe {
			goto _test_eof227
		}
	st_case_227:
		switch data[p] {
		case 44:
			goto st94
		case 54:
			goto st228
		case 59:
			goto st212
		case 95:
			goto st223
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st222
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st223
				}
			case data[p] >= 65:
				goto st223
			}
		default:
			goto st223
		}
		goto st197
	st228:
		if p++; p == pe {
			goto _test_eof228
		}
	st_case_228:
		switch data[p] {
		case 44:
			goto st94
		case 52:
			goto st229
		case 59:
			goto st212
		case 95:
			goto st223
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st222
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st223
				}
			case data[p] >= 65:
				goto st223
			}
		default:
			goto st223
		}
		goto st197
	st229:
		if p++; p == pe {
			goto _test_eof229
		}
	st_case_229:
		switch data[p] {
		case 44:
			goto tr137
		case 59:
			goto tr257
		case 95:
			goto st223
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto tr256
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st223
				}
			case data[p] >= 65:
				goto st223
			}
		default:
			goto st223
		}
		goto tr224
tr256:
//line xss_170.rl:13

            return true
        
	goto st300
	st300:
		if p++; p == pe {
			goto _test_eof300
		}
	st_case_300:
//line xss_170.go:7886
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st223
		case 98:
			goto st224
		case 99:
			goto st230
		case 100:
			goto st237
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st222
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st223
				}
			case data[p] >= 65:
				goto st223
			}
		default:
			goto st223
		}
		goto st197
	st230:
		if p++; p == pe {
			goto _test_eof230
		}
	st_case_230:
		switch data[p] {
		case 44:
			goto st94
		case 59:
			goto st212
		case 95:
			goto st223
		case 104:
			goto st231
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st222
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st223
				}
			case data[p] >= 65:
				goto st223
			}
		default:
			goto st223
		}
		goto st197
	st231:
		if p++; p == pe {
			goto _test_eof231
		}
	st_case_231:
		switch data[p] {
		case 44:
			goto st94
		case 59:
			goto st212
		case 95:
			goto st223
		case 97:
			goto st232
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st222
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st223
				}
			case data[p] >= 65:
				goto st223
			}
		default:
			goto st223
		}
		goto st197
	st232:
		if p++; p == pe {
			goto _test_eof232
		}
	st_case_232:
		switch data[p] {
		case 44:
			goto st94
		case 59:
			goto st212
		case 95:
			goto st223
		case 114:
			goto st233
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st222
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st223
				}
			case data[p] >= 65:
				goto st223
			}
		default:
			goto st223
		}
		goto st197
	st233:
		if p++; p == pe {
			goto _test_eof233
		}
	st_case_233:
		switch data[p] {
		case 44:
			goto st94
		case 59:
			goto st212
		case 95:
			goto st223
		case 115:
			goto st234
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st222
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st223
				}
			case data[p] >= 65:
				goto st223
			}
		default:
			goto st223
		}
		goto st197
	st234:
		if p++; p == pe {
			goto _test_eof234
		}
	st_case_234:
		switch data[p] {
		case 44:
			goto st94
		case 59:
			goto st212
		case 95:
			goto st223
		case 101:
			goto st235
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st222
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st223
				}
			case data[p] >= 65:
				goto st223
			}
		default:
			goto st223
		}
		goto st197
	st235:
		if p++; p == pe {
			goto _test_eof235
		}
	st_case_235:
		switch data[p] {
		case 44:
			goto st94
		case 59:
			goto st212
		case 95:
			goto st223
		case 116:
			goto st236
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st222
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st223
				}
			case data[p] >= 65:
				goto st223
			}
		default:
			goto st223
		}
		goto st197
	st236:
		if p++; p == pe {
			goto _test_eof236
		}
	st_case_236:
		switch data[p] {
		case 44:
			goto st94
		case 59:
			goto st212
		case 61:
			goto st212
		case 95:
			goto st223
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st222
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st223
				}
			case data[p] >= 65:
				goto st223
			}
		default:
			goto st223
		}
		goto st197
	st237:
		if p++; p == pe {
			goto _test_eof237
		}
	st_case_237:
		switch data[p] {
		case 44:
			goto st94
		case 59:
			goto st212
		case 95:
			goto st223
		case 97:
			goto st238
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st222
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st223
				}
			case data[p] >= 65:
				goto st223
			}
		default:
			goto st223
		}
		goto st197
	st238:
		if p++; p == pe {
			goto _test_eof238
		}
	st_case_238:
		switch data[p] {
		case 44:
			goto st94
		case 59:
			goto st212
		case 95:
			goto st223
		case 116:
			goto st239
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st222
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st223
				}
			case data[p] >= 65:
				goto st223
			}
		default:
			goto st223
		}
		goto st197
	st239:
		if p++; p == pe {
			goto _test_eof239
		}
	st_case_239:
		switch data[p] {
		case 44:
			goto st94
		case 59:
			goto st212
		case 95:
			goto st223
		case 97:
			goto st240
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st222
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st223
				}
			case data[p] >= 65:
				goto st223
			}
		default:
			goto st223
		}
		goto st197
	st240:
		if p++; p == pe {
			goto _test_eof240
		}
	st_case_240:
		switch data[p] {
		case 44:
			goto st94
		case 58:
			goto st217
		case 59:
			goto st212
		case 95:
			goto st223
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st222
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st223
				}
			case data[p] >= 65:
				goto st223
			}
		default:
			goto st223
		}
		goto st197
tr257:
//line xss_170.rl:13

            return true
        
	goto st301
	st301:
		if p++; p == pe {
			goto _test_eof301
		}
	st_case_301:
//line xss_170.go:8291
		switch data[p] {
		case 44:
			goto tr102
		case 95:
			goto st198
		case 98:
			goto st199
		case 99:
			goto st205
		case 100:
			goto st213
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st198
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st198
			}
		default:
			goto st198
		}
		goto tr224
	st241:
		if p++; p == pe {
			goto _test_eof241
		}
	st_case_241:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st223
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st222
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st223
				}
			case data[p] >= 65:
				goto st223
			}
		default:
			goto st223
		}
		goto st197
	st242:
		if p++; p == pe {
			goto _test_eof242
		}
	st_case_242:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st241
		case 97:
			goto st243
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st222
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st241
				}
			case data[p] >= 65:
				goto st241
			}
		default:
			goto st241
		}
		goto st197
	st243:
		if p++; p == pe {
			goto _test_eof243
		}
	st_case_243:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st223
		case 115:
			goto st226
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st222
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st223
				}
			case data[p] >= 65:
				goto st223
			}
		default:
			goto st223
		}
		goto st197
	st244:
		if p++; p == pe {
			goto _test_eof244
		}
	st_case_244:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st241
		case 104:
			goto st245
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st222
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st241
				}
			case data[p] >= 65:
				goto st241
			}
		default:
			goto st241
		}
		goto st197
	st245:
		if p++; p == pe {
			goto _test_eof245
		}
	st_case_245:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st223
		case 97:
			goto st232
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st222
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st223
				}
			case data[p] >= 65:
				goto st223
			}
		default:
			goto st223
		}
		goto st197
	st246:
		if p++; p == pe {
			goto _test_eof246
		}
	st_case_246:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st241
		case 97:
			goto st247
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st222
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st241
				}
			case data[p] >= 65:
				goto st241
			}
		default:
			goto st241
		}
		goto st197
	st247:
		if p++; p == pe {
			goto _test_eof247
		}
	st_case_247:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st223
		case 116:
			goto st239
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st222
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st223
				}
			case data[p] >= 65:
				goto st223
			}
		default:
			goto st223
		}
		goto st197
	st248:
		if p++; p == pe {
			goto _test_eof248
		}
	st_case_248:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st219
		case 97:
			goto st249
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st219
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st219
			}
		default:
			goto st198
		}
		goto st197
	st249:
		if p++; p == pe {
			goto _test_eof249
		}
	st_case_249:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st220
		case 95:
			goto st219
		case 115:
			goto st250
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st219
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st219
			}
		default:
			goto st198
		}
		goto st197
	st250:
		if p++; p == pe {
			goto _test_eof250
		}
	st_case_250:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st220
		case 95:
			goto st219
		case 101:
			goto st251
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st219
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st219
			}
		default:
			goto st198
		}
		goto st197
	st251:
		if p++; p == pe {
			goto _test_eof251
		}
	st_case_251:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st220
		case 54:
			goto st252
		case 95:
			goto st219
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st219
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st219
			}
		default:
			goto st198
		}
		goto st197
	st252:
		if p++; p == pe {
			goto _test_eof252
		}
	st_case_252:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st220
		case 52:
			goto st253
		case 95:
			goto st219
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st219
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st219
			}
		default:
			goto st198
		}
		goto st197
	st253:
		if p++; p == pe {
			goto _test_eof253
		}
	st_case_253:
		switch data[p] {
		case 44:
			goto tr102
		case 47:
			goto tr275
		case 95:
			goto st219
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st219
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st219
			}
		default:
			goto st198
		}
		goto tr224
tr275:
//line xss_170.rl:13

            return true
        
	goto st302
	st302:
		if p++; p == pe {
			goto _test_eof302
		}
	st_case_302:
//line xss_170.go:8707
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st221
		case 98:
			goto st242
		case 99:
			goto st244
		case 100:
			goto st246
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st221
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st221
			}
		default:
			goto st221
		}
		goto st197
	st254:
		if p++; p == pe {
			goto _test_eof254
		}
	st_case_254:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st219
		case 104:
			goto st255
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st219
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st219
			}
		default:
			goto st198
		}
		goto st197
	st255:
		if p++; p == pe {
			goto _test_eof255
		}
	st_case_255:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st220
		case 95:
			goto st219
		case 97:
			goto st256
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st219
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st219
			}
		default:
			goto st198
		}
		goto st197
	st256:
		if p++; p == pe {
			goto _test_eof256
		}
	st_case_256:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st220
		case 95:
			goto st219
		case 114:
			goto st257
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st219
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st219
			}
		default:
			goto st198
		}
		goto st197
	st257:
		if p++; p == pe {
			goto _test_eof257
		}
	st_case_257:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st220
		case 95:
			goto st219
		case 115:
			goto st258
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st219
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st219
			}
		default:
			goto st198
		}
		goto st197
	st258:
		if p++; p == pe {
			goto _test_eof258
		}
	st_case_258:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st220
		case 95:
			goto st219
		case 101:
			goto st259
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st219
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st219
			}
		default:
			goto st198
		}
		goto st197
	st259:
		if p++; p == pe {
			goto _test_eof259
		}
	st_case_259:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st220
		case 95:
			goto st219
		case 116:
			goto st260
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st219
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st219
			}
		default:
			goto st198
		}
		goto st197
	st260:
		if p++; p == pe {
			goto _test_eof260
		}
	st_case_260:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st220
		case 61:
			goto st212
		case 95:
			goto st219
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st219
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st219
			}
		default:
			goto st198
		}
		goto st197
	st261:
		if p++; p == pe {
			goto _test_eof261
		}
	st_case_261:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st219
		case 97:
			goto st262
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st219
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st219
			}
		default:
			goto st198
		}
		goto st197
	st262:
		if p++; p == pe {
			goto _test_eof262
		}
	st_case_262:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st220
		case 95:
			goto st219
		case 116:
			goto st263
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st219
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st219
			}
		default:
			goto st198
		}
		goto st197
	st263:
		if p++; p == pe {
			goto _test_eof263
		}
	st_case_263:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st220
		case 95:
			goto st219
		case 97:
			goto st264
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st219
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st219
			}
		default:
			goto st198
		}
		goto st197
	st264:
		if p++; p == pe {
			goto _test_eof264
		}
	st_case_264:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st220
		case 58:
			goto st217
		case 95:
			goto st219
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st219
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st219
			}
		default:
			goto st198
		}
		goto st197
	st265:
		if p++; p == pe {
			goto _test_eof265
		}
	st_case_265:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st197
		case 95:
			goto st196
		case 97:
			goto st266
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st196
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st196
			}
		default:
			goto st196
		}
		goto st7
	st266:
		if p++; p == pe {
			goto _test_eof266
		}
	st_case_266:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st197
		case 95:
			goto st196
		case 116:
			goto st267
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st196
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st196
			}
		default:
			goto st196
		}
		goto st7
	st267:
		if p++; p == pe {
			goto _test_eof267
		}
	st_case_267:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st197
		case 95:
			goto st196
		case 97:
			goto st268
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st196
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st196
			}
		default:
			goto st196
		}
		goto st7
	st268:
		if p++; p == pe {
			goto _test_eof268
		}
	st_case_268:
		switch data[p] {
		case 44:
			goto st8
		case 58:
			goto st6
		case 59:
			goto st197
		case 95:
			goto st196
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st196
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st196
			}
		default:
			goto st196
		}
		goto st7
	st269:
		if p++; p == pe {
			goto _test_eof269
		}
	st_case_269:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st197
		case 95:
			goto st270
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st270
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st270
			}
		default:
			goto st196
		}
		goto st7
	st270:
		if p++; p == pe {
			goto _test_eof270
		}
	st_case_270:
		switch data[p] {
		case 44:
			goto st8
		case 47:
			goto st271
		case 59:
			goto st197
		case 95:
			goto st270
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st270
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st270
			}
		default:
			goto st196
		}
		goto st7
	st271:
		if p++; p == pe {
			goto _test_eof271
		}
	st_case_271:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st197
		case 95:
			goto st272
		case 100:
			goto st280
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st272
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st272
			}
		default:
			goto st272
		}
		goto st7
	st272:
		if p++; p == pe {
			goto _test_eof272
		}
	st_case_272:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st197
		case 95:
			goto st279
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st273
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st279
				}
			case data[p] >= 65:
				goto st279
			}
		default:
			goto st279
		}
		goto st7
	st273:
		if p++; p == pe {
			goto _test_eof273
		}
	st_case_273:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st197
		case 95:
			goto st274
		case 100:
			goto st275
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st273
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st274
				}
			case data[p] >= 65:
				goto st274
			}
		default:
			goto st274
		}
		goto st7
	st274:
		if p++; p == pe {
			goto _test_eof274
		}
	st_case_274:
		switch data[p] {
		case 44:
			goto st178
		case 59:
			goto st212
		case 95:
			goto st274
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st273
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st274
				}
			case data[p] >= 65:
				goto st274
			}
		default:
			goto st274
		}
		goto st7
	st275:
		if p++; p == pe {
			goto _test_eof275
		}
	st_case_275:
		switch data[p] {
		case 44:
			goto st178
		case 59:
			goto st212
		case 95:
			goto st274
		case 97:
			goto st276
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st273
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st274
				}
			case data[p] >= 65:
				goto st274
			}
		default:
			goto st274
		}
		goto st7
	st276:
		if p++; p == pe {
			goto _test_eof276
		}
	st_case_276:
		switch data[p] {
		case 44:
			goto st178
		case 59:
			goto st212
		case 95:
			goto st274
		case 116:
			goto st277
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st273
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st274
				}
			case data[p] >= 65:
				goto st274
			}
		default:
			goto st274
		}
		goto st7
	st277:
		if p++; p == pe {
			goto _test_eof277
		}
	st_case_277:
		switch data[p] {
		case 44:
			goto st178
		case 59:
			goto st212
		case 95:
			goto st274
		case 97:
			goto st278
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st273
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st274
				}
			case data[p] >= 65:
				goto st274
			}
		default:
			goto st274
		}
		goto st7
	st278:
		if p++; p == pe {
			goto _test_eof278
		}
	st_case_278:
		switch data[p] {
		case 44:
			goto st178
		case 58:
			goto st6
		case 59:
			goto st212
		case 95:
			goto st274
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st273
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st274
				}
			case data[p] >= 65:
				goto st274
			}
		default:
			goto st274
		}
		goto st7
	st279:
		if p++; p == pe {
			goto _test_eof279
		}
	st_case_279:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st197
		case 95:
			goto st274
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st273
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st274
				}
			case data[p] >= 65:
				goto st274
			}
		default:
			goto st274
		}
		goto st7
	st280:
		if p++; p == pe {
			goto _test_eof280
		}
	st_case_280:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st197
		case 95:
			goto st279
		case 97:
			goto st281
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st273
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st279
				}
			case data[p] >= 65:
				goto st279
			}
		default:
			goto st279
		}
		goto st7
	st281:
		if p++; p == pe {
			goto _test_eof281
		}
	st_case_281:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st197
		case 95:
			goto st274
		case 116:
			goto st277
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st273
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st274
				}
			case data[p] >= 65:
				goto st274
			}
		default:
			goto st274
		}
		goto st7
	st282:
		if p++; p == pe {
			goto _test_eof282
		}
	st_case_282:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st197
		case 95:
			goto st270
		case 97:
			goto st283
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st270
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st270
			}
		default:
			goto st196
		}
		goto st7
	st283:
		if p++; p == pe {
			goto _test_eof283
		}
	st_case_283:
		switch data[p] {
		case 44:
			goto st8
		case 47:
			goto st271
		case 59:
			goto st197
		case 95:
			goto st270
		case 116:
			goto st284
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st270
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st270
			}
		default:
			goto st196
		}
		goto st7
	st284:
		if p++; p == pe {
			goto _test_eof284
		}
	st_case_284:
		switch data[p] {
		case 44:
			goto st8
		case 47:
			goto st271
		case 59:
			goto st197
		case 95:
			goto st270
		case 97:
			goto st285
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st270
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st270
			}
		default:
			goto st196
		}
		goto st7
	st285:
		if p++; p == pe {
			goto _test_eof285
		}
	st_case_285:
		switch data[p] {
		case 44:
			goto st8
		case 47:
			goto st271
		case 58:
			goto st6
		case 59:
			goto st197
		case 95:
			goto st270
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st270
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st270
			}
		default:
			goto st196
		}
		goto st7
	st_out:
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
	_test_eof286: cs = 286; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof287: cs = 287; goto _test_eof
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
	_test_eof288: cs = 288; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
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
	_test_eof65: cs = 65; goto _test_eof
	_test_eof66: cs = 66; goto _test_eof
	_test_eof67: cs = 67; goto _test_eof
	_test_eof68: cs = 68; goto _test_eof
	_test_eof69: cs = 69; goto _test_eof
	_test_eof289: cs = 289; goto _test_eof
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
	_test_eof290: cs = 290; goto _test_eof
	_test_eof87: cs = 87; goto _test_eof
	_test_eof88: cs = 88; goto _test_eof
	_test_eof89: cs = 89; goto _test_eof
	_test_eof90: cs = 90; goto _test_eof
	_test_eof91: cs = 91; goto _test_eof
	_test_eof92: cs = 92; goto _test_eof
	_test_eof93: cs = 93; goto _test_eof
	_test_eof94: cs = 94; goto _test_eof
	_test_eof291: cs = 291; goto _test_eof
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
	_test_eof292: cs = 292; goto _test_eof
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
	_test_eof293: cs = 293; goto _test_eof
	_test_eof123: cs = 123; goto _test_eof
	_test_eof124: cs = 124; goto _test_eof
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
	_test_eof294: cs = 294; goto _test_eof
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
	_test_eof295: cs = 295; goto _test_eof
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
	_test_eof296: cs = 296; goto _test_eof
	_test_eof173: cs = 173; goto _test_eof
	_test_eof174: cs = 174; goto _test_eof
	_test_eof175: cs = 175; goto _test_eof
	_test_eof176: cs = 176; goto _test_eof
	_test_eof177: cs = 177; goto _test_eof
	_test_eof178: cs = 178; goto _test_eof
	_test_eof297: cs = 297; goto _test_eof
	_test_eof298: cs = 298; goto _test_eof
	_test_eof179: cs = 179; goto _test_eof
	_test_eof180: cs = 180; goto _test_eof
	_test_eof181: cs = 181; goto _test_eof
	_test_eof182: cs = 182; goto _test_eof
	_test_eof183: cs = 183; goto _test_eof
	_test_eof184: cs = 184; goto _test_eof
	_test_eof185: cs = 185; goto _test_eof
	_test_eof186: cs = 186; goto _test_eof
	_test_eof187: cs = 187; goto _test_eof
	_test_eof188: cs = 188; goto _test_eof
	_test_eof189: cs = 189; goto _test_eof
	_test_eof190: cs = 190; goto _test_eof
	_test_eof191: cs = 191; goto _test_eof
	_test_eof192: cs = 192; goto _test_eof
	_test_eof193: cs = 193; goto _test_eof
	_test_eof194: cs = 194; goto _test_eof
	_test_eof195: cs = 195; goto _test_eof
	_test_eof196: cs = 196; goto _test_eof
	_test_eof197: cs = 197; goto _test_eof
	_test_eof198: cs = 198; goto _test_eof
	_test_eof199: cs = 199; goto _test_eof
	_test_eof200: cs = 200; goto _test_eof
	_test_eof201: cs = 201; goto _test_eof
	_test_eof202: cs = 202; goto _test_eof
	_test_eof203: cs = 203; goto _test_eof
	_test_eof204: cs = 204; goto _test_eof
	_test_eof299: cs = 299; goto _test_eof
	_test_eof205: cs = 205; goto _test_eof
	_test_eof206: cs = 206; goto _test_eof
	_test_eof207: cs = 207; goto _test_eof
	_test_eof208: cs = 208; goto _test_eof
	_test_eof209: cs = 209; goto _test_eof
	_test_eof210: cs = 210; goto _test_eof
	_test_eof211: cs = 211; goto _test_eof
	_test_eof212: cs = 212; goto _test_eof
	_test_eof213: cs = 213; goto _test_eof
	_test_eof214: cs = 214; goto _test_eof
	_test_eof215: cs = 215; goto _test_eof
	_test_eof216: cs = 216; goto _test_eof
	_test_eof217: cs = 217; goto _test_eof
	_test_eof218: cs = 218; goto _test_eof
	_test_eof219: cs = 219; goto _test_eof
	_test_eof220: cs = 220; goto _test_eof
	_test_eof221: cs = 221; goto _test_eof
	_test_eof222: cs = 222; goto _test_eof
	_test_eof223: cs = 223; goto _test_eof
	_test_eof224: cs = 224; goto _test_eof
	_test_eof225: cs = 225; goto _test_eof
	_test_eof226: cs = 226; goto _test_eof
	_test_eof227: cs = 227; goto _test_eof
	_test_eof228: cs = 228; goto _test_eof
	_test_eof229: cs = 229; goto _test_eof
	_test_eof300: cs = 300; goto _test_eof
	_test_eof230: cs = 230; goto _test_eof
	_test_eof231: cs = 231; goto _test_eof
	_test_eof232: cs = 232; goto _test_eof
	_test_eof233: cs = 233; goto _test_eof
	_test_eof234: cs = 234; goto _test_eof
	_test_eof235: cs = 235; goto _test_eof
	_test_eof236: cs = 236; goto _test_eof
	_test_eof237: cs = 237; goto _test_eof
	_test_eof238: cs = 238; goto _test_eof
	_test_eof239: cs = 239; goto _test_eof
	_test_eof240: cs = 240; goto _test_eof
	_test_eof301: cs = 301; goto _test_eof
	_test_eof241: cs = 241; goto _test_eof
	_test_eof242: cs = 242; goto _test_eof
	_test_eof243: cs = 243; goto _test_eof
	_test_eof244: cs = 244; goto _test_eof
	_test_eof245: cs = 245; goto _test_eof
	_test_eof246: cs = 246; goto _test_eof
	_test_eof247: cs = 247; goto _test_eof
	_test_eof248: cs = 248; goto _test_eof
	_test_eof249: cs = 249; goto _test_eof
	_test_eof250: cs = 250; goto _test_eof
	_test_eof251: cs = 251; goto _test_eof
	_test_eof252: cs = 252; goto _test_eof
	_test_eof253: cs = 253; goto _test_eof
	_test_eof302: cs = 302; goto _test_eof
	_test_eof254: cs = 254; goto _test_eof
	_test_eof255: cs = 255; goto _test_eof
	_test_eof256: cs = 256; goto _test_eof
	_test_eof257: cs = 257; goto _test_eof
	_test_eof258: cs = 258; goto _test_eof
	_test_eof259: cs = 259; goto _test_eof
	_test_eof260: cs = 260; goto _test_eof
	_test_eof261: cs = 261; goto _test_eof
	_test_eof262: cs = 262; goto _test_eof
	_test_eof263: cs = 263; goto _test_eof
	_test_eof264: cs = 264; goto _test_eof
	_test_eof265: cs = 265; goto _test_eof
	_test_eof266: cs = 266; goto _test_eof
	_test_eof267: cs = 267; goto _test_eof
	_test_eof268: cs = 268; goto _test_eof
	_test_eof269: cs = 269; goto _test_eof
	_test_eof270: cs = 270; goto _test_eof
	_test_eof271: cs = 271; goto _test_eof
	_test_eof272: cs = 272; goto _test_eof
	_test_eof273: cs = 273; goto _test_eof
	_test_eof274: cs = 274; goto _test_eof
	_test_eof275: cs = 275; goto _test_eof
	_test_eof276: cs = 276; goto _test_eof
	_test_eof277: cs = 277; goto _test_eof
	_test_eof278: cs = 278; goto _test_eof
	_test_eof279: cs = 279; goto _test_eof
	_test_eof280: cs = 280; goto _test_eof
	_test_eof281: cs = 281; goto _test_eof
	_test_eof282: cs = 282; goto _test_eof
	_test_eof283: cs = 283; goto _test_eof
	_test_eof284: cs = 284; goto _test_eof
	_test_eof285: cs = 285; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 15, 21, 45, 69, 86, 94, 111, 135, 150, 178, 204, 212, 229, 253, 287, 293, 296, 301:
//line xss_170.rl:13

            return true
        
//line xss_170.go:9986
		}
	}

	}

//line xss_170.rl:34


    return false
}
