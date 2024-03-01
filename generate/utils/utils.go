package utils

import (
	"solver/cube/colors"
	"solver/cube/cube1"
	"solver/cube/faces"
	"solver/cube/types"
)

type Node struct {
	C     types.ICube
	Depth int
}

func SetRouxFB(c *cube1.Cube1) {
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
