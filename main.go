package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"bbs-like-backend/handler"
	"bbs-like-backend/lib/security"
	"bbs-like-backend/middleware"
	"bbs-like-backend/repository"
	"bbs-like-backend/service"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

// Temporary setting declaration
// TODO: move to .env file
const (
	DB_HOST     = "127.0.0.1"
	DB_PORT     = 5432
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

func handleCors(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "successful"})
}

func main() {

	route := gin.Default()
	route.Use(cors.New(middleware.CorsSetting()))
	addr := fmt.Sprintf("host=localhost user=%v password=%v dbname=%v port=%d sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	db, err := gorm.Open(postgres.Open(addr), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	pr := repository.NewPostgreSQLRepository(db)
	us := service.NewUserService(pr)
	handler.NewUserHandle(route, us)
	// route.GET("/test", HandleHello)
	// route.POST("/user", handler.UserCreate)
	// route.GET("/userlist", handler.GetUsersList)
	// route.POST("/cors", handleCors)
	// route.POST("/login", handler.Login)
	route.Run(":5050")
}
