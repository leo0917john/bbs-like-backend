package handler

import (
	"bbs-like-backend/pkg/entity"
	"bbs-like-backend/pkg/model"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	UService entity.UserService
}

func NewUserHandle(e *gin.Engine, us entity.UserService) {
	handler := &userHandler{
		UService: us,
	}
	res := handler.UService.TableCheck()
	if res != nil {
		log.Panic("table not exist")
	}
	e.GET("/test", handler.test)
	e.POST("/user", handler.UserCreateHandler)
	e.GET("/users", handler.GetUsersList)
	e.GET("/login", handler.UserLoginHandler)
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

	if _, hit := params["username"]; !hit {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "ERROR entry name"})
		return
	}
	if _, hit := params["password"]; !hit {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "ERROR entry name"})
		return
	}
	username := params["username"].(string)
	password := params["password"].(string)
	new_user := &model.User{
		Username: username,
		Password: password,
	}
	create_err := h.UService.CreateUser(*new_user)

	if create_err != nil {
		log.Println(create_err)
		c.JSON(http.StatusConflict, gin.H{
			"error": create_err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "commit"})

	}

}

func (h *userHandler) UserLoginHandler(c *gin.Context) {
	var input model.UserLoginData
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	pass := h.UService.LoginCheck(input)
	if pass {
		c.JSON(http.StatusOK, gin.H{"msg": "Login success!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Login fail!"})
	}
}

func (h *userHandler) GetUsersList(c *gin.Context) {
	tmp_slice, err := h.UService.GetUsersList("")
	if err != nil {
		log.Panicln(err)
	}
	c.JSON(http.StatusOK, gin.H{"context": tmp_slice})

}
