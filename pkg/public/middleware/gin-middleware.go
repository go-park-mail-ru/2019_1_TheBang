package middleware

import (
	"2019_1_TheBang/config"
	"2019_1_TheBang/pkg/public/auth"

	"github.com/gin-gonic/gin"
)

func CorsMiddlewareGin(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", config.FrontentDst)
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	c.Next()
}

func AuthMiddlewareGin(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		return
	}

	check := urlMehtod{URL: c.Request.URL.Path, Method: c.Request.Method}
	if ok := ignorCheckAuth[check]; !ok {
		if _, ok := auth.CheckTocken(c.Request); !ok {
			return
		}
	}

	c.Next()
}
