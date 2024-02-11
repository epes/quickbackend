package server

import (
	"fmt"
	"net/http"
)

func (s *server) handleHealth() http.HandlerFunc {
	ok := struct {
		Ok bool `json:"ok"`
	}{
		Ok: true,
	}

	return func(w http.ResponseWriter, r *http.Request) {
		err := encode(w, r, http.StatusOK, ok)
		if err != nil {
			s.logger.Error(fmt.Errorf("health check encode: %w", err))
			http.Error(w, ErrPublicInternalError.Error(), http.StatusInternalServerError)
		}
	}
}
