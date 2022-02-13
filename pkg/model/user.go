package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
}

type UserLoginData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) TableName() string {
	return "users"
}
