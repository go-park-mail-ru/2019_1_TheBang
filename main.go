package main

import (
	"fmt"
	"math/rand"
)

type Position struct {
	X int
	Y int
}

const (
	Ground int = 0
	Wall   int = 1
	Gem    int = 2
)

type GameMap2 struct {
	Map    [][]int
	Height int
	Width  int
	Gems   int
}

func NewMap2(height, width int) GameMap2 {
	m := [][]int{}
	for i := 0; i < width; i++ {
		m = append(m, make([]int, height, height))
	}

	gamemap := GameMap2{
		Map:    m,
		Height: height,
		Width:  width,
	}

	gamemap.AddGems()
	gamemap.AddWalls()

	return gamemap
}

func (m *GameMap2) AddGems() {
	for i := 0; i < m.Height; {
		x := rand.Intn(m.Width)
		y := rand.Intn(m.Height)

		if m.Map[x][y] != Ground {
			continue
		}

		m.Map[x][y] = Gem
		m.Gems = m.Height
		i++
	}
}

func (m *GameMap2) AddWalls() {

}

func (m *GameMap2) CreateTeleport() Position {
	var x, y int

Loop:
	for {
		x = rand.Intn(m.Width)
		y = rand.Intn(m.Height)

		if m.Map[x][y] == Ground {
			break Loop
		}
	}

	return Position{
		X: x,
		Y: y,
	}
}

func main() {
	fmt.Println(NewMap2(12, 12))
}
