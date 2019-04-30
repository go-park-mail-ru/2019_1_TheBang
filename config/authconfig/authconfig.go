package authconfig

import (
	"2019_1_TheBang/config"
	"os"
)

var (
	AUTHPORT = getAuhtPort()
)

func getAuhtPort() string {
	port := os.Getenv("AUTHPORT")
	if port == "" {
		config.Logger.Warn("There is no AUTH_PORT!")
		port = "50061"
	}

	return port
}
