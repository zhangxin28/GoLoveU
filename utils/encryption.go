package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"strings"
)

// EncodeMd5 returns the md5 string
func EncodeMd5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return hex.EncodeToString(hash.Sum(nil))
}

// Base64Encode returns the base64 string
func Base64Encode(raw []byte) string {
	t := base64.StdEncoding.EncodeToString(raw)
	t = strings.TrimSpace(t)
	t = strings.Replace(t, "\r", "", -1)
	t = strings.Replace(t, "\n", "", -1)
	t = strings.Replace(t, "\n\r", "", -1)
	t = strings.Replace(t, "\r\n", "", -1)
	return t
}

// Base64Decode returns the raw byte array for a base64 string
func Base64Decode(base64Str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(base64Str)
}
