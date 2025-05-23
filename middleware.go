package main

import (
	"net/http"
)

func useCorsMiddleware(h http.Handler) http.HandlerFunc {
	var logger = CreateLogger("Route")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.TRACE(r.Method, r.URL.Path)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")
		h.ServeHTTP(w, r)
	})
}
