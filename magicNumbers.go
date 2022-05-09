package main

import "fmt"

// generate candidate magic number for testing
func generateCandidateMagicNumber() uint64 {
	return getRandom64BitNumber() & getRandom64BitNumber() & getRandom64BitNumber()
}

// look at candidate magic numbers and find an appropriate one
func findMagicNumber(square uint64, relevantBits uint64, piece int) uint64 {
	var attackMask uint64
	var occupancies [4096]uint64
	var attacks [4096]uint64
	var usedAttacks [4096]uint64

	// get attack mask
	if piece == 1 {
		// bishop
		attackMask = maskBishopAttacks(square)
	} else {
		// rook
		attackMask = maskRookAttacks(square)
	}

	occupancyIndices := uint64(1) << relevantBits

	// loop over occupancy indices
	for i := uint64(0); i < occupancyIndices; i++ {
		occupancies[i] = setOccupancy(i, relevantBits, attackMask)

		// init attacks
		if piece == 0 {
			// bishop
			attacks[i] = bishopAttacksOnTheFly(square, occupancies[i])
		} else {
			// rook
			attacks[i] = rookAttacksOnTheFly(square, occupancies[i])
		}
	}

	// test magic numbers
	for randomCount := uint64(0); randomCount < 100000000; randomCount++ {

		// get a magic number
		magicNumber := generateCandidateMagicNumber()

		// skip magic numbers with too few on bits
		if countBits((attackMask*magicNumber)&0xFF00000000000000) < 6 {
			continue
		}
		// test magic index
		validMagicNumber := true
		for i, fail := uint64(0), false; !fail && i < occupancyIndices; i++ {
			magicIndex := (occupancies[i] * magicNumber) >> (64 - relevantBits)

			// if the attack is empty at that index
			if usedAttacks[magicIndex] == 0 {
				// add the attack
				usedAttacks[magicIndex] = attacks[i]
				// otherwise check if the existing attack matches the correct attack
			} else if usedAttacks[magicIndex] != attacks[i] {
				// magic index does not work
				fail = true
				validMagicNumber = false
			}
		}
		if validMagicNumber {
			return magicNumber
		}
	}
	fmt.Println("  Failed to find magic number")
	return uint64(0)
}

// init magic numbers
func initMagicNumbers() {
	// loop over all the squares on the board
	for square := uint64(0); square < 64; square++ {
		// print rook magic numbers
		fmt.Printf(" 0x%x,\n", findMagicNumber(square, rookRelevantBits[square], rook))
	}

	fmt.Println("****************************************")

	// loop over all the squares on the board
	for square := uint64(0); square < 64; square++ {
		// print print magic numbers
		fmt.Printf(" 0x%x,\n", findMagicNumber(square, bishopRelevantBits[square], bishop))
	}
}
