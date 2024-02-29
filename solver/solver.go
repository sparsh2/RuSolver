package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"solver/cube/cube1"
	"solver/cube/faces"
	"solver/cube/moves"
	"solver/generate/utils"
	"solver/solver/lib"
	"strings"

	"github.com/gammazero/deque"
)

var genFb map[string][]moves.Move
var genDir = "/Users/ssarode/pers/cubesolver/target/gen.fb"

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter scramble: ")
	sc, _ := reader.ReadString('\n')
	sc = strings.Trim(sc, "\n")

	scrambleMoves := moves.Parse(sc)

	c := make(chan string, 1)

	for _, fb := range lib.FirstBlocksToSolve {
		go worker(c, fb, scrambleMoves)
	}

	for i := 0; i < len(lib.FirstBlocksToSolve); i++ {
		res := <-c
		fmt.Println(res)
	}
}

func worker(c chan<- string, fb *lib.FirstBlock, scrambleMoves []moves.Move) {
	ruCube := cube1.NewCube1()
	utils.SetRouxFB(ruCube)

	ruCube.ApplyMoveSequence(moves.InvertMoves(fb.Setup))
	ruCube.ApplyMoveSequence(scrambleMoves)
	ruCube.ApplyMoveSequence(fb.Setup)

	solvedpremoves := []moves.Move{}
	unorientedSol := solve(ruCube)

	// Following section only required for slice moves
	// -----------------------------------------------------------------
	// premoves := []moves.Move{}
	// for i := 0; i < 4; i++ {
	// 	for j := 0; j < 4; j++ {
	// 		mvs := solve(ruCube)
	// 		if len(unorientedSol) > len(mvs) {
	// 			unorientedSol = mvs
	// 			solvedpremoves = []moves.Move{}
	// 			solvedpremoves = append(solvedpremoves, premoves...)
	// 		}
	// 		premoves = append(premoves, moves.Y)
	// 		ruCube.ApplyMove(moves.Y)
	// 	}
	// 	premoves = append(premoves, moves.X)
	// 	ruCube.ApplyMove(moves.X)
	// }
	// -----------------------------------------------------------------

	finSol := []moves.Move{}

	finSol = append(finSol, fb.Setup...)
	finSol = append(finSol, solvedpremoves...)
	finSol = minimize(finSol)
	finSol = append(finSol, unorientedSol...)

	c <- fmt.Sprintf("%s:\t%v\t%v", fb.Name, len(unorientedSol), moves.ToString(finSol))
}

func solve(cq *cube1.Cube1) []moves.Move {
	// utils.Node
	q := deque.New[*utils.Node](2048, 32)
	cq.TrackMoves = true
	cq.Moves = []moves.Move{}

	c := cq.Duplicate()
	mp := map[string]bool{}

	q.PushBack(&utils.Node{
		C:     c,
		Depth: 0,
	})
	// mp[c.Encode()] = true

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

	for q.Len() != 0 {
		n := q.PopFront()
		if _, ok := mp[n.C.Encode()]; ok {
			continue
		} else {
			mp[n.C.Encode()] = true
		}
		if genSolution, ok := genFb[n.C.Encode()]; ok {
			finSolution := []moves.Move{}
			finSolution = append(finSolution, n.C.Moves...)
			finSolution = append(finSolution, genSolution...)
			return finSolution
		}
		for _, move := range movesList {
			newNode := &utils.Node{
				C:     n.C.Duplicate(),
				Depth: n.Depth + 1,
			}
			// log.Default().Printf("depth: %v\n", n.Depth+1)
			newNode.C.ApplyMove(move)
			q.PushBack(newNode)
		}
	}

	return []moves.Move{}
}

func minimize(moveList []moves.Move) []moves.Move {
	c := cube1.NewCube1()
	for i := 0; i < len(moveList)-1; i++ {
		dc1 := c.Duplicate()
		dc2 := c.Duplicate()

		dc2.ApplyMove(moveList[i])
		dc2.ApplyMove(moveList[i+1])

		for mv := moves.LEFT; mv <= moves.S2; mv++ {
			newDc := dc1.Duplicate()
			newDc.ApplyMove(mv)
			if isEqual(newDc, dc2) {
				newSlice := []moves.Move{}
				for j := 0; j < i; j++ {
					newSlice = append(newSlice, moveList[j])
				}
				newSlice = append(newSlice, mv)
				for j := i + 2; j < len(moveList); j++ {
					newSlice = append(newSlice, moveList[j])
				}
				return minimize(newSlice)
			}
		}
	}

	return moveList
}

func isEqual(c1 *cube1.Cube1, c2 *cube1.Cube1) bool {
	for f := faces.FRONT; f < faces.BOTTOM; f++ {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if c1.Cube[f][i][j] != c2.Cube[f][i][j] {
					return false
				}
			}
		}
	}
	return true
}

func init() {
	log.Default().Print("loading roux data")
	mp, err := read()
	genFb = map[string][]moves.Move{}
	if err != nil {
		log.Default().Fatal(err)
	}

	for k, v := range mp {
		if v == "" {
			genFb[k] = []moves.Move{}
			continue
		}
		genFb[k] = moves.Parse(v)
	}
	log.Default().Print("loading completed!")
	log.Default().Printf("size: %v", len(genFb))
	// fmt.Println(moves.ToString(minimize(moves.Parse("y y"))))
}

func read() (map[string]string, error) {
	bytes, err := ioutil.ReadFile(genDir)
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
