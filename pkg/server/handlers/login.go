package handlers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-park-mail-ru/2019_1_TheBang/api"
	"github.com/go-park-mail-ru/2019_1_TheBang/config"
	"github.com/go-park-mail-ru/2019_1_TheBang/pkg/server/models"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func LogInHandler(w http.ResponseWriter, r *http.Request) {
	login := api.Login{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("LogInHandler: %v\n", err.Error())

		return
	}

	err = json.Unmarshal(body, &login)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("LogInHandler: %v\n", err.Error())

		return
	}

	ss, status := LoginAcount(login.Nickname, login.Passwd)
	if status != http.StatusOK {
		w.WriteHeader(status)

		return
	}

	expiration := time.Now().Add(10 * time.Hour)
	cookie := http.Cookie{
		Name:     config.CookieName,
		Value:    ss,
		Expires:  expiration,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)

	prof, status := models.SelectUser(login.Nickname)
	if status != http.StatusOK {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("LogInHandler, can not search loged user, statsu: %v\n", status)

		return
	}

	err = json.NewEncoder(w).Encode(prof)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())

		return
	}

	}


func LoginAcount(username, passwd string) (ss string, status int) {
	ok := models.CheckUser(username, passwd)
	if !ok {
		return ss, http.StatusUnauthorized
	}

	claims := models.CustomClaims{
		username,
		jwt.StandardClaims{
			Issuer: config.ServerName,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(config.SECRET)
	if err != nil {
		log.Printf("Error with JWT tocken generation: %v\n", err.Error())

		return ss, http.StatusInternalServerError
	}

	return ss, http.StatusOK
}