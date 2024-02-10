package api

import (
	"context"
	"github.com/epes/quickbackend/config"
	"github.com/epes/quickbackend/logger"
)

type API interface {
	Health(ctx context.Context) error
}

type api struct {
	cfg    config.Config
	logger logger.Logger
}

func New(
	cfg config.Config,
	logger logger.Logger,
) (API, error) {
	a := &api{
		cfg:    cfg,
		logger: logger,
	}

	return a, nil
}
