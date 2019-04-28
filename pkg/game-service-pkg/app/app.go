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

var AppInst *App

func InitAppInst() {
	AppInst = NewApp()
}

type App struct {
	MaxRoomsCount uint                `json:"max_rooms_count"`
	Rooms         map[uint]*room.Room `json:"rooms"`
	RoomsCount    uint                `json:"rooms_count"`
	locker        sync.Mutex
}

func NewApp() *App {
	config.Logger.Infow("NewApp",
		"msg", "App was created",
	)

	app := &App{
		Rooms:         make(map[uint]*room.Room),
		MaxRoomsCount: gameconfig.MaxRoomsInGame,
	}

	return app
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

func (a *App) WrappedRoomsList() []room.RoomWrap {
	a.locker.Lock()
	defer a.locker.Unlock()

	wraps := []room.RoomWrap{}

	for id := range a.Rooms {
		roomNode, _ := a.Rooms[id]
		wrap := room.WrapedRoom(roomNode)

		wraps = append(wraps, wrap)
	}

	return wraps
}

func (a *App) RoomsList() []*room.Room {
	a.locker.Lock()
	defer a.locker.Unlock()

	rooms := []*room.Room{}
	for _, room := range a.Rooms {
		rooms = append(rooms, room)
	}

	return rooms
}

func (a *App) Room(id uint) (*room.Room, error) {
	a.locker.Lock()
	defer a.locker.Unlock()

	room, ok := a.Rooms[id]
	if !ok {
		return nil, ErrorRoomNotFound
	}

	return room, nil
}

func (a *App) WrappedRoom(id uint) (room.RoomWrap, error) {
	a.locker.Lock()
	defer a.locker.Unlock()

	gameRoom, ok := AppInst.Rooms[id]
	if !ok {
		return room.RoomWrap{}, ErrorRoomNotFound
	}

	wrap := room.WrapedRoom(gameRoom)

	return wrap, nil
}

// Изменить способ получения id комнаты, возможны коллизии
func (a *App) NewRoom() (room.RoomWrap, error) {
	a.locker.Lock()
	defer a.locker.Unlock()

	if a.RoomsCount == a.MaxRoomsCount {
		config.Logger.Warnw("NewRoom",
			"msg", "Rooms limit")

		return room.RoomWrap{}, ErrorMaxRoomsLimit
	}

	facker, _ := faker.New("en")
	roomName := facker.Name()

	id := a.RoomsCount + 1
	a.Rooms[id] = &room.Room{
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
	a.RoomsCount++

	go a.Rooms[id].RunRoom()

	config.Logger.Infow("NewRoom",
		"msg", fmt.Sprintf("New room [id:%v, name:%v] was created", id, roomName))

	wrap := room.WrapedRoom(a.Rooms[id])
	return wrap, nil
}

func (a *App) DeleteRoom(id uint) {
	a.locker.Lock()
	defer a.locker.Unlock()

}
