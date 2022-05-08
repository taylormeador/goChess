package main

import (
	"fmt"
)

//****************************************************************
//                         constants
//****************************************************************

// white = 1 or true, black = 0 or false
const (
	white = 1
	black = 0
)

// enumerate board squares
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

// bitboard masks
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

var algebraic = [64]string{
	"a8", "b8", "c8", "d8", "e8", "f8", "g8", "h8",
	"a7", "b7", "c7", "d7", "e7", "f7", "g7", "h7",
	"a6", "b6", "c6", "d6", "e6", "f6", "g6", "h6",
	"a5", "b5", "c5", "d5", "e5", "f5", "g5", "h5",
	"a4", "b4", "c4", "d4", "e4", "f4", "g4", "h4",
	"a3", "b3", "c3", "d3", "e3", "f3", "g3", "h3",
	"a2", "b2", "c2", "d2", "e2", "f2", "g2", "h2",
	"a1", "b1", "c1", "d1", "e1", "f1", "g1", "h1",
}

//****************************************************************
//                     bit manipulations
//****************************************************************

// check if a bit is on or off
func getBit(b uint64, square uint64) uint64 {
	return b & uint64(1) << square
}

// turn on a bit
func setBit(b uint64, square uint64) uint64 {
	return b | uint64(1)<<square
}

// turn off a bit
func popBit(b uint64, square uint64) uint64 {
	return b & ^(uint64(1) << square)
}

// count the number of bits on a bitboard
func countBits(b uint64) uint64 {
	count := uint64(0)
	for {
		b &= b - 1
		count += 1
		if b == 0 {
			break
		}
	}
	return count
}

func getLeastSignificantBitIndex(b uint64) uint64 {
	leastSignificantBit := b & -b
	leadingOnes := leastSignificantBit - 1
	return countBits(leadingOnes)
}

//****************************************************************
//                          attacks
//****************************************************************

// loop through all the squares on the board generating attacks for pawns, knights, and kings on those squares
func initSliderAttacks() {
	var pawnAttacks [2][64]uint64
	var knightAttacks [64]uint64
	var kingAttacks [64]uint64
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			square := uint64(rank*8 + file)
			pawnAttacks[white][square] = maskPawnAttacks(square, white)
			pawnAttacks[black][square] = maskPawnAttacks(square, black)
			knightAttacks[square] = maskKnightAttacks(square)
			kingAttacks[square] = maskKingAttacks(square)
		}
	}
}

