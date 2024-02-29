package main

import (
	"fmt"
	"solver/cube/cube1"
	"solver/cube/moves"
	"solver/generate/utils"
)

func main() {
	c := cube1.NewCube1()
	utils.SetRouxFB(c)

	moveStr := "U' L2 B2 F2 U' F2 D U F2 R2 U2 R' U' L' U2 B U B' D2 x'"
	// "U' L2 B2 F2 U' F2 D U F2 R2 U2 R' U' L' U2 B U B' D2"
	moveList := moves.Parse(moveStr)
	c.ApplyMoveSequence(moveList)
	c.PrettyPrint()
	fmt.Println(c.Encode())
}
