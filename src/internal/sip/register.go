package sip

import (
	"context"
)

func (c *Client) Register(ctx context.Context) error {
	c.log.Info(
		"starting register",
		"server", c.opts.ServerHost,
		"user", c.opts.Username,
	)

	c.log.Debug(
		"sending REGISTER",
		"expires", c.opts.Expire,
	)

	req, err := c.buildRegisterRequest()
	if err != nil {
		return err
	}

	res, err := c.client.Do(ctx, req)
	if err != nil {
		return err
	}

	c.log.Info(
		"register response",
		"status", res.StatusCode,
		"reason", res.Reason,
	)

	<-ctx.Done()

	c.log.Info("shutdown received")

	return nil
}
