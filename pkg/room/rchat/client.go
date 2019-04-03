package room

import "github.com/gorilla/websocket"

type Client struct {
	Nickname string
	PhotoURL string
	hub      *Hub
	conn     *websocket.Conn
	send     chan []byte
}
