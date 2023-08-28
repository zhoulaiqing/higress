package core

import "encoding/base64"

func Base64decode(data string) (string, bool, error) {
	dec, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		// Forgiving implementation, which ignores invalid characters
		return data, false, nil
	}
	return WrapUnsafe(dec), true, nil
}
