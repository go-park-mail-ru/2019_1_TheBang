package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var (
	storageAcc = CreateAccountStorage()
	storageProf = CreateProfileStorage()
	SECRET []byte
	CookieName string = "session_id"
	ServerName = "TheBang server"
)

//заглушка
func GetGreeting(r *http.Request) string{
	_, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		return "Hellow, unknown"
	}

	return fmt.Sprintf("Hellow, my friend")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	SECRET = []byte(os.Getenv("SECRET"))
	if string(SECRET) == "" {
		//toDo  вернуть строку
		SECRET = []byte("secret")
		log.Println("There is no SECRET!")
	}

	r := mux.NewRouter()

	r.HandleFunc("/", RootHandler).Methods("GET") //ok
	r.HandleFunc("/signup", SignupHandler).Methods("POST") //ok

	r.HandleFunc("/login", LogInHandler).Methods("POST") //ok
	r.HandleFunc("/logout", LogoutHandler).Methods("GET") //ok

	r.HandleFunc("/leaderbord", LeaderbordHandler).Methods("GET")

	r.HandleFunc("/profiles", ProfilesHandler).Methods("GET") //ok
	r.HandleFunc("/profiles/{id:[0-9]+}/details", ThisProfileHandler).Methods("GET") //ok
	r.HandleFunc("/profiles/{id:[0-9]+}/update", UpdateProfileInfoHandler).Methods("PUT")
	r.HandleFunc("/profiles/{id:[0-9]+}/avatar", ChangeProfileAvatarHandler).Methods("POST")

	http.ListenAndServe(":" + port, r)
}