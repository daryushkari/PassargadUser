package rest

import (
	"PassargadUser/entities/requestModel"
	"PassargadUser/pkg/crypt"
	"PassargadUser/pkg/messages"
	"PassargadUser/repository"
	"PassargadUser/usecase"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"gorm.io/gorm"
	"net/http"
)

func DeleteUser(ctx *gin.Context) {
	tr := otel.Tracer("deleteUser")
	_, span := tr.Start(ctx, "deleteUser-delivery")
	defer span.End()

	userType, _ := ctx.Get(crypt.UserTypeKey)
	userType, ok := userType.(crypt.UserType)
	if !ok || userType != crypt.LoggedIn {
		span.SetAttributes(attribute.Key("error").String(messages.UnAuthorized))
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": messages.UnAuthorized})
		return
	}

	uname, _ := ctx.Get(crypt.UsernameKey)
	username, ok := uname.(string)
	if !ok {
		span.SetAttributes(attribute.Key("error").String("user name not exist"))
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": messages.UnAuthorized})
		return
	}
	span.SetAttributes(attribute.Key("username").String(username))

	err, usr := repository.UsrRepo.GetByUsername(ctx, username)
	if err != nil {
		span.SetAttributes(attribute.Key("error").String(err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": messages.InternalServerError})
		return
	}

	err = repository.UsrRepo.Delete(ctx, usr.ID)
	if err != nil {
		span.SetAttributes(attribute.Key("error").String(err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": messages.InternalServerError})
		return
	}
	span.SetAttributes(attribute.Key("message").String("success"))
	ctx.JSON(http.StatusOK, gin.H{"data": requestModel.BasicResponse{
		Message: messages.APISuccess,
		Code:    http.StatusOK,
	}})
}

func CreateUser(ctx *gin.Context) {
	tr := otel.Tracer("createUser")
	_, span := tr.Start(ctx, "createUser-delivery")
	defer span.End()

	var input requestModel.CreateRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		span.SetAttributes(attribute.Key("error").String(err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": messages.BadRequest})
		return
	}
	span.SetAttributes(attribute.Key("username").String(input.Username))

	err, out := usecase.CreateUser(ctx, input)
	if err != nil {
		span.SetAttributes(attribute.Key("error").String(err.Error()))
		ctx.JSON(out.Code, gin.H{"error": out.Message})
		return
	}
	ctx.JSON(out.Code, gin.H{"data": out})
}

func UpdateUser(ctx *gin.Context) {
	tr := otel.Tracer("updateUser")
	_, span := tr.Start(ctx, "updateUser-delivery")
	defer span.End()

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
	span.SetAttributes(attribute.Key("username").String(username))

	var input requestModel.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		span.SetAttributes(attribute.Key("error").String(err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": messages.BadRequest})
		return
	}

	err, usr := repository.UsrRepo.GetByUsername(ctx, username)
	if err != nil {
		span.SetAttributes(attribute.Key("error").String(err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": messages.InternalServerError})
		return
	}

	usr.Password = crypt.GetMD5Hash(input.Password)
	usr.Firstname = input.FirstName
	usr.Lastname = input.LastName
	usr.Email = input.Email

	err = repository.UsrRepo.Update(ctx, usr)
	if err != nil {
		span.SetAttributes(attribute.Key("error").String(err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": messages.InternalServerError})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": requestModel.BasicResponse{
		Message: messages.APISuccess,
		Code:    http.StatusOK,
	}})
}

func GetUserInfo(ctx *gin.Context) {
	tr := otel.Tracer("getUser")
	_, span := tr.Start(ctx, "getUser-delivery")
	defer span.End()

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
	span.SetAttributes(attribute.Key("username").String(username))

	err, user := repository.UsrRepo.GetByUsername(ctx, username)
	if err != nil {
		span.SetAttributes(attribute.Key("error").String(err.Error()))
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
	tr := otel.Tracer("loginUser")
	_, span := tr.Start(ctx, "loginUser-delivery")
	defer span.End()

	var input requestModel.LoginRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		span.SetAttributes(attribute.Key("error").String(err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": messages.BadRequest})
		return
	}

	span.SetAttributes(attribute.Key("username").String(input.Username))

	err, out := usecase.LoginUser(ctx, input)
	if err != nil {
		span.SetAttributes(attribute.Key("error").String(err.Error()))
		ctx.JSON(out.Code, gin.H{"error": out.Message})
		return
	}
	span.SetAttributes(attribute.Key("jwt").String(out.JWTToken))
	ctx.JSON(out.Code, gin.H{"data": out})
}
