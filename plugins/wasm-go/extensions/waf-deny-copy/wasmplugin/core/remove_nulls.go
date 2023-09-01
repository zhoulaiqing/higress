package core

import (
	"strings"
	"unicode"
)

func RemoveNulls(data string) (string, bool, error) {
	transformedData := strings.ReplaceAll(data, "\x00", "")
	return transformedData, len(data) != len(transformedData), nil
}

func RemoveWhitespace(data string) (string, bool, error) {
	changed := false
	transformedData := strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			// if the character is a space, drop it
			changed = true
			return -1
		}
		// else keep it in the string
		return r
	}, data)

	return transformedData, changed, nil
}
