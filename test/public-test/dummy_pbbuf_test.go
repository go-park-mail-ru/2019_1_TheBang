package test

import (
	pb "2019_1_TheBang/pkg/public/protobuf"
	"testing"
)

func TestPbProtoDummy(t *testing.T) {
	cookie, err := GetTESTAdminCookie()
	if err != nil {
		t.Fatal(err.Error())
	}
	defer DeleteTESTAdmin()

	req := &pb.CookieRequest{
		JwtToken: cookie.Value,
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
	_ = req.GetJwtToken()
}
