package crypt

import (
	"github.com/golang-jwt/jwt"
	"time"
)

const (
	JWTExpireTime = 10 * time.Minute
)

func GenerateJWT(username string) (error, string) {
	var sampleSecretKey = []byte("SecretYouShouldHide")

	token := jwt.New(jwt.SigningMethodEdDSA)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(JWTExpireTime)
	claims["user"] = username

	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return err, ""
	}

	return nil, tokenString
}
