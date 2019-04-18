package main

import (
	"2019_1_TheBang/config"
	"2019_1_TheBang/pkg/auth-service-pkg/authchecker"
	pb "2019_1_TheBang/pkg/public/protobuf"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func main() {
	defer config.Logger.Sync()
	config.Logger.Info(fmt.Sprintf("AUTH_PORT: %v", config.AUTHPORT))

	lis, err := net.Listen("tcp", ":"+config.AUTHPORT)
	if err != nil {
		config.Logger.Fatalw("Listen port",
			"error:", err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterCookieCheckerServer(s, &authchecker.Server{})

	if err = s.Serve(lis); err != nil {
		config.Logger.Fatalw("serve port",
			"error:", err.Error())
	}

	config.Logger.Info(fmt.Sprint("Auth server started!"))
}
