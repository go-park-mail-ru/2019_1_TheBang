package middleware

import (
	"2019_1_TheBang/pkg/public/auth"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

type urlMehtod struct {
	URL    string
	Method string
}

var ignorCheckAuth = map[urlMehtod]bool{
	urlMehtod{URL: "/auth", Method: "POST"}:    true,
	urlMehtod{URL: "/user", Method: "POST"}:    true,
	urlMehtod{URL: "/room", Method: "POST"}:    true,
	urlMehtod{URL: "/chat", Method: "GET"}:     true,
	urlMehtod{URL: "/messages", Method: "GET"}: true,
	urlMehtod{URL: "/metrics", Method: "GET"}:  true,
}

func CorsMiddlewareGin(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	c.Next()
}

func AuthMiddlewareGin(c *gin.Context) {
	check := urlMehtod{URL: c.Request.URL.Path, Method: c.Request.Method}
	m, _ := regexp.Match(`/messages*`, []byte(check.URL))
	if m == true {
		c.Next()
	}

	m, _ = regexp.Match(`/room*`, []byte(check.URL))
	if m == true && check.Method == "GET" {
		c.Next()
	}

	if ok := ignorCheckAuth[check]; !ok {
		if _, ok := auth.CheckTocken(c.Request); !ok {
			c.AbortWithStatus(http.StatusUnauthorized)

			return
		}
	}

	c.Next()
}
