package gameconfig

import (
	"2019_1_TheBang/config"
	"os"
	"time"
)

const (
	MaxPlayersInRoom      uint = 1 // поправить и вернуть обратно
	MaxRoomsInGame        uint = 10
	RoomTickTime               = 5 * time.Second
	PlayerWritingTickTime      = 1 * time.Second
	PlayerReadingTickTime      = 1 * time.Second

	WriteDeadline = 10 * time.Second
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
		config.Logger.Warn("There is no GAMEPORT!")
		port = "8002"
	}

	return port
}
