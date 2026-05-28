package app

import (
	"context"
	"log/slog"

	"github.com/AsTR0I/sipcV2/internal/infra/config"
	"github.com/AsTR0I/sipcV2/internal/transport/sip"
)

type App struct {
	sip *sip.Client
	log *slog.Logger
}

func NewApp(
	log *slog.Logger,
	sipClient *sip.Client,
) *App {
	return &App{
		sip: sipClient,
		log: log,
	}
}

func (app *App) Options(
	ctx context.Context,
	cfg config.SIPConfig,
) error {
	return app.sip.Options(ctx, cfg)
}
