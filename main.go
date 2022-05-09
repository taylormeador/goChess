package main

import (
	"fmt"
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
	fmt.Printf("  Bitboard: %d\n\n", b)
}

// init
func initAll() {
	initLeapersAttacks()
	initSliderAttacks(bishop)
	initSliderAttacks(rook)
}

// main
func main() {
	initAll()

	// define test bitboard
	occupancy := uint64(0)
	occupancy |= setBit(occupancy, c5)
	occupancy |= setBit(occupancy, f2)
	occupancy |= setBit(occupancy, g7)
	occupancy |= setBit(occupancy, b2)
	occupancy |= setBit(occupancy, g5)
	occupancy |= setBit(occupancy, e2)
	occupancy |= setBit(occupancy, e7)

	// print
	printBitboard(occupancy)
	printBitboard(getBishopAttacks(d4, occupancy))
	printBitboard(getRookAttacks(e5, occupancy))
}
