package sip

import (
	"net"
	"strconv"

	"github.com/emiago/sipgo/sip"
)

func parseHostPort(raw string) (string, int, error) {
	host := raw
	port := 5060

	if h, p, err := net.SplitHostPort(raw); err == nil {
		host = h
		port, _ = strconv.Atoi(p)
	}

	return host, port, nil
}

func needsAuth(res *sip.Response) bool {
	return res.StatusCode == sip.StatusUnauthorized ||
		res.StatusCode == sip.StatusProxyAuthRequired
}
