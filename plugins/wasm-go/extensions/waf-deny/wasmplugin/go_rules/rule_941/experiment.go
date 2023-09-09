
//line experiment.rl:1
package rule_941

import (
    "fmt"
    "strings"
    "golang.org/x/exp/slices"
)

type Machine struct {

}

type XssChecker struct {

}

var riskTags = []string {
    "script", "style", "svg", "set", "form", "marquee", "meta", "link", "object", "embed", "applet", "audio", "animate",
    "param", "iframe", "frame", "base", "body", "bindings", "image", "img", "video",
}



func matchExp(data []byte) bool {

//line experiment.rl:26

//line experiment.go:31
const xss_start int = 0
const xss_first_final int = 77
const xss_error int = -1

const xss_en_main int = 0


//line experiment.rl:27
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof
    pb := 0
    maybeTag := false
    var tagNameBuilder strings.Builder
    var tagName string

    maybeAttr := false
    var attrName string

    
//line experiment.go:51
	{
	cs = xss_start
	}

//line experiment.go:56
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
	case 77:
		goto st_case_77
	case 10:
		goto st_case_10
	case 78:
		goto st_case_78
	case 79:
		goto st_case_79
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 80:
		goto st_case_80
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
	case 81:
		goto st_case_81
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
	case 82:
		goto st_case_82
	case 83:
		goto st_case_83
	}
	goto st_out
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		switch data[p] {
		case 34:
			goto st1
		case 39:
			goto st1
		case 60:
			goto st72
		}
		goto st0
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch data[p] {
		case 32:
			goto st1
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr5
			}
		default:
			goto tr5
		}
		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 32:
			goto st1
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st1
		}
		goto st2
tr9:
//line experiment.rl:45

            maybeTag = true
            tagNameBuilder.Reset()
        
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line experiment.go:309
		switch data[p] {
		case 32:
			goto tr7
		case 34:
			goto tr7
		case 39:
			goto tr7
		case 47:
			goto tr7
		case 60:
			goto tr9
		case 62:
			goto tr10
		case 95:
			goto tr8
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr7
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr11
				}
			case data[p] >= 65:
				goto tr11
			}
		default:
			goto tr8
		}
		goto tr6
tr6:
//line experiment.rl:45

            maybeTag = true
            tagNameBuilder.Reset()
        
	goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line experiment.go:356
		switch data[p] {
		case 32:
			goto st5
		case 34:
			goto st5
		case 39:
			goto st5
		case 47:
			goto st5
		case 60:
			goto st3
		case 62:
			goto st7
		case 95:
			goto st6
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
					goto tr16
				}
			case data[p] >= 65:
				goto tr16
			}
		default:
			goto st6
		}
		goto st4
tr7:
//line experiment.rl:45

            maybeTag = true
            tagNameBuilder.Reset()
        
	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line experiment.go:403
		switch data[p] {
		case 32:
			goto st5
		case 34:
			goto st5
		case 39:
			goto st5
		case 47:
			goto st5
		case 60:
			goto st3
		case 62:
			goto st7
		case 95:
			goto st6
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
					goto tr17
				}
			case data[p] >= 65:
				goto tr17
			}
		default:
			goto st6
		}
		goto st4
tr8:
//line experiment.rl:45

            maybeTag = true
            tagNameBuilder.Reset()
        
	goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line experiment.go:450
		switch data[p] {
		case 32:
			goto st1
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 58:
			goto st7
		case 60:
			goto st3
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st1
		}
		goto st2
tr10:
//line experiment.rl:45

            maybeTag = true
            tagNameBuilder.Reset()
        
	goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line experiment.go:481
		switch data[p] {
		case 32:
			goto st8
		case 34:
			goto st8
		case 39:
			goto st8
		case 47:
			goto st8
		case 60:
			goto st3
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st8
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr16
				}
			case data[p] >= 65:
				goto tr16
			}
		default:
			goto st2
		}
		goto st7
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		switch data[p] {
		case 32:
			goto st8
		case 34:
			goto st8
		case 39:
			goto st8
		case 47:
			goto st8
		case 60:
			goto st3
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st8
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr17
				}
			case data[p] >= 65:
				goto tr17
			}
		default:
			goto st2
		}
		goto st7
tr17:
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
//line experiment.rl:71

            maybeAttr = true
            attrName = ""
            pb = p
        
	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line experiment.go:570
		switch data[p] {
		case 32:
			goto tr20
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr20
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr22
				}
			case data[p] >= 65:
				goto tr22
			}
		default:
			goto st2
		}
		goto tr19
