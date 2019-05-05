package test

import (
	"2019_1_TheBang/pkg/game-service-pkg/room"
	"testing"
)

func TestMapGeneration(t *testing.T) {
	size := 12
	var count int

	m := room.NewMap(size, size)
	for _, i := range m.Map {
		for _, j := range i {
			if j == room.Gem {
				count++
			}
		}
	}

	if count != size {
		t.Errorf("Invalid gems count: expected %v, have %v", size, count)
	}

	telep := m.CreateTeleport()
	if m.Map[telep.X][telep.Y] != room.Ground {
		t.Errorf("Invalid telep position: expected %v (ground), have %v", room.Ground, m.Map[telep.X][telep.Y])
	}

	nickname := "test"
	player := &room.Player{
		Id:       100500,
		Nickname: nickname,
	}
	players := make(map[*room.Player]interface{}, 1)
	players[player] = nil

	pos, _ := m.AddPlayers(players)
	for _, i := range pos {
		if m.Map[i.X][i.Y] != room.Ground {
			t.Errorf("Invalid players position: expected %v (ground), have %v", room.Ground, m.Map[telep.X][telep.Y])
		}
	}
}
