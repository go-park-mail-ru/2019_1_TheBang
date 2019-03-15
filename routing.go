package main

import (
	"crypto/md5"
	_ "crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"io"
	_ "io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	_ "os"
	"strconv"
	_ "strconv"
	"time"
)

func MyProfileCreateHandler(w http.ResponseWriter, r *http.Request) {
	profile, err := CreateAccount(w, r)
	if err != nil {
		log.Println(err.Error())
		info := InfoText{Data: err.Error()}
		err := json.NewEncoder(w).Encode(info)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())

			return
		}

		return
	}

	claims := customClaims{
		profile.Nickname,
		jwt.StandardClaims{
			Issuer: ServerName,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(SECRET)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error with JWT tocken generation: %v\n", err.Error())

		return
	}

	expiration := time.Now().Add(10 * time.Hour)
	cookie := http.Cookie{
		Name:     CookieName,
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

func LeaderbordPageHandler(w http.ResponseWriter, r *http.Request) {
	///
	///
	///
	///
	///
	///
	///
	vars := mux.Vars(r)
	page, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	if page == 1 {
		w.Write([]byte(p1))

		return
	} else if page == 2 {
		w.Write([]byte(p2))

		return
	}
	w.WriteHeader(404)
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

var p1 = `[
    {
        "nickname": "admin1",
        "name": "admin1",
        "surname": "admin1",
        "dob": "0.0.0.0",
        "photo": "default_img",
        "score": 10000
    },{
        "nickname": "admin2",
        "name": "admin2",
        "surname": "admin2",
        "dob": "0.0.0.0",
        "photo": "default_img",
        "score": 1000
    }, {
        "nickname": "admin3",
        "name": "admin3",
        "surname": "admin3",
        "dob": "0.0.0.0",
        "photo": "default_img",
        "score": 900
    },  {
        "nickname": "admin4",
        "name": "admin4",
        "surname": "admin4",
        "dob": "0.0.0.0",
        "photo": "default_img",
        "score": 800
    },  {
        "nickname": "admin5",
        "name": "admin5",
        "surname": "admin5",
        "dob": "0.0.0.0",
        "photo": "default_img",
        "score": 700
    },  {
        "nickname": "admin6",
        "name": "admin6",
        "surname": "admin6",
        "dob": "0.0.0.0",
        "photo": "default_img",
        "score": 600
    }
]`

var p2 = `[
    {
        "nickname": "admin7",
        "name": "admin7",
        "surname": "admin7",
        "dob": "0.0.0.0",
        "photo": "default_img",
        "score": 500
    },{
        "nickname": "admin8",
        "name": "admin8",
        "surname": "admin8",
        "dob": "0.0.0.0",
        "photo": "default_img",
        "score": 400
    }, {
        "nickname": "admin9",
        "name": "admin9",
        "surname": "admin9",
        "dob": "0.0.0.0",
        "photo": "default_img",
        "score": 300
    },  {
        "nickname": "admin10",
        "name": "admin10",
        "surname": "admin10",
        "dob": "0.0.0.0",
        "photo": "default_img",
        "score": 200
    },  {
        "nickname": "admin11",
        "name": "admin11",
        "surname": "admin11",
        "dob": "0.0.0.0",
        "photo": "default_img",
        "score": 100
    },  {
        "nickname": "admin12",
        "name": "admin12",
        "surname": "admin12",
        "dob": "0.0.0.0",
        "photo": "default_img",
        "score": 90
    }
]`
