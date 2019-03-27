package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-park-mail-ru/2019_1_TheBang/config"
	"github.com/go-park-mail-ru/2019_1_TheBang/pkg/server/auth"
	"github.com/go-park-mail-ru/2019_1_TheBang/pkg/server/models"
)

func MyProfileInfoHandler(w http.ResponseWriter, r *http.Request) {
	token, ok := auth.CheckTocken(r)
	if !ok {
		w.WriteHeader(http.StatusForbidden)

		return
	}
	nickname, status := NicknameFromCookie(token)
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

func NicknameFromCookie(token *jwt.Token) (nickname string, status int) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		nickname = claims["nickname"].(string)
	} else {
		status = http.StatusInternalServerError
		config.Logger.Warnw("NicknameFromCookie",
			"warn", "Error with parsing token's claims")

		return nickname, status
	}

	return nickname, http.StatusOK
}
