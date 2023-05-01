package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
)

func JWTVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Header["Token"] == nil {
			ctx.Set("user-type", "guest")
		}

		//ctx.Header()
		ctx.Next()
	}
}

func VerifyToken(tokenString string, secretKey []byte) (*jwt.Token, error) {
	tk, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !tk.Valid {
		return nil, errors.New("invalid token signature")
	}

	claims, ok := tk.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token signature")
	}

	expireTime, ok := claims["exp"].(string)
	if !ok {
		log.Println("bad")
	}
	username, ok := claims["user"].(string)
	if !ok {
		log.Println("bad")
	}
	log.Println(username, expireTime)

	return tk, nil
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
