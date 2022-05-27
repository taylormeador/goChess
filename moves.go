package main

import "fmt"

func generateMoves(sourceSquare uint64) uint64 {
	var singlePawnMoves uint64
	var doublePawnMoves uint64
	var promotionPawnMoves uint64
	var pawnCaptures uint64

	// quiet pawn moves
	if side == white {
		singlePawnMoves = sourceSquare >> 8 & ^occupancies[both]
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

	// castling moves
	if side == white { // white
		if castle&wk != 0 { // kingside
			if getBit(occupancies[both], f1) == 0 && getBit(occupancies[both], g1) == 0 { // check that the f1 and g1 squares are empty
				if isSquareAttacked(f1, black) == 0 && isSquareAttacked(e1, black) == 0 { // check that f1 and the king are not attacked
					targetSquare := "g1"
					fmt.Printf("castle move: %s%s\n", algebraic[sourceSquare], targetSquare)
				}
			}
		}
		if castle&wq != 0 { // queenside
			if getBit(occupancies[both], b1) == 0 && getBit(occupancies[both], c1) == 0 && getBit(occupancies[both], d1) == 0 { // check that the b1, c1, and d1 squares are empty
				if isSquareAttacked(d1, black) == 0 && isSquareAttacked(e1, black) == 0 { // check that d1 and the king are not attacked
					targetSquare := "c1"
					fmt.Printf("castle move: %s%s\n", algebraic[sourceSquare], targetSquare)
				}
			}
		}
	} else if side == black { // black
		if castle&bk != 0 { // kingside
			if getBit(occupancies[both], f8) == 0 && getBit(occupancies[both], g8) == 0 { // check that the f8 and g8 squares are empty
				if isSquareAttacked(f8, white) == 0 && isSquareAttacked(e8, white) == 0 { // check that f8 and the king are not attacked
					targetSquare := "g8"
					fmt.Printf("castle move: %s%s\n", algebraic[sourceSquare], targetSquare)
				}
			}
		}
		if castle&bq != 0 { // queenside
			if getBit(occupancies[both], b8) == 0 && getBit(occupancies[both], c8) == 0 && getBit(occupancies[both], d8) == 0 { // check that the b8, c8, and d8 squares are empty
				if isSquareAttacked(d8, white) == 0 && isSquareAttacked(e8, white) == 0 { // check that d8 and the king are not attacked
					targetSquare := "c8"
					fmt.Printf("castle move: %s%s\n", algebraic[sourceSquare], targetSquare)
				}
			}
		}
	}

	moves := doublePawnMoves | singlePawnMoves | promotionPawnMoves | pawnCaptures
	return moves
}
