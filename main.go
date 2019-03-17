package main

import (
	"github.com/go-park-mail-ru/2019_1_TheBang/config"
	"github.com/go-park-mail-ru/2019_1_TheBang/pkg/server/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {


	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.HandleFunc("/auth", handlers.LogInHandler).Methods("POST")
	r.HandleFunc("/auth", handlers.LogoutHandler).Methods("DELETE", "OPTIONS")

	r.HandleFunc("/leaderbord/{page:[0-9]+}", handlers.LeaderbordHandler).Methods("GET")

	r.HandleFunc("/user", handlers.MyProfileCreateHandler).Methods("POST")
	r.HandleFunc("/user", handlers.MyProfileInfoHandler).Methods("GET")
	r.HandleFunc("/user", handlers.MyProfileInfoUpdateHandler).Methods("PUT", "OPTIONS")

	r.HandleFunc("/user/avatar", handlers.ChangeProfileAvatarHandler).Methods("POST", "OPTIONS")
	//
	r.HandleFunc("/icon/{filename}", handlers.GetIconHandler).Methods("GET")

	log.Println(config.FrontentDst)
	log.Fatal(http.ListenAndServe(":" + config.PORT, r))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", config.FrontentDst)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		next.ServeHTTP(w, r)
	})
}