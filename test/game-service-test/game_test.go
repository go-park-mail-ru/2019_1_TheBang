package test

import (
	"2019_1_TheBang/config/gameconfig"
	"2019_1_TheBang/pkg/game-service-pkg/app"
	"2019_1_TheBang/pkg/game-service-pkg/room"
	"testing"
)

func TestGameEasy(t *testing.T) {
	gameconfig.CONFIGPATH = "../../config/gameconfig"
	gameconfig.InitGameConfig()
	app.InitAppInst()
	r, err := app.AppInst.NewRoom()
	if err != nil {
		t.Error("Can not create new room")
	}

	g := room.NewGame(app.AppInst.Rooms[r.Id])
	snap := g.Snap()
	_ = snap
}
