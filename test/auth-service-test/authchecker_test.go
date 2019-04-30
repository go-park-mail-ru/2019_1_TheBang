package test

import (
	"2019_1_TheBang/pkg/auth-service-pkg/authchecker"
	pb "2019_1_TheBang/pkg/public/protobuf"
	"context"
	_ "google.golang.org/grpc"
	"testing"
)

func TestAuthcheckerServerSUCCESS(t *testing.T) {
	req := &pb.CookieRequest{
		JwtToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
			"eyJpZCI6MSwibmlja25hbWUiOiIxMiIsInBob3RvX3VybCI6IjNkMDEzMmI" +
			"4NTRmYmRlYWNmYzUwMmYyZmNjMWIyMjg2IiwiaXNzIjoiVGhlQmFuZyBzZXJ2ZX" +
			"IifQ.ghh_8keORscXMbHatzjTn2o8bPN0IcWdPgcAReSV7sA",
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	s := authchecker.Server{}
	res, err := s.CheckCookie(ctx, req)
	if err != nil {
		t.Errorf("TestAuthcheckerServerSUCCESS: %v", err.Error())
	}

	if !res.Valid {
		t.Error("TestAuthcheckerServerSUCCESS, valid user is invalid")
	}

	user := res.GetUser()
	if user.Nickname != "12" {
		t.Error("TestAuthcheckerServerSUCCESS, incorrect nickname")
	}
}

func TestAuthcheckerServerFAIL(t *testing.T) {
	req := &pb.CookieRequest{
		JwtToken: "invalid_token",
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := authchecker.Server{}
	res, err := s.CheckCookie(ctx, req)
	if err != nil {
		t.Errorf("TestAuthcheckerServerSUCCESS: %v", err.Error())
	}

	if res.Valid {
		t.Error("TestAuthcheckerServerSUCCESS, valid user is invalid")
	}
}
