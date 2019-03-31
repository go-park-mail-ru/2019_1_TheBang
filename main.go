package main

import (
	"net/http"

	"2019_1_TheBang/config"
	"2019_1_TheBang/pkg/server/handlers"
	"2019_1_TheBang/pkg/server/middlewares"

	"github.com/gorilla/mux"
)

func main() {
	defer config.Logger.Sync()
	err := config.DB.Ping()
	if err != nil {
		config.Logger.Fatal("Can not start connection with database")
	}

	r := mux.NewRouter()
	r.Use(middlewares.AccessLogMiddleware, middlewares.CommonMiddleware)

	r.HandleFunc("/auth", handlers.LogInHandler).Methods("POST")
	r.HandleFunc("/auth", middlewares.AuthMiddleware(handlers.LogoutHandler)).Methods("DELETE", "OPTIONS")

	r.HandleFunc("/leaderbord/{page:[0-9]+}", middlewares.AuthMiddleware(handlers.LeaderbordHandler)).Methods("GET")

	r.HandleFunc("/user", handlers.MyProfileCreateHandler).Methods("POST")
	r.HandleFunc("/user", middlewares.AuthMiddleware(handlers.MyProfileInfoHandler)).Methods("GET")
	r.HandleFunc("/user", middlewares.AuthMiddleware(handlers.MyProfileInfoUpdateHandler)).Methods("PUT", "OPTIONS")

	r.HandleFunc("/user/avatar", middlewares.AuthMiddleware(handlers.ChangeProfileAvatarHandler)).Methods("POST", "OPTIONS")

	r.HandleFunc("/icon/{filename}", handlers.GetIconHandler).Methods("GET")

	config.Logger.Infof("FrontentDst: %v", config.FrontentDst)
	config.Logger.Fatal(http.ListenAndServe(":"+config.PORT, r))
}
