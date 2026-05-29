package sip

import (
	"context"
	"net"
	"strconv"

	"github.com/AsTR0I/sipcV2/internal/domain"
	"github.com/emiago/sipgo"
	"github.com/emiago/sipgo/sip"
	"github.com/google/uuid"
)

type Transport interface {
	Do(ctx context.Context, req *sip.Request) (*sip.Response, error)
	DoDigestAuth(ctx context.Context, req *sip.Request, res *sip.Response, auth sipgo.DigestAuth) (*sip.Response, error)
}

func (c *Client) Register(ctx context.Context, cfg domain.RegisterRequest) error {
	host, port, err := parseHostPort(cfg.ServerHost)
	if err != nil {
		return err
	}

	dest := cfg.ServerHost
	if cfg.Proxy != "" {
		dest = cfg.Proxy
	}

	dest = net.JoinHostPort(dest, strconv.Itoa(port))

	req := buildRegisterReq(cfg, host, port, dest)

	c.dump(req, nil)

	res, err := c.client.Do(ctx, req)
	if err != nil {
		return err
	}

	c.dump(req, res)

	if needsAuth(res) {
		req2 := req.Clone()

		auth := sipgo.DigestAuth{
			Username: cfg.Username,
			Password: cfg.Password,
		}

		res2, err := c.client.DoDigestAuth(ctx, req2, res, auth)
		if err != nil {
			return err
		}

		c.dump(req2, res2)
		res = res2
	}

	return nil
}
func buildRegisterReq(cfg domain.RegisterRequest, host string, port int, destination string) *sip.Request {
	recipient := sip.Uri{
		Scheme: "sip",
		Host:   host,
		Port:   port,
	}

	req := sip.NewRequest(sip.REGISTER, recipient)
	req.SetDestination(destination)

	viaHeader := &sip.ViaHeader{
		ProtocolName:    "SIP",
		ProtocolVersion: "2.0",
		Transport:       "UDP",
		Host:            host,
		Port:            port,
		Params:          sip.NewParams(),
	}

	viaHeader.Params.Add("branch", sip.GenerateBranchN(10))
	req.AppendHeader(viaHeader)

	maxForwardsHeader := sip.MaxForwardsHeader(70)
	req.AppendHeader(&maxForwardsHeader)

	from := &sip.FromHeader{
		Address: sip.Uri{
			Scheme: "sip",
			User:   cfg.From,
			Host:   host,
			Port:   port,
		},
		Params: sip.NewParams(),
	}
	from.Params.Add("tag", uuid.NewString())
	req.AppendHeader(from)

	to := &sip.ToHeader{
		Address: sip.Uri{
			Scheme: "sip",
			User:   cfg.To,
			Host:   host,
			Port:   port,
		},
	}
	req.AppendHeader(to)

	callID := sip.CallIDHeader(uuid.NewString())
	req.AppendHeader(&callID)

	req.AppendHeader(&sip.CSeqHeader{
		SeqNo:      1,
		MethodName: sip.REGISTER,
	})

	req.AppendHeader(&sip.ContactHeader{
		Address: sip.Uri{
			Scheme: "sip",
			User:   cfg.From,
			Host:   host,
			Port:   port,
		},
	})

	req.AppendHeader(sip.NewHeader("Allow", "REGISTER"))

	expires := cfg.Expires
	if expires <= 0 {
		expires = 360
	}
	req.AppendHeader(sip.NewHeader("Expires", strconv.Itoa(expires)))

	req.AppendHeader(sip.NewHeader("User-Agent", cfg.UserAgent))

	req.AppendHeader(sip.NewHeader("Event", "keep-alive"))

	return req
}
