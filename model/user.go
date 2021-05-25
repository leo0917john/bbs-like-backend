package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
}

func (u *User) TableName() string {
	return "users"
}
