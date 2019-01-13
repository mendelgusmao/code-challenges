package middleware

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

func RequireSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := context.Get(r, "session").(*sessions.Session)

		if valid, ok := session.Values["valid"]; !ok || (ok && !valid.(bool)) {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}
