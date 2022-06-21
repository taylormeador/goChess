package main

import (
	"fmt"
	"time"
)

// init
func initAll() {
	initLeapersAttacks()
	initSlidersAttacks(bishop)
	initSlidersAttacks(rook)
}

// main
func main() {
	startTime := time.Now()
	initAll()
	parseFEN("Q3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R b KQkq - 0 1")
	printBoard()
	perftDriver(3)

	//for _, move := range moveList {
	//	copyBoard()
	//	//input := bufio.NewScanner(os.Stdin)
	//	//input.Scan()
	//	if makeMove(move) == 0 {
	//		continue
	//	}
	//	makeMove(move)
	//	restoreBoard()
	//}
	endTime := time.Now()
	fmt.Printf("Time taken: %d", endTime.Sub(startTime))
}
