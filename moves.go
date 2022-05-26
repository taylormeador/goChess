package main

func generateMoves(sourceSquare uint64) uint64 {
	var singlePawnMoves uint64
	var doublePawnMoves uint64
	var promotionPawnMoves uint64
	var pawnCaptures uint64

	// quiet pawn moves
	if side == white {
		singlePawnMoves = bitboards[P] >> 8 & ^occupancies[both]
		doublePawnMoves = (bitboards[P] & secondRank) >> 16 & ^occupancies[both]
		promotionPawnMoves = (bitboards[P] & seventhRank) >> 8 & ^occupancies[both]

	} else if side == black {
		singlePawnMoves = bitboards[p] << 8 & ^occupancies[both]
		doublePawnMoves = (bitboards[p] & seventhRank) << 16 & ^occupancies[both]
		promotionPawnMoves = (bitboards[p] & secondRank) << 8 & ^occupancies[both]
	}

	// pawn captures
	if side == white {
		pawnCaptures = pawnAttacks[white][sourceSquare] & occupancies[black]
		// en passant
		if enPassantSquare != noSquare {
			pawnCaptures |= pawnAttacks[white][sourceSquare] & (1 << enPassantSquare)
		}
	} else if side == black {
		pawnCaptures = pawnAttacks[black][sourceSquare] & occupancies[white]
		// en passant
		if enPassantSquare != noSquare {
			pawnCaptures |= pawnAttacks[black][sourceSquare] & (1 << enPassantSquare)
		}
	}
	return doublePawnMoves | singlePawnMoves | promotionPawnMoves | pawnCaptures
}
