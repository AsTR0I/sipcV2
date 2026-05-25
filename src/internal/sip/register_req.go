package sip

import (
	"fmt"

	"github.com/emiago/sipgo/sip"
	"github.com/google/uuid"
)

func (c *Client) buildRegisterRequest() (*sip.Request, error) {
	recipient := sip.Uri{
		Scheme: "sip",
		Host:   c.opts.ServerHost,
		Port:   c.opts.ServerPort,
	}

	req := sip.NewRequest(
		sip.REGISTER,
		recipient,
	)

	fromAddress := sip.Uri{
		Scheme: "sip",
		User:   c.opts.From,
		Host:   c.opts.Realm,
	}

	toAddress := sip.Uri{
		Scheme: "sip",
		User:   c.opts.To,
		Host:   c.opts.Realm,
	}

	req.AppendHeader(
		sip.NewHeader("Max-Forwards", "70"),
	)

	params := sip.NewParams()
	params.Add("tag", sip.GenerateTagN(16))
	req.AppendHeader(
		&sip.FromHeader{
			DisplayName: "",
			Address:     fromAddress,
			Params:      params,
		},
	)

	req.AppendHeader(
		&sip.ToHeader{
			DisplayName: "",
			Address:     toAddress,
			Params:      sip.NewParams(),
		},
	)

	req.AppendHeader(
		&sip.ContactHeader{
			Address: fromAddress,
		},
	)

	uuid, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	callID := sip.CallIDHeader(uuid.String())
	req.AppendHeader(&callID)

	cseq := sip.CSeqHeader{
		SeqNo:      1,
		MethodName: sip.REGISTER,
	}

	req.AppendHeader(&cseq)

	req.AppendHeader(
		sip.NewHeader(
			"Expires",
			fmt.Sprintf("%d", c.opts.Expire),
		),
	)

	req.AppendHeader(
		sip.NewHeader(
			"User-Agent",
			c.opts.UserAgentName,
		),
	)

	return req, nil
}
