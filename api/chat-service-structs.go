package api

type ChatMessage struct {
	Timestamp int    `json:"timestamp"`
	Message   string `json:"message"`
	Author    string `json:"author"`
}
