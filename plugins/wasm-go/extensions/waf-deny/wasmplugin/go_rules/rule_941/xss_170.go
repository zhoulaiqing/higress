
//line xss_170.rl:1
package rule_941


func matchXSS170(data []byte) bool {

//line xss_170.rl:6

//line xss_170.go:11
const xss_start int = 0
const xss_first_final int = 846
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
	case 846:
		goto st_case_846
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
	case 847:
		goto st_case_847
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
	case 848:
		goto st_case_848
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
	case 849:
		goto st_case_849
	case 70:
		goto st_case_70
	case 71:
		goto st_case_71
	case 850:
		goto st_case_850
	case 851:
		goto st_case_851
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
	case 852:
		goto st_case_852
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
	case 853:
		goto st_case_853
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
	case 854:
		goto st_case_854
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
	case 183:
		goto st_case_183
	case 184:
		goto st_case_184
	case 855:
		goto st_case_855
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
	case 856:
		goto st_case_856
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
	case 857:
		goto st_case_857
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
	case 230:
		goto st_case_230
	case 231:
		goto st_case_231
	case 232:
		goto st_case_232
	case 233:
		goto st_case_233
	case 858:
		goto st_case_858
	case 234:
		goto st_case_234
	case 235:
		goto st_case_235
	case 859:
		goto st_case_859
	case 236:
		goto st_case_236
	case 237:
		goto st_case_237
	case 860:
		goto st_case_860
	case 861:
		goto st_case_861
	case 238:
		goto st_case_238
	case 862:
		goto st_case_862
	case 863:
		goto st_case_863
	case 239:
		goto st_case_239
	case 240:
		goto st_case_240
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
	case 864:
		goto st_case_864
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
	case 865:
		goto st_case_865
	case 284:
		goto st_case_284
	case 285:
		goto st_case_285
	case 286:
		goto st_case_286
	case 287:
		goto st_case_287
	case 288:
		goto st_case_288
	case 289:
		goto st_case_289
	case 290:
		goto st_case_290
	case 291:
		goto st_case_291
	case 292:
		goto st_case_292
	case 293:
		goto st_case_293
	case 294:
		goto st_case_294
	case 295:
		goto st_case_295
	case 296:
		goto st_case_296
	case 297:
		goto st_case_297
	case 298:
		goto st_case_298
	case 299:
		goto st_case_299
	case 866:
		goto st_case_866
	case 300:
		goto st_case_300
	case 301:
		goto st_case_301
	case 302:
		goto st_case_302
	case 303:
		goto st_case_303
	case 304:
		goto st_case_304
	case 305:
		goto st_case_305
	case 306:
		goto st_case_306
	case 307:
		goto st_case_307
	case 308:
		goto st_case_308
	case 309:
		goto st_case_309
	case 310:
		goto st_case_310
	case 311:
		goto st_case_311
	case 312:
		goto st_case_312
	case 313:
		goto st_case_313
	case 314:
		goto st_case_314
	case 315:
		goto st_case_315
	case 316:
		goto st_case_316
	case 317:
		goto st_case_317
	case 318:
		goto st_case_318
	case 319:
		goto st_case_319
	case 320:
		goto st_case_320
	case 321:
		goto st_case_321
	case 322:
		goto st_case_322
	case 323:
		goto st_case_323
	case 324:
		goto st_case_324
	case 867:
		goto st_case_867
	case 325:
		goto st_case_325
	case 326:
		goto st_case_326
	case 327:
		goto st_case_327
	case 328:
		goto st_case_328
	case 329:
		goto st_case_329
	case 330:
		goto st_case_330
	case 331:
		goto st_case_331
	case 332:
		goto st_case_332
	case 333:
		goto st_case_333
	case 334:
		goto st_case_334
	case 335:
		goto st_case_335
	case 336:
		goto st_case_336
	case 337:
		goto st_case_337
	case 338:
		goto st_case_338
	case 339:
		goto st_case_339
	case 340:
		goto st_case_340
	case 341:
		goto st_case_341
	case 342:
		goto st_case_342
	case 343:
		goto st_case_343
	case 344:
		goto st_case_344
	case 345:
		goto st_case_345
	case 346:
		goto st_case_346
	case 347:
		goto st_case_347
	case 348:
		goto st_case_348
	case 349:
		goto st_case_349
	case 868:
		goto st_case_868
	case 350:
		goto st_case_350
	case 351:
		goto st_case_351
	case 352:
		goto st_case_352
	case 353:
		goto st_case_353
	case 354:
		goto st_case_354
	case 355:
		goto st_case_355
	case 356:
		goto st_case_356
	case 357:
		goto st_case_357
	case 358:
		goto st_case_358
	case 359:
		goto st_case_359
	case 360:
		goto st_case_360
	case 361:
		goto st_case_361
	case 362:
		goto st_case_362
	case 363:
		goto st_case_363
	case 364:
		goto st_case_364
	case 365:
		goto st_case_365
	case 366:
		goto st_case_366
	case 367:
		goto st_case_367
	case 368:
		goto st_case_368
	case 369:
		goto st_case_369
	case 370:
		goto st_case_370
	case 371:
		goto st_case_371
	case 372:
		goto st_case_372
	case 373:
		goto st_case_373
	case 374:
		goto st_case_374
	case 375:
		goto st_case_375
	case 376:
		goto st_case_376
	case 377:
		goto st_case_377
	case 869:
		goto st_case_869
	case 378:
		goto st_case_378
	case 379:
		goto st_case_379
	case 870:
		goto st_case_870
	case 871:
		goto st_case_871
	case 380:
		goto st_case_380
	case 381:
		goto st_case_381
	case 382:
		goto st_case_382
	case 383:
		goto st_case_383
	case 384:
		goto st_case_384
	case 385:
		goto st_case_385
	case 386:
		goto st_case_386
	case 387:
		goto st_case_387
	case 388:
		goto st_case_388
	case 389:
		goto st_case_389
	case 390:
		goto st_case_390
	case 391:
		goto st_case_391
	case 392:
		goto st_case_392
	case 393:
		goto st_case_393
	case 394:
		goto st_case_394
	case 395:
		goto st_case_395
	case 396:
		goto st_case_396
	case 397:
		goto st_case_397
	case 398:
		goto st_case_398
	case 399:
		goto st_case_399
	case 872:
		goto st_case_872
	case 400:
		goto st_case_400
	case 401:
		goto st_case_401
	case 402:
		goto st_case_402
	case 403:
		goto st_case_403
	case 404:
		goto st_case_404
	case 405:
		goto st_case_405
	case 406:
		goto st_case_406
	case 407:
		goto st_case_407
	case 408:
		goto st_case_408
	case 409:
		goto st_case_409
	case 410:
		goto st_case_410
	case 411:
		goto st_case_411
	case 412:
		goto st_case_412
	case 413:
		goto st_case_413
	case 414:
		goto st_case_414
	case 415:
		goto st_case_415
	case 416:
		goto st_case_416
	case 417:
		goto st_case_417
	case 418:
		goto st_case_418
	case 419:
		goto st_case_419
	case 420:
		goto st_case_420
	case 421:
		goto st_case_421
	case 422:
		goto st_case_422
	case 423:
		goto st_case_423
	case 424:
		goto st_case_424
	case 425:
		goto st_case_425
	case 426:
		goto st_case_426
	case 427:
		goto st_case_427
	case 428:
		goto st_case_428
	case 429:
		goto st_case_429
	case 430:
		goto st_case_430
	case 431:
		goto st_case_431
	case 432:
		goto st_case_432
	case 433:
		goto st_case_433
	case 434:
		goto st_case_434
	case 435:
		goto st_case_435
	case 436:
		goto st_case_436
	case 437:
		goto st_case_437
	case 438:
		goto st_case_438
	case 439:
		goto st_case_439
	case 440:
		goto st_case_440
	case 441:
		goto st_case_441
	case 442:
		goto st_case_442
	case 443:
		goto st_case_443
	case 444:
		goto st_case_444
	case 445:
		goto st_case_445
	case 446:
		goto st_case_446
	case 447:
		goto st_case_447
	case 448:
		goto st_case_448
	case 449:
		goto st_case_449
	case 873:
		goto st_case_873
	case 450:
		goto st_case_450
	case 451:
		goto st_case_451
	case 452:
		goto st_case_452
	case 453:
		goto st_case_453
	case 454:
		goto st_case_454
	case 455:
		goto st_case_455
	case 874:
		goto st_case_874
	case 456:
		goto st_case_456
	case 457:
		goto st_case_457
	case 458:
		goto st_case_458
	case 459:
		goto st_case_459
	case 460:
		goto st_case_460
	case 461:
		goto st_case_461
	case 462:
		goto st_case_462
	case 463:
		goto st_case_463
	case 464:
		goto st_case_464
	case 465:
		goto st_case_465
	case 466:
		goto st_case_466
	case 467:
		goto st_case_467
	case 468:
		goto st_case_468
	case 875:
		goto st_case_875
	case 469:
		goto st_case_469
	case 470:
		goto st_case_470
	case 876:
		goto st_case_876
	case 471:
		goto st_case_471
	case 472:
		goto st_case_472
	case 877:
		goto st_case_877
	case 878:
		goto st_case_878
	case 473:
		goto st_case_473
	case 879:
		goto st_case_879
	case 880:
		goto st_case_880
	case 474:
		goto st_case_474
	case 475:
		goto st_case_475
	case 476:
		goto st_case_476
	case 477:
		goto st_case_477
	case 478:
		goto st_case_478
	case 479:
		goto st_case_479
	case 480:
		goto st_case_480
	case 481:
		goto st_case_481
	case 482:
		goto st_case_482
	case 483:
		goto st_case_483
	case 484:
		goto st_case_484
	case 485:
		goto st_case_485
	case 486:
		goto st_case_486
	case 487:
		goto st_case_487
	case 488:
		goto st_case_488
	case 489:
		goto st_case_489
	case 490:
		goto st_case_490
	case 491:
		goto st_case_491
	case 492:
		goto st_case_492
	case 493:
		goto st_case_493
	case 881:
		goto st_case_881
	case 494:
		goto st_case_494
	case 495:
		goto st_case_495
	case 496:
		goto st_case_496
	case 497:
		goto st_case_497
	case 498:
		goto st_case_498
	case 499:
		goto st_case_499
	case 500:
		goto st_case_500
	case 501:
		goto st_case_501
	case 502:
		goto st_case_502
	case 503:
		goto st_case_503
	case 504:
		goto st_case_504
	case 505:
		goto st_case_505
	case 506:
		goto st_case_506
	case 507:
		goto st_case_507
	case 508:
		goto st_case_508
	case 509:
		goto st_case_509
	case 510:
		goto st_case_510
	case 511:
		goto st_case_511
	case 512:
		goto st_case_512
	case 513:
		goto st_case_513
	case 514:
		goto st_case_514
	case 882:
		goto st_case_882
	case 515:
		goto st_case_515
	case 516:
		goto st_case_516
	case 517:
		goto st_case_517
	case 518:
		goto st_case_518
	case 519:
		goto st_case_519
	case 520:
		goto st_case_520
	case 521:
		goto st_case_521
	case 522:
		goto st_case_522
	case 523:
		goto st_case_523
	case 524:
		goto st_case_524
	case 525:
		goto st_case_525
	case 526:
		goto st_case_526
	case 527:
		goto st_case_527
	case 528:
		goto st_case_528
	case 529:
		goto st_case_529
	case 530:
		goto st_case_530
	case 531:
		goto st_case_531
	case 532:
		goto st_case_532
	case 533:
		goto st_case_533
	case 534:
		goto st_case_534
	case 535:
		goto st_case_535
	case 536:
		goto st_case_536
	case 537:
		goto st_case_537
	case 538:
		goto st_case_538
	case 539:
		goto st_case_539
	case 540:
		goto st_case_540
	case 541:
		goto st_case_541
	case 542:
		goto st_case_542
	case 543:
		goto st_case_543
	case 544:
		goto st_case_544
	case 545:
		goto st_case_545
	case 546:
		goto st_case_546
	case 547:
		goto st_case_547
	case 548:
		goto st_case_548
	case 549:
		goto st_case_549
	case 550:
		goto st_case_550
	case 551:
		goto st_case_551
	case 552:
		goto st_case_552
	case 553:
		goto st_case_553
	case 554:
		goto st_case_554
	case 555:
		goto st_case_555
	case 556:
		goto st_case_556
	case 557:
		goto st_case_557
	case 558:
		goto st_case_558
	case 559:
		goto st_case_559
	case 560:
		goto st_case_560
	case 561:
		goto st_case_561
	case 562:
		goto st_case_562
	case 883:
		goto st_case_883
	case 563:
		goto st_case_563
	case 564:
		goto st_case_564
	case 565:
		goto st_case_565
	case 566:
		goto st_case_566
	case 567:
		goto st_case_567
	case 568:
		goto st_case_568
	case 569:
		goto st_case_569
	case 570:
		goto st_case_570
	case 571:
		goto st_case_571
	case 572:
		goto st_case_572
	case 573:
		goto st_case_573
	case 574:
		goto st_case_574
	case 575:
		goto st_case_575
	case 576:
		goto st_case_576
	case 577:
		goto st_case_577
	case 578:
		goto st_case_578
	case 579:
		goto st_case_579
	case 580:
		goto st_case_580
	case 581:
		goto st_case_581
	case 582:
		goto st_case_582
	case 583:
		goto st_case_583
	case 584:
		goto st_case_584
	case 585:
		goto st_case_585
	case 586:
		goto st_case_586
	case 587:
		goto st_case_587
	case 884:
		goto st_case_884
	case 588:
		goto st_case_588
	case 589:
		goto st_case_589
	case 590:
		goto st_case_590
	case 591:
		goto st_case_591
	case 592:
		goto st_case_592
	case 593:
		goto st_case_593
	case 594:
		goto st_case_594
	case 595:
		goto st_case_595
	case 596:
		goto st_case_596
	case 597:
		goto st_case_597
	case 598:
		goto st_case_598
	case 599:
		goto st_case_599
	case 600:
		goto st_case_600
	case 601:
		goto st_case_601
	case 602:
		goto st_case_602
	case 603:
		goto st_case_603
	case 604:
		goto st_case_604
	case 605:
		goto st_case_605
	case 606:
		goto st_case_606
	case 607:
		goto st_case_607
	case 608:
		goto st_case_608
	case 609:
		goto st_case_609
	case 610:
		goto st_case_610
	case 611:
		goto st_case_611
	case 885:
		goto st_case_885
	case 612:
		goto st_case_612
	case 613:
		goto st_case_613
	case 886:
		goto st_case_886
	case 887:
		goto st_case_887
	case 614:
		goto st_case_614
	case 615:
		goto st_case_615
	case 616:
		goto st_case_616
	case 617:
		goto st_case_617
	case 618:
		goto st_case_618
	case 619:
		goto st_case_619
	case 620:
		goto st_case_620
	case 621:
		goto st_case_621
	case 622:
		goto st_case_622
	case 623:
		goto st_case_623
	case 624:
		goto st_case_624
	case 625:
		goto st_case_625
	case 626:
		goto st_case_626
	case 627:
		goto st_case_627
	case 628:
		goto st_case_628
	case 629:
		goto st_case_629
	case 630:
		goto st_case_630
	case 631:
		goto st_case_631
	case 632:
		goto st_case_632
	case 633:
		goto st_case_633
	case 634:
		goto st_case_634
	case 635:
		goto st_case_635
	case 636:
		goto st_case_636
	case 888:
		goto st_case_888
	case 637:
		goto st_case_637
	case 638:
		goto st_case_638
	case 639:
		goto st_case_639
	case 640:
		goto st_case_640
	case 641:
		goto st_case_641
	case 642:
		goto st_case_642
	case 643:
		goto st_case_643
	case 644:
		goto st_case_644
	case 645:
		goto st_case_645
	case 646:
		goto st_case_646
	case 647:
		goto st_case_647
	case 648:
		goto st_case_648
	case 649:
		goto st_case_649
	case 650:
		goto st_case_650
	case 651:
		goto st_case_651
	case 652:
		goto st_case_652
	case 653:
		goto st_case_653
	case 654:
		goto st_case_654
	case 655:
		goto st_case_655
	case 656:
		goto st_case_656
	case 657:
		goto st_case_657
	case 658:
		goto st_case_658
	case 889:
		goto st_case_889
	case 659:
		goto st_case_659
	case 660:
		goto st_case_660
	case 661:
		goto st_case_661
	case 662:
		goto st_case_662
	case 663:
		goto st_case_663
	case 664:
		goto st_case_664
	case 665:
		goto st_case_665
	case 666:
		goto st_case_666
	case 667:
		goto st_case_667
	case 668:
		goto st_case_668
	case 669:
		goto st_case_669
	case 670:
		goto st_case_670
	case 671:
		goto st_case_671
	case 672:
		goto st_case_672
	case 673:
		goto st_case_673
	case 674:
		goto st_case_674
	case 890:
		goto st_case_890
	case 675:
		goto st_case_675
	case 676:
		goto st_case_676
	case 677:
		goto st_case_677
	case 678:
		goto st_case_678
	case 679:
		goto st_case_679
	case 680:
		goto st_case_680
	case 681:
		goto st_case_681
	case 682:
		goto st_case_682
	case 683:
		goto st_case_683
	case 684:
		goto st_case_684
	case 685:
		goto st_case_685
	case 686:
		goto st_case_686
	case 687:
		goto st_case_687
	case 688:
		goto st_case_688
	case 689:
		goto st_case_689
	case 690:
		goto st_case_690
	case 691:
		goto st_case_691
	case 692:
		goto st_case_692
	case 693:
		goto st_case_693
	case 694:
		goto st_case_694
	case 695:
		goto st_case_695
	case 696:
		goto st_case_696
	case 697:
		goto st_case_697
	case 698:
		goto st_case_698
	case 699:
		goto st_case_699
	case 891:
		goto st_case_891
	case 700:
		goto st_case_700
	case 701:
		goto st_case_701
	case 702:
		goto st_case_702
	case 703:
		goto st_case_703
	case 704:
		goto st_case_704
	case 705:
		goto st_case_705
	case 706:
		goto st_case_706
	case 707:
		goto st_case_707
	case 708:
		goto st_case_708
	case 709:
		goto st_case_709
	case 710:
		goto st_case_710
	case 711:
		goto st_case_711
	case 712:
		goto st_case_712
	case 713:
		goto st_case_713
	case 714:
		goto st_case_714
	case 715:
		goto st_case_715
	case 716:
		goto st_case_716
	case 717:
		goto st_case_717
	case 718:
		goto st_case_718
	case 719:
		goto st_case_719
	case 720:
		goto st_case_720
	case 721:
		goto st_case_721
	case 722:
		goto st_case_722
	case 723:
		goto st_case_723
	case 724:
		goto st_case_724
	case 725:
		goto st_case_725
	case 726:
		goto st_case_726
	case 727:
		goto st_case_727
	case 728:
		goto st_case_728
	case 729:
		goto st_case_729
	case 730:
		goto st_case_730
	case 731:
		goto st_case_731
	case 732:
		goto st_case_732
	case 733:
		goto st_case_733
	case 734:
		goto st_case_734
	case 735:
		goto st_case_735
	case 736:
		goto st_case_736
	case 737:
		goto st_case_737
	case 892:
		goto st_case_892
	case 738:
		goto st_case_738
	case 739:
		goto st_case_739
	case 893:
		goto st_case_893
	case 894:
		goto st_case_894
	case 740:
		goto st_case_740
	case 741:
		goto st_case_741
	case 742:
		goto st_case_742
	case 743:
		goto st_case_743
	case 744:
		goto st_case_744
	case 745:
		goto st_case_745
	case 746:
		goto st_case_746
	case 747:
		goto st_case_747
	case 748:
		goto st_case_748
	case 749:
		goto st_case_749
	case 750:
		goto st_case_750
	case 751:
		goto st_case_751
	case 752:
		goto st_case_752
	case 753:
		goto st_case_753
	case 754:
		goto st_case_754
	case 755:
		goto st_case_755
	case 756:
		goto st_case_756
	case 757:
		goto st_case_757
	case 758:
		goto st_case_758
	case 759:
		goto st_case_759
	case 895:
		goto st_case_895
	case 760:
		goto st_case_760
	case 761:
		goto st_case_761
	case 762:
		goto st_case_762
	case 763:
		goto st_case_763
	case 764:
		goto st_case_764
	case 765:
		goto st_case_765
	case 766:
		goto st_case_766
	case 767:
		goto st_case_767
	case 768:
		goto st_case_768
	case 769:
		goto st_case_769
	case 770:
		goto st_case_770
	case 771:
		goto st_case_771
	case 772:
		goto st_case_772
	case 773:
		goto st_case_773
	case 774:
		goto st_case_774
	case 775:
		goto st_case_775
	case 776:
		goto st_case_776
	case 777:
		goto st_case_777
	case 778:
		goto st_case_778
	case 779:
		goto st_case_779
	case 780:
		goto st_case_780
	case 781:
		goto st_case_781
	case 782:
		goto st_case_782
	case 783:
		goto st_case_783
	case 784:
		goto st_case_784
	case 785:
		goto st_case_785
	case 786:
		goto st_case_786
	case 787:
		goto st_case_787
	case 788:
		goto st_case_788
	case 789:
		goto st_case_789
	case 790:
		goto st_case_790
	case 791:
		goto st_case_791
	case 792:
		goto st_case_792
	case 793:
		goto st_case_793
	case 794:
		goto st_case_794
	case 795:
		goto st_case_795
	case 796:
		goto st_case_796
	case 797:
		goto st_case_797
	case 798:
		goto st_case_798
	case 799:
		goto st_case_799
	case 800:
		goto st_case_800
	case 801:
		goto st_case_801
	case 802:
		goto st_case_802
	case 803:
		goto st_case_803
	case 804:
		goto st_case_804
	case 805:
		goto st_case_805
	case 806:
		goto st_case_806
	case 807:
		goto st_case_807
	case 808:
		goto st_case_808
	case 809:
		goto st_case_809
	case 810:
		goto st_case_810
	case 811:
		goto st_case_811
	case 812:
		goto st_case_812
	case 813:
		goto st_case_813
	case 814:
		goto st_case_814
	case 815:
		goto st_case_815
	case 816:
		goto st_case_816
	case 817:
		goto st_case_817
	case 818:
		goto st_case_818
	case 819:
		goto st_case_819
	case 820:
		goto st_case_820
	case 821:
		goto st_case_821
	case 822:
		goto st_case_822
	case 823:
		goto st_case_823
	case 824:
		goto st_case_824
	case 825:
		goto st_case_825
	case 826:
		goto st_case_826
	case 827:
		goto st_case_827
	case 828:
		goto st_case_828
	case 829:
		goto st_case_829
	case 830:
		goto st_case_830
	case 831:
		goto st_case_831
	case 832:
		goto st_case_832
	case 896:
		goto st_case_896
	case 833:
		goto st_case_833
	case 834:
		goto st_case_834
	case 897:
		goto st_case_897
	case 898:
		goto st_case_898
	case 835:
		goto st_case_835
	case 836:
		goto st_case_836
	case 837:
		goto st_case_837
	case 838:
		goto st_case_838
	case 839:
		goto st_case_839
	case 840:
		goto st_case_840
	case 841:
		goto st_case_841
	case 842:
		goto st_case_842
	case 843:
		goto st_case_843
	case 844:
		goto st_case_844
	case 845:
		goto st_case_845
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
		case 106:
			goto st820
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
			goto st455
		case 59:
			goto st570
		case 95:
			goto st781
		case 100:
			goto st806
		case 106:
			goto st810
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st554
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st781
			}
		default:
			goto st554
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
			goto st555
		case 95:
			goto st554
		case 100:
			goto st721
		case 106:
			goto st725
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st554
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st554
			}
		default:
			goto st554
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
			goto st346
		case 95:
			goto st9
		case 100:
			goto st450
		case 106:
			goto st456
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
			goto st346
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
			goto st179
		case 99:
			goto st185
		case 100:
			goto st193
		case 106:
			goto st300
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
		case 106:
			goto st169
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
		case 106:
			goto st169
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
			goto tr38
		case 95:
			goto st13
		case 98:
			goto st16
		case 99:
			goto st22
		case 100:
			goto st29
		case 106:
			goto st169
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
		goto tr37
tr37:
//line xss_170.rl:13

            return true
        
	goto st846
	st846:
		if p++; p == pe {
			goto _test_eof846
		}
	st_case_846:
//line xss_170.go:2280
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
		case 106:
			goto st169
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
			goto tr38
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
		goto tr37
tr38:
//line xss_170.rl:13

            return true
        
	goto st847
	st847:
		if p++; p == pe {
			goto _test_eof847
		}
	st_case_847:
//line xss_170.go:2473
		switch data[p] {
		case 62:
			goto tr38
		case 95:
			goto st13
		case 98:
			goto st16
		case 99:
			goto st22
		case 100:
			goto st29
		case 106:
			goto st169
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
		goto tr37
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
			goto st142
		case 99:
			goto st148
		case 100:
			goto st155
		case 106:
			goto st159
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
			goto st134
		case 99:
			goto st136
		case 100:
			goto st138
		case 106:
			goto st140
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
			goto st133
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st133
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st133
			}
		default:
			goto st133
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
		case 106:
			goto st57
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
			goto tr38
		case 59:
			goto tr38
		case 62:
			goto tr38
		case 95:
			goto st39
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto tr78
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
		goto tr37
tr78:
//line xss_170.rl:13

            return true
        
	goto st848
	st848:
		if p++; p == pe {
			goto _test_eof848
		}
	st_case_848:
//line xss_170.go:3221
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
		case 106:
			goto st57
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
		case 44:
			goto st15
		case 59:
			goto st15
		case 62:
			goto st15
		case 95:
			goto st39
		case 97:
			goto st58
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
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		switch data[p] {
		case 44:
			goto st15
		case 59:
			goto st15
		case 62:
			goto st15
		case 95:
			goto st39
		case 118:
			goto st59
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
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
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
			goto st60
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
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
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
			goto st61
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
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		switch data[p] {
		case 44:
			goto st15
		case 59:
			goto st15
		case 62:
			goto st15
		case 95:
			goto st39
		case 99:
			goto st62
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
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
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
			goto st63
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
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		switch data[p] {
		case 44:
			goto st15
		case 59:
			goto st15
		case 62:
			goto st15
		case 95:
			goto st39
		case 105:
			goto st64
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
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		switch data[p] {
		case 44:
			goto st15
		case 59:
			goto st15
		case 62:
			goto st15
		case 95:
			goto st39
		case 112:
			goto st65
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
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
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
			goto st66
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
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
		switch data[p] {
		case 44:
			goto st15
		case 58:
			goto st67
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
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
		switch data[p] {
		case 62:
			goto st69
		case 92:
			goto st132
		case 95:
			goto st70
		case 98:
			goto st72
		case 100:
			goto st78
		case 110:
			goto st126
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st70
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st70
			}
		default:
			goto st70
		}
		goto st68
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 91:
			goto st69
		case 92:
			goto st71
		case 95:
			goto st70
		case 98:
			goto st72
		case 100:
			goto st78
		case 110:
			goto st126
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st70
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			case data[p] >= 65:
				goto st70
			}
		default:
			goto st69
		}
		goto st68
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		switch data[p] {
		case 32:
			goto tr107
		case 40:
			goto tr107
		case 46:
			goto tr107
		case 91:
			goto tr107
		case 92:
			goto tr108
		case 95:
			goto st70
		case 98:
			goto st72
		case 100:
			goto st78
		case 110:
			goto st126
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st70
				}
			case data[p] >= 9:
				goto tr107
			}
		case data[p] > 62:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			case data[p] >= 65:
				goto st70
			}
		default:
			goto tr107
		}
		goto tr106
tr106:
//line xss_170.rl:13

            return true
        
	goto st849
	st849:
		if p++; p == pe {
			goto _test_eof849
		}
	st_case_849:
//line xss_170.go:4128
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 91:
			goto st69
		case 92:
			goto st71
		case 95:
			goto st70
		case 98:
			goto st72
		case 100:
			goto st78
		case 110:
			goto st126
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st70
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			case data[p] >= 65:
				goto st70
			}
		default:
			goto st69
		}
		goto st68
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 95:
			goto st70
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st70
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
		switch data[p] {
		case 32:
			goto tr107
		case 40:
			goto tr107
		case 46:
			goto tr107
		case 91:
			goto tr107
		case 92:
			goto tr108
		case 95:
			goto st70
		case 98:
			goto st72
		case 100:
			goto st78
		case 110:
			goto st126
		case 117:
			goto st130
		case 120:
			goto st130
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st70
				}
			case data[p] >= 9:
				goto tr107
			}
		case data[p] > 62:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			case data[p] >= 65:
				goto st70
			}
		default:
			goto tr107
		}
		goto tr106
tr107:
//line xss_170.rl:13

            return true
        
	goto st850
	st850:
		if p++; p == pe {
			goto _test_eof850
		}
	st_case_850:
//line xss_170.go:4277
		switch data[p] {
		case 32:
			goto tr107
		case 40:
			goto tr107
		case 46:
			goto tr107
		case 91:
			goto tr107
		case 92:
			goto tr108
		case 95:
			goto st70
		case 98:
			goto st72
		case 100:
			goto st78
		case 110:
			goto st126
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st70
				}
			case data[p] >= 9:
				goto tr107
			}
		case data[p] > 62:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			case data[p] >= 65:
				goto st70
			}
		default:
			goto tr107
		}
		goto tr106
tr108:
//line xss_170.rl:13

            return true
        
	goto st851
	st851:
		if p++; p == pe {
			goto _test_eof851
		}
	st_case_851:
//line xss_170.go:4332
		switch data[p] {
		case 32:
			goto tr107
		case 40:
			goto tr107
		case 46:
			goto tr107
		case 91:
			goto tr107
		case 92:
			goto tr108
		case 95:
			goto st70
		case 98:
			goto st72
		case 100:
			goto st78
		case 110:
			goto st126
		case 117:
			goto st130
		case 120:
			goto st130
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st70
				}
			case data[p] >= 9:
				goto tr107
			}
		case data[p] > 62:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			case data[p] >= 65:
				goto st70
			}
		default:
			goto tr107
		}
		goto tr106
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 95:
			goto st70
		case 97:
			goto st73
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st70
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st70
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 95:
			goto st70
		case 115:
			goto st74
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st70
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 95:
			goto st70
		case 101:
			goto st75
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st70
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 54:
			goto st76
		case 95:
			goto st70
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st70
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 52:
			goto st77
		case 95:
			goto st70
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st70
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
		switch data[p] {
		case 32:
			goto tr107
		case 40:
			goto tr107
		case 46:
			goto tr107
		case 95:
			goto st70
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st70
				}
			case data[p] >= 9:
				goto tr107
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			default:
				goto tr107
			}
		default:
			goto tr107
		}
		goto tr106
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 95:
			goto st70
		case 97:
			goto st79
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st70
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st70
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 95:
			goto st70
		case 116:
			goto st80
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st70
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 95:
			goto st70
		case 97:
			goto st81
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st70
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st70
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 58:
			goto st82
		case 95:
			goto st70
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st70
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 44:
			goto st69
		case 46:
			goto st69
		case 91:
			goto st69
		case 92:
			goto st71
		case 95:
			goto st83
		case 98:
			goto st111
		case 100:
			goto st117
		case 110:
			goto st121
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st70
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st83
				}
			case data[p] >= 65:
				goto st70
			}
		default:
			goto st69
		}
		goto st68
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 95:
			goto st84
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st84
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st84
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 47:
			goto st85
		case 95:
			goto st84
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st84
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st84
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 91:
			goto st69
		case 92:
			goto st71
		case 95:
			goto st86
		case 98:
			goto st105
		case 100:
			goto st107
		case 110:
			goto st109
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st86
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st86
				}
			case data[p] >= 65:
				goto st86
			}
		default:
			goto st69
		}
		goto st68
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 43:
			goto st87
		case 45:
			goto st87
		case 46:
			goto st69
		case 95:
			goto st104
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st104
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st104
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st104
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 43:
			goto st87
		case 45:
			goto st87
		case 46:
			goto st69
		case 91:
			goto st69
		case 92:
			goto st71
		case 95:
			goto st88
		case 98:
			goto st89
		case 100:
			goto st95
		case 110:
			goto st99
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st88
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st88
				}
			case data[p] >= 65:
				goto st88
			}
		default:
			goto st69
		}
		goto st68
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 44:
			goto st69
		case 46:
			goto st69
		case 95:
			goto st88
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st69
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st88
				}
			default:
				goto st87
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st88
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st88
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 44:
			goto st69
		case 46:
			goto st69
		case 95:
			goto st88
		case 97:
			goto st90
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st69
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st88
				}
			default:
				goto st87
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st88
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st88
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 44:
			goto st69
		case 46:
			goto st69
		case 95:
			goto st88
		case 115:
			goto st91
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st69
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st88
				}
			default:
				goto st87
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st88
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st88
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 44:
			goto st69
		case 46:
			goto st69
		case 95:
			goto st88
		case 101:
			goto st92
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st69
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st88
				}
			default:
				goto st87
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st88
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st88
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 44:
			goto st69
		case 46:
			goto st69
		case 54:
			goto st93
		case 95:
			goto st88
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st69
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st88
				}
			default:
				goto st87
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st88
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st88
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 44:
			goto st69
		case 46:
			goto st69
		case 52:
			goto st94
		case 95:
			goto st88
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st69
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st88
				}
			default:
				goto st87
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st88
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st88
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
		switch data[p] {
		case 32:
			goto tr107
		case 40:
			goto tr107
		case 44:
			goto tr107
		case 46:
			goto tr107
		case 95:
			goto st88
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr107
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st88
				}
			default:
				goto tr140
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st88
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st88
				}
			default:
				goto tr107
			}
		default:
			goto tr107
		}
		goto tr106
tr140:
//line xss_170.rl:13

            return true
        
	goto st852
	st852:
		if p++; p == pe {
			goto _test_eof852
		}
	st_case_852:
//line xss_170.go:5457
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 43:
			goto st87
		case 45:
			goto st87
		case 46:
			goto st69
		case 91:
			goto st69
		case 92:
			goto st71
		case 95:
			goto st88
		case 98:
			goto st89
		case 100:
			goto st95
		case 110:
			goto st99
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st88
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st88
				}
			case data[p] >= 65:
				goto st88
			}
		default:
			goto st69
		}
		goto st68
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 44:
			goto st69
		case 46:
			goto st69
		case 95:
			goto st88
		case 97:
			goto st96
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st69
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st88
				}
			default:
				goto st87
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st88
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st88
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 44:
			goto st69
		case 46:
			goto st69
		case 95:
			goto st88
		case 116:
			goto st97
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st69
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st88
				}
			default:
				goto st87
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st88
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st88
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 44:
			goto st69
		case 46:
			goto st69
		case 95:
			goto st88
		case 97:
			goto st98
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st69
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st88
				}
			default:
				goto st87
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st88
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st88
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 44:
			goto st69
		case 46:
			goto st69
		case 58:
			goto st82
		case 95:
			goto st88
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st69
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st88
				}
			default:
				goto st87
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st88
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st88
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 44:
			goto st69
		case 46:
			goto st69
		case 95:
			goto st88
		case 97:
			goto st100
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st69
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st88
				}
			default:
				goto st87
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st88
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st88
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 44:
			goto st69
		case 46:
			goto st69
		case 95:
			goto st88
		case 109:
			goto st101
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st69
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st88
				}
			default:
				goto st87
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st88
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st88
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 44:
			goto st69
		case 46:
			goto st69
		case 95:
			goto st88
		case 101:
			goto st102
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st69
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st88
				}
			default:
				goto st87
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st88
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st88
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
		switch data[p] {
		case 43:
			goto st103
		case 45:
			goto st103
		case 95:
			goto st88
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st88
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st88
			}
		default:
			goto st88
		}
		goto st69
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
		switch data[p] {
		case 32:
			goto tr107
		case 40:
			goto tr107
		case 43:
			goto tr140
		case 45:
			goto tr140
		case 46:
			goto tr107
		case 91:
			goto tr107
		case 92:
			goto tr108
		case 95:
			goto st88
		case 98:
			goto st89
		case 100:
			goto st95
		case 110:
			goto st99
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st88
				}
			case data[p] >= 9:
				goto tr107
			}
		case data[p] > 62:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st88
				}
			case data[p] >= 65:
				goto st88
			}
		default:
			goto tr107
		}
		goto tr106
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 43:
			goto st87
		case 45:
			goto st87
		case 46:
			goto st69
		case 95:
			goto st88
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st88
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st88
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st88
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 43:
			goto st87
		case 45:
			goto st87
		case 46:
			goto st69
		case 95:
			goto st104
		case 97:
			goto st106
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st104
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st104
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st104
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 43:
			goto st87
		case 45:
			goto st87
		case 46:
			goto st69
		case 95:
			goto st88
		case 115:
			goto st91
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st88
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st88
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st88
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 43:
			goto st87
		case 45:
			goto st87
		case 46:
			goto st69
		case 95:
			goto st104
		case 97:
			goto st108
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st104
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st104
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st104
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 43:
			goto st87
		case 45:
			goto st87
		case 46:
			goto st69
		case 95:
			goto st88
		case 116:
			goto st97
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st88
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st88
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st88
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 43:
			goto st87
		case 45:
			goto st87
		case 46:
			goto st69
		case 95:
			goto st104
		case 97:
			goto st110
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st104
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st104
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st104
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 43:
			goto st87
		case 45:
			goto st87
		case 46:
			goto st69
		case 95:
			goto st88
		case 109:
			goto st101
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st88
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st88
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st88
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 95:
			goto st84
		case 97:
			goto st112
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st84
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st84
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 47:
			goto st85
		case 95:
			goto st84
		case 115:
			goto st113
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st84
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st84
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 47:
			goto st85
		case 95:
			goto st84
		case 101:
			goto st114
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st84
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st84
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 47:
			goto st85
		case 54:
			goto st115
		case 95:
			goto st84
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st84
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st84
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 47:
			goto st85
		case 52:
			goto st116
		case 95:
			goto st84
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st84
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st84
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
		switch data[p] {
		case 32:
			goto tr107
		case 40:
			goto tr107
		case 46:
			goto tr107
		case 47:
			goto tr156
		case 95:
			goto st84
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st84
				}
			case data[p] >= 9:
				goto tr107
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st84
				}
			default:
				goto tr107
			}
		default:
			goto tr107
		}
		goto tr106
