package middleware

import (
	"PassargadUser/config"
	"PassargadUser/pkg/crypt"
	"PassargadUser/pkg/messages"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

type JWTData struct {
	ExpireTime int64
	Username   string
}

func JWTVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Header["Token"] == nil {
			ctx.Set(crypt.UserTypeKey, crypt.Guest)
			ctx.Next()
			return
		}

		err, jwtData := VerifyToken(ctx.Request.Header["Token"][0], config.SampleSecretKey)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": messages.UnAuthorized})
			return
		}

		ctx.Set(crypt.UserTypeKey, crypt.LoggedIn)
		ctx.Set(crypt.ExpireKey, jwtData.ExpireTime)
		ctx.Set(crypt.UsernameKey, jwtData.Username)

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

	expireTime, ok := claims["exp"].(float64)
	if !ok {
		return errors.New("invalid JWT token"), nil
	}
	username, ok := claims["user"].(string)
	if !ok {
		return errors.New("invalid JWT token"), nil
	}

	return nil, &JWTData{
		ExpireTime: int64(expireTime),
		Username:   username,
	}
}
