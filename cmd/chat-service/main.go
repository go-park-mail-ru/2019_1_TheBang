package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"os"
)

func getPort() string {
	port := os.Getenv("CHATPORT")
	if port == "" {
		return "8003"
	}

	return port
}

func main() {
	r := gin.New()

	flag.Parse()
	hub := newHub()
	go hub.run()

	r.GET("/messages", func(c *gin.Context) {
		serveWs(hub, c.Writer, c.Request)
	})

	port := getPort()
	r.Run(":" + port)
}
