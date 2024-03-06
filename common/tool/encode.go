package tool

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strings"
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

// EncEmail 加密邮箱（只针对@前加密；<=3时填充3个*；4、5时保留前后1位，中间填充等量*；，>=6时保留前后2位，中间填充2个*）
func EncEmail(email string) string {
	idx := strings.Index(email, "@")
	if idx == -1 {
		return ""
	}

	prefix := email[:idx]
	suffix := email[idx:]
	l := len(prefix)
	if l <= 3 {
		return "***" + suffix
	} else if l <= 5 {
		return string(prefix[0]) + strings.Repeat("*", l-2) + string(prefix[l-1]) + suffix
	} else {
		return prefix[:2] + "**" + prefix[l-2:] + suffix
	}
}
