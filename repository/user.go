package repository

import (
	"PassargadUser/domain"
	"PassargadUser/pkg/crypt"
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
	Create(ctx context.Context, usr *domain.User) (err error, userId uint)
	Delete(ctx context.Context, userId uint) (err error)
	GetByUsername(ctx context.Context, userName string) (err error, usr *domain.User)
	GetById(ctx context.Context, userId uint) (err error, usr *domain.User)
	Update(ctx context.Context, usr *domain.User) (err error)
}

func (r user) Create(ctx context.Context, usr *domain.User) (err error, userId uint) {
	dbc := r.DB.Create(usr)
	if dbc.Error != nil {
		return dbc.Error, 0
	}
	return nil, usr.ID
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

func (r user) GetByUsername(ctx context.Context, userName string) (err error, usr *domain.User) {
	rdb := r.DB.First(usr, "username = ?", userName)
	if rdb.Error != nil {
		return rdb.Error, nil
	}
	return nil, usr
}

func (r user) GetById(ctx context.Context, userId uint) (err error, usr *domain.User) {
	rdb := r.DB.First(usr, userId)
	if rdb.Error != nil {
		return rdb.Error, nil
	}
	return nil, usr
}

func (r user) Update(ctx context.Context, usr *domain.User) (err error) {
	rdb := r.DB.Model(usr).Updates(domain.User{
		Password:  crypt.GetMD5Hash(usr.Password),
		Email:     usr.Email,
		Firstname: usr.Email,
		Lastname:  usr.Lastname,
	})
	return rdb.Error
}
