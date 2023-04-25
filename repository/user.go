package repository

import (
	"PassargadUser/domain"
	"context"
	"gorm.io/gorm"
)

var (
	User UserRepositoryInterface = user{}
)

type user struct {
	DB *gorm.DB
}

type UserRepositoryInterface interface {
	Create(ctx context.Context, user *domain.User) (err error, userId string)
	Delete(ctx context.Context, userId string) (err error)
	GetByUsername(ctx context.Context, userName string) (err error, user *domain.User)
	GetById(ctx context.Context, userId string) (err error, user *domain.User)
	Update(ctx context.Context, user *domain.User) (err error)
}

func (r user) Create(ctx context.Context, user *domain.User) (err error, userId string) {
	r.DB.Create(user)
}

func (r user) Delete(ctx context.Context, userId string) (err error) {

}

func (r user) GetByUsername(ctx context.Context, userName string) (err error, user *domain.User) {

}

func (r user) GetById(ctx context.Context, userId string) (err error, user *domain.User) {

}

func (r user) Update(ctx context.Context, user *domain.User) (err error) {

}
