package main

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	defaultImg = "default_img"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	hellowStr := GetGreeting(r)
	info := InfoText{Data: hellowStr + ", this is root!"}
	err := json.NewEncoder(w).Encode(info)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())

		return
	}
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := CreateAccount(w, r)
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

	w.WriteHeader(http.StatusCreated)
	info := InfoText{Data: "User was created!"}

	err = json.NewEncoder(w).Encode(info)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())

		return
	}
}

func CreateAccount(w http.ResponseWriter, r *http.Request) error {
	//toDo обработка ошибок
	signup := Signup{}
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &signup)
	_ = err

	user := Profile{
		Nickname: signup.Nickname,
		Name:     signup.Name,
		Surname:  signup.Surname,
		DOB:      signup.DOB,
	}
	passwd := signup.Passwd

	storageAcc.mu.Lock()
	defer storageAcc.mu.Unlock()

	if _, ok := storageAcc.data[user.Nickname]; ok {
		w.WriteHeader(http.StatusConflict)
		err := errors.New("This user already exists!")
		return err
	}

	user.Id = storageAcc.count
	user.Photo = defaultImg

	storageAcc.data[user.Nickname] = passwd
	storageProf.data[storageProf.count] = user

	storageAcc.count += 1
	storageProf.count += 1

	return nil
}

func LogInHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	login := Login{}
	//toDo обработать эту ошибку
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &login)

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
		Name:     "session_id",
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

//toDo добавить в токен id
type customClaims struct {
	Nickname string `json:"nickname"`
	jwt.StandardClaims
}

func LoginAcount(username, passwd string) (string, error) {
	storageAcc.mu.Lock()
	defer storageAcc.mu.Unlock()

	if pw, ok := storageAcc.data[username]; !ok || pw != passwd {
		err := errors.New("Wrong answer or password!")
		return "", err
	}

	claims := customClaims{
		ServerName,
		jwt.StandardClaims{
			Issuer: "theBang server",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(SECRET)
	if err != nil {
		log.Printf("Error with JWT tocken generation: %v\n", err.Error())
	}

	return ss, nil
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
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

// toDo сделать погинацию
func LeaderbordHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//toDo убрать заглушку лидерборда
	_, err := w.Write([]byte(Leaderboard))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())

		return
	}
}

func ProfilesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	hellowStr := GetGreeting(r)
	info := InfoText{Data: hellowStr + ", this is profiles!"}
	err := json.NewEncoder(w).Encode(info)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())

		return
	}
}

func ThisProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		info := InfoText{Data: "Incorrect user id!"}
		err := json.NewEncoder(w).Encode(info)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())

			return
		}

		return
	}

	storageProf.mu.Lock()
	defer storageProf.mu.Unlock()

	profile, ok := storageProf.data[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		info := InfoText{Data: "We have not this user!"}
		err := json.NewEncoder(w).Encode(info)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())

			return
		}

		return
	}

	err = json.NewEncoder(w).Encode(profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())

		return
	}
}

//toDo вместе с базами проверка на принадлежность пользователя
func UpdateProfileInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		info := InfoText{Data: "Incorrect user id!"}
		err := json.NewEncoder(w).Encode(info)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())

			return
		}

		return
	}

	if ok := CheckTocken(r); !ok {
		w.WriteHeader(http.StatusForbidden)
		info := InfoText{Data: "You can not change this profiles info!"}
		err := json.NewEncoder(w).Encode(info)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())

			return
		}

		return
	}

	storageProf.mu.Lock()
	defer storageProf.mu.Unlock()

	if _, ok := storageProf.data[id]; !ok {
		w.WriteHeader(http.StatusNotFound)
		info := InfoText{Data: "We have not this user!"}
		err := json.NewEncoder(w).Encode(info)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())

			return
		}

		return
	}

	//toDo обработка ошибок
	update := Update{}
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &update)

	updateProf := storageProf.data[id]
	updateProf.DOB = update.DOB
	updateProf.Surname = update.Surname
	updateProf.Name = update.Name

	storageProf.data[id] = updateProf

	w.WriteHeader(http.StatusAccepted)
	info := InfoText{Data: "User was updated!"}
	err = json.NewEncoder(w).Encode(info)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())

		return
	}
}

func CheckTocken(r *http.Request) bool {
	cookie, err := r.Cookie(CookieName)
	if err != nil {
		return false
	}

	tokenStr := cookie.Value

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return SECRET, nil
	})
	if err != nil {
		log.Printf("Error with check tocken: %v", err.Error())
		return false
	}

	if !token.Valid {
		log.Println("%v use faked cookie: %v", r.RemoteAddr, err)
		return false
	}

	return true
}

//toDo избавиться
func ChangeProfileAvatarHMTLHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(HTML))
}

func ChangeProfileAvatarHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		info := InfoText{Data: "Incorrect user id!"}
		err := json.NewEncoder(w).Encode(info)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())

			return
		}

		return
	}

	if _, ok := storageProf.data[id]; !ok {
		w.WriteHeader(http.StatusNotFound)
		info := InfoText{Data: "We have not this user!"}
		err := json.NewEncoder(w).Encode(info)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())

			return
		}

		return
	}

	if ok := CheckTocken(r); !ok {
		w.WriteHeader(http.StatusForbidden)
		info := InfoText{Data: "You can not change this profiles photo!"}
		err := json.NewEncoder(w).Encode(info)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())

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

	b, err := io.Copy(fileout, filein)
	if err != nil {
		_ = b // просто обрабатывать ошибку было нельзя
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("ChangeProfileAvatarHandler: ", "img was not saved on disk!")

		return
	}

	storageProf.mu.Lock()
	defer storageProf.mu.Unlock()

	updatedProf := storageProf.data[id]
	deletePhoto(updatedProf.Photo)

	updatedProf.Photo = filename
	storageProf.data[id] = updatedProf

	w.WriteHeader(http.StatusAccepted)
	info := InfoText{Data: "Photo was updated!"}
	err = json.NewEncoder(w).Encode(info)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())

		return
	}
}

func deletePhoto(filename string) {
	if filename == defaultImg {
		return
	}

	err := os.Remove("tmp/" + filename)
	if err != nil {
		log.Printf("Can not remove file tmp/%v\n", filename)
	}
}

var Leaderboard = `{
  "leaderbord": [
    {
      "position": 1,
      "nickname": "Andrey",
      "score": 10000,
      "photo": "default_img"
    },
    {
      "position": 2,
      "nickname": "Bob",
      "score": 5000,
      "photo": "default_img"
    },
    {
      "position": 3,
      "nickname": "Nick",
      "score": 2500,
      "photo": "default_img"
    },
    {
      "position": 4,
      "nickname": "Tom",
      "score": 1000,
      "photo": "default_img"
    },
    {
      "position": 5,
      "nickname": "Liza",
      "score": 10,
      "photo": "default_img"
    }
    ]
}`
