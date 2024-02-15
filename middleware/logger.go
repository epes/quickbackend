package middleware

import (
	"github.com/epes/quickbackend/logger"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func Logger(logger logger.Logger) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			requestUuid, _ := uuid.NewRandom()
			requestId := requestUuid.String()

			logger.Infow(
				"http request",
				"request_id", requestId,
				"url", r.URL.Path,
				"method", r.Method,
			)

			h.ServeHTTP(w, r)

			logger.Infow(
				"http response",
				"request_id", requestId,
				"timing", time.Now().Sub(start).String(),
			)
		})
	}
}
