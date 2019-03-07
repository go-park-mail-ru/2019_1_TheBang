package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

var (
	storageAcc = CreateAccountStorage()
	storageProf = CreateProfileStorage()
	defaultImg = "default_img"
	SECRET string
)

func GetGreeting(r *http.Request) string{
	cookie, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		return "Hellow, unknown"
	}

	name := cookie.Value
	return fmt.Sprintf("Hellow, %v", name)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	SECRET = os.Getenv("SECRET")
	if SECRET == "" {
		//toDo  вернуть строку
		//log.Fatal("There is no SECRET!")
	}

	r := mux.NewRouter()

	r.HandleFunc("/", RootHandler).Methods("GET") //ok
	r.HandleFunc("/signup", SignupHandler).Methods("POST") //ok

	r.HandleFunc("/login", LogInHandler).Methods("POST") //ok
	r.HandleFunc("/logout", LogoutHandler).Methods("GET") //ok

	r.HandleFunc("/leaderbord", LeaderbordHandler).Methods("GET") //ok

	r.HandleFunc("/profiles", ProfilesHandler).Methods("GET") //ok
	r.HandleFunc("/profiles/{id:[0-9]+}/details", ThisProfileHandler).Methods("GET") //ок
	r.HandleFunc("/profiles/{id:[0-9]+}/update", UpdateProfileInfoHandler).Methods("PUT") //ок
	r.HandleFunc("/profiles/{id:[0-9]+}/avatar", ChangeProfileAvatarHandler).Methods("POST")

	http.ListenAndServe(":" + port, r)
}