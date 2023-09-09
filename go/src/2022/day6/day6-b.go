package main

import (
	"fmt"
	"github.com/antonio-mikulic/advent-of-code/pkg/reader"
	"slices"
)

func main() {
	for _, line := range reader.ReadRows(false) {
		fmt.Println("Reading line: ", line)
		lastChars := make([]int32, 0)

		for index, char := range line {
			found := slices.Index(lastChars, char)
			if found != -1 {
				lastChars = lastChars[found+1:]
			}

			lastChars = append(lastChars, char)

			if len(lastChars) == 14 {
				fmt.Printf("Code in row starts at %d \n", index+1)
				break
			}
		}
	}
}
