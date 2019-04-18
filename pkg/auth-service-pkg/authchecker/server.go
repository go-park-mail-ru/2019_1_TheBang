package authchecker

import (
	"2019_1_TheBang/config"
	pb "2019_1_TheBang/pkg/public/protobuf"
	"context"
)

type Server struct{}

func (s *Server) CheckCookie(ctx context.Context, in *pb.CookieRequest) (*pb.UserInfoResponse, error) {
	config.Logger.Info("CheckCookie started")
	defer config.Logger.Info("CheckCookie finished")

	user, err := InfoFromCookie(in.JwtToken)
	if err != nil {
		return &pb.UserInfoResponse{
				Valid: false,
				User:  nil,
			},
			nil
	}

	return &pb.UserInfoResponse{
			Valid: true,
			User:  &user,
		},
		nil
}
