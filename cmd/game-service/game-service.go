package main

import (
	"2019_1_TheBang/config"
	"2019_1_TheBang/config/gameconfig"
	"2019_1_TheBang/pkg/game-service-pkg/app"
	_ "2019_1_TheBang/pkg/game-service-pkg/gamemonitoring"
	"fmt"
)

func main() {
	gameconfig.InitGameConfig()
	app.InitAppInst()

	defer config.Logger.Sync()
	config.Logger.Info(fmt.Sprintf("FrontenDest: %v", config.FrontentDst))
	config.Logger.Info(fmt.Sprintf("GAMEPORT: %v", config.GAMEPORT))

	router := setUpGameRouter()

	router.Run(":" + config.GAMEPORT)
}
