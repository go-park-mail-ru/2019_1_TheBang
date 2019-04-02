package room

import (
	"2019_1_TheBang/api"
	"2019_1_TheBang/config"
	"encoding/json"
	"net/http"
)

func CreateRoomHandle(w http.ResponseWriter, r *http.Request) {
	room, err := RoomsInfo.NewRoom()
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		error := api.MyError{
			Message: err.Error(),
		}

		err := json.NewEncoder(w).Encode(error)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			config.Logger.Warnw("CreateRoomHandle",
				"warn", err.Error())

			return
		}

		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(room)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		config.Logger.Warnw("CreateRoomHandle",
			"warn", err.Error())

		return
	}
}
