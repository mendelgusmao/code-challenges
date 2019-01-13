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
	}
}
