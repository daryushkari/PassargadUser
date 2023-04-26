package repository

import (
	"PassargadUser/domain"
	"context"
	"gorm.io/gorm"
	"time"
)

var (
	UsrRepo UserRepositoryInterface = &UserRepository{}
)

type UserRepository struct {
	DB *gorm.DB
}

type UserRepositoryInterface interface {
	InitDB(db *gorm.DB)
	Create(ctx context.Context, usr *domain.User) (err error, userId uint)
	Delete(ctx context.Context, userId uint) (err error)
	GetByUsername(ctx context.Context, userName string) (err error, usr *domain.User)
	GetById(ctx context.Context, userId uint) (err error, usr *domain.User)
	Update(ctx context.Context, usr *domain.User) (err error)
}

func (r *UserRepository) InitDB(db *gorm.DB) {
	r.DB = db
}

func (r *UserRepository) Create(ctx context.Context, usr *domain.User) (err error, userId uint) {
	tx := r.DB.WithContext(ctx)
	rdb := tx.Create(usr)
	if rdb.Error != nil {
		return rdb.Error, 0
	}
	return nil, usr.ID
}

func (r *UserRepository) Delete(ctx context.Context, userId uint) (err error) {
	usr := &domain.User{}
	tx := r.DB.WithContext(ctx)
	rdb := tx.First(usr, userId)
	if rdb.Error != nil {
		return rdb.Error
	}
	usr.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}
	rdb = tx.Save(usr)
	if rdb.Error != nil {
		return rdb.Error
	}
	return nil
}

func (r *UserRepository) GetByUsername(ctx context.Context, userName string) (err error, usr *domain.User) {
	tx := r.DB.WithContext(ctx)
	usr = &domain.User{}
	rdb := tx.First(usr, "username = ?", userName)
	if rdb.Error != nil {
		return rdb.Error, nil
	}
	return nil, usr
}

func (r *UserRepository) GetById(ctx context.Context, userId uint) (err error, usr *domain.User) {
	usr = &domain.User{}
	tx := r.DB.WithContext(ctx)
	rdb := tx.First(usr, userId)
	if rdb.Error != nil {
		return rdb.Error, nil
	}
	return nil, usr
}

func (r *UserRepository) Update(ctx context.Context, usr *domain.User) (err error) {
	usr = &domain.User{}
	tx := r.DB.WithContext(ctx)
	rdb := tx.Model(usr).Updates(domain.User{
		Password:  usr.Password,
		Email:     usr.Email,
		Firstname: usr.Email,
		Lastname:  usr.Lastname,
	})
	return rdb.Error
}
