package sip

// РАЗОБРАТЬСЯ С ПРОБЛЕМОЙ ОТМЕНЫ 	invite при ctx.done
import (
	"context"
	"net"
	"strconv"

	"github.com/AsTR0I/sipcV2/internal/domain"
	"github.com/emiago/diago"
	"github.com/emiago/sipgo/sip"
)

func (c *Client) Invite(ctx context.Context, cfg domain.InviteRequest) error {
	host, port, err := parseHostPort(cfg.ServerHost)
	if err != nil {
		return err
	}

	dest := cfg.ServerHost
	if cfg.Proxy != "" {
		dest = cfg.Proxy
	}

	dest = net.JoinHostPort(dest, strconv.Itoa(port))

	recipient := sip.Uri{
		Scheme: "sip",
		User:   cfg.To,
		Host:   host,
		Port:   port,
	}

	dg := diago.NewDiago(
		c.ua,
		diago.WithClient(c.client),
	)

	opts := diago.NewDialogOptions{
		Transport: "udp",
	}

	c.log.Info("1")

	d, err := dg.NewDialog(recipient, opts)
	if err != nil {
		return err
	}

	req := d.InviteRequest

	fromParams := sip.NewParams()

	req.RemoveHeader("From")
	req.AppendHeader(&sip.FromHeader{
		Address: sip.Uri{
			Scheme: "sip",
			User:   cfg.From,
			Host:   host,
		},
		Params: fromParams.Add(
			"tag",
			sip.GenerateTagN(16),
		),
	})

	req.RemoveHeader("To")
	req.ReplaceHeader(&sip.ToHeader{
		Address: sip.Uri{
			Scheme: "sip",
			User:   cfg.To,
			Host:   host,
			Port:   port,
		},
	})

	req.RemoveHeader("Contact")
	req.ReplaceHeader(&sip.ContactHeader{
		Address: sip.Uri{
			Scheme: "sip",
			User:   cfg.From,
			Host:   host,
			Port:   port,
		},
	})

	c.dump(d.InviteRequest, nil)

	if err := d.Invite(ctx, diago.InviteClientOptions{
		Username: cfg.Username,
		Password: cfg.Password,
	}); err != nil {
		return err
	}

	if err := d.Ack(ctx); err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		d.Hangup(context.Background())
		d.Close()
		return ctx.Err()
	case <-d.Context().Done():
		d.Hangup(d.Context())
		d.Close()
		return d.Context().Err()
	}
}
