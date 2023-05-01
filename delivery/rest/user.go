package rest

import (
	"PassargadUser/entities/requestModel"
	"PassargadUser/pkg/crypt"
	"PassargadUser/pkg/messages"
	"PassargadUser/repository"
	"PassargadUser/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func DeleteUser(ctx *gin.Context) {
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

	err, usr := repository.UsrRepo.GetByUsername(ctx, username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": messages.InternalServerError})
		return
	}

	err = repository.UsrRepo.Delete(ctx, usr.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": messages.InternalServerError})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": requestModel.BasicResponse{
		Message: messages.APISuccess,
		Code:    http.StatusOK,
	}})
}

func CreateUser(ctx *gin.Context) {
	var input requestModel.CreateRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": messages.BadRequest})
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

	var input requestModel.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": messages.BadRequest})
		return
	}

	err, usr := repository.UsrRepo.GetByUsername(ctx, username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": messages.InternalServerError})
		return
	}

	usr.Password = crypt.GetMD5Hash(input.Password)
	usr.Firstname = input.FirstName
	usr.Lastname = input.LastName
	usr.Email = input.Email

	err = repository.UsrRepo.Update(ctx, usr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": messages.InternalServerError})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": requestModel.BasicResponse{
		Message: messages.APISuccess,
		Code:    http.StatusOK,
	}})
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
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": messages.UserNameNotExist})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": messages.InternalServerError})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": requestModel.UserInfoResponse{
		Email:     user.Email,
		FirstName: user.Firstname,
		LastName:  user.Lastname,
	}})
}

func LoginUser(ctx *gin.Context) {
	var input requestModel.LoginRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": messages.BadRequest})
		return
	}

	err, out := usecase.LoginUser(ctx, input)
	if err != nil {
		ctx.JSON(out.Code, gin.H{"error": out.Message})
		return
	}
	ctx.JSON(out.Code, gin.H{"data": out})
}
