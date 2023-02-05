package rest

import (
	"net/http"

	"gateway/configs"
)

func SetupGlobalMiddleware(cfg *configs.Config, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	})
}
