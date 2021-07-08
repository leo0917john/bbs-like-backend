package handler

import (
	"bbs-like-backend/db"
	"bbs-like-backend/lib/security"
	"bbs-like-backend/model"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	var user model.User
	db.DB.Where("Username = ?", input.Username).First(&user)
	if input.Password != user.Password {
		c.JSON(http.StatusUnauthorized, errors.New("incorrect email or password"))
		return
	}
	if newToken, err := sign(user.ID); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		// update JWT Token
		c.Header("Authorization", newToken)
		// allow CORS
		c.Header("Access-Control-Expose-Headers", "Authorization")
		c.JSON(http.StatusOK, map[string]interface{}{"userID": user.ID})

	}
}

func sign(userID uint) (string, error) {
	claims := security.Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(security.TokenLifeTime).Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(security.SecretKey)
}
