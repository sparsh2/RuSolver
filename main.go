package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"solver/cube"
	"solver/cube/colors"
	"solver/cube/cube1"
	"solver/cube/faces"
	"solver/cube/moves"
	"solver/cube/types"

	"github.com/gammazero/deque"
)

type Node struct {
	c     types.ICube
	depth int
}

func main() {
	rouxCube := cube1.NewCube1()
	// rouxCube.TrackMoves = true
	rouxCube.SetTrackMoves(true)
	// setRouxFB(rouxCube)
	rouxCube.Decode(cube.RouxMask)

	q := deque.New[Node](65536, 32)
	mp := map[string][]moves.Move{}
	mp[rouxCube.Encode()] = []moves.Move{}

	movesList := []moves.Move{
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
	}

	maxDep := 5

	for _, move := range movesList {
		copyCube := rouxCube.GetCopy()
		copyCube.ApplyMove(move)

		if _, ok := mp[copyCube.Encode()]; !ok {
			q.PushBack(Node{
				c:     copyCube,
				depth: 1,
			})
		}
	}

	for q.Len() != 0 {
		node := q.PopFront()
		hash := node.c.Encode()
		if _, ok := mp[hash]; !ok {
			mp[hash] = append(mp[hash], node.c.GetMoves()...)
			if node.depth == maxDep {
				continue
			}
			for _, move := range movesList {
				copyCube := node.c.GetCopy()
				copyCube.ApplyMove(move)

				if _, ok := mp[copyCube.Encode()]; !ok {
					q.PushBack(Node{
						c:     copyCube,
						depth: node.depth + 1,
					})
				}
			}
		}
	}

	// fmt.Println(mp)

	fmt.Println(len(mp))
	save(mp)
}

func save(mp map[string][]moves.Move) {
	newMp := map[string]string{}

	for k, v := range mp {
		strMoves := moves.ToString(v)
		newMp[k] = strMoves
	}

	bytes, err := json.MarshalIndent(newMp, "", "  ")
	if err != nil {
		fmt.Println("couldn't marshal")
		return
	}

	err = ioutil.WriteFile("gen.fb", bytes, 0755)
	if err != nil {
		fmt.Println("couldn't save")
		return
	}
}

func setRouxFB(c *cube1.Cube1) {
	c.Cube[faces.TOP] = cube1.GetEmptyFace()
	c.Cube[faces.RIGHT] = cube1.GetEmptyFace()
	c.Cube[faces.LEFT] = cube1.GetEmptyFace()
	c.Cube[faces.BOTTOM] = cube1.GetEmptyFace()
	c.Cube[faces.FRONT] = cube1.GetEmptyFace()
	c.Cube[faces.BACK] = cube1.GetEmptyFace()

	for i := 1; i < 3; i++ {
		for j := 0; j < 3; j++ {
			c.Cube[faces.LEFT][i][j] = colors.GREEN
		}
	}
	for i := 0; i < 3; i++ {
		c.Cube[faces.BOTTOM][i][0] = colors.WHITE
	}
	c.Cube[faces.FRONT][1][0] = colors.ORANGE
	c.Cube[faces.FRONT][2][0] = colors.ORANGE

	c.Cube[faces.BACK][1][2] = colors.RED
	c.Cube[faces.BACK][2][2] = colors.RED

}
