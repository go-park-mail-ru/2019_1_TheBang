package test

import (
	"2019_1_TheBang/config"
	"2019_1_TheBang/pkg/game-service-pkg/room"
	"fmt"
)

const (
	gem      room.Cell = "gem"
	player   room.Cell = "player"
	groung   room.Cell = "ground"
	box      room.Cell = "box"
	teleport room.Cell = "teleport"

	left  = "left"
	right = "right"
	up    = "up"
	down  = "down"
)

func newmap() room.GameMap {
	config.Logger.Infow("newmap",
		"msg", fmt.Sprint("newmap was generated"))

	return room.GameMap{
		{player, groung, groung, groung, groung, groung, groung, groung, groung, groung},
		{groung, groung, groung, groung, groung, groung, groung, groung, groung, groung},
		{groung, groung, groung, groung, groung, groung, groung, groung, groung, groung},
		{groung, groung, groung, groung, groung, groung, groung, groung, groung, groung},
		{groung, groung, groung, groung, groung, groung, groung, groung, groung, groung},
		{groung, groung, groung, groung, groung, groung, groung, groung, groung, groung},
		{groung, groung, groung, groung, groung, groung, groung, groung, groung, groung},
		{groung, groung, groung, groung, groung, groung, groung, groung, groung, groung},
		{groung, groung, groung, groung, groung, groung, groung, groung, groung, groung},
		{groung, groung, groung, groung, groung, groung, groung, groung, groung, groung},
	}
}
