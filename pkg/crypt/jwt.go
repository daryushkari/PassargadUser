package crypt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	JWTExpireTime = 10 * time.Minute
)

func GenerateJWT(username string) (error, string) {
	var sampleSecretKey = []byte("SecretYouShouldsdfsdfsdfsdfsdffsdfsdfHide")

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(JWTExpireTime)
	claims["user"] = username

	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return err, ""
	}

	return nil, tokenString
}
