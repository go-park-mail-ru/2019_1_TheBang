package room

import (
	"2019_1_TheBang/config"
	"encoding/json"
	"net/http"
)

func RoomsListHandle(w http.ResponseWriter, r *http.Request) {
	roomsList := RoomsInfo.RoomsList()
	err := json.NewEncoder(w).Encode(roomsList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		config.Logger.Warnw("RoomsList",
			"warn", err.Error())

		return
	}
}
