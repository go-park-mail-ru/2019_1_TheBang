package handlers

import (
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-park-mail-ru/2019_1_TheBang/config"
	"github.com/go-park-mail-ru/2019_1_TheBang/models"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func MyProfileCreateHandler(w http.ResponseWriter, r *http.Request) {
	profile, err := CreateAccount(w, r)
	if err != nil {
		log.Println(err.Error())
		info := models.InfoText{Data: err.Error()}
		err := json.NewEncoder(w).Encode(info)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())

			return
		}

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

func CreateAccount(w http.ResponseWriter, r *http.Request) (prof models.Profile, err error) {
	signup := models.Signup{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())

		return prof, err
	}

	err = json.Unmarshal(body, &signup)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())

		return prof, err
	}

	prof = models.Profile{
		Nickname: signup.Nickname,
		Name:     signup.Name,
		Surname:  signup.Surname,
		DOB:      signup.DOB,
	}
	passwd := signup.Passwd

	if _, ok := storageAcc.data[prof.Nickname]; ok {
		w.WriteHeader(http.StatusConflict)
		err := errors.New("This user already exists!")

		return prof, err
	}

	prof.Photo = config.DefaultImg

	storageAcc.data[prof.Nickname] = passwd
	storageProf.data[prof.Nickname] = prof

	return prof, nil
}