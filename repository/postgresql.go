package repository

import (
	"bbs-like-backend/model"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type postgreSQLRepository struct {
	DB *gorm.DB
}

type postgreSQLRepositoryError struct {
	msg string
}

func NewPostgreSQLRepository(db *gorm.DB) model.UserRepository {
	return &postgreSQLRepository{
		DB: db,
	}
}

func NewPostgreSQLRepositoryError(msg string) error {
	return &postgreSQLRepositoryError{
		msg: msg,
	}
}

func (pe *postgreSQLRepositoryError) Error() string {
	return fmt.Sprintf("[PostgreSQLRepo]Error:%v", pe.msg)
}

func (pr *postgreSQLRepository) Init() {
	pr.DB.Debug().AutoMigrate(&model.User{})

}

func (pr *postgreSQLRepository) TableExist() error {
	migrator := pr.DB.Migrator()
	has := migrator.HasTable(&model.User{})
	if !has {
		return NewPostgreSQLRepositoryError("UserTable not exist")
	}
	return nil
}

func (pr *postgreSQLRepository) GetUser(name string) (model.User, error) {
	var target model.User
	result := pr.DB.Where("username=?", name).First(&target)
	if result.RowsAffected != 1 {
		return target, NewPostgreSQLRepositoryError("User Data Error")
	} else if result.Error != nil {
		return target, NewPostgreSQLRepositoryError(result.Error.Error())
	}
	return target, nil
}

func (pr *postgreSQLRepository) GetUsers(keyword string) ([]model.User, error) {
	var UsersList []model.User
	pr.DB.Find(&UsersList)
	return UsersList, nil
}

func (pr *postgreSQLRepository) CreateUser(user model.User) error {
	tx_result := pr.DB.Debug().Create(&user)
	if tx_result.Error != nil {
		log.Println(tx_result.Error)
	}
	if tx_result.RowsAffected != 1 {
		log.Println("RowsAffected Number failt")
	}
	return nil
}
func (pr *postgreSQLRepository) UpadeUser() {
	//todo
}
func (pr *postgreSQLRepository) DeleteUser() {
	//todo
}
func (pr *postgreSQLRepository) UserExist(name string) bool {
	var res bool = false
	pr.DB.Raw("SELECT 1 FROM users where username = ? LIMIT 1;", name).Scan(&res)
	return res
}
func (pr *postgreSQLRepository) Close() {
	//todo
}
