package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Firstname string
	Lastname  string
	Username  string
	Password  string
	Email     string
}
