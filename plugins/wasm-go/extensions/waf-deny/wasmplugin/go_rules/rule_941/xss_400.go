
//line xss_400.rl:1
package rule_941

import (
    "fmt"
)

func matchXSS400(data []byte) bool {

//line xss_400.rl:9

//line xss_400.go:14
var _xss400_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 7, 
	1, 8, 1, 9, 1, 10, 1, 11, 
	1, 12, 2, 2, 3, 2, 2, 4, 
	2, 2, 5, 2, 2, 6, 
}

var _xss400_key_offsets []byte = []byte{
	0, 3, 7, 10, 14, 15, 19, 23, 
	27, 31, 35, 39, 40, 41, 42, 43, 
	44, 45, 46, 47, 48, 49, 53, 56, 
	57, 58, 59, 60, 61, 
}

var _xss400_trans_keys []byte = []byte{
	97, 109, 115, 97, 109, 112, 115, 97, 
	109, 115, 97, 109, 112, 115, 46, 97, 
	109, 111, 115, 97, 109, 114, 115, 97, 
	109, 115, 116, 97, 109, 112, 115, 97, 
	108, 109, 115, 97, 109, 115, 121, 93, 
	108, 108, 96, 96, 102, 108, 101, 99, 
	116, 46, 91, 99, 114, 97, 109, 115, 
	93, 46, 97, 96, 101, 46, 
}

var _xss400_single_lengths []byte = []byte{
	3, 4, 3, 4, 1, 4, 4, 4, 
	4, 4, 4, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 4, 3, 1, 
	1, 1, 1, 1, 1, 
}

var _xss400_range_lengths []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 
}

var _xss400_index_offsets []byte = []byte{
	0, 4, 9, 13, 18, 20, 25, 30, 
	35, 40, 45, 50, 52, 54, 56, 58, 
	60, 62, 64, 66, 68, 70, 75, 79, 
	81, 83, 85, 87, 89, 
}

var _xss400_indicies []byte = []byte{
	2, 3, 4, 1, 2, 3, 5, 4, 
	1, 6, 3, 4, 1, 2, 3, 7, 
	4, 1, 8, 7, 2, 3, 9, 4, 
	1, 2, 3, 10, 4, 1, 2, 3, 
	4, 7, 1, 2, 3, 11, 4, 1, 
	2, 12, 3, 4, 1, 2, 3, 4, 
	7, 1, 14, 13, 16, 15, 17, 15, 
	18, 17, 19, 18, 20, 15, 21, 15, 
	22, 15, 23, 15, 24, 15, 26, 27, 
	28, 29, 25, 2, 3, 4, 1, 14, 
	13, 13, 14, 32, 30, 19, 18, 34, 
	30, 31, 24, 
}

var _xss400_trans_targs []byte = []byte{
	21, 0, 1, 2, 5, 8, 3, 4, 
	22, 6, 7, 9, 10, 11, 24, 21, 
	13, 14, 15, 26, 17, 18, 19, 20, 
	28, 21, 22, 23, 25, 27, 21, 21, 
	12, 21, 16, 
}

var _xss400_trans_actions []byte = []byte{
	17, 0, 0, 0, 0, 0, 0, 0, 
	22, 0, 0, 0, 0, 0, 19, 15, 
	0, 0, 0, 25, 0, 0, 0, 0, 
	0, 7, 28, 28, 28, 5, 13, 9, 
	0, 11, 0, 
}

var _xss400_to_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 1, 0, 0, 
	0, 0, 0, 0, 0, 
}

var _xss400_from_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 3, 0, 0, 
	0, 0, 0, 0, 0, 
}

var _xss400_eof_trans []byte = []byte{
	1, 1, 1, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 16, 16, 16, 1, 
	16, 16, 16, 16, 16, 0, 1, 31, 
	32, 31, 34, 31, 32, 
}

