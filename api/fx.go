package api

import (
	"context"
	"github.com/epes/quickbackend/config"
	"github.com/epes/quickbackend/logger"
	"github.com/epes/quickbackend/store/kvstore"
)

type API interface {
	Health(ctx context.Context) error

	Counter(ctx context.Context) (int, error)
	IncreaseCounter(ctx context.Context, addition int) (int, error)
}

type api struct {
	cfg    config.Config
	logger logger.Logger
	intKV  kvstore.KVStore[int]
}

func New(
	cfg config.Config,
	logger logger.Logger,
	intKV kvstore.KVStore[int],
) (API, error) {
	a := &api{
		cfg:    cfg,
		logger: logger,
		intKV:  intKV,
	}

	return a, nil
}
