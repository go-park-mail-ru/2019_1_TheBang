package main

import (
	"2019_1_TheBang/config"
	"2019_1_TheBang/config/gameconfig"
	"2019_1_TheBang/pkg/game-service-pkg/app"
	"2019_1_TheBang/pkg/public/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
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

func main() {
	defer config.Logger.Sync()
	config.Logger.Info(fmt.Sprintf("FrontenDest: %v", config.FrontentDst))
	config.Logger.Info(fmt.Sprintf("GAMEPORT: %v", gameconfig.GAMEPORT))

	router := setUpGameRouter()

	router.Run(":" + gameconfig.GAMEPORT)
}
