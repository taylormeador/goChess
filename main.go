package main

import (
	"bufio"
	"fmt"
	"os"
)

// init
func initAll() {
	initLeapersAttacks()
	initSlidersAttacks(bishop)
	initSlidersAttacks(rook)
}

// main
func main() {
	initAll()
	parseFEN(startPosition)
	whiteIsHuman := true
	blackIsHuman := false
	depth := 6
	for {
		printBoard()
		if (side == white && whiteIsHuman) || (side == black && blackIsHuman) {
			fmt.Print("\nenter move: ")
			input := bufio.NewScanner(os.Stdin)
			input.Scan()
			if makeMove(parseMove(input.Text())) == 0 {
				fmt.Printf("invalid move")
			}
		} else { // cpu move
			if makeMove(searchPosition(depth)) != 0 {
				fmt.Printf("\ncpu move: ")
				printMove(bestMove)
				fmt.Println()
				continue
			} else {
				fmt.Printf("error")
				printBoard()
				break
			}
		}
	}

}
