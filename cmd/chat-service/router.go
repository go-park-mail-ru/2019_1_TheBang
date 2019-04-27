package main

import (
	"2019_1_TheBang/pkg/public/middleware"
	"github.com/gin-gonic/gin"
)

func getChatRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CorsMiddlewareGin,
		middleware.AuthMiddlewareGin)

	hub := newHub()
	go hub.run()

	router.GET("/messages", func(c *gin.Context) {
		serveWs(hub, c.Writer, c.Request)
	})

	return router
}
