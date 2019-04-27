package middleware

import (
	"2019_1_TheBang/config"
	"2019_1_TheBang/pkg/public/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type urlMehtod struct {
	URL    string
	Method string
}

var ignorCheckAuth = map[urlMehtod]bool{
	urlMehtod{URL: "/auth", Method: "POST"}: true,
	urlMehtod{URL: "/user", Method: "POST"}: true,
	urlMehtod{URL: "/room", Method: "POST"}: true,
	urlMehtod{URL: "/oauth/vk/authorize", Method: "GET"}: true,
	urlMehtod{URL: "/oauth/vk/connect", Method: "GET"}: true,
}

func CorsMiddlewareGin(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", config.FrontentDst)
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	c.Next()
}

func AuthMiddlewareGin(c *gin.Context) {
	check := urlMehtod{URL: c.Request.URL.Path, Method: c.Request.Method}
	if ok := ignorCheckAuth[check]; !ok {
		if _, ok := auth.CheckTocken(c.Request); !ok {
			return
		}
	}

	c.Next()
}
