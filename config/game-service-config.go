package config

import (
	"os"
	"time"
)

const (
	MaxPlayersInRoom uint = 1 // поправить и вернуть обратно
	MaxRoomsInGame   uint = 10
	RoomTickTime          = 5 * time.Second
	// GameTickTime               = 1 * time.Second // fps стоит обсуждений
	PlayerWritingTickTime = 1 * time.Second
	PlayerReadingTickTime = 1 * time.Second

	WriteDeadline = 10 * time.Second
	// ReadingWait = 10 * time.Second
)

var (
	SocketReadBufferSize        = 1024
	SocketWriteBufferSize       = 1024
	MaxMessageSize        int64 = 512
	InOutBuffer                 = 10
)

var (
	GAMEPORT = getGamePort()
)

func getGamePort() string {
	port := os.Getenv("GAMEPORT")
	if port == "" {
		Logger.Warn("There is no GAMEPORT!")
		port = "8002"
	}
	return port
}
