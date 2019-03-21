package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/mendelgusmao/supereasy/backend/config"
)

type GenerateTokenFunc func(string)

func GenerateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		context.Set(r, "generateToken", GenerateTokenFunc(func(userID string) {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"user": userID,
			})

			tokenString, err := token.SignedString([]byte(config.Backend.JWTSecret))

			if err != nil {
				log.Printf("middleware.GenerateToken: %s", err)
				http.Error(w, "", http.StatusInternalServerError)
				json.NewEncoder(w).Encode(errorMessage{Error: "error generating token"})

				return
			}

			w.Header().Add("Authorization", tokenString)
			w.WriteHeader(http.StatusNoContent)

			return
		}))

		next.ServeHTTP(w, r)
	})
}
