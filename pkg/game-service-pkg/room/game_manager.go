package room

import (
	"2019_1_TheBang/api"
	"2019_1_TheBang/config"
	"2019_1_TheBang/config/gameconfig"
	"context"
	"time"

	pb "2019_1_TheBang/pkg/public/pbscore"
	"fmt"
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
	Move   string `json:"move" mapstructure:"move"`
}

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type EndGameInnerMsg struct {
	Msg    string `json:"msg"`
	Winner string `json:"winner"`
	Points int32  `json:"points"`
}

func (g *GameInst) Aggregation(actions ...Action) (bool, api.SocketMsg) {
	for _, action := range actions {
		ok, endGameMsg := g.AcceptAction(action)
		if ok {
			return true, endGameMsg
		}
	}

	return false, api.SocketMsg{}
}

func (g *GameInst) AcceptAction(action Action) (bool, api.SocketMsg) {
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
		return false, api.SocketMsg{}
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

	if newpos == g.Teleport && g.IsTeleport {
		g.PlayersScore[action.Player] += gameconfig.TeleportPoints
		inner := EndGameInnerMsg{
			Msg:    "Game was finished",
			Winner: action.Player,
			Points: g.PlayersScore[action.Player],
		}
		endGameMsg := api.SocketMsg{
			Type: api.GameFinish,
			Data: inner,
		}

		err := g.UpgradePoints(inner)
		if err != nil {
			config.Logger.Warn("AcceptAction", err.Error())
		}

		return true, endGameMsg
	}

	g.PlayersPos[action.Player] = newpos
	g.Map.Map[pos.X][pos.Y] = Ground

	if g.GemsCount == 0 && !g.IsTeleport {
		g.IsTeleport = true
	}

	return false, api.SocketMsg{}
}

func (g *GameInst) UpgradePoints(info EndGameInnerMsg) error {
	rw := WrapedRoom(g.Room)
	var player_id float64
	for _, player := range rw.Players {
		if player.Nickname == info.Winner {
			player_id = player.Id

			break
		}
	}

	client := pb.NewScoreUpdaterClient(gameconfig.PointsConn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ScoreRequest{
		PlayerId: player_id,
		Point:    info.Points,
	}

	res, err := client.UpdateScore(ctx, req)
	if err != nil {
		myerr := fmt.Errorf("could not update points: %v", err.Error())

		return myerr
	}

	if !res.Ok {
		myerr := fmt.Errorf("could not update points (invalid): %v", err.Error())

		return myerr
	}

	return nil
}