tr156:
//line xss_170.rl:13

            return true
        
	goto st853
	st853:
		if p++; p == pe {
			goto _test_eof853
		}
	st_case_853:
//line xss_170.go:6550
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 91:
			goto st69
		case 92:
			goto st71
		case 95:
			goto st86
		case 98:
			goto st105
		case 100:
			goto st107
		case 110:
			goto st109
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st86
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st86
				}
			case data[p] >= 65:
				goto st86
			}
		default:
			goto st69
		}
		goto st68
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 95:
			goto st84
		case 97:
			goto st118
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st84
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st84
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 47:
			goto st85
		case 95:
			goto st84
		case 116:
			goto st119
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st84
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st84
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 47:
			goto st85
		case 95:
			goto st84
		case 97:
			goto st120
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st84
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st84
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 47:
			goto st85
		case 58:
			goto st82
		case 95:
			goto st84
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st84
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st84
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 95:
			goto st84
		case 97:
			goto st122
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st84
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st84
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st122:
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 47:
			goto st85
		case 95:
			goto st84
		case 109:
			goto st123
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st84
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st84
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 47:
			goto st85
		case 95:
			goto st84
		case 101:
			goto st124
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st84
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st84
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
		switch data[p] {
		case 47:
			goto st125
		case 95:
			goto st84
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st84
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st84
			}
		default:
			goto st70
		}
		goto st69
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
		switch data[p] {
		case 32:
			goto tr107
		case 40:
			goto tr107
		case 46:
			goto tr107
		case 91:
			goto tr107
		case 92:
			goto tr108
		case 95:
			goto st86
		case 98:
			goto st105
		case 100:
			goto st107
		case 110:
			goto st109
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st86
				}
			case data[p] >= 9:
				goto tr107
			}
		case data[p] > 62:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st86
				}
			case data[p] >= 65:
				goto st86
			}
		default:
			goto tr107
		}
		goto tr106
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 95:
			goto st70
		case 97:
			goto st127
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st70
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st70
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 95:
			goto st70
		case 109:
			goto st128
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st70
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 95:
			goto st70
		case 101:
			goto st129
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st70
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
		if data[p] == 95 {
			goto st70
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st70
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st70
			}
		default:
			goto st70
		}
		goto st69
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 95:
			goto st70
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st131
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			default:
				goto st69
			}
		default:
			goto st69
		}
		goto st68
	st131:
		if p++; p == pe {
			goto _test_eof131
		}
	st_case_131:
		switch data[p] {
		case 32:
			goto tr107
		case 40:
			goto tr107
		case 46:
			goto tr107
		case 95:
			goto st70
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st131
				}
			case data[p] >= 9:
				goto tr107
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st70
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			default:
				goto tr107
			}
		default:
			goto tr107
		}
		goto tr106
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
		switch data[p] {
		case 32:
			goto st69
		case 40:
			goto st69
		case 46:
			goto st69
		case 91:
			goto st69
		case 92:
			goto st71
		case 95:
			goto st70
		case 98:
			goto st72
		case 100:
			goto st78
		case 110:
			goto st126
		case 117:
			goto st130
		case 120:
			goto st130
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st70
				}
			case data[p] >= 9:
				goto st69
			}
		case data[p] > 62:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			case data[p] >= 65:
				goto st70
			}
		default:
			goto st69
		}
		goto st68
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
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
	st134:
		if p++; p == pe {
			goto _test_eof134
		}
	st_case_134:
		switch data[p] {
		case 43:
			goto st38
		case 45:
			goto st38
		case 62:
			goto st15
		case 95:
			goto st133
		case 97:
			goto st135
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st133
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st133
			}
		default:
			goto st133
		}
		goto st14
	st135:
		if p++; p == pe {
			goto _test_eof135
		}
	st_case_135:
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
	st136:
		if p++; p == pe {
			goto _test_eof136
		}
	st_case_136:
		switch data[p] {
		case 43:
			goto st38
		case 45:
			goto st38
		case 62:
			goto st15
		case 95:
			goto st133
		case 104:
			goto st137
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st133
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st133
			}
		default:
			goto st133
		}
		goto st14
	st137:
		if p++; p == pe {
			goto _test_eof137
		}
	st_case_137:
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
	st138:
		if p++; p == pe {
			goto _test_eof138
		}
	st_case_138:
		switch data[p] {
		case 43:
			goto st38
		case 45:
			goto st38
		case 62:
			goto st15
		case 95:
			goto st133
		case 97:
			goto st139
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st133
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st133
			}
		default:
			goto st133
		}
		goto st14
	st139:
		if p++; p == pe {
			goto _test_eof139
		}
	st_case_139:
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
	st140:
		if p++; p == pe {
			goto _test_eof140
		}
	st_case_140:
		switch data[p] {
		case 43:
			goto st38
		case 45:
			goto st38
		case 62:
			goto st15
		case 95:
			goto st133
		case 97:
			goto st141
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st133
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st133
			}
		default:
			goto st133
		}
		goto st14
	st141:
		if p++; p == pe {
			goto _test_eof141
		}
	st_case_141:
		switch data[p] {
		case 43:
			goto st38
		case 45:
			goto st38
		case 62:
			goto st15
		case 95:
			goto st39
		case 118:
			goto st59
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
	st142:
		if p++; p == pe {
			goto _test_eof142
		}
	st_case_142:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st35
		case 97:
			goto st143
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
	st143:
		if p++; p == pe {
			goto _test_eof143
		}
	st_case_143:
		switch data[p] {
		case 47:
			goto st36
		case 62:
			goto st15
		case 95:
			goto st35
		case 115:
			goto st144
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
	st144:
		if p++; p == pe {
			goto _test_eof144
		}
	st_case_144:
		switch data[p] {
		case 47:
			goto st36
		case 62:
			goto st15
		case 95:
			goto st35
		case 101:
			goto st145
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
	st145:
		if p++; p == pe {
			goto _test_eof145
		}
	st_case_145:
		switch data[p] {
		case 47:
			goto st36
		case 54:
			goto st146
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
	st146:
		if p++; p == pe {
			goto _test_eof146
		}
	st_case_146:
		switch data[p] {
		case 47:
			goto st36
		case 52:
			goto st147
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
	st147:
		if p++; p == pe {
			goto _test_eof147
		}
	st_case_147:
		switch data[p] {
		case 47:
			goto tr177
		case 62:
			goto tr38
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
		goto tr37
tr177:
//line xss_170.rl:13

            return true
        
	goto st854
	st854:
		if p++; p == pe {
			goto _test_eof854
		}
	st_case_854:
//line xss_170.go:7716
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st37
		case 98:
			goto st134
		case 99:
			goto st136
		case 100:
			goto st138
		case 106:
			goto st140
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
	st148:
		if p++; p == pe {
			goto _test_eof148
		}
	st_case_148:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st35
		case 104:
			goto st149
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
	st149:
		if p++; p == pe {
			goto _test_eof149
		}
	st_case_149:
		switch data[p] {
		case 47:
			goto st36
		case 62:
			goto st15
		case 95:
			goto st35
		case 97:
			goto st150
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
	st150:
		if p++; p == pe {
			goto _test_eof150
		}
	st_case_150:
		switch data[p] {
		case 47:
			goto st36
		case 62:
			goto st15
		case 95:
			goto st35
		case 114:
			goto st151
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
	st151:
		if p++; p == pe {
			goto _test_eof151
		}
	st_case_151:
		switch data[p] {
		case 47:
			goto st36
		case 62:
			goto st15
		case 95:
			goto st35
		case 115:
			goto st152
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
	st152:
		if p++; p == pe {
			goto _test_eof152
		}
	st_case_152:
		switch data[p] {
		case 47:
			goto st36
		case 62:
			goto st15
		case 95:
			goto st35
		case 101:
			goto st153
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
	st153:
		if p++; p == pe {
			goto _test_eof153
		}
	st_case_153:
		switch data[p] {
		case 47:
			goto st36
		case 62:
			goto st15
		case 95:
			goto st35
		case 116:
			goto st154
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
	st154:
		if p++; p == pe {
			goto _test_eof154
		}
	st_case_154:
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
	st155:
		if p++; p == pe {
			goto _test_eof155
		}
	st_case_155:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st35
		case 97:
			goto st156
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
	st156:
		if p++; p == pe {
			goto _test_eof156
		}
	st_case_156:
		switch data[p] {
		case 47:
			goto st36
		case 62:
			goto st15
		case 95:
			goto st35
		case 116:
			goto st157
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
	st157:
		if p++; p == pe {
			goto _test_eof157
		}
	st_case_157:
		switch data[p] {
		case 47:
			goto st36
		case 62:
			goto st15
		case 95:
			goto st35
		case 97:
			goto st158
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
	st158:
		if p++; p == pe {
			goto _test_eof158
		}
	st_case_158:
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
	st159:
		if p++; p == pe {
			goto _test_eof159
		}
	st_case_159:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st35
		case 97:
			goto st160
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
	st160:
		if p++; p == pe {
			goto _test_eof160
		}
	st_case_160:
		switch data[p] {
		case 47:
			goto st36
		case 62:
			goto st15
		case 95:
			goto st35
		case 118:
			goto st161
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
	st161:
		if p++; p == pe {
			goto _test_eof161
		}
	st_case_161:
		switch data[p] {
		case 47:
			goto st36
		case 62:
			goto st15
		case 95:
			goto st35
		case 97:
			goto st162
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
	st162:
		if p++; p == pe {
			goto _test_eof162
		}
	st_case_162:
		switch data[p] {
		case 47:
			goto st36
		case 62:
			goto st15
		case 95:
			goto st35
		case 115:
			goto st163
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
	st163:
		if p++; p == pe {
			goto _test_eof163
		}
	st_case_163:
		switch data[p] {
		case 47:
			goto st36
		case 62:
			goto st15
		case 95:
			goto st35
		case 99:
			goto st164
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
	st164:
		if p++; p == pe {
			goto _test_eof164
		}
	st_case_164:
		switch data[p] {
		case 47:
			goto st36
		case 62:
			goto st15
		case 95:
			goto st35
		case 114:
			goto st165
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
	st165:
		if p++; p == pe {
			goto _test_eof165
		}
	st_case_165:
		switch data[p] {
		case 47:
			goto st36
		case 62:
			goto st15
		case 95:
			goto st35
		case 105:
			goto st166
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
	st166:
		if p++; p == pe {
			goto _test_eof166
		}
	st_case_166:
		switch data[p] {
		case 47:
			goto st36
		case 62:
			goto st15
		case 95:
			goto st35
		case 112:
			goto st167
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
	st167:
		if p++; p == pe {
			goto _test_eof167
		}
	st_case_167:
		switch data[p] {
		case 47:
			goto st36
		case 62:
			goto st15
		case 95:
			goto st35
		case 116:
			goto st168
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
	st168:
		if p++; p == pe {
			goto _test_eof168
		}
	st_case_168:
		switch data[p] {
		case 47:
			goto st36
		case 58:
			goto st67
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
	st169:
		if p++; p == pe {
			goto _test_eof169
		}
	st_case_169:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st13
		case 97:
			goto st170
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
	st170:
		if p++; p == pe {
			goto _test_eof170
		}
	st_case_170:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st13
		case 118:
			goto st171
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
	st171:
		if p++; p == pe {
			goto _test_eof171
		}
	st_case_171:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st13
		case 97:
			goto st172
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
	st172:
		if p++; p == pe {
			goto _test_eof172
		}
	st_case_172:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st13
		case 115:
			goto st173
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
	st173:
		if p++; p == pe {
			goto _test_eof173
		}
	st_case_173:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st13
		case 99:
			goto st174
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
	st174:
		if p++; p == pe {
			goto _test_eof174
		}
	st_case_174:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st13
		case 114:
			goto st175
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
	st175:
		if p++; p == pe {
			goto _test_eof175
		}
	st_case_175:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st13
		case 105:
			goto st176
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
	st176:
		if p++; p == pe {
			goto _test_eof176
		}
	st_case_176:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st13
		case 112:
			goto st177
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
	st177:
		if p++; p == pe {
			goto _test_eof177
		}
	st_case_177:
		switch data[p] {
		case 62:
			goto st15
		case 95:
			goto st13
		case 116:
			goto st178
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
	st178:
		if p++; p == pe {
			goto _test_eof178
		}
	st_case_178:
		switch data[p] {
		case 58:
			goto st67
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
	st179:
		if p++; p == pe {
			goto _test_eof179
		}
	st_case_179:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 97:
			goto st180
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
	st180:
		if p++; p == pe {
			goto _test_eof180
		}
	st_case_180:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 115:
			goto st181
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
	st181:
		if p++; p == pe {
			goto _test_eof181
		}
	st_case_181:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 101:
			goto st182
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
	st182:
		if p++; p == pe {
			goto _test_eof182
		}
	st_case_182:
		switch data[p] {
		case 54:
			goto st183
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
	st183:
		if p++; p == pe {
			goto _test_eof183
		}
	st_case_183:
		switch data[p] {
		case 52:
			goto st184
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
	st184:
		if p++; p == pe {
			goto _test_eof184
		}
	st_case_184:
		switch data[p] {
		case 60:
			goto tr211
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
		goto tr210
tr210:
//line xss_170.rl:13

            return true
        
	goto st855
	st855:
		if p++; p == pe {
			goto _test_eof855
		}
	st_case_855:
//line xss_170.go:8752
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 98:
			goto st179
		case 99:
			goto st185
		case 100:
			goto st193
		case 106:
			goto st300
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
	st185:
		if p++; p == pe {
			goto _test_eof185
		}
	st_case_185:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 104:
			goto st186
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
	st186:
		if p++; p == pe {
			goto _test_eof186
		}
	st_case_186:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 97:
			goto st187
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
	st187:
		if p++; p == pe {
			goto _test_eof187
		}
	st_case_187:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 114:
			goto st188
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
	st188:
		if p++; p == pe {
			goto _test_eof188
		}
	st_case_188:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 115:
			goto st189
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
	st189:
		if p++; p == pe {
			goto _test_eof189
		}
	st_case_189:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 101:
			goto st190
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
	st190:
		if p++; p == pe {
			goto _test_eof190
		}
	st_case_190:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 116:
			goto st191
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
	st191:
		if p++; p == pe {
			goto _test_eof191
		}
	st_case_191:
		switch data[p] {
		case 60:
			goto st12
		case 61:
			goto st192
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
	st192:
		if p++; p == pe {
			goto _test_eof192
		}
	st_case_192:
		switch data[p] {
		case 60:
			goto tr211
		case 95:
			goto st11
		case 98:
			goto st179
		case 99:
			goto st185
		case 100:
			goto st193
		case 106:
			goto st300
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
		goto tr210
tr211:
//line xss_170.rl:13

            return true
        
	goto st856
	st856:
		if p++; p == pe {
			goto _test_eof856
		}
	st_case_856:
//line xss_170.go:9005
		switch data[p] {
		case 95:
			goto st13
		case 98:
			goto st16
		case 99:
			goto st22
		case 100:
			goto st29
		case 106:
			goto st169
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
	st193:
		if p++; p == pe {
			goto _test_eof193
		}
	st_case_193:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 97:
			goto st194
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
	st194:
		if p++; p == pe {
			goto _test_eof194
		}
	st_case_194:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 116:
			goto st195
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
	st195:
		if p++; p == pe {
			goto _test_eof195
		}
	st_case_195:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 97:
			goto st196
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
	st196:
		if p++; p == pe {
			goto _test_eof196
		}
	st_case_196:
		switch data[p] {
		case 58:
			goto st197
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
	st197:
		if p++; p == pe {
			goto _test_eof197
		}
	st_case_197:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st192
		case 60:
			goto st12
		case 95:
			goto st198
		case 98:
			goto st319
		case 99:
			goto st325
		case 100:
			goto st332
		case 106:
			goto st336
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st198
			}
		default:
			goto st11
		}
		goto st10
	st198:
		if p++; p == pe {
			goto _test_eof198
		}
	st_case_198:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st199
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st199
			}
		default:
			goto st11
		}
		goto st10
	st199:
		if p++; p == pe {
			goto _test_eof199
		}
	st_case_199:
		switch data[p] {
		case 47:
			goto st200
		case 60:
			goto st12
		case 95:
			goto st199
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st199
			}
		default:
			goto st11
		}
		goto st10
	st200:
		if p++; p == pe {
			goto _test_eof200
		}
	st_case_200:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st201
		case 98:
			goto st311
		case 99:
			goto st313
		case 100:
			goto st315
		case 106:
			goto st317
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st201
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st201
			}
		default:
			goto st201
		}
		goto st10
	st201:
		if p++; p == pe {
			goto _test_eof201
		}
	st_case_201:
		switch data[p] {
		case 43:
			goto st202
		case 45:
			goto st202
		case 60:
			goto st12
		case 95:
			goto st310
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st310
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st310
			}
		default:
			goto st310
		}
		goto st10
	st202:
		if p++; p == pe {
			goto _test_eof202
		}
	st_case_202:
		switch data[p] {
		case 43:
			goto st202
		case 45:
			goto st202
		case 60:
			goto st12
		case 95:
			goto st203
		case 98:
			goto st204
		case 99:
			goto st210
		case 100:
			goto st217
		case 106:
			goto st221
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st203
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st203:
		if p++; p == pe {
			goto _test_eof203
		}
	st_case_203:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st192
		case 60:
			goto st12
		case 95:
			goto st203
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st202
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st203
				}
			case data[p] >= 65:
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st204:
		if p++; p == pe {
			goto _test_eof204
		}
	st_case_204:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st192
		case 60:
			goto st12
		case 95:
			goto st203
		case 97:
			goto st205
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st202
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st203
				}
			case data[p] >= 65:
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st205:
		if p++; p == pe {
			goto _test_eof205
		}
	st_case_205:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st192
		case 60:
			goto st12
		case 95:
			goto st203
		case 115:
			goto st206
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st202
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st203
				}
			case data[p] >= 65:
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st206:
		if p++; p == pe {
			goto _test_eof206
		}
	st_case_206:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st192
		case 60:
			goto st12
		case 95:
			goto st203
		case 101:
			goto st207
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st202
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st203
				}
			case data[p] >= 65:
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st207:
		if p++; p == pe {
			goto _test_eof207
		}
	st_case_207:
		switch data[p] {
		case 44:
			goto st192
		case 54:
			goto st208
		case 59:
			goto st192
		case 60:
			goto st12
		case 95:
			goto st203
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st202
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st203
				}
			case data[p] >= 65:
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st208:
		if p++; p == pe {
			goto _test_eof208
		}
	st_case_208:
		switch data[p] {
		case 44:
			goto st192
		case 52:
			goto st209
		case 59:
			goto st192
		case 60:
			goto st12
		case 95:
			goto st203
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st202
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st203
				}
			case data[p] >= 65:
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st209:
		if p++; p == pe {
			goto _test_eof209
		}
	st_case_209:
		switch data[p] {
		case 44:
			goto tr248
		case 59:
			goto tr248
		case 60:
			goto tr211
		case 95:
			goto st203
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto tr247
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st203
				}
			case data[p] >= 65:
				goto st203
			}
		default:
			goto st203
		}
		goto tr210
tr247:
//line xss_170.rl:13

            return true
        
	goto st857
	st857:
		if p++; p == pe {
			goto _test_eof857
		}
	st_case_857:
//line xss_170.go:9569
		switch data[p] {
		case 43:
			goto st202
		case 45:
			goto st202
		case 60:
			goto st12
		case 95:
			goto st203
		case 98:
			goto st204
		case 99:
			goto st210
		case 100:
			goto st217
		case 106:
			goto st221
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st203
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st210:
		if p++; p == pe {
			goto _test_eof210
		}
	st_case_210:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st192
		case 60:
			goto st12
		case 95:
			goto st203
		case 104:
			goto st211
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st202
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st203
				}
			case data[p] >= 65:
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st211:
		if p++; p == pe {
			goto _test_eof211
		}
	st_case_211:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st192
		case 60:
			goto st12
		case 95:
			goto st203
		case 97:
			goto st212
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st202
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st203
				}
			case data[p] >= 65:
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st212:
		if p++; p == pe {
			goto _test_eof212
		}
	st_case_212:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st192
		case 60:
			goto st12
		case 95:
			goto st203
		case 114:
			goto st213
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st202
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st203
				}
			case data[p] >= 65:
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st213:
		if p++; p == pe {
			goto _test_eof213
		}
	st_case_213:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st192
		case 60:
			goto st12
		case 95:
			goto st203
		case 115:
			goto st214
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st202
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st203
				}
			case data[p] >= 65:
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st214:
		if p++; p == pe {
			goto _test_eof214
		}
	st_case_214:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st192
		case 60:
			goto st12
		case 95:
			goto st203
		case 101:
			goto st215
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st202
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st203
				}
			case data[p] >= 65:
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st215:
		if p++; p == pe {
			goto _test_eof215
		}
	st_case_215:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st192
		case 60:
			goto st12
		case 95:
			goto st203
		case 116:
			goto st216
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st202
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st203
				}
			case data[p] >= 65:
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st216:
		if p++; p == pe {
			goto _test_eof216
		}
	st_case_216:
		switch data[p] {
		case 44:
			goto st192
		case 60:
			goto st12
		case 95:
			goto st203
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st203
				}
			case data[p] >= 43:
				goto st202
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st203
				}
			case data[p] >= 65:
				goto st203
			}
		default:
			goto st192
		}
		goto st10
	st217:
		if p++; p == pe {
			goto _test_eof217
		}
	st_case_217:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st192
		case 60:
			goto st12
		case 95:
			goto st203
		case 97:
			goto st218
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st202
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st203
				}
			case data[p] >= 65:
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st218:
		if p++; p == pe {
			goto _test_eof218
		}
	st_case_218:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st192
		case 60:
			goto st12
		case 95:
			goto st203
		case 116:
			goto st219
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st202
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st203
				}
			case data[p] >= 65:
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st219:
		if p++; p == pe {
			goto _test_eof219
		}
	st_case_219:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st192
		case 60:
			goto st12
		case 95:
			goto st203
		case 97:
			goto st220
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st202
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st203
				}
			case data[p] >= 65:
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st220:
		if p++; p == pe {
			goto _test_eof220
		}
	st_case_220:
		switch data[p] {
		case 44:
			goto st192
		case 58:
			goto st197
		case 59:
			goto st192
		case 60:
			goto st12
		case 95:
			goto st203
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st202
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st203
				}
			case data[p] >= 65:
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st221:
		if p++; p == pe {
			goto _test_eof221
		}
	st_case_221:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st192
		case 60:
			goto st12
		case 95:
			goto st203
		case 97:
			goto st222
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st202
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st203
				}
			case data[p] >= 65:
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st222:
		if p++; p == pe {
			goto _test_eof222
		}
	st_case_222:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st192
		case 60:
			goto st12
		case 95:
			goto st203
		case 118:
			goto st223
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st202
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st203
				}
			case data[p] >= 65:
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st223:
		if p++; p == pe {
			goto _test_eof223
		}
	st_case_223:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st192
		case 60:
			goto st12
		case 95:
			goto st203
		case 97:
			goto st224
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st202
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st203
				}
			case data[p] >= 65:
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st224:
		if p++; p == pe {
			goto _test_eof224
		}
	st_case_224:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st192
		case 60:
			goto st12
		case 95:
			goto st203
		case 115:
			goto st225
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st202
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st203
				}
			case data[p] >= 65:
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st225:
		if p++; p == pe {
			goto _test_eof225
		}
	st_case_225:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st192
		case 60:
			goto st12
		case 95:
			goto st203
		case 99:
			goto st226
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st202
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st203
				}
			case data[p] >= 65:
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st226:
		if p++; p == pe {
			goto _test_eof226
		}
	st_case_226:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st192
		case 60:
			goto st12
		case 95:
			goto st203
		case 114:
			goto st227
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st202
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st203
				}
			case data[p] >= 65:
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st227:
		if p++; p == pe {
			goto _test_eof227
		}
	st_case_227:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st192
		case 60:
			goto st12
		case 95:
			goto st203
		case 105:
			goto st228
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st202
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st203
				}
			case data[p] >= 65:
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st228:
		if p++; p == pe {
			goto _test_eof228
		}
	st_case_228:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st192
		case 60:
			goto st12
		case 95:
			goto st203
		case 112:
			goto st229
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st202
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st203
				}
			case data[p] >= 65:
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st229:
		if p++; p == pe {
			goto _test_eof229
		}
	st_case_229:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st192
		case 60:
			goto st12
		case 95:
			goto st203
		case 116:
			goto st230
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st202
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st203
				}
			case data[p] >= 65:
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st230:
		if p++; p == pe {
			goto _test_eof230
		}
	st_case_230:
		switch data[p] {
		case 44:
			goto st192
		case 58:
			goto st231
		case 59:
			goto st192
		case 60:
			goto st12
		case 95:
			goto st203
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st202
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st203
				}
			case data[p] >= 65:
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st231:
		if p++; p == pe {
			goto _test_eof231
		}
	st_case_231:
		switch data[p] {
		case 60:
			goto st236
		case 92:
			goto st299
		case 95:
			goto st234
		case 98:
			goto st239
		case 100:
			goto st245
		case 110:
			goto st293
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st234
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st234
			}
		default:
			goto st234
		}
		goto st232
	st232:
		if p++; p == pe {
			goto _test_eof232
		}
	st_case_232:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 61:
			goto st233
		case 91:
			goto st233
		case 92:
			goto st238
		case 95:
			goto st234
		case 98:
			goto st239
		case 100:
			goto st245
		case 110:
			goto st293
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st233
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st234
				}
			case data[p] >= 65:
				goto st234
			}
		default:
			goto st234
		}
		goto st232
	st233:
		if p++; p == pe {
			goto _test_eof233
		}
	st_case_233:
		switch data[p] {
		case 32:
			goto tr279
		case 40:
			goto tr279
		case 46:
			goto tr279
		case 60:
			goto tr280
		case 61:
			goto tr279
		case 91:
			goto tr279
		case 92:
			goto tr281
		case 95:
			goto st234
		case 98:
			goto st239
		case 100:
			goto st245
		case 110:
			goto st293
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr279
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st234
				}
			case data[p] >= 65:
				goto st234
			}
		default:
			goto st234
		}
		goto tr278
tr278:
//line xss_170.rl:13

            return true
        
	goto st858
	st858:
		if p++; p == pe {
			goto _test_eof858
		}
	st_case_858:
//line xss_170.go:10474
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 61:
			goto st233
		case 91:
			goto st233
		case 92:
			goto st238
		case 95:
			goto st234
		case 98:
			goto st239
		case 100:
			goto st245
		case 110:
			goto st293
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st233
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st234
				}
			case data[p] >= 65:
				goto st234
			}
		default:
			goto st234
		}
		goto st232
	st234:
		if p++; p == pe {
			goto _test_eof234
		}
	st_case_234:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st234
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st234
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st234
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st234
		}
		goto st232
	st235:
		if p++; p == pe {
			goto _test_eof235
		}
	st_case_235:
		switch data[p] {
		case 32:
			goto tr280
		case 40:
			goto tr280
		case 46:
			goto tr280
		case 91:
			goto tr280
		case 92:
			goto tr283
		case 95:
			goto st70
		case 98:
			goto st72
		case 100:
			goto st78
		case 110:
			goto st126
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st70
				}
			case data[p] >= 9:
				goto tr280
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			case data[p] >= 65:
				goto st70
			}
		default:
			goto tr280
		}
		goto tr282
tr282:
//line xss_170.rl:13

            return true
        
	goto st859
	st859:
		if p++; p == pe {
			goto _test_eof859
		}
	st_case_859:
//line xss_170.go:10618
		switch data[p] {
		case 32:
			goto st235
		case 40:
			goto st235
		case 46:
			goto st235
		case 91:
			goto st235
		case 92:
			goto st237
		case 95:
			goto st70
		case 98:
			goto st72
		case 100:
			goto st78
		case 110:
			goto st126
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st70
				}
			case data[p] >= 9:
				goto st235
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			case data[p] >= 65:
				goto st70
			}
		default:
			goto st235
		}
		goto st236
	st236:
		if p++; p == pe {
			goto _test_eof236
		}
	st_case_236:
		switch data[p] {
		case 32:
			goto st235
		case 40:
			goto st235
		case 46:
			goto st235
		case 91:
			goto st235
		case 92:
			goto st237
		case 95:
			goto st70
		case 98:
			goto st72
		case 100:
			goto st78
		case 110:
			goto st126
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st70
				}
			case data[p] >= 9:
				goto st235
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			case data[p] >= 65:
				goto st70
			}
		default:
			goto st235
		}
		goto st236
	st237:
		if p++; p == pe {
			goto _test_eof237
		}
	st_case_237:
		switch data[p] {
		case 32:
			goto tr280
		case 40:
			goto tr280
		case 46:
			goto tr280
		case 91:
			goto tr280
		case 92:
			goto tr283
		case 95:
			goto st70
		case 98:
			goto st72
		case 100:
			goto st78
		case 110:
			goto st126
		case 117:
			goto st130
		case 120:
			goto st130
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st70
				}
			case data[p] >= 9:
				goto tr280
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			case data[p] >= 65:
				goto st70
			}
		default:
			goto tr280
		}
		goto tr282
tr280:
//line xss_170.rl:13

            return true
        
	goto st860
	st860:
		if p++; p == pe {
			goto _test_eof860
		}
	st_case_860:
//line xss_170.go:10773
		switch data[p] {
		case 32:
			goto tr280
		case 40:
			goto tr280
		case 46:
			goto tr280
		case 91:
			goto tr280
		case 92:
			goto tr283
		case 95:
			goto st70
		case 98:
			goto st72
		case 100:
			goto st78
		case 110:
			goto st126
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st70
				}
			case data[p] >= 9:
				goto tr280
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			case data[p] >= 65:
				goto st70
			}
		default:
			goto tr280
		}
		goto tr282
tr283:
//line xss_170.rl:13

            return true
        
	goto st861
	st861:
		if p++; p == pe {
			goto _test_eof861
		}
	st_case_861:
//line xss_170.go:10828
		switch data[p] {
		case 32:
			goto tr280
		case 40:
			goto tr280
		case 46:
			goto tr280
		case 91:
			goto tr280
		case 92:
			goto tr283
		case 95:
			goto st70
		case 98:
			goto st72
		case 100:
			goto st78
		case 110:
			goto st126
		case 117:
			goto st130
		case 120:
			goto st130
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st70
				}
			case data[p] >= 9:
				goto tr280
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st70
				}
			case data[p] >= 65:
				goto st70
			}
		default:
			goto tr280
		}
		goto tr282
	st238:
		if p++; p == pe {
			goto _test_eof238
		}
	st_case_238:
		switch data[p] {
		case 32:
			goto tr279
		case 40:
			goto tr279
		case 46:
			goto tr279
		case 60:
			goto tr280
		case 61:
			goto tr279
		case 91:
			goto tr279
		case 92:
			goto tr281
		case 95:
			goto st234
		case 98:
			goto st239
		case 100:
			goto st245
		case 110:
			goto st293
		case 117:
			goto st297
		case 120:
			goto st297
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr279
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st234
				}
			case data[p] >= 65:
				goto st234
			}
		default:
			goto st234
		}
		goto tr278
tr279:
//line xss_170.rl:13

            return true
        
	goto st862
	st862:
		if p++; p == pe {
			goto _test_eof862
		}
	st_case_862:
//line xss_170.go:10938
		switch data[p] {
		case 32:
			goto tr279
		case 40:
			goto tr279
		case 46:
			goto tr279
		case 60:
			goto tr280
		case 61:
			goto tr279
		case 91:
			goto tr279
		case 92:
			goto tr281
		case 95:
			goto st234
		case 98:
			goto st239
		case 100:
			goto st245
		case 110:
			goto st293
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr279
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st234
				}
			case data[p] >= 65:
				goto st234
			}
		default:
			goto st234
		}
		goto tr278
tr281:
//line xss_170.rl:13

            return true
        
	goto st863
	st863:
		if p++; p == pe {
			goto _test_eof863
		}
	st_case_863:
//line xss_170.go:10992
		switch data[p] {
		case 32:
			goto tr279
		case 40:
			goto tr279
		case 46:
			goto tr279
		case 60:
			goto tr280
		case 61:
			goto tr279
		case 91:
			goto tr279
		case 92:
			goto tr281
		case 95:
			goto st234
		case 98:
			goto st239
		case 100:
			goto st245
		case 110:
			goto st293
		case 117:
			goto st297
		case 120:
			goto st297
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr279
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st234
				}
			case data[p] >= 65:
				goto st234
			}
		default:
			goto st234
		}
		goto tr278
	st239:
		if p++; p == pe {
			goto _test_eof239
		}
	st_case_239:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st234
		case 97:
			goto st240
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st234
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st234
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st234
		}
		goto st232
	st240:
		if p++; p == pe {
			goto _test_eof240
		}
	st_case_240:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st234
		case 115:
			goto st241
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st234
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st234
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st234
		}
		goto st232
	st241:
		if p++; p == pe {
			goto _test_eof241
		}
	st_case_241:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st234
		case 101:
			goto st242
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st234
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st234
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st234
		}
		goto st232
	st242:
		if p++; p == pe {
			goto _test_eof242
		}
	st_case_242:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 54:
			goto st243
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st234
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st234
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st234
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st234
		}
		goto st232
	st243:
		if p++; p == pe {
			goto _test_eof243
		}
	st_case_243:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 52:
			goto st244
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st234
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st234
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st234
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st234
		}
		goto st232
	st244:
		if p++; p == pe {
			goto _test_eof244
		}
	st_case_244:
		switch data[p] {
		case 32:
			goto tr279
		case 40:
			goto tr279
		case 46:
			goto tr279
		case 60:
			goto tr280
		case 61:
			goto tr279
		case 95:
			goto st234
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st234
				}
			case data[p] >= 9:
				goto tr279
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st234
				}
			case data[p] >= 91:
				goto tr279
			}
		default:
			goto st234
		}
		goto tr278
	st245:
		if p++; p == pe {
			goto _test_eof245
		}
	st_case_245:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st234
		case 97:
			goto st246
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st234
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st234
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st234
		}
		goto st232
	st246:
		if p++; p == pe {
			goto _test_eof246
		}
	st_case_246:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st234
		case 116:
			goto st247
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st234
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st234
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st234
		}
		goto st232
	st247:
		if p++; p == pe {
			goto _test_eof247
		}
	st_case_247:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st234
		case 97:
			goto st248
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st234
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st234
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st234
		}
		goto st232
	st248:
		if p++; p == pe {
			goto _test_eof248
		}
	st_case_248:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 58:
			goto st249
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st234
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st234
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st234
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st234
		}
		goto st232
	st249:
		if p++; p == pe {
			goto _test_eof249
		}
	st_case_249:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 44:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 91:
			goto st233
		case 92:
			goto st238
		case 95:
			goto st250
		case 98:
			goto st278
		case 100:
			goto st284
		case 110:
			goto st288
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st234
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st250
				}
			case data[p] >= 65:
				goto st234
			}
		default:
			goto st233
		}
		goto st232
	st250:
		if p++; p == pe {
			goto _test_eof250
		}
	st_case_250:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st251
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st251
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st251
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st234
		}
		goto st232
	st251:
		if p++; p == pe {
			goto _test_eof251
		}
	st_case_251:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 47:
			goto st252
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st251
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st251
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st251
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st234
		}
		goto st232
	st252:
		if p++; p == pe {
			goto _test_eof252
		}
	st_case_252:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 61:
			goto st233
		case 91:
			goto st233
		case 92:
			goto st238
		case 95:
			goto st253
		case 98:
			goto st272
		case 100:
			goto st274
		case 110:
			goto st276
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st233
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st253
				}
			case data[p] >= 65:
				goto st253
			}
		default:
			goto st253
		}
		goto st232
	st253:
		if p++; p == pe {
			goto _test_eof253
		}
	st_case_253:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 43:
			goto st254
		case 45:
			goto st254
		case 46:
			goto st233
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st271
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st271
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st271
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st271
		}
		goto st232
	st254:
		if p++; p == pe {
			goto _test_eof254
		}
	st_case_254:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 43:
			goto st254
		case 45:
			goto st254
		case 46:
			goto st233
		case 60:
			goto st235
		case 61:
			goto st233
		case 91:
			goto st233
		case 92:
			goto st238
		case 95:
			goto st255
		case 98:
			goto st256
		case 100:
			goto st262
		case 110:
			goto st266
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st233
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st255
				}
			case data[p] >= 65:
				goto st255
			}
		default:
			goto st255
		}
		goto st232
	st255:
		if p++; p == pe {
			goto _test_eof255
		}
	st_case_255:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 44:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 95:
			goto st255
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st233
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st255
				}
			default:
				goto st254
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st255
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st255
				}
			default:
				goto st233
			}
		default:
			goto st233
		}
		goto st232
	st256:
		if p++; p == pe {
			goto _test_eof256
		}
	st_case_256:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 44:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 95:
			goto st255
		case 97:
			goto st257
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st233
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st255
				}
			default:
				goto st254
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st255
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st255
				}
			default:
				goto st233
			}
		default:
			goto st233
		}
		goto st232
	st257:
		if p++; p == pe {
			goto _test_eof257
		}
	st_case_257:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 44:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 95:
			goto st255
		case 115:
			goto st258
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st233
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st255
				}
			default:
				goto st254
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st255
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st255
				}
			default:
				goto st233
			}
		default:
			goto st233
		}
		goto st232
	st258:
		if p++; p == pe {
			goto _test_eof258
		}
	st_case_258:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 44:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 95:
			goto st255
		case 101:
			goto st259
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st233
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st255
				}
			default:
				goto st254
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st255
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st255
				}
			default:
				goto st233
			}
		default:
			goto st233
		}
		goto st232
	st259:
		if p++; p == pe {
			goto _test_eof259
		}
	st_case_259:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 44:
			goto st233
		case 46:
			goto st233
		case 54:
			goto st260
		case 60:
			goto st235
		case 95:
			goto st255
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st233
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st255
				}
			default:
				goto st254
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st255
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st255
				}
			default:
				goto st233
			}
		default:
			goto st233
		}
		goto st232
	st260:
		if p++; p == pe {
			goto _test_eof260
		}
	st_case_260:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 44:
			goto st233
		case 46:
			goto st233
		case 52:
			goto st261
		case 60:
			goto st235
		case 95:
			goto st255
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st233
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st255
				}
			default:
				goto st254
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st255
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st255
				}
			default:
				goto st233
			}
		default:
			goto st233
		}
		goto st232
	st261:
		if p++; p == pe {
			goto _test_eof261
		}
	st_case_261:
		switch data[p] {
		case 32:
			goto tr279
		case 40:
			goto tr279
		case 44:
			goto tr279
		case 46:
			goto tr279
		case 60:
			goto tr280
		case 95:
			goto st255
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr279
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st255
				}
			default:
				goto tr316
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st255
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st255
				}
			default:
				goto tr279
			}
		default:
			goto tr279
		}
		goto tr278
