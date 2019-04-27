package gameconfig

import (
	"2019_1_TheBang/config"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"time"
)

var (
	MaxPlayersInRoom uint
	MaxRoomsInGame   uint

	RoomTickTime          = 200 * time.Microsecond
	PlayerWritingTickTime = 200 * time.Microsecond
	PlayerReadingTickTime = 200 * time.Microsecond
	WriteDeadline         = 10 * time.Second

	GameWidth      int
	GameHeight     int
	TeleportPoints uint

	SocketReadBufferSize  int
	SocketWriteBufferSize int
	MaxMessageSize        int64
	InOutBuffer           int
)

var (
	GAMEPORT = getGamePort()
)

func InitGameConfig() {
	viper.AddConfigPath("config/gameconfig")
	viper.SetConfigName("gameconfig")
	err := viper.ReadInConfig()
	if err != nil {
		config.Logger.Fatal(fmt.Sprintf("Fatal error config file: %s \n", err.Error()))
	}

	MaxPlayersInRoom = uint(viper.GetInt("app.room.max_players_in_room"))
	MaxRoomsInGame = uint(viper.GetInt("app.max_rooms_in_game"))

	GameWidth = viper.GetInt("app.room.game.map.width")
	GameHeight = viper.GetInt("app.room.game.map.height")
	TeleportPoints = uint(viper.GetInt("app.room.game.teleport_points"))

	SocketReadBufferSize = viper.GetInt("networt.socket.read")
	SocketWriteBufferSize = viper.GetInt("networt.socket.write")
	MaxMessageSize = viper.GetInt64("networt.message.size")
	InOutBuffer = viper.GetInt("networt.chan_buffer")
}

func getGamePort() string {
	port := os.Getenv("GAMEPORT")
	if port == "" {
		config.Logger.Warn("There is no GAMEPORT!")
		port = "8002"
	}

	return port
}
