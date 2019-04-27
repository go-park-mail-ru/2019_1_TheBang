package hub

import (
	"2019_1_TheBang/config"
	"2019_1_TheBang/pkg/public/auth"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Client struct {
	Nickname string
	PhotoURL string
	Hub      *Hub
	Conn     *websocket.Conn
	Send     chan []byte
}

func playerInfoFromCookie(ctx *gin.Context) *Client {
	info, ok := auth.CheckTocken(ctx.Request)
	if !ok {
		return &Client{
			Nickname: "undef",
			PhotoURL: config.DefaultImg,
		}
	}

	return &Client{
		Nickname: info.Nickname,
		PhotoURL: info.PhotoUrl,
	}
}

func clientFromContext(ctx *gin.Context, conn *websocket.Conn) *Client {
	info := playerInfoFromCookie(ctx)
	client := &Client{
		Nickname: info.Nickname,
		PhotoURL: info.PhotoURL,
		Conn:     conn,
		Send:     make(chan []byte, 256),
	}

	return client
}
