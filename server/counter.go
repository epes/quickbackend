package server

import (
	"context"
	"fmt"
	"net/http"
)

func (s *server) handleCounterGet() http.HandlerFunc {
	type response struct {
		Count int `json:"count"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		counter, err := s.api.Counter(r.Context())
		if err != nil {
			s.logger.Error(fmt.Errorf("counter get from api: %w", err))
			http.Error(w, ErrPublicInternalError.Error(), http.StatusInternalServerError)
			return
		}

		err = encode(w, r, http.StatusOK, response{
			Count: counter,
		})
		if err != nil {
			s.logger.Error(fmt.Errorf("counter get encode: %w", err))
			http.Error(w, ErrPublicInternalError.Error(), http.StatusInternalServerError)
			return
		}
	}
}

type CounterPostRequest struct {
	Addition int
}

func (r CounterPostRequest) Valid(ctx context.Context) map[string]string {
	if r.Addition <= 0 {
		return map[string]string{
			"Addition": "addition has to be a positive number",
		}
	}

	return nil
}

func (s *server) handleCounterPost() http.HandlerFunc {
	type response struct {
		Count int `json:"count"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		decoded, fields, err := decodeValid[CounterPostRequest](r)
		if err != nil {
			s.logger.Info(fmt.Errorf("counter post decode: %s: %v", err, fields))
			http.Error(w, ErrPublicInternalError.Error(), http.StatusInternalServerError)
			return
		}

		count, err := s.api.IncreaseCounter(r.Context(), decoded.Addition)
		if err != nil {
			s.logger.Error(fmt.Errorf("counter post from api: %w", err))
			http.Error(w, ErrPublicInternalError.Error(), http.StatusInternalServerError)
			return
		}

		err = encode(w, r, http.StatusOK, response{
			Count: count,
		})
		if err != nil {
			s.logger.Error(fmt.Errorf("counter post encode: %w", err))
			http.Error(w, ErrPublicInternalError.Error(), http.StatusInternalServerError)
		}
	}
}
