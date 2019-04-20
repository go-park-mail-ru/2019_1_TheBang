package main

import (
	"fmt"
	"net"
	"sync"

	"github.com/gin-gonic/gin"

	"2019_1_TheBang/config"
	"2019_1_TheBang/config/mainconfig"
	"2019_1_TheBang/pkg/main-serivce-pkg/leaderboard"
	"2019_1_TheBang/pkg/main-serivce-pkg/login"
	"2019_1_TheBang/pkg/main-serivce-pkg/logout"
	"2019_1_TheBang/pkg/main-serivce-pkg/user"
	"2019_1_TheBang/pkg/public/middleware"

	pb "2019_1_TheBang/pkg/public/pbscore"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func setUpRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.AccessLogMiddleware,
		middleware.CommonMiddleware,
		middleware.AuthMiddleware)

	r.HandleFunc("/leaderbord/{page:[0-9]+}", leaderboard.LeaderbordHandler).Methods("GET")

	r.HandleFunc("/user", user.MyProfileCreateHandler).Methods("POST")
	r.HandleFunc("/user", user.MyProfileInfoHandler).Methods("GET")
	r.HandleFunc("/user", user.MyProfileInfoUpdateHandler).Methods("PUT", "OPTIONS")

	r.HandleFunc("/user/avatar", user.ChangeProfileAvatarHandler).Methods("POST", "OPTIONS")

	r.HandleFunc("/icon/{filename}", user.GetIconHandler).Methods("GET")

	return r
}

func setUpMainRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CorsMiddlewareGin,
		middleware.AuthMiddlewareGin)

	router.POST("/auth", login.LogInHandler)
	router.DELETE("/auth", logout.LogoutHandler)

	// router.GET("/leaderbord/{:page}", leaderboard.LeaderbordHandler)

	// router.GET("/room", app.RoomsListHandle)
	// router.POST("/room", app.CreateRoomHandle)
	// router.GET("/room/:id", app.ConnectRoomHandle)

	return router
}

func main() {
	wg := sync.WaitGroup{}

	defer config.Logger.Sync()
	config.Logger.Info(fmt.Sprintf("FrontenDest: %v", config.FrontentDst))
	config.Logger.Info(fmt.Sprintf("MAINPORT: %v", mainconfig.MAINPORT))
	config.Logger.Info(fmt.Sprintf("POINTSPORT: %v", mainconfig.POINTSPORT))

	err := mainconfig.DB.Ping()
	if err != nil {
		config.Logger.Fatal("Can not start connection with database")
	}

	r := setUpMainRouter()

	wg.Add(1)
	// go http.ListenAndServe(":"+mainconfig.MAINPORT, r)
	go r.Run(":" + mainconfig.MAINPORT)

	fmt.Println("HERE")

	lis, err := net.Listen("tcp", ":"+mainconfig.POINTSPORT)
	if err != nil {
		config.Logger.Fatalw("Listen port",
			"error:", err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterScoreUpdaterServer(s, &user.Server{})

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err = s.Serve(lis); err != nil {
			config.Logger.Fatalw("serve port",
				"error:", err.Error())
		}
	}()

	config.Logger.Info(fmt.Sprint("ScoreSserver started!"))

	wg.Wait()
}
