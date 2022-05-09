package main

// check if a bit is on or off
func getBit(b uint64, square uint64) uint64 {
	return b & uint64(1) << square
}

// turn on a bit
func setBit(b uint64, square uint64) uint64 {
	return b | uint64(1)<<square
}

// turn off a bit
func popBit(b uint64, square uint64) uint64 {
	return b & ^(uint64(1) << square)
}

// count the number of bits on a bitboard
func countBits(b uint64) uint64 {
	count := uint64(0)
	for {
		b &= b - 1
		count += 1
		if b == 0 {
			break
		}
	}
	return count
}

// get the index of the least significant bit that is on
func getLeastSignificantBitIndex(b uint64) uint64 {
	// check for a non empty bitboard
	if b >= 0 {
		leastSignificantBit := b & -b
		leadingOnes := leastSignificantBit - 1
		return countBits(leadingOnes)
	} else {
		// return out of range index if the board is empty
		return uint64(64)
	}

}
