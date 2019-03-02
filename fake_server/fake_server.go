package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"sync"
	"time"
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

func GetGreeting(r *http.Request) string{
	cookie, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		return "Hellow, unknown\n"
	}

	name := cookie.Value
	return fmt.Sprintf("Hellow, %v\n", name)
}


func RootHandler(w http.ResponseWriter, r *http.Request) {
	hellowStr := GetGreeting(r)
	w.Write([]byte(hellowStr))
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

	hellowStr := GetGreeting(r)
	w.Write([]byte(hellowStr))
	w.Write([]byte("this is signup!"))
}

func LoginAcount(username, passwd string) error {
	storage.mu.Lock()
	if _, ok := storage.data[username]; !ok {
		err := errors.New("Wrong answer or password!")
		storage.mu.Unlock()
		return err
	}
	storage.mu.Unlock()

	return nil
}

func LogInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		passwd := r.FormValue("passwd")

		err := LoginAcount(username, passwd)
		if err != nil {
			w.Write([]byte(err.Error()))
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
		w.Write([]byte(answer))

		return
	}

	hellowStr := GetGreeting(r)
	w.Write([]byte(hellowStr))
	w.Write([]byte("this is login!"))
}

func LeaderbordHandler(w http.ResponseWriter, r *http.Request) {
	hellowStr := GetGreeting(r)
	w.Write([]byte(hellowStr))
	w.Write([]byte("this is leaderbord!"))
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// намеренно сначало отдаю приветствие, а затем уже убиваю печеньку!(
	hellowStr := GetGreeting(r)

	session, err := r.Cookie("session_id")
	if err == nil {
		session.Expires = time.Now().AddDate(0, 0, -1)
		http.SetCookie(w, session)
	}

	w.Write([]byte(hellowStr))
	w.Write([]byte("this is logout!"))
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	hellowStr := GetGreeting(r)
	w.Write([]byte(hellowStr))
	w.Write([]byte("this is profile!"))
}

func main() {
		r := mux.NewRouter()
		r.HandleFunc("/", RootHandler).Methods("GET")
		r.HandleFunc("/signup", SignupHandler).Methods("GET", "POST")
		r.HandleFunc("/login", LogInHandler).Methods("GET", "POST")
		r.HandleFunc("/leaderbord", LeaderbordHandler).Methods("GET")
		r.HandleFunc("/logout", LogoutHandler).Methods("GET")
		r.HandleFunc("/profile", ProfileHandler).Methods("GET")


	http.ListenAndServe(":8080", r)
}