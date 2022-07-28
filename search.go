package main

import "fmt"

// half move counter
var ply int

// best move
var bestMove uint64

// init
var bestMoveSoFar uint64

func negaMax(alpha int, beta int, depth int) int {
	// recursion escape condition
	if depth == 0 {
		return evaluate()
	}

	// increment node counter
	nodes++

	// old value of alpha
	oldAlpha := alpha

	// generate moves
	generateAllMoves()

	// loop over moves within a movelist
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

			// if root move
			if ply == 0 {
				bestMoveSoFar = moveList[count]
			}
		}
	}

	// found better move
	if oldAlpha != alpha {
		bestMove = bestMoveSoFar
		fmt.Printf("new best move so far: ")
		printMove(bestMove)
		fmt.Println()
	}

	// node (move) fails low
	return alpha
}

// search position for the best move
func searchPosition(depth int) {
	negaMax(-50000, 50000, depth)

	// best move placeholder
	fmt.Printf("bestmove ")
	printMove(bestMove)
	fmt.Printf("\n")
}
