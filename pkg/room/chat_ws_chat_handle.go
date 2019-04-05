package room

import (
	"2019_1_TheBang/config"
	"2019_1_TheBang/pkg/auth"
	"2019_1_TheBang/pkg/user"
	"net/http"
)

func ServeWsChat(hub *HubChat, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		config.Logger.Warnw("ServeWsChat",
			"warn", err.Error())

		return
	}

	nickname, status := auth.NicknameFromCookie(auth.TokenFromCookie(r))
	if status == http.StatusInternalServerError {
		w.WriteHeader(http.StatusInternalServerError)
		config.Logger.Warnw("ServeWsChat",
			"warn", "Can not get nickname form valid user's cookie")

		return
	}

	profile, status := user.SelectUser(nickname)
	if status != http.StatusOK {
		w.WriteHeader(status)
		config.Logger.Warnw("ServeWsChat",
			"warn", "Can not get profile with this nickname")
	}

	client := NewClient(hub, conn, profile)
	client.Hub.Register <- client

	go client.Writing()
	go client.Reading()
}
