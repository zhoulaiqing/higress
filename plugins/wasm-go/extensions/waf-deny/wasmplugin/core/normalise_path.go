package core

import (
	"path/filepath"
	"strings"
)

func NormalisePath(data string) (string, bool, error) {
	leng := len(data)
	if leng < 1 {
		return data, false, nil
	}
	clean := filepath.Clean(data)
	if clean == "." {
		return "", true, nil
	}
	if data[len(data)-1] == '/' {
		return clean + "/", true, nil
	}
	return clean, data != clean, nil
}

func NormalisePathWin(data string) (string, bool, error) {
	leng := len(data)
	if leng < 1 {
		return data, false, nil
	}
	data = strings.ReplaceAll(data, "\\", "/")
	return NormalisePath(data)
}
