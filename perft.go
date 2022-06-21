package main

var nodes uint64

func perftDriver(depth int) {
	if depth == 0 {
		nodes++
		return
	}
	generateAllMoves()
	for _, move := range moveList {
		copyBoard()
		if makeMove(move) == 0 {
			continue
		}
		perftDriver(depth - 1)
		restoreBoard()
	}
}
