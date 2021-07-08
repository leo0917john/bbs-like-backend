package security

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID uint `json:"userID"`
	jwt.StandardClaims
}

var (
	SecretKey     []byte
	TokenLifeTime time.Duration
)

func Init(secret []byte, lifeTime time.Duration) {
	SecretKey = secret
	TokenLifeTime = lifeTime
	fmt.Println("Init : jwt init done")
}
