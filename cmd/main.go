package main

import (
	"github.com/epes/quickbackend/api"
	"github.com/epes/quickbackend/config"
	"github.com/epes/quickbackend/environment"
	"github.com/epes/quickbackend/logger"
	"github.com/epes/quickbackend/server"
	"github.com/epes/quickbackend/store/kvstore"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			environment.New,
			config.New,
			logger.New,
			kvstore.NewStringKV,
			kvstore.NewIntKV,
			api.New,
		),
		fx.Invoke(server.Invoke),
	).Run()
}
