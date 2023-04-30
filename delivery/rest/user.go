package rest

import (
	"PassargadUser/entities/requestModel"
	"PassargadUser/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteUser(ctx *gin.Context) {

}

func CreateUser(ctx *gin.Context) {
	var input requestModel.CreateRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err, out := usecase.CreateUser(ctx, input)
	if err != nil {
		ctx.JSON(out.Code, gin.H{"error": out})
		return
	}
	ctx.JSON(out.Code, gin.H{"data": out})
}

func UpdateUser(ctx *gin.Context) {

}

func GetUserInfo(ctx *gin.Context) {

}

func LoginUser(ctx *gin.Context) {
	var input requestModel.LoginRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err, out := usecase.LoginUser(ctx, input)
	if err != nil {
		ctx.JSON(out.Code, gin.H{"error": out})
	}
	ctx.JSON(out.Code, gin.H{"data": out})
}
