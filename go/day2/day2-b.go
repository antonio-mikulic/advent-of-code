package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(file)

	score := 0

	for fileScanner.Scan() {
		game := fileScanner.Text()
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
