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

func MyProfileCreateHandler(w http.ResponseWriter, r *http.Request) {
	signup := api.Signup{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())

		return
	}

	err = json.Unmarshal(body, &signup)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())

		return
	}

	profile, status := models.CreateUser(&signup)
	if status != http.StatusCreated {
		w.WriteHeader(status)

		return
	}

	claims := models.CustomClaims{
		profile.Nickname,
		jwt.StandardClaims{
			Issuer: config.ServerName,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(config.SECRET)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error with JWT tocken generation: %v\n", err.Error())

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
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())

		return
	}
}