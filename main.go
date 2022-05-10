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
	bitboards[P] = setBit(bitboards[P], e2)
	printBitboard(bitboards[P])
	printBoard()
}
