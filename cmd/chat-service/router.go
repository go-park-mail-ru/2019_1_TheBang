package main

import (
	"2019_1_TheBang/pkg/chat-service-pkg/hub"
	"2019_1_TheBang/pkg/public/middleware"
	"github.com/gin-gonic/gin"
)

func getChatRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CorsMiddlewareGin,
		middleware.AuthMiddlewareGin)

	chatHub := hub.NewHub()
	go chatHub.Run()

	router.GET("/chat", func(c *gin.Context) {
		hub.ServeChat(chatHub, c)
	})
	router.GET("/messages", hub.MessagesHandle)
	router.PUT("/message", hub.UpdateMessageHandle)
	router.DELETE("/message", hub.DeleteMessageHandle)
	router.OPTIONS("/message", func (c *gin.Context) {})

	return router
}
