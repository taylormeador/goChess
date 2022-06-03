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
	parseFEN("q3k2r/p1ppqpb1/bn2pnp1/3PN3/Pp2P3/2N2Q1p/PPPBBPPP/R3K2R b KQkq b5 0 1")

	//for _, move := range moveList {
	//	printBoard()
	//	copyBoard()
	//	makeMove(move)
	//
	//	printBoard()
	//	printMove(move)
	//	input := bufio.NewScanner(os.Stdin)
	//	input.Scan()
	//	restoreBoard()
	//}
}
