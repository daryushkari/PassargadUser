package crypt

import (
	"PassargadUser/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	JWTExpireTime = 10 * time.Minute
)

type UserType string

const (
	LoggedIn UserType = "logged-in"
	Guest    UserType = "guest"
)

const (
	UserTypeKey = "user-type"
	ExpireKey   = "exp"
	UsernameKey = "user"
)

func GenerateJWT(username string) (error, string) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(JWTExpireTime).Unix()
	claims["user"] = username

	tokenString, err := token.SignedString(config.SampleSecretKey)
	if err != nil {
		return err, ""
	}

	return nil, tokenString
}
