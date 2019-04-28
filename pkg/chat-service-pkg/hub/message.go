package hub

import (
	"2019_1_TheBang/api"
	"2019_1_TheBang/config"
	"2019_1_TheBang/config/chatconfig"
	"errors"
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

func DeleteMessage(timestamp int, author string) (err error) {
	row, err := chatconfig.DB.Exec(sqlDeleteMessage,
		timestamp,
		author)
	if err != nil {
		config.Logger.Warnw("DeleteMessage",
			"warn", err.Error())

		return err
	}

	rows, _ := row.RowsAffected()
	if rows != 1 {
		err = errors.New("Invalid message's deleted count")

		return
	}

	return
}

func EditMessage(msg api.ChatMessage) (err error) {
	row, err := chatconfig.DB.Exec(sqlEditMessage,
		msg.Message,
		msg.Author,
		msg.Timestamp)
	if err != nil {
		config.Logger.Warnw("EditMessage",
			"warn", err.Error())

		return err
	}

	rows, _ := row.RowsAffected()
	if rows != 1 {
		err = errors.New("Invalid message's deleted count")

		return
	}

	return
}

var sqlInsertMessage = `insert into messages (author, [timestamp], message, photo_url) values ($1, $2, $3, $4)`

var sqlDeleteMessage = `delete from messages where timestamp = $1 and author = $2`

var sqlEditMessage = `update messages set message = $1 where author = $2 and timestamp = $3`

var sqlSelectMessage = `select
							author,
							[timestamp],
							message,
							photo_url

						from messages
						where timestamp < $1
						order by timestamp DESC
						limit $2`
