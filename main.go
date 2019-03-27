package main

import (
	"net/http"

	"github.com/go-park-mail-ru/2019_1_TheBang/config"
	"github.com/go-park-mail-ru/2019_1_TheBang/pkg/server/handlers"
	"github.com/gorilla/mux"
)

func main() {
	defer config.Logger.Sync()

	r := mux.NewRouter()
	r.Use(accessLogMiddleware, commonMiddleware)

	r.HandleFunc("/auth", handlers.LogInHandler).Methods("POST")
	r.HandleFunc("/auth", handlers.LogoutHandler).Methods("DELETE", "OPTIONS")

	r.HandleFunc("/leaderbord/{page:[0-9]+}", handlers.LeaderbordHandler).Methods("GET")

	r.HandleFunc("/user", handlers.MyProfileCreateHandler).Methods("POST")
	r.HandleFunc("/user", handlers.MyProfileInfoHandler).Methods("GET")
	r.HandleFunc("/user", handlers.MyProfileInfoUpdateHandler).Methods("PUT", "OPTIONS")

	r.HandleFunc("/user/avatar", handlers.ChangeProfileAvatarHandler).Methods("POST", "OPTIONS")

	r.HandleFunc("/icon/{filename}", handlers.GetIconHandler).Methods("GET")

	config.Logger.Infof("FrontentDst: %v", config.FrontentDst)
	config.Logger.Fatal(http.ListenAndServe(":"+config.PORT, r))
}
