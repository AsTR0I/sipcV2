package sip

import (
	"context"

	"github.com/emiago/sipgo"
	"github.com/emiago/sipgo/sip"
)

type Transport interface {
	Do(ctx context.Context, req *sip.Request) (*sip.Response, error)
	DoDigestAuth(ctx context.Context, req *sip.Request, res *sip.Response, auth sipgo.DigestAuth) (*sip.Response, error)
}
