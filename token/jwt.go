package token

import (
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const secretKey = "demoService"

func NewToken(email string) (string, error) {
	t := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"iat": t,
		"exp": t.Add(time.Hour * 1),
	})

	return token.SignedString([]byte(secretKey))
}

func TokenVerify(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		log.Println(err)
		return false
	}

	return token.Valid
}
