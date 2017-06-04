// Author: Vincent Composieux <vincent.composieux@gmail.com>

package main

import (
	"../app"

	jwt "github.com/dgrijalva/jwt-go"

	"fmt"
)

// Generates and prints a JWT token with the secret given in application configuration.
func main() {
	config := app.GetConfig()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"app": config.JWT.Name,
	})

	tokenString, _ := token.SignedString([]byte(config.JWT.Secret))

	fmt.Println(tokenString)
}
