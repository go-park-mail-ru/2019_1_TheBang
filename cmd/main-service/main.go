package main

import (
	"fmt"
	"net/http"

	"2019_1_TheBang/config"
	"2019_1_TheBang/config/mainconfig"
	"2019_1_TheBang/pkg/main-serivce-pkg/leaderboard"
	"2019_1_TheBang/pkg/main-serivce-pkg/login"
	"2019_1_TheBang/pkg/main-serivce-pkg/logout"
	"2019_1_TheBang/pkg/main-serivce-pkg/user"
	"2019_1_TheBang/pkg/public/middleware"

	"github.com/gorilla/mux"
)

func main() {
	defer config.Logger.Sync()
	config.Logger.Info(fmt.Sprintf("FrontenDest: %v", config.FrontentDst))
	config.Logger.Info(fmt.Sprintf("PORT: %v", mainconfig.MAINPORT))

	err := mainconfig.DB.Ping()
	if err != nil {
		config.Logger.Fatal("Can not start connection with database")
	}

	r := mux.NewRouter()
	r.Use(middleware.AccessLogMiddleware,
		middleware.CommonMiddleware,
		middleware.AuthMiddleware)

	r.HandleFunc("/auth", login.LogInHandler).Methods("POST")
	r.HandleFunc("/auth", logout.LogoutHandler).Methods("DELETE", "OPTIONS")

	r.HandleFunc("/leaderbord/{page:[0-9]+}", leaderboard.LeaderbordHandler).Methods("GET")

	r.HandleFunc("/user", user.MyProfileCreateHandler).Methods("POST")
	r.HandleFunc("/user", user.MyProfileInfoHandler).Methods("GET")
	r.HandleFunc("/user", user.MyProfileInfoUpdateHandler).Methods("PUT", "OPTIONS")

	r.HandleFunc("/user/avatar", user.ChangeProfileAvatarHandler).Methods("POST", "OPTIONS")

	r.HandleFunc("/icon/{filename}", user.GetIconHandler).Methods("GET")

	config.Logger.Fatal(http.ListenAndServe(":"+mainconfig.MAINPORT, r))
}