tr19:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:56

            if maybeTag {
                tagName = tagNameBuilder.String()
                fmt.Println(tagName)

                if slices.Contains(riskTags, tagName) {
                    fmt.Println("Matched")
                    return true
                }

                tagNameBuilder.Reset()
                maybeTag = false
            }
        
	goto st77
tr26:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:56

            if maybeTag {
                tagName = tagNameBuilder.String()
                fmt.Println(tagName)

                if slices.Contains(riskTags, tagName) {
                    fmt.Println("Matched")
                    return true
                }

                tagNameBuilder.Reset()
                maybeTag = false
            }
        
//line experiment.rl:77

            if maybeAttr {
                attrName = string(data[pb:p])
                fmt.Printf("Attr name: %d %d %s \n", pb, p, attrName)
                maybeAttr = false
            }
        
//line experiment.rl:85

            if len(attrName) > 0 {
                return true
            }
        
	goto st77
tr29:
//line experiment.rl:85

            if len(attrName) > 0 {
                return true
            }
        
	goto st77
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
//line experiment.go:676
		switch data[p] {
		case 32:
			goto st8
		case 34:
			goto st8
		case 39:
			goto st8
		case 47:
			goto st8
		case 60:
			goto st3
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st8
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr16
				}
			case data[p] >= 65:
				goto tr16
			}
		default:
			goto st2
		}
		goto st7
tr11:
//line experiment.rl:45

            maybeTag = true
            tagNameBuilder.Reset()
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st10
tr16:
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st10
tr23:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line experiment.go:749
		switch data[p] {
		case 32:
			goto tr20
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr20
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr23
				}
			case data[p] >= 65:
				goto tr23
			}
		default:
			goto st2
		}
		goto tr19
tr20:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:56

            if maybeTag {
                tagName = tagNameBuilder.String()
                fmt.Println(tagName)

                if slices.Contains(riskTags, tagName) {
                    fmt.Println("Matched")
                    return true
                }

                tagNameBuilder.Reset()
                maybeTag = false
            }
        
	goto st78
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
//line experiment.go:810
		switch data[p] {
		case 32:
			goto st8
		case 34:
			goto st8
		case 39:
			goto st8
		case 47:
			goto st8
		case 60:
			goto st3
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st8
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr17
				}
			case data[p] >= 65:
				goto tr17
			}
		default:
			goto st2
		}
		goto st7
tr21:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:56

            if maybeTag {
                tagName = tagNameBuilder.String()
                fmt.Println(tagName)

                if slices.Contains(riskTags, tagName) {
                    fmt.Println("Matched")
                    return true
                }

                tagNameBuilder.Reset()
                maybeTag = false
            }
        
	goto st79
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
//line experiment.go:871
		switch data[p] {
		case 32:
			goto tr7
		case 34:
			goto tr7
		case 39:
			goto tr7
		case 47:
			goto tr7
		case 60:
			goto tr9
		case 62:
			goto tr10
		case 95:
			goto tr8
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr7
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr11
				}
			case data[p] >= 65:
				goto tr11
			}
		default:
			goto tr8
		}
		goto tr6
tr22:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line experiment.go:925
		switch data[p] {
		case 32:
			goto tr20
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr20
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr24
				}
			case data[p] >= 65:
				goto tr24
			}
		default:
			goto st2
		}
		goto tr19
tr24:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line experiment.go:977
		switch data[p] {
		case 32:
			goto tr25
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 61:
			goto tr26
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr27
				}
			case data[p] >= 65:
				goto tr27
			}
		default:
			goto st2
		}
		goto tr19
tr25:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:56

            if maybeTag {
                tagName = tagNameBuilder.String()
                fmt.Println(tagName)

                if slices.Contains(riskTags, tagName) {
                    fmt.Println("Matched")
                    return true
                }

                tagNameBuilder.Reset()
                maybeTag = false
            }
        
