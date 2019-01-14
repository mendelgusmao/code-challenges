package templates

import (
	"html/template"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

func funcMap(w http.ResponseWriter, r *http.Request) template.FuncMap {
	return template.FuncMap{
		"loggedIn": func() bool {
			session := context.Get(r, "session").(*sessions.Session)
			valid, ok := session.Values["valid"]

			return ok && valid.(bool)
		},
		"email": func() string {
			session := context.Get(r, "session").(*sessions.Session)
			valid, ok := session.Values["valid"]

			if ok && valid.(bool) {
				return session.Values["email"].(string)
			}

			return ""
		},
		"flashes": func() []interface{} {
			session := context.Get(r, "session").(*sessions.Session)
			flashes := session.Flashes()
			session.Save(r, w)

			return flashes
		},
	}
}
