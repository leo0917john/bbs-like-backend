package handler

import (
	"bbs-like-backend/db"
	"bbs-like-backend/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// var cnt int

func UserCreate(c *gin.Context) {
	// var params map[string]interface{}

	// err := c.ShouldBindJSON(&params)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	// 	return
	// }

	// username := params["username"].(string)
	// password := params["password"].(string)
	// resp := fmt.Sprintf("user:%s pwd:%s", username, password)
	// fmt.Printf("user:%s \t pwd:%s\n", username, password)
	// c.JSON(http.StatusOK, gin.H{"msg": resp})
	var user *model.User
	var params map[string]interface{}

	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	username := params["username"].(string)
	password := params["password"].(string)
	if userExists(db.DB, username) {
		user = &model.User{Username: username, Password: password}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Username exist"})
		return
	}

	tx := db.DB.Begin()
	if db.DB != nil {
		tx.Create(user)
	} else {
		db.Show()
		fmt.Println("db == nil")
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"msg": "commit"})
}

func GetList(c *gin.Context) {
	type Username struct {
		Username string
	}
	var users []Username
	db.DB.Model(&model.User{}).Select("Username").Find(&users)
	c.JSON(http.StatusOK, gin.H{"context": users})
}

func userExists(db *gorm.DB, username string) bool {
	// sqlStmt := `SELECT username FROM userinfo WHERE username = ?`
	// err := db.QueryRow(sqlStmt, username).Scan(&username)
	var users []model.User
	db.Find(&users, "Username = ?", username)
	// if err != nil {
	// 	if err != sql.ErrNoRows {
	// 		// a real error happened! you should change your function return
	// 		// to "(bool, error)" and return "false, err" here
	// 		log.Print(err)
	// 	}

	// 	return false
	// }
	return len(users) == 0
}
