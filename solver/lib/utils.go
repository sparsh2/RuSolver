package lib

import "solver/cube/moves"

type FirstBlock struct {
	Setup []moves.Move // to bring to std roux position
	Name  string
}

var FirstBlocksToSolve []*FirstBlock

func init() {
	FirstBlocksToSolve = []*FirstBlock{
		{
			Setup: []moves.Move{moves.Y, moves.X2},
			Name:  "white-green",
		},
		{
			Setup: []moves.Move{moves.X2},
			Name:  "white-orange",
		},
		{
			Setup: []moves.Move{moves.Y_INVERTED, moves.X2},
			Name:  "white-blue",
		},
		{
			Setup: []moves.Move{moves.Y2, moves.X2},
			Name:  "white-red",
		},
		{
			Setup: []moves.Move{moves.Y},
			Name:  "yellow-green",
		},
		{
			Setup: []moves.Move{},
			Name:  "yellow-orange",
		},
		{
			Setup: []moves.Move{moves.Y_INVERTED},
			Name:  "yellow-blue",
		},
		{
			Setup: []moves.Move{moves.Y2},
			Name:  "yellow-red",
		},
		{
			Setup: []moves.Move{moves.Y, moves.X},
			Name:  "orange-green",
		},
		{
			Setup: []moves.Move{moves.Y, moves.X_INVERTED},
			Name:  "red-green",
		},
		{
			Setup: []moves.Move{moves.X_INVERTED, moves.Y, moves.X},
			Name:  "orange-white",
		},
	}
}
