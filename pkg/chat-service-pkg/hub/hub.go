package hub

import (
	"2019_1_TheBang/config"
)

var HubInst *Hub

type Hub struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

func InitChatHub() {
	HubInst = NewHub()
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
			config.Logger.Infof("user %v connected", client.Nickname)

		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				config.Logger.Infof("user %v disconnected", client.Nickname)

				delete(h.Clients, client)
				close(client.Send)
			}

		case message := <-h.Broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}
