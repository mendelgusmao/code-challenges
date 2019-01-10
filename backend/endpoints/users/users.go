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
	id, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)

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
	db := context.Get(r, "db")

	userDAO := newDAO(db.(*sql.DB))
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("createUser: %s", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if errs := user.validate(); errs != nil {
		// TODO
	}

	if err := user.encryptPassword(); err != nil {
		log.Printf("createUser: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	if err := userDAO.create(&user); err != nil {
		log.Printf("createUser: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
