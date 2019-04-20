package user

import (
	pb "2019_1_TheBang/pkg/public/pbscore"
	"context"
)

type Server struct{}

func (s *Server) UpdateScore(ctx context.Context, in *pb.ScoreRequest) (*pb.ScoreResponse, error) {
	ok := UpdateUserScore(in.PlayerId, in.Point)
	return &pb.ScoreResponse{
			Ok: ok,
		},
		nil
}
