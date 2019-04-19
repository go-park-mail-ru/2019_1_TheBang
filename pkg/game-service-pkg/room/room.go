package room

import (
	"2019_1_TheBang/api"
	"2019_1_TheBang/config"
	"2019_1_TheBang/config/gameconfig"
	"fmt"
	"sync"
	"time"

	"github.com/mitchellh/mapstructure"
)

type RoomWrap struct {
	Id           uint     `json:"id"`
	Name         string   `json:"room"`
	MaxPlayers   uint     `json:"max_players"`
	PlayersCount uint     `json:"players_count"`
	Players      []Player `json:"players"`
}

func WrapedRoom(room *Room) RoomWrap {
	room.locker.Lock()
	defer room.locker.Unlock()

	palyers := []Player{}
	for player := range room.Players {
		palyers = append(palyers, *player)
	}

	wrap := RoomWrap{
		Id:           room.Id,
		Name:         room.Name,
		MaxPlayers:   room.MaxPlayers,
		PlayersCount: room.PlayersCount,
		Players:      palyers,
	}

	return wrap
}

// Подумать о том, как это все таки будет передаваться json-ом
type Room struct {
	Id           uint                    `json:"id"`
	Name         string                  `json:"room"`
	MaxPlayers   uint                    `json:"max_players"`
	PlayersCount uint                    `json:"players_count"`
	Players      map[*Player]interface{} `json:"players"`
	Register     chan *Player            `json:"-"`
	Unregister   chan *Player            `json:"-"`
	Broadcast    chan api.SocketMsg      `json:"-"`
	Closer       chan bool               `json:"-"`
	Start        bool                    `json:"-"`
	GameInst     GameInst                `json:"-"`
	locker       sync.Mutex              `json:"-"`
}

func (r *Room) Conection(player *Player) {
	if r.PlayersCount == r.MaxPlayers {
		player.Conn.WriteJSON(api.TooManyPlayersMsg)
		player.Conn.Close()
	}

	r.locker.Lock()
	r.Players[player] = nil
	r.PlayersCount++
	r.locker.Unlock()

	player.Room = r
	player.In <- api.ConectionMsg

	config.Logger.Infow("Conection",
		"msg", fmt.Sprintf("Player [id: %v, nick: %v] was connected to room [id: %v, name: %v]",
			player.Id, player.Nickname, r.Id, r.Name))
}

func (r *Room) Disconection(player *Player) {
	r.locker.Lock()
	delete(r.Players, player)
	r.PlayersCount--
	r.locker.Unlock()

	player.Conn.Close()

	config.Logger.Infow("Conection",
		"msg", fmt.Sprintf("Player [id: %v, nick: %v] was disconnected from room [id: %v, name: %v]",
			player.Id, player.Nickname, r.Id, r.Name))
}

func (r *Room) Distribution(msg api.SocketMsg) {
	for player := range r.Players {
		player.In <- msg
	}
}

func (r *Room) RunRoom() {
	config.Logger.Infow("RunRoom",
		"msg", fmt.Sprintf("Room  [id: %v name: %v] opened", r.Id, r.Name))

	defer config.Logger.Infow("RunRoom",
		"msg", fmt.Sprintf("Room [id: %v name: %v] closed", r.Id, r.Name))

	ticker := time.NewTicker(gameconfig.RoomTickTime)
	defer ticker.Stop()

	// ToDo удаление разорвавших соединение пользователей
Loop:
	for {
		select {
		case player := <-r.Register:
			r.Conection(player)

		case player := <-r.Unregister:
			r.Disconection(player)

		case msg := <-r.Broadcast:
			if r.Start == true {

				action := Action{}
				err := mapstructure.Decode(msg.Data, &action)
				if err != nil {
					config.Logger.Warnw("GameInst Run",
						"warn", fmt.Sprintf("Invalid action: %v", err.Error()))

					continue
				}

				ok := r.GameInst.Aggregation(action)
				if ok {
					r.Distribution(api.GameFinishedMsg)

					break Loop
				}
			}

		case <-ticker.C:
			if r.Start == true && r.PlayersCount == 0 {
				break Loop
			}

			if r.Start == true {
				r.Distribution(
					api.SocketMsg{
						Type: api.GameState,
						Data: r.GameInst.Snap(),
					})

				continue
			}

			r.Distribution(api.SocketMsg{
				Type: api.RoomState,
				Data: WrapedRoom(r),
			})

			if r.PlayersCount == r.MaxPlayers {
				r.Start = true
				r.GameInst = NewGame(r)

				config.Logger.Infow("GameInst Run",
					"msg", fmt.Sprintf("Game in room [id: %v, name: %v] was started", r.Id, r.Name))

				defer config.Logger.Infow("GameInst Run",
					"msg", fmt.Sprintf("Game in room [id: %v, name: %v] was finished", r.Id, r.Name))

				msg := api.SocketMsg{
					Type: api.GameStarted,
					Data: struct {
						Msg string  `json:"msg"`
						Map GameMap `json:"game_map"`
					}{
						Msg: "Game was started",
						Map: r.GameInst.Map,
					},
				}

				r.Distribution(msg)
			}

		case <-r.Closer:
			break Loop
		}
	}

	// отрубаем всех от игры
	for player := range r.Players {
		r.Disconection(player)
	}
}
