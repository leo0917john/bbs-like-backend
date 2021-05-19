package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserCreate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "hello user"})
}
