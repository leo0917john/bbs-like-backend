package db

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_HOST     = "0.0.0.0:5432"
	DB_NAME     = "test"
	DB_USER     = "postgre_user"
	DB_PASSWORD = "postgre_pwd"
)

var DB *gorm.DB

func Open() {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", DB_USER, DB_PASSWORD, DB_HOST, DB_NAME)
	sqldb, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = sqldb.Ping()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Got db response")
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqldb,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting database")
		panic(err.Error())
	} else {
		fmt.Println("DB connection successful")
	}
	DB = db
}

func Show() {
	fmt.Println(DB)
}

func Close() {
	sqlDB, _ := DB.DB()
	sqlDB.Close()
}
