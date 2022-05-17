package main

import (
	"fmt"
	"strconv"
)

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

// parse FEN string
func parseFEN(FEN string) {
	// reset variables
	bitboards = [12]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	occupancyBitboards = [3]uint64{0, 0, 0}
	side = 0
	enPassantSquare = noSquare
	castle = 0

	counter := uint64(0)
	for _, value := range FEN {
		// debug fmt.Printf("%d, %d, %q\n", i, counter, value)
		// match letters with pieces
		if (value >= 'a' && value <= 'z') || (value >= 'A' && value <= 'Z') {
			// get the piece
			piece := charPieces[byte(value)]
			if counter < 64 {
				bitboards[piece] = setBit(bitboards[piece], counter)
			} else if counter == 65 { // side to move
				intValue, _ := strconv.Atoi(string(value))
				side = intValue

			} else if counter > 65 && counter < 71 { // castling rights
				castle |= pieceToCastle[byte(value)]
			} else if counter == 73 { // en passant square
				intValue, _ := strconv.Atoi(string(value))
				enPassantSquare = uint64(intValue)
			}

		} else if value == '/' { // new rank
			continue
		} else { // skip ahead the number of empty squares
			intValue, _ := strconv.Atoi(string(value))
			if intValue <= 8 && intValue >= 1 {
				counter += uint64(intValue - 1)
			}
		}
		counter++
	}

	printBoard()
	fmt.Printf("side: %d\n", side)
	fmt.Printf("castling: %d\n", castle)
	fmt.Printf("en-passant: %d\n", enPassantSquare)

	/*// loop through all the squares
	for rank := uint64(0); rank < 8; rank++ {
		for file := uint64(0); file < 8; file++ {
			square := rank*8 + file
			if (FEN >= 'a' && FEN <= 'z') || (FEN >= 'A' && FEN <= 'Z') {
				// get the piece
				piece := charPieces[FEN]

				// place the piece on the square
				bitboards[piece] = setBit(bitboards[piece], square)
			}
		}
	}*/
}
