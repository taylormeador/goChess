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

	parseFEN(trickyPosition)
	printBoard()
	printAttackedSquares(white)
}
