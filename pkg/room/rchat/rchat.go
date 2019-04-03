package room

type ChatMessage struct {
	User     string
	PhotoURL string
	Time     string
	Message  string
}

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}
