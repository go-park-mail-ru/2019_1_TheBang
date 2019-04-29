package test

import (
	"2019_1_TheBang/config/gameconfig"
	"testing"
)

func TestGameConfig(t *testing.T) {
	gameconfig.CONFIGPATH = "../../config/gameconfig"
	gameconfig.InitGameConfig()

	if gameconfig.InOutBuffer == 0 {
		t.Error("InOutBuffer was not init")
	}

	if gameconfig.GameHeight == 0 || gameconfig.GameWidth == 0 {
		t.Error("GameHeight/GameWidth was not init")
	}

	if gameconfig.MaxRoomsInGame == 0 || gameconfig.MaxPlayersInRoom == uint(0) {
		t.Error("MaxRoomsInGame/MaxPlayersInRoom was not init")
	}
}
