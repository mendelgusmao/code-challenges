package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(time.Now().Format(time.Stamp), r.RemoteAddr, r.Method, r.RequestURI)

		next.ServeHTTP(w, r)
	})
}
