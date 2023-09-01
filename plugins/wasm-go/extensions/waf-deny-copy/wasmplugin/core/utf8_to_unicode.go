package core

import (
	"strconv"
	"unicode/utf8"
)

func Utf8ToUnicode(str string) (string, bool, error) {
	for i, c := range str {
		if c >= utf8.RuneSelf {
			return doUTF8ToUnicode(str, i), true, nil
		}
	}
	return str, false, nil
}

func doUTF8ToUnicode(input string, pos int) string {
	// Preallocate to length of input, the encoded string will be at least
	// as long.
	res := make([]byte, pos, len(input))
	copy(res, input[0:pos])

	for _, c := range input[pos:] {
		if c < utf8.RuneSelf {
			res = append(res, byte(c))
			continue
		}
		cHexLen := numHexDigits(c)
		res = append(res, '%', 'u')
		// Pad to 4 characters
		for i := 0; i < 4-cHexLen; i++ {
			res = append(res, '0')
		}
		res = strconv.AppendUint(res, uint64(c), 16)
	}

	return WrapUnsafe(res)
}

func numHexDigits(c rune) int {
	switch {
	case c <= 0xf:
		return 1
	case c <= 0xff:
		return 2
	case c <= 0xfff:
		return 3
	}
	return 4
}
