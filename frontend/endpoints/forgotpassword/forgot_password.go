package forgotpassword

import (
	"fmt"
	"log"
	"net/http"

	"bitbucket.org/mendelgusmao/me_gu/frontend/router"
	"bitbucket.org/mendelgusmao/me_gu/frontend/templates"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	resty "gopkg.in/resty.v1"
)

func init() {
	subrouter := router.Router.PathPrefix("/forgot-password").Subrouter()

	subrouter.HandleFunc("", show).Methods("GET")
	subrouter.HandleFunc("", createResetToken).Methods("POST")
	subrouter.HandleFunc("/{token}", newPasswordForm).Methods("GET")
	subrouter.HandleFunc("/{token}", setNewPassword).Methods("POST")
}

func show(w http.ResponseWriter, r *http.Request) {
	templates.
		NewRenderer("forgotpassword/show").
		Do(w, r, nil)
}

func createResetToken(w http.ResponseWriter, r *http.Request) {
	session := context.Get(r, "session").(*sessions.Session)

	userData := struct {
		Email string `json:"email"`
	}{
		Email: r.FormValue("email"),
	}

	response, err := resty.R().
		SetBody(userData).
		Post("/users/password-reset")

	if err != nil {
		log.Printf("forgotpassword.createResetToken: %s", err)
		session.AddFlash("There was an error creating a password reset token.")
	} else {
		switch response.StatusCode() {
		case http.StatusCreated:
			session.AddFlash("Password reset token successfully created. Check your inbox.")
			session.Save(r, w)

			http.Redirect(w, r, "/forgot-password", http.StatusSeeOther)
			return

		default:
			log.Printf("forgotpassword.createResetToken: backend returned %d", response.StatusCode())
			session.AddFlash("There was an error creating a password reset token.")
		}
	}

	session.Save(r, w)

	templates.
		NewRenderer("forgotpassword/show").
		Do(w, r, nil)
}

func newPasswordForm(w http.ResponseWriter, r *http.Request) {
	session := context.Get(r, "session").(*sessions.Session)

	token := mux.Vars(r)["token"]
	uri := fmt.Sprintf("/users/password-reset/token/%s", token)

	response, err := resty.R().
		Get(uri)

	if err != nil {
		log.Printf("forgotpassword.newPasswordForm: %s", err)
		session.AddFlash("There was an error checking validity of the token.")
	} else {
		switch response.StatusCode() {
		case http.StatusNoContent:
			templateData := struct {
				Token string
			}{
				Token: token,
			}

			templates.
				NewRenderer("forgotpassword/new_password").
				Do(w, r, templateData)
			return

		default:
			log.Printf("forgotpassword.newPasswordForm: backend returned %d", response.StatusCode())
			session.AddFlash("There was an error checking validity of the token.")
		}
	}

	session.Save(r, w)

	templates.
		NewRenderer("forgotpassword/show").
		Do(w, r, nil)
}

func setNewPassword(w http.ResponseWriter, r *http.Request) {
	session := context.Get(r, "session").(*sessions.Session)

	userData := struct {
		Token    string `json:"token"`
		Password string `json:"password"`
	}{
		Token:    mux.Vars(r)["token"],
		Password: r.FormValue("password"),
	}

	response, err := resty.R().
		SetBody(userData).
		Post("/users/password-reset/password")

	if err != nil {
		log.Printf("forgotpassword.createResetToken: %s", err)
		session.AddFlash("There was an error changing your password.")
	} else {
		switch response.StatusCode() {
		case http.StatusNoContent:
			session.AddFlash("Password successfully changed. You should login now.")
			session.Save(r, w)

			http.Redirect(w, r, "/session/new", http.StatusSeeOther)
			return

		case http.StatusForbidden:
			session.AddFlash("There was an error changing your password. Token expired.")

		default:
			log.Printf("forgotpassword.createResetToken: backend returned %d", response.StatusCode())
			session.AddFlash("There was an error changing your password.")
		}
	}

	session.Save(r, w)

	templates.
		NewRenderer("forgotpassword/show").
		Do(w, r, nil)
}
