package main

// individual pieces
var bitboards [12]uint64

// by colors
var occupancyBitboards [3]uint64

// side to move
var side = -1

// en passant square
var enPassantSquare = noSquare
