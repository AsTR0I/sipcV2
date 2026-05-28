package sip

import (
	"log/slog"

	"github.com/emiago/sipgo"
)

type Client struct {
	log    *slog.Logger
	ua     *sipgo.UserAgent
	client *sipgo.Client
}

func NewClient(
	log *slog.Logger,
	userAgent string,
) (*Client, error) {

	ua, err := sipgo.NewUA(
		sipgo.WithUserAgent(userAgent),
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

	return &Client{
		log:    log,
		ua:     ua,
		client: client,
	}, nil
}
