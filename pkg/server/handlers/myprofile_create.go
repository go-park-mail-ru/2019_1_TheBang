package handlers

import (
	"encoding/json"
	//"github.com/dgrijalva/jwt-go"
	"github.com/go-park-mail-ru/2019_1_TheBang/api"
	//"github.com/go-park-mail-ru/2019_1_TheBang/config"
	"github.com/go-park-mail-ru/2019_1_TheBang/pkg/server/models"
	"io/ioutil"
	"log"
	"net/http"
	//"time"
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

	//toDo загулшка, пока не будет прилетать дата
	//signup.DOB = time.Now().String()
	signup.DOB = "2018-01-01"

	_, status := models.CreateUser(&signup)
	if status != http.StatusCreated {
		w.WriteHeader(status)

		return
	}

	w.WriteHeader(http.StatusCreated)
}