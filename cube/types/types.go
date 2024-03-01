package types

import (
	"solver/cube/moves"
)

type ICube interface {
	ApplyMove(moves.Move)
	ApplyMoveSequence([]moves.Move)
	// ApplyMask([]colors.Color)
	Encode() string
	Decode(string)
	// Reset()
	SetTrackMoves(bool)
	PrettyPrint()
	GetCopy() ICube
	GetMoves() []moves.Move
	ResetMoves()
}
