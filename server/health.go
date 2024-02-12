package server

import (
	"fmt"
	"net/http"
	"time"
)

func (s *server) handleHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		health := struct {
			Ok     bool  `json:"ok"`
			Uptime int64 `json:"uptime"`
		}{
			Ok:     true,
			Uptime: time.Now().Sub(s.cfg.ServerStartTime).Milliseconds(),
		}

		err := encode(w, r, http.StatusOK, health)
		if err != nil {
			s.logger.Error(fmt.Errorf("health check encode: %w", err))
			http.Error(w, ErrPublicInternalError.Error(), http.StatusInternalServerError)
		}
	}
}
