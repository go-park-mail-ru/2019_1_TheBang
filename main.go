package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var (
	storageAcc  = CreateAccountStorage()
	storageProf = CreateProfileStorage()
	SECRET      []byte
	CookieName  = "bang_token"
	ServerName  = "TheBang server"
	FrontentDst = "localhost:3000"
	DefaultImg  = "default_img"
)

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
	r.Use(commonMiddleware)

	r.HandleFunc("/auth", LogInHandler).Methods("POST")
	r.HandleFunc("/auth", LogoutHandler).Methods("DELETE")

	r.HandleFunc("/leaderbord", LeaderbordHandler).Methods("GET")

	r.HandleFunc("/user", MyProfileCreateHandler).Methods("POST")
	r.HandleFunc("/user", MyProfileInfoHandler).Methods("GET")
	r.HandleFunc("/user", MyProfileInfoUpdateHandler).Methods("PUT")

	//toDo зашлушка для ручного естирования
	r.HandleFunc("/user/avatar", ChangeProfileAvatarHMTLHandler).Methods("GET")
	r.HandleFunc("/user/avatar", ChangeProfileAvatarHandler).Methods("POST")

	r.HandleFunc("/icon/{filename}", GetIconHandler).Methods("GET")


	log.Fatal(http.ListenAndServe(":"+port, r))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		next.ServeHTTP(w, r)
	})
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
<form action="/user/avatar" method="post" enctype="multipart/form-data">
    <div>photo:</div>
    <input type="file" name="photo">
    <br>
    <br>
    <input type="submit" value="Upload">
</form>
</body>
</html>`
