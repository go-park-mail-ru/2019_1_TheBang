package room

import (
	"2019_1_TheBang/config/gameconfig"
)

const (
	left  = "left"
	right = "right"
	up    = "up"
	down  = "down"
)

type Action struct {
	Time   string `json:"time" mapstructure:"time"`
	Player string `json:"player" mapstructure:"player"`
	Move   string `json:"move" mapstructure:"move"` // left | right | up | down
}

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (g *GameInst) Aggregation(actions ...Action) bool {
	for _, action := range actions {
		ok := g.AcceptAction(action)
		if ok {
			return true
		}
	}

	return false
}

func (g *GameInst) AcceptAction(action Action) bool {
	var (
		leftBorder  int = 0
		rightBorder int = g.Map.Width - 1
		upBorder    int = g.Map.Height - 1
		downBorder  int = 0
	)

	var (
		pos Position
		ok  bool
	)

	if pos, ok = g.PlayersPos[action.Player]; !ok {
		return false
	}

	newpos := pos

	switch {
	case action.Move == left:
		if newpos.X > leftBorder {
			newpos.X--
		}

	case action.Move == right:
		if newpos.X < rightBorder {
			newpos.X++
		}

	case action.Move == down:
		if newpos.Y < upBorder {
			newpos.Y++
		}

	case action.Move == up:
		if newpos.Y > downBorder {
			newpos.Y--
		}
	}

	if g.Map.Map[newpos.X][newpos.Y] == Gem {
		g.PlayersScore[action.Player]++
		g.GemsCount--
		delete(g.GemsPosMap, newpos)

		sliceGems := []Position{}
		for gempos := range g.GemsPosMap {
			sliceGems = append(sliceGems, gempos)
		}

		g.GemsPos = sliceGems
	}

	//  заметка: ели телепорт отркылся, то не важно кто на него наступил, тому + 5 баллов
	if newpos == g.Teleport && g.IsTeleport {
		g.PlayersScore[action.Player] += gameconfig.TeleportPoints

		return true
	}

	g.PlayersPos[action.Player] = newpos
	g.Map.Map[pos.X][pos.Y] = Ground

	if g.GemsCount == 0 && !g.IsTeleport {
		g.IsTeleport = true
	}

	return false
}