// generate pawn attacks for a single square
func maskPawnAttacks(square uint64, side int) uint64 {
	var attacks uint64
	var leftAttack uint64
	var rightAttack uint64

	// set the pawn on an empty bitboard
	bitboard := setBit(0, square)

	// white pawns move up the board, black pawns move down
	// file masks prevent off board attacks
	if side == white {
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

// generate knight attacks for a single square
func maskKnightAttacks(square uint64) uint64 {
	var attacks uint64

	// set the knight on an empty bitboard
	bitboard := setBit(0, square)

	// add offsets to attacks, leaving out off board attacks
	attacks |= bitboard >> 15 & notAFile
	attacks |= bitboard >> 6 & notABFile
	attacks |= bitboard << 10 & notABFile
	attacks |= bitboard << 17 & notAFile
	attacks |= bitboard << 15 & notHFile
	attacks |= bitboard << 6 & notGHFile
	attacks |= bitboard >> 10 & notGHFile
	attacks |= bitboard >> 17 & notHFile

	return attacks
}

// generate king attacks for a single square
func maskKingAttacks(square uint64) uint64 {
	var attacks uint64

	// set the king on an empty bitboard
	bitboard := setBit(0, square)

	// on the on board attacks, starting from 12 o clock and moving clockwise
	attacks |= bitboard >> 8
	attacks |= bitboard >> 7 & notAFile
	attacks |= bitboard << 1 & notAFile
	attacks |= bitboard << 9 & notAFile
	attacks |= bitboard << 8
	attacks |= bitboard << 7 & notHFile
	attacks |= bitboard >> 1 & notHFile
	attacks |= bitboard >> 9 & notHFile

	return attacks
}

// generate bishop attack mask for a single square for magic bitboard
func maskBishopAttacks(square uint64) uint64 {
	var attacks uint64

	// current rank and file
	currentRank := square / 8
	currentFile := square % 8

	// loop through squares in the four directions, stopping before hitting the edge of the board
	for targetRank, targetFile := currentRank+1, currentFile+1; targetRank < 7 && targetFile < 7; targetRank, targetFile = targetRank+1, targetFile+1 {
		targetSquare := targetRank*8 + targetFile
		attacks |= uint64(1) << targetSquare
	}

	for targetRank, targetFile := currentRank-1, currentFile+1; targetRank > 0 && targetFile < 7; targetRank, targetFile = targetRank-1, targetFile+1 {
		targetSquare := targetRank*8 + targetFile
		attacks |= uint64(1) << targetSquare
	}

	for targetRank, targetFile := currentRank-1, currentFile-1; targetRank > 0 && targetFile < 7 && targetFile > 0; targetRank, targetFile = targetRank-1, targetFile-1 {
		targetSquare := targetRank*8 + targetFile
		attacks |= uint64(1) << targetSquare
	}

	for targetRank, targetFile := currentRank+1, currentFile-1; targetRank < 7 && targetFile < 7 && targetFile > 0; targetRank, targetFile = targetRank+1, targetFile-1 {
		targetSquare := targetRank*8 + targetFile
		attacks |= uint64(1) << targetSquare
	}
	return attacks
}

// generate bishop attacks for a single square on the fly
func bishopAttacksOnTheFly(square uint64, blockers uint64) uint64 {
	var attacks uint64

	// current rank and file
	currentRank := square / 8
	currentFile := square % 8

	// loop through squares in the four directions, stopping after you hit a blocker
	for targetRank, targetFile := currentRank+1, currentFile+1; targetRank <= 7 && targetFile <= 7; targetRank, targetFile = targetRank+1, targetFile+1 {
		targetSquare := targetRank*8 + targetFile
		targetBitboard := uint64(1) << targetSquare
		attacks |= targetBitboard
		if (targetBitboard & blockers) != 0 {
			break
		}
	}

	for targetRank, targetFile := currentRank-1, currentFile+1; targetRank >= 0 && targetFile <= 7; targetRank, targetFile = targetRank-1, targetFile+1 {
		targetSquare := targetRank*8 + targetFile
		targetBitboard := uint64(1) << targetSquare
		attacks |= targetBitboard
		if (targetBitboard & blockers) != 0 {
			break
		}
	}

	for targetRank, targetFile := currentRank-1, currentFile-1; targetRank >= 0 && targetFile <= 7 && targetFile >= 0; targetRank, targetFile = targetRank-1, targetFile-1 {
		targetSquare := targetRank*8 + targetFile
		targetBitboard := uint64(1) << targetSquare
		attacks |= targetBitboard
		if (targetBitboard & blockers) != 0 {
			break
		}
	}

	for targetRank, targetFile := currentRank+1, currentFile-1; targetRank <= 7 && targetFile <= 7 && targetFile >= 0; targetRank, targetFile = targetRank+1, targetFile-1 {
		targetSquare := targetRank*8 + targetFile
		targetBitboard := uint64(1) << targetSquare
		attacks |= targetBitboard
		if (targetBitboard & blockers) != 0 {
			break
		}
	}
	return attacks
}

// generate rook attack mask for a single square for magic bitboard
func maskRookAttacks(square uint64) uint64 {
	var attacks uint64

	// current rank and file
	currentRank := square / 8
	currentFile := square % 8

	// loop through squares in the four directions, stopping before hitting the edge of the board
	for targetRank := currentRank + 1; targetRank < 7; targetRank++ {
		targetSquare := targetRank*8 + currentFile
		attacks |= uint64(1) << targetSquare

	}
	for targetRank := currentRank - 1; targetRank > 0 && targetRank < 7; targetRank-- {
		targetSquare := targetRank*8 + currentFile
		attacks |= uint64(1) << targetSquare

	}

	for targetFile := currentFile - 1; targetFile > 0 && targetFile < 7; targetFile-- {
		targetSquare := currentRank*8 + targetFile
		attacks |= uint64(1) << targetSquare

	}

	for targetFile := currentFile + 1; targetFile < 7; targetFile++ {
		targetSquare := currentRank*8 + targetFile
		attacks |= uint64(1) << targetSquare
	}
	printBitboard(attacks)
	return attacks
}

// generate rook attacks for a single square oin the fly
func rookAttacksOnTheFly(square uint64, blockers uint64) uint64 {
	var attacks uint64

	// current rank and file
	currentRank := square / 8
	currentFile := square % 8

	// loop through squares in the four directions, stopping after hitting a blocker
	for targetRank := currentRank + 1; targetRank <= 7; targetRank++ {
		targetSquare := targetRank*8 + currentFile
		targetBitboard := uint64(1) << targetSquare
		attacks |= targetBitboard
		if (targetBitboard & blockers) != 0 {
			break
		}
	}

	for targetRank := currentRank - 1; targetRank >= 0 && targetRank <= 7; targetRank-- {
		targetSquare := targetRank*8 + currentFile
		targetBitboard := uint64(1) << targetSquare
		attacks |= targetBitboard
		if (targetBitboard & blockers) != 0 {
			break
		}
	}

	for targetFile := currentFile - 1; targetFile >= 0 && targetFile <= 7; targetFile-- {
		targetSquare := currentRank*8 + targetFile
		targetBitboard := uint64(1) << targetSquare
		attacks |= targetBitboard
		if (targetBitboard & blockers) != 0 {
			break
		}
	}

	for targetFile := currentFile + 1; targetFile <= 7; targetFile++ {
		targetSquare := currentRank*8 + targetFile
		targetBitboard := uint64(1) << targetSquare
		attacks |= targetBitboard
		if (targetBitboard & blockers) != 0 {
			break
		}
	}
	printBitboard(attacks)
	return attacks
}

//****************************************************************
//                           main
//****************************************************************

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
	fmt.Printf("    Index: %d   Coordinate: %s\n\n", 4, algebraic[4])
}

// main
func main() {
	//bitboard := uint64(0)
	//pawnAttacks := initPawnAttacks()
	//knightAttacks := initKnightAttacks()

	if false {
		for rank := 0; rank < 8; rank++ {
			for file := 0; file < 8; file++ {
				square := uint64(rank*8 + file)
				rookAttacksOnTheFly(square, d6)
			}
		}
	}

	bitboard := uint64(0)
	bitboard = setBit(bitboard, e4)
	bitboard = setBit(bitboard, h1)
	//bitboard = setBit(bitboard, a8)
	bitboard = setBit(bitboard, e1)
	bitboard = setBit(bitboard, e7)
	bitboard = setBit(bitboard, a1)
	printBitboard(bitboard)
}
