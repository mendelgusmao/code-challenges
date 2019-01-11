package users

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

func init() {
	subrouter.HandleFunc("/password-reset", createPasswordResetRequest).Methods("POST")
	subrouter.HandleFunc("/password-reset/token/{token}", checkPasswordResetTokenValidity).Methods("GET")
	subrouter.HandleFunc("/password-reset/password", updatePassword).Methods("POST")
}

func createPasswordResetRequest(w http.ResponseWriter, r *http.Request) {
	db := context.Get(r, "db")
	userDAO := newDAO(db.(*sql.DB))

	credentials := struct {
		Email string `json:"email"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		log.Printf("createPasswordResetRequest: %s", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	user, err := userDAO.findByEmail(credentials.Email)

	if err != nil {
		log.Printf("createPasswordResetRequest: %s", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	token, err := uuid.NewV4()

	if err != nil {
		log.Printf("createPasswordResetRequest: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	if err := userDAO.updatePasswordResetToken(user.Email, token.String()); err != nil {
		log.Printf("createPasswordResetRequest: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func checkPasswordResetTokenValidity(w http.ResponseWriter, r *http.Request) {
	db := context.Get(r, "db")
	userDAO := newDAO(db.(*sql.DB))

	token := mux.Vars(r)["token"]
	user, err := userDAO.findByToken(token)

	if err != nil {
		log.Printf("checkPasswordResetTokenValidity: %s", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	if !user.validPasswordResetToken() {
		log.Printf("checkPasswordResetTokenValidity: expired token")
		http.Error(w, "", http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func updatePassword(w http.ResponseWriter, r *http.Request) {
	db := context.Get(r, "db")
	userDAO := newDAO(db.(*sql.DB))

	credentials := struct {
		Token    string `json:"token"`
		Password string `json:"password"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		log.Printf("createPasswordResetRequest: %s", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	user, err := userDAO.findByToken(credentials.Token)

	if err != nil || user == nil {
		log.Printf("updatePassword: %s", err)
		http.Error(w, "", http.StatusForbidden)
		return
	}

	if !user.validPasswordResetToken() {
		log.Printf("updatePassword: expired token")
		http.Error(w, "", http.StatusForbidden)
		return
	}

	userRequest := UserRequest(*user)
	userRequest.Password = credentials.Password

	if err := userRequest.encryptPassword(); err != nil {
		log.Printf("updatePassword: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	userRequest.PasswordResetToken = nil
	userRequest.PasswordResetTokenExpiration = nil
	user.apply(&userRequest)

	if err := userDAO.update(user); err != nil {
		log.Printf("updatePassword: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
