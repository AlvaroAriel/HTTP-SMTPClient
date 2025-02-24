package middleware

import (
	"net/http"
	"os"
)

func CorsMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Access-Control-Allow-Origin", os.Getenv("ALLOWED_HOSTS"))
		w.Header().Add("Access-Control-Allow-Headers", "origin, content-type")
		w.Header().Add("Access-Control-Allow-Methods", "POST")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)

	})
}
