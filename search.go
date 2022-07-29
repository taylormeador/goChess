package main

import "fmt"

// half move counter
var ply int

// best move
var bestMove uint64

// init
var bestMoveSoFar uint64

// quiescence search
func quiescence(alpha int, beta int) int {
	// evaluate current position
	evaluation := evaluate()

	// fail-hard beta cutoff
	if evaluation >= beta {
		// node (move) fails high
		return beta
	}

	// found a better move
	if evaluation > alpha {
		// PV node (move)
		alpha = evaluation
	}

	// generate moves
	generateAllMoves()

	// loop over moves within a move list
	for count := 0; count < len(moveList); count++ {

		// preserve board state
		copyBoard()

		// increment half move counter
		ply++

		// check that move is legal
		if makeMove(moveList[count]) == 0 {
			// decrement half move counter if move is illegal
			ply--
			continue
		}

		// score current move
		score := -quiescence(-beta, -alpha)

		// decrement ply
		ply--

		// take move back
		restoreBoard()

		// fail-hard beta cutoff
		if score >= beta {
			// node (move) fails high
			return beta
		}

		// found a better move
		if score > alpha {
			// PV node (move)
			alpha = score
		}

	}
	// node (move) fails low
	return alpha
}

// negamax with alpha beta pruning
func negaMax(alpha int, beta int, depth int) int {
	// recursion escape condition
	if depth == 0 {
		return quiescence(alpha, beta)
	}

	// increment node counter
	nodes++

	// init check vars
	currentSide := white
	currentKingBitboard := bitboards[K]
	if side == black {
		currentSide = black
		currentKingBitboard = bitboards[k]
	}

	// is king in check
	inCheck := isSquareAttacked(getLeastSignificantBitIndex(currentKingBitboard), currentSide)

	// init vars
	legalMoves := 0
	bestMoveSoFar = 0
	oldAlpha := alpha
	generateAllMoves()

	// loop over moves within a move list
	for count := 0; count < len(moveList); count++ {

		// preserve board state
		copyBoard()

		// increment half move counter
		ply++

		// check that move is legal
		if makeMove(moveList[count]) == 0 {
			// decrement half move counter if move is illegal
			ply--
			continue
		}

		// increment legalMoves counter
		legalMoves++

		// score current move
		score := -negaMax(-beta, -alpha, depth-1)

		// decrement ply
		ply--

		// take move back
		restoreBoard()

		// fail-hard beta cutoff
		if score >= beta {
			// node (move) fails high
			return beta
		}

		// found a better move
		if score > alpha {
			// PV node (move)
			alpha = score

			// root node
			if ply == 0 {
				bestMoveSoFar = moveList[count]
			}
		}

	}

	// we don't have any legal moves to make in the current postion
	if legalMoves == 0 {
		// king is in check
		if inCheck != 0 {
			// return mating score (assuming closest distance to mating position)
			fmt.Println("checkmate found")
			return -49000 + ply
		} else { // king not in check
			// return stalemate score
			fmt.Println("stalemate found")
			return 0
		}
	}

	// update alpha
	if oldAlpha != alpha {
		bestMove = bestMoveSoFar
	}
	// node (move) fails low
	return alpha
}

// search position for the best move
func searchPosition(depth int) {
	negaMax(-50000, 50000, depth)

	// best move placeholder
	fmt.Printf("bestmove ")
	if bestMove != 0 {
		printMove(bestMove)
	}
	fmt.Printf("\n")
}
