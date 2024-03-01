package cube1

import (
	"solver/cube/colors"
	"solver/cube/faces"
)

func getColorsArrary(c colors.Color) [][]colors.Color {
	bytes := []colors.Color{}
	for i := 0; i < 9; i++ {
		bytes = append(bytes, c)
	}
	return [][]colors.Color{
		bytes[:3],
		bytes[3:6],
		bytes[6:9],
	}
}

func GetEmptyFace() [][]colors.Color {
	return [][]colors.Color{
		{colors.BLACK, colors.BLACK, colors.BLACK},
		{colors.BLACK, colors.BLACK, colors.BLACK},
		{colors.BLACK, colors.BLACK, colors.BLACK},
	}
}

func (c *Cube1) getStripe(f faces.Face, rc bool, dir bool, sr, sc int) []colors.Color {
	by := []colors.Color{}
	if rc {
		for i := 0; i < 3; i++ {
			if dir {
				by = append(by, c.Cube[f][sr][i])
			} else {
				by = append(by, c.Cube[f][sr][2-i])
			}
		}
	} else {
		for i := 0; i < 3; i++ {
			if dir {
				by = append(by, c.Cube[f][i][sc])
			} else {
				by = append(by, c.Cube[f][2-i][sc])
			}
		}
	}
	return by
}

func (c *Cube1) moveRightInverted() {
	c.rotateFaceCounterClockwise(faces.RIGHT)

	strip := []colors.Color{}

	strip = append(
		strip,
		c.getStripe(faces.FRONT, false, true, 0, 2)...,
	)
	strip = append(
		strip,
		c.getStripe(faces.BOTTOM, false, true, 0, 2)...,
	)
	strip = append(
		strip,
		c.getStripe(faces.BACK, false, false, 2, 0)...,
	)
	strip = append(
		strip,
		c.getStripe(faces.TOP, false, true, 0, 2)...,
	)

	c.Cube[faces.BOTTOM][0][2] = strip[0]
	c.Cube[faces.BOTTOM][1][2] = strip[1]
	c.Cube[faces.BOTTOM][2][2] = strip[2]

	c.Cube[faces.BACK][2][0] = strip[3]
	c.Cube[faces.BACK][1][0] = strip[4]
	c.Cube[faces.BACK][0][0] = strip[5]

	c.Cube[faces.TOP][0][2] = strip[6]
	c.Cube[faces.TOP][1][2] = strip[7]
	c.Cube[faces.TOP][2][2] = strip[8]

	c.Cube[faces.FRONT][0][2] = strip[9]
	c.Cube[faces.FRONT][1][2] = strip[10]
	c.Cube[faces.FRONT][2][2] = strip[11]
}

func (c *Cube1) moveLeftInverted() {
	c.rotateFaceCounterClockwise(faces.LEFT)

	strip := []colors.Color{}

	strip = append(
		strip,
		c.getStripe(faces.FRONT, false, false, 2, 0)...,
	)
	strip = append(
		strip,
		c.getStripe(faces.TOP, false, false, 2, 0)...,
	)
	strip = append(
		strip,
		c.getStripe(faces.BACK, false, true, 0, 2)...,
	)
	strip = append(
		strip,
		c.getStripe(faces.BOTTOM, false, false, 2, 0)...,
	)

	c.Cube[faces.TOP][2][0] = strip[0]
	c.Cube[faces.TOP][1][0] = strip[1]
	c.Cube[faces.TOP][0][0] = strip[2]

	c.Cube[faces.BACK][0][2] = strip[3]
	c.Cube[faces.BACK][1][2] = strip[4]
	c.Cube[faces.BACK][2][2] = strip[5]

	c.Cube[faces.BOTTOM][2][0] = strip[6]
	c.Cube[faces.BOTTOM][1][0] = strip[7]
	c.Cube[faces.BOTTOM][0][0] = strip[8]

	c.Cube[faces.FRONT][2][0] = strip[9]
	c.Cube[faces.FRONT][1][0] = strip[10]
	c.Cube[faces.FRONT][0][0] = strip[11]
}

