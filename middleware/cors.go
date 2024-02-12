package middleware

import (
	"github.com/epes/quickbackend/config"
	"github.com/epes/quickbackend/logger"
	"net/http"
	"strings"
)

func Cors(logger logger.Logger, config config.Config) Middleware {
	origins := strings.Join(config.AccessControlAllowOrigin, ", ")
	methods := strings.Join(config.AccessControlAllowMethods, ", ")
	headers := strings.Join(config.AccessControlAllowHeaders, ", ")

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", origins)
			w.Header().Set("Access-Control-Allow-Methods", methods)
			w.Header().Set("Access-Control-Allow-Headers", headers)

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			} else {
				h.ServeHTTP(w, r)
			}
		})
	}
}
