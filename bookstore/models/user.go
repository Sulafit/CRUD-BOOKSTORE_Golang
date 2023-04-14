package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int
	Username string
	Password string
}


func NewUser(username, password string) *User {
	return &User{
		Username: username,
		Password: password,
	}
}
