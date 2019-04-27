package chatconfig

import (
	"os"
)

var CHATPORT = getPort()

func getPort() string {
	port := os.Getenv("CHATPORT")
	if port == "" {
		return "8003"
	}

	return port
}
