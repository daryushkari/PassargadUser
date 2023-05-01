package middleware

import (
	"PassargadUser/config"
	"PassargadUser/pkg/messages"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type JWTData struct {
	ExpireTime string
	Username   string
}

func JWTVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println(ctx.Request.Header["Token"])
		if ctx.Request.Header["Token"] == nil {
			ctx.Set("user-type", "guest")
			ctx.Next()
			return
		}

		err, jwtData := VerifyToken(ctx.Request.Header["Token"][0], config.SampleSecretKey)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": messages.UnAuthorized})
			return
		}

		expTime, err := time.Parse("2023-05-01T11:23:14", jwtData.ExpireTime)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": messages.UnAuthorized})
			return
		}

		if time.Now().After(expTime) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": messages.TokenExpired})
			return
		}

		ctx.Set("user-type", "logged-in")
		ctx.Set("exp", jwtData.ExpireTime)
		ctx.Set("user", jwtData.Username)

		ctx.Next()
	}
}

func VerifyToken(tokenString string, secretKey []byte) (err error, jwtData *JWTData) {
	tk, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return err, nil
	}

	if !tk.Valid {
		return errors.New("invalid token signature"), nil
	}

	claims, ok := tk.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("invalid JWT token"), nil
	}

	expireTime, ok := claims["exp"].(string)
	if !ok {
		return errors.New("invalid JWT token"), nil
	}
	username, ok := claims["user"].(string)
	if !ok {
		return errors.New("invalid JWT token"), nil
	}

	return nil, &JWTData{
		ExpireTime: expireTime,
		Username:   username,
	}
}

//func JWTParser(tkString string) {
//	tk, err := jwt.Parse(tkString, func(token *jwt.Token) (interface{}, error) {
//		*jwt.SigningMethodHS256.Alg()
//		ok := token.Method.(*jwt.SigningMethodECDSA)
//		if !ok {
//			log.Println("sidjfisdjfijds")
//		}
//		return "", nil
//	})
//
//	if err != nil {
//		log.Println(err, "bad error")
//	}
//
//	if !tk.Valid || tk.Claims == nil {
//		log.Println("nok")
//	}
//
//	claims, ok := tk.Claims.(jwt.MapClaims)
//	if !ok {
//		log.Println("nok")
//	}
//
//	expireTime, ok := claims["exp"].(time.Time)
//	if !ok {
//		log.Println("bad")
//	}
//	username, ok := claims["user"].(string)
//	if !ok {
//		log.Println("bad")
//	}
//	log.Println(username, expireTime)
//
//}
