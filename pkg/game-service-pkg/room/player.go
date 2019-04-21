package room

import (
	"2019_1_TheBang/api"
	"2019_1_TheBang/config"
	"2019_1_TheBang/config/gameconfig"
	"2019_1_TheBang/pkg/public/auth"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type UserInfo struct {
	Id       float64 `json:"id"`
	Nickname string  `json:"nickname"`
	PhotoURL string  `json:"photo_url"`
}

type Player struct {
	Id       float64            `json:"id"`
	Nickname string             `json:"nickname"`
	PhotoURL string             `json:"photo_url"`
	Conn     *websocket.Conn    `json:"-"`
	In       chan api.SocketMsg `json:"-"`
	Out      chan api.SocketMsg `json:"-"`
	Room     *Room              `json:"-"`
}

func (p *Player) Reading() {
	ticker := time.NewTicker(gameconfig.PlayerReadingTickTime)
	defer func() {
		ticker.Stop()
		p.Conn.Close()
	}()

Loop:
	for {
		msg, ok := <-p.In
		if !ok {
			break Loop
		}

		err := p.Conn.WriteJSON(msg)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func (p *Player) Writing() {
	go func() {
		for {
			msg := &api.SocketMsg{}
			err := p.Conn.ReadJSON(msg)
			if websocket.IsUnexpectedCloseError(err) {
				p.Room.Unregister <- p
				return
			}

			p.Out <- *msg
		}
	}()

Loop:
	for {
		msg, ok := <-p.Out
		if !ok {
			break Loop
		}

		p.Room.Broadcast <- msg
	}
}

func PlayerFromCtx(ctx *gin.Context, conn *websocket.Conn) *Player {
	info := playerInfoFromCookie(ctx)
	player := &Player{
		Id:       info.Id,
		Nickname: info.Nickname,
		PhotoURL: info.PhotoURL,
		Conn:     conn,
		In:       make(chan api.SocketMsg, gameconfig.InOutBuffer),
		Out:      make(chan api.SocketMsg, gameconfig.InOutBuffer),
	}

	config.Logger.Infow("PlayerFromCtx",
		"msg", fmt.Sprintf("Player [id: %v, nick: %v] was initialized", player.Id, player.Nickname))

	return player
}

func playerInfoFromCookie(ctx *gin.Context) UserInfo {
	// todo убрать дебаг
	info, ok := auth.CheckTocken(ctx.Request)
	if !ok {
		return UserInfo{
			Id:       666,
			Nickname: "debug",
			PhotoURL: "debug",
		}
	}

	return UserInfo{
		Id:       info.Id,
		Nickname: info.Nickname,
		PhotoURL: info.PhotoUrl,
	}
}
