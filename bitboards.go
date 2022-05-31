package main

// individual pieces
var bitboards [12]uint64

// by colors
var occupancies [3]uint64

// side to move
var side = -1

// en passant square
var enPassantSquare = noSquare

// castling rights
var castle = -1

// custom type for collecting board info
type gameState struct {
	bitboards       [12]uint64
	occupancies     [3]uint64
	enPassantSquare uint64
	side            int
	castle          int
}

// returns a copy of the gamestate
func copyBoard() gameState {
	var state gameState

	state.bitboards = bitboards
	state.occupancies = occupancies
	state.enPassantSquare = enPassantSquare
	state.side = side
	state.castle = castle

	return state
}

// sets the relevant global vars to reflect the given gamestate
func restoreBoard(state gameState) {
	bitboards = state.bitboards
	occupancies = state.occupancies
	enPassantSquare = state.enPassantSquare
	side = state.side
	castle = state.castle
}
