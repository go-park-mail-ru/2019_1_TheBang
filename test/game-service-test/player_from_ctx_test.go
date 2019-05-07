package test

import (
	"2019_1_TheBang/pkg/game-service-pkg/room"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPlayerFromCtx(t *testing.T) {
	rr := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rr)
	ctx.Request = httptest.NewRequest("GET", "/dummy", nil)

	cookie, err := GetTESTAdminCookie()
	if err != nil {
		t.Error("TestPlayerFromCtx: can not get admin cookie")
	}
	defer DeleteTESTAdmin()

	ctx.SetCookie(cookie.Name, cookie.Value,
		cookie.MaxAge,
		cookie.Path,
		cookie.Domain,
		cookie.Secure,
		cookie.HttpOnly)
	player := room.PlayerFromCtx(ctx, nil)
	dummy := room.UserInfo{}

	if player.Nickname == dummy.Nickname {
		t.Error("TestPlayerFromCtx: can not get player")
	}
}
