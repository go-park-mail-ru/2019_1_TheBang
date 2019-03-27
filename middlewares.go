package main

import (
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2019_1_TheBang/config"
)

func accessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		config.Logger.Infow("AppLog",
			"method=", r.Method,
			"path=", r.URL.Path,
			"remAddr=", r.RemoteAddr,
			"service=", time.Since(start))
	})
}
