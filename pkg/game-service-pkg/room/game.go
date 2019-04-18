package room

import (
	"2019_1_TheBang/config"
	"fmt"
)

const (
	Width  uint = 10
	Height uint = 10

	left  = "left"
	right = "right"
	up    = "up"
	down  = "down"
)

type Cell string

const (
	gem      Cell = "gem"
	player   Cell = "player"
	groung   Cell = "ground"
	box      Cell = "box"
	teleport Cell = "teleport"
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

// todo изменить на слайс слайсов
// с возможностью генерации определенного размера карты
type GameMap [Height][Width]Cell

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
