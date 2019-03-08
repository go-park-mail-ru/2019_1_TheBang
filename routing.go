package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strconv"
	"time"
)

var (
	defaultImg = "default_img"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
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
		err := CreateAccount(w, r)
		if err != nil {
			log.Println(err.Error())

			err := json.NewEncoder(w).Encode(err.Error())
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
	//toDo валидация формы
	var user = Profile{
		Nickname: r.FormValue("nickname"),
		Name: r.FormValue("name"),
		Surname: r.FormValue("surname"),
		DOB: r.FormValue("dob"),
	}
	passwd := r.FormValue("passwd")

	storageAcc.mu.Lock()
	defer storageAcc.mu.Unlock()

	if _, ok := storageAcc.data[user.Nickname]; ok {
		w.WriteHeader(http.StatusConflict)
		err := errors.New("This user already exists!")
		return err
	}

		//file, header, err := r.FormFile("photo")
		//if err != nil {
		//	err := errors.New("image was failed in form!")
		//	return err
		//}
		//defer file.Close()
		//
		//hasher := md5.New()
		//io.Copy(hasher, file)
		//filename := string(hasher.Sum(nil))
		//
		////toDo при фейле удалить созданный фаил
		////toDo если у 2 пользователей одинаковые изображение, обработка коллизий
		//
		//filein, err := header.Open()
		//if err != nil {
		//	err := errors.New("image was failed!")
		//	return err
		//}
		//defer filein.Close()
		//
		//fileout, err := os.OpenFile("tmp/" + filename, os.O_WRONLY|os.O_CREATE, 0644)
		//if err != nil {
		//	//toDo тут скорее всего 500-я ошибка
		//	w.WriteHeader(http.StatusInternalServerError)
		//	err := errors.New("image was not saved on disk!")
		//	return err
		//}
		//defer fileout.Close()
		//
		//b, err := io.Copy(fileout, filein)
		//if err != nil {
		//	_ = b // просто обрабатывать ошибку было нельзя
		//	//toDo тут скорее всего 500-я ошибка
		//	w.WriteHeader(http.StatusInternalServerError)
		//	err := errors.New("image was not saved!")
		//	return err
		//}
		//
		//user.Photo = filename

	user.Id = storageAcc.count
	user.Photo = defaultImg

	storageAcc.data[user.Nickname] = passwd
	storageProf.data[storageProf.count] = user

	storageAcc.count += 1
	storageProf.count += 1

	return nil
}

func LogInHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("nickname")
	passwd := r.FormValue("passwd")

	token, err := LoginAcount(username, passwd)
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
		Name:    "session_id",
		Value:   token,
		Expires: expiration,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)

	answer := fmt.Sprintf("User %v was login!", username)
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
			Issuer:    "theBang server",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(SECRET)
	if err != nil {
		log.Println("Error with JWT tocken generation!")
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
	hellowStr := GetGreeting(r)
	info := InfoText{Data: hellowStr + ", this is leaderbord!"}
	err := json.NewEncoder(w).Encode(info)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())

		return
	}
}

func ProfilesHandler(w http.ResponseWriter, r *http.Request) {
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

func UpdateProfileInfoHandler(w http.ResponseWriter, r *http.Request) {
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

	updateProf := Profile{
		Id: id,
		Nickname: r.FormValue("nickname"),
		Name:  r.FormValue("name"),
		Surname: r.FormValue("surname"),
		DOB: r.FormValue("DOB"),
		Photo: storageProf.data[id].Photo,
	}

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

	_ = cookie
	return true

	//jwt.Parse(cookie.Value, jwt.SigningMethodHS256)
	//cookie.Value
}


func ChangeProfileAvatarHandler(w http.ResponseWriter, r *http.Request) {

}

