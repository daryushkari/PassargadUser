package usecase

import (
	"PassargadUser/entities/domain"
	"PassargadUser/entities/requestModel"
	"PassargadUser/pkg/crypt"
	"PassargadUser/pkg/messages"
	"PassargadUser/repository"
	"errors"
	"github.com/gin-gonic/gin"
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
		return err, nil
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
		return err, nil
	}
	return nil, &requestModel.BasicResponse{
		Message: messages.CreatedSuccessfully,
		Code:    http.StatusCreated,
	}
}
