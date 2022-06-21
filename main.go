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

	initAll()
	parseFEN(trickyPosition)
	printBoard()
	startTime := time.Now()
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
	fmt.Printf("%d nodes\n", nodes)
	fmt.Printf("Time taken: %d", endTime.Sub(startTime))
}