//line experiment.rl:77

            if maybeAttr {
                attrName = string(data[pb:p])
                fmt.Printf("Attr name: %d %d %s \n", pb, p, attrName)
                maybeAttr = false
            }
        
	goto st80
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
//line experiment.go:1048
		switch data[p] {
		case 32:
			goto st13
		case 34:
			goto st8
		case 39:
			goto st8
		case 47:
			goto st8
		case 60:
			goto st3
		case 61:
			goto tr29
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st13
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr17
				}
			case data[p] >= 65:
				goto tr17
			}
		default:
			goto st2
		}
		goto st7
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		switch data[p] {
		case 32:
			goto st13
		case 34:
			goto st8
		case 39:
			goto st8
		case 47:
			goto st8
		case 60:
			goto st3
		case 61:
			goto tr29
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st13
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr17
				}
			case data[p] >= 65:
				goto tr17
			}
		default:
			goto st2
		}
		goto st7
tr27:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line experiment.go:1141
		switch data[p] {
		case 32:
			goto tr25
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 61:
			goto tr26
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr30
				}
			case data[p] >= 65:
				goto tr30
			}
		default:
			goto st2
		}
		goto tr19
tr30:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line experiment.go:1195
		switch data[p] {
		case 32:
			goto tr25
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 61:
			goto tr26
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr31
				}
			case data[p] >= 65:
				goto tr31
			}
		default:
			goto st2
		}
		goto tr19
tr31:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line experiment.go:1249
		switch data[p] {
		case 32:
			goto tr25
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 61:
			goto tr26
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr32
				}
			case data[p] >= 65:
				goto tr32
			}
		default:
			goto st2
		}
		goto tr19
tr32:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line experiment.go:1303
		switch data[p] {
		case 32:
			goto tr25
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 61:
			goto tr26
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr33
				}
			case data[p] >= 65:
				goto tr33
			}
		default:
			goto st2
		}
		goto tr19
tr33:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line experiment.go:1357
		switch data[p] {
		case 32:
			goto tr25
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 61:
			goto tr26
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr34
				}
			case data[p] >= 65:
				goto tr34
			}
		default:
			goto st2
		}
		goto tr19
tr34:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line experiment.go:1411
		switch data[p] {
		case 32:
			goto tr25
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 61:
			goto tr26
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr35
				}
			case data[p] >= 65:
				goto tr35
			}
		default:
			goto st2
		}
		goto tr19
tr35:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line experiment.go:1465
		switch data[p] {
		case 32:
			goto tr25
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 61:
			goto tr26
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr36
				}
			case data[p] >= 65:
				goto tr36
			}
		default:
			goto st2
		}
		goto tr19
tr36:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line experiment.go:1519
		switch data[p] {
		case 32:
			goto tr25
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 61:
			goto tr26
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr37
				}
			case data[p] >= 65:
				goto tr37
			}
		default:
			goto st2
		}
		goto tr19
tr37:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line experiment.go:1573
		switch data[p] {
		case 32:
			goto tr25
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 61:
			goto tr26
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr38
				}
			case data[p] >= 65:
				goto tr38
			}
		default:
			goto st2
		}
		goto tr19
tr38:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line experiment.go:1627
		switch data[p] {
		case 32:
			goto tr25
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 61:
			goto tr26
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr39
				}
			case data[p] >= 65:
				goto tr39
			}
		default:
			goto st2
		}
		goto tr19
tr39:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line experiment.go:1681
		switch data[p] {
		case 32:
			goto tr25
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 61:
			goto tr26
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr40
				}
			case data[p] >= 65:
				goto tr40
			}
		default:
			goto st2
		}
		goto tr19
tr40:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line experiment.go:1735
		switch data[p] {
		case 32:
			goto tr25
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 61:
			goto tr26
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr41
				}
			case data[p] >= 65:
				goto tr41
			}
		default:
			goto st2
		}
		goto tr19
tr41:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line experiment.go:1789
		switch data[p] {
		case 32:
			goto tr25
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 61:
			goto tr26
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr42
				}
			case data[p] >= 65:
				goto tr42
			}
		default:
			goto st2
		}
		goto tr19
tr42:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
//line experiment.go:1843
		switch data[p] {
		case 32:
			goto tr25
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 61:
			goto tr26
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr43
				}
			case data[p] >= 65:
				goto tr43
			}
		default:
			goto st2
		}
		goto tr19
tr43:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line experiment.go:1897
		switch data[p] {
		case 32:
			goto tr25
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 61:
			goto tr26
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr44
				}
			case data[p] >= 65:
				goto tr44
			}
		default:
			goto st2
		}
		goto tr19
tr44:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
//line experiment.go:1951
		switch data[p] {
		case 32:
			goto tr25
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 61:
			goto tr26
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr45
				}
			case data[p] >= 65:
				goto tr45
			}
		default:
			goto st2
		}
		goto tr19