tr316:
//line xss_170.rl:13

            return true
        
	goto st864
	st864:
		if p++; p == pe {
			goto _test_eof864
		}
	st_case_864:
//line xss_170.go:12130
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 43:
			goto st254
		case 45:
			goto st254
		case 46:
			goto st233
		case 60:
			goto st235
		case 61:
			goto st233
		case 91:
			goto st233
		case 92:
			goto st238
		case 95:
			goto st255
		case 98:
			goto st256
		case 100:
			goto st262
		case 110:
			goto st266
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st233
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st255
				}
			case data[p] >= 65:
				goto st255
			}
		default:
			goto st255
		}
		goto st232
	st262:
		if p++; p == pe {
			goto _test_eof262
		}
	st_case_262:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 44:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 95:
			goto st255
		case 97:
			goto st263
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st233
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st255
				}
			default:
				goto st254
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st255
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st255
				}
			default:
				goto st233
			}
		default:
			goto st233
		}
		goto st232
	st263:
		if p++; p == pe {
			goto _test_eof263
		}
	st_case_263:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 44:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 95:
			goto st255
		case 116:
			goto st264
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st233
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st255
				}
			default:
				goto st254
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st255
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st255
				}
			default:
				goto st233
			}
		default:
			goto st233
		}
		goto st232
	st264:
		if p++; p == pe {
			goto _test_eof264
		}
	st_case_264:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 44:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 95:
			goto st255
		case 97:
			goto st265
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st233
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st255
				}
			default:
				goto st254
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st255
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st255
				}
			default:
				goto st233
			}
		default:
			goto st233
		}
		goto st232
	st265:
		if p++; p == pe {
			goto _test_eof265
		}
	st_case_265:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 44:
			goto st233
		case 46:
			goto st233
		case 58:
			goto st249
		case 60:
			goto st235
		case 95:
			goto st255
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st233
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st255
				}
			default:
				goto st254
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st255
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st255
				}
			default:
				goto st233
			}
		default:
			goto st233
		}
		goto st232
	st266:
		if p++; p == pe {
			goto _test_eof266
		}
	st_case_266:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 44:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 95:
			goto st255
		case 97:
			goto st267
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st233
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st255
				}
			default:
				goto st254
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st255
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st255
				}
			default:
				goto st233
			}
		default:
			goto st233
		}
		goto st232
	st267:
		if p++; p == pe {
			goto _test_eof267
		}
	st_case_267:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 44:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 95:
			goto st255
		case 109:
			goto st268
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st233
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st255
				}
			default:
				goto st254
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st255
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st255
				}
			default:
				goto st233
			}
		default:
			goto st233
		}
		goto st232
	st268:
		if p++; p == pe {
			goto _test_eof268
		}
	st_case_268:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 44:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 95:
			goto st255
		case 101:
			goto st269
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st233
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st255
				}
			default:
				goto st254
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st255
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st255
				}
			default:
				goto st233
			}
		default:
			goto st233
		}
		goto st232
	st269:
		if p++; p == pe {
			goto _test_eof269
		}
	st_case_269:
		switch data[p] {
		case 43:
			goto st270
		case 45:
			goto st270
		case 60:
			goto st235
		case 95:
			goto st255
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st255
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st255
			}
		default:
			goto st255
		}
		goto st233
	st270:
		if p++; p == pe {
			goto _test_eof270
		}
	st_case_270:
		switch data[p] {
		case 32:
			goto tr279
		case 40:
			goto tr279
		case 43:
			goto tr316
		case 45:
			goto tr316
		case 46:
			goto tr279
		case 60:
			goto tr280
		case 61:
			goto tr279
		case 91:
			goto tr279
		case 92:
			goto tr281
		case 95:
			goto st255
		case 98:
			goto st256
		case 100:
			goto st262
		case 110:
			goto st266
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr279
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st255
				}
			case data[p] >= 65:
				goto st255
			}
		default:
			goto st255
		}
		goto tr278
	st271:
		if p++; p == pe {
			goto _test_eof271
		}
	st_case_271:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 43:
			goto st254
		case 45:
			goto st254
		case 46:
			goto st233
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st255
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st255
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st255
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st255
		}
		goto st232
	st272:
		if p++; p == pe {
			goto _test_eof272
		}
	st_case_272:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 43:
			goto st254
		case 45:
			goto st254
		case 46:
			goto st233
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st271
		case 97:
			goto st273
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st271
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st271
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st271
		}
		goto st232
	st273:
		if p++; p == pe {
			goto _test_eof273
		}
	st_case_273:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 43:
			goto st254
		case 45:
			goto st254
		case 46:
			goto st233
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st255
		case 115:
			goto st258
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st255
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st255
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st255
		}
		goto st232
	st274:
		if p++; p == pe {
			goto _test_eof274
		}
	st_case_274:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 43:
			goto st254
		case 45:
			goto st254
		case 46:
			goto st233
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st271
		case 97:
			goto st275
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st271
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st271
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st271
		}
		goto st232
	st275:
		if p++; p == pe {
			goto _test_eof275
		}
	st_case_275:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 43:
			goto st254
		case 45:
			goto st254
		case 46:
			goto st233
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st255
		case 116:
			goto st264
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st255
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st255
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st255
		}
		goto st232
	st276:
		if p++; p == pe {
			goto _test_eof276
		}
	st_case_276:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 43:
			goto st254
		case 45:
			goto st254
		case 46:
			goto st233
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st271
		case 97:
			goto st277
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st271
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st271
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st271
		}
		goto st232
	st277:
		if p++; p == pe {
			goto _test_eof277
		}
	st_case_277:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 43:
			goto st254
		case 45:
			goto st254
		case 46:
			goto st233
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st255
		case 109:
			goto st268
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st255
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st255
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st255
		}
		goto st232
	st278:
		if p++; p == pe {
			goto _test_eof278
		}
	st_case_278:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st251
		case 97:
			goto st279
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st251
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st251
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st234
		}
		goto st232
	st279:
		if p++; p == pe {
			goto _test_eof279
		}
	st_case_279:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 47:
			goto st252
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st251
		case 115:
			goto st280
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st251
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st251
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st234
		}
		goto st232
	st280:
		if p++; p == pe {
			goto _test_eof280
		}
	st_case_280:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 47:
			goto st252
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st251
		case 101:
			goto st281
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st251
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st251
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st234
		}
		goto st232
	st281:
		if p++; p == pe {
			goto _test_eof281
		}
	st_case_281:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 47:
			goto st252
		case 54:
			goto st282
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st251
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st251
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st251
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st234
		}
		goto st232
	st282:
		if p++; p == pe {
			goto _test_eof282
		}
	st_case_282:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 47:
			goto st252
		case 52:
			goto st283
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st251
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st251
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st251
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st234
		}
		goto st232
	st283:
		if p++; p == pe {
			goto _test_eof283
		}
	st_case_283:
		switch data[p] {
		case 32:
			goto tr279
		case 40:
			goto tr279
		case 46:
			goto tr279
		case 47:
			goto tr332
		case 60:
			goto tr280
		case 61:
			goto tr279
		case 95:
			goto st251
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st251
				}
			case data[p] >= 9:
				goto tr279
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st251
				}
			case data[p] >= 91:
				goto tr279
			}
		default:
			goto st234
		}
		goto tr278
tr332:
//line xss_170.rl:13

            return true
        
	goto st865
	st865:
		if p++; p == pe {
			goto _test_eof865
		}
	st_case_865:
//line xss_170.go:13237
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 61:
			goto st233
		case 91:
			goto st233
		case 92:
			goto st238
		case 95:
			goto st253
		case 98:
			goto st272
		case 100:
			goto st274
		case 110:
			goto st276
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st233
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st253
				}
			case data[p] >= 65:
				goto st253
			}
		default:
			goto st253
		}
		goto st232
	st284:
		if p++; p == pe {
			goto _test_eof284
		}
	st_case_284:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st251
		case 97:
			goto st285
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st251
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st251
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st234
		}
		goto st232
	st285:
		if p++; p == pe {
			goto _test_eof285
		}
	st_case_285:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 47:
			goto st252
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st251
		case 116:
			goto st286
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st251
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st251
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st234
		}
		goto st232
	st286:
		if p++; p == pe {
			goto _test_eof286
		}
	st_case_286:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 47:
			goto st252
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st251
		case 97:
			goto st287
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st251
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st251
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st234
		}
		goto st232
	st287:
		if p++; p == pe {
			goto _test_eof287
		}
	st_case_287:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 47:
			goto st252
		case 58:
			goto st249
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st251
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st251
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st251
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st234
		}
		goto st232
	st288:
		if p++; p == pe {
			goto _test_eof288
		}
	st_case_288:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st251
		case 97:
			goto st289
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st251
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st251
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st234
		}
		goto st232
	st289:
		if p++; p == pe {
			goto _test_eof289
		}
	st_case_289:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 47:
			goto st252
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st251
		case 109:
			goto st290
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st251
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st251
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st234
		}
		goto st232
	st290:
		if p++; p == pe {
			goto _test_eof290
		}
	st_case_290:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 47:
			goto st252
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st251
		case 101:
			goto st291
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st251
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st251
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st234
		}
		goto st232
	st291:
		if p++; p == pe {
			goto _test_eof291
		}
	st_case_291:
		switch data[p] {
		case 47:
			goto st292
		case 60:
			goto st235
		case 95:
			goto st251
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st251
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st251
			}
		default:
			goto st234
		}
		goto st233
	st292:
		if p++; p == pe {
			goto _test_eof292
		}
	st_case_292:
		switch data[p] {
		case 32:
			goto tr279
		case 40:
			goto tr279
		case 46:
			goto tr279
		case 60:
			goto tr280
		case 61:
			goto tr279
		case 91:
			goto tr279
		case 92:
			goto tr281
		case 95:
			goto st253
		case 98:
			goto st272
		case 100:
			goto st274
		case 110:
			goto st276
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr279
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st253
				}
			case data[p] >= 65:
				goto st253
			}
		default:
			goto st253
		}
		goto tr278
	st293:
		if p++; p == pe {
			goto _test_eof293
		}
	st_case_293:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st234
		case 97:
			goto st294
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st234
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st234
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st234
		}
		goto st232
	st294:
		if p++; p == pe {
			goto _test_eof294
		}
	st_case_294:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st234
		case 109:
			goto st295
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st234
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st234
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st234
		}
		goto st232
	st295:
		if p++; p == pe {
			goto _test_eof295
		}
	st_case_295:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st234
		case 101:
			goto st296
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st234
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st234
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st234
		}
		goto st232
	st296:
		if p++; p == pe {
			goto _test_eof296
		}
	st_case_296:
		switch data[p] {
		case 60:
			goto st235
		case 95:
			goto st234
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st234
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st234
			}
		default:
			goto st234
		}
		goto st233
	st297:
		if p++; p == pe {
			goto _test_eof297
		}
	st_case_297:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 61:
			goto st233
		case 95:
			goto st234
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st298
				}
			case data[p] >= 9:
				goto st233
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st234
				}
			case data[p] >= 91:
				goto st233
			}
		default:
			goto st234
		}
		goto st232
	st298:
		if p++; p == pe {
			goto _test_eof298
		}
	st_case_298:
		switch data[p] {
		case 32:
			goto tr279
		case 40:
			goto tr279
		case 46:
			goto tr279
		case 60:
			goto tr280
		case 61:
			goto tr279
		case 95:
			goto st234
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st298
				}
			case data[p] >= 9:
				goto tr279
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st234
				}
			case data[p] >= 91:
				goto tr279
			}
		default:
			goto st234
		}
		goto tr278
	st299:
		if p++; p == pe {
			goto _test_eof299
		}
	st_case_299:
		switch data[p] {
		case 32:
			goto st233
		case 40:
			goto st233
		case 46:
			goto st233
		case 60:
			goto st235
		case 61:
			goto st233
		case 91:
			goto st233
		case 92:
			goto st238
		case 95:
			goto st234
		case 98:
			goto st239
		case 100:
			goto st245
		case 110:
			goto st293
		case 117:
			goto st297
		case 120:
			goto st297
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st233
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st234
				}
			case data[p] >= 65:
				goto st234
			}
		default:
			goto st234
		}
		goto st232
tr248:
//line xss_170.rl:13

            return true
        
	goto st866
	st866:
		if p++; p == pe {
			goto _test_eof866
		}
	st_case_866:
//line xss_170.go:13973
		switch data[p] {
		case 60:
			goto tr211
		case 95:
			goto st11
		case 98:
			goto st179
		case 99:
			goto st185
		case 100:
			goto st193
		case 106:
			goto st300
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
		goto tr210
	st300:
		if p++; p == pe {
			goto _test_eof300
		}
	st_case_300:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 97:
			goto st301
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
	st301:
		if p++; p == pe {
			goto _test_eof301
		}
	st_case_301:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 118:
			goto st302
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
	st302:
		if p++; p == pe {
			goto _test_eof302
		}
	st_case_302:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 97:
			goto st303
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
	st303:
		if p++; p == pe {
			goto _test_eof303
		}
	st_case_303:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 115:
			goto st304
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
	st304:
		if p++; p == pe {
			goto _test_eof304
		}
	st_case_304:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 99:
			goto st305
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
	st305:
		if p++; p == pe {
			goto _test_eof305
		}
	st_case_305:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 114:
			goto st306
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
	st306:
		if p++; p == pe {
			goto _test_eof306
		}
	st_case_306:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 105:
			goto st307
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
	st307:
		if p++; p == pe {
			goto _test_eof307
		}
	st_case_307:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 112:
			goto st308
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
	st308:
		if p++; p == pe {
			goto _test_eof308
		}
	st_case_308:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st11
		case 116:
			goto st309
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
	st309:
		if p++; p == pe {
			goto _test_eof309
		}
	st_case_309:
		switch data[p] {
		case 58:
			goto st231
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
	st310:
		if p++; p == pe {
			goto _test_eof310
		}
	st_case_310:
		switch data[p] {
		case 43:
			goto st202
		case 45:
			goto st202
		case 60:
			goto st12
		case 95:
			goto st203
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st203
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st311:
		if p++; p == pe {
			goto _test_eof311
		}
	st_case_311:
		switch data[p] {
		case 43:
			goto st202
		case 45:
			goto st202
		case 60:
			goto st12
		case 95:
			goto st310
		case 97:
			goto st312
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st310
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st310
			}
		default:
			goto st310
		}
		goto st10
	st312:
		if p++; p == pe {
			goto _test_eof312
		}
	st_case_312:
		switch data[p] {
		case 43:
			goto st202
		case 45:
			goto st202
		case 60:
			goto st12
		case 95:
			goto st203
		case 115:
			goto st206
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st203
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st313:
		if p++; p == pe {
			goto _test_eof313
		}
	st_case_313:
		switch data[p] {
		case 43:
			goto st202
		case 45:
			goto st202
		case 60:
			goto st12
		case 95:
			goto st310
		case 104:
			goto st314
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st310
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st310
			}
		default:
			goto st310
		}
		goto st10
	st314:
		if p++; p == pe {
			goto _test_eof314
		}
	st_case_314:
		switch data[p] {
		case 43:
			goto st202
		case 45:
			goto st202
		case 60:
			goto st12
		case 95:
			goto st203
		case 97:
			goto st212
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st203
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st315:
		if p++; p == pe {
			goto _test_eof315
		}
	st_case_315:
		switch data[p] {
		case 43:
			goto st202
		case 45:
			goto st202
		case 60:
			goto st12
		case 95:
			goto st310
		case 97:
			goto st316
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st310
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st310
			}
		default:
			goto st310
		}
		goto st10
	st316:
		if p++; p == pe {
			goto _test_eof316
		}
	st_case_316:
		switch data[p] {
		case 43:
			goto st202
		case 45:
			goto st202
		case 60:
			goto st12
		case 95:
			goto st203
		case 116:
			goto st219
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st203
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st317:
		if p++; p == pe {
			goto _test_eof317
		}
	st_case_317:
		switch data[p] {
		case 43:
			goto st202
		case 45:
			goto st202
		case 60:
			goto st12
		case 95:
			goto st310
		case 97:
			goto st318
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st310
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st310
			}
		default:
			goto st310
		}
		goto st10
	st318:
		if p++; p == pe {
			goto _test_eof318
		}
	st_case_318:
		switch data[p] {
		case 43:
			goto st202
		case 45:
			goto st202
		case 60:
			goto st12
		case 95:
			goto st203
		case 118:
			goto st223
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st203
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st203
			}
		default:
			goto st203
		}
		goto st10
	st319:
		if p++; p == pe {
			goto _test_eof319
		}
	st_case_319:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st199
		case 97:
			goto st320
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st199
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st199
			}
		default:
			goto st11
		}
		goto st10
	st320:
		if p++; p == pe {
			goto _test_eof320
		}
	st_case_320:
		switch data[p] {
		case 47:
			goto st200
		case 60:
			goto st12
		case 95:
			goto st199
		case 115:
			goto st321
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st199
			}
		default:
			goto st11
		}
		goto st10
	st321:
		if p++; p == pe {
			goto _test_eof321
		}
	st_case_321:
		switch data[p] {
		case 47:
			goto st200
		case 60:
			goto st12
		case 95:
			goto st199
		case 101:
			goto st322
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st199
			}
		default:
			goto st11
		}
		goto st10
	st322:
		if p++; p == pe {
			goto _test_eof322
		}
	st_case_322:
		switch data[p] {
		case 47:
			goto st200
		case 54:
			goto st323
		case 60:
			goto st12
		case 95:
			goto st199
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st199
			}
		default:
			goto st11
		}
		goto st10
	st323:
		if p++; p == pe {
			goto _test_eof323
		}
	st_case_323:
		switch data[p] {
		case 47:
			goto st200
		case 52:
			goto st324
		case 60:
			goto st12
		case 95:
			goto st199
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st199
			}
		default:
			goto st11
		}
		goto st10
	st324:
		if p++; p == pe {
			goto _test_eof324
		}
	st_case_324:
		switch data[p] {
		case 47:
			goto tr362
		case 60:
			goto tr211
		case 95:
			goto st199
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st199
			}
		default:
			goto st11
		}
		goto tr210
tr362:
//line xss_170.rl:13

            return true
        
	goto st867
	st867:
		if p++; p == pe {
			goto _test_eof867
		}
	st_case_867:
//line xss_170.go:14704
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st201
		case 98:
			goto st311
		case 99:
			goto st313
		case 100:
			goto st315
		case 106:
			goto st317
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st201
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st201
			}
		default:
			goto st201
		}
		goto st10
	st325:
		if p++; p == pe {
			goto _test_eof325
		}
	st_case_325:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st199
		case 104:
			goto st326
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st199
			}
		default:
			goto st11
		}
		goto st10
	st326:
		if p++; p == pe {
			goto _test_eof326
		}
	st_case_326:
		switch data[p] {
		case 47:
			goto st200
		case 60:
			goto st12
		case 95:
			goto st199
		case 97:
			goto st327
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st199
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st199
			}
		default:
			goto st11
		}
		goto st10
	st327:
		if p++; p == pe {
			goto _test_eof327
		}
	st_case_327:
		switch data[p] {
		case 47:
			goto st200
		case 60:
			goto st12
		case 95:
			goto st199
		case 114:
			goto st328
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st199
			}
		default:
			goto st11
		}
		goto st10
	st328:
		if p++; p == pe {
			goto _test_eof328
		}
	st_case_328:
		switch data[p] {
		case 47:
			goto st200
		case 60:
			goto st12
		case 95:
			goto st199
		case 115:
			goto st329
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st199
			}
		default:
			goto st11
		}
		goto st10
	st329:
		if p++; p == pe {
			goto _test_eof329
		}
	st_case_329:
		switch data[p] {
		case 47:
			goto st200
		case 60:
			goto st12
		case 95:
			goto st199
		case 101:
			goto st330
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st199
			}
		default:
			goto st11
		}
		goto st10
	st330:
		if p++; p == pe {
			goto _test_eof330
		}
	st_case_330:
		switch data[p] {
		case 47:
			goto st200
		case 60:
			goto st12
		case 95:
			goto st199
		case 116:
			goto st331
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st199
			}
		default:
			goto st11
		}
		goto st10
	st331:
		if p++; p == pe {
			goto _test_eof331
		}
	st_case_331:
		switch data[p] {
		case 47:
			goto st200
		case 60:
			goto st12
		case 61:
			goto st192
		case 95:
			goto st199
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st199
			}
		default:
			goto st11
		}
		goto st10
	st332:
		if p++; p == pe {
			goto _test_eof332
		}
	st_case_332:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st199
		case 97:
			goto st333
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st199
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st199
			}
		default:
			goto st11
		}
		goto st10
	st333:
		if p++; p == pe {
			goto _test_eof333
		}
	st_case_333:
		switch data[p] {
		case 47:
			goto st200
		case 60:
			goto st12
		case 95:
			goto st199
		case 116:
			goto st334
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st199
			}
		default:
			goto st11
		}
		goto st10
	st334:
		if p++; p == pe {
			goto _test_eof334
		}
	st_case_334:
		switch data[p] {
		case 47:
			goto st200
		case 60:
			goto st12
		case 95:
			goto st199
		case 97:
			goto st335
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st199
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st199
			}
		default:
			goto st11
		}
		goto st10
	st335:
		if p++; p == pe {
			goto _test_eof335
		}
	st_case_335:
		switch data[p] {
		case 47:
			goto st200
		case 58:
			goto st197
		case 60:
			goto st12
		case 95:
			goto st199
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st199
			}
		default:
			goto st11
		}
		goto st10
	st336:
		if p++; p == pe {
			goto _test_eof336
		}
	st_case_336:
		switch data[p] {
		case 60:
			goto st12
		case 95:
			goto st199
		case 97:
			goto st337
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st199
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st199
			}
		default:
			goto st11
		}
		goto st10
	st337:
		if p++; p == pe {
			goto _test_eof337
		}
	st_case_337:
		switch data[p] {
		case 47:
			goto st200
		case 60:
			goto st12
		case 95:
			goto st199
		case 118:
			goto st338
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st199
			}
		default:
			goto st11
		}
		goto st10
	st338:
		if p++; p == pe {
			goto _test_eof338
		}
	st_case_338:
		switch data[p] {
		case 47:
			goto st200
		case 60:
			goto st12
		case 95:
			goto st199
		case 97:
			goto st339
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st199
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st199
			}
		default:
			goto st11
		}
		goto st10
	st339:
		if p++; p == pe {
			goto _test_eof339
		}
	st_case_339:
		switch data[p] {
		case 47:
			goto st200
		case 60:
			goto st12
		case 95:
			goto st199
		case 115:
			goto st340
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st199
			}
		default:
			goto st11
		}
		goto st10
	st340:
		if p++; p == pe {
			goto _test_eof340
		}
	st_case_340:
		switch data[p] {
		case 47:
			goto st200
		case 60:
			goto st12
		case 95:
			goto st199
		case 99:
			goto st341
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st199
			}
		default:
			goto st11
		}
		goto st10
	st341:
		if p++; p == pe {
			goto _test_eof341
		}
	st_case_341:
		switch data[p] {
		case 47:
			goto st200
		case 60:
			goto st12
		case 95:
			goto st199
		case 114:
			goto st342
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st199
			}
		default:
			goto st11
		}
		goto st10
	st342:
		if p++; p == pe {
			goto _test_eof342
		}
	st_case_342:
		switch data[p] {
		case 47:
			goto st200
		case 60:
			goto st12
		case 95:
			goto st199
		case 105:
			goto st343
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st199
			}
		default:
			goto st11
		}
		goto st10
	st343:
		if p++; p == pe {
			goto _test_eof343
		}
	st_case_343:
		switch data[p] {
		case 47:
			goto st200
		case 60:
			goto st12
		case 95:
			goto st199
		case 112:
			goto st344
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st199
			}
		default:
			goto st11
		}
		goto st10
	st344:
		if p++; p == pe {
			goto _test_eof344
		}
	st_case_344:
		switch data[p] {
		case 47:
			goto st200
		case 60:
			goto st12
		case 95:
			goto st199
		case 116:
			goto st345
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st199
			}
		default:
			goto st11
		}
		goto st10
	st345:
		if p++; p == pe {
			goto _test_eof345
		}
	st_case_345:
		switch data[p] {
		case 47:
			goto st200
		case 58:
			goto st231
		case 60:
			goto st12
		case 95:
			goto st199
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st199
			}
		default:
			goto st11
		}
		goto st10
	st346:
		if p++; p == pe {
			goto _test_eof346
		}
	st_case_346:
		switch data[p] {
		case 59:
			goto st12
		case 95:
			goto st347
		case 100:
			goto st350
		case 106:
			goto st440
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st347
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st347
			}
		default:
			goto st347
		}
		goto st346
	st347:
		if p++; p == pe {
			goto _test_eof347
		}
	st_case_347:
		switch data[p] {
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st347
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st347
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st347
			}
		default:
			goto st347
		}
		goto st348
	st348:
		if p++; p == pe {
			goto _test_eof348
		}
	st_case_348:
		switch data[p] {
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st347
		case 100:
			goto st350
		case 106:
			goto st440
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st347
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st347
			}
		default:
			goto st347
		}
		goto st348
	st349:
		if p++; p == pe {
			goto _test_eof349
		}
	st_case_349:
		switch data[p] {
		case 59:
			goto tr37
		case 62:
			goto tr387
		case 95:
			goto st347
		case 100:
			goto st350
		case 106:
			goto st440
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st347
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st347
			}
		default:
			goto st347
		}
		goto tr386
tr386:
//line xss_170.rl:13

            return true
        
	goto st868
	st868:
		if p++; p == pe {
			goto _test_eof868
		}
	st_case_868:
//line xss_170.go:15439
		switch data[p] {
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st347
		case 100:
			goto st350
		case 106:
			goto st440
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st347
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st347
			}
		default:
			goto st347
		}
		goto st348
	st350:
		if p++; p == pe {
			goto _test_eof350
		}
	st_case_350:
		switch data[p] {
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st347
		case 97:
			goto st351
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st347
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st347
			}
		default:
			goto st347
		}
		goto st348
	st351:
		if p++; p == pe {
			goto _test_eof351
		}
	st_case_351:
		switch data[p] {
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st347
		case 116:
			goto st352
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st347
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st347
			}
		default:
			goto st347
		}
		goto st348
	st352:
		if p++; p == pe {
			goto _test_eof352
		}
	st_case_352:
		switch data[p] {
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st347
		case 97:
			goto st353
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st347
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st347
			}
		default:
			goto st347
		}
		goto st348
	st353:
		if p++; p == pe {
			goto _test_eof353
		}
	st_case_353:
		switch data[p] {
		case 58:
			goto st354
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st347
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st347
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st347
			}
		default:
			goto st347
		}
		goto st348
	st354:
		if p++; p == pe {
			goto _test_eof354
		}
	st_case_354:
		switch data[p] {
		case 44:
			goto st349
		case 59:
			goto st15
		case 62:
			goto st349
		case 95:
			goto st355
		case 100:
			goto st426
		case 106:
			goto st430
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st347
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st355
			}
		default:
			goto st347
		}
		goto st348
	st355:
		if p++; p == pe {
			goto _test_eof355
		}
	st_case_355:
		switch data[p] {
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st356
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st356
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st356
			}
		default:
			goto st347
		}
		goto st348
	st356:
		if p++; p == pe {
			goto _test_eof356
		}
	st_case_356:
		switch data[p] {
		case 47:
			goto st357
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st356
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st356
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st356
			}
		default:
			goto st347
		}
		goto st348
	st357:
		if p++; p == pe {
			goto _test_eof357
		}
	st_case_357:
		switch data[p] {
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st358
		case 100:
			goto st422
		case 106:
			goto st424
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st358
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st358
			}
		default:
			goto st358
		}
		goto st348
	st358:
		if p++; p == pe {
			goto _test_eof358
		}
	st_case_358:
		switch data[p] {
		case 43:
			goto st359
		case 45:
			goto st359
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st421
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st421
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st421
			}
		default:
			goto st421
		}
		goto st348
	st359:
		if p++; p == pe {
			goto _test_eof359
		}
	st_case_359:
		switch data[p] {
		case 43:
			goto st359
		case 45:
			goto st359
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st360
		case 100:
			goto st361
		case 106:
			goto st365
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st360
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st360
			}
		default:
			goto st360
		}
		goto st348
	st360:
		if p++; p == pe {
			goto _test_eof360
		}
	st_case_360:
		switch data[p] {
		case 44:
			goto st349
		case 59:
			goto st15
		case 62:
			goto st349
		case 95:
			goto st360
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st359
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st360
				}
			case data[p] >= 65:
				goto st360
			}
		default:
			goto st360
		}
		goto st348
	st361:
		if p++; p == pe {
			goto _test_eof361
		}
	st_case_361:
		switch data[p] {
		case 44:
			goto st349
		case 59:
			goto st15
		case 62:
			goto st349
		case 95:
			goto st360
		case 97:
			goto st362
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st359
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st360
				}
			case data[p] >= 65:
				goto st360
			}
		default:
			goto st360
		}
		goto st348
	st362:
		if p++; p == pe {
			goto _test_eof362
		}
	st_case_362:
		switch data[p] {
		case 44:
			goto st349
		case 59:
			goto st15
		case 62:
			goto st349
		case 95:
			goto st360
		case 116:
			goto st363
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st359
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st360
				}
			case data[p] >= 65:
				goto st360
			}
		default:
			goto st360
		}
		goto st348
	st363:
		if p++; p == pe {
			goto _test_eof363
		}
	st_case_363:
		switch data[p] {
		case 44:
			goto st349
		case 59:
			goto st15
		case 62:
			goto st349
		case 95:
			goto st360
		case 97:
			goto st364
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st359
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st360
				}
			case data[p] >= 65:
				goto st360
			}
		default:
			goto st360
		}
		goto st348
	st364:
		if p++; p == pe {
			goto _test_eof364
		}
	st_case_364:
		switch data[p] {
		case 44:
			goto st349
		case 58:
			goto st354
		case 59:
			goto st15
		case 62:
			goto st349
		case 95:
			goto st360
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st359
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st360
				}
			case data[p] >= 65:
				goto st360
			}
		default:
			goto st360
		}
		goto st348
	st365:
		if p++; p == pe {
			goto _test_eof365
		}
	st_case_365:
		switch data[p] {
		case 44:
			goto st349
		case 59:
			goto st15
		case 62:
			goto st349
		case 95:
			goto st360
		case 97:
			goto st366
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st359
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st360
				}
			case data[p] >= 65:
				goto st360
			}
		default:
			goto st360
		}
		goto st348
	st366:
		if p++; p == pe {
			goto _test_eof366
		}
	st_case_366:
		switch data[p] {
		case 44:
			goto st349
		case 59:
			goto st15
		case 62:
			goto st349
		case 95:
			goto st360
		case 118:
			goto st367
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st359
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st360
				}
			case data[p] >= 65:
				goto st360
			}
		default:
			goto st360
		}
		goto st348
	st367:
		if p++; p == pe {
			goto _test_eof367
		}
	st_case_367:
		switch data[p] {
		case 44:
			goto st349
		case 59:
			goto st15
		case 62:
			goto st349
		case 95:
			goto st360
		case 97:
			goto st368
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st359
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st360
				}
			case data[p] >= 65:
				goto st360
			}
		default:
			goto st360
		}
		goto st348
	st368:
		if p++; p == pe {
			goto _test_eof368
		}
	st_case_368:
		switch data[p] {
		case 44:
			goto st349
		case 59:
			goto st15
		case 62:
			goto st349
		case 95:
			goto st360
		case 115:
			goto st369
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st359
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st360
				}
			case data[p] >= 65:
				goto st360
			}
		default:
			goto st360
		}
		goto st348
	st369:
		if p++; p == pe {
			goto _test_eof369
		}
	st_case_369:
		switch data[p] {
		case 44:
			goto st349
		case 59:
			goto st15
		case 62:
			goto st349
		case 95:
			goto st360
		case 99:
			goto st370
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st359
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st360
				}
			case data[p] >= 65:
				goto st360
			}
		default:
			goto st360
		}
		goto st348
	st370:
		if p++; p == pe {
			goto _test_eof370
		}
	st_case_370:
		switch data[p] {
		case 44:
			goto st349
		case 59:
			goto st15
		case 62:
			goto st349
		case 95:
			goto st360
		case 114:
			goto st371
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st359
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st360
				}
			case data[p] >= 65:
				goto st360
			}
		default:
			goto st360
		}
		goto st348
	st371:
		if p++; p == pe {
			goto _test_eof371
		}
	st_case_371:
		switch data[p] {
		case 44:
			goto st349
		case 59:
			goto st15
		case 62:
			goto st349
		case 95:
			goto st360
		case 105:
			goto st372
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st359
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st360
				}
			case data[p] >= 65:
				goto st360
			}
		default:
			goto st360
		}
		goto st348
	st372:
		if p++; p == pe {
			goto _test_eof372
		}
	st_case_372:
		switch data[p] {
		case 44:
			goto st349
		case 59:
			goto st15
		case 62:
			goto st349
		case 95:
			goto st360
		case 112:
			goto st373
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st359
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st360
				}
			case data[p] >= 65:
				goto st360
			}
		default:
			goto st360
		}
		goto st348
	st373:
		if p++; p == pe {
			goto _test_eof373
		}
	st_case_373:
		switch data[p] {
		case 44:
			goto st349
		case 59:
			goto st15
		case 62:
			goto st349
		case 95:
			goto st360
		case 116:
			goto st374
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st359
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st360
				}
			case data[p] >= 65:
				goto st360
			}
		default:
			goto st360
		}
		goto st348
	st374:
		if p++; p == pe {
			goto _test_eof374
		}
	st_case_374:
		switch data[p] {
		case 44:
			goto st349
		case 58:
			goto st375
		case 59:
			goto st15
		case 62:
			goto st349
		case 95:
			goto st360
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st359
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st360
				}
			case data[p] >= 65:
				goto st360
			}
		default:
			goto st360
		}
		goto st348
	st375:
		if p++; p == pe {
			goto _test_eof375
		}
	st_case_375:
		switch data[p] {
		case 59:
			goto st68
		case 62:
			goto st377
		case 92:
			goto st420
		case 95:
			goto st378
		case 100:
			goto st380
		case 110:
			goto st414
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st378
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st378
			}
		default:
			goto st378
		}
		goto st376
	st376:
		if p++; p == pe {
			goto _test_eof376
		}
	st_case_376:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 46:
			goto st377
		case 59:
			goto st68
		case 91:
			goto st377
		case 92:
			goto st379
		case 95:
			goto st378
		case 100:
			goto st380
		case 110:
			goto st414
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st378
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st378
				}
			case data[p] >= 65:
				goto st378
			}
		default:
			goto st377
		}
		goto st376
	st377:
		if p++; p == pe {
			goto _test_eof377
		}
	st_case_377:
		switch data[p] {
		case 32:
			goto tr426
		case 40:
			goto tr426
		case 46:
			goto tr426
		case 59:
			goto tr106
		case 91:
			goto tr426
		case 92:
			goto tr427
		case 95:
			goto st378
		case 100:
			goto st380
		case 110:
			goto st414
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st378
				}
			case data[p] >= 9:
				goto tr426
			}
		case data[p] > 62:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st378
				}
			case data[p] >= 65:
				goto st378
			}
		default:
			goto tr426
		}
		goto tr425
