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

var subrouter = router.Router.PathPrefix("/users").Subrouter()

func init() {
	subrouter.HandleFunc("/{id:[0-9]+}", getUser).Methods("GET")
	subrouter.HandleFunc("", createUser).Methods("PUT")
	subrouter.HandleFunc("/{id:[0-9]+}", updateUser).Methods("PATCH")
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

	json.NewEncoder(w).Encode(user.filtered())
}

func createUser(w http.ResponseWriter, r *http.Request) {
	db := context.Get(r, "db")
	userDAO := newDAO(db.(*sql.DB))

	var userRequest UserRequest

	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		log.Printf("createUser: %s", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if u, _ := userDAO.findByEmail(userRequest.Email); u != nil {
		log.Printf("createUser: email in use")
		http.Error(w, "", http.StatusConflict)
		return
	}

	if errs := userRequest.validate(); errs != nil {
		// TODO
	}

	if err := userRequest.encryptPassword(); err != nil {
		log.Printf("createUser: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	user := User(userRequest)

	if err := userDAO.create(&user); err != nil {
		log.Printf("createUser: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user.filtered())
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	db := context.Get(r, "db")
	userDAO := newDAO(db.(*sql.DB))

	var userRequest UserRequest

	id, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	user, err := userDAO.findByID(id)

	if err != nil {
		log.Printf("getUser: %s", err)

		if err == sql.ErrNoRows {
			http.Error(w, "", http.StatusNotFound)
			return
		}

		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		log.Printf("updateUser: %s", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if errs := userRequest.validate(); errs != nil {
		// TODO
	}

	if err := userRequest.encryptPassword(); err != nil {
		log.Printf("createUser: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	user.apply(&userRequest)

	if err := userDAO.update(user); err != nil {
		log.Printf("updateUser: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