tr45:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
//line experiment.go:2005
		switch data[p] {
		case 32:
			goto tr25
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 61:
			goto tr26
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr46
				}
			case data[p] >= 65:
				goto tr46
			}
		default:
			goto st2
		}
		goto tr19
tr46:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
//line experiment.go:2059
		switch data[p] {
		case 32:
			goto tr25
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 61:
			goto tr26
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr47
				}
			case data[p] >= 65:
				goto tr47
			}
		default:
			goto st2
		}
		goto tr19
tr47:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
//line experiment.go:2113
		switch data[p] {
		case 32:
			goto tr25
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 61:
			goto tr26
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr48
				}
			case data[p] >= 65:
				goto tr48
			}
		default:
			goto st2
		}
		goto tr19
tr48:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st33
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
//line experiment.go:2167
		switch data[p] {
		case 32:
			goto tr25
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 61:
			goto tr26
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr49
				}
			case data[p] >= 65:
				goto tr49
			}
		default:
			goto st2
		}
		goto tr19
tr49:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st34
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
//line experiment.go:2221
		switch data[p] {
		case 32:
			goto tr25
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 61:
			goto tr26
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr50
				}
			case data[p] >= 65:
				goto tr50
			}
		default:
			goto st2
		}
		goto tr19
tr50:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st35
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
//line experiment.go:2275
		switch data[p] {
		case 32:
			goto tr25
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 61:
			goto tr26
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr51
				}
			case data[p] >= 65:
				goto tr51
			}
		default:
			goto st2
		}
		goto tr19
tr51:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
//line experiment.go:2329
		switch data[p] {
		case 32:
			goto tr25
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 61:
			goto tr26
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr52
				}
			case data[p] >= 65:
				goto tr52
			}
		default:
			goto st2
		}
		goto tr19
tr52:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st37
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
//line experiment.go:2383
		switch data[p] {
		case 32:
			goto tr25
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 61:
			goto tr26
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr53
				}
			case data[p] >= 65:
				goto tr53
			}
		default:
			goto st2
		}
		goto tr19
tr53:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st38
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
//line experiment.go:2437
		switch data[p] {
		case 32:
			goto tr25
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 61:
			goto tr26
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr54
				}
			case data[p] >= 65:
				goto tr54
			}
		default:
			goto st2
		}
		goto tr19
tr54:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
//line experiment.go:2491
		switch data[p] {
		case 32:
			goto tr25
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 61:
			goto tr26
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr55
				}
			case data[p] >= 65:
				goto tr55
			}
		default:
			goto st2
		}
		goto tr19
tr55:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st40
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
//line experiment.go:2545
		switch data[p] {
		case 32:
			goto tr25
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 47:
			goto tr20
		case 60:
			goto tr21
		case 61:
			goto tr26
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr23
				}
			case data[p] >= 65:
				goto tr23
			}
		default:
			goto st2
		}
		goto tr19
tr5:
//line experiment.rl:71

            maybeAttr = true
            attrName = ""
            pb = p
        
	goto st41
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
//line experiment.go:2593
		switch data[p] {
		case 32:
			goto st1
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st42
			}
		default:
			goto st42
		}
		goto st2
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		switch data[p] {
		case 32:
			goto st1
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st43
			}
		default:
			goto st43
		}
		goto st2
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		switch data[p] {
		case 32:
			goto tr58
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		case 61:
			goto tr59
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr58
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st45
			}
		default:
			goto st45
		}
		goto st2
tr58:
//line experiment.rl:77

            if maybeAttr {
                attrName = string(data[pb:p])
                fmt.Printf("Attr name: %d %d %s \n", pb, p, attrName)
                maybeAttr = false
            }
        
	goto st44
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
//line experiment.go:2696
		switch data[p] {
		case 32:
			goto st44
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		case 61:
			goto tr62
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st44
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr5
			}
		default:
			goto tr5
		}
		goto st2
tr62:
//line experiment.rl:85

            if len(attrName) > 0 {
                return true
            }
        
	goto st81
tr59:
//line experiment.rl:77

            if maybeAttr {
                attrName = string(data[pb:p])
                fmt.Printf("Attr name: %d %d %s \n", pb, p, attrName)
                maybeAttr = false
            }
        
