package main

import (
	"net/http"

	"2019_1_TheBang/config"
	"2019_1_TheBang/pkg/chat"
	"2019_1_TheBang/pkg/leaderboard"
	"2019_1_TheBang/pkg/login"
	"2019_1_TheBang/pkg/logout"
	"2019_1_TheBang/pkg/middleware"
	"2019_1_TheBang/pkg/room"
	"2019_1_TheBang/pkg/user"

	"github.com/gorilla/mux"
)

func main() {
	defer config.Logger.Sync()
	err := config.DB.Ping()
	if err != nil {
		config.Logger.Fatal("Can not start connection with database")
	}

	hub := chat.NewHub()
	go hub.Run()

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

	r.HandleFunc("/rooms", room.RoomsListHandle).Methods("GET")
	r.HandleFunc("/rooms", room.CreateRoomHandle).Methods("POST")
	r.HandleFunc("/rooms/{room}", room.ConnectRoomHandle).Methods("GET")
	r.HandleFunc("/rooms/{room}/chat", room.RoomChatHandle)

	r.HandleFunc("/", chat.ServeHome).Methods("GET")
	r.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		chat.ServeWs(hub, w, r)
	})

	config.Logger.Infof("FrontentDst: %v", config.FrontentDst)
	config.Logger.Fatal(http.ListenAndServe(":"+config.PORT, r))
}
