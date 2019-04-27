package app

import (
	"2019_1_TheBang/api"
	"2019_1_TheBang/config"
	"2019_1_TheBang/config/gameconfig"
	"2019_1_TheBang/pkg/game-service-pkg/room"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/manveru/faker"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  gameconfig.SocketReadBufferSize,
	WriteBufferSize: gameconfig.SocketWriteBufferSize,
	CheckOrigin: func(r *http.Request) bool {
		// todo проверять. что это локал хост
		return true
	},
}

var AppInst = NewApp()

type Game struct {
	MaxRoomsCount uint                `json:"max_rooms_count"`
	Rooms         map[uint]*room.Room `json:"rooms"`
	RoomsCount    uint                `json:"rooms_count"`
	locker        sync.Mutex
}

func NewApp() *Game {
	config.Logger.Infow("NewApp",
		"msg", "Game was created",
	)

	return &Game{
		Rooms:         make(map[uint]*room.Room),
		MaxRoomsCount: gameconfig.MaxRoomsInGame,
	}
}

func checkRoomID(id string) bool {
	ID, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	AppInst.locker.Lock()
	defer AppInst.locker.Unlock()

	if _, ok := AppInst.Rooms[uint(ID)]; !ok {
		return false
	}

	return true
}

func (g *Game) WrappedRoomsList() []room.RoomWrap {
	g.locker.Lock()
	defer g.locker.Unlock()

	wraps := []room.RoomWrap{}

	for id := range g.Rooms {
		roomNode, _ := g.Rooms[id]
		wrap := room.WrapedRoom(roomNode)

		wraps = append(wraps, wrap)
	}

	return wraps
}

func (g *Game) RoomsList() []*room.Room {
	g.locker.Lock()
	defer g.locker.Unlock()

	rooms := []*room.Room{}
	for _, room := range g.Rooms {
		rooms = append(rooms, room)
	}

	return rooms
}

func (g *Game) Room(id uint) (*room.Room, error) {
	g.locker.Lock()
	defer g.locker.Unlock()

	room, ok := g.Rooms[id]
	if !ok {
		return nil, ErrorRoomNotFound
	}

	return room, nil
}

func (g *Game) WrappedRoom(id uint) (room.RoomWrap, error) {
	g.locker.Lock()
	defer g.locker.Unlock()

	gameRoom, ok := AppInst.Rooms[id]
	if !ok {
		return room.RoomWrap{}, ErrorRoomNotFound
	}

	wrap := room.WrapedRoom(gameRoom)

	return wrap, nil
}

// Изменить способ получения id комнаты, возможны коллизии
func (g *Game) NewRoom() (room.RoomWrap, error) {
	g.locker.Lock()
	defer g.locker.Unlock()

	// todo fix in constructor
	if g.RoomsCount == gameconfig.MaxRoomsInGame {
		config.Logger.Warnw("NewRoom",
			"msg", "Rooms limit")

		return room.RoomWrap{}, ErrorMaxRoomsLimit
	}

	facker, _ := faker.New("en")
	roomName := facker.Name()

	id := g.RoomsCount + 1
	g.Rooms[id] = &room.Room{
		Id:         id,
		Name:       roomName,
		MaxPlayers: gameconfig.MaxPlayersInRoom,
		Register:   make(chan *room.Player),
		Unregister: make(chan *room.Player),
		Players:    make(map[*room.Player]interface{}),
		Broadcast:  make(chan api.SocketMsg),
		Closer:     make(chan bool, 1),
		Start:      false,
	}
	g.RoomsCount++

	// Запуск комнаты
	go g.Rooms[id].RunRoom()

	config.Logger.Infow("NewRoom",
		"msg", fmt.Sprintf("New room [id:%v, name:%v] was created", id, roomName))

	wrap := room.WrapedRoom(g.Rooms[id])
	return wrap, nil
}

func (g *Game) DeleteRoom(id uint) {
	g.locker.Lock()
	defer g.locker.Unlock()

}
