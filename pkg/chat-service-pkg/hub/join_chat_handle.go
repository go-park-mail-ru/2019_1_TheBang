package hub

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ServeChat(chatHub *Hub, c *gin.Context) {
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

	client := clientFromContext(c, conn)
	client.Hub = chatHub
	client.Hub.Register <- client

	go client.WritePump()
	go client.ReadPump()
}
