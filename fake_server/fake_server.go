package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"sync"
)

type AccountStorage struct {
	data map[string]string
	mu sync.Mutex
}

func CreateStorage() AccountStorage {
	acc := AccountStorage{}
	acc.data = make(map[string]string)

	return acc
}

var (
	storage = CreateStorage()
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header)
	w.Write([]byte("this is root!"))
}

func CreateAccount(username, passwd string) error {
	storage.mu.Lock()
	if _, ok := storage.data[username]; ok {
		err := errors.New("This user already exists!")
		storage.mu.Unlock()
		return err
	}

	storage.data[username] = passwd

	storage.mu.Unlock()

	return nil
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		passwd := r.FormValue("passwd")

		err := CreateAccount(username, passwd)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		answer := fmt.Sprintf("User %v was created!", username)
		w.Write([]byte(answer))

		return
	}

	w.Write([]byte("this is signup!"))
}

func LogInHandler(w http.ResponseWriter, r *http.Request) {


	w.Write([]byte("this is login!"))
}

func main() {
		r := mux.NewRouter()
		r.HandleFunc("/", RootHandler).Methods("GET")
		r.HandleFunc("/signup", SignupHandler).Methods("GET", "POST")
		r.HandleFunc("/login", LogInHandler).Methods("GET", "POST")


	http.ListenAndServe(":8080", r)
}