package core

import "strings"

func RemoveNulls(data string) (string, bool, error) {
	transformedData := strings.ReplaceAll(data, "\x00", "")
	return transformedData, len(data) != len(transformedData), nil
}
