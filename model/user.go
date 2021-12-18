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

type UserService interface {
	CreateUser(User) error
	LoginCheck(User) bool
}

type UserRepository interface {
	GetUser()
	CreateUser(User) error
	UpadeUser()
	DeleteUser()
	FindUser() bool
	Close()
}
