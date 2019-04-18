package main

import (
	// r "2019_1_TheBang/pkg/game-service-pkg/room"
	"fmt"
)

type GameMap2 struct {
	Map    [][]int
	Height int
	Width  int
}

func NewMap2(height, width int) GameMap2 {
	m := [][]int{}
	for i := 0; i < width; i++ {
		m = append(m, make([]int, height, height))
	}

	return GameMap2{
		Map:    m,
		Height: height,
		Width:  width,
	}
}

func main() {
	fmt.Println(NewMap2(4, 2))
}
