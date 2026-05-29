package app

import (
	"context"

	"github.com/AsTR0I/sipcV2/internal/domain"
)

type SIPClient interface {
	Register(ctx context.Context, req domain.RegisterRequest) error
	Options(ctx context.Context, req domain.OptionsRequest) error
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