func (c *Cube1) moveE() {
	strip := []colors.Color{}

	strip = append(
		strip,
		c.getStripe(faces.FRONT, true, true, 1, 0)...,
	)
	strip = append(
		strip,
		c.getStripe(faces.RIGHT, true, true, 1, 0)...,
	)
	strip = append(
		strip,
		c.getStripe(faces.BACK, true, true, 1, 0)...,
	)
	strip = append(
		strip,
		c.getStripe(faces.LEFT, true, true, 1, 0)...,
	)

	c.Cube[faces.RIGHT][1][0] = strip[0]
	c.Cube[faces.RIGHT][1][1] = strip[1]
	c.Cube[faces.RIGHT][1][2] = strip[2]

	c.Cube[faces.BACK][1][0] = strip[3]
	c.Cube[faces.BACK][1][1] = strip[4]
	c.Cube[faces.BACK][1][2] = strip[5]

	c.Cube[faces.LEFT][1][0] = strip[6]
	c.Cube[faces.LEFT][1][1] = strip[7]
	c.Cube[faces.LEFT][1][2] = strip[8]

	c.Cube[faces.FRONT][1][0] = strip[9]
	c.Cube[faces.FRONT][1][1] = strip[10]
	c.Cube[faces.FRONT][1][2] = strip[11]
}

func (c *Cube1) moveM() {
	strip := []colors.Color{}

	strip = append(
		strip,
		c.getStripe(faces.FRONT, false, true, 0, 1)...,
	)
	strip = append(
		strip,
		c.getStripe(faces.BOTTOM, false, true, 0, 1)...,
	)
	strip = append(
		strip,
		c.getStripe(faces.BACK, false, false, 2, 1)...,
	)
	strip = append(
		strip,
		c.getStripe(faces.TOP, false, true, 0, 1)...,
	)

	c.Cube[faces.BOTTOM][0][1] = strip[0]
	c.Cube[faces.BOTTOM][1][1] = strip[1]
	c.Cube[faces.BOTTOM][2][1] = strip[2]

	c.Cube[faces.BACK][2][1] = strip[3]
	c.Cube[faces.BACK][1][1] = strip[4]
	c.Cube[faces.BACK][0][1] = strip[5]

	c.Cube[faces.TOP][0][1] = strip[6]
	c.Cube[faces.TOP][1][1] = strip[7]
	c.Cube[faces.TOP][2][1] = strip[8]

	c.Cube[faces.FRONT][0][1] = strip[9]
	c.Cube[faces.FRONT][1][1] = strip[10]
	c.Cube[faces.FRONT][2][1] = strip[11]
}

func (c *Cube1) moveDown() {
	c.rotateFaceClockwise(faces.BOTTOM)

	strip := []colors.Color{}

	strip = append(
		strip,
		c.getStripe(faces.FRONT, true, true, 2, 0)...,
	)
	strip = append(
		strip,
		c.getStripe(faces.RIGHT, true, true, 2, 0)...,
	)
	strip = append(
		strip,
		c.getStripe(faces.BACK, true, true, 2, 0)...,
	)
	strip = append(
		strip,
		c.getStripe(faces.LEFT, true, true, 2, 0)...,
	)

	c.Cube[faces.RIGHT][2][0] = strip[0]
	c.Cube[faces.RIGHT][2][1] = strip[1]
	c.Cube[faces.RIGHT][2][2] = strip[2]

	c.Cube[faces.BACK][2][0] = strip[3]
	c.Cube[faces.BACK][2][1] = strip[4]
	c.Cube[faces.BACK][2][2] = strip[5]

	c.Cube[faces.LEFT][2][0] = strip[6]
	c.Cube[faces.LEFT][2][1] = strip[7]
	c.Cube[faces.LEFT][2][2] = strip[8]

	c.Cube[faces.FRONT][2][0] = strip[9]
	c.Cube[faces.FRONT][2][1] = strip[10]
	c.Cube[faces.FRONT][2][2] = strip[11]
}

