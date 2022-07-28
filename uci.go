package main

// parse user/gui move string input e.g. "f7f8q"
func parseMove(moveString string) uint64 {
	generateAllMoves()
	sourceSquare := moveString[:2]
	targetSquare := moveString[2:4]

	// loop through all the possible moves and find the given move
	for _, move := range moveList {
		if sourceSquare == algebraic[getMoveAttr(move, "source")] &&
			targetSquare == algebraic[getMoveAttr(move, "target")] {
			promotedPiece := getMoveAttr(move, "promoted")

			// check if a piece was promoted
			if promotedPiece != 0 {

				// queen promotion
				if (promotedPiece == Q || promotedPiece == q) && (moveString[4] == 'q') {
					return move
				}

				// rook promotion
				if (promotedPiece == R || promotedPiece == r) && (moveString[4] == 'r') {
					return move
				}

				// bishop promotion
				if (promotedPiece == B || promotedPiece == b) && (moveString[4] == 'b') {
					return move
				}

				// knight promotion
				if (promotedPiece == N || promotedPiece == n) && (moveString[4] == 'n') {
					return move
				}
				continue
			}
			return move
		}
	}
	return 0
}
