package moves

import (
	"fmt"
	"strings"
)

func ToString(moves []Move) string {
	str := ""
	for _, move := range moves {
		str += ToStringMove(move) + " "
	}
	return strings.Trim(str, " ")
}

func ToStringMove(move Move) string {
	switch move {
	case RIGHT:
		return "R"
	case LEFT:
		return "L"
	case FRONT:
		return "F"
	case BACK:
		return "B"
	case UP:
		return "U"
	case DOWN:
		return "D"
	case M:
		return "M"
	case E:
		return "E"
	case S:
		return "S"
	case X:
		return "x"
	case Y:
		return "y"
	case Z:
		return "z"

	case RIGHT2:
		return "R2"
	case LEFT2:
		return "L2"
	case FRONT2:
		return "F2"
	case BACK2:
		return "B2"
	case UP2:
		return "U2"
	case DOWN2:
		return "D2"
	case M2:
		return "M2"
	case E2:
		return "E2"
	case S2:
		return "S2"
	case X2:
		return "x2"
	case Y2:
		return "y2"
	case Z2:
		return "z2"

	case RIGHT_INVERTED:
		return "R'"
	case LEFT_INVERTED:
		return "L'"
	case FRONT_INVERTED:
		return "F'"
	case BACK_INVERTED:
		return "B'"
	case UP_INVERTED:
		return "U'"
	case DOWN_INVERTED:
		return "D'"
	case M_INVERTED:
		return "M'"
	case E_INVERTED:
		return "E'"
	case S_INVERTED:
		return "s'"
	case X_INVERTED:
		return "x'"
	case Y_INVERTED:
		return "y'"
	case Z_INVERTED:
		return "z'"

	default:
		fmt.Println(move)
		return "-"
	}
}

func Parse(str string) []Move {
	strMoves := strings.Split(strings.Trim(str, " "), " ")
	moves := []Move{}

	for _, move := range strMoves {
		moves = append(moves, ParseSingleMove(move))
	}

	return moves
}

func ParseSingleMove(str string) Move {
	switch str {
	case "R":
		return RIGHT
	case "L":
		return LEFT
	case "U":
		return UP
	case "D":
		return DOWN
	case "F":
		return FRONT
	case "B":
		return BACK
	case "x":
		return X
	case "y":
		return Y
	case "z":
		return Z
	case "M":
		return M
	case "E":
		return E
	case "S":
		return S

	case "R'":
		return RIGHT_INVERTED
	case "L'":
		return LEFT_INVERTED
	case "U'":
		return UP_INVERTED
	case "D'":
		return DOWN_INVERTED
	case "F'":
		return FRONT_INVERTED
	case "B'":
		return BACK_INVERTED
	case "x'":
		return X_INVERTED
	case "y'":
		return Y_INVERTED
	case "z'":
		return Z_INVERTED
	case "M'":
		return M_INVERTED
	case "E'":
		return E_INVERTED
	case "S'":
		return S_INVERTED

	case "R2":
		return RIGHT2
	case "L2":
		return LEFT2
	case "U2":
		return UP2
	case "D2":
		return DOWN2
	case "F2":
		return FRONT2
	case "B2":
		return BACK2
	case "x2":
		return X2
	case "y2":
		return Y2
	case "z2":
		return Z2
	case "M2":
		return M2
	case "E2":
		return E2
	case "S2":
		return S2

	default:
		fmt.Println(str)
		fmt.Println("----------------------------------------")
		return INVALID
	}
}

func InvertMoves(moves []Move) []Move {
	invertedMoves := []Move{}
	n := len(moves)

	for ind := range moves {
		invertedMoves = append(invertedMoves, InvertSingleMove(moves[n-ind-1]))
	}
	return invertedMoves
}

func InvertSingleMove(move Move) Move {
	switch move {
	case RIGHT:
		return RIGHT_INVERTED
	case LEFT:
		return LEFT_INVERTED
	case UP:
		return UP_INVERTED
	case DOWN:
		return DOWN_INVERTED
	case FRONT:
		return FRONT_INVERTED
	case BACK:
		return BACK_INVERTED
	case M:
		return M_INVERTED
	case E:
		return E_INVERTED
	case S:
		return S_INVERTED
	case X:
		return X_INVERTED
	case Y:
		return Y_INVERTED
	case Z:
		return Z_INVERTED

	case RIGHT_INVERTED:
		return RIGHT
	case LEFT_INVERTED:
		return LEFT
	case UP_INVERTED:
		return UP
	case DOWN_INVERTED:
		return DOWN
	case FRONT_INVERTED:
		return FRONT
	case BACK_INVERTED:
		return BACK
	case M_INVERTED:
		return M
	case E_INVERTED:
		return E
	case S_INVERTED:
		return S_INVERTED
	case X_INVERTED:
		return X
	case Y_INVERTED:
		return Y
	case Z_INVERTED:
		return Z

	default:
		return move
	}
}
