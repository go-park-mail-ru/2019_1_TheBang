package main

import (
	"2019_1_TheBang/config"
	"2019_1_TheBang/config/chatconfig"
)

func main() {
	router := getChatRouter()
	config.Logger.Fatal(router.Run(":" + chatconfig.CHATPORT))
}
