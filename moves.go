package main

import "fmt"

var moveList []uint64

func generateMoves(sourceSquare uint64) {
	startSquare := uint64(1) << sourceSquare
	var targetSquare, move uint64
	var promotionRank, promotionRankMinusOne, pawnStartRank, pawnPush, doublePawnPush, pawnPromotion uint64
	var pawnPushOffset, doublePawnPushOffset, kingStartSquare uint64
	var pawn, knight, bishop, rook, queen, king uint64
	var bSquare, cSquare, dSquare, fSquare, gSquare uint64
	var castleKingside, castleQueenside, enemyColor int

	// assign side related variables
	if side == white {
		pawn = P
		knight = N
		bishop = B
		rook = R
		queen = Q
		king = K
		enemyColor = black
		promotionRank = eighthRank
		promotionRankMinusOne = seventhRank
		pawnStartRank = secondRank
		pawnPush = startSquare >> 8
		pawnPushOffset = sourceSquare - 8
		doublePawnPush = (startSquare & pawnStartRank) >> 16
		doublePawnPushOffset = sourceSquare - 16
		pawnPromotion = (startSquare & promotionRankMinusOne) >> 8
		castleKingside = castle & wk
		castleQueenside = castle & wq
		kingStartSquare = e1
		bSquare = b1
		cSquare = c1
		dSquare = d1
		fSquare = f1
		gSquare = g1

	} else if side == black {
		pawn = p
		knight = n
		bishop = b
		rook = r
		queen = q
		king = k
		enemyColor = white
		promotionRank = firstRank
		promotionRankMinusOne = secondRank
		pawnStartRank = seventhRank
		pawnPush = startSquare << 8
		pawnPushOffset = sourceSquare + 8
		doublePawnPush = (startSquare & pawnStartRank) << 16
		doublePawnPushOffset = sourceSquare + 16
		pawnPromotion = (startSquare & promotionRankMinusOne) << 8
		castleKingside = castle & bk
		castleQueenside = castle & bq
		kingStartSquare = e8
		bSquare = b8
		cSquare = c8
		dSquare = d8
		fSquare = f8
		gSquare = g8
	}

	// quiet pawn moves
	if bitboards[pawn]&startSquare != 0 { // if there is a white pawn on the start square
		// single pawn push
		targetSquare = (pawnPush & ^occupancies[both]) & ^promotionRank
		if targetSquare != 0 {
			move = encodeMove(sourceSquare, pawnPushOffset, pawn, 0, 0, 0, 0, 0)
			addMove(move)
		}
		// double pawn push
		targetSquare = doublePawnPush & ^occupancies[both]
		squareInFront := pawnPush & ^occupancies[both]
		if targetSquare != 0 && squareInFront != 0 { // check that the squares in front and two squares in front are empty
			addMove(encodeMove(sourceSquare, doublePawnPushOffset, pawn, 0, 0, 1, 0, 0))
		}
		// pawn promotion
		targetSquare = pawnPromotion & ^occupancies[both]
		if targetSquare != 0 {
			addMove(encodeMove(sourceSquare, pawnPushOffset, pawn, queen, 0, 0, 0, 0))
			addMove(encodeMove(sourceSquare, pawnPushOffset, pawn, rook, 0, 0, 0, 0))
			addMove(encodeMove(sourceSquare, pawnPushOffset, pawn, bishop, 0, 0, 0, 0))
			addMove(encodeMove(sourceSquare, pawnPushOffset, pawn, knight, 0, 0, 0, 0))
		}
	}

	// pawn captures
	if bitboards[pawn]&startSquare != 0 {
		pawnCaptures := pawnAttacks[side][sourceSquare] & occupancies[enemyColor]
		for {
			if pawnCaptures != 0 {
				targetSquare = getLeastSignificantBitIndex(pawnCaptures)
				if startSquare&promotionRankMinusOne != 0 { // promotion capture
					addMove(encodeMove(sourceSquare, targetSquare, pawn, queen, 1, 0, 0, 0))
					addMove(encodeMove(sourceSquare, targetSquare, pawn, rook, 1, 0, 0, 0))
					addMove(encodeMove(sourceSquare, targetSquare, pawn, knight, 1, 0, 0, 0))
					addMove(encodeMove(sourceSquare, targetSquare, pawn, bishop, 1, 0, 0, 0))
				} else { // regular capture
					addMove(encodeMove(sourceSquare, targetSquare, pawn, 0, 1, 0, 0, 0))
				}
				pawnCaptures = popBit(pawnCaptures, targetSquare)
			} else {
				break
			}
		}
		// en passant
		if enPassantSquare != noSquare {
			enPassantCapture := pawnAttacks[side][sourceSquare] & (1 << enPassantSquare)
			if enPassantCapture != 0 {
				targetSquare = getLeastSignificantBitIndex(enPassantCapture)
				addMove(encodeMove(sourceSquare, targetSquare, pawn, 0, 1, 0, 0, 0))
			}
		}
	}

	// castling moves
	if sourceSquare == kingStartSquare { // white
		if castleKingside != 0 { // kingside
			if getBit(occupancies[both], fSquare) == 0 && getBit(occupancies[both], gSquare) == 0 { // check that the kingside squares are empty
				if isSquareAttacked(fSquare, enemyColor) == 0 && isSquareAttacked(kingStartSquare, enemyColor) == 0 { // check that travel square and the king are not attacked
					move = encodeMove(sourceSquare, gSquare, king, 0, 0, 0, 0, 1)
					addMove(move)
				}
			}
		}
		if castleQueenside != 0 { // queenside
			if getBit(occupancies[both], bSquare) == 0 && getBit(occupancies[both], cSquare) == 0 && getBit(occupancies[both], dSquare) == 0 { // check that the queenside squares are empty
				if isSquareAttacked(dSquare, enemyColor) == 0 && isSquareAttacked(kingStartSquare, enemyColor) == 0 { // check that travel squares and the king are not attacked
					move = encodeMove(sourceSquare, cSquare, king, 0, 0, 0, 0, 1)
					addMove(move)
				}
			}
		}
	}

	// workin on the Night Moves
	if bitboards[knight]&startSquare != 0 { // there is a knight on the sourceSquare
		// quiet knight moves
		quietMoves := knightAttacks[sourceSquare] &^ occupancies[both]
		for {
			if quietMoves != 0 { // knight moves to empty squares
				targetSquare = getLeastSignificantBitIndex(quietMoves)
				addMove(encodeMove(sourceSquare, targetSquare, knight, 0, 0, 0, 0, 0))
				quietMoves = popBit(quietMoves, targetSquare)
			} else {
				break
			}
		}
		// knight attacks
		attackMoves := knightAttacks[sourceSquare] & occupancies[enemyColor]
		for {
			if attackMoves != 0 { // the knight is attacking an enemy piece
				targetSquare = getLeastSignificantBitIndex(attackMoves)
				addMove(encodeMove(sourceSquare, targetSquare, knight, 0, 1, 0, 0, 0))
				attackMoves = popBit(attackMoves, targetSquare)
			} else {
				break
			}
		}
	}

	// bishop moves
	if bitboards[bishop]&startSquare != 0 { // if there's a bishop on the square
		bishopMoves := getBishopAttacks(sourceSquare, occupancies[both]) & ^occupancies[side]
		for {
			if bishopMoves != 0 {
				targetSquare = getLeastSignificantBitIndex(bishopMoves)
				if occupancies[enemyColor]&(1<<targetSquare) != 0 {
					addMove(encodeMove(sourceSquare, targetSquare, bishop, 0, 1, 0, 0, 0))
				} else {
					addMove(encodeMove(sourceSquare, targetSquare, bishop, 0, 0, 0, 0, 0))
				}
				bishopMoves = popBit(bishopMoves, targetSquare)
			} else {
				break
			}
		}
	}

	// rook moves
	if bitboards[rook]&startSquare != 0 { // there is a rook on the square
		rookMoves := getRookAttacks(sourceSquare, occupancies[both]) & ^occupancies[side]
		for {
			if rookMoves != 0 {
				targetSquare = getLeastSignificantBitIndex(rookMoves)
				if occupancies[enemyColor]&(1<<targetSquare) != 0 {
					addMove(encodeMove(sourceSquare, targetSquare, rook, 0, 1, 0, 0, 0))
				} else {
					addMove(encodeMove(sourceSquare, targetSquare, rook, 0, 0, 0, 0, 0))
				}
				rookMoves = popBit(rookMoves, targetSquare)
			} else {
				break
			}
		}
	}

	// queen moves
	if bitboards[queen]&startSquare != 0 { // it is white's turn and there is a white queen on the square
		queenMoves := getQueenAttacks(sourceSquare, occupancies[both]) & ^occupancies[side]
		for {
			if queenMoves != 0 {
				targetSquare = getLeastSignificantBitIndex(queenMoves)
				if occupancies[enemyColor]&(1<<targetSquare) != 0 { // queen attacking enemy piece
					addMove(encodeMove(sourceSquare, targetSquare, queen, 0, 1, 0, 0, 0))
				} else {
					addMove(encodeMove(sourceSquare, targetSquare, queen, 0, 0, 0, 0, 0))
				}
				queenMoves = popBit(queenMoves, targetSquare)
			} else {
				break
			}
		}
	}
}

