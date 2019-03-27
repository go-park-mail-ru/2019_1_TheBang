package handlers

import (
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2019_1_TheBang/config"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}

	session, err := r.Cookie(config.CookieName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
}
