package hub

type messageFromClient struct {
	Author    string `json:"author"`
	Message   string `json:"message"`
	Timestamp int    `json:"timestamp"`
	PhotoURL  string `json:"photo_url"`
}

type messageToClient struct {
	Author    string `json:"author"`
	Message   string `json:"message"`
	Timestamp int    `json:"timestamp"`
	PhotoURL  string `json:"photo_url"`
}

func InserMessage(msg messageFromClient) {

}

func GetMessages(timestamp int) (messages []messageToClient) {
	return
}