const xss400_start int = 21
const xss400_first_final int = 21
const xss400_error int = -1

const xss400_en_main int = 21


//line xss_400.rl:10
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    var ts, te, act int
        _, _, _ = ts, te, act

    step := 0

    
//line xss_400.go:130
	{
	cs = xss400_start
	ts = 0
	te = 0
	act = 0
	}

//line xss_400.go:138
	{
	var _klen int
	var _trans int
	var _acts int
	var _nacts uint
	var _keys int
	if p == pe {
		goto _test_eof
	}
_resume:
	_acts = int(_xss400_from_state_actions[cs])
	_nacts = uint(_xss400_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		 _acts++
		switch _xss400_actions[_acts - 1] {
		case 1:
//line NONE:1
ts = p

//line xss_400.go:158
		}
	}

	_keys = int(_xss400_key_offsets[cs])
	_trans = int(_xss400_index_offsets[cs])

	_klen = int(_xss400_single_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + _klen - 1)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + ((_upper - _lower) >> 1)
			switch {
			case data[p] < _xss400_trans_keys[_mid]:
				_upper = _mid - 1
			case data[p] > _xss400_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_xss400_range_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + (_klen << 1) - 2)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + (((_upper - _lower) >> 1) & ^1)
			switch {
			case data[p] < _xss400_trans_keys[_mid]:
				_upper = _mid - 2
			case data[p] > _xss400_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
	_trans = int(_xss400_indicies[_trans])
_eof_trans:
	cs = int(_xss400_trans_targs[_trans])

	if _xss400_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_xss400_trans_actions[_trans])
	_nacts = uint(_xss400_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _xss400_actions[_acts-1] {
		case 2:
//line NONE:1
te = p+1

		case 3:
//line xss_400.rl:28
act = 1;
		case 4:
//line xss_400.rl:33
act = 2;
		case 5:
//line xss_400.rl:43
act = 3;
		case 6:
//line xss_400.rl:53
act = 4;
		case 7:
//line xss_400.rl:53
te = p+1

		case 8:
//line xss_400.rl:28
te = p
p--
{
                fmt.Printf("match step1 %s \n", string(data[ts:te]))
                step = 1
            }
		case 9:
//line xss_400.rl:43
te = p
p--
{
                if step == 2 {
                    fmt.Printf("match step3 %s \n", string(data[ts:te]))
                    return true
                } else {
                    fmt.Println("back when match 3")
                    step = 0
                }
            }
		case 10:
//line xss_400.rl:53
te = p
p--

		case 11:
//line xss_400.rl:53
p = (te) - 1

		case 12:
//line NONE:1
	switch act {
	case 1:
	{p = (te) - 1

                fmt.Printf("match step1 %s \n", string(data[ts:te]))
                step = 1
            }
	case 2:
	{p = (te) - 1

                if step == 1 {
                    fmt.Printf("match step2 %s \n", string(data[ts:te]))
                    step = 2
                } else {
                    fmt.Println("back when match 2")
                    step = 0
                }
            }
	case 3:
	{p = (te) - 1

                if step == 2 {
                    fmt.Printf("match step3 %s \n", string(data[ts:te]))
                    return true
                } else {
                    fmt.Println("back when match 3")
                    step = 0
                }
            }
	default:
	{p = (te) - 1
}
	}
	
//line xss_400.go:314
		}
	}

_again:
	_acts = int(_xss400_to_state_actions[cs])
	_nacts = uint(_xss400_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _xss400_actions[_acts-1] {
		case 0:
//line NONE:1
ts = 0

//line xss_400.go:328
		}
	}

	p++
	if p != pe {
		goto _resume
	}
	_test_eof: {}
	if p == eof {
		if _xss400_eof_trans[cs] > 0 {
			_trans = int(_xss400_eof_trans[cs] - 1)
			goto _eof_trans
		}
	}

	}

//line xss_400.rl:58


    return false
}
