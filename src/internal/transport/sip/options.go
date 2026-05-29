package sip

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"github.com/AsTR0I/sipcV2/internal/domain"
	"github.com/emiago/sipgo/sip"
	"github.com/google/uuid"
)

func (c *Client) Options(
	ctx context.Context,
	cfg domain.OptionsRequest,
) error {
	host := cfg.ServerHost
	port := "5060"

	if parsedHost, parsedPort, err := net.SplitHostPort(cfg.ServerHost); err == nil {
		host = parsedHost
		port = parsedPort
	}

	portInt, err := strconv.Atoi(port)
	if err != nil {
		return err
	}

	destination := net.JoinHostPort(host, port)

	if cfg.Proxy != "" {
		destination = cfg.Proxy
	}

	recipient := sip.Uri{
		Scheme: "sip",
		Host:   host,
		Port:   portInt,
	}

	req := sip.NewRequest(
		sip.OPTIONS,
		recipient,
	)

	req.SetDestination(destination)

	viaHeader := &sip.ViaHeader{
		ProtocolName:    "SIP",
		ProtocolVersion: "2.0",
		Transport:       "UDP",
		Params:          sip.NewParams(),
	}

	viaHeader.Params.Add(
		"branch",
		sip.GenerateBranch(),
	)

	req.AppendHeader(viaHeader)

	maxForwardsHeader := sip.MaxForwardsHeader(70)
	req.AppendHeader(&maxForwardsHeader)

	fromHeader := &sip.FromHeader{
		Address: sip.Uri{
			Scheme: "sip",
			User:   cfg.From,
			Host:   host,
			Port:   portInt,
		},
		Params: sip.NewParams(),
	}
	fromHeader.Params.Add("tag", uuid.NewString())
	req.AppendHeader(fromHeader)

	toHeader := &sip.ToHeader{
		Address: sip.Uri{
			Scheme: "sip",
			User:   cfg.To,
			Host:   host,
			Port:   portInt,
		},
	}
	req.AppendHeader(toHeader)

	callID := sip.CallIDHeader(uuid.NewString())
	req.AppendHeader(&callID)

	cSeqHeader := &sip.CSeqHeader{
		SeqNo:      1,
		MethodName: sip.OPTIONS,
	}
	req.AppendHeader(cSeqHeader)

	allowHeader := sip.NewHeader("Allow", "OPTIONS")
	req.AppendHeader(allowHeader)

	req.AppendHeader(
		sip.NewHeader(
			"User-Agent",
			cfg.UserAgent,
		),
	)

	eventHeader := sip.NewHeader("Event", "keep-alive")
	req.AppendHeader(eventHeader)

	fmt.Println(DumpRequest(req, req.Laddr.String(), req.Destination()))

	res, err := c.client.Do(ctx, req)
	if err != nil {
		return err
	}

	fmt.Println(DumpResponse(res, req.Destination(), req.Laddr.String()))

	c.log.Info(
		"SIP OPTIONS completed",
		"status",
		fmt.Sprintf("%d %s", res.StatusCode, res.Reason),
	)

	return nil
}
