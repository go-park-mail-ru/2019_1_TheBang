package app

import (
	"2019_1_TheBang/pkg/game-service-pkg/room"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ConnectRoomHandle(c *gin.Context) {
	id := c.Param("id")
	if ok := CheckRoomID(id); !ok {
		c.AbortWithStatus(http.StatusNotFound)

		return
	}
	param, _ := strconv.Atoi(id)
	ID := uint(param)

	AppInst.locker.Lock()
	defer AppInst.locker.Unlock()

	if ok := c.IsWebsocket(); !ok {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"message": "can not upgrade to websocket",
		})

		return
	}

	player := room.PlayerFromCtx(c, conn)
	go player.Reading()
	go player.Writing()
	AppInst.Rooms[ID].Register <- player
}
