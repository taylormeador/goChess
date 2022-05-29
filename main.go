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

	addMove(encodeMove(e2, e4, k, Q, 0, 0, 0, 0))
	addMove(encodeMove(e2, e5, P, 0, 1, 0, 0, 0))
	addMove(encodeMove(a1, h8, B, 0, 0, 0, 0, 0))
	printMoveList()
}
