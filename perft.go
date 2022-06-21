package main

import "fmt"

var nodes uint64

func perftDriver(depth int) {
	// recursion excape considition
	if depth == 0 {
		nodes++
		return
	}
	moveList = []uint64{}
	generateAllMoves()
	fmt.Println(side)
	for _, move := range moveList {
		copyBoard()
		if makeMove(move) == 0 {
			continue
		}
		// recursive call
		perftDriver(depth - 1)
		restoreBoard()
	}
}
