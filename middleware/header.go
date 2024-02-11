package middleware

import (
	"net/http"
)

func Header(key string, value string) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(key, value)

			h.ServeHTTP(w, r)
		})
	}
}
