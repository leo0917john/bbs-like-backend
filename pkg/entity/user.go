package entity

import (
	"bbs-like-backend/pkg/model"
)

type UserService interface {
	CreateUser(model.User) error
	LoginCheck(model.UserLoginData) bool
	TableCheck() error
	GetUsersList(string) ([]model.User, error)
}

type UserRepository interface {
	Init()
	TableExist() error
	GetUser(string) (model.User, error)
	GetUsers(string) ([]model.User, error)
	CreateUser(model.User) error
	UpadeUser()
	DeleteUser()
	UserExist(string) bool
	Close()
}
