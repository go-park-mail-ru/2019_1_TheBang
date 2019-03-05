package main

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

type Profile struct {
	Id int `json:"user_id, string"`
	Nickname string
	Name string
	Surname string
	DOB string //toDo нужно заменить на time.Time
	Photo string
}

type InfoText struct {
	Data string
}

// toDO заменить на бд
type AccountStorage struct {
	data map[string]string
	mu sync.Mutex
	count int // костыль для id
}

// toDO заменить на бд
type ProfileStorage struct {
	data map[int]Profile
	mu sync.Mutex
	count int // костыль для id
}

// toDO заменить на бд
func CreateAccountStorage() AccountStorage {
	acc := AccountStorage{}
	acc.data = make(map[string]string)

	return acc
}

// toDO заменить на бд
func CreateProfileStorage() ProfileStorage {
	prof := ProfileStorage{}
	prof.data = make(map[int]Profile)

	return prof
}

// toDO заменить на бд
var (
	storageAcc = CreateAccountStorage()
	storageProf = CreateProfileStorage()
	defaultImg = "default_img"
)

func InfoTextToJson(data string) []byte {
	infotext := InfoText{Data: data}
	result, _ := json.Marshal(&infotext) // намеренное игнорирование ошибки
	return result
}

func GetGreeting(r *http.Request) string{
	cookie, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		return "Hellow, unknown"
	}

	name := cookie.Value
	return fmt.Sprintf("Hellow, %v", name)
}


func RootHandler(w http.ResponseWriter, r *http.Request) {
	hellowStr := GetGreeting(r)
	dataJson := InfoTextToJson(hellowStr + ", this is root!")
	w.Write(dataJson)
}

func CreateAccount(w http.ResponseWriter, r *http.Request) error {
	user := Profile{
		Nickname: r.FormValue("nickname"),
		Name:  r.FormValue("name"),
		Surname: r.FormValue("surname"),
		DOB: r.FormValue("DOB"),
	}

	passwd := r.FormValue("passwd")

	storageAcc.mu.Lock()
	defer storageAcc.mu.Unlock()

	if _, ok := storageAcc.data[user.Nickname]; ok {
		err := errors.New("This user already exists!")
		return err
	}

	// toDo сделать ограничение по размеру
	//toDo привести код в порядок
	withPhoto := r.FormValue("with_photo")
	if withPhoto == "yes" {
		file, header, err := r.FormFile("photo")
		if err != nil {
			err := errors.New("image was failed in form!")
			return err
		}
		defer file.Close()

		hasher := md5.New()
		io.Copy(hasher, file)
		filename := string(hasher.Sum(nil))

		//toDo при фейле удалить созданный фаил
		//toDo при сборке из консоил изменить путь
		//toDo если у 2 пользователей одинаковые изображение, обработка коллизий

		filein, err := header.Open()
		if err != nil {
			err := errors.New("image was failed!")
			return err
		}
		defer filein.Close()

		fileout, err := os.OpenFile("fake_server/tmp/" + filename, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			//toDo тут скорее всего 500-я ошибка
			w.WriteHeader(http.StatusInternalServerError)
			err := errors.New("image was not saved on disk!")
			return err
		}
		defer fileout.Close()

		b, err := io.Copy(fileout, filein)
		if err != nil {
			_ = b // просто обрабатывать ошибку было нельзя
			//toDo тут скорее всего 500-я ошибка
			w.WriteHeader(http.StatusInternalServerError)
			err := errors.New("image was not saved!")
			return err
		}

		user.Photo = filename
	} else {
		user.Photo = defaultImg
	}

	user.Id = storageAcc.count

	storageAcc.data[user.Nickname] = passwd
	storageProf.data[storageProf.count] = user

	storageAcc.count += 1
	storageProf.count += 1

	return nil
}

