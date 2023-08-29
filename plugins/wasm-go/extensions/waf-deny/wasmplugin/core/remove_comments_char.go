package core

func RemoveCommentsChar(value string) (string, bool, error) {
	inputLen := len(value)
	res := make([]byte, 0, inputLen)
	changed := false
	for i := 0; i < inputLen; {
		switch {
		case value[i] == '/' && (i+1 < inputLen) && value[i+1] == '*':
			i += 2
			changed = true
		case value[i] == '*' && (i+1 < inputLen) && value[i+1] == '/':
			i += 2
			changed = true
		case value[i] == '<' &&
			(i+1 < inputLen) &&
			value[i+1] == '!' &&
			(i+2 < inputLen) &&
			value[i+2] == '-' &&
			(i+3 < inputLen) &&
			value[i+3] == '-':
			i += 4
			changed = true
		case value[i] == '-' &&
			(i+1 < inputLen) && value[i+1] == '-' &&
			(i+2 < inputLen) && value[i+2] == '>':
			i += 3
			changed = true
		case value[i] == '-' && (i+1 < inputLen) && value[i+1] == '-':
			i += 2
			changed = true
		case value[i] == '#':
			i += 1
			changed = true
		default:
			res = append(res, value[i])
			i += 1
		}
	}
	return WrapUnsafe(res), changed, nil
}
