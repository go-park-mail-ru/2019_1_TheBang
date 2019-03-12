package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var (
	storageAcc    = CreateAccountStorage()
	storageProf  = CreateProfileStorage()
	SECRET      []byte
	CookieName   = "bang_token"
	ServerName         = "TheBang server"
	FrontentDst        = "localhost:3000"
	DefaultImg         = "default_img"
)

//заглушка
func GetGreeting(r *http.Request) string {
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
		SECRET = []byte("secret")
		log.Println("There is no SECRET!")
	}

	r := mux.NewRouter()

	r.HandleFunc("/", RootHandler).Methods("GET")

	r.HandleFunc("/auth", LogInHandler).Methods("POST")
	r.HandleFunc("/auth", LogoutHandler).Methods("DELETE")

	r.HandleFunc("/leaderbord", LeaderbordHandler).Methods("GET")

	r.HandleFunc("/user", MyProfileCreateHandler).Methods("POST")
	r.HandleFunc("/user", MyProfileInfoHandler).Methods("GET")
	//r.HandleFunc("/user", MyProfileInfoUpdateHandler).Methods("PUT")

	r.HandleFunc("/profiles", ProfilesHandler).Methods("GET")
	//r.HandleFunc("/profiles/{id:[0-9]+}/details", ThisProfileHandler).Methods("GET")

	http.ListenAndServe(":"+port, r)
}

var HTML = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
    <style>
        body {
            display: flex;
            align-items: center;
            justify-content: center;
        }
        form {
            width: 400px;
            height: 500px;
            background-color: lightblue;
            display: flex;
            flex-direction: column;
            padding: 50px;
            box-sizing: border-box;
        }
        form div {
            flex-grow: 13;
        }
    </style>
</head>
<body>
<form action="/profiles/0/avatar" method="post" enctype="multipart/form-data">
    <div>photo:</div>
    <input type="file" name="photo">
    <br>
    <br>
    <input type="submit" value="Upload">
</form>
</body>
</html>`
