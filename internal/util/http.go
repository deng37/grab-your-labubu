package util

import (
	"net/http"
	"net"
)

func UpdateHeaderJson(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func GetUserIP(r *http.Request) string {
	userIp := r.Header.Get("Fly-Client-IP")	// First checking on Fly-Client-IP

	if userIp != "" {
		return userIp
	}

	userIp = r.Header.Get("X-Forwarded-For") // Second checking on X-Forwarded-For
	if userIp != "" {
		return userIp
	}

	host, _, _ := net.SplitHostPort(r.RemoteAddr) // Last checking on RemoteAddr
	return host
}