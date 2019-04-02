package room

import (
	"2019_1_TheBang/config"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func ConnectRoomHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomName := vars["room"]

	if ok := RoomsInfo.CheckRoom(roomName); !ok {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	// ToDo переход на сокеты (заглушка)
	room := RoomsInfo.GetRoom(roomName)
	err := json.NewEncoder(w).Encode(room)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		config.Logger.Warnw("RoomsList",
			"warn", err.Error())

		return
	}
}
