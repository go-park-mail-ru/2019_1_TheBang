package room

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

//  todo ок заменить на финиш (имя переменной)
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

		// up и down инвертированы
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
	}

	if g.Map.Map[newpos.X][newpos.Y] == teleport {
		g.Map.Map[pos.X][pos.Y] = Ground
		g.Map.Map[newpos.X][newpos.Y] = player

		return true
	}

	g.PlayersPos[action.Player] = newpos
	g.Map.Map[pos.X][pos.Y] = Ground
	g.Map.Map[newpos.X][newpos.Y] = player

	if g.GemsCount == 0 && !g.Teleport {
		// хардкод телепорта
		g.Map[2][2] = teleport
		g.Teleport = true
	}

	return false
}
