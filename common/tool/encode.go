package tool

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func Base64Encode(b []byte) string {
	return base64.RawURLEncoding.EncodeToString(b)
}

func Base64Decode(s string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(s)
}

func Sha256(s, secret, salt string) (string, error) {
	h := hmac.New(sha256.New, []byte(secret+salt))
	if _, err := h.Write([]byte(s)); err != nil {
		return "", err
	}
	return Base64Encode(h.Sum(nil)), nil
}
