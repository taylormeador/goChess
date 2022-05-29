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

	parseFEN("r3k2r/p1ppqpb1/bn2pnp1/3PN3/Pp2P3/2N2Q1p/PPPBBPPP/R3K2R b KQkq - 0 1")
	//parseFEN(startPosition)
	printBoard()

	//for rank := uint64(0); rank < 8; rank++ {
	//	for file := uint64(0); file < 8; file++ {
	//		square := rank*8 + file
	//		generateMoves(square)
	//	}
	//}
	move := encodeMove(e2, e4, k, 1, 1, 0, 0, 0)
	//printBitboard(1 << getMoveAttr(move, 1))
	fmt.Println(getMoveAttr(move, 7))
}
