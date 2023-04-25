package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	firstname string
	lastname  string
	username  string
	password  string
	email     string
}
