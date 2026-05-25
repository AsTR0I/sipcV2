package sip

import (
	"log/slog"

	"github.com/emiago/sipgo"
)

type Options struct {
	ServerHost    string
	ServerPort    int
	From          string
	To            string
	Username      string
	Password      string
	Realm         string
	Expire        int
	UserAgentName string
}

type Client struct {
	log  *slog.Logger
	opts Options

	ua     *sipgo.UserAgent
	client *sipgo.Client
}

func NewClient(
	log *slog.Logger,
	opts Options,

) (*Client, error) {
	ua, err := sipgo.NewUA(
		sipgo.WithUserAgent(opts.UserAgentName),
	)
	if err != nil {
		return nil, err
	}

	uaClient, err := sipgo.NewClient(
		ua,
		sipgo.WithClientLogger(log),
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		log:    log,
		opts:   opts,
		ua:     ua,
		client: uaClient,
	}, nil
}
