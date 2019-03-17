package handlers

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2019_1_TheBang/api"
	"github.com/go-park-mail-ru/2019_1_TheBang/pkg/server/auth"
	"github.com/go-park-mail-ru/2019_1_TheBang/pkg/server/models"
	"io/ioutil"
	"log"
	"net/http"
)

func MyProfileInfoUpdateHandler(w http.ResponseWriter, r *http.Request) {
	token, ok := auth.CheckTocken(r)
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		log.Println("User with not valid cookie")

		return
	}
	nickname, status := NicknameFromCookie(token)
	if status == http.StatusInternalServerError {
		w.WriteHeader(status)

		return
	}

	update := api.Update{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())

		return
	}

	err = json.Unmarshal(body, &update)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())

		return
	}

	profile, status := models.UpdateUser(nickname, update)
	if status != http.StatusOK {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("MyProfileInfoUpdateHandler: can not update valid user's info")

		return
	}

	err = json.NewEncoder(w).Encode(profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("MyProfileInfoUpdateHandler: %v\n", err.Error())

		return
	}
}
