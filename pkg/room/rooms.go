package room

import (
	"github.com/manveru/faker"
)

var RoomsInfo = Rooms{
	List: map[string]Room{
		"test": Room{
			Name:             "test",
			MaxPlayersInRoom: MaxPlayersInRoom,
			Hub:              NewHubChat(),
		},
	},
}

type Rooms struct {
	List  map[string]Room
	Count int
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
		Hub:              NewHubChat(),
	}

	r.List[name] = new
	r.Count += 1

	return new, nil
}
