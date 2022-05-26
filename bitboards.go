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
