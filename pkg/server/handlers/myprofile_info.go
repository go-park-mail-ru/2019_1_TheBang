package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"2019_1_TheBang/config"
	"2019_1_TheBang/pkg/server/auth"
	"2019_1_TheBang/pkg/server/models"
)

func MyProfileInfoHandler(w http.ResponseWriter, r *http.Request) {
	token := auth.TokenFromCookie(r)
	nickname, status := auth.NicknameFromCookie(token)
	if status == http.StatusInternalServerError {
		w.WriteHeader(status)
		config.Logger.Warnw("MyProfileInfoHandler",
			"RemoteAddr", r.RemoteAddr,
			"status", http.StatusInternalServerError)

		return
	}

	profile, status := models.SelectUser(nickname)
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
