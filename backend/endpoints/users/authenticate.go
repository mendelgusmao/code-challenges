package users

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/context"
)

func init() {
	subrouter.HandleFunc("/authenticate", authenticateUser).Methods("POST")
}

func authenticateUser(w http.ResponseWriter, r *http.Request) {
	db := context.Get(r, "db")
	userDAO := newDAO(db.(*sql.DB))

	credentials := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		log.Printf("authenticateUser: %s", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	user, err := userDAO.findByEmail(credentials.Email)

	if err != nil {
		log.Printf("authenticateUser: %s", err)
		http.Error(w, "", http.StatusForbidden)
		return
	}

	if user.authenticate(credentials.Password) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user.filtered())
		return
	}

	w.WriteHeader(http.StatusForbidden)
}
