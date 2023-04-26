package repository

import (
	"PassargadUser/domain"
	"context"
	"gorm.io/gorm"
	"time"
)

var (
	User UserRepositoryInterface = user{}
)

type user struct {
	DB *gorm.DB
}

type UserRepositoryInterface interface {
	Create(ctx context.Context, user *domain.User) (err error, userId uint)
	Delete(ctx context.Context, userId uint) (err error)
	GetByUsername(ctx context.Context, userName string) (err error, user *domain.User)
	GetById(ctx context.Context, userId uint) (err error, user *domain.User)
	Update(ctx context.Context, user *domain.User) (err error)
}

func (r user) Create(ctx context.Context, user *domain.User) (err error, userId uint) {
	dbc := r.DB.Create(user)
	if dbc.Error != nil {
		return dbc.Error, 0
	}
	return nil, user.ID
}

func (r user) Delete(ctx context.Context, userId uint) (err error) {
	usr := &domain.User{}
	rdb := r.DB.First(usr, userId)
	if rdb.Error != nil {
		return rdb.Error
	}
	usr.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}
	rdb = r.DB.Save(usr)
	if rdb.Error != nil {
		return rdb.Error
	}
	return nil
}

func (r user) GetByUsername(ctx context.Context, userName string) (err error, user *domain.User) {

}

func (r user) GetById(ctx context.Context, userId uint) (err error, user *domain.User) {

}

func (r user) Update(ctx context.Context, user *domain.User) (err error) {

}
