package server

import (
	"fmt"
	"net/http"
	"time"
)

type HealthResponse struct {
	Ok     bool  `json:"ok"`
	Uptime int64 `json:"uptime"`
}

func (s *server) handleHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ok := true
		healthErr := s.api.Health(r.Context())
		if healthErr != nil {
			s.logger.Errorw("health check error", healthErr)
			ok = false
		}

		health := HealthResponse{
			Ok:     ok,
			Uptime: time.Now().Sub(s.cfg.ServerStartTime).Milliseconds(),
		}

		err := encode(w, r, http.StatusOK, health)
		if err != nil {
			s.logger.Error(fmt.Errorf("health check encode: %w", err))
			http.Error(w, ErrPublicInternalError.Error(), http.StatusInternalServerError)
		}
	}
}
