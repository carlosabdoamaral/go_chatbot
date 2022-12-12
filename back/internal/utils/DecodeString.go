package utils

import "encoding/base64"

func DecodeString(s string) string {
	return _FromBase64(s)
}

func _FromBase64(s string) string {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic("[!] Error while decoding string")
	}

	return string(decoded)
}
