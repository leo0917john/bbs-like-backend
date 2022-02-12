package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"bbs-like-backend/pkg/handler"
	"bbs-like-backend/pkg/middleware"
	"bbs-like-backend/pkg/module/security/jwt"
	"bbs-like-backend/pkg/repository"
	"bbs-like-backend/pkg/service"
	"bbs-like-backend/pkg/version"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

const (
	sqliteDiskMode = 1
	sqliteMemNode  = 2
)

var (
	printVersion bool
)

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

	flag.BoolVar(&printVersion, "version", false, "print program build version")
	flag.Parse()
	// jwt setting
	secretKey := []byte(SECRET_KEY)
	tokenLifeTime := time.Duration(TOKEN_LIFETIME) * time.Minute

	// jwt init setting
	jwt.Init(secretKey, tokenLifeTime)
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

	if printVersion {
		fmt.Printf("%v \n", version.PrintCLIVersion())
		return
	}

	route := gin.Default()
	route.Use(cors.New(middleware.CorsSetting()))
	// db := createPostgreGormDBInstance()
	db := createSqlite3GormDBInstance(sqliteMemNode)

	pr := repository.NewPostgreSQLRepository(db)
	pr.Init()
	us := service.NewUserService(pr)
	handler.NewUserHandle(route, us)

	route.Run(":5050")
}

func createSqlite3GormDBInstance(mode int) *gorm.DB {
	db_path := ""
	switch mode {
	case sqliteDiskMode:
		db_path = "./db/gorm.db"
	case sqliteMemNode:
		db_path = ":memory:"
	}
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func createPostgreGormDBInstance() *gorm.DB {
	postgre_url := fmt.Sprintf("host=localhost user=%v password=%v dbname=%v port=%d sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	db, err := gorm.Open(postgres.Open(postgre_url), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	return db
}