tr425:
//line xss_170.rl:13

            return true
        
	goto st869
	st869:
		if p++; p == pe {
			goto _test_eof869
		}
	st_case_869:
//line xss_170.go:16419
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 46:
			goto st377
		case 59:
			goto st68
		case 91:
			goto st377
		case 92:
			goto st379
		case 95:
			goto st378
		case 100:
			goto st380
		case 110:
			goto st414
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st378
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st378
				}
			case data[p] >= 65:
				goto st378
			}
		default:
			goto st377
		}
		goto st376
	st378:
		if p++; p == pe {
			goto _test_eof378
		}
	st_case_378:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 46:
			goto st377
		case 59:
			goto st68
		case 95:
			goto st378
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st378
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st378
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st378
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st379:
		if p++; p == pe {
			goto _test_eof379
		}
	st_case_379:
		switch data[p] {
		case 32:
			goto tr426
		case 40:
			goto tr426
		case 46:
			goto tr426
		case 59:
			goto tr106
		case 91:
			goto tr426
		case 92:
			goto tr427
		case 95:
			goto st378
		case 100:
			goto st380
		case 110:
			goto st414
		case 117:
			goto st418
		case 120:
			goto st418
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st378
				}
			case data[p] >= 9:
				goto tr426
			}
		case data[p] > 62:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st378
				}
			case data[p] >= 65:
				goto st378
			}
		default:
			goto tr426
		}
		goto tr425
tr426:
//line xss_170.rl:13

            return true
        
	goto st870
	st870:
		if p++; p == pe {
			goto _test_eof870
		}
	st_case_870:
//line xss_170.go:16570
		switch data[p] {
		case 32:
			goto tr426
		case 40:
			goto tr426
		case 46:
			goto tr426
		case 59:
			goto tr106
		case 91:
			goto tr426
		case 92:
			goto tr427
		case 95:
			goto st378
		case 100:
			goto st380
		case 110:
			goto st414
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st378
				}
			case data[p] >= 9:
				goto tr426
			}
		case data[p] > 62:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st378
				}
			case data[p] >= 65:
				goto st378
			}
		default:
			goto tr426
		}
		goto tr425
tr427:
//line xss_170.rl:13

            return true
        
	goto st871
	st871:
		if p++; p == pe {
			goto _test_eof871
		}
	st_case_871:
//line xss_170.go:16625
		switch data[p] {
		case 32:
			goto tr426
		case 40:
			goto tr426
		case 46:
			goto tr426
		case 59:
			goto tr106
		case 91:
			goto tr426
		case 92:
			goto tr427
		case 95:
			goto st378
		case 100:
			goto st380
		case 110:
			goto st414
		case 117:
			goto st418
		case 120:
			goto st418
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st378
				}
			case data[p] >= 9:
				goto tr426
			}
		case data[p] > 62:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st378
				}
			case data[p] >= 65:
				goto st378
			}
		default:
			goto tr426
		}
		goto tr425
	st380:
		if p++; p == pe {
			goto _test_eof380
		}
	st_case_380:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 46:
			goto st377
		case 59:
			goto st68
		case 95:
			goto st378
		case 97:
			goto st381
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st378
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st378
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st378
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st381:
		if p++; p == pe {
			goto _test_eof381
		}
	st_case_381:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 46:
			goto st377
		case 59:
			goto st68
		case 95:
			goto st378
		case 116:
			goto st382
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st378
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st378
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st378
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st382:
		if p++; p == pe {
			goto _test_eof382
		}
	st_case_382:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 46:
			goto st377
		case 59:
			goto st68
		case 95:
			goto st378
		case 97:
			goto st383
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st378
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st378
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st378
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st383:
		if p++; p == pe {
			goto _test_eof383
		}
	st_case_383:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 46:
			goto st377
		case 58:
			goto st384
		case 59:
			goto st68
		case 95:
			goto st378
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st378
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st378
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st378
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st384:
		if p++; p == pe {
			goto _test_eof384
		}
	st_case_384:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 44:
			goto st377
		case 46:
			goto st377
		case 59:
			goto st69
		case 91:
			goto st377
		case 92:
			goto st379
		case 95:
			goto st385
		case 100:
			goto st405
		case 110:
			goto st409
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st378
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st385
				}
			case data[p] >= 65:
				goto st378
			}
		default:
			goto st377
		}
		goto st376
	st385:
		if p++; p == pe {
			goto _test_eof385
		}
	st_case_385:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 46:
			goto st377
		case 59:
			goto st68
		case 95:
			goto st386
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st386
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st378
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st386
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st386:
		if p++; p == pe {
			goto _test_eof386
		}
	st_case_386:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 46:
			goto st377
		case 47:
			goto st387
		case 59:
			goto st68
		case 95:
			goto st386
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st386
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st378
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st386
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st387:
		if p++; p == pe {
			goto _test_eof387
		}
	st_case_387:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 46:
			goto st377
		case 59:
			goto st68
		case 91:
			goto st377
		case 92:
			goto st379
		case 95:
			goto st388
		case 100:
			goto st401
		case 110:
			goto st403
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st388
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st388
				}
			case data[p] >= 65:
				goto st388
			}
		default:
			goto st377
		}
		goto st376
	st388:
		if p++; p == pe {
			goto _test_eof388
		}
	st_case_388:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 43:
			goto st389
		case 45:
			goto st389
		case 46:
			goto st377
		case 59:
			goto st68
		case 95:
			goto st400
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st400
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st400
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st400
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st389:
		if p++; p == pe {
			goto _test_eof389
		}
	st_case_389:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 43:
			goto st389
		case 45:
			goto st389
		case 46:
			goto st377
		case 59:
			goto st68
		case 91:
			goto st377
		case 92:
			goto st379
		case 95:
			goto st390
		case 100:
			goto st391
		case 110:
			goto st395
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st390
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st390
				}
			case data[p] >= 65:
				goto st390
			}
		default:
			goto st377
		}
		goto st376
	st390:
		if p++; p == pe {
			goto _test_eof390
		}
	st_case_390:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 44:
			goto st377
		case 46:
			goto st377
		case 59:
			goto st69
		case 95:
			goto st390
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st377
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st390
				}
			default:
				goto st389
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st390
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st390
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st391:
		if p++; p == pe {
			goto _test_eof391
		}
	st_case_391:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 44:
			goto st377
		case 46:
			goto st377
		case 59:
			goto st69
		case 95:
			goto st390
		case 97:
			goto st392
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st377
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st390
				}
			default:
				goto st389
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st390
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st390
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st392:
		if p++; p == pe {
			goto _test_eof392
		}
	st_case_392:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 44:
			goto st377
		case 46:
			goto st377
		case 59:
			goto st69
		case 95:
			goto st390
		case 116:
			goto st393
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st377
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st390
				}
			default:
				goto st389
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st390
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st390
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st393:
		if p++; p == pe {
			goto _test_eof393
		}
	st_case_393:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 44:
			goto st377
		case 46:
			goto st377
		case 59:
			goto st69
		case 95:
			goto st390
		case 97:
			goto st394
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st377
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st390
				}
			default:
				goto st389
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st390
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st390
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st394:
		if p++; p == pe {
			goto _test_eof394
		}
	st_case_394:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 44:
			goto st377
		case 46:
			goto st377
		case 58:
			goto st384
		case 59:
			goto st69
		case 95:
			goto st390
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st377
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st390
				}
			default:
				goto st389
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st390
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st390
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st395:
		if p++; p == pe {
			goto _test_eof395
		}
	st_case_395:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 44:
			goto st377
		case 46:
			goto st377
		case 59:
			goto st69
		case 95:
			goto st390
		case 97:
			goto st396
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st377
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st390
				}
			default:
				goto st389
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st390
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st390
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st396:
		if p++; p == pe {
			goto _test_eof396
		}
	st_case_396:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 44:
			goto st377
		case 46:
			goto st377
		case 59:
			goto st69
		case 95:
			goto st390
		case 109:
			goto st397
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st377
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st390
				}
			default:
				goto st389
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st390
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st390
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st397:
		if p++; p == pe {
			goto _test_eof397
		}
	st_case_397:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 44:
			goto st377
		case 46:
			goto st377
		case 59:
			goto st69
		case 95:
			goto st390
		case 101:
			goto st398
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st377
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st390
				}
			default:
				goto st389
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st390
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st390
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st398:
		if p++; p == pe {
			goto _test_eof398
		}
	st_case_398:
		switch data[p] {
		case 43:
			goto st399
		case 45:
			goto st399
		case 59:
			goto st69
		case 95:
			goto st390
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st390
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st390
			}
		default:
			goto st390
		}
		goto st377
	st399:
		if p++; p == pe {
			goto _test_eof399
		}
	st_case_399:
		switch data[p] {
		case 32:
			goto tr426
		case 40:
			goto tr426
		case 43:
			goto tr453
		case 45:
			goto tr453
		case 46:
			goto tr426
		case 59:
			goto tr106
		case 91:
			goto tr426
		case 92:
			goto tr427
		case 95:
			goto st390
		case 100:
			goto st391
		case 110:
			goto st395
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st390
				}
			case data[p] >= 9:
				goto tr426
			}
		case data[p] > 62:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st390
				}
			case data[p] >= 65:
				goto st390
			}
		default:
			goto tr426
		}
		goto tr425
tr453:
//line xss_170.rl:13

            return true
        
	goto st872
	st872:
		if p++; p == pe {
			goto _test_eof872
		}
	st_case_872:
//line xss_170.go:17650
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 43:
			goto st389
		case 45:
			goto st389
		case 46:
			goto st377
		case 59:
			goto st68
		case 91:
			goto st377
		case 92:
			goto st379
		case 95:
			goto st390
		case 100:
			goto st391
		case 110:
			goto st395
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st390
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st390
				}
			case data[p] >= 65:
				goto st390
			}
		default:
			goto st377
		}
		goto st376
	st400:
		if p++; p == pe {
			goto _test_eof400
		}
	st_case_400:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 43:
			goto st389
		case 45:
			goto st389
		case 46:
			goto st377
		case 59:
			goto st68
		case 95:
			goto st390
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st390
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st390
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st390
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st401:
		if p++; p == pe {
			goto _test_eof401
		}
	st_case_401:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 43:
			goto st389
		case 45:
			goto st389
		case 46:
			goto st377
		case 59:
			goto st68
		case 95:
			goto st400
		case 97:
			goto st402
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st400
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st400
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st400
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st402:
		if p++; p == pe {
			goto _test_eof402
		}
	st_case_402:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 43:
			goto st389
		case 45:
			goto st389
		case 46:
			goto st377
		case 59:
			goto st68
		case 95:
			goto st390
		case 116:
			goto st393
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st390
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st390
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st390
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st403:
		if p++; p == pe {
			goto _test_eof403
		}
	st_case_403:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 43:
			goto st389
		case 45:
			goto st389
		case 46:
			goto st377
		case 59:
			goto st68
		case 95:
			goto st400
		case 97:
			goto st404
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st400
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st400
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st400
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st404:
		if p++; p == pe {
			goto _test_eof404
		}
	st_case_404:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 43:
			goto st389
		case 45:
			goto st389
		case 46:
			goto st377
		case 59:
			goto st68
		case 95:
			goto st390
		case 109:
			goto st397
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st390
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st390
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st390
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st405:
		if p++; p == pe {
			goto _test_eof405
		}
	st_case_405:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 46:
			goto st377
		case 59:
			goto st68
		case 95:
			goto st386
		case 97:
			goto st406
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st386
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st378
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st386
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st406:
		if p++; p == pe {
			goto _test_eof406
		}
	st_case_406:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 46:
			goto st377
		case 47:
			goto st387
		case 59:
			goto st68
		case 95:
			goto st386
		case 116:
			goto st407
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st386
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st378
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st386
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st407:
		if p++; p == pe {
			goto _test_eof407
		}
	st_case_407:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 46:
			goto st377
		case 47:
			goto st387
		case 59:
			goto st68
		case 95:
			goto st386
		case 97:
			goto st408
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st386
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st378
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st386
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st408:
		if p++; p == pe {
			goto _test_eof408
		}
	st_case_408:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 46:
			goto st377
		case 47:
			goto st387
		case 58:
			goto st384
		case 59:
			goto st68
		case 95:
			goto st386
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st386
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st378
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st386
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st409:
		if p++; p == pe {
			goto _test_eof409
		}
	st_case_409:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 46:
			goto st377
		case 59:
			goto st68
		case 95:
			goto st386
		case 97:
			goto st410
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st386
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st378
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st386
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st410:
		if p++; p == pe {
			goto _test_eof410
		}
	st_case_410:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 46:
			goto st377
		case 47:
			goto st387
		case 59:
			goto st68
		case 95:
			goto st386
		case 109:
			goto st411
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st386
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st378
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st386
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st411:
		if p++; p == pe {
			goto _test_eof411
		}
	st_case_411:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 46:
			goto st377
		case 47:
			goto st387
		case 59:
			goto st68
		case 95:
			goto st386
		case 101:
			goto st412
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st386
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st378
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st386
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st412:
		if p++; p == pe {
			goto _test_eof412
		}
	st_case_412:
		switch data[p] {
		case 47:
			goto st413
		case 59:
			goto st69
		case 95:
			goto st386
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st386
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st386
			}
		default:
			goto st378
		}
		goto st377
	st413:
		if p++; p == pe {
			goto _test_eof413
		}
	st_case_413:
		switch data[p] {
		case 32:
			goto tr426
		case 40:
			goto tr426
		case 46:
			goto tr426
		case 59:
			goto tr106
		case 91:
			goto tr426
		case 92:
			goto tr427
		case 95:
			goto st388
		case 100:
			goto st401
		case 110:
			goto st403
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st388
				}
			case data[p] >= 9:
				goto tr426
			}
		case data[p] > 62:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st388
				}
			case data[p] >= 65:
				goto st388
			}
		default:
			goto tr426
		}
		goto tr425
	st414:
		if p++; p == pe {
			goto _test_eof414
		}
	st_case_414:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 46:
			goto st377
		case 59:
			goto st68
		case 95:
			goto st378
		case 97:
			goto st415
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st378
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st378
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st378
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st415:
		if p++; p == pe {
			goto _test_eof415
		}
	st_case_415:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 46:
			goto st377
		case 59:
			goto st68
		case 95:
			goto st378
		case 109:
			goto st416
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st378
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st378
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st378
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st416:
		if p++; p == pe {
			goto _test_eof416
		}
	st_case_416:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 46:
			goto st377
		case 59:
			goto st68
		case 95:
			goto st378
		case 101:
			goto st417
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st378
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st378
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st378
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st417:
		if p++; p == pe {
			goto _test_eof417
		}
	st_case_417:
		switch data[p] {
		case 59:
			goto st69
		case 95:
			goto st378
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st378
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st378
			}
		default:
			goto st378
		}
		goto st377
	st418:
		if p++; p == pe {
			goto _test_eof418
		}
	st_case_418:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 46:
			goto st377
		case 59:
			goto st68
		case 95:
			goto st378
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st419
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st378
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st378
				}
			default:
				goto st377
			}
		default:
			goto st377
		}
		goto st376
	st419:
		if p++; p == pe {
			goto _test_eof419
		}
	st_case_419:
		switch data[p] {
		case 32:
			goto tr426
		case 40:
			goto tr426
		case 46:
			goto tr426
		case 59:
			goto tr106
		case 95:
			goto st378
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st419
				}
			case data[p] >= 9:
				goto tr426
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st378
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st378
				}
			default:
				goto tr426
			}
		default:
			goto tr426
		}
		goto tr425
	st420:
		if p++; p == pe {
			goto _test_eof420
		}
	st_case_420:
		switch data[p] {
		case 32:
			goto st377
		case 40:
			goto st377
		case 46:
			goto st377
		case 59:
			goto st68
		case 91:
			goto st377
		case 92:
			goto st379
		case 95:
			goto st378
		case 100:
			goto st380
		case 110:
			goto st414
		case 117:
			goto st418
		case 120:
			goto st418
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st378
				}
			case data[p] >= 9:
				goto st377
			}
		case data[p] > 62:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st378
				}
			case data[p] >= 65:
				goto st378
			}
		default:
			goto st377
		}
		goto st376
	st421:
		if p++; p == pe {
			goto _test_eof421
		}
	st_case_421:
		switch data[p] {
		case 43:
			goto st359
		case 45:
			goto st359
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st360
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st360
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st360
			}
		default:
			goto st360
		}
		goto st348
	st422:
		if p++; p == pe {
			goto _test_eof422
		}
	st_case_422:
		switch data[p] {
		case 43:
			goto st359
		case 45:
			goto st359
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st421
		case 97:
			goto st423
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st421
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st421
			}
		default:
			goto st421
		}
		goto st348
	st423:
		if p++; p == pe {
			goto _test_eof423
		}
	st_case_423:
		switch data[p] {
		case 43:
			goto st359
		case 45:
			goto st359
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st360
		case 116:
			goto st363
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st360
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st360
			}
		default:
			goto st360
		}
		goto st348
	st424:
		if p++; p == pe {
			goto _test_eof424
		}
	st_case_424:
		switch data[p] {
		case 43:
			goto st359
		case 45:
			goto st359
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st421
		case 97:
			goto st425
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st421
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st421
			}
		default:
			goto st421
		}
		goto st348
	st425:
		if p++; p == pe {
			goto _test_eof425
		}
	st_case_425:
		switch data[p] {
		case 43:
			goto st359
		case 45:
			goto st359
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st360
		case 118:
			goto st367
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st360
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st360
			}
		default:
			goto st360
		}
		goto st348
	st426:
		if p++; p == pe {
			goto _test_eof426
		}
	st_case_426:
		switch data[p] {
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st356
		case 97:
			goto st427
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st356
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st356
			}
		default:
			goto st347
		}
		goto st348
	st427:
		if p++; p == pe {
			goto _test_eof427
		}
	st_case_427:
		switch data[p] {
		case 47:
			goto st357
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st356
		case 116:
			goto st428
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st356
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st356
			}
		default:
			goto st347
		}
		goto st348
	st428:
		if p++; p == pe {
			goto _test_eof428
		}
	st_case_428:
		switch data[p] {
		case 47:
			goto st357
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st356
		case 97:
			goto st429
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st356
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st356
			}
		default:
			goto st347
		}
		goto st348
	st429:
		if p++; p == pe {
			goto _test_eof429
		}
	st_case_429:
		switch data[p] {
		case 47:
			goto st357
		case 58:
			goto st354
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st356
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st356
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st356
			}
		default:
			goto st347
		}
		goto st348
	st430:
		if p++; p == pe {
			goto _test_eof430
		}
	st_case_430:
		switch data[p] {
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st356
		case 97:
			goto st431
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st356
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st356
			}
		default:
			goto st347
		}
		goto st348
	st431:
		if p++; p == pe {
			goto _test_eof431
		}
	st_case_431:
		switch data[p] {
		case 47:
			goto st357
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st356
		case 118:
			goto st432
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st356
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st356
			}
		default:
			goto st347
		}
		goto st348
	st432:
		if p++; p == pe {
			goto _test_eof432
		}
	st_case_432:
		switch data[p] {
		case 47:
			goto st357
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st356
		case 97:
			goto st433
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st356
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st356
			}
		default:
			goto st347
		}
		goto st348
	st433:
		if p++; p == pe {
			goto _test_eof433
		}
	st_case_433:
		switch data[p] {
		case 47:
			goto st357
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st356
		case 115:
			goto st434
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st356
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st356
			}
		default:
			goto st347
		}
		goto st348
	st434:
		if p++; p == pe {
			goto _test_eof434
		}
	st_case_434:
		switch data[p] {
		case 47:
			goto st357
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st356
		case 99:
			goto st435
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st356
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st356
			}
		default:
			goto st347
		}
		goto st348
	st435:
		if p++; p == pe {
			goto _test_eof435
		}
	st_case_435:
		switch data[p] {
		case 47:
			goto st357
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st356
		case 114:
			goto st436
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st356
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st356
			}
		default:
			goto st347
		}
		goto st348
	st436:
		if p++; p == pe {
			goto _test_eof436
		}
	st_case_436:
		switch data[p] {
		case 47:
			goto st357
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st356
		case 105:
			goto st437
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st356
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st356
			}
		default:
			goto st347
		}
		goto st348
	st437:
		if p++; p == pe {
			goto _test_eof437
		}
	st_case_437:
		switch data[p] {
		case 47:
			goto st357
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st356
		case 112:
			goto st438
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st356
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st356
			}
		default:
			goto st347
		}
		goto st348
	st438:
		if p++; p == pe {
			goto _test_eof438
		}
	st_case_438:
		switch data[p] {
		case 47:
			goto st357
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st356
		case 116:
			goto st439
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st356
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st356
			}
		default:
			goto st347
		}
		goto st348
	st439:
		if p++; p == pe {
			goto _test_eof439
		}
	st_case_439:
		switch data[p] {
		case 47:
			goto st357
		case 58:
			goto st375
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st356
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st356
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st356
			}
		default:
			goto st347
		}
		goto st348
	st440:
		if p++; p == pe {
			goto _test_eof440
		}
	st_case_440:
		switch data[p] {
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st347
		case 97:
			goto st441
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st347
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st347
			}
		default:
			goto st347
		}
		goto st348
	st441:
		if p++; p == pe {
			goto _test_eof441
		}
	st_case_441:
		switch data[p] {
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st347
		case 118:
			goto st442
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st347
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st347
			}
		default:
			goto st347
		}
		goto st348
	st442:
		if p++; p == pe {
			goto _test_eof442
		}
	st_case_442:
		switch data[p] {
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st347
		case 97:
			goto st443
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st347
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st347
			}
		default:
			goto st347
		}
		goto st348
	st443:
		if p++; p == pe {
			goto _test_eof443
		}
	st_case_443:
		switch data[p] {
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st347
		case 115:
			goto st444
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st347
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st347
			}
		default:
			goto st347
		}
		goto st348
	st444:
		if p++; p == pe {
			goto _test_eof444
		}
	st_case_444:
		switch data[p] {
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st347
		case 99:
			goto st445
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st347
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st347
			}
		default:
			goto st347
		}
		goto st348
	st445:
		if p++; p == pe {
			goto _test_eof445
		}
	st_case_445:
		switch data[p] {
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st347
		case 114:
			goto st446
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st347
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st347
			}
		default:
			goto st347
		}
		goto st348
	st446:
		if p++; p == pe {
			goto _test_eof446
		}
	st_case_446:
		switch data[p] {
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st347
		case 105:
			goto st447
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st347
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st347
			}
		default:
			goto st347
		}
		goto st348
	st447:
		if p++; p == pe {
			goto _test_eof447
		}
	st_case_447:
		switch data[p] {
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st347
		case 112:
			goto st448
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st347
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st347
			}
		default:
			goto st347
		}
		goto st348
	st448:
		if p++; p == pe {
			goto _test_eof448
		}
	st_case_448:
		switch data[p] {
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st347
		case 116:
			goto st449
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st347
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st347
			}
		default:
			goto st347
		}
		goto st348
	st449:
		if p++; p == pe {
			goto _test_eof449
		}
	st_case_449:
		switch data[p] {
		case 58:
			goto st375
		case 59:
			goto st14
		case 62:
			goto st349
		case 95:
			goto st347
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st347
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st347
			}
		default:
			goto st347
		}
		goto st348
tr387:
//line xss_170.rl:13

            return true
        
	goto st873
	st873:
		if p++; p == pe {
			goto _test_eof873
		}
	st_case_873:
//line xss_170.go:19519
		switch data[p] {
		case 59:
			goto tr37
		case 62:
			goto tr387
		case 95:
			goto st347
		case 100:
			goto st350
		case 106:
			goto st440
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st347
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st347
			}
		default:
			goto st347
		}
		goto tr386
	st450:
		if p++; p == pe {
			goto _test_eof450
		}
	st_case_450:
		switch data[p] {
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st9
		case 97:
			goto st451
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
	st451:
		if p++; p == pe {
			goto _test_eof451
		}
	st_case_451:
		switch data[p] {
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st9
		case 116:
			goto st452
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
	st452:
		if p++; p == pe {
			goto _test_eof452
		}
	st_case_452:
		switch data[p] {
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st9
		case 97:
			goto st453
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
	st453:
		if p++; p == pe {
			goto _test_eof453
		}
	st_case_453:
		switch data[p] {
		case 58:
			goto st454
		case 59:
			goto st10
		case 60:
			goto st346
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
	st454:
		if p++; p == pe {
			goto _test_eof454
		}
	st_case_454:
		switch data[p] {
		case 44:
			goto st455
		case 59:
			goto st192
		case 60:
			goto st346
		case 95:
			goto st515
		case 100:
			goto st540
		case 106:
			goto st544
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st9
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st515
			}
		default:
			goto st9
		}
		goto st8
	st455:
		if p++; p == pe {
			goto _test_eof455
		}
	st_case_455:
		switch data[p] {
		case 59:
			goto tr210
		case 60:
			goto tr498
		case 95:
			goto st9
		case 100:
			goto st450
		case 106:
			goto st456
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
		goto tr497
tr497:
//line xss_170.rl:13

            return true
        
	goto st874
	st874:
		if p++; p == pe {
			goto _test_eof874
		}
	st_case_874:
//line xss_170.go:19730
		switch data[p] {
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st9
		case 100:
			goto st450
		case 106:
			goto st456
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
	st456:
		if p++; p == pe {
			goto _test_eof456
		}
	st_case_456:
		switch data[p] {
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st9
		case 97:
			goto st457
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
	st457:
		if p++; p == pe {
			goto _test_eof457
		}
	st_case_457:
		switch data[p] {
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st9
		case 118:
			goto st458
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
	st458:
		if p++; p == pe {
			goto _test_eof458
		}
	st_case_458:
		switch data[p] {
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st9
		case 97:
			goto st459
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
	st459:
		if p++; p == pe {
			goto _test_eof459
		}
	st_case_459:
		switch data[p] {
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st9
		case 115:
			goto st460
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
	st460:
		if p++; p == pe {
			goto _test_eof460
		}
	st_case_460:
		switch data[p] {
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st9
		case 99:
			goto st461
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
	st461:
		if p++; p == pe {
			goto _test_eof461
		}
	st_case_461:
		switch data[p] {
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st9
		case 114:
			goto st462
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
	st462:
		if p++; p == pe {
			goto _test_eof462
		}
	st_case_462:
		switch data[p] {
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st9
		case 105:
			goto st463
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
	st463:
		if p++; p == pe {
			goto _test_eof463
		}
	st_case_463:
		switch data[p] {
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st9
		case 112:
			goto st464
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
	st464:
		if p++; p == pe {
			goto _test_eof464
		}
	st_case_464:
		switch data[p] {
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st9
		case 116:
			goto st465
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
	st465:
		if p++; p == pe {
			goto _test_eof465
		}
	st_case_465:
		switch data[p] {
		case 58:
			goto st466
		case 59:
			goto st10
		case 60:
			goto st346
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
	st466:
		if p++; p == pe {
			goto _test_eof466
		}
	st_case_466:
		switch data[p] {
		case 59:
			goto st232
		case 60:
			goto st471
		case 92:
			goto st514
		case 95:
			goto st469
		case 100:
			goto st474
		case 110:
			goto st508
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st469
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st469
			}
		default:
			goto st469
		}
		goto st467
	st467:
		if p++; p == pe {
			goto _test_eof467
		}
	st_case_467:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 46:
			goto st468
		case 59:
			goto st232
		case 60:
			goto st470
		case 61:
			goto st468
		case 91:
			goto st468
		case 92:
			goto st473
		case 95:
			goto st469
		case 100:
			goto st474
		case 110:
			goto st508
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st468
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st469
				}
			case data[p] >= 65:
				goto st469
			}
		default:
			goto st469
		}
		goto st467
	st468:
		if p++; p == pe {
			goto _test_eof468
		}
	st_case_468:
		switch data[p] {
		case 32:
			goto tr519
		case 40:
			goto tr519
		case 46:
			goto tr519
		case 59:
			goto tr278
		case 60:
			goto tr520
		case 61:
			goto tr519
		case 91:
			goto tr519
		case 92:
			goto tr521
		case 95:
			goto st469
		case 100:
			goto st474
		case 110:
			goto st508
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr519
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st469
				}
			case data[p] >= 65:
				goto st469
			}
		default:
			goto st469
		}
		goto tr518
tr518:
//line xss_170.rl:13

            return true
        
	goto st875
	st875:
		if p++; p == pe {
			goto _test_eof875
		}
	st_case_875:
//line xss_170.go:20173
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 46:
			goto st468
		case 59:
			goto st232
		case 60:
			goto st470
		case 61:
			goto st468
		case 91:
			goto st468
		case 92:
			goto st473
		case 95:
			goto st469
		case 100:
			goto st474
		case 110:
			goto st508
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st468
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st469
				}
			case data[p] >= 65:
				goto st469
			}
		default:
			goto st469
		}
		goto st467
	st469:
		if p++; p == pe {
			goto _test_eof469
		}
	st_case_469:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 46:
			goto st468
		case 59:
			goto st232
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st469
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st469
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st469
				}
			case data[p] >= 91:
				goto st468
			}
		default:
			goto st469
		}
		goto st467
	st470:
		if p++; p == pe {
			goto _test_eof470
		}
	st_case_470:
		switch data[p] {
		case 32:
			goto tr520
		case 40:
			goto tr520
		case 46:
			goto tr520
		case 59:
			goto tr282
		case 91:
			goto tr520
		case 92:
			goto tr523
		case 95:
			goto st378
		case 100:
			goto st380
		case 110:
			goto st414
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st378
				}
			case data[p] >= 9:
				goto tr520
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st378
				}
			case data[p] >= 65:
				goto st378
			}
		default:
			goto tr520
		}
		goto tr522
tr522:
//line xss_170.rl:13

            return true
        
	goto st876
	st876:
		if p++; p == pe {
			goto _test_eof876
		}
	st_case_876:
//line xss_170.go:20319
		switch data[p] {
		case 32:
			goto st470
		case 40:
			goto st470
		case 46:
			goto st470
		case 59:
			goto st236
		case 91:
			goto st470
		case 92:
			goto st472
		case 95:
			goto st378
		case 100:
			goto st380
		case 110:
			goto st414
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st378
				}
			case data[p] >= 9:
				goto st470
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st378
				}
			case data[p] >= 65:
				goto st378
			}
		default:
			goto st470
		}
		goto st471
	st471:
		if p++; p == pe {
			goto _test_eof471
		}
	st_case_471:
		switch data[p] {
		case 32:
			goto st470
		case 40:
			goto st470
		case 46:
			goto st470
		case 59:
			goto st236
		case 91:
			goto st470
		case 92:
			goto st472
		case 95:
			goto st378
		case 100:
			goto st380
		case 110:
			goto st414
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st378
				}
			case data[p] >= 9:
				goto st470
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st378
				}
			case data[p] >= 65:
				goto st378
			}
		default:
			goto st470
		}
		goto st471
	st472:
		if p++; p == pe {
			goto _test_eof472
		}
	st_case_472:
		switch data[p] {
		case 32:
			goto tr520
		case 40:
			goto tr520
		case 46:
			goto tr520
		case 59:
			goto tr282
		case 91:
			goto tr520
		case 92:
			goto tr523
		case 95:
			goto st378
		case 100:
			goto st380
		case 110:
			goto st414
		case 117:
			goto st418
		case 120:
			goto st418
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st378
				}
			case data[p] >= 9:
				goto tr520
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st378
				}
			case data[p] >= 65:
				goto st378
			}
		default:
			goto tr520
		}
		goto tr522
tr520:
//line xss_170.rl:13

            return true
        
	goto st877
	st877:
		if p++; p == pe {
			goto _test_eof877
		}
	st_case_877:
//line xss_170.go:20474
		switch data[p] {
		case 32:
			goto tr520
		case 40:
			goto tr520
		case 46:
			goto tr520
		case 59:
			goto tr282
		case 91:
			goto tr520
		case 92:
			goto tr523
		case 95:
			goto st378
		case 100:
			goto st380
		case 110:
			goto st414
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st378
				}
			case data[p] >= 9:
				goto tr520
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st378
				}
			case data[p] >= 65:
				goto st378
			}
		default:
			goto tr520
		}
		goto tr522
tr523:
//line xss_170.rl:13

            return true
        
	goto st878
	st878:
		if p++; p == pe {
			goto _test_eof878
		}
	st_case_878:
//line xss_170.go:20529
		switch data[p] {
		case 32:
			goto tr520
		case 40:
			goto tr520
		case 46:
			goto tr520
		case 59:
			goto tr282
		case 91:
			goto tr520
		case 92:
			goto tr523
		case 95:
			goto st378
		case 100:
			goto st380
		case 110:
			goto st414
		case 117:
			goto st418
		case 120:
			goto st418
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st378
				}
			case data[p] >= 9:
				goto tr520
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st378
				}
			case data[p] >= 65:
				goto st378
			}
		default:
			goto tr520
		}
		goto tr522
	st473:
		if p++; p == pe {
			goto _test_eof473
		}
	st_case_473:
		switch data[p] {
		case 32:
			goto tr519
		case 40:
			goto tr519
		case 46:
			goto tr519
		case 59:
			goto tr278
		case 60:
			goto tr520
		case 61:
			goto tr519
		case 91:
			goto tr519
		case 92:
			goto tr521
		case 95:
			goto st469
		case 100:
			goto st474
		case 110:
			goto st508
		case 117:
			goto st512
		case 120:
			goto st512
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr519
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st469
				}
			case data[p] >= 65:
				goto st469
			}
		default:
			goto st469
		}
		goto tr518
tr519:
//line xss_170.rl:13

            return true
        
	goto st879
	st879:
		if p++; p == pe {
			goto _test_eof879
		}
	st_case_879:
//line xss_170.go:20639
		switch data[p] {
		case 32:
			goto tr519
		case 40:
			goto tr519
		case 46:
			goto tr519
		case 59:
			goto tr278
		case 60:
			goto tr520
		case 61:
			goto tr519
		case 91:
			goto tr519
		case 92:
			goto tr521
		case 95:
			goto st469
		case 100:
			goto st474
		case 110:
			goto st508
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr519
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st469
				}
			case data[p] >= 65:
				goto st469
			}
		default:
			goto st469
		}
		goto tr518
tr521:
//line xss_170.rl:13

            return true
        
	goto st880
	st880:
		if p++; p == pe {
			goto _test_eof880
		}
	st_case_880:
