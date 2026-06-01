package app

import (
	"context"

	"github.com/AsTR0I/sipcV2/internal/domain"
)

type SIPClient interface {
	Register(ctx context.Context, cfg domain.RegisterRequest) error
	Options(ctx context.Context, cfg domain.OptionsRequest) error
	Invite(ctx context.Context, cfg domain.InviteRequest) error
}

type App struct {
	sip SIPClient
}

func NewApp(sipClient SIPClient) *App {
	return &App{
		sip: sipClient,
	}
}

func (a *App) Register(
	ctx context.Context,
	req domain.RegisterRequest,
) error {
	return a.sip.Register(ctx, req)
}

func (a *App) Options(
	ctx context.Context,
	req domain.OptionsRequest,
) error {
	return a.sip.Options(ctx, req)
}

func (a *App) Invite(
	ctx context.Context,
	req domain.InviteRequest,
) error {
	return a.sip.Invite(ctx, req)
}
