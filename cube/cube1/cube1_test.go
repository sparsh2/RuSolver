package cube1

import (
	"fmt"
	"solver/cube/faces"
	"solver/cube/moves"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCube1_TestBasicMoves(t *testing.T) {
	movesToTest := []moves.Move{
		moves.RIGHT,
		moves.RIGHT_INVERTED,
		moves.RIGHT2,

		moves.LEFT,
		moves.LEFT2,
		moves.LEFT_INVERTED,

		moves.BACK,
		moves.BACK2,
		moves.BACK_INVERTED,

		moves.FRONT,
		moves.FRONT2,
		moves.FRONT_INVERTED,

		moves.DOWN,
		moves.DOWN2,
		moves.DOWN_INVERTED,

		moves.UP,
		moves.UP2,
		moves.UP_INVERTED,

		moves.M,
		moves.M2,
		moves.M_INVERTED,

		moves.X,
		moves.X2,
		moves.X_INVERTED,

		moves.Y,
		moves.Y2,
		moves.Y_INVERTED,

		moves.E,
		moves.E2,
		moves.E_INVERTED,
	}

	for _, move := range movesToTest {
		t.Run(fmt.Sprintf("test move %v", moves.ToString([]moves.Move{move})), func(t *testing.T) {
			cube, mp := getTestCube()
			for i := 0; i < 4; i++ {
				cube.ApplyMove(move)
			}

			for f := faces.FRONT; f <= faces.BOTTOM; f++ {
				for i := 0; i < 3; i++ {
					for j := 0; j < 3; j++ {
						assert.Equal(t, mp[f][i][j], cube.Cube[f][i][j], fmt.Sprintf("i: %d, j: %d, face: %d", i, j, f))
					}
				}
			}
		})
	}
}

func TestCube1_TestCombinationMoves(t *testing.T) {
	combinationMovesToTest := [][]moves.Move{
		{
			moves.RIGHT,
			moves.UP,
			moves.RIGHT_INVERTED,
			moves.UP_INVERTED,
		},
		{
			moves.LEFT_INVERTED,
			moves.UP_INVERTED,
			moves.LEFT,
			moves.UP,
		},
	}

	for id, movesToTest := range combinationMovesToTest {
		cube, mp := getTestCube()
		for i := 0; i < 6; i++ {
			cube.ApplyMoveSequence(movesToTest)
		}

		for f := faces.FRONT; f <= faces.BOTTOM; f++ {
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					assert.Equal(t, mp[f][i][j], cube.Cube[f][i][j], fmt.Sprintf("moves to test id: %v\nface: %v\ni: %v\nj: %v", id, f, i, j))
				}
			}
		}
	}
}

func TestCube1_Encode(t *testing.T) {
	cube := NewCube1()
	str := cube.Encode()

	mp := map[rune]int{}

	for _, chars := range str {
		mp[chars]++
	}

	for _, v := range mp {
		assert.Equal(t, 9, v)
	}
}

func TestCube1_Decode_SolvedCube(t *testing.T) {
	cube := NewCube1()
	solved := cube.Encode()

	cube.Decode(solved)
	for f := faces.FRONT; f <= faces.BOTTOM; f++ {
		intf := int(f)
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				assert.Equal(t, string(solved[9*intf+3*i+j]), cube.Cube[f][i][j].ToString())
			}
		}
	}
}

func TestCube1_Decode_UnsolvedCube(t *testing.T) {
	cube := NewCube1()
	solved := cube.Encode()
	fmt.Println(solved)
	movesToTest := []moves.Move{
		moves.RIGHT,
		moves.UP,
		moves.RIGHT_INVERTED,
		moves.UP_INVERTED,
	}
	cube.ApplyMoveSequence(movesToTest)
	unsolved := cube.Encode()
	fmt.Println(unsolved)

	cube1 := NewCube1()
	cube1.Decode(unsolved)
	for i := 0; i < 5; i++ {
		cube1.ApplyMoveSequence(movesToTest)
	}
	for f := faces.FRONT; f <= faces.BOTTOM; f++ {
		intf := int(f)
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				assert.Equal(t, string(solved[9*intf+3*i+j]), cube1.Cube[f][i][j].ToString(), fmt.Sprintf("i: %v, j: %v, face: %v", i, j, f))
			}
		}
	}
}