func (c *Cube1) moveUp() {
	c.rotateFaceClockwise(faces.TOP)

	strip := []colors.Color{}

	strip = append(
		strip,
		c.getStripe(faces.FRONT, true, false, 0, 2)...,
	)
	strip = append(
		strip,
		c.getStripe(faces.LEFT, true, false, 0, 2)...,
	)
	strip = append(
		strip,
		c.getStripe(faces.BACK, true, false, 0, 2)...,
	)
	strip = append(
		strip,
		c.getStripe(faces.RIGHT, true, false, 0, 2)...,
	)

	c.Cube[faces.LEFT][0][2] = strip[0]
	c.Cube[faces.LEFT][0][1] = strip[1]
	c.Cube[faces.LEFT][0][0] = strip[2]

	c.Cube[faces.BACK][0][2] = strip[3]
	c.Cube[faces.BACK][0][1] = strip[4]
	c.Cube[faces.BACK][0][0] = strip[5]

	c.Cube[faces.RIGHT][0][2] = strip[6]
	c.Cube[faces.RIGHT][0][1] = strip[7]
	c.Cube[faces.RIGHT][0][0] = strip[8]

	c.Cube[faces.FRONT][0][2] = strip[9]
	c.Cube[faces.FRONT][0][1] = strip[10]
	c.Cube[faces.FRONT][0][0] = strip[11]
}

func (c *Cube1) moveBack() {
	c.rotateFaceClockwise(faces.BACK)

	strip := []colors.Color{}

	strip = append(
		strip,
		c.getStripe(faces.TOP, true, false, 0, 2)...,
	)
	strip = append(
		strip,
		c.getStripe(faces.LEFT, false, true, 0, 0)...,
	)
	strip = append(
		strip,
		c.getStripe(faces.BOTTOM, true, true, 2, 0)...,
	)
	strip = append(
		strip,
		c.getStripe(faces.RIGHT, false, false, 2, 2)...,
	)

	c.Cube[faces.LEFT][0][0] = strip[0]
	c.Cube[faces.LEFT][1][0] = strip[1]
	c.Cube[faces.LEFT][2][0] = strip[2]

	c.Cube[faces.BOTTOM][2][0] = strip[3]
	c.Cube[faces.BOTTOM][2][1] = strip[4]
	c.Cube[faces.BOTTOM][2][2] = strip[5]

	c.Cube[faces.RIGHT][2][2] = strip[6]
	c.Cube[faces.RIGHT][1][2] = strip[7]
	c.Cube[faces.RIGHT][0][2] = strip[8]

	c.Cube[faces.TOP][0][2] = strip[9]
	c.Cube[faces.TOP][0][1] = strip[10]
	c.Cube[faces.TOP][0][0] = strip[11]
}

func (c *Cube1) moveFront() {
	c.rotateFaceClockwise(faces.FRONT)

	strip := []colors.Color{}

	strip = append(
		strip,
		c.getStripe(faces.TOP, true, true, 2, 0)...,
	)
	strip = append(
		strip,
		c.getStripe(faces.RIGHT, false, true, 0, 0)...,
	)
	strip = append(
		strip,
		c.getStripe(faces.BOTTOM, true, false, 0, 2)...,
	)
	strip = append(
		strip,
		c.getStripe(faces.LEFT, false, false, 2, 2)...,
	)

	c.Cube[faces.RIGHT][0][0] = strip[0]
	c.Cube[faces.RIGHT][1][0] = strip[1]
	c.Cube[faces.RIGHT][2][0] = strip[2]

	c.Cube[faces.BOTTOM][0][2] = strip[3]
	c.Cube[faces.BOTTOM][0][1] = strip[4]
	c.Cube[faces.BOTTOM][0][0] = strip[5]

	c.Cube[faces.LEFT][2][2] = strip[6]
	c.Cube[faces.LEFT][1][2] = strip[7]
	c.Cube[faces.LEFT][0][2] = strip[8]

	c.Cube[faces.TOP][2][0] = strip[9]
	c.Cube[faces.TOP][2][1] = strip[10]
	c.Cube[faces.TOP][2][2] = strip[11]
}

