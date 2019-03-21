package users

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/mendelgusmao/supereasy/backend/endpoints/messages"
	"github.com/mendelgusmao/supereasy/backend/router"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	router.Router.HandleFunc("/users/authenticate", authenticateUser).Methods("POST")
}

func authenticateUser(w http.ResponseWriter, r *http.Request) {
	db := context.Get(r, "db").(*mongo.Database)
	dao := newDAO(db)

	credentials := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		log.Printf("authenticateUser: %s", err)
		http.Error(w, "", http.StatusBadRequest)
		json.NewEncoder(w).Encode(messages.Error{Error: "invalid payload"})

		return
	}

	_, err := dao.findByEmailAndPassword(credentials.Email, credentials.Password)

	if err != nil {
		log.Printf("authenticateUser: %s", err)

		if err == mongo.ErrNoDocuments {
			http.Error(w, "", http.StatusUnauthorized)
			json.NewEncoder(w).Encode(messages.Error{Error: "wrong email and password combination"})

			return
		}

		http.Error(w, "", http.StatusInternalServerError)
		json.NewEncoder(w).Encode(messages.Error{Error: "unexpected error"})

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
