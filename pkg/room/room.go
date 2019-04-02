package room

type Room struct {
	Name string `json:"name"`
}

var Rooms = map[string]Room{
	"test": Room{
		Name: "test",
	},
}

func RoomsList() []Room {
	list := []Room{}

	for room := range Rooms {
		append
	}
}
