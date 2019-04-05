package room

type Room struct {
	Name             string `json:"name"`
	Players          int    `json:"players"`
	MaxPlayersInRoom int    `json:"max_players_in_room"`
	Hub              *HubChat
}
