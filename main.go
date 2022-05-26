package main

// init
func initAll() {
	initLeapersAttacks()
	initSlidersAttacks(bishop)
	initSlidersAttacks(rook)
}

// main
func main() {
	initAll()

	parseFEN("rnbqkbnr/pppppppp/8/8/2PpP3/8/PPPPPPPP/RNBQKBNR b KQkq e3 0 1")
	//parseFEN(trickyPosition)
	printBoard()
	printBitboard(generateMoves(d4))

}
