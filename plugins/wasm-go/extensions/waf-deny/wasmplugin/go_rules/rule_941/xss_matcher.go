package rule_941

func matchXss(value string) bool {
	data := []byte(value)

	if matchExp(data) {
		return true
	}

	if matchXSS130(data) {
		return true
	}
	if matchXSS140(data) {
		return true
	}
	if matchXSS170(data) {
		return true
	}

	return false
}
