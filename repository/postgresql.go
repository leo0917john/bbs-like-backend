package repository

import (
	"bbs-like-backend/model"
	"log"

	"gorm.io/gorm"
)

type postgreSQLRepository struct {
	DB *gorm.DB
}

func NewPostgreSQLRepository(db *gorm.DB) model.UserRepository {
	return &postgreSQLRepository{
		DB: db,
	}
}

func (pr *postgreSQLRepository) GetUser() {

}
func (pr *postgreSQLRepository) CreateUser(user model.User) error {
	//
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
func (pr *postgreSQLRepository) FindUser() bool {
	return false
}
func (pr *postgreSQLRepository) Close() {
	//todo
}
