package room

import "2019_1_TheBang/api"

var GameStartedMsg = api.SocketMsg{
	Type: api.GameStarted,
	Data: struct {
		Msg string  `json:"msg"`
		Map GameMap `json:"game_map"`
	}{
		Msg: "Game was started",
	},
}

var GameFinishedMsg = api.SocketMsg{
	Type: api.GameFinish,
	Data: struct {
		Msg string `json:"msg"`
	}{
		Msg: "Game was finished",
	},
}
