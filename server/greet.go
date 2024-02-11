package server

import (
	"context"
	"fmt"
	"net/http"
)

func (s *server) handleGreetGet() http.HandlerFunc {
	type response struct {
		Message string `json:"message"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")

		err := encode(w, r, http.StatusOK, response{
			Message: fmt.Sprintf("Hello %s", name),
		})

		if err != nil {
			s.logger.Error(fmt.Errorf("greet get encode: %w", err))
			http.Error(w, ErrPublicInternalError.Error(), http.StatusInternalServerError)
		}
	}
}

type GreetPostRequest struct {
	Name string
}

func (r GreetPostRequest) Valid(ctx context.Context) map[string]string {
	if r.Name == "" {
		return map[string]string{
			"Name": "empty name parameter",
		}
	}

	return nil
}

func (s *server) handleGreetPost() http.HandlerFunc {
	type response struct {
		Message string `json:"message"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		decoded, fields, err := decodeValid[GreetPostRequest](r)
		if err != nil {
			s.logger.Info(fmt.Errorf("greet post decode: %s: %v", err, fields))
			http.Error(w, ErrPublicInternalError.Error(), http.StatusInternalServerError)
			return
		}

		err = encode(w, r, http.StatusOK, response{
			Message: fmt.Sprintf("Hello %s", decoded.Name),
		})
		if err != nil {
			s.logger.Error(fmt.Errorf("greet post encode: %w", err))
			http.Error(w, ErrPublicInternalError.Error(), http.StatusInternalServerError)
		}
	}
}
