package main

import (
	"net/http"

	"2019_1_TheBang/config"
	"2019_1_TheBang/pkg/leaderboard"
	"2019_1_TheBang/pkg/login"
	"2019_1_TheBang/pkg/logout"
	"2019_1_TheBang/pkg/middleware"
	"2019_1_TheBang/pkg/user"

	"github.com/gorilla/mux"
)

func main() {
	defer config.Logger.Sync()
	err := config.DB.Ping()
	if err != nil {
		config.Logger.Fatal("Can not start connection with database")
	}

	r := mux.NewRouter()
	r.Use(middleware.AccessLogMiddleware, middleware.CommonMiddleware)

	r.HandleFunc("/auth", login.LogInHandler).Methods("POST")
	r.HandleFunc("/auth", logout.LogoutHandler).Methods("DELETE", "OPTIONS")

	r.HandleFunc("/leaderbord/{page:[0-9]+}", middleware.AuthMiddleware(leaderboard.LeaderbordHandler)).Methods("GET")

	r.HandleFunc("/user", user.MyProfileCreateHandler).Methods("POST")
	r.HandleFunc("/user", middleware.AuthMiddleware(user.MyProfileInfoHandler)).Methods("GET")
	r.HandleFunc("/user", middleware.AuthMiddleware(user.MyProfileInfoUpdateHandler)).Methods("PUT", "OPTIONS")

	r.HandleFunc("/user/avatar", middleware.AuthMiddleware(user.ChangeProfileAvatarHandler)).Methods("POST", "OPTIONS")

	r.HandleFunc("/icon/{filename}", user.GetIconHandler).Methods("GET")

	config.Logger.Infof("FrontentDst: %v", config.FrontentDst)
	config.Logger.Fatal(http.ListenAndServe(":"+config.PORT, r))
}
