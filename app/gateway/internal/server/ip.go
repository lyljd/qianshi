package server

import (
	"net"
	"net/http"
	"strings"
)

func getIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		if ip := net.ParseIP(strings.TrimSpace(strings.Split(xff, ",")[0])); ip != nil {
			return ip.String()
		}
	}

	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		if ip := net.ParseIP(strings.TrimSpace(xri)); ip != nil {
			return ip.String()
		}
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err != nil {
		return ""
	} else {
		return ip
	}
}
