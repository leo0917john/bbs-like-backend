package service

import (
	"bbs-like-backend/pkg/entity"
	"bbs-like-backend/pkg/model"
	"errors"
	"log"
)

type userservice struct {
	userRepo entity.UserRepository
}

func NewUserService(ur entity.UserRepository) entity.UserService {
	return &userservice{
		userRepo: ur,
	}
}

func (us *userservice) TableCheck() error {
	res := us.userRepo.TableExist()
	return res
}

func (us *userservice) CreateUser(user model.User) error {
	if us.userRepo.UserExist(user.Username) {
		return errors.New("UserName is exist")
	} else {
		us.userRepo.CreateUser(user)
	}
	return nil
}

func (us *userservice) LoginCheck(loginer model.UserLoginData) bool {
	user, err := us.userRepo.GetUser(loginer.Username)
	if err != nil {
		log.Println(err)
		return false
	}
	if user.Password != loginer.Password {
		log.Printf("Loginer:%v password authentication fail", loginer.Username)
		return false
	}
	return true
}

func (us *userservice) GetUsersList(keyword string) ([]model.User, error) {
	var tmp_slice []model.User
	tmp_slice, err := us.userRepo.GetUsers("")
	if err != nil {
		return tmp_slice, err
	}
	return tmp_slice, nil
}
