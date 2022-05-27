package main

import "fmt"

func generateMoves(sourceSquare uint64) uint64 {
	startSquare := uint64(1) << sourceSquare
	var targetSquare uint64
	var moves uint64

	// quiet pawn moves
	if side == white {
		if bitboards[P]&startSquare != 0 { // if there is a white pawn on the start square
			// single pawn push
			targetSquare = startSquare >> 8 & ^occupancies[both]
			if targetSquare != 0 {
				fmt.Printf("pawn push: %s%s\n", algebraic[sourceSquare], algebraic[sourceSquare-8])
			}
			// double pawn push
			targetSquare = (startSquare & secondRank) >> 16 & ^occupancies[both]
			squareInFront := startSquare >> 8 & ^occupancies[both]
			if targetSquare != 0 && squareInFront != 0 { // check that the squares in front and two squares in front are empty
				fmt.Printf("double pawn push: %s%s\n", algebraic[sourceSquare], algebraic[sourceSquare-16])
			}
			// pawn promotion
			targetSquare = (startSquare & seventhRank) >> 8 & ^occupancies[both]
			if targetSquare != 0 {
				fmt.Printf("promotion pawn push: %s%s\n", algebraic[sourceSquare], algebraic[sourceSquare-8])
			}
		}
	} else if side == black {
		if bitboards[p]&startSquare != 0 { // if there is a black pawn on the start square
			// single pawn push
			targetSquare = startSquare << 8 & ^occupancies[both]
			if targetSquare != 0 {
				fmt.Printf("pawn push: %s%s\n", algebraic[sourceSquare], algebraic[sourceSquare+8])
			}
			// double pawn push
			targetSquare = (startSquare & seventhRank) << 16 & ^occupancies[both]
			squareInFront := startSquare << 8 & ^occupancies[both]
			if targetSquare != 0 && squareInFront != 0 {
				fmt.Printf("double pawn push: %s%s\n", algebraic[sourceSquare], algebraic[sourceSquare+16])
			}
			// castle
			targetSquare = (startSquare & secondRank) << 8 & ^occupancies[both]
			if targetSquare != 0 {
				fmt.Printf("promotion pawn push: %s%s\n", algebraic[sourceSquare], algebraic[sourceSquare+8])
			}
		}
	}

	//// pawn captures
	//if side == white {
	//	pawnCaptures = pawnAttacks[white][sourceSquare] & occupancies[black]
	//	// en passant
	//	if enPassantSquare != noSquare {
	//		pawnCaptures |= pawnAttacks[white][sourceSquare] & (1 << enPassantSquare)
	//	}
	//} else if side == black {
	//	pawnCaptures = pawnAttacks[black][sourceSquare] & occupancies[white]
	//	// en passant
	//	if enPassantSquare != noSquare {
	//		pawnCaptures |= pawnAttacks[black][sourceSquare] & (1 << enPassantSquare)
	//	}
	//}

	//// castling moves
	//if side == white { // white
	//	if castle&wk != 0 { // kingside
	//		if getBit(occupancies[both], f1) == 0 && getBit(occupancies[both], g1) == 0 { // check that the f1 and g1 squares are empty
	//			if isSquareAttacked(f1, black) == 0 && isSquareAttacked(e1, black) == 0 { // check that f1 and the king are not attacked
	//				targetSquare := "g1"
	//				fmt.Printf("castle move: %s%s\n", algebraic[sourceSquare], targetSquare)
	//			}
	//		}
	//	}
	//	if castle&wq != 0 { // queenside
	//		if getBit(occupancies[both], b1) == 0 && getBit(occupancies[both], c1) == 0 && getBit(occupancies[both], d1) == 0 { // check that the b1, c1, and d1 squares are empty
	//			if isSquareAttacked(d1, black) == 0 && isSquareAttacked(e1, black) == 0 { // check that d1 and the king are not attacked
	//				targetSquare := "c1"
	//				fmt.Printf("castle move: %s%s\n", algebraic[sourceSquare], targetSquare)
	//			}
	//		}
	//	}
	//} else if side == black { // black
	//	if castle&bk != 0 { // kingside
	//		if getBit(occupancies[both], f8) == 0 && getBit(occupancies[both], g8) == 0 { // check that the f8 and g8 squares are empty
	//			if isSquareAttacked(f8, white) == 0 && isSquareAttacked(e8, white) == 0 { // check that f8 and the king are not attacked
	//				targetSquare := "g8"
	//				fmt.Printf("castle move: %s%s\n", algebraic[sourceSquare], targetSquare)
	//			}
	//		}
	//	}
	//	if castle&bq != 0 { // queenside
	//		if getBit(occupancies[both], b8) == 0 && getBit(occupancies[both], c8) == 0 && getBit(occupancies[both], d8) == 0 { // check that the b8, c8, and d8 squares are empty
	//			if isSquareAttacked(d8, white) == 0 && isSquareAttacked(e8, white) == 0 { // check that d8 and the king are not attacked
	//				targetSquare := "c8"
	//				fmt.Printf("castle move: %s%s\n", algebraic[sourceSquare], targetSquare)
	//			}
	//		}
	//	}
	//}

	return moves
}
