package handler

import (
	"bbs-like-backend/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	UService model.UserService
}

func NewUserHandle(e *gin.Engine, us model.UserService) {
	handler := &userHandler{
		UService: us,
	}
	e.GET("/test", handler.test)
	// route.GET("/test", HandleHello)
	e.POST("/user", handler.UserCreateHandler)
	// route.GET("/userlist", handler.GetUsersList)
	// route.POST("/cors", handleCors)
	// route.POST("/login", handler.Login)
}

func (h *userHandler) test(c *gin.Context) {
	fmt.Println("hello")
	c.JSON(http.StatusOK, gin.H{"msg": "hello"})
}

func (h *userHandler) UserCreateHandler(c *gin.Context) {
	// var user *model.User
	var params map[string]interface{}

	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	username := params["username"].(string)
	password := params["password"].(string)
	new_user := &model.User{
		Username: username,
		Password: password,
	}
	h.UService.CreateUser(*new_user)
	// if userNotExists(db.DB, username) {
	// 	user = &model.User{
	// 		Username: username,
	// 		Password: password,
	// 	}
	// } else {
	// 	c.JSON(http.StatusBadRequest, gin.H{"msg": "Username exist"})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"msg": "commit"})
}
