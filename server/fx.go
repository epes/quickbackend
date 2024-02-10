package server

import (
	"context"
	"fmt"
	"github.com/epes/quickbackend/api"
	"github.com/epes/quickbackend/config"
	"github.com/epes/quickbackend/logger"

	"go.uber.org/fx"
)

func Invoke(
	cfg config.Config,
	lifecycle fx.Lifecycle,
	logger logger.Logger,
	api api.API,
) error {
	server, err := newServer(cfg, logger, api)
	if err != nil {
		return fmt.Errorf("creating new web server: %w", err)
	}

	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go server.Serve()

				return nil
			},
		},
	)

	return nil
}
