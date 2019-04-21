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

	"google.golang.org/grpc"
)

func setUpMainRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CorsMiddlewareGin,
		middleware.AuthMiddlewareGin)

	router.POST("/auth", login.LogInHandler)
	router.DELETE("/auth", logout.LogoutHandler)

	router.GET("/leaderbord/:page", leaderboard.LeaderbordHandler)

	router.POST("/user", user.MyProfileCreateHandler)
	router.GET("/user", user.MyProfileInfoHandler)
	router.PUT("/user", user.MyProfileInfoUpdateHandler)

	router.POST("/user/avatar", user.ChangeProfileAvatarHandler)

	router.GET("/icon/:filename", user.GetIconHandler)

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
	go r.Run(":" + mainconfig.MAINPORT)

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
