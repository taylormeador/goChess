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

	for rank := uint64(0); rank < 8; rank++ {
		for file := uint64(0); file < 8; file++ {
			square := rank*8 + file
			generateMoves(square)
		}
	}
	printMoveList()
	printBoard()

}
