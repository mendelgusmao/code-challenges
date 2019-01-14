package signup

import (
	"net/http"

	"bitbucket.org/mendelgusmao/me_gu/frontend/router"
	"bitbucket.org/mendelgusmao/me_gu/frontend/templates"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	resty "gopkg.in/resty.v1"
)

func init() {
	subrouter := router.Router.PathPrefix("/signup").Subrouter()

	subrouter.HandleFunc("", show).Methods("GET")
	subrouter.HandleFunc("", signup).Methods("POST")
}

func show(w http.ResponseWriter, r *http.Request) {
	templates.
		NewRenderer("signup/signup").
		Do(w, r, nil)
}

func signup(w http.ResponseWriter, r *http.Request) {
	session := context.Get(r, "session").(*sessions.Session)

	userData := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	response, err := resty.R().
		SetResult(map[string]interface{}{}).
		SetBody(userData).
		Put("/users")

	if err != nil {
		session.AddFlash("There was an error creating your account.")
	} else {
		switch response.StatusCode() {
		case http.StatusCreated:
			user := response.Result().(*map[string]interface{})

			session.Values["id"] = (*user)["id"]
			session.Values["email"] = (*user)["email"]
			session.Values["valid"] = true

			session.AddFlash("Account created successfully. Complete your profile.")
			session.Save(r, w)

			http.Redirect(w, r, "/profile/edit", http.StatusSeeOther)
			return

		case http.StatusConflict:
			session.AddFlash("Couldn't create account: e-mail already in use.")

		default:
			session.AddFlash("There was an error creating your account.")
		}
	}

	session.Save(r, w)

	templates.
		NewRenderer("signup/signup").
		Do(w, r, nil)
}
