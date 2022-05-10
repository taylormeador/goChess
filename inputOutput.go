package main

import "fmt"

// prints uint64 as bitboard
func printBitboard(bitboard uint64) {
	fmt.Println()
	// loop through ranks and files
	for rank := uint64(0); rank < 8; rank++ {
		fmt.Printf(" %d   ", 8-rank)
		for file := uint64(0); file < 8; file++ {
			// convert rank and file to index
			index := rank*8 + file

			// check whether the bit should be on or off
			printBit := 0
			if bitboard&(uint64(1)<<index) != 0 {
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
	fmt.Printf("  Bitboard: %d\n\n", bitboard)
}

// print board
func printBoard() {
	fmt.Println()
	for rank := uint64(0); rank < 8; rank++ {
		fmt.Printf(" %d   ", 8-rank)
		for file := uint64(0); file < 8; file++ {
			square := rank*8 + file

			piece := -1

			for bitboardPiece := P; bitboardPiece <= k; bitboardPiece++ {
				if getBit(bitboards[bitboardPiece], square) != 0 {
					piece = bitboardPiece
				}
			}

			if piece == -1 {
				fmt.Printf(" %s ", ".")
			} else {
				fmt.Printf(" %s ", unicodePieces[piece])
			}

		}
		fmt.Println()
	}
	fmt.Println()
	// print files and bitboard integer value
	fmt.Printf("      a  b  c  d  e  f  g  h\n\n")
}
