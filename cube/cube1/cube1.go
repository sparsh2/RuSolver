package cube1

import (
	"fmt"
	"solver/cube/colors"
	"solver/cube/faces"
	"solver/cube/moves"
	"solver/cube/types"
)

type Row []byte

type Cube1 struct {
	Cube       map[faces.Face][][]colors.Color
	TrackMoves bool
	Moves      []moves.Move
}

func NewCube1() types.ICube {
	mp := map[faces.Face][][]colors.Color{}
	mp[faces.FRONT] = getColorsArrary(colors.GREEN)
	mp[faces.TOP] = getColorsArrary(colors.WHITE)
	mp[faces.BACK] = getColorsArrary(colors.BLUE)
	mp[faces.LEFT] = getColorsArrary(colors.ORANGE)
	mp[faces.RIGHT] = getColorsArrary(colors.RED)
	mp[faces.BOTTOM] = getColorsArrary(colors.YELLOW)

	return &Cube1{
		Cube:       mp,
		TrackMoves: false,
		Moves:      []moves.Move{},
	}
}

func (c *Cube1) PrettyPrint() {
	for f := faces.FRONT; f <= faces.BOTTOM; f++ {
		fmt.Println(f.ToString(), ":")
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				fmt.Printf("%s ", c.Cube[f][i][j].ToString())
			}
			fmt.Println()
		}
	}
}

func (c *Cube1) GetMoves() []moves.Move {
	newMoves := make([]moves.Move, len(c.Moves))
	copy(newMoves, c.Moves)
	return newMoves
}

func (c *Cube1) SetTrackMoves(val bool) {
	c.TrackMoves = val
}

func (c *Cube1) GetCopy() types.ICube {
	newC := NewCube1().(*Cube1)
	for f := faces.FRONT; f <= faces.BOTTOM; f++ {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				newC.Cube[f][i][j] = c.Cube[f][i][j]
			}
		}
	}
	newC.TrackMoves = c.TrackMoves
	newC.Moves = make([]moves.Move, len(c.Moves))
	copy(newC.Moves, c.Moves)
	return newC
}

func (c *Cube1) ResetMoves() {
	c.Moves = []moves.Move{}
}

func (c *Cube1) Encode() string {
	str := ""

	for f := faces.FRONT; f <= faces.BOTTOM; f++ {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				// str += getColor(c.Cube[f][i][j])
				str += c.Cube[f][i][j].ToString()
			}
		}
	}
	return str
}

func (c *Cube1) Decode(hash string) {
	mp := map[faces.Face][][]colors.Color{}

	for f := faces.FRONT; f <= faces.BOTTOM; f++ {
		intf := int(f)
		mp[f] = GetEmptyFace()

		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				mp[f][i][j] = colors.Parse(string(hash[intf*9+3*i+j]))
			}
		}
	}

	c.Cube = mp
}

func (c *Cube1) ApplyMoveSequence(moves []moves.Move) {
	for _, move := range moves {
		c.ApplyMove(move)
	}
}

func (c *Cube1) ApplyMove(move moves.Move) {
	if c.TrackMoves {
		c.Moves = append(c.Moves, move)
	}

	switch move {
	case moves.RIGHT:
		c.moveRight()
	case moves.LEFT:
		c.moveLeft()
	case moves.FRONT:
		c.moveFront()
	case moves.BACK:
		c.moveBack()
	case moves.UP:
		c.moveUp()
	case moves.DOWN:
		c.moveDown()
	case moves.X:
		c.moveRight()
		c.moveLeftInverted()
		c.moveM()
		c.moveM()
		c.moveM()
	case moves.Y:
		c.moveUp()
		c.ApplyMove(moves.DOWN_INVERTED)
		c.ApplyMove(moves.E_INVERTED)
	case moves.M:
		c.moveM()
	case moves.E:
		c.moveE()

	case moves.RIGHT_INVERTED:
		c.moveRightInverted()
	case moves.LEFT_INVERTED:
		c.moveLeftInverted()
	case moves.FRONT_INVERTED:
		c.moveFront()
		c.moveFront()
		c.moveFront()
	case moves.BACK_INVERTED:
		c.moveBack()
		c.moveBack()
		c.moveBack()
	case moves.UP_INVERTED:
		c.moveUp()
		c.moveUp()
		c.moveUp()
	case moves.DOWN_INVERTED:
		c.moveDown()
		c.moveDown()
		c.moveDown()
	case moves.X_INVERTED:
		c.moveRightInverted()
		c.moveLeft()
		c.moveM()
	case moves.Y_INVERTED:
		c.ApplyMove(moves.UP_INVERTED)
		c.moveDown()
		c.moveE()
	case moves.M_INVERTED:
		c.moveM()
		c.moveM()
		c.moveM()
	case moves.E_INVERTED:
		c.moveE()
		c.moveE()
		c.moveE()

	case moves.RIGHT2:
		c.moveRight()
		c.moveRight()
	case moves.LEFT2:
		c.moveLeft()
		c.moveLeft()
	case moves.FRONT2:
		c.moveFront()
		c.moveFront()
	case moves.BACK2:
		c.moveBack()
		c.moveBack()
	case moves.UP2:
		c.moveUp()
		c.moveUp()
	case moves.DOWN2:
		c.moveDown()
		c.moveDown()
	case moves.X2:
		c.ApplyMove(moves.X_INVERTED)
		c.ApplyMove(moves.X_INVERTED)
	case moves.Y2:
		c.ApplyMove(moves.Y)
		c.ApplyMove(moves.Y)
	case moves.M2:
		c.moveM()
		c.moveM()
	case moves.E2:
		c.moveE()
		c.moveE()
	}
}
