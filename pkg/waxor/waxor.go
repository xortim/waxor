package waxor

import "encoding/base64"

// EncodeXor {xor} encodes a string
func EncodeXor(str string) string {
	var result string
	for _, b := range []byte(str) {
		result = result + string(b^95)
	}
	result = "{xor}" + base64.StdEncoding.EncodeToString([]byte(result))
	return result
}

// DecodeXor {xor} decodes a string
func DecodeXor(str string) string {
	var result string
	bytes, _ := base64.StdEncoding.DecodeString(str[5:])
	for _, b := range bytes {
		result = result + string(b^95)
	}
	return result
}
