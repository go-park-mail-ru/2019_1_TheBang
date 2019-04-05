package room

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func ConnectRoomHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomName := vars["room"]
	w.Header().Add("Content-Type", "text/html")

	if ok := RoomsInfo.CheckRoom(roomName); !ok {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	tmpl := template.Must(template.ParseFiles("room_chat.html"))
	tmpl.Execute(w, nil)

	// ToDo переход на сокеты (заглушка)
	// room := RoomsInfo.GetRoom(roomName)
	// err := json.NewEncoder(w).Encode(room)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	config.Logger.Warnw("RoomsList",
	// 		"warn", err.Error())

	// 	return
	// }
}
