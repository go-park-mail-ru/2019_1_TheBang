package room

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"github.com/manveru/faker"
)

func newMessage() []byte {
	fack, _ := faker.New("en")
	data, _ := json.Marshal(ChatMessage{
		User:     fack.Name(),
		PhotoURL: fack.URL(),
		Time:     fmt.Sprintf("%v", time.Now()),
		Message:  fack.Sentence(5, true),
	})
	return data
}

func SendNewMsg(client *websocket.Conn) {
	ticker := time.NewTicker(3 * time.Second)
	for {
		w, err := client.NextWriter(websocket.TextMessage)
		if err != nil {
			ticker.Stop()
			break
		}

		msg := newMessage()
		w.Write(msg)
		w.Close()

		<-ticker.C
	}
}