//line experiment.rl:85

            if len(attrName) > 0 {
                return true
            }
        
	goto st81
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
//line experiment.go:2753
		switch data[p] {
		case 32:
			goto st1
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st1
		}
		goto st2
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		switch data[p] {
		case 32:
			goto tr58
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		case 61:
			goto tr59
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr58
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st46
			}
		default:
			goto st46
		}
		goto st2
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		switch data[p] {
		case 32:
			goto tr58
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		case 61:
			goto tr59
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr58
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st47
			}
		default:
			goto st47
		}
		goto st2
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		switch data[p] {
		case 32:
			goto tr58
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		case 61:
			goto tr59
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr58
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st48
			}
		default:
			goto st48
		}
		goto st2
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		switch data[p] {
		case 32:
			goto tr58
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		case 61:
			goto tr59
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr58
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st49
			}
		default:
			goto st49
		}
		goto st2
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		switch data[p] {
		case 32:
			goto tr58
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		case 61:
			goto tr59
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr58
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st50
			}
		default:
			goto st50
		}
		goto st2
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		switch data[p] {
		case 32:
			goto tr58
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		case 61:
			goto tr59
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr58
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st51
			}
		default:
			goto st51
		}
		goto st2
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
		switch data[p] {
		case 32:
			goto tr58
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		case 61:
			goto tr59
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr58
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st52
			}
		default:
			goto st52
		}
		goto st2
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
		switch data[p] {
		case 32:
			goto tr58
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		case 61:
			goto tr59
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr58
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st53
			}
		default:
			goto st53
		}
		goto st2
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
		switch data[p] {
		case 32:
			goto tr58
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		case 61:
			goto tr59
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr58
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st54
			}
		default:
			goto st54
		}
		goto st2
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		switch data[p] {
		case 32:
			goto tr58
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		case 61:
			goto tr59
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr58
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st55
			}
		default:
			goto st55
		}
		goto st2
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
		switch data[p] {
		case 32:
			goto tr58
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		case 61:
			goto tr59
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr58
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st56
			}
		default:
			goto st56
		}
		goto st2
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		switch data[p] {
		case 32:
			goto tr58
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		case 61:
			goto tr59
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr58
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st57
			}
		default:
			goto st57
		}
		goto st2
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		switch data[p] {
		case 32:
			goto tr58
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		case 61:
			goto tr59
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr58
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st58
			}
		default:
			goto st58
		}
		goto st2
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		switch data[p] {
		case 32:
			goto tr58
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		case 61:
			goto tr59
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr58
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st59
			}
		default:
			goto st59
		}
		goto st2
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		switch data[p] {
		case 32:
			goto tr58
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		case 61:
			goto tr59
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr58
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st60
			}
		default:
			goto st60
		}
		goto st2
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
		switch data[p] {
		case 32:
			goto tr58
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		case 61:
			goto tr59
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr58
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st61
			}
		default:
			goto st61
		}
		goto st2
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		switch data[p] {
		case 32:
			goto tr58
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		case 61:
			goto tr59
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr58
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st62
			}
		default:
			goto st62
		}
		goto st2
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		switch data[p] {
		case 32:
			goto tr58
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		case 61:
			goto tr59
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr58
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st63
			}
		default:
			goto st63
		}
		goto st2
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		switch data[p] {
		case 32:
			goto tr58
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		case 61:
			goto tr59
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr58
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st64
			}
		default:
			goto st64
		}
		goto st2
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		switch data[p] {
		case 32:
			goto tr58
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		case 61:
			goto tr59
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr58
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st65
			}
		default:
			goto st65
		}
		goto st2
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		switch data[p] {
		case 32:
			goto tr58
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		case 61:
			goto tr59
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr58
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st66
			}
		default:
			goto st66
		}
		goto st2
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
		switch data[p] {
		case 32:
			goto tr58
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		case 61:
			goto tr59
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr58
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st67
			}
		default:
			goto st67
		}
		goto st2
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
		switch data[p] {
		case 32:
			goto tr58
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		case 61:
			goto tr59
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr58
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st68
			}
		default:
			goto st68
		}
		goto st2
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
		switch data[p] {
		case 32:
			goto tr58
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		case 61:
			goto tr59
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr58
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st69
			}
		default:
			goto st69
		}
		goto st2
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		switch data[p] {
		case 32:
			goto tr58
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		case 61:
			goto tr59
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr58
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st70
			}
		default:
			goto st70
		}
		goto st2
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
		switch data[p] {
		case 32:
			goto tr58
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		case 61:
			goto tr59
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr58
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st71
			}
		default:
			goto st71
		}
		goto st2
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
		switch data[p] {
		case 32:
			goto tr58
		case 34:
			goto st1
		case 39:
			goto st1
		case 47:
			goto st1
		case 60:
			goto st3
		case 61:
			goto tr59
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr58
		}
		goto st2
