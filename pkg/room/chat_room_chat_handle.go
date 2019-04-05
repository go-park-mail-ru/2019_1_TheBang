package room

import (
	"2019_1_TheBang/config"
	"net/http"

	"github.com/gorilla/mux"
)

func RoomChatHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomName := vars["room"]

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		config.Logger.Warnw("RoomChatHandle",
			"warn", err.Error())

		return
	}
	_ = conn
	_ = roomName

}