func (c *Cube1) moveLeft() {
	c.rotateFaceClockwise(faces.LEFT)

	strip := []colors.Color{}

	strip = append(strip, c.getStripe(faces.FRONT, false, true, 0, 0)...)
	strip = append(strip, c.getStripe(faces.BOTTOM, false, true, 0, 0)...)
	strip = append(strip, c.getStripe(faces.BACK, false, false, 2, 2)...)
	strip = append(strip, c.getStripe(faces.TOP, false, true, 0, 0)...)

	c.Cube[faces.BOTTOM][0][0] = strip[0]
	c.Cube[faces.BOTTOM][1][0] = strip[1]
	c.Cube[faces.BOTTOM][2][0] = strip[2]

	c.Cube[faces.BACK][2][2] = strip[3]
	c.Cube[faces.BACK][1][2] = strip[4]
	c.Cube[faces.BACK][0][2] = strip[5]

	c.Cube[faces.TOP][0][0] = strip[6]
	c.Cube[faces.TOP][1][0] = strip[7]
	c.Cube[faces.TOP][2][0] = strip[8]

	c.Cube[faces.FRONT][0][0] = strip[9]
	c.Cube[faces.FRONT][1][0] = strip[10]
	c.Cube[faces.FRONT][2][0] = strip[11]
}

func (c *Cube1) moveRight() {
	c.rotateFaceClockwise(faces.RIGHT)

	strip := []colors.Color{}

	strip = append(strip, c.getStripe(faces.FRONT, false, false, 2, 2)...)
	strip = append(strip, c.getStripe(faces.TOP, false, false, 2, 2)...)
	strip = append(strip, c.getStripe(faces.BACK, false, true, 0, 0)...)
	strip = append(strip, c.getStripe(faces.BOTTOM, false, false, 2, 2)...)

	c.Cube[faces.TOP][2][2] = strip[0]
	c.Cube[faces.TOP][1][2] = strip[1]
	c.Cube[faces.TOP][0][2] = strip[2]

	c.Cube[faces.BACK][0][0] = strip[3]
	c.Cube[faces.BACK][1][0] = strip[4]
	c.Cube[faces.BACK][2][0] = strip[5]

	c.Cube[faces.BOTTOM][2][2] = strip[6]
	c.Cube[faces.BOTTOM][1][2] = strip[7]
	c.Cube[faces.BOTTOM][0][2] = strip[8]

	c.Cube[faces.FRONT][2][2] = strip[9]
	c.Cube[faces.FRONT][1][2] = strip[10]
	c.Cube[faces.FRONT][0][2] = strip[11]
}

func (c *Cube1) rotateFaceCounterClockwise(f faces.Face) {
	face := GetEmptyFace()
	face[0][0] = c.Cube[f][0][2]
	face[2][0] = c.Cube[f][0][0]
	face[2][2] = c.Cube[f][2][0]
	face[0][2] = c.Cube[f][2][2]

	face[0][1] = c.Cube[f][1][2]
	face[1][0] = c.Cube[f][0][1]
	face[2][1] = c.Cube[f][1][0]
	face[1][2] = c.Cube[f][2][1]

	face[1][1] = c.Cube[f][1][1]

	c.Cube[f] = face
}

func (c *Cube1) rotateFaceClockwise(f faces.Face) {
	face := GetEmptyFace()
	face[0][0] = c.Cube[f][2][0]
	face[0][2] = c.Cube[f][0][0]
	face[2][2] = c.Cube[f][0][2]
	face[2][0] = c.Cube[f][2][2]

	face[0][1] = c.Cube[f][1][0]
	face[1][2] = c.Cube[f][0][1]
	face[2][1] = c.Cube[f][1][2]
	face[1][0] = c.Cube[f][2][1]

	face[1][1] = c.Cube[f][1][1]

	c.Cube[f] = face
}
