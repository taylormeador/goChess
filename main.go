package main

import (
	"fmt"
)

//****************************************************************
//                         constants
//****************************************************************

// Enumerate board squares
const (
	a8 = uint64(iota)
	b8
	c8
	d8
	e8
	f8
	g8
	h8
	a7
	b7
	c7
	d7
	e7
	f7
	g7
	h7
	a6
	b6
	c6
	d6
	e6
	f6
	g6
	h6
	a5
	b5
	c5
	d5
	e5
	f5
	g5
	h5
	a4
	b4
	c4
	d4
	e4
	f4
	g4
	h4
	a3
	b3
	c3
	d3
	e3
	f3
	g3
	h3
	a2
	b2
	c2
	d2
	e2
	f2
	g2
	h2
	a1
	b1
	c1
	d1
	e1
	f1
	g1
	h1
)

// Bitboard masks
const (
	//  e.g. all 0's in the "a" file, all 1's elsewhere
	//
	//  8    0  1  1  1  1  1  1  1
	//  7    0  1  1  1  1  1  1  1
	//  6    0  1  1  1  1  1  1  1
	//  5    0  1  1  1  1  1  1  1
	//  4    0  1  1  1  1  1  1  1
	//  3    0  1  1  1  1  1  1  1
	//  2    0  1  1  1  1  1  1  1
	//  1    0  1  1  1  1  1  1  1
	//
	//       a  b  c  d  e  f  g  h
	//       Bitboard: 18374403900871474942

	notAFile  = uint64(18374403900871474942)
	notHFile  = uint64(9187201950435737471)
	notABFile = uint64(18229723555195321596)
	notGHFile = uint64(4557430888798830399)
)

//****************************************************************
//                     bit manipulations
//****************************************************************

func getBit(b uint64, square uint64) uint64 {
	return b & uint64(1) << square
}

func setBit(b uint64, square uint64) uint64 {
	return b | uint64(1)<<square
}

func popBit(b uint64, square uint64) uint64 {
	return b & ^(uint64(1) << square)
}

//****************************************************************
//                        attacks
//****************************************************************

func maskPawnAttacks(square uint64, side bool) uint64 {
	var attacks uint64
	var leftAttack uint64
	var rightAttack uint64

	// set the pawn on an empty bitboard
	bitboard := setBit(0, square)

	// white pawns move up the board, black pawns move down
	if side {
		leftAttack = bitboard >> 9 & notHFile
		rightAttack = bitboard >> 7 & notAFile
	} else {
		leftAttack = bitboard << 7 & notHFile
		rightAttack = bitboard << 9 & notAFile
	}

	// merge bitboards
	attacks |= leftAttack | rightAttack
	return attacks
}

func initPawnAttacks() [2][64]uint64 {
	var attacks [2][64]uint64
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			square := uint64(rank*8 + file)
			attacks[1][square] = maskPawnAttacks(square, true)
			attacks[0][square] = maskPawnAttacks(square, false)
			printBitboard(attacks[0][square])
		}
	}
	return attacks
}

//****************************************************************
//                         main
//****************************************************************

// prints uint64 as bitboard
func printBitboard(b uint64) {
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
	fmt.Printf("      Bitboard: %d\n\n", b)
}

func main() {
	//bitboard := uint64(0)
	initPawnAttacks()
}
