package main

import "fmt"

func generateMoves(sourceSquare uint64) {
	startSquare := uint64(1) << sourceSquare
	var targetSquare uint64

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

	// pawn captures
	if side == white {
		if bitboards[P]&startSquare != 0 {
			pawnCaptures := pawnAttacks[white][sourceSquare] & occupancies[black]
			for {
				if pawnCaptures != 0 {
					targetSquare = getLeastSignificantBitIndex(pawnCaptures)
					fmt.Printf("pawn capture: %s%s\n", algebraic[sourceSquare], algebraic[targetSquare])
					pawnCaptures = popBit(pawnCaptures, targetSquare)
				} else {
					break
				}
			}
			// en passant
			if enPassantSquare != noSquare {
				enPassantCapture := pawnAttacks[white][sourceSquare] & (1 << enPassantSquare)
				if enPassantCapture != 0 {
					targetSquare = getLeastSignificantBitIndex(enPassantCapture)
					fmt.Printf("pawn capture en passant: %s%s\n", algebraic[sourceSquare], algebraic[targetSquare])
				}
			}
		}
	} else if side == black {
		if bitboards[p]&startSquare != 0 {
			pawnCaptures := pawnAttacks[black][sourceSquare] & occupancies[white]
			for {
				if pawnCaptures != 0 {
					targetSquare = getLeastSignificantBitIndex(pawnCaptures)
					fmt.Printf("pawn capture: %s%s\n", algebraic[sourceSquare], algebraic[targetSquare])
					pawnCaptures = popBit(pawnCaptures, targetSquare)
				} else {
					break
				}
			}
			// en passant
			if enPassantSquare != noSquare {
				enPassantCapture := pawnAttacks[black][sourceSquare] & (1 << enPassantSquare)
				if enPassantCapture != 0 {
					targetSquare = getLeastSignificantBitIndex(enPassantCapture)
					fmt.Printf("pawn capture en passant: %s%s\n", algebraic[sourceSquare], algebraic[targetSquare])
				}
			}
		}
	}

	// castling moves
	if side == white && sourceSquare == e1 { // white
		if castle&wk != 0 { // kingside
			if getBit(occupancies[both], f1) == 0 && getBit(occupancies[both], g1) == 0 { // check that the f1 and g1 squares are empty
				if isSquareAttacked(f1, black) == 0 && isSquareAttacked(e1, black) == 0 { // check that f1 and the king are not attacked
					fmt.Printf("castle move: %s%s\n", algebraic[sourceSquare], "g1")
				}
			}
		}
		if castle&wq != 0 { // queenside
			if getBit(occupancies[both], b1) == 0 && getBit(occupancies[both], c1) == 0 && getBit(occupancies[both], d1) == 0 { // check that the b1, c1, and d1 squares are empty
				if isSquareAttacked(d1, black) == 0 && isSquareAttacked(e1, black) == 0 { // check that d1 and the king are not attacked
					fmt.Printf("castle move: %s%s\n", algebraic[sourceSquare], "c1")
				}
			}
		}
	} else if side == black && sourceSquare == e8 { // black
		if castle&bk != 0 { // kingside
			if getBit(occupancies[both], f8) == 0 && getBit(occupancies[both], g8) == 0 { // check that the f8 and g8 squares are empty
				if isSquareAttacked(f8, white) == 0 && isSquareAttacked(e8, white) == 0 { // check that f8 and the king are not attacked
					fmt.Printf("castle move: %s%s\n", algebraic[sourceSquare], "e8")
				}
			}
		}
		if castle&bq != 0 { // queenside
			if getBit(occupancies[both], b8) == 0 && getBit(occupancies[both], c8) == 0 && getBit(occupancies[both], d8) == 0 { // check that the b8, c8, and d8 squares are empty
				if isSquareAttacked(d8, white) == 0 && isSquareAttacked(e8, white) == 0 { // check that d8 and the king are not attacked
					fmt.Printf("castle move: %s%s\n", algebraic[sourceSquare], "c8")
				}
			}
		}
	}

	// workin on the Night Moves
	if side == white { // white's turn
		if bitboards[N]&startSquare != 0 { // there is a knight on the sourceSquare
			// quiet knight moves
			quietMoves := knightAttacks[sourceSquare] &^ occupancies[both]
			for {
				if quietMoves != 0 { // knight moves to empty squares
					targetSquare = getLeastSignificantBitIndex(quietMoves)
					fmt.Printf("knight move: %s%s\n", algebraic[sourceSquare], algebraic[targetSquare])
					quietMoves = popBit(quietMoves, targetSquare)
				} else {
					break
				}
			}
			// knight attacks
			attackMoves := knightAttacks[sourceSquare] & occupancies[black]
			for {
				if attackMoves != 0 { // the knight is attacking an enemy piece
					targetSquare = getLeastSignificantBitIndex(attackMoves)
					fmt.Printf("knight capture: %s%s\n", algebraic[sourceSquare], algebraic[targetSquare])
					attackMoves = popBit(attackMoves, targetSquare)
				} else {
					break
				}
			}
		}
	} else if side == black { // black's turn
		if bitboards[n]&startSquare != 0 { // there is a knight on the sourceSquare
			// quiet knight moves
			quietMoves := knightAttacks[sourceSquare] &^ occupancies[both]
			for {
				if quietMoves != 0 { // knight moves to empty squares
					targetSquare = getLeastSignificantBitIndex(quietMoves)
					fmt.Printf("knight move: %s%s\n", algebraic[sourceSquare], algebraic[targetSquare])
					quietMoves = popBit(quietMoves, targetSquare)
				} else {
					break
				}
			}
			// knight attacks
			attackMoves := knightAttacks[sourceSquare] & occupancies[white]
			for {
				if attackMoves != 0 { // the knight is attacking an enemy piece
					targetSquare = getLeastSignificantBitIndex(attackMoves)
					fmt.Printf("knight capture: %s%s\n", algebraic[sourceSquare], algebraic[targetSquare])
					attackMoves = popBit(attackMoves, targetSquare)
				} else {
					break
				}
			}
		}
	}

	// bishop moves
	if side == white && bitboards[B]&startSquare != 0 { // it is white's turn and there's a white bishop on the square
		bishopMoves := getBishopAttacks(sourceSquare, occupancies[both]) & ^occupancies[white]
		for {
			if bishopMoves != 0 {
				targetSquare = getLeastSignificantBitIndex(bishopMoves)
				if occupancies[black]&(1<<targetSquare) != 0 {
					fmt.Printf("bishop capture: %s%s\n", algebraic[sourceSquare], algebraic[targetSquare])
				} else {
					fmt.Printf("bishop move: %s%s\n", algebraic[sourceSquare], algebraic[targetSquare])
				}
				bishopMoves = popBit(bishopMoves, targetSquare)
			} else {
				break
			}
		}
	} else if side == black && bitboards[b]&startSquare != 0 { // it is white's turn and there's a white bishop on the square
		bishopMoves := getBishopAttacks(sourceSquare, occupancies[both]) & ^occupancies[black]
		for {
			if bishopMoves != 0 {
				targetSquare = getLeastSignificantBitIndex(bishopMoves)
				if occupancies[white]&(1<<targetSquare) != 0 {
					fmt.Printf("bishop capture: %s%s\n", algebraic[sourceSquare], algebraic[targetSquare])
				} else {
					fmt.Printf("bishop move: %s%s\n", algebraic[sourceSquare], algebraic[targetSquare])
				}
				bishopMoves = popBit(bishopMoves, targetSquare)
			} else {
				break
			}
		}
	}

	// rook moves
	if side == white && bitboards[R]&startSquare != 0 { // it is white's turn and there is a white rook on the square
		rookMoves := getRookAttacks(sourceSquare, occupancies[both]) & ^occupancies[white]
		for {
			if rookMoves != 0 {
				targetSquare = getLeastSignificantBitIndex(rookMoves)
				if occupancies[black]&(1<<targetSquare) != 0 {
					fmt.Printf("rook capture: %s%s\n", algebraic[sourceSquare], algebraic[targetSquare])
				} else {
					fmt.Printf("rook move: %s%s\n", algebraic[sourceSquare], algebraic[targetSquare])
				}
				rookMoves = popBit(rookMoves, targetSquare)
			} else {
				break
			}
		}
	} else if side == black && bitboards[r]&startSquare != 0 { // it is white's turn and there is a white rook on the square
		rookMoves := getRookAttacks(sourceSquare, occupancies[both]) & ^occupancies[black]
		for {
			if rookMoves != 0 {
				targetSquare = getLeastSignificantBitIndex(rookMoves)
				if occupancies[white]&(1<<targetSquare) != 0 {
					fmt.Printf("rook capture: %s%s\n", algebraic[sourceSquare], algebraic[targetSquare])
				} else {
					fmt.Printf("rook move: %s%s\n", algebraic[sourceSquare], algebraic[targetSquare])
				}
				rookMoves = popBit(rookMoves, targetSquare)
			} else {
				break
			}
		}
	}

	// queen moves
	if side == white && bitboards[Q]&startSquare != 0 { // it is white's turn and there is a white queen on the square
		queenMoves := getQueenAttacks(sourceSquare, occupancies[both]) & ^occupancies[white]
		for {
			if queenMoves != 0 {
				targetSquare = getLeastSignificantBitIndex(queenMoves)
				if occupancies[black]&(1<<targetSquare) != 0 { // queen attacking black piece
					fmt.Printf("queen capture: %s%s\n", algebraic[sourceSquare], algebraic[targetSquare])
				} else {
					fmt.Printf("queen move: %s%s\n", algebraic[sourceSquare], algebraic[targetSquare])
				}
				queenMoves = popBit(queenMoves, targetSquare)
			} else {
				break
			}
		}
	} else if side == black && bitboards[q]&startSquare != 0 { // it is white's turn and there is a white queen on the square
		queenMoves := getQueenAttacks(sourceSquare, occupancies[both]) & ^occupancies[black]
		for {
			if queenMoves != 0 {
				targetSquare = getLeastSignificantBitIndex(queenMoves)
				if occupancies[black]&(1<<targetSquare) != 0 { // queen attacking black piece
					fmt.Printf("queen capture: %s%s\n", algebraic[sourceSquare], algebraic[targetSquare])
				} else {
					fmt.Printf("queen move: %s%s\n", algebraic[sourceSquare], algebraic[targetSquare])
				}
				queenMoves = popBit(queenMoves, targetSquare)
			} else {
				break
			}
		}
	}
}
