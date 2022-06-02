package main

import (
	"bufio"
	"os"
)

// init
func initAll() {
	initLeapersAttacks()
	initSlidersAttacks(bishop)
	initSlidersAttacks(rook)
}

// main
func main() {
	initAll()

	parseFEN("r3k2r/pPppqpb1/bn2pnp1/3PN3/Pp2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq b5 0 1")
	//fmt.Println(getMoveAttr(encodeMove(b7, b8, P, Q, 0, 0, 0, 0), "promoted"))
	//makeMove(encodeMove(e8, g8, k, 0, 0, 0, 0, 1))
	generateAllMoves()
	printBoard()
	for _, move := range moveList {
		printBoard()
		copyBoard()
		makeMove(move)

		printBoard()
		printMove(move)
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		restoreBoard()
	}
}
