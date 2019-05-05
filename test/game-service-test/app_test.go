package test

import (
	"2019_1_TheBang/config/gameconfig"
	"2019_1_TheBang/pkg/game-service-pkg/app"
	"testing"
)

func TestAppEasy(t *testing.T) {
	a := app.NewApp()
	rooms := a.RoomsList()

	count := len(rooms)
	expected := 0
	if count != expected {
		t.Errorf("Invalid rooms count: expected %v, have %v", expected, count)
	}

	room, _ := a.NewRoom()
	_ = room
}

func TestApp(t *testing.T) {
	gameconfig.CONFIGPATH = "../../config/gameconfig"
	gameconfig.InitGameConfig()
	app.InitAppInst()
	room, err := app.AppInst.NewRoom()
	if err != nil {
		t.Errorf("TestApp: %v", err.Error())
	}

	_, err = app.AppInst.WrappedRoom(room.Id)
	if err != nil {
		t.Errorf("TestApp: %v", err.Error())
	}

	app.AppInst.MaxRoomsCount = 1
	_, err = app.AppInst.NewRoom()
	if err == nil {
		t.Errorf("App generated more then %v rooms", gameconfig.MaxRoomsInGame)
	}
}
