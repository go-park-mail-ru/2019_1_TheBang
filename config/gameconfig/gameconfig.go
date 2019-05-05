package gameconfig

import (
	"2019_1_TheBang/config"
	"fmt"
	"time"

	"github.com/spf13/viper"
)

var (
	MaxPlayersInRoom uint
	MaxRoomsInGame   uint

	RoomTickTime          = 20 * time.Millisecond
	PlayerWritingTickTime = 20 * time.Millisecond
	PlayerReadingTickTime = 20 * time.Millisecond
	// RoomTickTime          = 1 * time.Second
	// PlayerWritingTickTime = 1 * time.Second
	// PlayerReadingTickTime = 1 * time.Second
	WriteDeadline = 10 * time.Second

	GameWidth      int
	GameHeight     int
	TeleportPoints int32

	SocketReadBufferSize  int
	SocketWriteBufferSize int
	MaxMessageSize        int64
	InOutBuffer           int
)

var (
	CONFIGPATH string = "config/gameconfig"
	CONFIGNAME        = "gameconfig"
)

func InitGameConfig() {
	viper.AddConfigPath(CONFIGPATH)
	viper.SetConfigName(CONFIGNAME)
	err := viper.ReadInConfig()
	if err != nil {
		config.Logger.Fatal(fmt.Sprintf("Fatal error config file: %s \n", err.Error()))
	}

	MaxPlayersInRoom = uint(viper.GetInt("app.room.max_players_in_room"))
	MaxRoomsInGame = uint(viper.GetInt("app.max_rooms_in_game"))

	GameWidth = viper.GetInt("app.room.game.map.width")
	GameHeight = viper.GetInt("app.room.game.map.height")
	TeleportPoints = int32(viper.GetInt("app.room.game.teleport_points"))

	SocketReadBufferSize = viper.GetInt("networt.socket.read")
	SocketWriteBufferSize = viper.GetInt("networt.socket.write")
	MaxMessageSize = viper.GetInt64("networt.message.size")
	InOutBuffer = viper.GetInt("networt.chan_buffer")
}
