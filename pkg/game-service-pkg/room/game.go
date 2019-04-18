package room

import (
	"2019_1_TheBang/config"
	"fmt"
)

const (
	width  uint = 10
	height uint = 10

	left  = "left"
	right = "right"
	up    = "up"
	down  = "down"
)

type cell string

const (
	gem      cell = "gem"
	player   cell = "player"
	groung   cell = "ground"
	box      cell = "box"
	teleport cell = "teleport"
)

type Action struct {
	Time   string `json:"time" mapstructure:"time"`
	Player string `json:"player" mapstructure:"player"`
	Move   string `json:"move" mapstructure:"move"` // left | right | up | down
}

type Position struct {
	X uint
	Y uint
}

type GameSnap struct {
	Map          GameMap         `json:"map"`
	PlayersScore map[string]uint `json:"players_score"`
	GemsCount    uint            `json:"gems_count"`
	MaxGemsCount uint            `json:"max_gems_count"`
}

type GameMap [height][width]cell

func NewMap() GameMap {
	config.Logger.Infow("NewMap",
		"msg", fmt.Sprint("NewMap was generated"))

	return GameMap{
		{player, groung, groung, groung, groung, groung, groung, groung, groung, groung},
		{groung, groung, groung, groung, groung, groung, groung, groung, groung, groung},
		{groung, groung, groung, groung, groung, groung, groung, groung, groung, groung},
		{groung, groung, groung, groung, groung, groung, groung, groung, groung, groung},
		{groung, groung, groung, groung, groung, groung, groung, groung, groung, groung},
		{groung, groung, groung, groung, groung, groung, groung, groung, groung, groung},
		{groung, groung, groung, groung, groung, groung, groung, groung, groung, groung},
		{groung, groung, groung, groung, groung, groung, groung, groung, groung, groung},
		{groung, groung, groung, groung, groung, groung, groung, groung, groung, groung},
		{gem, groung, groung, groung, groung, groung, groung, groung, groung, groung}, // захадкожены гемы
	}
}
