package server

import (
	"fmt"
	"github.com/epes/quickbackend/api"
	"github.com/epes/quickbackend/config"
	"github.com/epes/quickbackend/logger"
	"github.com/epes/quickbackend/middleware"
	"net/http"
)

type server struct {
	cfg    config.Config
	logger logger.Logger
	api    api.API
}

func newServer(cfg config.Config, logger logger.Logger, api api.API) (*server, error) {
	s := &server{
		cfg:    cfg,
		logger: logger,
		api:    api,
	}

	return s, nil
}

func (s *server) Serve() {
	mux := http.NewServeMux()

	s.logger.Debug(fmt.Sprintf("running web server on :%d", s.cfg.Port))

	middlewareMux := middleware.Apply(
		mux,
		middleware.Header("Content-Type", "application/json"),
		middleware.Logger(s.logger),
	)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", s.cfg.Port), middlewareMux); err != nil {
		s.logger.Fatalw("web server", "err", err)
	}
}
