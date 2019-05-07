package main

import (
	"2019_1_TheBang/config"
	"2019_1_TheBang/config/chatconfig"
)

func main() {
	err := chatconfig.DB.Ping()
	if err != nil {
		config.Logger.Fatal("Can not start connection with database")
	}

	router := getChatRouter()
	config.Logger.Fatal(router.Run(":" + config.CHATPORT))
}
