package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"2019_1_TheBang/api"
	"2019_1_TheBang/config"
	"2019_1_TheBang/pkg/auth"
)

func MyProfileInfoUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}

	token := auth.TokenFromCookie(r)
	nickname, status := auth.NicknameFromCookie(token)
	if status == http.StatusInternalServerError {
		w.WriteHeader(status)
		config.Logger.Warnw("MyProfileInfoUpdateHandler",
			"RemoteAddr", r.RemoteAddr,
			"status", http.StatusInternalServerError)

		return
	}

	update := api.Update{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		config.Logger.Warnw("MyProfileInfoUpdateHandler",
			"RemoteAddr", r.RemoteAddr,
			"status", http.StatusInternalServerError,
			"warn", err.Error())

		return
	}
	//toDo жду от фронта
	update.DOB = "2018-01-01"

	err = json.Unmarshal(body, &update)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		config.Logger.Warnw("MyProfileInfoUpdateHandler",
			"RemoteAddr", r.RemoteAddr,
			"status", http.StatusInternalServerError)

		return
	}

	profile, status := UpdateUser(nickname, update)
	if status != http.StatusOK {
		w.WriteHeader(http.StatusInternalServerError)
		config.Logger.Warnw("MyProfileInfoUpdateHandler",
			"RemoteAddr", r.RemoteAddr,
			"status", http.StatusInternalServerError,
			"warn", "can not update valid user's info")

		return
	}

	err = json.NewEncoder(w).Encode(profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		config.Logger.Warnw("MyProfileInfoUpdateHandler",
			"RemoteAddr", r.RemoteAddr,
			"status", http.StatusInternalServerError)

		return
	}
}
