package gen_test

import (
	"encoding/json"
	"io/ioutil"
	"solver/cube/cube1"
	"solver/cube/faces"
	"solver/cube/moves"
	"solver/generate/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

const readDir = "../../target/gen.fb"

func TestGen(t *testing.T) {
	// mapCubeToMoves := map[*cube1.Cube1][]moves.Move{}
	mp, err := read()
	if err != nil {
		t.Errorf("couldn't parse generated file: %v", err)
	}

	rouxCube := cube1.NewCube1().(*cube1.Cube1)
	utils.SetRouxFB(rouxCube)

	for k, v := range mp {
		// TODO: figure out testing
		c := cube1.NewCube1().(*cube1.Cube1)
		c.Decode(k)

		if v == "---O--O---G----R------GG-GG--------W-R-------W--W----G" {
			c.PrettyPrint()
		}

		c.ApplyMoveSequence(moves.Parse(v))

		for f := faces.FRONT; f <= faces.BOTTOM; f++ {
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					assert.Equal(t, rouxCube.Cube[f][i][j], c.Cube[f][i][j], "k:%v\nv:%v\n", k, v)
					if rouxCube.Cube[f][i][j] != c.Cube[f][i][j] {
						c.ApplyMoveSequence(moves.InvertMoves(moves.Parse(v)))
						c.PrettyPrint()
						t.FailNow()
					}
				}
			}
		}
	}

}

func read() (map[string]string, error) {
	bytes, err := ioutil.ReadFile(readDir)
	mp := map[string]string{}
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &mp)
	if err != nil {
		return nil, err
	}

	return mp, nil
}
