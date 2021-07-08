package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
)

var IsDev bool

func CorsSetting() cors.Config {

	// config := cors.DefaultConfig()
	corsConf := cors.Config{
		MaxAge:                 12 * time.Hour,
		AllowBrowserExtensions: true,
	}
	if IsDev {

		corsConf.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"}
		corsConf.AllowHeaders = []string{"Authorization", "Content-Type", "Origin",
			"Connection", "Accept-Encoding", "Accept-Language", "Host"}
		corsConf.AllowOrigins = []string{"https://www.example.com"}
	} else {
		corsConf.AllowAllOrigins = true
		corsConf.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"}
		corsConf.AllowHeaders = []string{"Authorization", "Content-Type", "Upgrade", "Origin",
			"Connection", "Accept-Encoding", "Accept-Language", "Host"}
	}
	return corsConf
}
