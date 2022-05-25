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
	//printAttackedSquares(white)
	//printAttackedSquares(black)
	parseFEN(emptyBoard)
	printBoard()
	parseFEN(startPosition)
	printBoard()
	parseFEN(trickyPosition)
	printBoard()
	parseFEN(killerPosition)
	printBoard()
	parseFEN(cmkPosition)
	printBoard()
}
