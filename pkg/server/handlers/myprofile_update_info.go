package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-park-mail-ru/2019_1_TheBang/api"
	"github.com/go-park-mail-ru/2019_1_TheBang/config"
	"github.com/go-park-mail-ru/2019_1_TheBang/pkg/server/auth"
	"github.com/go-park-mail-ru/2019_1_TheBang/pkg/server/models"
)

func MyProfileInfoUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}

	token, ok := auth.CheckTocken(r)
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		config.Logger.Infow("MyProfileInfoUpdateHandler",
			"RemoteAddr", r.RemoteAddr,
			"status", http.StatusForbidden)

		config.Logger.Warnw("MyProfileInfoUpdateHandler",
			"RemoteAddr", r.RemoteAddr,
			"warning", "User with not valid cookie")

		return
	}
	nickname, status := NicknameFromCookie(token)
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
			"status", http.StatusInternalServerError)

		return
	}
	//toDo жду от фронтами
	update.DOB = "2018-01-01"

	err = json.Unmarshal(body, &update)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		config.Logger.Warnw("MyProfileInfoUpdateHandler",
			"RemoteAddr", r.RemoteAddr,
			"status", http.StatusInternalServerError)

		return
	}

	profile, status := models.UpdateUser(nickname, update)
	if status != http.StatusOK {
		w.WriteHeader(http.StatusInternalServerError)
		config.Logger.Warnw("MyProfileInfoUpdateHandler",
			"RemoteAddr", r.RemoteAddr,
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
