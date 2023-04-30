package usecase

import (
	"PassargadUser/entities/domain"
	"PassargadUser/entities/requestModel"
	"PassargadUser/pkg/crypt"
	"PassargadUser/pkg/messages"
	"PassargadUser/repository"
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-jwt/jwt"
	"gorm.io/gorm"
	"net/http"
)

func CreateUser(ctx *gin.Context, input requestModel.CreateRequest) (err error, output *requestModel.BasicResponse) {
	err, _ = repository.UsrRepo.GetByUsername(ctx, input.Username)
	if err == nil {
		return errors.New(messages.UserNameExists), &requestModel.BasicResponse{
			Message: messages.UserNameExists,
			Code:    http.StatusBadRequest,
		}
	}
	if err != gorm.ErrRecordNotFound {
		return err, &requestModel.BasicResponse{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}
	usrCreate := &domain.User{
		Username:  input.Username,
		Password:  crypt.GetMD5Hash(input.Password),
		Lastname:  input.Lastname,
		Firstname: input.Firstname,
		Email:     input.Email,
	}
	err = repository.UsrRepo.Create(ctx, usrCreate)
	if err != nil {
		return err, &requestModel.BasicResponse{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}
	return nil, &requestModel.BasicResponse{
		Message: messages.CreatedSuccessfully,
		Code:    http.StatusCreated,
	}
}

func LoginUser(ctx *gin.Context, input requestModel.LoginRequest) (err error, output *requestModel.LoginResponse) {
	err, usr := repository.UsrRepo.GetByUsername(ctx, input.Username)

	if err == gorm.ErrRecordNotFound {
		return err, &requestModel.LoginResponse{
			Message: messages.UserNameNotExist,
			Code:    http.StatusBadRequest,
		}
	}
	if err != nil {
		return err, &requestModel.LoginResponse{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	if crypt.GetMD5Hash(input.Password) == usr.Password {
		err, token := crypt.GenerateJWT(usr.Username)
		if err != nil {
			return err, &requestModel.LoginResponse{
				Message: err.Error(),
				Code:    http.StatusInternalServerError,
			}
		}
		return nil, &requestModel.LoginResponse{
			JWTToken: token,
			Message:  messages.LoginSuccessful,
			Code:     http.StatusOK,
		}
	}

	return err, &requestModel.LoginResponse{
		Message: messages.WrongPassword,
		Code:    http.StatusBadRequest,
	}
}
