package test

import (
	"2019_1_TheBang/pkg/main-serivce-pkg/user"
	pb "2019_1_TheBang/pkg/public/pbscore"
	"testing"
)

func TestPbScoreDummy(t *testing.T) {
	_, err := GetTESTAdminCookie()
	if err != nil {
		t.Fatal(err.Error())
	}
	defer DeleteTESTAdmin()

	// s := user.Server{}
	// _, cancel := context.WithCancel(context.Background())
	// defer cancel()

	prof, _ := user.SelectUser(testAdminNick)
	var points int32 = 10

	req := &pb.ScoreRequest{
		PlayerId: float64(prof.Id),
		Point:    points,
	}

	req.Reset()
	_ = req.String()
	_, _ = req.Descriptor()
	req.XXX_DiscardUnknown()
	_, _ = req.XXX_Marshal([]byte{}, false)
	_ = req.XXX_Unmarshal([]byte{})

	dumm := dummy{}
	req.XXX_Merge(dumm)
	_ = req.XXX_Size()
	req.XXX_DiscardUnknown()
	_ = req.GetPlayerId()
	_ = req.GetPoint()
}

type dummy struct{}

func (d dummy) Reset()               {}
func (d dummy) ProtoMessage()        {}
func (d dummy) String() (str string) { return }
