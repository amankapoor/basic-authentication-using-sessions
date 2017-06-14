package middleware

import (
	"net/http"

	"log"
)

func Middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Print("Middleware starts")
		next.ServeHTTP(w, r)
		log.Print("Middleware ends")
	}
	return http.HandlerFunc(fn)
}
