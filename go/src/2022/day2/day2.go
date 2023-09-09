package main

import (
	"fmt"
	"github.com/antonio-mikulic/advent-of-code/pkg/reader"
)

func main() {

	// Rock = A && X
	// Paper = B && Y
	// Scissors = C && Z
	winning := map[uint8]uint8{
		'A': 'Y',
		'B': 'Z',
		'C': 'X',
	}

	draw := map[uint8]uint8{
		'A': 'X',
		'B': 'Y',
		'C': 'Z',
	}

	score := 0
	rockScore := 1
	paperScore := 2
	scissorsScore := 3

	for _, word := range reader.ReadRows(false) {
		they := word[0]
		me := word[2]
		fmt.Printf("%c - %c, starting at %d. ", they, me, score)

		if winning[they] == me {
			score += 6
			fmt.Printf("we win +6")
		} else if draw[they] == me {
			score += 3
			fmt.Printf("we draw +3")
		} else {
			fmt.Printf("we lose +0")
		}

		switch me {
		case 'X':
			score += rockScore
			fmt.Printf(", and %d for rock", rockScore)
		case 'Y':
			score += paperScore
			fmt.Printf(", and %d for paper", paperScore)
		case 'Z':
			score += scissorsScore
			fmt.Printf(", and %d for scissors", scissorsScore)
		}

		fmt.Printf(", ending with %d \n", score)
	}

	fmt.Print(score)

}
