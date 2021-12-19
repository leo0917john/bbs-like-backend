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

type UserService interface {
	CreateUser(User) error
	LoginCheck(UserLoginData) bool
	TableCheck() error
	GetUsersList(string) ([]User, error)
}

type UserRepository interface {
	Init()
	TableExist() error
	GetUser(string) (User, error)
	GetUsers(string) ([]User, error)
	CreateUser(User) error
	UpadeUser()
	DeleteUser()
	UserExist(string) bool
	Close()
}
