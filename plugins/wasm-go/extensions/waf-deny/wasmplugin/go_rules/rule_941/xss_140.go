
//line xss_140.rl:1
package rule_941


func matchXSS140(data []byte) bool {

//line xss_140.rl:6

//line xss_140.go:11
var _xss140_actions []byte = []byte{
	0, 1, 0, 
}

var _xss140_key_offsets []byte = []byte{
	0, 2, 5, 9, 13, 16, 19, 23, 
	27, 31, 35, 40, 45, 50, 54, 59, 
	64, 69, 74, 79, 84, 89, 94, 99, 
	103, 107, 111, 114, 118, 122, 126, 130, 
	135, 140, 145, 149, 154, 159, 164, 169, 
	174, 179, 184, 189, 194, 
}

var _xss140_trans_keys []byte = []byte{
	97, 122, 61, 97, 122, 58, 61, 97, 
	122, 58, 61, 97, 122, 117, 97, 122, 
	59, 97, 122, 58, 61, 97, 122, 58, 
	61, 97, 122, 59, 117, 97, 122, 59, 
	61, 97, 122, 59, 61, 114, 97, 122, 
	59, 61, 108, 97, 122, 40, 59, 61, 
	97, 122, 59, 106, 97, 122, 59, 61, 
	97, 98, 122, 59, 61, 118, 97, 122, 
	59, 61, 97, 98, 122, 59, 61, 115, 
	97, 122, 59, 61, 99, 97, 122, 59, 
	61, 114, 97, 122, 59, 61, 105, 97, 
	122, 59, 61, 112, 97, 122, 59, 61, 
	116, 97, 122, 58, 61, 97, 122, 58, 
	61, 97, 122, 59, 61, 97, 122, 59, 
	97, 122, 58, 61, 97, 122, 58, 61, 
	97, 122, 59, 117, 97, 122, 59, 61, 
	97, 122, 59, 61, 114, 97, 122, 59, 
	61, 108, 97, 122, 40, 59, 61, 97, 
	122, 59, 106, 97, 122, 59, 61, 97, 
	98, 122, 59, 61, 118, 97, 122, 59, 
	61, 97, 98, 122, 59, 61, 115, 97, 
	122, 59, 61, 99, 97, 122, 59, 61, 
	114, 97, 122, 59, 61, 105, 97, 122, 
	59, 61, 112, 97, 122, 59, 61, 116, 
	97, 122, 58, 61, 97, 122, 
}

var _xss140_single_lengths []byte = []byte{
	0, 1, 2, 2, 1, 1, 2, 2, 
	2, 2, 3, 3, 3, 2, 3, 3, 
	3, 3, 3, 3, 3, 3, 3, 2, 
	2, 2, 1, 2, 2, 2, 2, 3, 
	3, 3, 2, 3, 3, 3, 3, 3, 
	3, 3, 3, 3, 2, 
}

var _xss140_range_lengths []byte = []byte{
	1, 1, 1, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 
}

var _xss140_index_offsets []byte = []byte{
	0, 2, 5, 9, 13, 16, 19, 23, 
	27, 31, 35, 40, 45, 50, 54, 59, 
	64, 69, 74, 79, 84, 89, 94, 99, 
	103, 107, 111, 114, 118, 122, 126, 130, 
	135, 140, 145, 149, 154, 159, 164, 169, 
	174, 179, 184, 189, 194, 
}

var _xss140_indicies []byte = []byte{
	1, 0, 2, 1, 0, 0, 0, 4, 
	3, 5, 0, 4, 3, 8, 7, 6, 
	9, 7, 6, 6, 6, 11, 10, 12, 
	6, 11, 10, 9, 8, 7, 6, 9, 
	9, 7, 6, 9, 9, 13, 7, 6, 
	9, 9, 14, 7, 6, 15, 9, 9, 
	7, 6, 9, 16, 7, 6, 9, 9, 
	17, 7, 6, 9, 9, 18, 7, 6, 
	9, 9, 19, 7, 6, 9, 9, 20, 
	7, 6, 9, 9, 21, 7, 6, 9, 
	9, 22, 7, 6, 9, 9, 23, 7, 
	6, 9, 9, 24, 7, 6, 9, 9, 
	25, 7, 6, 12, 9, 11, 10, 5, 
	2, 4, 3, 27, 27, 28, 26, 30, 
	31, 29, 29, 29, 33, 32, 34, 29, 
	33, 32, 30, 35, 31, 29, 30, 30, 
	31, 29, 30, 30, 36, 31, 29, 30, 
	30, 37, 31, 29, 38, 30, 30, 31, 
	29, 30, 39, 31, 29, 30, 30, 40, 
	31, 29, 30, 30, 41, 31, 29, 30, 
	30, 42, 31, 29, 30, 30, 43, 31, 
	29, 30, 30, 44, 31, 29, 30, 30, 
	45, 31, 29, 30, 30, 46, 31, 29, 
	30, 30, 47, 31, 29, 30, 30, 25, 
	31, 29, 34, 30, 33, 32, 
}

var _xss140_trans_targs []byte = []byte{
	0, 1, 2, 3, 24, 4, 5, 9, 
	10, 6, 7, 23, 8, 11, 12, 13, 
	14, 15, 16, 17, 18, 19, 20, 21, 
	22, 25, 26, 27, 30, 26, 27, 30, 
	28, 44, 29, 31, 32, 33, 34, 35, 
	36, 37, 38, 39, 40, 41, 42, 43, 
}

var _xss140_trans_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 1, 1, 1, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
}

var _xss140_eof_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 1, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 
}

const xss140_start int = 0
const xss140_first_final int = 25
const xss140_error int = -1

const xss140_en_main int = 0


//line xss_140.rl:7
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof
    
//line xss_140.go:146
	{
	cs = xss140_start
	}

//line xss_140.go:151
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
	_keys = int(_xss140_key_offsets[cs])
	_trans = int(_xss140_index_offsets[cs])

	_klen = int(_xss140_single_lengths[cs])
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
			case data[p] < _xss140_trans_keys[_mid]:
				_upper = _mid - 1
			case data[p] > _xss140_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_xss140_range_lengths[cs])
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
			case data[p] < _xss140_trans_keys[_mid]:
				_upper = _mid - 2
			case data[p] > _xss140_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
	_trans = int(_xss140_indicies[_trans])
	cs = int(_xss140_trans_targs[_trans])

	if _xss140_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_xss140_trans_actions[_trans])
	_nacts = uint(_xss140_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _xss140_actions[_acts-1] {
		case 0:
//line xss_140.rl:10

            return true
        
//line xss_140.go:232
		}
	}

_again:
	p++
	if p != pe {
		goto _resume
	}
	_test_eof: {}
	if p == eof {
		__acts := _xss140_eof_actions[cs]
		__nacts := uint(_xss140_actions[__acts]); __acts++
		for ; __nacts > 0; __nacts-- {
			__acts++
			switch _xss140_actions[__acts-1] {
			case 0:
//line xss_140.rl:10

            return true
        
//line xss_140.go:253
			}
		}
	}

	}

//line xss_140.rl:19


    return false
}
