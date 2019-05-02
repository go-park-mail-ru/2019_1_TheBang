package main

import (
	"fmt"
	"net"
	"sync"

	"2019_1_TheBang/config"
	"2019_1_TheBang/config/mainconfig"
	"2019_1_TheBang/pkg/main-serivce-pkg/user"

	pb "2019_1_TheBang/pkg/public/pbscore"

	"google.golang.org/grpc"
)

func main() {
	wg := sync.WaitGroup{}

	defer config.Logger.Sync()
	config.Logger.Info(fmt.Sprintf("FrontenDest: %v", config.FrontentDst))
	config.Logger.Info(fmt.Sprintf("MAINPORT: %v", config.MAINPORT))
	config.Logger.Info(fmt.Sprintf("POINTSPORT: %v", config.POINTSPORT))

	err := mainconfig.DB.Ping()
	if err != nil {
		config.Logger.Fatal("Can not start connection with database")
	}

	r := setUpMainRouter()

	wg.Add(1)
	go r.Run(":" + config.MAINPORT)

	lis, err := net.Listen("tcp", ":"+config.POINTSPORT)
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
