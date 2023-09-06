package rule_941

func matchXSS(data []byte) bool {

	//line xss.rl:5

	//line xss.go:10
	var _xss_actions []byte = []byte{
		0, 1, 0,
	}

	var _xss_key_offsets []byte = []byte{
		0, 1, 3, 5, 7, 9, 11, 13,
		14,
	}

	var _xss_trans_keys []byte = []byte{
		60, 60, 115, 60, 99, 60, 114, 60,
		105, 60, 112, 60, 116, 62, 62,
	}

	var _xss_single_lengths []byte = []byte{
		1, 2, 2, 2, 2, 2, 2, 1,
		1,
	}

	var _xss_range_lengths []byte = []byte{
		0, 0, 0, 0, 0, 0, 0, 0,
		0,
	}

	var _xss_index_offsets []byte = []byte{
		0, 2, 5, 8, 11, 14, 17, 20,
		22,
	}

	var _xss_indicies []byte = []byte{
		1, 0, 1, 2, 0, 1, 3, 0,
		1, 4, 0, 1, 5, 0, 1, 6,
		0, 1, 7, 0, 8, 7, 8, 7,
	}

	var _xss_trans_targs []byte = []byte{
		0, 1, 2, 3, 4, 5, 6, 7,
		8,
	}

	var _xss_trans_actions []byte = []byte{
		0, 0, 0, 0, 0, 0, 0, 0,
		1,
	}

	const xss_start int = 0
	const xss_first_final int = 8
	const xss_error int = -1

	const xss_en_main int = 0

	//line xss.rl:6
	cs, p, pe, eof := 0, 0, len(data), len(data)
	_ = eof

	//line xss.go:68
	{
		cs = xss_start
	}

	//line xss.go:73
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
		_keys = int(_xss_key_offsets[cs])
		_trans = int(_xss_index_offsets[cs])

		_klen = int(_xss_single_lengths[cs])
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
				case data[p] < _xss_trans_keys[_mid]:
					_upper = _mid - 1
				case data[p] > _xss_trans_keys[_mid]:
					_lower = _mid + 1
				default:
					_trans += int(_mid - int(_keys))
					goto _match
				}
			}
			_keys += _klen
			_trans += _klen
		}

		_klen = int(_xss_range_lengths[cs])
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
				case data[p] < _xss_trans_keys[_mid]:
					_upper = _mid - 2
				case data[p] > _xss_trans_keys[_mid+1]:
					_lower = _mid + 2
				default:
					_trans += int((_mid - int(_keys)) >> 1)
					goto _match
				}
			}
			_trans += _klen
		}

	_match:
		_trans = int(_xss_indicies[_trans])
		cs = int(_xss_trans_targs[_trans])

		if _xss_trans_actions[_trans] == 0 {
			goto _again
		}

		_acts = int(_xss_trans_actions[_trans])
		_nacts = uint(_xss_actions[_acts])
		_acts++
		for ; _nacts > 0; _nacts-- {
			_acts++
			switch _xss_actions[_acts-1] {
			case 0:
				//line xss.rl:9
				return true
				//line xss.go:152
			}
		}

	_again:
		p++
		if p != pe {
			goto _resume
		}
	_test_eof:
		{
		}
	}

	//line xss.rl:12

	return false
}
