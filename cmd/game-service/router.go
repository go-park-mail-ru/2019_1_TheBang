package main

import (
	"github.com/gin-gonic/gin"
	"2019_1_TheBang/pkg/game-service-pkg/app"
	"2019_1_TheBang/pkg/public/middleware"
)

func setUpGameRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CorsMiddlewareGin,
		middleware.AuthMiddlewareGin)

	router.GET("/room", app.RoomsListHandle)
	router.POST("/room", app.CreateRoomHandle)
	router.GET("/room/:id", app.ConnectRoomHandle)

	return router
}
