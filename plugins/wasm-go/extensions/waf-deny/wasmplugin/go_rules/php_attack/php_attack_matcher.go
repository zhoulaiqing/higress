package php_attack

import "strings"

func matchDefault(value string) bool {

	lowerV := strings.ToLower(value)

	data := []byte(lowerV)

	if matchPhpInjection(data) {
		return true
	}

	if matchPhpFunctions(data) {
		return true
	}

	return false
}

func matchFile(value string) bool {
	lowerV := strings.ToLower(value)
	data := []byte(lowerV)

	if matchPhpScriptUploads(data) {
		return true
	}

	return false
}