//toDo заменить тип DOB на time.Time
//toDo добавить возможность загрузки картинки
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := CreateAccount(w, r)

		if err != nil {
			dataJson := InfoTextToJson(err.Error())
			w.Write(dataJson)
			return
		}

		w.WriteHeader(http.StatusCreated)
		dataJson := InfoTextToJson("User as created!")
		w.Write(dataJson)

		return
	}

	hellowStr := GetGreeting(r)
	dataJson := InfoTextToJson(hellowStr + ", this is signup!")
	w.Write(dataJson)
}

func LoginAcount(username, passwd string) error {
	storageAcc.mu.Lock()
	defer storageAcc.mu.Unlock()

	if _, ok := storageAcc.data[username]; !ok {
		err := errors.New("Wrong answer or password!")
		return err
	}

	return nil
}

//toDo более сложную обработку cookie
func LogInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("nickname")
		passwd := r.FormValue("passwd")

		err := LoginAcount(username, passwd)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			dataJson := InfoTextToJson("Wrong nickname or passwd!")
			w.Write(dataJson)

			return
		}

		expiration := time.Now().Add(10 * time.Hour)
		cookie := http.Cookie{
			Name:    "session_id",
			Value:   username,
			Expires: expiration,
		}

		http.SetCookie(w, &cookie)

		answer := fmt.Sprintf("User %v was login!", username)
		dataJson := InfoTextToJson(answer + ", this is login!")
		w.Write(dataJson)
		return
	}

	hellowStr := GetGreeting(r)
	dataJson := InfoTextToJson(hellowStr + ", this is login!")
	w.Write(dataJson)
}

func LeaderbordHandler(w http.ResponseWriter, r *http.Request) {
	hellowStr := GetGreeting(r)
	dataJson := InfoTextToJson(hellowStr + ", this is leaderbord!")
	w.Write(dataJson)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// намеренно сначало отдаю приветствие, а затем уже убиваю печеньку!(
	hellowStr := GetGreeting(r)

	session, err := r.Cookie("session_id")
	if err == nil {
		session.Expires = time.Now().AddDate(0, 0, -1)
		http.SetCookie(w, session)
	}

	dataJson := InfoTextToJson(hellowStr + ", you successfully logout!")
	w.Write(dataJson)
}

func ProfilesHandler(w http.ResponseWriter, r *http.Request) {
	hellowStr := GetGreeting(r)
	dataJson := InfoTextToJson(hellowStr + ", this is profiles!")
	w.Write(dataJson)
}

func ThisProfileHandler(w http.ResponseWriter, r *http.Request) {
	//toDo реализовать проверку прав на изменение профиля + доп логика
	//toDo добавить возможность редактирования изображения

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		dataJson := InfoTextToJson("Wrong user id! " + err.Error())
		w.Write(dataJson)
		return
	}

	storageProf.mu.Lock()
	defer storageProf.mu.Unlock()

	profile, ok := storageProf.data[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		dataJson := InfoTextToJson("We have not this user!")
		w.Write(dataJson)
		return
	}

	if r.Method == http.MethodPut {
		updateProf := Profile{
			Nickname: r.FormValue("nickname"),
			Name:  r.FormValue("name"),
			Surname: r.FormValue("surname"),
			DOB: r.FormValue("DOB"),
		}

		storageProf.data[id] = updateProf

		w.WriteHeader(http.StatusAccepted)
		dataJson := InfoTextToJson("Userinfo was updated!")
		w.Write(dataJson)
		return
	}

	jsonPorf, _ := json.Marshal(&profile) //ошибку намеренно не обрабатываем
	w.Write([]byte(jsonPorf))
}

func main() {
		r := mux.NewRouter()
		r.HandleFunc("/", RootHandler).Methods("GET")
		r.HandleFunc("/signup", SignupHandler).Methods("GET", "POST")

		r.HandleFunc("/login", LogInHandler).Methods("GET", "POST")
		r.HandleFunc("/logout", LogoutHandler).Methods("GET")

		r.HandleFunc("/leaderbord", LeaderbordHandler).Methods("GET")

		r.HandleFunc("/profiles", ProfilesHandler).Methods("GET")
		r.HandleFunc("/profiles/{id:[0-9]+}", ThisProfileHandler).Methods("GET", "PUT")


	http.ListenAndServe(":8080", r)
}