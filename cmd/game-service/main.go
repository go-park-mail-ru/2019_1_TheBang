package main

import (
	"2019_1_TheBang/config"
	"2019_1_TheBang/config/gameconfig"
	"fmt"

)

func main() {
	defer config.Logger.Sync()
	config.Logger.Info(fmt.Sprintf("FrontenDest: %v", config.FrontentDst))
	config.Logger.Info(fmt.Sprintf("GAMEPORT: %v", gameconfig.GAMEPORT))

	router := setUpGameRouter()

	router.Run(":" + gameconfig.GAMEPORT)
}
