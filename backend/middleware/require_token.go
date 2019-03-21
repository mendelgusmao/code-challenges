package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/mendelgusmao/supereasy/backend/config"
)

func RequireToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(config.Backend.JWTSecret), nil
		})

		if err != nil || (err == nil && !token.Valid) {
			http.Error(w, "", http.StatusForbidden)
			json.NewEncoder(w).Encode(errorMessage{Error: "invalid authorization. please authenticate"})

			return
		}

		next.ServeHTTP(w, r)
	})
}
