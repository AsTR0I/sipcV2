package sip

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/emiago/sipgo"
	"github.com/emiago/sipgo/sip"
)

type Client struct {
	log    *slog.Logger
	ua     *sipgo.UserAgent
	client *sipgo.Client

	localUDP net.PacketConn
}

func NewClient(
	log *slog.Logger,
	userAgent string,
	userPort string,
) (*Client, error) {

	ua, err := sipgo.NewUA(
		sipgo.WithUserAgent(userAgent),
	)
	if err != nil {
		return nil, err
	}

	port := "5060"
	if userPort != "" {
		port = userPort
	}
	localAddr := fmt.Sprintf("0.0.0.0:%s", port)

	client, err := sipgo.NewClient(
		ua,
		sipgo.WithClientLogger(log),
		sipgo.WithClientAddr(localAddr),
	)
	if err != nil {
		return nil, err
	}

	c := &Client{
		log:    log,
		ua:     ua,
		client: client,
	}

	return c, nil
}

func (c *Client) dump(req *sip.Request, res *sip.Response) {
	if req != nil {
		fmt.Println(DumpRequest(req, req.Destination(), req.Laddr.String()))
	}

	if res != nil {
		fmt.Println(DumpResponse(res, req.Destination(), req.Laddr.String()))
	}
}
