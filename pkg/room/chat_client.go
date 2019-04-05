package room

import (
	"2019_1_TheBang/api"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	Nickname string
	PhotoURL string
	Hub      *HubChat
	Conn     *websocket.Conn
	Send     chan []byte
}

func NewClient(hub *HubChat, conn *websocket.Conn, prof api.Profile) *Client {
	return &Client{
		Nickname: prof.Nickname,
		PhotoURL: prof.Photo,
		Hub:      hub,
		Conn:     conn,
		Send:     make(chan []byte, ClientBufferSize),
	}
}

func (c *Client) Writing() {

}

func (c *Client) Reading() {

}
