package main

// white = 1 or true, black = 0 or false
const (
	white = 1
	black = 0
)

// enumerate board squares
const (
	a8 = uint64(iota)
	b8
	c8
	d8
	e8
	f8
	g8
	h8
	a7
	b7
	c7
	d7
	e7
	f7
	g7
	h7
	a6
	b6
	c6
	d6
	e6
	f6
	g6
	h6
	a5
	b5
	c5
	d5
	e5
	f5
	g5
	h5
	a4
	b4
	c4
	d4
	e4
	f4
	g4
	h4
	a3
	b3
	c3
	d3
	e3
	f3
	g3
	h3
	a2
	b2
	c2
	d2
	e2
	f2
	g2
	h2
	a1
	b1
	c1
	d1
	e1
	f1
	g1
	h1
)

// bitboard masks
const (
	//  e.g. all 0's in the "a" file, all 1's elsewhere
	//
	//  8    0  1  1  1  1  1  1  1
	//  7    0  1  1  1  1  1  1  1
	//  6    0  1  1  1  1  1  1  1
	//  5    0  1  1  1  1  1  1  1
	//  4    0  1  1  1  1  1  1  1
	//  3    0  1  1  1  1  1  1  1
	//  2    0  1  1  1  1  1  1  1
	//  1    0  1  1  1  1  1  1  1
	//
	//       a  b  c  d  e  f  g  h
	//       Bitboard: 18374403900871474942

	notAFile  = uint64(18374403900871474942)
	notHFile  = uint64(9187201950435737471)
	notABFile = uint64(18229723555195321596)
	notGHFile = uint64(4557430888798830399)
)

// easy way to lookup name of square from index in bitboard
var algebraic = [64]string{
	"a8", "b8", "c8", "d8", "e8", "f8", "g8", "h8",
	"a7", "b7", "c7", "d7", "e7", "f7", "g7", "h7",
	"a6", "b6", "c6", "d6", "e6", "f6", "g6", "h6",
	"a5", "b5", "c5", "d5", "e5", "f5", "g5", "h5",
	"a4", "b4", "c4", "d4", "e4", "f4", "g4", "h4",
	"a3", "b3", "c3", "d3", "e3", "f3", "g3", "h3",
	"a2", "b2", "c2", "d2", "e2", "f2", "g2", "h2",
	"a1", "b1", "c1", "d1", "e1", "f1", "g1", "h1",
}

// relevant occupancy bit counts for a bishop at every square on board
var bishopRelevantBits = [64]uint64{
	6, 5, 5, 5, 5, 5, 5, 6,
	5, 5, 5, 5, 5, 5, 5, 5,
	5, 5, 7, 7, 7, 7, 5, 5,
	5, 5, 7, 9, 9, 7, 5, 5,
	5, 5, 7, 9, 9, 7, 5, 5,
	5, 5, 7, 7, 7, 7, 5, 5,
	5, 5, 5, 5, 5, 5, 5, 5,
	6, 5, 5, 5, 5, 5, 5, 6,
}

// relevant occupancy bit counts for a rook at every square on board
var rookRelevantBits = [64]uint64{
	12, 11, 11, 11, 11, 11, 11, 12,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	12, 11, 11, 11, 11, 11, 11, 12,
}
