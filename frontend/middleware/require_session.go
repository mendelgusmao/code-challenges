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
			session.AddFlash("You should log in.")
			session.Save(r, w)

			http.Redirect(w, r, "/session/new", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}
