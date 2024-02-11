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
	cfg     config.Config
	logger  logger.Logger
	api     api.API
	handler http.Handler
}

func newServer(cfg config.Config, logger logger.Logger, api api.API) (*server, error) {
	s := &server{
		cfg:    cfg,
		logger: logger,
		api:    api,
	}

	s.handler = middleware.Apply(
		s.routedMux(),
		middleware.Header("Content-Type", "application/json"),
		middleware.Logger(s.logger),
	)

	return s, nil
}

func (s *server) Serve() {

	s.logger.Debug(fmt.Sprintf("server listening on :%d", s.cfg.Port))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", s.cfg.Port), s.handler); err != nil && err != http.ErrServerClosed {
		s.logger.Fatalw("web server", "err", err)
	}
}
