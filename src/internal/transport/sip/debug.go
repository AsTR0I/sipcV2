package sip

import (
	"fmt"
	"strings"

	"github.com/emiago/sipgo/sip"
	"github.com/fatih/color"
)

var (
	cEnvelope = color.New(color.FgHiBlack)

	cRequest = color.New(color.FgCyan)

	c1xx = color.New(color.FgYellow)
	c2xx = color.New(color.FgGreen)
	c3xx = color.New(color.FgGreen)
	c4xx = color.New(color.FgRed)
	c5xx = color.New(color.FgRed)
	c6xx = color.New(color.FgRed)
)

func DumpRequest(req *sip.Request, src, dst string) string {
	col := cRequest

	var b strings.Builder

	b.WriteString(cEnvelope.Sprintf("%s -> %s\n", src, dst))

	b.WriteString(col.Sprintf(buildRequest(req)))

	return b.String()
}

func DumpResponse(res *sip.Response, src, dst string) string {
	col := statusColor(int(res.StatusCode))

	var b strings.Builder

	b.WriteString(cEnvelope.Sprintf("%s <- %s\n", src, dst))

	// ВСЁ тело пакета одним цветом
	b.WriteString(col.Sprintf(buildResponse(res)))

	return b.String()
}

func buildRequest(req *sip.Request) string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf(
		"%s %s SIP/2.0\n",
		req.Method,
		uriString(req.Recipient),
	))

	for _, h := range req.Headers() {
		b.WriteString(fmt.Sprintf("%s: %s\n", h.Name(), h.Value()))
	}

	if len(req.Body()) > 0 {
		b.WriteString("\n")
		b.Write(req.Body())
	}

	return b.String()
}

func buildResponse(res *sip.Response) string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf(
		"SIP/2.0 %d %s\n",
		res.StatusCode,
		res.Reason,
	))

	for _, h := range res.Headers() {
		b.WriteString(fmt.Sprintf("%s: %s\n", h.Name(), h.Value()))
	}

	if len(res.Body()) > 0 {
		b.WriteString("\n")
		b.Write(res.Body())
	}

	return b.String()
}

func statusColor(code int) *color.Color {
	switch {
	case code >= 100 && code < 200:
		return c1xx
	case code >= 200 && code < 300:
		return c2xx
	case code >= 300 && code < 400:
		return c3xx
	case code >= 400 && code < 500:
		return c4xx
	case code >= 500 && code < 600:
		return c5xx
	case code >= 600 && code < 700:
		return c6xx
	default:
		return color.New(color.FgWhite)
	}
}

func uriString(u sip.Uri) string {
	var b strings.Builder

	b.WriteString(u.Scheme + ":")

	if u.Wildcard {
		return "*"
	}

	if u.User != "" {
		b.WriteString(u.User + "@")
	}

	b.WriteString(u.Host)

	if u.Port != 0 {
		b.WriteString(fmt.Sprintf(":%d", u.Port))
	}

	return b.String()
}
