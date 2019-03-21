package users

import (
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/mendelgusmao/supereasy/backend/middleware"
	"github.com/mendelgusmao/supereasy/backend/router"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	subrouter := router.Router.PathPrefix("/users/authenticate").Subrouter()
	subrouter.Use(middleware.GenerateToken)

	subrouter.HandleFunc("", authenticateUser).Methods("POST")
}

func authenticateUser(w http.ResponseWriter, r *http.Request) {
	db := context.Get(r, "db").(*mongo.Database)
	jsonDecoder := context.Get(r, "jsonDecoder").(middleware.JSONDecoderFunc)
	error := context.Get(r, "error").(middleware.ErrorFunc)

	dao := newDAO(db)

	credentials := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if !jsonDecoder(&credentials) {
		return
	}

	user, err := dao.findByEmailAndPassword(credentials.Email, credentials.Password)

	if err != nil {
		log.Printf("authenticateUser: %s", err)

		if err == mongo.ErrNoDocuments {
			error(http.StatusUnauthorized, "wrong email and password combination")

			return
		}

		error(http.StatusInternalServerError, "unexpected error")

		return
	}

	generateToken := context.Get(r, "generateToken").(middleware.GenerateTokenFunc)
	generateToken(user.ID)
}
