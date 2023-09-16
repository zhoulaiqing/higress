package rule_941

import (
	libinjection2 "github.com/wasilibs/go-libinjection"
)

func matchXss(value string) bool {
	//fmt.Printf("Try to match : %s \n", value)
	if doMatchXssOriginal(value) {
		return true
	}

	//value2 := strings.ReplaceAll(value, "\xbc", "<")
	//value2 = strings.ReplaceAll(value2, "\xbe", ">")
	//value2 = strings.ReplaceAll(value2, "\xa2", "\"")
	//
	//if value != value2 {
	//	if doMatchXssOriginal(value2) {
	//		return true
	//	}
	//}

	return false
}

func doMatchXssOriginal(value string) bool {
	//data := []byte(value)

	if libinjection2.IsXSS(value) {
		return true
	}

	//if matchExp(data) {
	//	return true
	//}
	//
	//if matchXSS130(data) {
	//	return true
	//}
	//if matchXSS140(data) {
	//	return true
	//}
	//if matchXSS170(data) {
	//	return true
	//}
	//if matchXSS200(data) {
	//	return true
	//}
	//if matchXSS210(data) {
	//	return true
	//}
	//
	//if matchXSS370(data) {
	//	return true
	//}
	//if matchXSS390(data) {
	//	return true
	//}
	//if matchXSS400(data) {
	//	return true
	//}

	return false
}
