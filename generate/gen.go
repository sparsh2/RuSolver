package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"solver/cube/cube1"
	"solver/cube/moves"
	"solver/generate/utils"
	"sort"
	"strings"

	"github.com/gammazero/deque"
)

const trainingDataPath = "../target/gen_train.csv"
const outputDir = "../target/gen.fb"
const dataDir = "../target/gendata.fb"

func genTrainingSet() {
	rouxCube := cube1.NewCube1()
	rouxCube.TrackMoves = true
	utils.SetRouxFB(rouxCube)

	q := deque.New[utils.Node](65536, 32)
	mp := map[string][]moves.Move{}
	// mp[rouxCube.Encode()] = []moves.Move{}

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

		// moves.M,
		// moves.M2,
		// moves.M_INVERTED,

		// moves.E,
		// moves.E2,
		// moves.E_INVERTED,
	}

	maxDep := 4

	q.PushBack(utils.Node{
		C:     rouxCube,
		Depth: 0,
	})

	// for _, move := range movesList {
	// 	copyCube := rouxCube.Duplicate()
	// 	copyCube.ApplyMove(move)

	// 	if !canSkip(copyCube, mp) {
	// 		q.PushBack(utils.Node{
	// 			C:     copyCube,
	// 			Depth: 1,
	// 		})
	// 		// mp[copyCube.Encode()] = append(mp[copyCube.Encode()], copyCube.Moves...)
	// 	}
	// }

	listOfMoves := [][]moves.Move{}
	for q.Len() != 0 {
		node := q.PopFront()
		hash := node.C.Encode()
		if !canSkip(node.C, mp) {
			mp[hash] = append(mp[hash], node.C.Moves...)
			cpyMv := make([]moves.Move, len(node.C.Moves))
			copy(cpyMv, node.C.Moves)
			listOfMoves = append(listOfMoves, cpyMv)
			if node.Depth == maxDep {
				continue
			}
			for _, move := range movesList {
				copyCube := node.C.Duplicate()
				copyCube.ApplyMove(move)

				q.PushBack(utils.Node{
					C:     copyCube,
					Depth: node.Depth + 1,
				})
				// if _, ok := mp[copyCube.Encode()]; !ok {
				// }
			}
		}
	}

	listOfMoves2 := filter(listOfMoves, 1)
	listOfMoves3 := filter(listOfMoves, 2)
	listOfMoves4 := filter(listOfMoves, 3)
	listOfMoves5 := filter(listOfMoves, 4)

	trainingMoves := []string{}
	trainingMoves = append(trainingMoves, getSortedStrings(listOfMoves2)...)
	trainingMoves = append(trainingMoves, getSortedStrings(listOfMoves3)...)
	trainingMoves = append(trainingMoves, getSortedStrings(listOfMoves4)...)
	trainingMoves = append(trainingMoves, getSortedStrings(listOfMoves5)...)

	saveTraining(trainingMoves)
}

func saveTraining(data []string) {
	var b strings.Builder
	for _, mv := range data {
		fmt.Fprintf(&b, "%s,,\n", mv)
	}
	// fmt.Println(b.String()[0])
	ioutil.WriteFile(trainingDataPath, []byte(b.String()), 0700)
}

func getSortedStrings(moveSeqs [][]moves.Move) []string {
	res := []string{}
	for _, mvs := range moveSeqs {
		res = append(res, moves.ToString(mvs))
	}
	sort.Strings(res)
	return res
}

func filter(moveSeqs [][]moves.Move, size int) [][]moves.Move {
	filtered := [][]moves.Move{}
	for _, mvs := range moveSeqs {
		if len(mvs) == size {
			filtered = append(filtered, mvs)
		}
	}
	return filtered
}

func main() {

	// gen training
	// genTrainingSet()
	// return

	rouxCube := cube1.NewCube1()
	rouxCube.TrackMoves = true
	utils.SetRouxFB(rouxCube)

	q := deque.New[utils.Node](65536, 32)
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

		// moves.M,
		// moves.M2,
		// moves.M_INVERTED,

		// moves.E,
		// moves.E2,
		// moves.E_INVERTED,
	}

	maxDep := 6

	for _, move := range movesList {
		copyCube := rouxCube.Duplicate()
		copyCube.ApplyMove(move)

		if _, ok := mp[copyCube.Encode()]; !ok {
			q.PushBack(utils.Node{
				C:     copyCube,
				Depth: 1,
			})
		}
	}

	for q.Len() != 0 {
		node := q.PopFront()
		hash := node.C.Encode()
		if _, ok := mp[hash]; !ok {
			if canSkip(node.C, mp) {
				continue
			}
			mp[hash] = append(mp[hash], node.C.Moves...)
			if node.Depth == maxDep {
				continue
			}
			for _, move := range movesList {
				copyCube := node.C.Duplicate()
				copyCube.ApplyMove(move)

				if _, ok := mp[copyCube.Encode()]; !ok {
					q.PushBack(utils.Node{
						C:     copyCube,
						Depth: node.Depth + 1,
					})
				}
			}
		}
	}

	fmt.Printf("size : %d\n", len(mp))

	save(mp)
}

func canSkip(cube *cube1.Cube1, mp map[string][]moves.Move) bool {
	c := cube.Duplicate()
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if _, ok := mp[c.Encode()]; ok {
				return true
			}
			c.ApplyMove(moves.Y)
		}
		c.ApplyMove(moves.X)
	}
	return false
}

func save(mp map[string][]moves.Move) {
	isJson := true
	newMp := map[string]string{}

	for k, v := range mp {
		strMoves := moves.ToString(moves.InvertMoves(v))
		newMp[k] = strMoves
	}

	if !isJson {
		isFirst := false
		for k, v := range mp {
			lines := ""
			for _, r := range k {
				lines += encodeColor(string(r))
			}
			lines += ","
			for _, mv := range v {
				lines += fmt.Sprintf("%v ", moves.ToStringMove(mv))
			}
			lines += "\n"
			// fmt.Print(lines)
			if isFirst {
				x := ""
				fmt.Scanf("%s", &x)
				isFirst = false
			}
			bytes, err := ioutil.ReadFile(dataDir)
			if err != nil {
				fmt.Println(err)
				return
			}
			bytes = append(bytes, []byte(lines)...)

			err = ioutil.WriteFile(dataDir, bytes, 0755)
			if err != nil {
				fmt.Println("couldn't save")
				fmt.Println(err)
				return
			}
		}
		return
	}

	bytes, err := json.MarshalIndent(newMp, "", "  ")
	if err != nil {
		fmt.Println("couldn't marshal")
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(outputDir, bytes, 0755)
	if err != nil {
		fmt.Println("couldn't save")
		fmt.Println(err)
		return
	}
}

func encodeColor(color string) string {
	switch color {
	case "G":
		return "0"
	case "W":
		return "1"
	case "O":
		return "2"
	case "R":
		return "3"
	default:
		return "4"
	}
}
