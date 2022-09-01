package middleware

import (
	"log"
	"net/http"
)

func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		log.Print("UTC [INFO] LOG:", r)
		next.ServeHTTP(w, r)
	})

}