//line xss_170.go:20693
		switch data[p] {
		case 32:
			goto tr519
		case 40:
			goto tr519
		case 46:
			goto tr519
		case 59:
			goto tr278
		case 60:
			goto tr520
		case 61:
			goto tr519
		case 91:
			goto tr519
		case 92:
			goto tr521
		case 95:
			goto st469
		case 100:
			goto st474
		case 110:
			goto st508
		case 117:
			goto st512
		case 120:
			goto st512
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr519
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st469
				}
			case data[p] >= 65:
				goto st469
			}
		default:
			goto st469
		}
		goto tr518
	st474:
		if p++; p == pe {
			goto _test_eof474
		}
	st_case_474:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 46:
			goto st468
		case 59:
			goto st232
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st469
		case 97:
			goto st475
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st469
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st469
				}
			case data[p] >= 91:
				goto st468
			}
		default:
			goto st469
		}
		goto st467
	st475:
		if p++; p == pe {
			goto _test_eof475
		}
	st_case_475:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 46:
			goto st468
		case 59:
			goto st232
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st469
		case 116:
			goto st476
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st469
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st469
				}
			case data[p] >= 91:
				goto st468
			}
		default:
			goto st469
		}
		goto st467
	st476:
		if p++; p == pe {
			goto _test_eof476
		}
	st_case_476:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 46:
			goto st468
		case 59:
			goto st232
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st469
		case 97:
			goto st477
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st469
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st469
				}
			case data[p] >= 91:
				goto st468
			}
		default:
			goto st469
		}
		goto st467
	st477:
		if p++; p == pe {
			goto _test_eof477
		}
	st_case_477:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 46:
			goto st468
		case 58:
			goto st478
		case 59:
			goto st232
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st469
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st469
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st469
				}
			case data[p] >= 91:
				goto st468
			}
		default:
			goto st469
		}
		goto st467
	st478:
		if p++; p == pe {
			goto _test_eof478
		}
	st_case_478:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 44:
			goto st468
		case 46:
			goto st468
		case 59:
			goto st233
		case 60:
			goto st470
		case 61:
			goto st468
		case 91:
			goto st468
		case 92:
			goto st473
		case 95:
			goto st479
		case 100:
			goto st499
		case 110:
			goto st503
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st468
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st479
				}
			case data[p] >= 65:
				goto st469
			}
		default:
			goto st469
		}
		goto st467
	st479:
		if p++; p == pe {
			goto _test_eof479
		}
	st_case_479:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 46:
			goto st468
		case 59:
			goto st232
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st480
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st480
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st480
				}
			case data[p] >= 91:
				goto st468
			}
		default:
			goto st469
		}
		goto st467
	st480:
		if p++; p == pe {
			goto _test_eof480
		}
	st_case_480:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 46:
			goto st468
		case 47:
			goto st481
		case 59:
			goto st232
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st480
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st480
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st480
				}
			case data[p] >= 91:
				goto st468
			}
		default:
			goto st469
		}
		goto st467
	st481:
		if p++; p == pe {
			goto _test_eof481
		}
	st_case_481:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 46:
			goto st468
		case 59:
			goto st232
		case 60:
			goto st470
		case 61:
			goto st468
		case 91:
			goto st468
		case 92:
			goto st473
		case 95:
			goto st482
		case 100:
			goto st495
		case 110:
			goto st497
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st468
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st482
				}
			case data[p] >= 65:
				goto st482
			}
		default:
			goto st482
		}
		goto st467
	st482:
		if p++; p == pe {
			goto _test_eof482
		}
	st_case_482:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 43:
			goto st483
		case 45:
			goto st483
		case 46:
			goto st468
		case 59:
			goto st232
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st494
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st494
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st494
				}
			case data[p] >= 91:
				goto st468
			}
		default:
			goto st494
		}
		goto st467
	st483:
		if p++; p == pe {
			goto _test_eof483
		}
	st_case_483:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 43:
			goto st483
		case 45:
			goto st483
		case 46:
			goto st468
		case 59:
			goto st232
		case 60:
			goto st470
		case 61:
			goto st468
		case 91:
			goto st468
		case 92:
			goto st473
		case 95:
			goto st484
		case 100:
			goto st485
		case 110:
			goto st489
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st468
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st484
				}
			case data[p] >= 65:
				goto st484
			}
		default:
			goto st484
		}
		goto st467
	st484:
		if p++; p == pe {
			goto _test_eof484
		}
	st_case_484:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 44:
			goto st468
		case 46:
			goto st468
		case 59:
			goto st233
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st484
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 13:
				if 43 <= data[p] && data[p] <= 45 {
					goto st483
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st484
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st484
				}
			default:
				goto st468
			}
		default:
			goto st484
		}
		goto st467
	st485:
		if p++; p == pe {
			goto _test_eof485
		}
	st_case_485:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 44:
			goto st468
		case 46:
			goto st468
		case 59:
			goto st233
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st484
		case 97:
			goto st486
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 13:
				if 43 <= data[p] && data[p] <= 45 {
					goto st483
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st484
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st484
				}
			default:
				goto st468
			}
		default:
			goto st484
		}
		goto st467
	st486:
		if p++; p == pe {
			goto _test_eof486
		}
	st_case_486:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 44:
			goto st468
		case 46:
			goto st468
		case 59:
			goto st233
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st484
		case 116:
			goto st487
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 13:
				if 43 <= data[p] && data[p] <= 45 {
					goto st483
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st484
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st484
				}
			default:
				goto st468
			}
		default:
			goto st484
		}
		goto st467
	st487:
		if p++; p == pe {
			goto _test_eof487
		}
	st_case_487:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 44:
			goto st468
		case 46:
			goto st468
		case 59:
			goto st233
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st484
		case 97:
			goto st488
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 13:
				if 43 <= data[p] && data[p] <= 45 {
					goto st483
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st484
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st484
				}
			default:
				goto st468
			}
		default:
			goto st484
		}
		goto st467
	st488:
		if p++; p == pe {
			goto _test_eof488
		}
	st_case_488:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 44:
			goto st468
		case 46:
			goto st468
		case 58:
			goto st478
		case 59:
			goto st233
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st484
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 13:
				if 43 <= data[p] && data[p] <= 45 {
					goto st483
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st484
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st484
				}
			default:
				goto st468
			}
		default:
			goto st484
		}
		goto st467
	st489:
		if p++; p == pe {
			goto _test_eof489
		}
	st_case_489:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 44:
			goto st468
		case 46:
			goto st468
		case 59:
			goto st233
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st484
		case 97:
			goto st490
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 13:
				if 43 <= data[p] && data[p] <= 45 {
					goto st483
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st484
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st484
				}
			default:
				goto st468
			}
		default:
			goto st484
		}
		goto st467
	st490:
		if p++; p == pe {
			goto _test_eof490
		}
	st_case_490:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 44:
			goto st468
		case 46:
			goto st468
		case 59:
			goto st233
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st484
		case 109:
			goto st491
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 13:
				if 43 <= data[p] && data[p] <= 45 {
					goto st483
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st484
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st484
				}
			default:
				goto st468
			}
		default:
			goto st484
		}
		goto st467
	st491:
		if p++; p == pe {
			goto _test_eof491
		}
	st_case_491:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 44:
			goto st468
		case 46:
			goto st468
		case 59:
			goto st233
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st484
		case 101:
			goto st492
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 13:
				if 43 <= data[p] && data[p] <= 45 {
					goto st483
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st484
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st484
				}
			default:
				goto st468
			}
		default:
			goto st484
		}
		goto st467
	st492:
		if p++; p == pe {
			goto _test_eof492
		}
	st_case_492:
		switch data[p] {
		case 43:
			goto st493
		case 45:
			goto st493
		case 59:
			goto st233
		case 60:
			goto st470
		case 95:
			goto st484
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st484
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st484
			}
		default:
			goto st484
		}
		goto st468
	st493:
		if p++; p == pe {
			goto _test_eof493
		}
	st_case_493:
		switch data[p] {
		case 32:
			goto tr519
		case 40:
			goto tr519
		case 43:
			goto tr550
		case 45:
			goto tr550
		case 46:
			goto tr519
		case 59:
			goto tr278
		case 60:
			goto tr520
		case 61:
			goto tr519
		case 91:
			goto tr519
		case 92:
			goto tr521
		case 95:
			goto st484
		case 100:
			goto st485
		case 110:
			goto st489
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr519
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st484
				}
			case data[p] >= 65:
				goto st484
			}
		default:
			goto st484
		}
		goto tr518
tr550:
//line xss_170.rl:13

            return true
        
	goto st881
	st881:
		if p++; p == pe {
			goto _test_eof881
		}
	st_case_881:
//line xss_170.go:21715
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 43:
			goto st483
		case 45:
			goto st483
		case 46:
			goto st468
		case 59:
			goto st232
		case 60:
			goto st470
		case 61:
			goto st468
		case 91:
			goto st468
		case 92:
			goto st473
		case 95:
			goto st484
		case 100:
			goto st485
		case 110:
			goto st489
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st468
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st484
				}
			case data[p] >= 65:
				goto st484
			}
		default:
			goto st484
		}
		goto st467
	st494:
		if p++; p == pe {
			goto _test_eof494
		}
	st_case_494:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 43:
			goto st483
		case 45:
			goto st483
		case 46:
			goto st468
		case 59:
			goto st232
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st484
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st484
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st484
				}
			case data[p] >= 91:
				goto st468
			}
		default:
			goto st484
		}
		goto st467
	st495:
		if p++; p == pe {
			goto _test_eof495
		}
	st_case_495:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 43:
			goto st483
		case 45:
			goto st483
		case 46:
			goto st468
		case 59:
			goto st232
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st494
		case 97:
			goto st496
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st494
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st494
				}
			case data[p] >= 91:
				goto st468
			}
		default:
			goto st494
		}
		goto st467
	st496:
		if p++; p == pe {
			goto _test_eof496
		}
	st_case_496:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 43:
			goto st483
		case 45:
			goto st483
		case 46:
			goto st468
		case 59:
			goto st232
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st484
		case 116:
			goto st487
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st484
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st484
				}
			case data[p] >= 91:
				goto st468
			}
		default:
			goto st484
		}
		goto st467
	st497:
		if p++; p == pe {
			goto _test_eof497
		}
	st_case_497:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 43:
			goto st483
		case 45:
			goto st483
		case 46:
			goto st468
		case 59:
			goto st232
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st494
		case 97:
			goto st498
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st494
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st494
				}
			case data[p] >= 91:
				goto st468
			}
		default:
			goto st494
		}
		goto st467
	st498:
		if p++; p == pe {
			goto _test_eof498
		}
	st_case_498:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 43:
			goto st483
		case 45:
			goto st483
		case 46:
			goto st468
		case 59:
			goto st232
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st484
		case 109:
			goto st491
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st484
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st484
				}
			case data[p] >= 91:
				goto st468
			}
		default:
			goto st484
		}
		goto st467
	st499:
		if p++; p == pe {
			goto _test_eof499
		}
	st_case_499:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 46:
			goto st468
		case 59:
			goto st232
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st480
		case 97:
			goto st500
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st480
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st480
				}
			case data[p] >= 91:
				goto st468
			}
		default:
			goto st469
		}
		goto st467
	st500:
		if p++; p == pe {
			goto _test_eof500
		}
	st_case_500:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 46:
			goto st468
		case 47:
			goto st481
		case 59:
			goto st232
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st480
		case 116:
			goto st501
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st480
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st480
				}
			case data[p] >= 91:
				goto st468
			}
		default:
			goto st469
		}
		goto st467
	st501:
		if p++; p == pe {
			goto _test_eof501
		}
	st_case_501:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 46:
			goto st468
		case 47:
			goto st481
		case 59:
			goto st232
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st480
		case 97:
			goto st502
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st480
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st480
				}
			case data[p] >= 91:
				goto st468
			}
		default:
			goto st469
		}
		goto st467
	st502:
		if p++; p == pe {
			goto _test_eof502
		}
	st_case_502:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 46:
			goto st468
		case 47:
			goto st481
		case 58:
			goto st478
		case 59:
			goto st232
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st480
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st480
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st480
				}
			case data[p] >= 91:
				goto st468
			}
		default:
			goto st469
		}
		goto st467
	st503:
		if p++; p == pe {
			goto _test_eof503
		}
	st_case_503:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 46:
			goto st468
		case 59:
			goto st232
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st480
		case 97:
			goto st504
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st480
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st480
				}
			case data[p] >= 91:
				goto st468
			}
		default:
			goto st469
		}
		goto st467
	st504:
		if p++; p == pe {
			goto _test_eof504
		}
	st_case_504:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 46:
			goto st468
		case 47:
			goto st481
		case 59:
			goto st232
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st480
		case 109:
			goto st505
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st480
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st480
				}
			case data[p] >= 91:
				goto st468
			}
		default:
			goto st469
		}
		goto st467
	st505:
		if p++; p == pe {
			goto _test_eof505
		}
	st_case_505:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 46:
			goto st468
		case 47:
			goto st481
		case 59:
			goto st232
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st480
		case 101:
			goto st506
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st480
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st480
				}
			case data[p] >= 91:
				goto st468
			}
		default:
			goto st469
		}
		goto st467
	st506:
		if p++; p == pe {
			goto _test_eof506
		}
	st_case_506:
		switch data[p] {
		case 47:
			goto st507
		case 59:
			goto st233
		case 60:
			goto st470
		case 95:
			goto st480
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st480
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st480
			}
		default:
			goto st469
		}
		goto st468
	st507:
		if p++; p == pe {
			goto _test_eof507
		}
	st_case_507:
		switch data[p] {
		case 32:
			goto tr519
		case 40:
			goto tr519
		case 46:
			goto tr519
		case 59:
			goto tr278
		case 60:
			goto tr520
		case 61:
			goto tr519
		case 91:
			goto tr519
		case 92:
			goto tr521
		case 95:
			goto st482
		case 100:
			goto st495
		case 110:
			goto st497
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr519
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st482
				}
			case data[p] >= 65:
				goto st482
			}
		default:
			goto st482
		}
		goto tr518
	st508:
		if p++; p == pe {
			goto _test_eof508
		}
	st_case_508:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 46:
			goto st468
		case 59:
			goto st232
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st469
		case 97:
			goto st509
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st469
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st469
				}
			case data[p] >= 91:
				goto st468
			}
		default:
			goto st469
		}
		goto st467
	st509:
		if p++; p == pe {
			goto _test_eof509
		}
	st_case_509:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 46:
			goto st468
		case 59:
			goto st232
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st469
		case 109:
			goto st510
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st469
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st469
				}
			case data[p] >= 91:
				goto st468
			}
		default:
			goto st469
		}
		goto st467
	st510:
		if p++; p == pe {
			goto _test_eof510
		}
	st_case_510:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 46:
			goto st468
		case 59:
			goto st232
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st469
		case 101:
			goto st511
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st469
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st469
				}
			case data[p] >= 91:
				goto st468
			}
		default:
			goto st469
		}
		goto st467
	st511:
		if p++; p == pe {
			goto _test_eof511
		}
	st_case_511:
		switch data[p] {
		case 59:
			goto st233
		case 60:
			goto st470
		case 95:
			goto st469
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st469
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st469
			}
		default:
			goto st469
		}
		goto st468
	st512:
		if p++; p == pe {
			goto _test_eof512
		}
	st_case_512:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 46:
			goto st468
		case 59:
			goto st232
		case 60:
			goto st470
		case 61:
			goto st468
		case 95:
			goto st469
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st513
				}
			case data[p] >= 9:
				goto st468
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st469
				}
			case data[p] >= 91:
				goto st468
			}
		default:
			goto st469
		}
		goto st467
	st513:
		if p++; p == pe {
			goto _test_eof513
		}
	st_case_513:
		switch data[p] {
		case 32:
			goto tr519
		case 40:
			goto tr519
		case 46:
			goto tr519
		case 59:
			goto tr278
		case 60:
			goto tr520
		case 61:
			goto tr519
		case 95:
			goto st469
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st513
				}
			case data[p] >= 9:
				goto tr519
			}
		case data[p] > 90:
			switch {
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st469
				}
			case data[p] >= 91:
				goto tr519
			}
		default:
			goto st469
		}
		goto tr518
	st514:
		if p++; p == pe {
			goto _test_eof514
		}
	st_case_514:
		switch data[p] {
		case 32:
			goto st468
		case 40:
			goto st468
		case 46:
			goto st468
		case 59:
			goto st232
		case 60:
			goto st470
		case 61:
			goto st468
		case 91:
			goto st468
		case 92:
			goto st473
		case 95:
			goto st469
		case 100:
			goto st474
		case 110:
			goto st508
		case 117:
			goto st512
		case 120:
			goto st512
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st468
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st469
				}
			case data[p] >= 65:
				goto st469
			}
		default:
			goto st469
		}
		goto st467
tr498:
//line xss_170.rl:13

            return true
        
	goto st882
	st882:
		if p++; p == pe {
			goto _test_eof882
		}
	st_case_882:
//line xss_170.go:22731
		switch data[p] {
		case 59:
			goto st12
		case 95:
			goto st347
		case 100:
			goto st350
		case 106:
			goto st440
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st347
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st347
			}
		default:
			goto st347
		}
		goto st346
	st515:
		if p++; p == pe {
			goto _test_eof515
		}
	st_case_515:
		switch data[p] {
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st516
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st516
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st516
			}
		default:
			goto st9
		}
		goto st8
	st516:
		if p++; p == pe {
			goto _test_eof516
		}
	st_case_516:
		switch data[p] {
		case 47:
			goto st517
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st516
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st516
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st516
			}
		default:
			goto st9
		}
		goto st8
	st517:
		if p++; p == pe {
			goto _test_eof517
		}
	st_case_517:
		switch data[p] {
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st518
		case 100:
			goto st536
		case 106:
			goto st538
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st518
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st518
			}
		default:
			goto st518
		}
		goto st8
	st518:
		if p++; p == pe {
			goto _test_eof518
		}
	st_case_518:
		switch data[p] {
		case 43:
			goto st519
		case 45:
			goto st519
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st535
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st535
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st535
			}
		default:
			goto st535
		}
		goto st8
	st519:
		if p++; p == pe {
			goto _test_eof519
		}
	st_case_519:
		switch data[p] {
		case 43:
			goto st519
		case 45:
			goto st519
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st520
		case 100:
			goto st521
		case 106:
			goto st525
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st520
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st520
			}
		default:
			goto st520
		}
		goto st8
	st520:
		if p++; p == pe {
			goto _test_eof520
		}
	st_case_520:
		switch data[p] {
		case 44:
			goto st455
		case 59:
			goto st192
		case 60:
			goto st346
		case 95:
			goto st520
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st519
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st520
				}
			case data[p] >= 65:
				goto st520
			}
		default:
			goto st520
		}
		goto st8
	st521:
		if p++; p == pe {
			goto _test_eof521
		}
	st_case_521:
		switch data[p] {
		case 44:
			goto st455
		case 59:
			goto st192
		case 60:
			goto st346
		case 95:
			goto st520
		case 97:
			goto st522
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st519
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st520
				}
			case data[p] >= 65:
				goto st520
			}
		default:
			goto st520
		}
		goto st8
	st522:
		if p++; p == pe {
			goto _test_eof522
		}
	st_case_522:
		switch data[p] {
		case 44:
			goto st455
		case 59:
			goto st192
		case 60:
			goto st346
		case 95:
			goto st520
		case 116:
			goto st523
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st519
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st520
				}
			case data[p] >= 65:
				goto st520
			}
		default:
			goto st520
		}
		goto st8
	st523:
		if p++; p == pe {
			goto _test_eof523
		}
	st_case_523:
		switch data[p] {
		case 44:
			goto st455
		case 59:
			goto st192
		case 60:
			goto st346
		case 95:
			goto st520
		case 97:
			goto st524
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st519
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st520
				}
			case data[p] >= 65:
				goto st520
			}
		default:
			goto st520
		}
		goto st8
	st524:
		if p++; p == pe {
			goto _test_eof524
		}
	st_case_524:
		switch data[p] {
		case 44:
			goto st455
		case 58:
			goto st454
		case 59:
			goto st192
		case 60:
			goto st346
		case 95:
			goto st520
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st519
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st520
				}
			case data[p] >= 65:
				goto st520
			}
		default:
			goto st520
		}
		goto st8
	st525:
		if p++; p == pe {
			goto _test_eof525
		}
	st_case_525:
		switch data[p] {
		case 44:
			goto st455
		case 59:
			goto st192
		case 60:
			goto st346
		case 95:
			goto st520
		case 97:
			goto st526
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st519
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st520
				}
			case data[p] >= 65:
				goto st520
			}
		default:
			goto st520
		}
		goto st8
	st526:
		if p++; p == pe {
			goto _test_eof526
		}
	st_case_526:
		switch data[p] {
		case 44:
			goto st455
		case 59:
			goto st192
		case 60:
			goto st346
		case 95:
			goto st520
		case 118:
			goto st527
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st519
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st520
				}
			case data[p] >= 65:
				goto st520
			}
		default:
			goto st520
		}
		goto st8
	st527:
		if p++; p == pe {
			goto _test_eof527
		}
	st_case_527:
		switch data[p] {
		case 44:
			goto st455
		case 59:
			goto st192
		case 60:
			goto st346
		case 95:
			goto st520
		case 97:
			goto st528
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st519
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st520
				}
			case data[p] >= 65:
				goto st520
			}
		default:
			goto st520
		}
		goto st8
	st528:
		if p++; p == pe {
			goto _test_eof528
		}
	st_case_528:
		switch data[p] {
		case 44:
			goto st455
		case 59:
			goto st192
		case 60:
			goto st346
		case 95:
			goto st520
		case 115:
			goto st529
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st519
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st520
				}
			case data[p] >= 65:
				goto st520
			}
		default:
			goto st520
		}
		goto st8
	st529:
		if p++; p == pe {
			goto _test_eof529
		}
	st_case_529:
		switch data[p] {
		case 44:
			goto st455
		case 59:
			goto st192
		case 60:
			goto st346
		case 95:
			goto st520
		case 99:
			goto st530
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st519
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st520
				}
			case data[p] >= 65:
				goto st520
			}
		default:
			goto st520
		}
		goto st8
	st530:
		if p++; p == pe {
			goto _test_eof530
		}
	st_case_530:
		switch data[p] {
		case 44:
			goto st455
		case 59:
			goto st192
		case 60:
			goto st346
		case 95:
			goto st520
		case 114:
			goto st531
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st519
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st520
				}
			case data[p] >= 65:
				goto st520
			}
		default:
			goto st520
		}
		goto st8
	st531:
		if p++; p == pe {
			goto _test_eof531
		}
	st_case_531:
		switch data[p] {
		case 44:
			goto st455
		case 59:
			goto st192
		case 60:
			goto st346
		case 95:
			goto st520
		case 105:
			goto st532
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st519
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st520
				}
			case data[p] >= 65:
				goto st520
			}
		default:
			goto st520
		}
		goto st8
	st532:
		if p++; p == pe {
			goto _test_eof532
		}
	st_case_532:
		switch data[p] {
		case 44:
			goto st455
		case 59:
			goto st192
		case 60:
			goto st346
		case 95:
			goto st520
		case 112:
			goto st533
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st519
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st520
				}
			case data[p] >= 65:
				goto st520
			}
		default:
			goto st520
		}
		goto st8
	st533:
		if p++; p == pe {
			goto _test_eof533
		}
	st_case_533:
		switch data[p] {
		case 44:
			goto st455
		case 59:
			goto st192
		case 60:
			goto st346
		case 95:
			goto st520
		case 116:
			goto st534
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st519
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st520
				}
			case data[p] >= 65:
				goto st520
			}
		default:
			goto st520
		}
		goto st8
	st534:
		if p++; p == pe {
			goto _test_eof534
		}
	st_case_534:
		switch data[p] {
		case 44:
			goto st455
		case 58:
			goto st466
		case 59:
			goto st192
		case 60:
			goto st346
		case 95:
			goto st520
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st519
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st520
				}
			case data[p] >= 65:
				goto st520
			}
		default:
			goto st520
		}
		goto st8
	st535:
		if p++; p == pe {
			goto _test_eof535
		}
	st_case_535:
		switch data[p] {
		case 43:
			goto st519
		case 45:
			goto st519
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st520
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st520
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st520
			}
		default:
			goto st520
		}
		goto st8
	st536:
		if p++; p == pe {
			goto _test_eof536
		}
	st_case_536:
		switch data[p] {
		case 43:
			goto st519
		case 45:
			goto st519
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st535
		case 97:
			goto st537
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st535
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st535
			}
		default:
			goto st535
		}
		goto st8
	st537:
		if p++; p == pe {
			goto _test_eof537
		}
	st_case_537:
		switch data[p] {
		case 43:
			goto st519
		case 45:
			goto st519
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st520
		case 116:
			goto st523
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st520
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st520
			}
		default:
			goto st520
		}
		goto st8
	st538:
		if p++; p == pe {
			goto _test_eof538
		}
	st_case_538:
		switch data[p] {
		case 43:
			goto st519
		case 45:
			goto st519
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st535
		case 97:
			goto st539
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st535
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st535
			}
		default:
			goto st535
		}
		goto st8
	st539:
		if p++; p == pe {
			goto _test_eof539
		}
	st_case_539:
		switch data[p] {
		case 43:
			goto st519
		case 45:
			goto st519
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st520
		case 118:
			goto st527
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st520
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st520
			}
		default:
			goto st520
		}
		goto st8
	st540:
		if p++; p == pe {
			goto _test_eof540
		}
	st_case_540:
		switch data[p] {
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st516
		case 97:
			goto st541
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st516
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st516
			}
		default:
			goto st9
		}
		goto st8
	st541:
		if p++; p == pe {
			goto _test_eof541
		}
	st_case_541:
		switch data[p] {
		case 47:
			goto st517
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st516
		case 116:
			goto st542
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st516
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st516
			}
		default:
			goto st9
		}
		goto st8
	st542:
		if p++; p == pe {
			goto _test_eof542
		}
	st_case_542:
		switch data[p] {
		case 47:
			goto st517
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st516
		case 97:
			goto st543
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st516
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st516
			}
		default:
			goto st9
		}
		goto st8
	st543:
		if p++; p == pe {
			goto _test_eof543
		}
	st_case_543:
		switch data[p] {
		case 47:
			goto st517
		case 58:
			goto st454
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st516
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st516
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st516
			}
		default:
			goto st9
		}
		goto st8
	st544:
		if p++; p == pe {
			goto _test_eof544
		}
	st_case_544:
		switch data[p] {
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st516
		case 97:
			goto st545
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st516
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st516
			}
		default:
			goto st9
		}
		goto st8
	st545:
		if p++; p == pe {
			goto _test_eof545
		}
	st_case_545:
		switch data[p] {
		case 47:
			goto st517
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st516
		case 118:
			goto st546
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st516
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st516
			}
		default:
			goto st9
		}
		goto st8
	st546:
		if p++; p == pe {
			goto _test_eof546
		}
	st_case_546:
		switch data[p] {
		case 47:
			goto st517
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st516
		case 97:
			goto st547
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st516
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st516
			}
		default:
			goto st9
		}
		goto st8
	st547:
		if p++; p == pe {
			goto _test_eof547
		}
	st_case_547:
		switch data[p] {
		case 47:
			goto st517
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st516
		case 115:
			goto st548
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st516
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st516
			}
		default:
			goto st9
		}
		goto st8
	st548:
		if p++; p == pe {
			goto _test_eof548
		}
	st_case_548:
		switch data[p] {
		case 47:
			goto st517
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st516
		case 99:
			goto st549
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st516
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st516
			}
		default:
			goto st9
		}
		goto st8
	st549:
		if p++; p == pe {
			goto _test_eof549
		}
	st_case_549:
		switch data[p] {
		case 47:
			goto st517
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st516
		case 114:
			goto st550
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st516
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st516
			}
		default:
			goto st9
		}
		goto st8
	st550:
		if p++; p == pe {
			goto _test_eof550
		}
	st_case_550:
		switch data[p] {
		case 47:
			goto st517
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st516
		case 105:
			goto st551
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st516
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st516
			}
		default:
			goto st9
		}
		goto st8
	st551:
		if p++; p == pe {
			goto _test_eof551
		}
	st_case_551:
		switch data[p] {
		case 47:
			goto st517
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st516
		case 112:
			goto st552
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st516
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st516
			}
		default:
			goto st9
		}
		goto st8
	st552:
		if p++; p == pe {
			goto _test_eof552
		}
	st_case_552:
		switch data[p] {
		case 47:
			goto st517
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st516
		case 116:
			goto st553
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st516
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st516
			}
		default:
			goto st9
		}
		goto st8
	st553:
		if p++; p == pe {
			goto _test_eof553
		}
	st_case_553:
		switch data[p] {
		case 47:
			goto st517
		case 58:
			goto st466
		case 59:
			goto st10
		case 60:
			goto st346
		case 95:
			goto st516
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st516
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st516
			}
		default:
			goto st9
		}
		goto st8
	st554:
		if p++; p == pe {
			goto _test_eof554
		}
	st_case_554:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st555
		case 95:
			goto st554
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st554
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st554
			}
		default:
			goto st554
		}
		goto st7
	st555:
		if p++; p == pe {
			goto _test_eof555
		}
	st_case_555:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st556
		case 98:
			goto st557
		case 99:
			goto st563
		case 100:
			goto st571
		case 106:
			goto st675
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto st555
	st556:
		if p++; p == pe {
			goto _test_eof556
		}
	st_case_556:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st556
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto st555
	st557:
		if p++; p == pe {
			goto _test_eof557
		}
	st_case_557:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st556
		case 97:
			goto st558
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto st555
	st558:
		if p++; p == pe {
			goto _test_eof558
		}
	st_case_558:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st556
		case 115:
			goto st559
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto st555
	st559:
		if p++; p == pe {
			goto _test_eof559
		}
	st_case_559:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st556
		case 101:
			goto st560
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto st555
	st560:
		if p++; p == pe {
			goto _test_eof560
		}
	st_case_560:
		switch data[p] {
		case 44:
			goto st10
		case 54:
			goto st561
		case 95:
			goto st556
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto st555
	st561:
		if p++; p == pe {
			goto _test_eof561
		}
	st_case_561:
		switch data[p] {
		case 44:
			goto st10
		case 52:
			goto st562
		case 95:
			goto st556
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto st555
	st562:
		if p++; p == pe {
			goto _test_eof562
		}
	st_case_562:
		switch data[p] {
		case 44:
			goto tr210
		case 95:
			goto st556
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto tr610
tr610:
//line xss_170.rl:13

            return true
        
	goto st883
	st883:
		if p++; p == pe {
			goto _test_eof883
		}
	st_case_883:
//line xss_170.go:24247
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st556
		case 98:
			goto st557
		case 99:
			goto st563
		case 100:
			goto st571
		case 106:
			goto st675
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto st555
	st563:
		if p++; p == pe {
			goto _test_eof563
		}
	st_case_563:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st556
		case 104:
			goto st564
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto st555
	st564:
		if p++; p == pe {
			goto _test_eof564
		}
	st_case_564:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st556
		case 97:
			goto st565
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto st555
	st565:
		if p++; p == pe {
			goto _test_eof565
		}
	st_case_565:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st556
		case 114:
			goto st566
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto st555
	st566:
		if p++; p == pe {
			goto _test_eof566
		}
	st_case_566:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st556
		case 115:
			goto st567
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto st555
	st567:
		if p++; p == pe {
			goto _test_eof567
		}
	st_case_567:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st556
		case 101:
			goto st568
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto st555
	st568:
		if p++; p == pe {
			goto _test_eof568
		}
	st_case_568:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st556
		case 116:
			goto st569
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto st555
	st569:
		if p++; p == pe {
			goto _test_eof569
		}
	st_case_569:
		switch data[p] {
		case 44:
			goto st10
		case 61:
			goto st570
		case 95:
			goto st556
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto st555
	st570:
		if p++; p == pe {
			goto _test_eof570
		}
	st_case_570:
		switch data[p] {
		case 44:
			goto tr210
		case 95:
			goto st556
		case 98:
			goto st557
		case 99:
			goto st563
		case 100:
			goto st571
		case 106:
			goto st675
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto tr610
	st571:
		if p++; p == pe {
			goto _test_eof571
		}
	st_case_571:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st556
		case 97:
			goto st572
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto st555
	st572:
		if p++; p == pe {
			goto _test_eof572
		}
	st_case_572:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st556
		case 116:
			goto st573
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto st555
	st573:
		if p++; p == pe {
			goto _test_eof573
		}
	st_case_573:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st556
		case 97:
			goto st574
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto st555
	st574:
		if p++; p == pe {
			goto _test_eof574
		}
	st_case_574:
		switch data[p] {
		case 44:
			goto st10
		case 58:
			goto st575
		case 95:
			goto st556
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto st555
	st575:
		if p++; p == pe {
			goto _test_eof575
		}
	st_case_575:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st570
		case 95:
			goto st576
		case 98:
			goto st694
		case 99:
			goto st700
		case 100:
			goto st707
		case 106:
			goto st711
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st576
			}
		default:
			goto st556
		}
		goto st555
	st576:
		if p++; p == pe {
			goto _test_eof576
		}
	st_case_576:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st577
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st577
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st577
			}
		default:
			goto st556
		}
		goto st555
	st577:
		if p++; p == pe {
			goto _test_eof577
		}
	st_case_577:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st578
		case 95:
			goto st577
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st577
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st577
			}
		default:
			goto st556
		}
		goto st555
	st578:
		if p++; p == pe {
			goto _test_eof578
		}
	st_case_578:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st579
		case 98:
			goto st686
		case 99:
			goto st688
		case 100:
			goto st690
		case 106:
			goto st692
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st579
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st579
			}
		default:
			goto st579
		}
		goto st555
	st579:
		if p++; p == pe {
			goto _test_eof579
		}
	st_case_579:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st685
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st685
				}
			case data[p] >= 65:
				goto st685
			}
		default:
			goto st685
		}
		goto st555
	st580:
		if p++; p == pe {
			goto _test_eof580
		}
	st_case_580:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st581
		case 98:
			goto st582
		case 99:
			goto st588
		case 100:
			goto st595
		case 106:
			goto st599
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st581:
		if p++; p == pe {
			goto _test_eof581
		}
	st_case_581:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st570
		case 95:
			goto st581
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st582:
		if p++; p == pe {
			goto _test_eof582
		}
	st_case_582:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st570
		case 95:
			goto st581
		case 97:
			goto st583
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st583:
		if p++; p == pe {
			goto _test_eof583
		}
	st_case_583:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st570
		case 95:
			goto st581
		case 115:
			goto st584
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st584:
		if p++; p == pe {
			goto _test_eof584
		}
	st_case_584:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st570
		case 95:
			goto st581
		case 101:
			goto st585
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st585:
		if p++; p == pe {
			goto _test_eof585
		}
	st_case_585:
		switch data[p] {
		case 44:
			goto st192
		case 54:
			goto st586
		case 59:
			goto st570
		case 95:
			goto st581
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st586:
		if p++; p == pe {
			goto _test_eof586
		}
	st_case_586:
		switch data[p] {
		case 44:
			goto st192
		case 52:
			goto st587
		case 59:
			goto st570
		case 95:
			goto st581
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st587:
		if p++; p == pe {
			goto _test_eof587
		}
	st_case_587:
		switch data[p] {
		case 44:
			goto tr248
		case 59:
			goto tr646
		case 95:
			goto st581
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto tr645
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto tr610
tr645:
//line xss_170.rl:13

            return true
        
	goto st884
	st884:
		if p++; p == pe {
			goto _test_eof884
		}
	st_case_884:
//line xss_170.go:25013
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st581
		case 98:
			goto st582
		case 99:
			goto st588
		case 100:
			goto st595
		case 106:
			goto st599
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st588:
		if p++; p == pe {
			goto _test_eof588
		}
	st_case_588:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st570
		case 95:
			goto st581
		case 104:
			goto st589
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st589:
		if p++; p == pe {
			goto _test_eof589
		}
	st_case_589:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st570
		case 95:
			goto st581
		case 97:
			goto st590
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st590:
		if p++; p == pe {
			goto _test_eof590
		}
	st_case_590:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st570
		case 95:
			goto st581
		case 114:
			goto st591
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st591:
		if p++; p == pe {
			goto _test_eof591
		}
	st_case_591:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st570
		case 95:
			goto st581
		case 115:
			goto st592
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st592:
		if p++; p == pe {
			goto _test_eof592
		}
	st_case_592:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st570
		case 95:
			goto st581
		case 101:
			goto st593
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st593:
		if p++; p == pe {
			goto _test_eof593
		}
	st_case_593:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st570
		case 95:
			goto st581
		case 116:
			goto st594
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st594:
		if p++; p == pe {
			goto _test_eof594
		}
	st_case_594:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st570
		case 61:
			goto st570
		case 95:
			goto st581
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st595:
		if p++; p == pe {
			goto _test_eof595
		}
	st_case_595:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st570
		case 95:
			goto st581
		case 97:
			goto st596
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st596:
		if p++; p == pe {
			goto _test_eof596
		}
	st_case_596:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st570
		case 95:
			goto st581
		case 116:
			goto st597
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st597:
		if p++; p == pe {
			goto _test_eof597
		}
	st_case_597:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st570
		case 95:
			goto st581
		case 97:
			goto st598
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st598:
		if p++; p == pe {
			goto _test_eof598
		}
	st_case_598:
		switch data[p] {
		case 44:
			goto st192
		case 58:
			goto st575
		case 59:
			goto st570
		case 95:
			goto st581
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st599:
		if p++; p == pe {
			goto _test_eof599
		}
	st_case_599:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st570
		case 95:
			goto st581
		case 97:
			goto st600
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st600:
		if p++; p == pe {
			goto _test_eof600
		}
	st_case_600:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st570
		case 95:
			goto st581
		case 118:
			goto st601
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st601:
		if p++; p == pe {
			goto _test_eof601
		}
	st_case_601:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st570
		case 95:
			goto st581
		case 97:
			goto st602
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st602:
		if p++; p == pe {
			goto _test_eof602
		}
	st_case_602:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st570
		case 95:
			goto st581
		case 115:
			goto st603
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st603:
		if p++; p == pe {
			goto _test_eof603
		}
	st_case_603:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st570
		case 95:
			goto st581
		case 99:
			goto st604
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st604:
		if p++; p == pe {
			goto _test_eof604
		}
	st_case_604:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st570
		case 95:
			goto st581
		case 114:
			goto st605
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st605:
		if p++; p == pe {
			goto _test_eof605
		}
	st_case_605:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st570
		case 95:
			goto st581
		case 105:
			goto st606
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st606:
		if p++; p == pe {
			goto _test_eof606
		}
	st_case_606:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st570
		case 95:
			goto st581
		case 112:
			goto st607
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st607:
		if p++; p == pe {
			goto _test_eof607
		}
	st_case_607:
		switch data[p] {
		case 44:
			goto st192
		case 59:
			goto st570
		case 95:
			goto st581
		case 116:
			goto st608
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st608:
		if p++; p == pe {
			goto _test_eof608
		}
	st_case_608:
		switch data[p] {
		case 44:
			goto st192
		case 58:
			goto st609
		case 59:
			goto st570
		case 95:
			goto st581
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st609:
		if p++; p == pe {
			goto _test_eof609
		}
	st_case_609:
		switch data[p] {
		case 44:
			goto st232
		case 92:
			goto st674
		case 95:
			goto st612
		case 98:
			goto st614
		case 100:
			goto st620
		case 110:
			goto st668
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st612
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st612
			}
		default:
			goto st612
		}
		goto st610
	st610:
		if p++; p == pe {
			goto _test_eof610
		}
	st_case_610:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 91:
			goto st611
		case 92:
			goto st613
		case 95:
			goto st612
		case 98:
			goto st614
		case 100:
			goto st620
		case 110:
			goto st668
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st612
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st612
				}
			case data[p] >= 65:
				goto st612
			}
		default:
			goto st611
		}
		goto st610
	st611:
		if p++; p == pe {
			goto _test_eof611
		}
	st_case_611:
		switch data[p] {
		case 32:
			goto tr675
		case 40:
			goto tr675
		case 44:
			goto tr278
		case 46:
			goto tr675
		case 91:
			goto tr675
		case 92:
			goto tr676
		case 95:
			goto st612
		case 98:
			goto st614
		case 100:
			goto st620
		case 110:
			goto st668
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st612
				}
			case data[p] >= 9:
				goto tr675
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st612
				}
			case data[p] >= 65:
				goto st612
			}
		default:
			goto tr675
		}
		goto tr674
