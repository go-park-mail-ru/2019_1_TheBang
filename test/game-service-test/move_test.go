package test

import (
	"2019_1_TheBang/pkg/game-service-pkg/room"
	"testing"
)

func preGame(x, y uint) (game room.GameInst, user room.Player) {
	width := 5
	heith := 5
	user = room.Player{
		Id:       1,
		Nickname: "test",
		PhotoURL: "test",
	}

	game = room.GameInst{
		Map: room.NewMap(width, heith),
		PlayersScore: map[string]uint{
			user.Nickname: 0,
		},
		GemsCount: 0,
		PlayersPos: map[string]room.Position{
			user.Nickname: room.Position{X: x, Y: y},
		},
	}
	game.Map[0][0] = player

	return
}

func TestMoveRight(t *testing.T) {
	var x uint = 0
	var y uint = 0

	game, user := preGame(x, y)
	action := room.Action{
		Time:   "test",
		Player: user.Nickname,
		Move:   right,
	}

	game.Aggregation(action)

	if game.PlayersPos[user.Nickname].X != x+1 {
		t.Errorf("Error: faild %v move", right)
	}
}

func TestMoveLeft(t *testing.T) {
	var x uint = 0
	var y uint = 0

	game, user := preGame(x, y)
	action := room.Action{
		Time:   "test",
		Player: user.Nickname,
		Move:   left,
	}

	game.Aggregation(action)

	if game.PlayersPos[user.Nickname].X != x {
		t.Errorf("Error: faild %v move", left)
	}
}

func TestMoveUp(t *testing.T) {
	var x uint = 0
	var y uint = 0

	game, user := preGame(x, y)
	action := room.Action{
		Time:   "test",
		Player: user.Nickname,
		Move:   up,
	}

	game.Aggregation(action)

	if game.PlayersPos[user.Nickname].Y != y {
		t.Errorf("Error: faild %v move", up)
	}
}

func TestMoveDown(t *testing.T) {
	var x uint = 0
	var y uint = 0

	game, user := preGame(x, y)
	action := room.Action{
		Time:   "test",
		Player: user.Nickname,
		Move:   down,
	}

	game.Aggregation(action)

	if game.PlayersPos[user.Nickname].Y != y+1 {
		t.Errorf("Error: faild %v move", down)
	}
}
