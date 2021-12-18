package service

import (
	"bbs-like-backend/model"
)

type userservice struct {
	userRepo model.UserRepository
}

func NewUserService(ur model.UserRepository) model.UserService {
	return &userservice{
		userRepo: ur,
	}
}

func (us *userservice) CreateUser(user model.User) error {
	us.userRepo.CreateUser(user)
	return nil
}

func (us *userservice) LoginCheck(user model.User) bool {
	return false
}

func GetUsersList() {
	//todo
}
