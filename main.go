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

	parseFEN("r3k2r/pPppqpb1/bn2pnp1/3PN3/Pp2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq b5 0 1")
	printBoard()
	board := copyBoard()

	parseFEN(emptyBoard)
	printBoard()

	restoreBoard(board)
	printBoard()

}
