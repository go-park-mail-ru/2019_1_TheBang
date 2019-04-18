package room

var (
	leftBorder  uint = 0
	rightBorder uint = width - 1
	upBorder    uint = height - 1
	downBorder  uint = 0
)

type GameInst struct {
	Map          GameMap
	PlayersPos   map[string]Position
	PlayersScore map[string]uint
	GemsCount    uint // захардкодить число гемов
	MaxGemsCount uint
	Room         *Room
	Teleport     bool // наличие тп на карте
}

func NewGame(r *Room) GameInst {
	score := make(map[string]uint)
	pos := make(map[string]Position)

	for player := range r.Players {
		score[player.Nickname] = 0
		pos[player.Nickname] = Position{}
	}

	return GameInst{
		Map:          NewMap(),
		PlayersPos:   pos,
		PlayersScore: score,
		GemsCount:    1, // захардкодить число гемов
		MaxGemsCount: 1, // захардкодить число гемов
		Room:         r,
	}
}

func (g *GameInst) Snap() GameSnap {
	return GameSnap{
		Map:          g.Map,
		PlayersScore: g.PlayersScore,
		GemsCount:    g.GemsCount,
		MaxGemsCount: g.MaxGemsCount,
	}
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

	if g.Map[newpos.X][newpos.Y] == gem {
		g.PlayersScore[action.Player]++
		g.GemsCount--
	}

	if g.Map[newpos.X][newpos.Y] == teleport {
		g.Map[pos.X][pos.Y] = groung
		g.Map[newpos.X][newpos.Y] = player

		return true
	}

	g.PlayersPos[action.Player] = newpos
	g.Map[pos.X][pos.Y] = groung
	g.Map[newpos.X][newpos.Y] = player

	if g.GemsCount == 0 && !g.Teleport {
		// хардкод телепорта
		g.Map[2][2] = teleport
		g.Teleport = true
	}

	return false
}
