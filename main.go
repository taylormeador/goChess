package main

import "fmt"

// init
func initAll() {
	initLeapersAttacks()
	initSlidersAttacks(bishop)
	initSlidersAttacks(rook)
}

// main
func main() {
	initAll()
	parseFEN(trickyPosition)
	move := parseMove("b7b8b")
	if move != 0 {
		makeMove(move)
		printMove(move)
	} else {
		fmt.Printf("Illegal move")
	}
}