tr90:
//line experiment.rl:45

            maybeTag = true
            tagNameBuilder.Reset()
        
	goto st72
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
//line experiment.go:3637
		switch data[p] {
		case 34:
			goto tr7
		case 39:
			goto tr7
		case 60:
			goto tr90
		case 62:
			goto tr91
		case 95:
			goto tr8
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr8
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr89
tr89:
//line experiment.rl:45

            maybeTag = true
            tagNameBuilder.Reset()
        
	goto st73
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
//line experiment.go:3675
		switch data[p] {
		case 34:
			goto st5
		case 39:
			goto st5
		case 60:
			goto st72
		case 62:
			goto st75
		case 95:
			goto st74
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st74
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr95
			}
		default:
			goto tr95
		}
		goto st73
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
		switch data[p] {
		case 34:
			goto st1
		case 39:
			goto st1
		case 58:
			goto st75
		case 60:
			goto st72
		}
		goto st0
tr91:
//line experiment.rl:45

            maybeTag = true
            tagNameBuilder.Reset()
        
	goto st75
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
//line experiment.go:3729
		switch data[p] {
		case 34:
			goto st8
		case 39:
			goto st8
		case 60:
			goto st72
		case 95:
			goto st0
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st0
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr95
			}
		default:
			goto tr95
		}
		goto st75
tr95:
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st76
tr98:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:39

            if maybeTag {
                pb = p
            }
        
	goto st76
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
//line experiment.go:3780
		switch data[p] {
		case 34:
			goto tr20
		case 39:
			goto tr20
		case 60:
			goto tr97
		case 95:
			goto st0
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st0
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr98
			}
		default:
			goto tr98
		}
		goto tr96
tr96:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:56

            if maybeTag {
                tagName = tagNameBuilder.String()
                fmt.Println(tagName)

                if slices.Contains(riskTags, tagName) {
                    fmt.Println("Matched")
                    return true
                }

                tagNameBuilder.Reset()
                maybeTag = false
            }
        
	goto st82
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
//line experiment.go:3832
		switch data[p] {
		case 34:
			goto st8
		case 39:
			goto st8
		case 60:
			goto st72
		case 95:
			goto st0
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st0
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr95
			}
		default:
			goto tr95
		}
		goto st75
tr97:
//line experiment.rl:50

            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        
//line experiment.rl:56

            if maybeTag {
                tagName = tagNameBuilder.String()
                fmt.Println(tagName)

                if slices.Contains(riskTags, tagName) {
                    fmt.Println("Matched")
                    return true
                }

                tagNameBuilder.Reset()
                maybeTag = false
            }
        
	goto st83
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
//line experiment.go:3884
		switch data[p] {
		case 34:
			goto tr7
		case 39:
			goto tr7
		case 60:
			goto tr90
		case 62:
			goto tr91
		case 95:
			goto tr8
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr8
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr89
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
	_test_eof77: cs = 77; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof78: cs = 78; goto _test_eof
	_test_eof79: cs = 79; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof80: cs = 80; goto _test_eof
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
	_test_eof81: cs = 81; goto _test_eof
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
	_test_eof70: cs = 70; goto _test_eof
	_test_eof71: cs = 71; goto _test_eof
	_test_eof72: cs = 72; goto _test_eof
	_test_eof73: cs = 73; goto _test_eof
	_test_eof74: cs = 74; goto _test_eof
	_test_eof75: cs = 75; goto _test_eof
	_test_eof76: cs = 76; goto _test_eof
	_test_eof82: cs = 82; goto _test_eof
	_test_eof83: cs = 83; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 9, 10, 11, 12, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 76:
//line experiment.rl:56

            if maybeTag {
                tagName = tagNameBuilder.String()
                fmt.Println(tagName)

                if slices.Contains(riskTags, tagName) {
                    fmt.Println("Matched")
                    return true
                }

                tagNameBuilder.Reset()
                maybeTag = false
            }
        
//line experiment.go:4015
		}
	}

	}

//line experiment.rl:118

        return false
}
