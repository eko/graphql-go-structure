// Author: Vincent Composieux <vincent.composieux@gmail.com>

package security

import (
	"../app"
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

// Handle security middleware aims to implement a JWT authentication.
func Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")[7:]

		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			config := app.GetConfig()

			return []byte(config.JWT.Secret), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			log.Printf("JWT Authenticated OK (app: %s)", claims["app"])

			next.ServeHTTP(w, r)
		}
	})
}
