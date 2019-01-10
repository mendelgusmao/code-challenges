package users

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"bitbucket.org/mendelgusmao/me_gu/backend/router"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func init() {
	subrouter := router.Router.PathPrefix("/users").Subrouter()

	subrouter.HandleFunc("/{id:[0-9]+}", getUser).Methods("GET")
	subrouter.HandleFunc("", createUser).Methods("POST")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	db := context.Get(r, "db")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	userDAO := newDAO(db.(*sql.DB))
	user, err := userDAO.findByID(id)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "", http.StatusNotFound)
			return
		}

		log.Printf("getUser: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
}
