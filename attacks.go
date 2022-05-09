package main

// loop through all the squares on the board generating attacks for pawns, knights, and kings on those squares
func initLeapersAttacks() {
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
	return attacks
}

// set relevant occupancy bits
func setOccupancy(index uint64, bitsInMask uint64, attackMask uint64) uint64 {
	var occupancy uint64

	// loop over bits in attack mask
	for count := uint64(0); count < bitsInMask; count++ {
		square := getLeastSignificantBitIndex(attackMask)
		attackMask = popBit(attackMask, square)
		if (index & (uint64(1) << count)) != 0 {
			occupancy |= uint64(1) << square
		}
	}
	return occupancy
}
