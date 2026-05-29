package sip

import (
	"errors"
	"fmt"
	"log/slog"
	"net"
	"strconv"

	"github.com/AsTR0I/sipcV2/internal/infra/config"
	"github.com/emiago/sipgo"
)

type Client struct {
	log    *slog.Logger
	ua     *sipgo.UserAgent
	client *sipgo.Client

	localUDP net.PacketConn
}

func NewClient(
	log *slog.Logger,
	cfg config.SIPConfig,
) (*Client, error) {

	ua, err := sipgo.NewUA(
		sipgo.WithUserAgent(cfg.UserAgent),
	)
	if err != nil {
		return nil, err
	}

	client, err := sipgo.NewClient(
		ua,
		sipgo.WithClientLogger(log),
	)
	if err != nil {
		return nil, err
	}

	c := &Client{
		log:    log,
		ua:     ua,
		client: client,
	}

	if cfg.UserPort != "" {
		intUserPort, err := strconv.Atoi(cfg.UserPort)
		if err != nil {
			return nil, err
		}
		if intUserPort < 1 || intUserPort > 65535 {
			return nil, errors.New("user-port cant not be < 1 or > 65535")
		}

		addr := fmt.Sprintf("0.0.0.0:%s", cfg.UserPort)

		conn, err := net.ListenPacket("udp", addr)
		if err != nil {
			return nil, err
		}

		c.localUDP = conn

		go func() {
			if err := ua.TransportLayer().ServeUDP(conn); err != nil {
				log.Error("udp transport stopped", slog.Any("err", err))
			}
		}()
	}

	return c, nil
}
