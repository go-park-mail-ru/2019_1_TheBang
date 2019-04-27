package hub

import (
	"2019_1_TheBang/config"
	"2019_1_TheBang/config/chatconfig"
)

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

func InserMessage(msg messageFromClient) bool {
	_, err := chatconfig.DB.Exec(sqlInsertMessage,
		msg.Author,
		msg.Timestamp,
		msg.Message,
		msg.PhotoURL)
	if err != nil {
		config.Logger.Warnw("InserMessage",
			"warn", err.Error())

		return false
	}

	return true
}

func GetMessages(timestamp int) (messages []messageToClient) {
	rows, err := chatconfig.DB.Query(sqlSelectMessage,
		timestamp,
		chatconfig.MessagesLimit)
	if err != nil {
		config.Logger.Warnw("GetMessages",
			"warn", err.Error())
	}
	defer rows.Close()

	messages = []messageToClient{}

	for rows.Next() {
		msg := messageToClient{}
		if err := rows.Scan(
			&msg.Author,
			&msg.Timestamp,
			&msg.Message,
			&msg.PhotoURL); err != nil {
			config.Logger.Warnw("GetMessages",
				"warn", err.Error())

			return
		}

		messages = append(messages, msg)
	}

	return
}

var sqlInsertMessage = `insert into messages (author, [timestamp], message, photo_url) values ($1, $2, $3, $4);`

var sqlSelectMessage = `select
							author,
							[timestamp],
							message,
							photo_url

						from messages
						where timestamp < $1
						order by timestamp DESC
						limit $2`
