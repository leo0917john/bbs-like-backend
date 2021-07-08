package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"bbs-like-backend/db"
	"bbs-like-backend/handler"
	"bbs-like-backend/lib/security"
	"bbs-like-backend/middleware"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

// Temporary setting declaration
// TODO: move to .env file
const (
	DB_HOST     = "0.0.0.0:5432"
	DB_NAME     = "test"
	DB_USER     = "postgre_user"
	DB_PASSWORD = "postgre_pwd"
	//  JWT setting
	SECRET_KEY     = "secret"
	TOKEN_LIFETIME = 120
)

func init() {

	// jwt setting
	secretKey := []byte(SECRET_KEY)
	tokenLifeTime := time.Duration(TOKEN_LIFETIME) * time.Minute

	// jwt init setting
	security.Init(secretKey, tokenLifeTime)
}

func HandleHello(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})

}

func DBconnection() (db *gorm.DB) {
	// connStr := "postgres://pqgotest:password@localhost/pqgotest"
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", DB_USER, DB_PASSWORD, DB_HOST, DB_NAME)
	sqldb, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	err = sqldb.Ping()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println("successfull connected!")
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqldb,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting database")
		panic(err.Error())
	} else {
		fmt.Println("Connected to database")
	}
	return gormDB
}

func handleCors(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "successful"})
}

func main() {

	db.Open()
	route := gin.Default()
	route.Use(cors.New(middleware.CorsSetting()))

	route.GET("/test", HandleHello)
	route.POST("/user", handler.UserCreate)
	route.GET("/userlist", handler.GetUsersList)
	route.POST("/cors", handleCors)
	route.POST("/login", handler.Login)
	route.Run(":5050")
}
