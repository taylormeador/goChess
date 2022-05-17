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
	printBitboard(getQueenAttacks(d5, uint64(18303478847064064385)))
}
