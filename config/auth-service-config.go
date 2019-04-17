package config

import "os"

var (
	AUTHPORT = getAuhtPort()
)

func getAuhtPort() string {
	port := os.Getenv("AUTHPORT")
	if port == "" {
		Logger.Warn("There is no AUTH_PORT!")
		port = "50051"
	}

	return port
}