tr674:
//line xss_170.rl:13

            return true
        
	goto st885
	st885:
		if p++; p == pe {
			goto _test_eof885
		}
	st_case_885:
//line xss_170.go:25882
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 91:
			goto st611
		case 92:
			goto st613
		case 95:
			goto st612
		case 98:
			goto st614
		case 100:
			goto st620
		case 110:
			goto st668
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st612
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st612
				}
			case data[p] >= 65:
				goto st612
			}
		default:
			goto st611
		}
		goto st610
	st612:
		if p++; p == pe {
			goto _test_eof612
		}
	st_case_612:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 95:
			goto st612
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st612
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st612
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st613:
		if p++; p == pe {
			goto _test_eof613
		}
	st_case_613:
		switch data[p] {
		case 32:
			goto tr675
		case 40:
			goto tr675
		case 44:
			goto tr278
		case 46:
			goto tr675
		case 91:
			goto tr675
		case 92:
			goto tr676
		case 95:
			goto st612
		case 98:
			goto st614
		case 100:
			goto st620
		case 110:
			goto st668
		case 117:
			goto st672
		case 120:
			goto st672
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st612
				}
			case data[p] >= 9:
				goto tr675
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st612
				}
			case data[p] >= 65:
				goto st612
			}
		default:
			goto tr675
		}
		goto tr674
tr675:
//line xss_170.rl:13

            return true
        
	goto st886
	st886:
		if p++; p == pe {
			goto _test_eof886
		}
	st_case_886:
//line xss_170.go:26037
		switch data[p] {
		case 32:
			goto tr675
		case 40:
			goto tr675
		case 44:
			goto tr278
		case 46:
			goto tr675
		case 91:
			goto tr675
		case 92:
			goto tr676
		case 95:
			goto st612
		case 98:
			goto st614
		case 100:
			goto st620
		case 110:
			goto st668
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st612
				}
			case data[p] >= 9:
				goto tr675
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st612
				}
			case data[p] >= 65:
				goto st612
			}
		default:
			goto tr675
		}
		goto tr674
tr676:
//line xss_170.rl:13

            return true
        
	goto st887
	st887:
		if p++; p == pe {
			goto _test_eof887
		}
	st_case_887:
//line xss_170.go:26094
		switch data[p] {
		case 32:
			goto tr675
		case 40:
			goto tr675
		case 44:
			goto tr278
		case 46:
			goto tr675
		case 91:
			goto tr675
		case 92:
			goto tr676
		case 95:
			goto st612
		case 98:
			goto st614
		case 100:
			goto st620
		case 110:
			goto st668
		case 117:
			goto st672
		case 120:
			goto st672
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st612
				}
			case data[p] >= 9:
				goto tr675
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st612
				}
			case data[p] >= 65:
				goto st612
			}
		default:
			goto tr675
		}
		goto tr674
	st614:
		if p++; p == pe {
			goto _test_eof614
		}
	st_case_614:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 95:
			goto st612
		case 97:
			goto st615
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st612
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st612
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st615:
		if p++; p == pe {
			goto _test_eof615
		}
	st_case_615:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 95:
			goto st612
		case 115:
			goto st616
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st612
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st612
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st616:
		if p++; p == pe {
			goto _test_eof616
		}
	st_case_616:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 95:
			goto st612
		case 101:
			goto st617
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st612
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st612
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st617:
		if p++; p == pe {
			goto _test_eof617
		}
	st_case_617:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 54:
			goto st618
		case 95:
			goto st612
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st612
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st612
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st618:
		if p++; p == pe {
			goto _test_eof618
		}
	st_case_618:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 52:
			goto st619
		case 95:
			goto st612
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st612
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st612
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st619:
		if p++; p == pe {
			goto _test_eof619
		}
	st_case_619:
		switch data[p] {
		case 32:
			goto tr675
		case 40:
			goto tr675
		case 44:
			goto tr278
		case 46:
			goto tr675
		case 95:
			goto st612
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st612
				}
			case data[p] >= 9:
				goto tr675
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st612
				}
			default:
				goto tr675
			}
		default:
			goto tr675
		}
		goto tr674
	st620:
		if p++; p == pe {
			goto _test_eof620
		}
	st_case_620:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 95:
			goto st612
		case 97:
			goto st621
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st612
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st612
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st621:
		if p++; p == pe {
			goto _test_eof621
		}
	st_case_621:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 95:
			goto st612
		case 116:
			goto st622
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st612
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st612
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st622:
		if p++; p == pe {
			goto _test_eof622
		}
	st_case_622:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 95:
			goto st612
		case 97:
			goto st623
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st612
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st612
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st623:
		if p++; p == pe {
			goto _test_eof623
		}
	st_case_623:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 58:
			goto st624
		case 95:
			goto st612
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st612
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st612
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st624:
		if p++; p == pe {
			goto _test_eof624
		}
	st_case_624:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st233
		case 46:
			goto st611
		case 91:
			goto st611
		case 92:
			goto st613
		case 95:
			goto st625
		case 98:
			goto st653
		case 100:
			goto st659
		case 110:
			goto st663
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st612
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st625
				}
			case data[p] >= 65:
				goto st612
			}
		default:
			goto st611
		}
		goto st610
	st625:
		if p++; p == pe {
			goto _test_eof625
		}
	st_case_625:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 95:
			goto st626
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st626
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st626
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st626:
		if p++; p == pe {
			goto _test_eof626
		}
	st_case_626:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 47:
			goto st627
		case 95:
			goto st626
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st626
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st626
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st627:
		if p++; p == pe {
			goto _test_eof627
		}
	st_case_627:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 91:
			goto st611
		case 92:
			goto st613
		case 95:
			goto st628
		case 98:
			goto st647
		case 100:
			goto st649
		case 110:
			goto st651
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st628
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st628
				}
			case data[p] >= 65:
				goto st628
			}
		default:
			goto st611
		}
		goto st610
	st628:
		if p++; p == pe {
			goto _test_eof628
		}
	st_case_628:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 95:
			goto st646
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st611
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st646
				}
			default:
				goto st629
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st646
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st646
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st629:
		if p++; p == pe {
			goto _test_eof629
		}
	st_case_629:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 91:
			goto st611
		case 92:
			goto st613
		case 95:
			goto st630
		case 98:
			goto st631
		case 100:
			goto st637
		case 110:
			goto st641
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 13:
				if 43 <= data[p] && data[p] <= 45 {
					goto st629
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 57:
			switch {
			case data[p] < 65:
				if 60 <= data[p] && data[p] <= 61 {
					goto st611
				}
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st630
				}
			default:
				goto st630
			}
		default:
			goto st630
		}
		goto st610
	st630:
		if p++; p == pe {
			goto _test_eof630
		}
	st_case_630:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st233
		case 46:
			goto st611
		case 95:
			goto st630
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st611
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st630
				}
			default:
				goto st629
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st630
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st630
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st631:
		if p++; p == pe {
			goto _test_eof631
		}
	st_case_631:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st233
		case 46:
			goto st611
		case 95:
			goto st630
		case 97:
			goto st632
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st611
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st630
				}
			default:
				goto st629
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st630
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st630
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st632:
		if p++; p == pe {
			goto _test_eof632
		}
	st_case_632:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st233
		case 46:
			goto st611
		case 95:
			goto st630
		case 115:
			goto st633
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st611
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st630
				}
			default:
				goto st629
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st630
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st630
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st633:
		if p++; p == pe {
			goto _test_eof633
		}
	st_case_633:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st233
		case 46:
			goto st611
		case 95:
			goto st630
		case 101:
			goto st634
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st611
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st630
				}
			default:
				goto st629
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st630
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st630
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st634:
		if p++; p == pe {
			goto _test_eof634
		}
	st_case_634:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st233
		case 46:
			goto st611
		case 54:
			goto st635
		case 95:
			goto st630
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st611
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st630
				}
			default:
				goto st629
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st630
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st630
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st635:
		if p++; p == pe {
			goto _test_eof635
		}
	st_case_635:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st233
		case 46:
			goto st611
		case 52:
			goto st636
		case 95:
			goto st630
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st611
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st630
				}
			default:
				goto st629
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st630
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st630
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st636:
		if p++; p == pe {
			goto _test_eof636
		}
	st_case_636:
		switch data[p] {
		case 32:
			goto tr675
		case 40:
			goto tr675
		case 44:
			goto tr279
		case 46:
			goto tr675
		case 95:
			goto st630
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr675
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st630
				}
			default:
				goto tr708
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st630
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st630
				}
			default:
				goto tr675
			}
		default:
			goto tr675
		}
		goto tr674
tr708:
//line xss_170.rl:13

            return true
        
	goto st888
	st888:
		if p++; p == pe {
			goto _test_eof888
		}
	st_case_888:
//line xss_170.go:27251
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 91:
			goto st611
		case 92:
			goto st613
		case 95:
			goto st630
		case 98:
			goto st631
		case 100:
			goto st637
		case 110:
			goto st641
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 13:
				if 43 <= data[p] && data[p] <= 45 {
					goto st629
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 57:
			switch {
			case data[p] < 65:
				if 60 <= data[p] && data[p] <= 61 {
					goto st611
				}
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st630
				}
			default:
				goto st630
			}
		default:
			goto st630
		}
		goto st610
	st637:
		if p++; p == pe {
			goto _test_eof637
		}
	st_case_637:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st233
		case 46:
			goto st611
		case 95:
			goto st630
		case 97:
			goto st638
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st611
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st630
				}
			default:
				goto st629
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st630
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st630
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st638:
		if p++; p == pe {
			goto _test_eof638
		}
	st_case_638:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st233
		case 46:
			goto st611
		case 95:
			goto st630
		case 116:
			goto st639
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st611
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st630
				}
			default:
				goto st629
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st630
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st630
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st639:
		if p++; p == pe {
			goto _test_eof639
		}
	st_case_639:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st233
		case 46:
			goto st611
		case 95:
			goto st630
		case 97:
			goto st640
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st611
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st630
				}
			default:
				goto st629
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st630
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st630
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st640:
		if p++; p == pe {
			goto _test_eof640
		}
	st_case_640:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st233
		case 46:
			goto st611
		case 58:
			goto st624
		case 95:
			goto st630
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st611
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st630
				}
			default:
				goto st629
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st630
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st630
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st641:
		if p++; p == pe {
			goto _test_eof641
		}
	st_case_641:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st233
		case 46:
			goto st611
		case 95:
			goto st630
		case 97:
			goto st642
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st611
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st630
				}
			default:
				goto st629
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st630
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st630
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st642:
		if p++; p == pe {
			goto _test_eof642
		}
	st_case_642:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st233
		case 46:
			goto st611
		case 95:
			goto st630
		case 109:
			goto st643
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st611
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st630
				}
			default:
				goto st629
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st630
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st630
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st643:
		if p++; p == pe {
			goto _test_eof643
		}
	st_case_643:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st233
		case 46:
			goto st611
		case 95:
			goto st630
		case 101:
			goto st644
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st611
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st630
				}
			default:
				goto st629
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st630
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st630
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st644:
		if p++; p == pe {
			goto _test_eof644
		}
	st_case_644:
		switch data[p] {
		case 44:
			goto st233
		case 95:
			goto st630
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st645
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st630
				}
			case data[p] >= 65:
				goto st630
			}
		default:
			goto st630
		}
		goto st611
	st645:
		if p++; p == pe {
			goto _test_eof645
		}
	st_case_645:
		switch data[p] {
		case 32:
			goto tr675
		case 40:
			goto tr675
		case 44:
			goto tr278
		case 46:
			goto tr675
		case 91:
			goto tr675
		case 92:
			goto tr676
		case 95:
			goto st630
		case 98:
			goto st631
		case 100:
			goto st637
		case 110:
			goto st641
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 13:
				if 43 <= data[p] && data[p] <= 45 {
					goto tr708
				}
			case data[p] >= 9:
				goto tr675
			}
		case data[p] > 57:
			switch {
			case data[p] < 65:
				if 60 <= data[p] && data[p] <= 61 {
					goto tr675
				}
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st630
				}
			default:
				goto st630
			}
		default:
			goto st630
		}
		goto tr674
	st646:
		if p++; p == pe {
			goto _test_eof646
		}
	st_case_646:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 95:
			goto st630
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st611
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st630
				}
			default:
				goto st629
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st630
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st630
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st647:
		if p++; p == pe {
			goto _test_eof647
		}
	st_case_647:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 95:
			goto st646
		case 97:
			goto st648
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st611
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st646
				}
			default:
				goto st629
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st646
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st646
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st648:
		if p++; p == pe {
			goto _test_eof648
		}
	st_case_648:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 95:
			goto st630
		case 115:
			goto st633
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st611
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st630
				}
			default:
				goto st629
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st630
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st630
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st649:
		if p++; p == pe {
			goto _test_eof649
		}
	st_case_649:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 95:
			goto st646
		case 97:
			goto st650
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st611
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st646
				}
			default:
				goto st629
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st646
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st646
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st650:
		if p++; p == pe {
			goto _test_eof650
		}
	st_case_650:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 95:
			goto st630
		case 116:
			goto st639
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st611
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st630
				}
			default:
				goto st629
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st630
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st630
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st651:
		if p++; p == pe {
			goto _test_eof651
		}
	st_case_651:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 95:
			goto st646
		case 97:
			goto st652
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st611
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st646
				}
			default:
				goto st629
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st646
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st646
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st652:
		if p++; p == pe {
			goto _test_eof652
		}
	st_case_652:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 95:
			goto st630
		case 109:
			goto st643
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st611
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st630
				}
			default:
				goto st629
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st630
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st630
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st653:
		if p++; p == pe {
			goto _test_eof653
		}
	st_case_653:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 95:
			goto st626
		case 97:
			goto st654
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st626
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st626
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st654:
		if p++; p == pe {
			goto _test_eof654
		}
	st_case_654:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 47:
			goto st627
		case 95:
			goto st626
		case 115:
			goto st655
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st626
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st626
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st655:
		if p++; p == pe {
			goto _test_eof655
		}
	st_case_655:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 47:
			goto st627
		case 95:
			goto st626
		case 101:
			goto st656
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st626
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st626
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st656:
		if p++; p == pe {
			goto _test_eof656
		}
	st_case_656:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 47:
			goto st627
		case 54:
			goto st657
		case 95:
			goto st626
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st626
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st626
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st657:
		if p++; p == pe {
			goto _test_eof657
		}
	st_case_657:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 47:
			goto st627
		case 52:
			goto st658
		case 95:
			goto st626
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st626
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st626
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st658:
		if p++; p == pe {
			goto _test_eof658
		}
	st_case_658:
		switch data[p] {
		case 32:
			goto tr675
		case 40:
			goto tr675
		case 44:
			goto tr278
		case 46:
			goto tr675
		case 47:
			goto tr724
		case 95:
			goto st626
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st626
				}
			case data[p] >= 9:
				goto tr675
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st626
				}
			default:
				goto tr675
			}
		default:
			goto tr675
		}
		goto tr674
tr724:
//line xss_170.rl:13

            return true
        
	goto st889
	st889:
		if p++; p == pe {
			goto _test_eof889
		}
	st_case_889:
//line xss_170.go:28377
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 91:
			goto st611
		case 92:
			goto st613
		case 95:
			goto st628
		case 98:
			goto st647
		case 100:
			goto st649
		case 110:
			goto st651
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st628
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st628
				}
			case data[p] >= 65:
				goto st628
			}
		default:
			goto st611
		}
		goto st610
	st659:
		if p++; p == pe {
			goto _test_eof659
		}
	st_case_659:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 95:
			goto st626
		case 97:
			goto st660
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st626
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st626
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st660:
		if p++; p == pe {
			goto _test_eof660
		}
	st_case_660:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 47:
			goto st627
		case 95:
			goto st626
		case 116:
			goto st661
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st626
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st626
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st661:
		if p++; p == pe {
			goto _test_eof661
		}
	st_case_661:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 47:
			goto st627
		case 95:
			goto st626
		case 97:
			goto st662
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st626
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st626
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st662:
		if p++; p == pe {
			goto _test_eof662
		}
	st_case_662:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 47:
			goto st627
		case 58:
			goto st624
		case 95:
			goto st626
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st626
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st626
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st663:
		if p++; p == pe {
			goto _test_eof663
		}
	st_case_663:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 95:
			goto st626
		case 97:
			goto st664
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st626
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st626
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st664:
		if p++; p == pe {
			goto _test_eof664
		}
	st_case_664:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 47:
			goto st627
		case 95:
			goto st626
		case 109:
			goto st665
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st626
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st626
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st665:
		if p++; p == pe {
			goto _test_eof665
		}
	st_case_665:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 47:
			goto st627
		case 95:
			goto st626
		case 101:
			goto st666
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st626
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st626
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st666:
		if p++; p == pe {
			goto _test_eof666
		}
	st_case_666:
		switch data[p] {
		case 44:
			goto st233
		case 47:
			goto st667
		case 95:
			goto st626
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st626
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st626
			}
		default:
			goto st612
		}
		goto st611
	st667:
		if p++; p == pe {
			goto _test_eof667
		}
	st_case_667:
		switch data[p] {
		case 32:
			goto tr675
		case 40:
			goto tr675
		case 44:
			goto tr278
		case 46:
			goto tr675
		case 91:
			goto tr675
		case 92:
			goto tr676
		case 95:
			goto st628
		case 98:
			goto st647
		case 100:
			goto st649
		case 110:
			goto st651
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st628
				}
			case data[p] >= 9:
				goto tr675
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st628
				}
			case data[p] >= 65:
				goto st628
			}
		default:
			goto tr675
		}
		goto tr674
	st668:
		if p++; p == pe {
			goto _test_eof668
		}
	st_case_668:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 95:
			goto st612
		case 97:
			goto st669
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st612
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st612
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st669:
		if p++; p == pe {
			goto _test_eof669
		}
	st_case_669:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 95:
			goto st612
		case 109:
			goto st670
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st612
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st612
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st670:
		if p++; p == pe {
			goto _test_eof670
		}
	st_case_670:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 95:
			goto st612
		case 101:
			goto st671
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st612
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st612
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st671:
		if p++; p == pe {
			goto _test_eof671
		}
	st_case_671:
		switch data[p] {
		case 44:
			goto st233
		case 95:
			goto st612
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st612
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st612
			}
		default:
			goto st612
		}
		goto st611
	st672:
		if p++; p == pe {
			goto _test_eof672
		}
	st_case_672:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 95:
			goto st612
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st673
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st612
				}
			default:
				goto st611
			}
		default:
			goto st611
		}
		goto st610
	st673:
		if p++; p == pe {
			goto _test_eof673
		}
	st_case_673:
		switch data[p] {
		case 32:
			goto tr675
		case 40:
			goto tr675
		case 44:
			goto tr278
		case 46:
			goto tr675
		case 95:
			goto st612
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st673
				}
			case data[p] >= 9:
				goto tr675
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st612
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st612
				}
			default:
				goto tr675
			}
		default:
			goto tr675
		}
		goto tr674
	st674:
		if p++; p == pe {
			goto _test_eof674
		}
	st_case_674:
		switch data[p] {
		case 32:
			goto st611
		case 40:
			goto st611
		case 44:
			goto st232
		case 46:
			goto st611
		case 91:
			goto st611
		case 92:
			goto st613
		case 95:
			goto st612
		case 98:
			goto st614
		case 100:
			goto st620
		case 110:
			goto st668
		case 117:
			goto st672
		case 120:
			goto st672
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st612
				}
			case data[p] >= 9:
				goto st611
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st612
				}
			case data[p] >= 65:
				goto st612
			}
		default:
			goto st611
		}
		goto st610
tr646:
//line xss_170.rl:13

            return true
        
	goto st890
	st890:
		if p++; p == pe {
			goto _test_eof890
		}
	st_case_890:
//line xss_170.go:29146
		switch data[p] {
		case 44:
			goto tr210
		case 95:
			goto st556
		case 98:
			goto st557
		case 99:
			goto st563
		case 100:
			goto st571
		case 106:
			goto st675
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto tr610
	st675:
		if p++; p == pe {
			goto _test_eof675
		}
	st_case_675:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st556
		case 97:
			goto st676
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto st555
	st676:
		if p++; p == pe {
			goto _test_eof676
		}
	st_case_676:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st556
		case 118:
			goto st677
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto st555
	st677:
		if p++; p == pe {
			goto _test_eof677
		}
	st_case_677:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st556
		case 97:
			goto st678
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto st555
	st678:
		if p++; p == pe {
			goto _test_eof678
		}
	st_case_678:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st556
		case 115:
			goto st679
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto st555
	st679:
		if p++; p == pe {
			goto _test_eof679
		}
	st_case_679:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st556
		case 99:
			goto st680
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto st555
	st680:
		if p++; p == pe {
			goto _test_eof680
		}
	st_case_680:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st556
		case 114:
			goto st681
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto st555
	st681:
		if p++; p == pe {
			goto _test_eof681
		}
	st_case_681:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st556
		case 105:
			goto st682
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto st555
	st682:
		if p++; p == pe {
			goto _test_eof682
		}
	st_case_682:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st556
		case 112:
			goto st683
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto st555
	st683:
		if p++; p == pe {
			goto _test_eof683
		}
	st_case_683:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st556
		case 116:
			goto st684
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto st555
	st684:
		if p++; p == pe {
			goto _test_eof684
		}
	st_case_684:
		switch data[p] {
		case 44:
			goto st10
		case 58:
			goto st609
		case 95:
			goto st556
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st556
			}
		default:
			goto st556
		}
		goto st555
	st685:
		if p++; p == pe {
			goto _test_eof685
		}
	st_case_685:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st581
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st686:
		if p++; p == pe {
			goto _test_eof686
		}
	st_case_686:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st685
		case 97:
			goto st687
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st685
				}
			case data[p] >= 65:
				goto st685
			}
		default:
			goto st685
		}
		goto st555
	st687:
		if p++; p == pe {
			goto _test_eof687
		}
	st_case_687:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st581
		case 115:
			goto st584
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st688:
		if p++; p == pe {
			goto _test_eof688
		}
	st_case_688:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st685
		case 104:
			goto st689
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st685
				}
			case data[p] >= 65:
				goto st685
			}
		default:
			goto st685
		}
		goto st555
	st689:
		if p++; p == pe {
			goto _test_eof689
		}
	st_case_689:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st581
		case 97:
			goto st590
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st690:
		if p++; p == pe {
			goto _test_eof690
		}
	st_case_690:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st685
		case 97:
			goto st691
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st685
				}
			case data[p] >= 65:
				goto st685
			}
		default:
			goto st685
		}
		goto st555
	st691:
		if p++; p == pe {
			goto _test_eof691
		}
	st_case_691:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st581
		case 116:
			goto st597
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st692:
		if p++; p == pe {
			goto _test_eof692
		}
	st_case_692:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st685
		case 97:
			goto st693
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st685
				}
			case data[p] >= 65:
				goto st685
			}
		default:
			goto st685
		}
		goto st555
	st693:
		if p++; p == pe {
			goto _test_eof693
		}
	st_case_693:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st581
		case 118:
			goto st601
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st580
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st581
				}
			case data[p] >= 65:
				goto st581
			}
		default:
			goto st581
		}
		goto st555
	st694:
		if p++; p == pe {
			goto _test_eof694
		}
	st_case_694:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st577
		case 97:
			goto st695
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st577
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st577
			}
		default:
			goto st556
		}
		goto st555
	st695:
		if p++; p == pe {
			goto _test_eof695
		}
	st_case_695:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st578
		case 95:
			goto st577
		case 115:
			goto st696
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st577
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st577
			}
		default:
			goto st556
		}
		goto st555
	st696:
		if p++; p == pe {
			goto _test_eof696
		}
	st_case_696:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st578
		case 95:
			goto st577
		case 101:
			goto st697
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st577
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st577
			}
		default:
			goto st556
		}
		goto st555
	st697:
		if p++; p == pe {
			goto _test_eof697
		}
	st_case_697:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st578
		case 54:
			goto st698
		case 95:
			goto st577
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st577
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st577
			}
		default:
			goto st556
		}
		goto st555
	st698:
		if p++; p == pe {
			goto _test_eof698
		}
	st_case_698:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st578
		case 52:
			goto st699
		case 95:
			goto st577
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st577
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st577
			}
		default:
			goto st556
		}
		goto st555
	st699:
		if p++; p == pe {
			goto _test_eof699
		}
	st_case_699:
		switch data[p] {
		case 44:
			goto tr210
		case 47:
			goto tr754
		case 95:
			goto st577
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st577
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st577
			}
		default:
			goto st556
		}
		goto tr610
tr754:
//line xss_170.rl:13

            return true
        
	goto st891
	st891:
		if p++; p == pe {
			goto _test_eof891
		}
	st_case_891:
//line xss_170.go:29886
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st579
		case 98:
			goto st686
		case 99:
			goto st688
		case 100:
			goto st690
		case 106:
			goto st692
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st579
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st579
			}
		default:
			goto st579
		}
		goto st555
	st700:
		if p++; p == pe {
			goto _test_eof700
		}
	st_case_700:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st577
		case 104:
			goto st701
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st577
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st577
			}
		default:
			goto st556
		}
		goto st555
	st701:
		if p++; p == pe {
			goto _test_eof701
		}
	st_case_701:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st578
		case 95:
			goto st577
		case 97:
			goto st702
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st577
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st577
			}
		default:
			goto st556
		}
		goto st555
	st702:
		if p++; p == pe {
			goto _test_eof702
		}
	st_case_702:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st578
		case 95:
			goto st577
		case 114:
			goto st703
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st577
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st577
			}
		default:
			goto st556
		}
		goto st555
	st703:
		if p++; p == pe {
			goto _test_eof703
		}
	st_case_703:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st578
		case 95:
			goto st577
		case 115:
			goto st704
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st577
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st577
			}
		default:
			goto st556
		}
		goto st555
	st704:
		if p++; p == pe {
			goto _test_eof704
		}
	st_case_704:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st578
		case 95:
			goto st577
		case 101:
			goto st705
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st577
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st577
			}
		default:
			goto st556
		}
		goto st555
	st705:
		if p++; p == pe {
			goto _test_eof705
		}
	st_case_705:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st578
		case 95:
			goto st577
		case 116:
			goto st706
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st577
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st577
			}
		default:
			goto st556
		}
		goto st555
	st706:
		if p++; p == pe {
			goto _test_eof706
		}
	st_case_706:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st578
		case 61:
			goto st570
		case 95:
			goto st577
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st577
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st577
			}
		default:
			goto st556
		}
		goto st555
	st707:
		if p++; p == pe {
			goto _test_eof707
		}
	st_case_707:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st577
		case 97:
			goto st708
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st577
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st577
			}
		default:
			goto st556
		}
		goto st555
	st708:
		if p++; p == pe {
			goto _test_eof708
		}
	st_case_708:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st578
		case 95:
			goto st577
		case 116:
			goto st709
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st577
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st577
			}
		default:
			goto st556
		}
		goto st555
	st709:
		if p++; p == pe {
			goto _test_eof709
		}
	st_case_709:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st578
		case 95:
			goto st577
		case 97:
			goto st710
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st577
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st577
			}
		default:
			goto st556
		}
		goto st555
	st710:
		if p++; p == pe {
			goto _test_eof710
		}
	st_case_710:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st578
		case 58:
			goto st575
		case 95:
			goto st577
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st577
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st577
			}
		default:
			goto st556
		}
		goto st555
	st711:
		if p++; p == pe {
			goto _test_eof711
		}
	st_case_711:
		switch data[p] {
		case 44:
			goto st10
		case 95:
			goto st577
		case 97:
			goto st712
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st577
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st577
			}
		default:
			goto st556
		}
		goto st555
	st712:
		if p++; p == pe {
			goto _test_eof712
		}
	st_case_712:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st578
		case 95:
			goto st577
		case 118:
			goto st713
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st577
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st577
			}
		default:
			goto st556
		}
		goto st555
	st713:
		if p++; p == pe {
			goto _test_eof713
		}
	st_case_713:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st578
		case 95:
			goto st577
		case 97:
			goto st714
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st577
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st577
			}
		default:
			goto st556
		}
		goto st555
	st714:
		if p++; p == pe {
			goto _test_eof714
		}
	st_case_714:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st578
		case 95:
			goto st577
		case 115:
			goto st715
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st577
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st577
			}
		default:
			goto st556
		}
		goto st555
	st715:
		if p++; p == pe {
			goto _test_eof715
		}
	st_case_715:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st578
		case 95:
			goto st577
		case 99:
			goto st716
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st577
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st577
			}
		default:
			goto st556
		}
		goto st555
	st716:
		if p++; p == pe {
			goto _test_eof716
		}
	st_case_716:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st578
		case 95:
			goto st577
		case 114:
			goto st717
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st577
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st577
			}
		default:
			goto st556
		}
		goto st555
	st717:
		if p++; p == pe {
			goto _test_eof717
		}
	st_case_717:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st578
		case 95:
			goto st577
		case 105:
			goto st718
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st577
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st577
			}
		default:
			goto st556
		}
		goto st555
	st718:
		if p++; p == pe {
			goto _test_eof718
		}
	st_case_718:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st578
		case 95:
			goto st577
		case 112:
			goto st719
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st577
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st577
			}
		default:
			goto st556
		}
		goto st555
	st719:
		if p++; p == pe {
			goto _test_eof719
		}
	st_case_719:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st578
		case 95:
			goto st577
		case 116:
			goto st720
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st577
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st577
			}
		default:
			goto st556
		}
		goto st555
	st720:
		if p++; p == pe {
			goto _test_eof720
		}
	st_case_720:
		switch data[p] {
		case 44:
			goto st10
		case 47:
			goto st578
		case 58:
			goto st609
		case 95:
			goto st577
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st577
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st577
			}
		default:
			goto st556
		}
		goto st555
	st721:
		if p++; p == pe {
			goto _test_eof721
		}
	st_case_721:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st555
		case 95:
			goto st554
		case 97:
			goto st722
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st554
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st554
			}
		default:
			goto st554
		}
		goto st7
	st722:
		if p++; p == pe {
			goto _test_eof722
		}
	st_case_722:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st555
		case 95:
			goto st554
		case 116:
			goto st723
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st554
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st554
			}
		default:
			goto st554
		}
		goto st7
	st723:
		if p++; p == pe {
			goto _test_eof723
		}
	st_case_723:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st555
		case 95:
			goto st554
		case 97:
			goto st724
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st554
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st554
			}
		default:
			goto st554
		}
		goto st7
	st724:
		if p++; p == pe {
			goto _test_eof724
		}
	st_case_724:
		switch data[p] {
		case 44:
			goto st8
		case 58:
			goto st6
		case 59:
			goto st555
		case 95:
			goto st554
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st554
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st554
			}
		default:
			goto st554
		}
		goto st7
	st725:
		if p++; p == pe {
			goto _test_eof725
		}
	st_case_725:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st555
		case 95:
			goto st554
		case 97:
			goto st726
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st554
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st554
			}
		default:
			goto st554
		}
		goto st7
	st726:
		if p++; p == pe {
			goto _test_eof726
		}
	st_case_726:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st555
		case 95:
			goto st554
		case 118:
			goto st727
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st554
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st554
			}
		default:
			goto st554
		}
		goto st7
	st727:
		if p++; p == pe {
			goto _test_eof727
		}
	st_case_727:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st555
		case 95:
			goto st554
		case 97:
			goto st728
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st554
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st554
			}
		default:
			goto st554
		}
		goto st7
	st728:
		if p++; p == pe {
			goto _test_eof728
		}
	st_case_728:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st555
		case 95:
			goto st554
		case 115:
			goto st729
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st554
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st554
			}
		default:
			goto st554
		}
		goto st7
	st729:
		if p++; p == pe {
			goto _test_eof729
		}
	st_case_729:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st555
		case 95:
			goto st554
		case 99:
			goto st730
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st554
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st554
			}
		default:
			goto st554
		}
		goto st7
	st730:
		if p++; p == pe {
			goto _test_eof730
		}
	st_case_730:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st555
		case 95:
			goto st554
		case 114:
			goto st731
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st554
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st554
			}
		default:
			goto st554
		}
		goto st7
	st731:
		if p++; p == pe {
			goto _test_eof731
		}
	st_case_731:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st555
		case 95:
			goto st554
		case 105:
			goto st732
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st554
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st554
			}
		default:
			goto st554
		}
		goto st7
	st732:
		if p++; p == pe {
			goto _test_eof732
		}
	st_case_732:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st555
		case 95:
			goto st554
		case 112:
			goto st733
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st554
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st554
			}
		default:
			goto st554
		}
		goto st7
	st733:
		if p++; p == pe {
			goto _test_eof733
		}
	st_case_733:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st555
		case 95:
			goto st554
		case 116:
			goto st734
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st554
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st554
			}
		default:
			goto st554
		}
		goto st7
	st734:
		if p++; p == pe {
			goto _test_eof734
		}
	st_case_734:
		switch data[p] {
		case 44:
			goto st8
		case 58:
			goto st735
		case 59:
			goto st555
		case 95:
			goto st554
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st554
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st554
			}
		default:
			goto st554
		}
		goto st7
	st735:
		if p++; p == pe {
			goto _test_eof735
		}
	st_case_735:
		switch data[p] {
		case 44:
			goto st467
		case 59:
			goto st610
		case 92:
			goto st780
		case 95:
			goto st738
		case 100:
			goto st740
		case 110:
			goto st774
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st738
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st738
			}
		default:
			goto st738
		}
		goto st736
	st736:
		if p++; p == pe {
			goto _test_eof736
		}
	st_case_736:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st467
		case 46:
			goto st737
		case 59:
			goto st610
		case 91:
			goto st737
		case 92:
			goto st739
		case 95:
			goto st738
		case 100:
			goto st740
		case 110:
			goto st774
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st738
				}
			case data[p] >= 9:
				goto st737
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st738
				}
			case data[p] >= 65:
				goto st738
			}
		default:
			goto st737
		}
		goto st736
	st737:
		if p++; p == pe {
			goto _test_eof737
		}
	st_case_737:
		switch data[p] {
		case 32:
			goto tr794
		case 40:
			goto tr794
		case 44:
			goto tr518
		case 46:
			goto tr794
		case 59:
			goto tr674
		case 91:
			goto tr794
		case 92:
			goto tr795
		case 95:
			goto st738
		case 100:
			goto st740
		case 110:
			goto st774
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st738
				}
			case data[p] >= 9:
				goto tr794
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st738
				}
			case data[p] >= 65:
				goto st738
			}
		default:
			goto tr794
		}
		goto tr793
tr793:
//line xss_170.rl:13

            return true
        
	goto st892
	st892:
		if p++; p == pe {
			goto _test_eof892
		}
	st_case_892:
//line xss_170.go:31031
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st467
		case 46:
			goto st737
		case 59:
			goto st610
		case 91:
			goto st737
		case 92:
			goto st739
		case 95:
			goto st738
		case 100:
			goto st740
		case 110:
			goto st774
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st738
				}
			case data[p] >= 9:
				goto st737
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st738
				}
			case data[p] >= 65:
				goto st738
			}
		default:
			goto st737
		}
		goto st736
	st738:
		if p++; p == pe {
			goto _test_eof738
		}
	st_case_738:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st467
		case 46:
			goto st737
		case 59:
			goto st610
		case 95:
			goto st738
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st738
				}
			case data[p] >= 9:
				goto st737
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st738
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st738
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st739:
		if p++; p == pe {
			goto _test_eof739
		}
	st_case_739:
		switch data[p] {
		case 32:
			goto tr794
		case 40:
			goto tr794
		case 44:
			goto tr518
		case 46:
			goto tr794
		case 59:
			goto tr674
		case 91:
			goto tr794
		case 92:
			goto tr795
		case 95:
			goto st738
		case 100:
			goto st740
		case 110:
			goto st774
		case 117:
			goto st778
		case 120:
			goto st778
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st738
				}
			case data[p] >= 9:
				goto tr794
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st738
				}
			case data[p] >= 65:
				goto st738
			}
		default:
			goto tr794
		}
		goto tr793
tr794:
//line xss_170.rl:13

            return true
        
	goto st893
	st893:
		if p++; p == pe {
			goto _test_eof893
		}
	st_case_893:
//line xss_170.go:31188
		switch data[p] {
		case 32:
			goto tr794
		case 40:
			goto tr794
		case 44:
			goto tr518
		case 46:
			goto tr794
		case 59:
			goto tr674
		case 91:
			goto tr794
		case 92:
			goto tr795
		case 95:
			goto st738
		case 100:
			goto st740
		case 110:
			goto st774
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st738
				}
			case data[p] >= 9:
				goto tr794
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st738
				}
			case data[p] >= 65:
				goto st738
			}
		default:
			goto tr794
		}
		goto tr793
