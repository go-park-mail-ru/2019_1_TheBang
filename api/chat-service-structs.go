package api

type ChatMessage struct {
	Timestamp int    `json:"timestamp"`
	Message   string `json:"message"`
	Author    string `json:"author"`
	Edited    bool   `json:"edited"`
	Deleted   bool   `json:"deleted"`
}
