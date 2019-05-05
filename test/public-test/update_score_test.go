package test

import (
	"2019_1_TheBang/pkg/main-serivce-pkg/user"
	pb "2019_1_TheBang/pkg/public/pbscore"
	"context"
	"testing"
)

func TestScoreUpdate(t *testing.T) {
	_, err := GetTESTAdminCookie()
	if err != nil {
		t.Fatal(err.Error())
	}
	defer DeleteTESTAdmin()

	s := user.Server{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	prof, _ := user.SelectUser(testAdminNick)
	var points int32 = 10

	req := &pb.ScoreRequest{
		PlayerId: float64(prof.Id),
		Point:    points,
	}

	res, err := s.UpdateScore(ctx, req)
	if err != nil {
		t.Errorf("TestScoreUpdate: %v", err.Error())
	}

	if !res.GetOk() {
		t.Errorf("TestScoreUpdate (invalid): %v", err.Error())
	}

	newprof, _ := user.SelectUser(testAdminNick)
	if prof.Score == newprof.Score {
		t.Error("TestScoreUpdate, score was not updated")
	}
}