tr795:
//line xss_170.rl:13

            return true
        
	goto st894
	st894:
		if p++; p == pe {
			goto _test_eof894
		}
	st_case_894:
//line xss_170.go:31245
		switch data[p] {
		case 32:
			goto tr794
		case 40:
			goto tr794
		case 44:
			goto tr518
		case 46:
			goto tr794
		case 59:
			goto tr674
		case 91:
			goto tr794
		case 92:
			goto tr795
		case 95:
			goto st738
		case 100:
			goto st740
		case 110:
			goto st774
		case 117:
			goto st778
		case 120:
			goto st778
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st738
				}
			case data[p] >= 9:
				goto tr794
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st738
				}
			case data[p] >= 65:
				goto st738
			}
		default:
			goto tr794
		}
		goto tr793
	st740:
		if p++; p == pe {
			goto _test_eof740
		}
	st_case_740:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st467
		case 46:
			goto st737
		case 59:
			goto st610
		case 95:
			goto st738
		case 97:
			goto st741
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st738
				}
			case data[p] >= 9:
				goto st737
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st738
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st738
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st741:
		if p++; p == pe {
			goto _test_eof741
		}
	st_case_741:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st467
		case 46:
			goto st737
		case 59:
			goto st610
		case 95:
			goto st738
		case 116:
			goto st742
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st738
				}
			case data[p] >= 9:
				goto st737
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st738
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st738
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st742:
		if p++; p == pe {
			goto _test_eof742
		}
	st_case_742:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st467
		case 46:
			goto st737
		case 59:
			goto st610
		case 95:
			goto st738
		case 97:
			goto st743
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st738
				}
			case data[p] >= 9:
				goto st737
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st738
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st738
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st743:
		if p++; p == pe {
			goto _test_eof743
		}
	st_case_743:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st467
		case 46:
			goto st737
		case 58:
			goto st744
		case 59:
			goto st610
		case 95:
			goto st738
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st738
				}
			case data[p] >= 9:
				goto st737
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st738
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st738
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st744:
		if p++; p == pe {
			goto _test_eof744
		}
	st_case_744:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st468
		case 46:
			goto st737
		case 59:
			goto st611
		case 91:
			goto st737
		case 92:
			goto st739
		case 95:
			goto st745
		case 100:
			goto st765
		case 110:
			goto st769
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st738
				}
			case data[p] >= 9:
				goto st737
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st745
				}
			case data[p] >= 65:
				goto st738
			}
		default:
			goto st737
		}
		goto st736
	st745:
		if p++; p == pe {
			goto _test_eof745
		}
	st_case_745:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st467
		case 46:
			goto st737
		case 59:
			goto st610
		case 95:
			goto st746
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st746
				}
			case data[p] >= 9:
				goto st737
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st738
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st746
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st746:
		if p++; p == pe {
			goto _test_eof746
		}
	st_case_746:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st467
		case 46:
			goto st737
		case 47:
			goto st747
		case 59:
			goto st610
		case 95:
			goto st746
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st746
				}
			case data[p] >= 9:
				goto st737
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st738
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st746
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st747:
		if p++; p == pe {
			goto _test_eof747
		}
	st_case_747:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st467
		case 46:
			goto st737
		case 59:
			goto st610
		case 91:
			goto st737
		case 92:
			goto st739
		case 95:
			goto st748
		case 100:
			goto st761
		case 110:
			goto st763
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st748
				}
			case data[p] >= 9:
				goto st737
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st748
				}
			case data[p] >= 65:
				goto st748
			}
		default:
			goto st737
		}
		goto st736
	st748:
		if p++; p == pe {
			goto _test_eof748
		}
	st_case_748:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st467
		case 46:
			goto st737
		case 59:
			goto st610
		case 95:
			goto st760
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st737
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st760
				}
			default:
				goto st749
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st760
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st760
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st749:
		if p++; p == pe {
			goto _test_eof749
		}
	st_case_749:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st467
		case 46:
			goto st737
		case 59:
			goto st610
		case 91:
			goto st737
		case 92:
			goto st739
		case 95:
			goto st750
		case 100:
			goto st751
		case 110:
			goto st755
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 13:
				if 43 <= data[p] && data[p] <= 45 {
					goto st749
				}
			case data[p] >= 9:
				goto st737
			}
		case data[p] > 57:
			switch {
			case data[p] < 65:
				if 60 <= data[p] && data[p] <= 61 {
					goto st737
				}
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st750
				}
			default:
				goto st750
			}
		default:
			goto st750
		}
		goto st736
	st750:
		if p++; p == pe {
			goto _test_eof750
		}
	st_case_750:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st468
		case 46:
			goto st737
		case 59:
			goto st611
		case 95:
			goto st750
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st737
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st750
				}
			default:
				goto st749
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st750
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st750
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st751:
		if p++; p == pe {
			goto _test_eof751
		}
	st_case_751:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st468
		case 46:
			goto st737
		case 59:
			goto st611
		case 95:
			goto st750
		case 97:
			goto st752
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st737
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st750
				}
			default:
				goto st749
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st750
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st750
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st752:
		if p++; p == pe {
			goto _test_eof752
		}
	st_case_752:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st468
		case 46:
			goto st737
		case 59:
			goto st611
		case 95:
			goto st750
		case 116:
			goto st753
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st737
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st750
				}
			default:
				goto st749
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st750
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st750
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st753:
		if p++; p == pe {
			goto _test_eof753
		}
	st_case_753:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st468
		case 46:
			goto st737
		case 59:
			goto st611
		case 95:
			goto st750
		case 97:
			goto st754
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st737
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st750
				}
			default:
				goto st749
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st750
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st750
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st754:
		if p++; p == pe {
			goto _test_eof754
		}
	st_case_754:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st468
		case 46:
			goto st737
		case 58:
			goto st744
		case 59:
			goto st611
		case 95:
			goto st750
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st737
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st750
				}
			default:
				goto st749
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st750
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st750
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st755:
		if p++; p == pe {
			goto _test_eof755
		}
	st_case_755:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st468
		case 46:
			goto st737
		case 59:
			goto st611
		case 95:
			goto st750
		case 97:
			goto st756
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st737
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st750
				}
			default:
				goto st749
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st750
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st750
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st756:
		if p++; p == pe {
			goto _test_eof756
		}
	st_case_756:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st468
		case 46:
			goto st737
		case 59:
			goto st611
		case 95:
			goto st750
		case 109:
			goto st757
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st737
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st750
				}
			default:
				goto st749
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st750
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st750
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st757:
		if p++; p == pe {
			goto _test_eof757
		}
	st_case_757:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st468
		case 46:
			goto st737
		case 59:
			goto st611
		case 95:
			goto st750
		case 101:
			goto st758
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st737
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st750
				}
			default:
				goto st749
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st750
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st750
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st758:
		if p++; p == pe {
			goto _test_eof758
		}
	st_case_758:
		switch data[p] {
		case 44:
			goto st468
		case 59:
			goto st611
		case 95:
			goto st750
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st759
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st750
				}
			case data[p] >= 65:
				goto st750
			}
		default:
			goto st750
		}
		goto st737
	st759:
		if p++; p == pe {
			goto _test_eof759
		}
	st_case_759:
		switch data[p] {
		case 32:
			goto tr794
		case 40:
			goto tr794
		case 44:
			goto tr518
		case 46:
			goto tr794
		case 59:
			goto tr674
		case 91:
			goto tr794
		case 92:
			goto tr795
		case 95:
			goto st750
		case 100:
			goto st751
		case 110:
			goto st755
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 13:
				if 43 <= data[p] && data[p] <= 45 {
					goto tr821
				}
			case data[p] >= 9:
				goto tr794
			}
		case data[p] > 57:
			switch {
			case data[p] < 65:
				if 60 <= data[p] && data[p] <= 61 {
					goto tr794
				}
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st750
				}
			default:
				goto st750
			}
		default:
			goto st750
		}
		goto tr793
tr821:
//line xss_170.rl:13

            return true
        
	goto st895
	st895:
		if p++; p == pe {
			goto _test_eof895
		}
	st_case_895:
//line xss_170.go:32295
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st467
		case 46:
			goto st737
		case 59:
			goto st610
		case 91:
			goto st737
		case 92:
			goto st739
		case 95:
			goto st750
		case 100:
			goto st751
		case 110:
			goto st755
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 13:
				if 43 <= data[p] && data[p] <= 45 {
					goto st749
				}
			case data[p] >= 9:
				goto st737
			}
		case data[p] > 57:
			switch {
			case data[p] < 65:
				if 60 <= data[p] && data[p] <= 61 {
					goto st737
				}
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st750
				}
			default:
				goto st750
			}
		default:
			goto st750
		}
		goto st736
	st760:
		if p++; p == pe {
			goto _test_eof760
		}
	st_case_760:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st467
		case 46:
			goto st737
		case 59:
			goto st610
		case 95:
			goto st750
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st737
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st750
				}
			default:
				goto st749
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st750
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st750
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st761:
		if p++; p == pe {
			goto _test_eof761
		}
	st_case_761:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st467
		case 46:
			goto st737
		case 59:
			goto st610
		case 95:
			goto st760
		case 97:
			goto st762
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st737
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st760
				}
			default:
				goto st749
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st760
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st760
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st762:
		if p++; p == pe {
			goto _test_eof762
		}
	st_case_762:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st467
		case 46:
			goto st737
		case 59:
			goto st610
		case 95:
			goto st750
		case 116:
			goto st753
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st737
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st750
				}
			default:
				goto st749
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st750
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st750
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st763:
		if p++; p == pe {
			goto _test_eof763
		}
	st_case_763:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st467
		case 46:
			goto st737
		case 59:
			goto st610
		case 95:
			goto st760
		case 97:
			goto st764
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st737
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st760
				}
			default:
				goto st749
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st760
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st760
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st764:
		if p++; p == pe {
			goto _test_eof764
		}
	st_case_764:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st467
		case 46:
			goto st737
		case 59:
			goto st610
		case 95:
			goto st750
		case 109:
			goto st757
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] < 43:
				if 9 <= data[p] && data[p] <= 13 {
					goto st737
				}
			case data[p] > 45:
				if 48 <= data[p] && data[p] <= 57 {
					goto st750
				}
			default:
				goto st749
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st750
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st750
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st765:
		if p++; p == pe {
			goto _test_eof765
		}
	st_case_765:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st467
		case 46:
			goto st737
		case 59:
			goto st610
		case 95:
			goto st746
		case 97:
			goto st766
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st746
				}
			case data[p] >= 9:
				goto st737
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st738
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st746
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st766:
		if p++; p == pe {
			goto _test_eof766
		}
	st_case_766:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st467
		case 46:
			goto st737
		case 47:
			goto st747
		case 59:
			goto st610
		case 95:
			goto st746
		case 116:
			goto st767
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st746
				}
			case data[p] >= 9:
				goto st737
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st738
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st746
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st767:
		if p++; p == pe {
			goto _test_eof767
		}
	st_case_767:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st467
		case 46:
			goto st737
		case 47:
			goto st747
		case 59:
			goto st610
		case 95:
			goto st746
		case 97:
			goto st768
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st746
				}
			case data[p] >= 9:
				goto st737
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st738
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st746
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st768:
		if p++; p == pe {
			goto _test_eof768
		}
	st_case_768:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st467
		case 46:
			goto st737
		case 47:
			goto st747
		case 58:
			goto st744
		case 59:
			goto st610
		case 95:
			goto st746
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st746
				}
			case data[p] >= 9:
				goto st737
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st738
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st746
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st769:
		if p++; p == pe {
			goto _test_eof769
		}
	st_case_769:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st467
		case 46:
			goto st737
		case 59:
			goto st610
		case 95:
			goto st746
		case 97:
			goto st770
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st746
				}
			case data[p] >= 9:
				goto st737
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st738
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st746
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st770:
		if p++; p == pe {
			goto _test_eof770
		}
	st_case_770:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st467
		case 46:
			goto st737
		case 47:
			goto st747
		case 59:
			goto st610
		case 95:
			goto st746
		case 109:
			goto st771
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st746
				}
			case data[p] >= 9:
				goto st737
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st738
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st746
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st771:
		if p++; p == pe {
			goto _test_eof771
		}
	st_case_771:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st467
		case 46:
			goto st737
		case 47:
			goto st747
		case 59:
			goto st610
		case 95:
			goto st746
		case 101:
			goto st772
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st746
				}
			case data[p] >= 9:
				goto st737
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st738
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st746
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st772:
		if p++; p == pe {
			goto _test_eof772
		}
	st_case_772:
		switch data[p] {
		case 44:
			goto st468
		case 47:
			goto st773
		case 59:
			goto st611
		case 95:
			goto st746
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st746
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st746
			}
		default:
			goto st738
		}
		goto st737
	st773:
		if p++; p == pe {
			goto _test_eof773
		}
	st_case_773:
		switch data[p] {
		case 32:
			goto tr794
		case 40:
			goto tr794
		case 44:
			goto tr518
		case 46:
			goto tr794
		case 59:
			goto tr674
		case 91:
			goto tr794
		case 92:
			goto tr795
		case 95:
			goto st748
		case 100:
			goto st761
		case 110:
			goto st763
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st748
				}
			case data[p] >= 9:
				goto tr794
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st748
				}
			case data[p] >= 65:
				goto st748
			}
		default:
			goto tr794
		}
		goto tr793
	st774:
		if p++; p == pe {
			goto _test_eof774
		}
	st_case_774:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st467
		case 46:
			goto st737
		case 59:
			goto st610
		case 95:
			goto st738
		case 97:
			goto st775
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st738
				}
			case data[p] >= 9:
				goto st737
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st738
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st738
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st775:
		if p++; p == pe {
			goto _test_eof775
		}
	st_case_775:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st467
		case 46:
			goto st737
		case 59:
			goto st610
		case 95:
			goto st738
		case 109:
			goto st776
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st738
				}
			case data[p] >= 9:
				goto st737
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st738
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st738
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st776:
		if p++; p == pe {
			goto _test_eof776
		}
	st_case_776:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st467
		case 46:
			goto st737
		case 59:
			goto st610
		case 95:
			goto st738
		case 101:
			goto st777
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st738
				}
			case data[p] >= 9:
				goto st737
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st738
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st738
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st777:
		if p++; p == pe {
			goto _test_eof777
		}
	st_case_777:
		switch data[p] {
		case 44:
			goto st468
		case 59:
			goto st611
		case 95:
			goto st738
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st738
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st738
			}
		default:
			goto st738
		}
		goto st737
	st778:
		if p++; p == pe {
			goto _test_eof778
		}
	st_case_778:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st467
		case 46:
			goto st737
		case 59:
			goto st610
		case 95:
			goto st738
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st779
				}
			case data[p] >= 9:
				goto st737
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st738
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st738
				}
			default:
				goto st737
			}
		default:
			goto st737
		}
		goto st736
	st779:
		if p++; p == pe {
			goto _test_eof779
		}
	st_case_779:
		switch data[p] {
		case 32:
			goto tr794
		case 40:
			goto tr794
		case 44:
			goto tr518
		case 46:
			goto tr794
		case 59:
			goto tr674
		case 95:
			goto st738
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st779
				}
			case data[p] >= 9:
				goto tr794
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st738
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st738
				}
			default:
				goto tr794
			}
		default:
			goto tr794
		}
		goto tr793
	st780:
		if p++; p == pe {
			goto _test_eof780
		}
	st_case_780:
		switch data[p] {
		case 32:
			goto st737
		case 40:
			goto st737
		case 44:
			goto st467
		case 46:
			goto st737
		case 59:
			goto st610
		case 91:
			goto st737
		case 92:
			goto st739
		case 95:
			goto st738
		case 100:
			goto st740
		case 110:
			goto st774
		case 117:
			goto st778
		case 120:
			goto st778
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st738
				}
			case data[p] >= 9:
				goto st737
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st738
				}
			case data[p] >= 65:
				goto st738
			}
		default:
			goto st737
		}
		goto st736
	st781:
		if p++; p == pe {
			goto _test_eof781
		}
	st_case_781:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st555
		case 95:
			goto st782
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st782
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st782
			}
		default:
			goto st554
		}
		goto st7
	st782:
		if p++; p == pe {
			goto _test_eof782
		}
	st_case_782:
		switch data[p] {
		case 44:
			goto st8
		case 47:
			goto st783
		case 59:
			goto st555
		case 95:
			goto st782
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st782
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st782
			}
		default:
			goto st554
		}
		goto st7
	st783:
		if p++; p == pe {
			goto _test_eof783
		}
	st_case_783:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st555
		case 95:
			goto st784
		case 100:
			goto st802
		case 106:
			goto st804
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st784
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st784
			}
		default:
			goto st784
		}
		goto st7
	st784:
		if p++; p == pe {
			goto _test_eof784
		}
	st_case_784:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st555
		case 95:
			goto st801
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st785
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st801
				}
			case data[p] >= 65:
				goto st801
			}
		default:
			goto st801
		}
		goto st7
	st785:
		if p++; p == pe {
			goto _test_eof785
		}
	st_case_785:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st555
		case 95:
			goto st786
		case 100:
			goto st787
		case 106:
			goto st791
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st785
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st786
				}
			case data[p] >= 65:
				goto st786
			}
		default:
			goto st786
		}
		goto st7
	st786:
		if p++; p == pe {
			goto _test_eof786
		}
	st_case_786:
		switch data[p] {
		case 44:
			goto st455
		case 59:
			goto st570
		case 95:
			goto st786
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st785
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st786
				}
			case data[p] >= 65:
				goto st786
			}
		default:
			goto st786
		}
		goto st7
	st787:
		if p++; p == pe {
			goto _test_eof787
		}
	st_case_787:
		switch data[p] {
		case 44:
			goto st455
		case 59:
			goto st570
		case 95:
			goto st786
		case 97:
			goto st788
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st785
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st786
				}
			case data[p] >= 65:
				goto st786
			}
		default:
			goto st786
		}
		goto st7
	st788:
		if p++; p == pe {
			goto _test_eof788
		}
	st_case_788:
		switch data[p] {
		case 44:
			goto st455
		case 59:
			goto st570
		case 95:
			goto st786
		case 116:
			goto st789
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st785
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st786
				}
			case data[p] >= 65:
				goto st786
			}
		default:
			goto st786
		}
		goto st7
	st789:
		if p++; p == pe {
			goto _test_eof789
		}
	st_case_789:
		switch data[p] {
		case 44:
			goto st455
		case 59:
			goto st570
		case 95:
			goto st786
		case 97:
			goto st790
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st785
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st786
				}
			case data[p] >= 65:
				goto st786
			}
		default:
			goto st786
		}
		goto st7
	st790:
		if p++; p == pe {
			goto _test_eof790
		}
	st_case_790:
		switch data[p] {
		case 44:
			goto st455
		case 58:
			goto st6
		case 59:
			goto st570
		case 95:
			goto st786
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st785
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st786
				}
			case data[p] >= 65:
				goto st786
			}
		default:
			goto st786
		}
		goto st7
	st791:
		if p++; p == pe {
			goto _test_eof791
		}
	st_case_791:
		switch data[p] {
		case 44:
			goto st455
		case 59:
			goto st570
		case 95:
			goto st786
		case 97:
			goto st792
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st785
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st786
				}
			case data[p] >= 65:
				goto st786
			}
		default:
			goto st786
		}
		goto st7
	st792:
		if p++; p == pe {
			goto _test_eof792
		}
	st_case_792:
		switch data[p] {
		case 44:
			goto st455
		case 59:
			goto st570
		case 95:
			goto st786
		case 118:
			goto st793
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st785
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st786
				}
			case data[p] >= 65:
				goto st786
			}
		default:
			goto st786
		}
		goto st7
	st793:
		if p++; p == pe {
			goto _test_eof793
		}
	st_case_793:
		switch data[p] {
		case 44:
			goto st455
		case 59:
			goto st570
		case 95:
			goto st786
		case 97:
			goto st794
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st785
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st786
				}
			case data[p] >= 65:
				goto st786
			}
		default:
			goto st786
		}
		goto st7
	st794:
		if p++; p == pe {
			goto _test_eof794
		}
	st_case_794:
		switch data[p] {
		case 44:
			goto st455
		case 59:
			goto st570
		case 95:
			goto st786
		case 115:
			goto st795
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st785
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st786
				}
			case data[p] >= 65:
				goto st786
			}
		default:
			goto st786
		}
		goto st7
	st795:
		if p++; p == pe {
			goto _test_eof795
		}
	st_case_795:
		switch data[p] {
		case 44:
			goto st455
		case 59:
			goto st570
		case 95:
			goto st786
		case 99:
			goto st796
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st785
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st786
				}
			case data[p] >= 65:
				goto st786
			}
		default:
			goto st786
		}
		goto st7
	st796:
		if p++; p == pe {
			goto _test_eof796
		}
	st_case_796:
		switch data[p] {
		case 44:
			goto st455
		case 59:
			goto st570
		case 95:
			goto st786
		case 114:
			goto st797
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st785
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st786
				}
			case data[p] >= 65:
				goto st786
			}
		default:
			goto st786
		}
		goto st7
	st797:
		if p++; p == pe {
			goto _test_eof797
		}
	st_case_797:
		switch data[p] {
		case 44:
			goto st455
		case 59:
			goto st570
		case 95:
			goto st786
		case 105:
			goto st798
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st785
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st786
				}
			case data[p] >= 65:
				goto st786
			}
		default:
			goto st786
		}
		goto st7
	st798:
		if p++; p == pe {
			goto _test_eof798
		}
	st_case_798:
		switch data[p] {
		case 44:
			goto st455
		case 59:
			goto st570
		case 95:
			goto st786
		case 112:
			goto st799
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st785
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st786
				}
			case data[p] >= 65:
				goto st786
			}
		default:
			goto st786
		}
		goto st7
	st799:
		if p++; p == pe {
			goto _test_eof799
		}
	st_case_799:
		switch data[p] {
		case 44:
			goto st455
		case 59:
			goto st570
		case 95:
			goto st786
		case 116:
			goto st800
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st785
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st786
				}
			case data[p] >= 65:
				goto st786
			}
		default:
			goto st786
		}
		goto st7
	st800:
		if p++; p == pe {
			goto _test_eof800
		}
	st_case_800:
		switch data[p] {
		case 44:
			goto st455
		case 58:
			goto st735
		case 59:
			goto st570
		case 95:
			goto st786
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st785
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st786
				}
			case data[p] >= 65:
				goto st786
			}
		default:
			goto st786
		}
		goto st7
	st801:
		if p++; p == pe {
			goto _test_eof801
		}
	st_case_801:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st555
		case 95:
			goto st786
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st785
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st786
				}
			case data[p] >= 65:
				goto st786
			}
		default:
			goto st786
		}
		goto st7
	st802:
		if p++; p == pe {
			goto _test_eof802
		}
	st_case_802:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st555
		case 95:
			goto st801
		case 97:
			goto st803
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st785
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st801
				}
			case data[p] >= 65:
				goto st801
			}
		default:
			goto st801
		}
		goto st7
	st803:
		if p++; p == pe {
			goto _test_eof803
		}
	st_case_803:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st555
		case 95:
			goto st786
		case 116:
			goto st789
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st785
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st786
				}
			case data[p] >= 65:
				goto st786
			}
		default:
			goto st786
		}
		goto st7
	st804:
		if p++; p == pe {
			goto _test_eof804
		}
	st_case_804:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st555
		case 95:
			goto st801
		case 97:
			goto st805
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st785
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto st801
				}
			case data[p] >= 65:
				goto st801
			}
		default:
			goto st801
		}
		goto st7
	st805:
		if p++; p == pe {
			goto _test_eof805
		}
	st_case_805:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st555
		case 95:
			goto st786
		case 118:
			goto st793
		}
		switch {
		case data[p] < 48:
			if 43 <= data[p] && data[p] <= 45 {
				goto st785
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st786
				}
			case data[p] >= 65:
				goto st786
			}
		default:
			goto st786
		}
		goto st7
	st806:
		if p++; p == pe {
			goto _test_eof806
		}
	st_case_806:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st555
		case 95:
			goto st782
		case 97:
			goto st807
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st782
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st782
			}
		default:
			goto st554
		}
		goto st7
	st807:
		if p++; p == pe {
			goto _test_eof807
		}
	st_case_807:
		switch data[p] {
		case 44:
			goto st8
		case 47:
			goto st783
		case 59:
			goto st555
		case 95:
			goto st782
		case 116:
			goto st808
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st782
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st782
			}
		default:
			goto st554
		}
		goto st7
	st808:
		if p++; p == pe {
			goto _test_eof808
		}
	st_case_808:
		switch data[p] {
		case 44:
			goto st8
		case 47:
			goto st783
		case 59:
			goto st555
		case 95:
			goto st782
		case 97:
			goto st809
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st782
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st782
			}
		default:
			goto st554
		}
		goto st7
	st809:
		if p++; p == pe {
			goto _test_eof809
		}
	st_case_809:
		switch data[p] {
		case 44:
			goto st8
		case 47:
			goto st783
		case 58:
			goto st6
		case 59:
			goto st555
		case 95:
			goto st782
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st782
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st782
			}
		default:
			goto st554
		}
		goto st7
	st810:
		if p++; p == pe {
			goto _test_eof810
		}
	st_case_810:
		switch data[p] {
		case 44:
			goto st8
		case 59:
			goto st555
		case 95:
			goto st782
		case 97:
			goto st811
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st782
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st782
			}
		default:
			goto st554
		}
		goto st7
	st811:
		if p++; p == pe {
			goto _test_eof811
		}
	st_case_811:
		switch data[p] {
		case 44:
			goto st8
		case 47:
			goto st783
		case 59:
			goto st555
		case 95:
			goto st782
		case 118:
			goto st812
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st782
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st782
			}
		default:
			goto st554
		}
		goto st7
	st812:
		if p++; p == pe {
			goto _test_eof812
		}
	st_case_812:
		switch data[p] {
		case 44:
			goto st8
		case 47:
			goto st783
		case 59:
			goto st555
		case 95:
			goto st782
		case 97:
			goto st813
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st782
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st782
			}
		default:
			goto st554
		}
		goto st7
	st813:
		if p++; p == pe {
			goto _test_eof813
		}
	st_case_813:
		switch data[p] {
		case 44:
			goto st8
		case 47:
			goto st783
		case 59:
			goto st555
		case 95:
			goto st782
		case 115:
			goto st814
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st782
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st782
			}
		default:
			goto st554
		}
		goto st7
	st814:
		if p++; p == pe {
			goto _test_eof814
		}
	st_case_814:
		switch data[p] {
		case 44:
			goto st8
		case 47:
			goto st783
		case 59:
			goto st555
		case 95:
			goto st782
		case 99:
			goto st815
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st782
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st782
			}
		default:
			goto st554
		}
		goto st7
	st815:
		if p++; p == pe {
			goto _test_eof815
		}
	st_case_815:
		switch data[p] {
		case 44:
			goto st8
		case 47:
			goto st783
		case 59:
			goto st555
		case 95:
			goto st782
		case 114:
			goto st816
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st782
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st782
			}
		default:
			goto st554
		}
		goto st7
	st816:
		if p++; p == pe {
			goto _test_eof816
		}
	st_case_816:
		switch data[p] {
		case 44:
			goto st8
		case 47:
			goto st783
		case 59:
			goto st555
		case 95:
			goto st782
		case 105:
			goto st817
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st782
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st782
			}
		default:
			goto st554
		}
		goto st7
	st817:
		if p++; p == pe {
			goto _test_eof817
		}
	st_case_817:
		switch data[p] {
		case 44:
			goto st8
		case 47:
			goto st783
		case 59:
			goto st555
		case 95:
			goto st782
		case 112:
			goto st818
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st782
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st782
			}
		default:
			goto st554
		}
		goto st7
	st818:
		if p++; p == pe {
			goto _test_eof818
		}
	st_case_818:
		switch data[p] {
		case 44:
			goto st8
		case 47:
			goto st783
		case 59:
			goto st555
		case 95:
			goto st782
		case 116:
			goto st819
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st782
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st782
			}
		default:
			goto st554
		}
		goto st7
	st819:
		if p++; p == pe {
			goto _test_eof819
		}
	st_case_819:
		switch data[p] {
		case 44:
			goto st8
		case 47:
			goto st783
		case 58:
			goto st735
		case 59:
			goto st555
		case 95:
			goto st782
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st782
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st782
			}
		default:
			goto st554
		}
		goto st7
	st820:
		if p++; p == pe {
			goto _test_eof820
		}
	st_case_820:
		switch data[p] {
		case 95:
			goto st1
		case 97:
			goto st821
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
	st821:
		if p++; p == pe {
			goto _test_eof821
		}
	st_case_821:
		switch data[p] {
		case 95:
			goto st1
		case 118:
			goto st822
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
	st822:
		if p++; p == pe {
			goto _test_eof822
		}
	st_case_822:
		switch data[p] {
		case 95:
			goto st1
		case 97:
			goto st823
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
	st823:
		if p++; p == pe {
			goto _test_eof823
		}
	st_case_823:
		switch data[p] {
		case 95:
			goto st1
		case 115:
			goto st824
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
	st824:
		if p++; p == pe {
			goto _test_eof824
		}
	st_case_824:
		switch data[p] {
		case 95:
			goto st1
		case 99:
			goto st825
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
	st825:
		if p++; p == pe {
			goto _test_eof825
		}
	st_case_825:
		switch data[p] {
		case 95:
			goto st1
		case 114:
			goto st826
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
	st826:
		if p++; p == pe {
			goto _test_eof826
		}
	st_case_826:
		switch data[p] {
		case 95:
			goto st1
		case 105:
			goto st827
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
	st827:
		if p++; p == pe {
			goto _test_eof827
		}
	st_case_827:
		switch data[p] {
		case 95:
			goto st1
		case 112:
			goto st828
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
	st828:
		if p++; p == pe {
			goto _test_eof828
		}
	st_case_828:
		switch data[p] {
		case 95:
			goto st1
		case 116:
			goto st829
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
	st829:
		if p++; p == pe {
			goto _test_eof829
		}
	st_case_829:
		switch data[p] {
		case 58:
			goto st830
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
	st830:
		if p++; p == pe {
			goto _test_eof830
		}
	st_case_830:
		switch data[p] {
		case 92:
			goto st845
		case 95:
			goto st833
		case 100:
			goto st835
		case 110:
			goto st839
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st833
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st833
			}
		default:
			goto st833
		}
		goto st831
	st831:
		if p++; p == pe {
			goto _test_eof831
		}
	st_case_831:
		switch data[p] {
		case 32:
			goto st832
		case 40:
			goto st832
		case 46:
			goto st832
		case 91:
			goto st832
		case 92:
			goto st834
		case 95:
			goto st833
		case 100:
			goto st835
		case 110:
			goto st839
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st833
				}
			case data[p] >= 9:
				goto st832
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st833
				}
			case data[p] >= 65:
				goto st833
			}
		default:
			goto st832
		}
		goto st831
	st832:
		if p++; p == pe {
			goto _test_eof832
		}
	st_case_832:
		switch data[p] {
		case 32:
			goto tr889
		case 40:
			goto tr889
		case 46:
			goto tr889
		case 91:
			goto tr889
		case 92:
			goto tr890
		case 95:
			goto st833
		case 100:
			goto st835
		case 110:
			goto st839
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st833
				}
			case data[p] >= 9:
				goto tr889
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st833
				}
			case data[p] >= 65:
				goto st833
			}
		default:
			goto tr889
		}
		goto tr888
tr888:
//line xss_170.rl:13

            return true
        
	goto st896
	st896:
		if p++; p == pe {
			goto _test_eof896
		}
	st_case_896:
//line xss_170.go:34936
		switch data[p] {
		case 32:
			goto st832
		case 40:
			goto st832
		case 46:
			goto st832
		case 91:
			goto st832
		case 92:
			goto st834
		case 95:
			goto st833
		case 100:
			goto st835
		case 110:
			goto st839
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st833
				}
			case data[p] >= 9:
				goto st832
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st833
				}
			case data[p] >= 65:
				goto st833
			}
		default:
			goto st832
		}
		goto st831
	st833:
		if p++; p == pe {
			goto _test_eof833
		}
	st_case_833:
		switch data[p] {
		case 32:
			goto st832
		case 40:
			goto st832
		case 46:
			goto st832
		case 95:
			goto st833
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st833
				}
			case data[p] >= 9:
				goto st832
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st833
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st833
				}
			default:
				goto st832
			}
		default:
			goto st832
		}
		goto st831
	st834:
		if p++; p == pe {
			goto _test_eof834
		}
	st_case_834:
		switch data[p] {
		case 32:
			goto tr889
		case 40:
			goto tr889
		case 46:
			goto tr889
		case 91:
			goto tr889
		case 92:
			goto tr890
		case 95:
			goto st833
		case 100:
			goto st835
		case 110:
			goto st839
		case 117:
			goto st843
		case 120:
			goto st843
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st833
				}
			case data[p] >= 9:
				goto tr889
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st833
				}
			case data[p] >= 65:
				goto st833
			}
		default:
			goto tr889
		}
		goto tr888
tr889:
//line xss_170.rl:13

            return true
        
	goto st897
	st897:
		if p++; p == pe {
			goto _test_eof897
		}
	st_case_897:
//line xss_170.go:35081
		switch data[p] {
		case 32:
			goto tr889
		case 40:
			goto tr889
		case 46:
			goto tr889
		case 91:
			goto tr889
		case 92:
			goto tr890
		case 95:
			goto st833
		case 100:
			goto st835
		case 110:
			goto st839
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st833
				}
			case data[p] >= 9:
				goto tr889
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st833
				}
			case data[p] >= 65:
				goto st833
			}
		default:
			goto tr889
		}
		goto tr888
tr890:
//line xss_170.rl:13

            return true
        
	goto st898
	st898:
		if p++; p == pe {
			goto _test_eof898
		}
	st_case_898:
