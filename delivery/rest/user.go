package rest

import (
	"PassargadUser/entities/requestModel"
	"PassargadUser/pkg/crypt"
	"PassargadUser/pkg/messages"
	"PassargadUser/repository"
	"PassargadUser/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteUser(ctx *gin.Context) {

}

func CreateUser(ctx *gin.Context) {
	var input requestModel.CreateRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": messages.InternalServerError})
		return
	}

	err, out := usecase.CreateUser(ctx, input)
	if err != nil {
		ctx.JSON(out.Code, gin.H{"error": out.Message})
		return
	}
	ctx.JSON(out.Code, gin.H{"data": out})
}

func UpdateUser(ctx *gin.Context) {

}

func GetUserInfo(ctx *gin.Context) {
	userType, _ := ctx.Get(crypt.UserTypeKey)
	userType, ok := userType.(crypt.UserType)
	if !ok || userType != crypt.LoggedIn {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": messages.UnAuthorized})
		return
	}

	uname, _ := ctx.Get(crypt.UsernameKey)
	username, ok := uname.(string)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": messages.UnAuthorized})
		return
	}

	err, user := repository.UsrRepo.GetByUsername(ctx, username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": messages.InternalServerError})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": requestModel.UserInfoResponse{
		Email: user.Email,
	}})
}

func LoginUser(ctx *gin.Context) {
	var input requestModel.LoginRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": messages.InternalServerError})
		return
	}

	err, out := usecase.LoginUser(ctx, input)
	if err != nil {
		ctx.JSON(out.Code, gin.H{"error": out.Message})
		return
	}
	ctx.JSON(out.Code, gin.H{"data": out})
}
