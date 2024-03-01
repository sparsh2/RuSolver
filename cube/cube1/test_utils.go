package cube1

import (
	"solver/cube/colors"
	"solver/cube/faces"
	"solver/cube/moves"
)

func getTestCube() (*Cube1, map[faces.Face][][]colors.Color) {
	cube := &Cube1{
		Cube:       make(map[faces.Face][][]colors.Color),
		TrackMoves: false,
		Moves:      []moves.Move{},
	}

	cube.Cube[faces.FRONT] = getColorsArrary(colors.WHITE)
	cube.Cube[faces.BACK] = getColorsArrary(colors.YELLOW)
	cube.Cube[faces.LEFT] = getColorsArrary(colors.RED)
	cube.Cube[faces.RIGHT] = getColorsArrary(colors.ORANGE)
	cube.Cube[faces.TOP] = getColorsArrary(colors.BLUE)
	cube.Cube[faces.BOTTOM] = getColorsArrary(colors.GREEN)

	mp := map[faces.Face][][]colors.Color{}
	mp[faces.FRONT] = GetEmptyFace()
	mp[faces.BACK] = GetEmptyFace()
	mp[faces.LEFT] = GetEmptyFace()
	mp[faces.RIGHT] = GetEmptyFace()
	mp[faces.TOP] = GetEmptyFace()
	mp[faces.BOTTOM] = GetEmptyFace()

	for f := faces.FRONT; f <= faces.BOTTOM; f++ {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				sum := i*3 + j
				cube.Cube[f][i][j] = colors.Color(sum%6 + 1)
				mp[f][i][j] = colors.Color(sum%6 + 1)
			}
		}
	}

	return cube, mp
}
