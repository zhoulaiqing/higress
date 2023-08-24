package core

import "golang.org/x/net/html"

func HtmlEntityDecode(data string) (string, bool, error) {
	transformedData := html.UnescapeString(data)
	return transformedData, len(data) != len(transformedData), nil
}
