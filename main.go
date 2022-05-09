package main

import (
	"bufio"
	"fmt"
	"os"
)

// prints uint64 as bitboard
func printBitboard(b uint64) {
	fmt.Println()
	// loop through ranks and files
	for rank := uint64(0); rank < 8; rank++ {
		fmt.Printf(" %d   ", 8-rank)
		for file := uint64(0); file < 8; file++ {
			// convert rank and file to index
			index := rank*8 + file

			// check whether the bit should be on or off
			printBit := 0
			if b&(uint64(1)<<index) != 0 {
				printBit = 1
			}

			// print 1 or 0 based on previously generated bool
			fmt.Printf(" %d ", printBit)
		}
		fmt.Println()
	}
	fmt.Println()
	// print files and bitboard integer value
	fmt.Println("      a  b  c  d  e  f  g  h")
	fmt.Printf("  Bitboard: %d\n", b)
	lsbIndex := getLeastSignificantBitIndex(b)
	lsbAlgebraic := "Out of range"
	if lsbIndex <= 63 && lsbIndex >= 0 {
		lsbAlgebraic = algebraic[getLeastSignificantBitIndex(b)]
	}
	fmt.Printf("    Index: %d   Coordinate: %s\n\n", lsbIndex, lsbAlgebraic)
}

// main
func main() {
	//bitboard := uint64(0)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(getRandomNumber())
		reader.ReadString('\n')
	}
}
