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
	parsePosition("position fen r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1 move a2a4")
	printBoard()

	//parseFEN(trickyPosition)
	//move := parseMove("b7b8b")
	//if move != 0 {
	//	makeMove(move)
	//	printMove(move)
	//} else {
	//	fmt.Printf("Illegal move")
	//}
}