/*  encode moves in binary
      binary move bits                               hexidecimal constants

0000 0000 0000 0000 0011 1111    source square       0x3f
0000 0000 0000 1111 1100 0000    target square       0xfc0
0000 0000 1111 0000 0000 0000    piece               0xf000
0000 1111 0000 0000 0000 0000    promoted piece      0xf0000
0001 0000 0000 0000 0000 0000    capture flag        0x100000
0010 0000 0000 0000 0000 0000    double push flag    0x200000
0100 0000 0000 0000 0000 0000    enpassant flag      0x400000
1000 0000 0000 0000 0000 0000    castling flag       0x800000
*/
func encodeMove(source uint64, target uint64, piece uint64, promoted uint64, capture uint64,
	double uint64, enPassant uint64, castling uint64) uint64 {

	move := source | (target << 6) | (piece << 12) | (promoted << 16) | (capture << 20) |
		(double << 21) | (enPassant << 22) | (castling << 23)

	return move
}

// return the specified attribute of an encoded move
func getMoveAttr(move uint64, attr string) uint64 {
	switch attr {
	case "source":
		return move & 0x3f
	case "target": // target
		return (move & 0xfc0) >> 6
	case "piece": // piece
		return (move & 0xf000) >> 12
	case "promoted": // promoted
		return (move & 0xf0000) >> 16
	case "capture": // capture
		if move&0x100000 != 0 {
			return 1
		}
		return 0

	case "double": // double
		if move&0x200000 != 0 {
			return 1
		}
		return 0
	case "enPassant": // enPassant
		if move&0x400000 != 0 {
			return 1
		}
		return 0
	case "castling": // castling
		if move&0x800000 != 0 {
			return 1
		}
		return 0
	}
	return 33554431
}

// append move to moveList
func addMove(move uint64) {
	moveList = append(moveList, move)
}

// print move source, target, and promoted piece
func printMove(move uint64) {
	fmt.Printf("%s%s%s\n", algebraic[getMoveAttr(move, "source")],
		algebraic[getMoveAttr(move, "target")], stringPieces[getMoveAttr(move, "promoted")],
	)
}

// loop through all moves in move list and print
func printMoveList() {
	// formatting
	fmt.Printf("\n    move    piece   capture   double    enpass    castling\n\n")

	// loop through movesList
	for _, move := range moveList {
		fmt.Printf("    %s%s%s   %s       %d         %d         %d         %d\n",
			algebraic[getMoveAttr(move, "source")],
			algebraic[getMoveAttr(move, "target")],
			promotedPieces[getMoveAttr(move, "promoted")],
			stringPieces[getMoveAttr(move, "piece")],
			getMoveAttr(move, "capture"),
			getMoveAttr(move, "double"),
			getMoveAttr(move, "enPassant"),
			getMoveAttr(move, "castling"),
		)
	}
	// print total number of moves
	fmt.Println()
	fmt.Printf("    Total number of moves: %d\n\n", len(moveList))
}
