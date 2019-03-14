package main

import (
	"crypto/md5"
	_ "crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"io"
	_ "io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	_ "os"
	_ "strconv"
	"time"
)

func MyProfileInfoHandler(w http.ResponseWriter, r *http.Request) {
	nickname, err := NicknameFromCookie(w, r)
	if err != nil {
		info := InfoText{Data: err.Error()}
		err = json.NewEncoder(w).Encode(info)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("MyProfileInfoHandler: %v\n", err.Error())

			return
		}

		return
	}

	storageProf.mu.Lock()
	defer storageProf.mu.Unlock()

	profile, ok := storageProf.data[nickname]
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("MyProfileInfoHandler: can not find user with valid token")

		return
	}

	err = json.NewEncoder(w).Encode(profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("MyProfileInfoHandler: %v\n", err.Error())

		return
	}
}

func MyProfileInfoUpdateHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "OPTIONS" {
		return
	}

	nickname, err := NicknameFromCookie(w, r)
	if err != nil {

		return
	}

	update := Update{}
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

	storageProf.mu.Lock()
	defer storageProf.mu.Unlock()

	prof := storageProf.data[nickname]

	prof.Name = update.Name
	prof.Surname = update.Surname
	prof.DOB = update.DOB

	storageProf.data[nickname] = prof

	w.WriteHeader(http.StatusAccepted)
	err = json.NewEncoder(w).Encode(prof)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())

		return
	}
}

func LogInHandler(w http.ResponseWriter, r *http.Request) {
	login := Login{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())

		return
	}
	err = json.Unmarshal(body, &login)

	token, err := LoginAcount(login.Nickname, login.Passwd)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		info := InfoText{Data: "Wrong nickname or password!"}
		err := json.NewEncoder(w).Encode(info)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())

			return
		}

		return
	}

	expiration := time.Now().Add(10 * time.Hour)
	cookie := http.Cookie{
		Name:     CookieName,
		Value:    token,
		Expires:  expiration,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)

	answer := fmt.Sprintf("User %v was login!", login.Nickname)
	info := InfoText{Data: answer}
	err = json.NewEncoder(w).Encode(info)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())

		return
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie(CookieName)

	if r.Method == "OPTIONS" {
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		info := InfoText{Data: "A not logged in user cannot log out!"}
		err := json.NewEncoder(w).Encode(info)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())

			return
		}

		return
	}

	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)

	info := InfoText{Data: "You successfully logged out!"}
	err = json.NewEncoder(w).Encode(info)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())

		return
	}
}

func LeaderbordHandler(w http.ResponseWriter, r *http.Request) {
	profSlice := []Profile{}

	storageProf.mu.Lock()
	defer storageProf.mu.Unlock()

	for _, prof := range storageProf.data {
		profSlice = append(profSlice, prof)
	}

	err := json.NewEncoder(w).Encode(profSlice)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())

		return
	}
}

func ChangeProfileAvatarHMTLHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(HTML))
}

func ChangeProfileAvatarHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	nickname, err := NicknameFromCookie(w, r)
	if err != nil {
		info := InfoText{Data: err.Error()}
		err = json.NewEncoder(w).Encode(info)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("MyProfileInfoHandler: %v\n", err.Error())

			return
		}

		return
	}

	file, header, err := r.FormFile("photo")
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		info := InfoText{Data: "image was failed in form!"}
		err := json.NewEncoder(w).Encode(info)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())

			return
		}

		return
	}
	defer file.Close()

	hasher := md5.New()
	_, err = io.Copy(hasher, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())

		return
	}
	filename := string(hasher.Sum(nil))

	filein, err := header.Open()
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		info := InfoText{Data: "image was failed in form!"}
		err := json.NewEncoder(w).Encode(info)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())

			return
		}

		return
	}
	defer filein.Close()

	fileout, err := os.OpenFile("tmp/"+filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("ChangeProfileAvatarHandler: ", "file for img was not created!")

		return
	}
	defer fileout.Close()

	_, err = io.Copy(fileout, filein)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("ChangeProfileAvatarHandler: ", "img was not saved on disk!")

		return
	}

	storageProf.mu.Lock()
	defer storageProf.mu.Unlock()

	updatedProf := storageProf.data[nickname]
	deletePhoto(updatedProf.Photo)

	updatedProf.Photo = filename
	storageProf.data[nickname] = updatedProf

	w.WriteHeader(http.StatusAccepted)
	//toDo возвращать профиль
	info := InfoText{Data: filename}
	err = json.NewEncoder(w).Encode(info)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())

		return
	}
}

func GetIconHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/jpeg")

	vars := mux.Vars(r)
	filename := vars["filename"]

	filepath := "tmp/" + filename
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("GetIconHandler: we can not read image")

		return
	}

	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("GetIconHandler: we can not write image")

		return
	}
}
