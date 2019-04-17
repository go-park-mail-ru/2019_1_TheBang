package user

import (
	"encoding/json"
	"log"
	"net/http"

	"2019_1_TheBang/config"
)

func MyProfileInfoHandler(w http.ResponseWriter, r *http.Request) {
	token := TokenFromCookie(r)
	info, status := InfoFromCookie(token)
	if status == http.StatusInternalServerError {
		w.WriteHeader(status)
		config.Logger.Warnw("MyProfileInfoHandler",
			"RemoteAddr", r.RemoteAddr,
			"status", http.StatusInternalServerError)

		return
	}

	profile, status := SelectUser(info.Nickname)
	if status != http.StatusOK {
		w.WriteHeader(status)

		return
	}

	err := json.NewEncoder(w).Encode(profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())

		return
	}
}
