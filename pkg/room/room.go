package room

import (
	"github.com/manveru/faker"
)

type Rooms struct {
	List  map[string]Room
	Count int
}

var RoomsInfo = Rooms{
	List: map[string]Room{
		"test": Room{
			Name:             "test",
			MaxPlayersInRoom: MaxPlayersInRoom,
		},
	},
}

type Room struct {
	Name             string `json:"name"`
	Players          int    `json:"players"`
	MaxPlayersInRoom int    `json:"max_players_in_room"`
}

func (r *Rooms) CheckRoom(name string) bool {
	if _, ok := r.List[name]; ok {
		return true
	}

	return false
}

func (r *Rooms) GetRoom(name string) Room {
	return r.List[name]
}

func (r *Rooms) RoomsList() []Room {
	list := []Room{}
	for _, room := range r.List {
		list = append(list, room)
	}

	return list
}

func (r *Rooms) NewRoom() (Room, error) {
	if r.Count == MaxRooms {
		return Room{}, ErrorMaxRooms
	}

	fack, _ := faker.New("en")

	name := fack.DomainName()
	if ok := r.CheckRoom(name); ok {
		return Room{}, ErrorConflictName
	}

	new := Room{
		Name:             name,
		MaxPlayersInRoom: MaxPlayersInRoom,
	}

	r.List[name] = new
	r.Count += 1

	return new, nil
}
