package main

import (
	"2019_1_TheBang/pkg/game-service-pkg/app"
	"2019_1_TheBang/pkg/public/middleware"

	"github.com/gin-gonic/gin"
)

func setUpGameRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CorsMiddlewareGin,
		middleware.AuthMiddlewareGin,
		middleware.MetricMiddleware)

	router.GET("/room", app.RoomsListHandle)
	router.POST("/room", app.CreateRoomHandle)
	router.GET("/room/:id", app.ConnectRoomHandle)

	return router
}
