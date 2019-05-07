package room

import "2019_1_TheBang/config/gameconfig"

type GameInst struct {
	Map          GameMap
	PlayersPos   map[string]Position
	PlayersScore map[string]int32
	GemsCount    int
	MaxGemsCount int
	Room         *Room
	IsTeleport   bool
	Teleport     Position
	GemsPosMap   map[Position]bool
	GemsPos      []Position
}

func NewGame(r *Room) GameInst {
	m := NewMap(gameconfig.GameHeight, gameconfig.GameWidth)
	pos, score := m.AddPlayers(r.Players)
	teleport := m.CreateTeleport()

	return GameInst{
		Map:          m,
		PlayersPos:   pos,
		PlayersScore: score,
		GemsCount:    m.Gems,
		MaxGemsCount: m.Gems,
		Room:         r,
		IsTeleport:   false,
		Teleport:     teleport,
		GemsPos:      m.GemsPos,
		GemsPosMap:   m.GemsPosMap,
	}
}

type GameSnap struct {
	PlayersPos   map[string]Position `json:"players_positions"`
	PlayersScore map[string]int32    `json:"players_score"`
	GemsCount    int                 `json:"gems_count"`
	MaxGemsCount int                 `json:"max_gems_count"`
	IsTeleport   bool                `json:"is_teleport"`
	Teleport     Position            `json:"teleport"`
	GemsPosMap   map[Position]bool   `json:"-"`
	GemsPos      []Position          `json:"gems_positions"`
}

func (g *GameInst) Snap() GameSnap {
	return GameSnap{
		PlayersPos:   g.PlayersPos,
		PlayersScore: g.PlayersScore,
		GemsCount:    g.GemsCount,
		MaxGemsCount: g.MaxGemsCount,
		IsTeleport:   g.IsTeleport,
		Teleport:     g.Teleport,
		GemsPos:      g.GemsPos,
	}
}
