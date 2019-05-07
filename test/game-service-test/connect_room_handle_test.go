package test

import (
	"2019_1_TheBang/config"
	"2019_1_TheBang/config/gameconfig"
	"2019_1_TheBang/pkg/game-service-pkg/app"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func TestConnectRoomHandle(t *testing.T) {
	gameconfig.CONFIGPATH = "../../config/gameconfig"
	gameconfig.InitGameConfig()
	app.InitAppInst()
	r, _ := app.AppInst.NewRoom()

	router := gin.Default()
	router.GET("/room/:id", app.ConnectRoomHandle)
	go router.Run(":" + config.GAMEPORT)

	path := "ws://127.0.0.1:" + config.GAMEPORT + "/room/" + strconv.Itoa(int(r.Id))
	ws, _, err := websocket.DefaultDialer.Dial(path, nil)
	if err != nil {
		t.Fatalf("%v", err.Error())
	}
	defer ws.Close()

	time.Sleep(1 * time.Second)
}
