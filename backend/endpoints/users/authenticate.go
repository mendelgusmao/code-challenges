package users

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/mendelgusmao/supereasy/backend/config"
	"github.com/mendelgusmao/supereasy/backend/middleware"
	"github.com/mendelgusmao/supereasy/backend/router"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	router.Router.HandleFunc("/users/authenticate", authenticateUser).Methods("POST")
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user.ID,
	})

	tokenString, err := token.SignedString([]byte(config.Backend.JWTSecret))

	if err != nil {
		log.Printf("authenticateUser: %s", err)
		error(http.StatusInternalServerError, "error generating token")

		return
	}

	w.Header().Add("Authorization", tokenString)
	w.WriteHeader(http.StatusNoContent)
}
