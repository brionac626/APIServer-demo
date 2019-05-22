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
		"iat":   t,
		"exp":   t.Add(time.Hour * 1),
		"email": email,
	})

	return token.SignedString([]byte(secretKey))
}

func TokenVerify(tokenString, email string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		log.Println(err)
		return false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("claims not exists")
		return false
	}
	if claims["email"] != email {
		log.Println("email not equal")
		return false
	}

	return token.Valid
}
