package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"bbs-like-backend/db"
	"bbs-like-backend/handler"
	"bbs-like-backend/middleware"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

const (
	DB_HOST     = "0.0.0.0:5432"
	DB_NAME     = "test"
	DB_USER     = "postgre_user"
	DB_PASSWORD = "postgre_pwd"
)

// var user_map = make(map[string]string)

func HandleHello(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})

}

func HandleRegister(c *gin.Context) {
	var user User
	method := c.Request.Method
	if method == "OPTIONS" {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization") //自定义 Header
		// c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		// c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		// c.Header("Access-Control-Allow-Credentials", "true")
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	c.Header("Access-Control-Allow-Origin", "*")
	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println("register failed")
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "successful"})
	fmt.Println("register success")
}

func HandleUserList(c *gin.Context) {

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

	// db.AutoMigrate(&User{})
	// user := User{Username: "Leo", Password: "123"}

	// result := db.Create(&user)
	// if result.Error != nil {
	// 	log.Fatal(result.Error)
	// 	panic(result.Error)
	// } else {
	// 	fmt.Println("Data Create complete")
	// }
	db.Open()
	route := gin.Default()
	route.Use(cors.New(middleware.Cors_init()))

	route.GET("/test", HandleHello)
	route.OPTIONS("/register", HandleRegister)
	route.POST("/register", HandleRegister)
	// route.OPTIONS("/user", handler.UserCreate)
	route.POST("/user", handler.UserCreate)
	route.POST("/userlist", handler.GetList)
	route.POST("/cors", handleCors)
	route.Run(":5050")
}
