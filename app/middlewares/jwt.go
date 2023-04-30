package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
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

func JWTParser(tkString string) {
	tk, err := jwt.Parse(tkString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodECDSA)
		if !ok {
			log.Println("sidjfisdjfijds")
		}
		return "", nil
	})

	if err != nil {
		log.Println(err, "bad error")
	}

	if !tk.Valid || tk.Claims == nil {
		log.Println("nok")
	}

	claims, ok := tk.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("nok")
	}

	expireTime, ok := claims["exp"].(time.Time)
	if !ok {
		log.Println("bad")
	}
	username, ok := claims["user"].(string)
	if !ok {
		log.Println("bad")
	}
	log.Println(username, expireTime)

}
