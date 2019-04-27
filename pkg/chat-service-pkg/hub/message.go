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
	messages = []messageToClient{}

	return
}

var sqlInsertMessage = `insert into messages (author, [timestamp], message, photo_url) values ($1, $2, $3, $4);`
