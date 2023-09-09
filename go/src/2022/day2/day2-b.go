package main

import (
	"fmt"
	"github.com/antonio-mikulic/advent-of-code/pkg/reader"
)

func main() {
	score := 0

	for _, game := range reader.ReadRows(false) {
		me := game[2]

		switch me {
		case 'Z':
			score += 6
		case 'Y':
			score += 3
		case 'X':
			score += 0
		}

		// Rock = A
		// Paper = B
		// Scissors = C
		// We loose = X
		// Draw = Y
		// We win = Z

		switch game {
		case "A Y", "B X", "C Z":
			score += 1
		case "A Z", "B Y", "C X":
			score += 2
		case "A X", "B Z", "C Y":
			score += 3
		}
	}

	fmt.Print(score)

}