//line xss_170.go:35134
		switch data[p] {
		case 32:
			goto tr889
		case 40:
			goto tr889
		case 46:
			goto tr889
		case 91:
			goto tr889
		case 92:
			goto tr890
		case 95:
			goto st833
		case 100:
			goto st835
		case 110:
			goto st839
		case 117:
			goto st843
		case 120:
			goto st843
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st833
				}
			case data[p] >= 9:
				goto tr889
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st833
				}
			case data[p] >= 65:
				goto st833
			}
		default:
			goto tr889
		}
		goto tr888
	st835:
		if p++; p == pe {
			goto _test_eof835
		}
	st_case_835:
		switch data[p] {
		case 32:
			goto st832
		case 40:
			goto st832
		case 46:
			goto st832
		case 95:
			goto st833
		case 97:
			goto st836
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st833
				}
			case data[p] >= 9:
				goto st832
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st833
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st833
				}
			default:
				goto st832
			}
		default:
			goto st832
		}
		goto st831
	st836:
		if p++; p == pe {
			goto _test_eof836
		}
	st_case_836:
		switch data[p] {
		case 32:
			goto st832
		case 40:
			goto st832
		case 46:
			goto st832
		case 95:
			goto st833
		case 116:
			goto st837
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st833
				}
			case data[p] >= 9:
				goto st832
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st833
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st833
				}
			default:
				goto st832
			}
		default:
			goto st832
		}
		goto st831
	st837:
		if p++; p == pe {
			goto _test_eof837
		}
	st_case_837:
		switch data[p] {
		case 32:
			goto st832
		case 40:
			goto st832
		case 46:
			goto st832
		case 95:
			goto st833
		case 97:
			goto st838
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st833
				}
			case data[p] >= 9:
				goto st832
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st833
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st833
				}
			default:
				goto st832
			}
		default:
			goto st832
		}
		goto st831
	st838:
		if p++; p == pe {
			goto _test_eof838
		}
	st_case_838:
		switch data[p] {
		case 32:
			goto st832
		case 40:
			goto st832
		case 46:
			goto st832
		case 58:
			goto st744
		case 95:
			goto st833
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st833
				}
			case data[p] >= 9:
				goto st832
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st833
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st833
				}
			default:
				goto st832
			}
		default:
			goto st832
		}
		goto st831
	st839:
		if p++; p == pe {
			goto _test_eof839
		}
	st_case_839:
		switch data[p] {
		case 32:
			goto st832
		case 40:
			goto st832
		case 46:
			goto st832
		case 95:
			goto st833
		case 97:
			goto st840
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st833
				}
			case data[p] >= 9:
				goto st832
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st833
				}
			case data[p] > 92:
				if 98 <= data[p] && data[p] <= 122 {
					goto st833
				}
			default:
				goto st832
			}
		default:
			goto st832
		}
		goto st831
	st840:
		if p++; p == pe {
			goto _test_eof840
		}
	st_case_840:
		switch data[p] {
		case 32:
			goto st832
		case 40:
			goto st832
		case 46:
			goto st832
		case 95:
			goto st833
		case 109:
			goto st841
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st833
				}
			case data[p] >= 9:
				goto st832
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st833
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st833
				}
			default:
				goto st832
			}
		default:
			goto st832
		}
		goto st831
	st841:
		if p++; p == pe {
			goto _test_eof841
		}
	st_case_841:
		switch data[p] {
		case 32:
			goto st832
		case 40:
			goto st832
		case 46:
			goto st832
		case 95:
			goto st833
		case 101:
			goto st842
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st833
				}
			case data[p] >= 9:
				goto st832
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st833
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st833
				}
			default:
				goto st832
			}
		default:
			goto st832
		}
		goto st831
	st842:
		if p++; p == pe {
			goto _test_eof842
		}
	st_case_842:
		if data[p] == 95 {
			goto st833
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st833
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st833
			}
		default:
			goto st833
		}
		goto st832
	st843:
		if p++; p == pe {
			goto _test_eof843
		}
	st_case_843:
		switch data[p] {
		case 32:
			goto st832
		case 40:
			goto st832
		case 46:
			goto st832
		case 95:
			goto st833
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st844
				}
			case data[p] >= 9:
				goto st832
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st833
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st833
				}
			default:
				goto st832
			}
		default:
			goto st832
		}
		goto st831
	st844:
		if p++; p == pe {
			goto _test_eof844
		}
	st_case_844:
		switch data[p] {
		case 32:
			goto tr889
		case 40:
			goto tr889
		case 46:
			goto tr889
		case 95:
			goto st833
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st844
				}
			case data[p] >= 9:
				goto tr889
			}
		case data[p] > 61:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st833
				}
			case data[p] > 92:
				if 97 <= data[p] && data[p] <= 122 {
					goto st833
				}
			default:
				goto tr889
			}
		default:
			goto tr889
		}
		goto tr888
	st845:
		if p++; p == pe {
			goto _test_eof845
		}
	st_case_845:
		switch data[p] {
		case 32:
			goto st832
		case 40:
			goto st832
		case 46:
			goto st832
		case 91:
			goto st832
		case 92:
			goto st834
		case 95:
			goto st833
		case 100:
			goto st835
		case 110:
			goto st839
		case 117:
			goto st843
		case 120:
			goto st843
		}
		switch {
		case data[p] < 60:
			switch {
			case data[p] > 13:
				if 48 <= data[p] && data[p] <= 57 {
					goto st833
				}
			case data[p] >= 9:
				goto st832
			}
		case data[p] > 61:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st833
				}
			case data[p] >= 65:
				goto st833
			}
		default:
			goto st832
		}
		goto st831
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
	_test_eof846: cs = 846; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof847: cs = 847; goto _test_eof
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
	_test_eof848: cs = 848; goto _test_eof
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
	_test_eof849: cs = 849; goto _test_eof
	_test_eof70: cs = 70; goto _test_eof
	_test_eof71: cs = 71; goto _test_eof
	_test_eof850: cs = 850; goto _test_eof
	_test_eof851: cs = 851; goto _test_eof
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
	_test_eof852: cs = 852; goto _test_eof
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
	_test_eof853: cs = 853; goto _test_eof
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
	_test_eof854: cs = 854; goto _test_eof
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
	_test_eof183: cs = 183; goto _test_eof
	_test_eof184: cs = 184; goto _test_eof
	_test_eof855: cs = 855; goto _test_eof
	_test_eof185: cs = 185; goto _test_eof
	_test_eof186: cs = 186; goto _test_eof
	_test_eof187: cs = 187; goto _test_eof
	_test_eof188: cs = 188; goto _test_eof
	_test_eof189: cs = 189; goto _test_eof
	_test_eof190: cs = 190; goto _test_eof
	_test_eof191: cs = 191; goto _test_eof
	_test_eof192: cs = 192; goto _test_eof
	_test_eof856: cs = 856; goto _test_eof
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
	_test_eof205: cs = 205; goto _test_eof
	_test_eof206: cs = 206; goto _test_eof
	_test_eof207: cs = 207; goto _test_eof
	_test_eof208: cs = 208; goto _test_eof
	_test_eof209: cs = 209; goto _test_eof
	_test_eof857: cs = 857; goto _test_eof
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
	_test_eof230: cs = 230; goto _test_eof
	_test_eof231: cs = 231; goto _test_eof
	_test_eof232: cs = 232; goto _test_eof
	_test_eof233: cs = 233; goto _test_eof
	_test_eof858: cs = 858; goto _test_eof
	_test_eof234: cs = 234; goto _test_eof
	_test_eof235: cs = 235; goto _test_eof
	_test_eof859: cs = 859; goto _test_eof
	_test_eof236: cs = 236; goto _test_eof
	_test_eof237: cs = 237; goto _test_eof
	_test_eof860: cs = 860; goto _test_eof
	_test_eof861: cs = 861; goto _test_eof
	_test_eof238: cs = 238; goto _test_eof
	_test_eof862: cs = 862; goto _test_eof
	_test_eof863: cs = 863; goto _test_eof
	_test_eof239: cs = 239; goto _test_eof
	_test_eof240: cs = 240; goto _test_eof
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
	_test_eof254: cs = 254; goto _test_eof
	_test_eof255: cs = 255; goto _test_eof
	_test_eof256: cs = 256; goto _test_eof
	_test_eof257: cs = 257; goto _test_eof
	_test_eof258: cs = 258; goto _test_eof
	_test_eof259: cs = 259; goto _test_eof
	_test_eof260: cs = 260; goto _test_eof
	_test_eof261: cs = 261; goto _test_eof
	_test_eof864: cs = 864; goto _test_eof
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
	_test_eof865: cs = 865; goto _test_eof
	_test_eof284: cs = 284; goto _test_eof
	_test_eof285: cs = 285; goto _test_eof
	_test_eof286: cs = 286; goto _test_eof
	_test_eof287: cs = 287; goto _test_eof
	_test_eof288: cs = 288; goto _test_eof
	_test_eof289: cs = 289; goto _test_eof
	_test_eof290: cs = 290; goto _test_eof
	_test_eof291: cs = 291; goto _test_eof
	_test_eof292: cs = 292; goto _test_eof
	_test_eof293: cs = 293; goto _test_eof
	_test_eof294: cs = 294; goto _test_eof
	_test_eof295: cs = 295; goto _test_eof
	_test_eof296: cs = 296; goto _test_eof
	_test_eof297: cs = 297; goto _test_eof
	_test_eof298: cs = 298; goto _test_eof
	_test_eof299: cs = 299; goto _test_eof
	_test_eof866: cs = 866; goto _test_eof
	_test_eof300: cs = 300; goto _test_eof
	_test_eof301: cs = 301; goto _test_eof
	_test_eof302: cs = 302; goto _test_eof
	_test_eof303: cs = 303; goto _test_eof
	_test_eof304: cs = 304; goto _test_eof
	_test_eof305: cs = 305; goto _test_eof
	_test_eof306: cs = 306; goto _test_eof
	_test_eof307: cs = 307; goto _test_eof
	_test_eof308: cs = 308; goto _test_eof
	_test_eof309: cs = 309; goto _test_eof
	_test_eof310: cs = 310; goto _test_eof
	_test_eof311: cs = 311; goto _test_eof
	_test_eof312: cs = 312; goto _test_eof
	_test_eof313: cs = 313; goto _test_eof
	_test_eof314: cs = 314; goto _test_eof
	_test_eof315: cs = 315; goto _test_eof
	_test_eof316: cs = 316; goto _test_eof
	_test_eof317: cs = 317; goto _test_eof
	_test_eof318: cs = 318; goto _test_eof
	_test_eof319: cs = 319; goto _test_eof
	_test_eof320: cs = 320; goto _test_eof
	_test_eof321: cs = 321; goto _test_eof
	_test_eof322: cs = 322; goto _test_eof
	_test_eof323: cs = 323; goto _test_eof
	_test_eof324: cs = 324; goto _test_eof
	_test_eof867: cs = 867; goto _test_eof
	_test_eof325: cs = 325; goto _test_eof
	_test_eof326: cs = 326; goto _test_eof
	_test_eof327: cs = 327; goto _test_eof
	_test_eof328: cs = 328; goto _test_eof
	_test_eof329: cs = 329; goto _test_eof
	_test_eof330: cs = 330; goto _test_eof
	_test_eof331: cs = 331; goto _test_eof
	_test_eof332: cs = 332; goto _test_eof
	_test_eof333: cs = 333; goto _test_eof
	_test_eof334: cs = 334; goto _test_eof
	_test_eof335: cs = 335; goto _test_eof
	_test_eof336: cs = 336; goto _test_eof
	_test_eof337: cs = 337; goto _test_eof
	_test_eof338: cs = 338; goto _test_eof
	_test_eof339: cs = 339; goto _test_eof
	_test_eof340: cs = 340; goto _test_eof
	_test_eof341: cs = 341; goto _test_eof
	_test_eof342: cs = 342; goto _test_eof
	_test_eof343: cs = 343; goto _test_eof
	_test_eof344: cs = 344; goto _test_eof
	_test_eof345: cs = 345; goto _test_eof
	_test_eof346: cs = 346; goto _test_eof
	_test_eof347: cs = 347; goto _test_eof
	_test_eof348: cs = 348; goto _test_eof
	_test_eof349: cs = 349; goto _test_eof
	_test_eof868: cs = 868; goto _test_eof
	_test_eof350: cs = 350; goto _test_eof
	_test_eof351: cs = 351; goto _test_eof
	_test_eof352: cs = 352; goto _test_eof
	_test_eof353: cs = 353; goto _test_eof
	_test_eof354: cs = 354; goto _test_eof
	_test_eof355: cs = 355; goto _test_eof
	_test_eof356: cs = 356; goto _test_eof
	_test_eof357: cs = 357; goto _test_eof
	_test_eof358: cs = 358; goto _test_eof
	_test_eof359: cs = 359; goto _test_eof
	_test_eof360: cs = 360; goto _test_eof
	_test_eof361: cs = 361; goto _test_eof
	_test_eof362: cs = 362; goto _test_eof
	_test_eof363: cs = 363; goto _test_eof
	_test_eof364: cs = 364; goto _test_eof
	_test_eof365: cs = 365; goto _test_eof
	_test_eof366: cs = 366; goto _test_eof
	_test_eof367: cs = 367; goto _test_eof
	_test_eof368: cs = 368; goto _test_eof
	_test_eof369: cs = 369; goto _test_eof
	_test_eof370: cs = 370; goto _test_eof
	_test_eof371: cs = 371; goto _test_eof
	_test_eof372: cs = 372; goto _test_eof
	_test_eof373: cs = 373; goto _test_eof
	_test_eof374: cs = 374; goto _test_eof
	_test_eof375: cs = 375; goto _test_eof
	_test_eof376: cs = 376; goto _test_eof
	_test_eof377: cs = 377; goto _test_eof
	_test_eof869: cs = 869; goto _test_eof
	_test_eof378: cs = 378; goto _test_eof
	_test_eof379: cs = 379; goto _test_eof
	_test_eof870: cs = 870; goto _test_eof
	_test_eof871: cs = 871; goto _test_eof
	_test_eof380: cs = 380; goto _test_eof
	_test_eof381: cs = 381; goto _test_eof
	_test_eof382: cs = 382; goto _test_eof
	_test_eof383: cs = 383; goto _test_eof
	_test_eof384: cs = 384; goto _test_eof
	_test_eof385: cs = 385; goto _test_eof
	_test_eof386: cs = 386; goto _test_eof
	_test_eof387: cs = 387; goto _test_eof
	_test_eof388: cs = 388; goto _test_eof
	_test_eof389: cs = 389; goto _test_eof
	_test_eof390: cs = 390; goto _test_eof
	_test_eof391: cs = 391; goto _test_eof
	_test_eof392: cs = 392; goto _test_eof
	_test_eof393: cs = 393; goto _test_eof
	_test_eof394: cs = 394; goto _test_eof
	_test_eof395: cs = 395; goto _test_eof
	_test_eof396: cs = 396; goto _test_eof
	_test_eof397: cs = 397; goto _test_eof
	_test_eof398: cs = 398; goto _test_eof
	_test_eof399: cs = 399; goto _test_eof
	_test_eof872: cs = 872; goto _test_eof
	_test_eof400: cs = 400; goto _test_eof
	_test_eof401: cs = 401; goto _test_eof
	_test_eof402: cs = 402; goto _test_eof
	_test_eof403: cs = 403; goto _test_eof
	_test_eof404: cs = 404; goto _test_eof
	_test_eof405: cs = 405; goto _test_eof
	_test_eof406: cs = 406; goto _test_eof
	_test_eof407: cs = 407; goto _test_eof
	_test_eof408: cs = 408; goto _test_eof
	_test_eof409: cs = 409; goto _test_eof
	_test_eof410: cs = 410; goto _test_eof
	_test_eof411: cs = 411; goto _test_eof
	_test_eof412: cs = 412; goto _test_eof
	_test_eof413: cs = 413; goto _test_eof
	_test_eof414: cs = 414; goto _test_eof
	_test_eof415: cs = 415; goto _test_eof
	_test_eof416: cs = 416; goto _test_eof
	_test_eof417: cs = 417; goto _test_eof
	_test_eof418: cs = 418; goto _test_eof
	_test_eof419: cs = 419; goto _test_eof
	_test_eof420: cs = 420; goto _test_eof
	_test_eof421: cs = 421; goto _test_eof
	_test_eof422: cs = 422; goto _test_eof
	_test_eof423: cs = 423; goto _test_eof
	_test_eof424: cs = 424; goto _test_eof
	_test_eof425: cs = 425; goto _test_eof
	_test_eof426: cs = 426; goto _test_eof
	_test_eof427: cs = 427; goto _test_eof
	_test_eof428: cs = 428; goto _test_eof
	_test_eof429: cs = 429; goto _test_eof
	_test_eof430: cs = 430; goto _test_eof
	_test_eof431: cs = 431; goto _test_eof
	_test_eof432: cs = 432; goto _test_eof
	_test_eof433: cs = 433; goto _test_eof
	_test_eof434: cs = 434; goto _test_eof
	_test_eof435: cs = 435; goto _test_eof
	_test_eof436: cs = 436; goto _test_eof
	_test_eof437: cs = 437; goto _test_eof
	_test_eof438: cs = 438; goto _test_eof
	_test_eof439: cs = 439; goto _test_eof
	_test_eof440: cs = 440; goto _test_eof
	_test_eof441: cs = 441; goto _test_eof
	_test_eof442: cs = 442; goto _test_eof
	_test_eof443: cs = 443; goto _test_eof
	_test_eof444: cs = 444; goto _test_eof
	_test_eof445: cs = 445; goto _test_eof
	_test_eof446: cs = 446; goto _test_eof
	_test_eof447: cs = 447; goto _test_eof
	_test_eof448: cs = 448; goto _test_eof
	_test_eof449: cs = 449; goto _test_eof
	_test_eof873: cs = 873; goto _test_eof
	_test_eof450: cs = 450; goto _test_eof
	_test_eof451: cs = 451; goto _test_eof
	_test_eof452: cs = 452; goto _test_eof
	_test_eof453: cs = 453; goto _test_eof
	_test_eof454: cs = 454; goto _test_eof
	_test_eof455: cs = 455; goto _test_eof
	_test_eof874: cs = 874; goto _test_eof
	_test_eof456: cs = 456; goto _test_eof
	_test_eof457: cs = 457; goto _test_eof
	_test_eof458: cs = 458; goto _test_eof
	_test_eof459: cs = 459; goto _test_eof
	_test_eof460: cs = 460; goto _test_eof
	_test_eof461: cs = 461; goto _test_eof
	_test_eof462: cs = 462; goto _test_eof
	_test_eof463: cs = 463; goto _test_eof
	_test_eof464: cs = 464; goto _test_eof
	_test_eof465: cs = 465; goto _test_eof
	_test_eof466: cs = 466; goto _test_eof
	_test_eof467: cs = 467; goto _test_eof
	_test_eof468: cs = 468; goto _test_eof
	_test_eof875: cs = 875; goto _test_eof
	_test_eof469: cs = 469; goto _test_eof
	_test_eof470: cs = 470; goto _test_eof
	_test_eof876: cs = 876; goto _test_eof
	_test_eof471: cs = 471; goto _test_eof
	_test_eof472: cs = 472; goto _test_eof
	_test_eof877: cs = 877; goto _test_eof
	_test_eof878: cs = 878; goto _test_eof
	_test_eof473: cs = 473; goto _test_eof
	_test_eof879: cs = 879; goto _test_eof
	_test_eof880: cs = 880; goto _test_eof
	_test_eof474: cs = 474; goto _test_eof
	_test_eof475: cs = 475; goto _test_eof
	_test_eof476: cs = 476; goto _test_eof
	_test_eof477: cs = 477; goto _test_eof
	_test_eof478: cs = 478; goto _test_eof
	_test_eof479: cs = 479; goto _test_eof
	_test_eof480: cs = 480; goto _test_eof
	_test_eof481: cs = 481; goto _test_eof
	_test_eof482: cs = 482; goto _test_eof
	_test_eof483: cs = 483; goto _test_eof
	_test_eof484: cs = 484; goto _test_eof
	_test_eof485: cs = 485; goto _test_eof
	_test_eof486: cs = 486; goto _test_eof
	_test_eof487: cs = 487; goto _test_eof
	_test_eof488: cs = 488; goto _test_eof
	_test_eof489: cs = 489; goto _test_eof
	_test_eof490: cs = 490; goto _test_eof
	_test_eof491: cs = 491; goto _test_eof
	_test_eof492: cs = 492; goto _test_eof
	_test_eof493: cs = 493; goto _test_eof
	_test_eof881: cs = 881; goto _test_eof
	_test_eof494: cs = 494; goto _test_eof
	_test_eof495: cs = 495; goto _test_eof
	_test_eof496: cs = 496; goto _test_eof
	_test_eof497: cs = 497; goto _test_eof
	_test_eof498: cs = 498; goto _test_eof
	_test_eof499: cs = 499; goto _test_eof
	_test_eof500: cs = 500; goto _test_eof
	_test_eof501: cs = 501; goto _test_eof
	_test_eof502: cs = 502; goto _test_eof
	_test_eof503: cs = 503; goto _test_eof
	_test_eof504: cs = 504; goto _test_eof
	_test_eof505: cs = 505; goto _test_eof
	_test_eof506: cs = 506; goto _test_eof
	_test_eof507: cs = 507; goto _test_eof
	_test_eof508: cs = 508; goto _test_eof
	_test_eof509: cs = 509; goto _test_eof
	_test_eof510: cs = 510; goto _test_eof
	_test_eof511: cs = 511; goto _test_eof
	_test_eof512: cs = 512; goto _test_eof
	_test_eof513: cs = 513; goto _test_eof
	_test_eof514: cs = 514; goto _test_eof
	_test_eof882: cs = 882; goto _test_eof
	_test_eof515: cs = 515; goto _test_eof
	_test_eof516: cs = 516; goto _test_eof
	_test_eof517: cs = 517; goto _test_eof
	_test_eof518: cs = 518; goto _test_eof
	_test_eof519: cs = 519; goto _test_eof
	_test_eof520: cs = 520; goto _test_eof
	_test_eof521: cs = 521; goto _test_eof
	_test_eof522: cs = 522; goto _test_eof
	_test_eof523: cs = 523; goto _test_eof
	_test_eof524: cs = 524; goto _test_eof
	_test_eof525: cs = 525; goto _test_eof
	_test_eof526: cs = 526; goto _test_eof
	_test_eof527: cs = 527; goto _test_eof
	_test_eof528: cs = 528; goto _test_eof
	_test_eof529: cs = 529; goto _test_eof
	_test_eof530: cs = 530; goto _test_eof
	_test_eof531: cs = 531; goto _test_eof
	_test_eof532: cs = 532; goto _test_eof
	_test_eof533: cs = 533; goto _test_eof
	_test_eof534: cs = 534; goto _test_eof
	_test_eof535: cs = 535; goto _test_eof
	_test_eof536: cs = 536; goto _test_eof
	_test_eof537: cs = 537; goto _test_eof
	_test_eof538: cs = 538; goto _test_eof
	_test_eof539: cs = 539; goto _test_eof
	_test_eof540: cs = 540; goto _test_eof
	_test_eof541: cs = 541; goto _test_eof
	_test_eof542: cs = 542; goto _test_eof
	_test_eof543: cs = 543; goto _test_eof
	_test_eof544: cs = 544; goto _test_eof
	_test_eof545: cs = 545; goto _test_eof
	_test_eof546: cs = 546; goto _test_eof
	_test_eof547: cs = 547; goto _test_eof
	_test_eof548: cs = 548; goto _test_eof
	_test_eof549: cs = 549; goto _test_eof
	_test_eof550: cs = 550; goto _test_eof
	_test_eof551: cs = 551; goto _test_eof
	_test_eof552: cs = 552; goto _test_eof
	_test_eof553: cs = 553; goto _test_eof
	_test_eof554: cs = 554; goto _test_eof
	_test_eof555: cs = 555; goto _test_eof
	_test_eof556: cs = 556; goto _test_eof
	_test_eof557: cs = 557; goto _test_eof
	_test_eof558: cs = 558; goto _test_eof
	_test_eof559: cs = 559; goto _test_eof
	_test_eof560: cs = 560; goto _test_eof
	_test_eof561: cs = 561; goto _test_eof
	_test_eof562: cs = 562; goto _test_eof
	_test_eof883: cs = 883; goto _test_eof
	_test_eof563: cs = 563; goto _test_eof
	_test_eof564: cs = 564; goto _test_eof
	_test_eof565: cs = 565; goto _test_eof
	_test_eof566: cs = 566; goto _test_eof
	_test_eof567: cs = 567; goto _test_eof
	_test_eof568: cs = 568; goto _test_eof
	_test_eof569: cs = 569; goto _test_eof
	_test_eof570: cs = 570; goto _test_eof
	_test_eof571: cs = 571; goto _test_eof
	_test_eof572: cs = 572; goto _test_eof
	_test_eof573: cs = 573; goto _test_eof
	_test_eof574: cs = 574; goto _test_eof
	_test_eof575: cs = 575; goto _test_eof
	_test_eof576: cs = 576; goto _test_eof
	_test_eof577: cs = 577; goto _test_eof
	_test_eof578: cs = 578; goto _test_eof
	_test_eof579: cs = 579; goto _test_eof
	_test_eof580: cs = 580; goto _test_eof
	_test_eof581: cs = 581; goto _test_eof
	_test_eof582: cs = 582; goto _test_eof
	_test_eof583: cs = 583; goto _test_eof
	_test_eof584: cs = 584; goto _test_eof
	_test_eof585: cs = 585; goto _test_eof
	_test_eof586: cs = 586; goto _test_eof
	_test_eof587: cs = 587; goto _test_eof
	_test_eof884: cs = 884; goto _test_eof
	_test_eof588: cs = 588; goto _test_eof
	_test_eof589: cs = 589; goto _test_eof
	_test_eof590: cs = 590; goto _test_eof
	_test_eof591: cs = 591; goto _test_eof
	_test_eof592: cs = 592; goto _test_eof
	_test_eof593: cs = 593; goto _test_eof
	_test_eof594: cs = 594; goto _test_eof
	_test_eof595: cs = 595; goto _test_eof
	_test_eof596: cs = 596; goto _test_eof
	_test_eof597: cs = 597; goto _test_eof
	_test_eof598: cs = 598; goto _test_eof
	_test_eof599: cs = 599; goto _test_eof
	_test_eof600: cs = 600; goto _test_eof
	_test_eof601: cs = 601; goto _test_eof
	_test_eof602: cs = 602; goto _test_eof
	_test_eof603: cs = 603; goto _test_eof
	_test_eof604: cs = 604; goto _test_eof
	_test_eof605: cs = 605; goto _test_eof
	_test_eof606: cs = 606; goto _test_eof
	_test_eof607: cs = 607; goto _test_eof
	_test_eof608: cs = 608; goto _test_eof
	_test_eof609: cs = 609; goto _test_eof
	_test_eof610: cs = 610; goto _test_eof
	_test_eof611: cs = 611; goto _test_eof
	_test_eof885: cs = 885; goto _test_eof
	_test_eof612: cs = 612; goto _test_eof
	_test_eof613: cs = 613; goto _test_eof
	_test_eof886: cs = 886; goto _test_eof
	_test_eof887: cs = 887; goto _test_eof
	_test_eof614: cs = 614; goto _test_eof
	_test_eof615: cs = 615; goto _test_eof
	_test_eof616: cs = 616; goto _test_eof
	_test_eof617: cs = 617; goto _test_eof
	_test_eof618: cs = 618; goto _test_eof
	_test_eof619: cs = 619; goto _test_eof
	_test_eof620: cs = 620; goto _test_eof
	_test_eof621: cs = 621; goto _test_eof
	_test_eof622: cs = 622; goto _test_eof
	_test_eof623: cs = 623; goto _test_eof
	_test_eof624: cs = 624; goto _test_eof
	_test_eof625: cs = 625; goto _test_eof
	_test_eof626: cs = 626; goto _test_eof
	_test_eof627: cs = 627; goto _test_eof
	_test_eof628: cs = 628; goto _test_eof
	_test_eof629: cs = 629; goto _test_eof
	_test_eof630: cs = 630; goto _test_eof
	_test_eof631: cs = 631; goto _test_eof
	_test_eof632: cs = 632; goto _test_eof
	_test_eof633: cs = 633; goto _test_eof
	_test_eof634: cs = 634; goto _test_eof
	_test_eof635: cs = 635; goto _test_eof
	_test_eof636: cs = 636; goto _test_eof
	_test_eof888: cs = 888; goto _test_eof
	_test_eof637: cs = 637; goto _test_eof
	_test_eof638: cs = 638; goto _test_eof
	_test_eof639: cs = 639; goto _test_eof
	_test_eof640: cs = 640; goto _test_eof
	_test_eof641: cs = 641; goto _test_eof
	_test_eof642: cs = 642; goto _test_eof
	_test_eof643: cs = 643; goto _test_eof
	_test_eof644: cs = 644; goto _test_eof
	_test_eof645: cs = 645; goto _test_eof
	_test_eof646: cs = 646; goto _test_eof
	_test_eof647: cs = 647; goto _test_eof
	_test_eof648: cs = 648; goto _test_eof
	_test_eof649: cs = 649; goto _test_eof
	_test_eof650: cs = 650; goto _test_eof
	_test_eof651: cs = 651; goto _test_eof
	_test_eof652: cs = 652; goto _test_eof
	_test_eof653: cs = 653; goto _test_eof
	_test_eof654: cs = 654; goto _test_eof
	_test_eof655: cs = 655; goto _test_eof
	_test_eof656: cs = 656; goto _test_eof
	_test_eof657: cs = 657; goto _test_eof
	_test_eof658: cs = 658; goto _test_eof
	_test_eof889: cs = 889; goto _test_eof
	_test_eof659: cs = 659; goto _test_eof
	_test_eof660: cs = 660; goto _test_eof
	_test_eof661: cs = 661; goto _test_eof
	_test_eof662: cs = 662; goto _test_eof
	_test_eof663: cs = 663; goto _test_eof
	_test_eof664: cs = 664; goto _test_eof
	_test_eof665: cs = 665; goto _test_eof
	_test_eof666: cs = 666; goto _test_eof
	_test_eof667: cs = 667; goto _test_eof
	_test_eof668: cs = 668; goto _test_eof
	_test_eof669: cs = 669; goto _test_eof
	_test_eof670: cs = 670; goto _test_eof
	_test_eof671: cs = 671; goto _test_eof
	_test_eof672: cs = 672; goto _test_eof
	_test_eof673: cs = 673; goto _test_eof
	_test_eof674: cs = 674; goto _test_eof
	_test_eof890: cs = 890; goto _test_eof
	_test_eof675: cs = 675; goto _test_eof
	_test_eof676: cs = 676; goto _test_eof
	_test_eof677: cs = 677; goto _test_eof
	_test_eof678: cs = 678; goto _test_eof
	_test_eof679: cs = 679; goto _test_eof
	_test_eof680: cs = 680; goto _test_eof
	_test_eof681: cs = 681; goto _test_eof
	_test_eof682: cs = 682; goto _test_eof
	_test_eof683: cs = 683; goto _test_eof
	_test_eof684: cs = 684; goto _test_eof
	_test_eof685: cs = 685; goto _test_eof
	_test_eof686: cs = 686; goto _test_eof
	_test_eof687: cs = 687; goto _test_eof
	_test_eof688: cs = 688; goto _test_eof
	_test_eof689: cs = 689; goto _test_eof
	_test_eof690: cs = 690; goto _test_eof
	_test_eof691: cs = 691; goto _test_eof
	_test_eof692: cs = 692; goto _test_eof
	_test_eof693: cs = 693; goto _test_eof
	_test_eof694: cs = 694; goto _test_eof
	_test_eof695: cs = 695; goto _test_eof
	_test_eof696: cs = 696; goto _test_eof
	_test_eof697: cs = 697; goto _test_eof
	_test_eof698: cs = 698; goto _test_eof
	_test_eof699: cs = 699; goto _test_eof
	_test_eof891: cs = 891; goto _test_eof
	_test_eof700: cs = 700; goto _test_eof
	_test_eof701: cs = 701; goto _test_eof
	_test_eof702: cs = 702; goto _test_eof
	_test_eof703: cs = 703; goto _test_eof
	_test_eof704: cs = 704; goto _test_eof
	_test_eof705: cs = 705; goto _test_eof
	_test_eof706: cs = 706; goto _test_eof
	_test_eof707: cs = 707; goto _test_eof
	_test_eof708: cs = 708; goto _test_eof
	_test_eof709: cs = 709; goto _test_eof
	_test_eof710: cs = 710; goto _test_eof
	_test_eof711: cs = 711; goto _test_eof
	_test_eof712: cs = 712; goto _test_eof
	_test_eof713: cs = 713; goto _test_eof
	_test_eof714: cs = 714; goto _test_eof
	_test_eof715: cs = 715; goto _test_eof
	_test_eof716: cs = 716; goto _test_eof
	_test_eof717: cs = 717; goto _test_eof
	_test_eof718: cs = 718; goto _test_eof
	_test_eof719: cs = 719; goto _test_eof
	_test_eof720: cs = 720; goto _test_eof
	_test_eof721: cs = 721; goto _test_eof
	_test_eof722: cs = 722; goto _test_eof
	_test_eof723: cs = 723; goto _test_eof
	_test_eof724: cs = 724; goto _test_eof
	_test_eof725: cs = 725; goto _test_eof
	_test_eof726: cs = 726; goto _test_eof
	_test_eof727: cs = 727; goto _test_eof
	_test_eof728: cs = 728; goto _test_eof
	_test_eof729: cs = 729; goto _test_eof
	_test_eof730: cs = 730; goto _test_eof
	_test_eof731: cs = 731; goto _test_eof
	_test_eof732: cs = 732; goto _test_eof
	_test_eof733: cs = 733; goto _test_eof
	_test_eof734: cs = 734; goto _test_eof
	_test_eof735: cs = 735; goto _test_eof
	_test_eof736: cs = 736; goto _test_eof
	_test_eof737: cs = 737; goto _test_eof
	_test_eof892: cs = 892; goto _test_eof
	_test_eof738: cs = 738; goto _test_eof
	_test_eof739: cs = 739; goto _test_eof
	_test_eof893: cs = 893; goto _test_eof
	_test_eof894: cs = 894; goto _test_eof
	_test_eof740: cs = 740; goto _test_eof
	_test_eof741: cs = 741; goto _test_eof
	_test_eof742: cs = 742; goto _test_eof
	_test_eof743: cs = 743; goto _test_eof
	_test_eof744: cs = 744; goto _test_eof
	_test_eof745: cs = 745; goto _test_eof
	_test_eof746: cs = 746; goto _test_eof
	_test_eof747: cs = 747; goto _test_eof
	_test_eof748: cs = 748; goto _test_eof
	_test_eof749: cs = 749; goto _test_eof
	_test_eof750: cs = 750; goto _test_eof
	_test_eof751: cs = 751; goto _test_eof
	_test_eof752: cs = 752; goto _test_eof
	_test_eof753: cs = 753; goto _test_eof
	_test_eof754: cs = 754; goto _test_eof
	_test_eof755: cs = 755; goto _test_eof
	_test_eof756: cs = 756; goto _test_eof
	_test_eof757: cs = 757; goto _test_eof
	_test_eof758: cs = 758; goto _test_eof
	_test_eof759: cs = 759; goto _test_eof
	_test_eof895: cs = 895; goto _test_eof
	_test_eof760: cs = 760; goto _test_eof
	_test_eof761: cs = 761; goto _test_eof
	_test_eof762: cs = 762; goto _test_eof
	_test_eof763: cs = 763; goto _test_eof
	_test_eof764: cs = 764; goto _test_eof
	_test_eof765: cs = 765; goto _test_eof
	_test_eof766: cs = 766; goto _test_eof
	_test_eof767: cs = 767; goto _test_eof
	_test_eof768: cs = 768; goto _test_eof
	_test_eof769: cs = 769; goto _test_eof
	_test_eof770: cs = 770; goto _test_eof
	_test_eof771: cs = 771; goto _test_eof
	_test_eof772: cs = 772; goto _test_eof
	_test_eof773: cs = 773; goto _test_eof
	_test_eof774: cs = 774; goto _test_eof
	_test_eof775: cs = 775; goto _test_eof
	_test_eof776: cs = 776; goto _test_eof
	_test_eof777: cs = 777; goto _test_eof
	_test_eof778: cs = 778; goto _test_eof
	_test_eof779: cs = 779; goto _test_eof
	_test_eof780: cs = 780; goto _test_eof
	_test_eof781: cs = 781; goto _test_eof
	_test_eof782: cs = 782; goto _test_eof
	_test_eof783: cs = 783; goto _test_eof
	_test_eof784: cs = 784; goto _test_eof
	_test_eof785: cs = 785; goto _test_eof
	_test_eof786: cs = 786; goto _test_eof
	_test_eof787: cs = 787; goto _test_eof
	_test_eof788: cs = 788; goto _test_eof
	_test_eof789: cs = 789; goto _test_eof
	_test_eof790: cs = 790; goto _test_eof
	_test_eof791: cs = 791; goto _test_eof
	_test_eof792: cs = 792; goto _test_eof
	_test_eof793: cs = 793; goto _test_eof
	_test_eof794: cs = 794; goto _test_eof
	_test_eof795: cs = 795; goto _test_eof
	_test_eof796: cs = 796; goto _test_eof
	_test_eof797: cs = 797; goto _test_eof
	_test_eof798: cs = 798; goto _test_eof
	_test_eof799: cs = 799; goto _test_eof
	_test_eof800: cs = 800; goto _test_eof
	_test_eof801: cs = 801; goto _test_eof
	_test_eof802: cs = 802; goto _test_eof
	_test_eof803: cs = 803; goto _test_eof
	_test_eof804: cs = 804; goto _test_eof
	_test_eof805: cs = 805; goto _test_eof
	_test_eof806: cs = 806; goto _test_eof
	_test_eof807: cs = 807; goto _test_eof
	_test_eof808: cs = 808; goto _test_eof
	_test_eof809: cs = 809; goto _test_eof
	_test_eof810: cs = 810; goto _test_eof
	_test_eof811: cs = 811; goto _test_eof
	_test_eof812: cs = 812; goto _test_eof
	_test_eof813: cs = 813; goto _test_eof
	_test_eof814: cs = 814; goto _test_eof
	_test_eof815: cs = 815; goto _test_eof
	_test_eof816: cs = 816; goto _test_eof
	_test_eof817: cs = 817; goto _test_eof
	_test_eof818: cs = 818; goto _test_eof
	_test_eof819: cs = 819; goto _test_eof
	_test_eof820: cs = 820; goto _test_eof
	_test_eof821: cs = 821; goto _test_eof
	_test_eof822: cs = 822; goto _test_eof
	_test_eof823: cs = 823; goto _test_eof
	_test_eof824: cs = 824; goto _test_eof
	_test_eof825: cs = 825; goto _test_eof
	_test_eof826: cs = 826; goto _test_eof
	_test_eof827: cs = 827; goto _test_eof
	_test_eof828: cs = 828; goto _test_eof
	_test_eof829: cs = 829; goto _test_eof
	_test_eof830: cs = 830; goto _test_eof
	_test_eof831: cs = 831; goto _test_eof
	_test_eof832: cs = 832; goto _test_eof
	_test_eof896: cs = 896; goto _test_eof
	_test_eof833: cs = 833; goto _test_eof
	_test_eof834: cs = 834; goto _test_eof
	_test_eof897: cs = 897; goto _test_eof
	_test_eof898: cs = 898; goto _test_eof
	_test_eof835: cs = 835; goto _test_eof
	_test_eof836: cs = 836; goto _test_eof
	_test_eof837: cs = 837; goto _test_eof
	_test_eof838: cs = 838; goto _test_eof
	_test_eof839: cs = 839; goto _test_eof
	_test_eof840: cs = 840; goto _test_eof
	_test_eof841: cs = 841; goto _test_eof
	_test_eof842: cs = 842; goto _test_eof
	_test_eof843: cs = 843; goto _test_eof
	_test_eof844: cs = 844; goto _test_eof
	_test_eof845: cs = 845; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 15, 21, 45, 69, 71, 77, 94, 103, 116, 125, 131, 147, 184, 192, 209, 233, 235, 237, 238, 244, 261, 270, 283, 292, 298, 324, 349, 377, 379, 399, 413, 419, 455, 468, 470, 472, 473, 493, 507, 513, 562, 570, 587, 611, 613, 619, 636, 645, 658, 667, 673, 699, 737, 739, 759, 773, 779, 832, 834, 844, 847, 850, 851, 860, 861, 862, 863, 866, 870, 871, 873, 877, 878, 879, 880, 886, 887, 890, 893, 894, 897, 898:
//line xss_170.rl:13

            return true
        
//line xss_170.go:36552
		}
	}

	}

//line xss_170.rl:39


    return false
}
