package room

import "math/rand"

const (
	Ground int = 0
	Wall   int = 1
	Gem    int = 2
)

type GameMap struct {
	Map        [][]int           `json:"map"`
	Height     int               `json:"height"`
	Width      int               `json:"width"`
	Gems       int               `json:"gems"`
	GemsPosMap map[Position]bool `json:"-"`
	GemsPos    []Position        `json:"gems_positions"`
}

func NewMap(height, width int) GameMap {
	m := [][]int{}
	for i := 0; i < width; i++ {
		m = append(m, make([]int, height, height))
	}

	gamemap := GameMap{
		Map:        m,
		Height:     height,
		Width:      width,
		GemsPos:    make([]Position, width, width),
		GemsPosMap: make(map[Position]bool),
	}

	gamemap.AddGems()
	gamemap.AddWalls()

	return gamemap
}

func (m *GameMap) AddGems() {
	for i := 0; i < m.Height; {
		x := rand.Intn(m.Width)
		y := rand.Intn(m.Height)

		if m.Map[x][y] != Ground {
			continue
		}

		m.Map[x][y] = Gem
		pos := Position{X: x, Y: y}
		m.GemsPosMap[pos] = true
		m.GemsPos = append(m.GemsPos, pos)
		i++
	}

	m.Gems = m.Height
}

func (m *GameMap) AddWalls() {

}

func (m *GameMap) AddPlayers(players map[*Player]interface{}) (positions map[string]Position, score map[string]int32) {
	positions = make(map[string]Position)
	score = make(map[string]int32)
	used := make(map[Position]interface{})

	for player := range players {
	Loop:
		for {
			pos := Position{
				X: rand.Intn(m.Width),
				Y: rand.Intn(m.Height),
			}

			if _, in := used[pos]; m.Map[pos.X][pos.Y] != Ground && !in {
				continue
			}

			positions[player.Nickname] = pos
			score[player.Nickname] = 0
			used[pos] = nil

			break Loop
		}
	}

	return
}

func (m *GameMap) CreateTeleport() Position {
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
