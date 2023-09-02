package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	dat, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var res uint16 = 0

	scanner := bufio.NewScanner(dat)

	for scanner.Scan() {
		sack := scanner.Text()
		middlePoint := len(sack) / 2
		used := make(map[uint8]bool)

		firstHalf := sack[:middlePoint]
		for i := 0; i < len(firstHalf); i++ {
			used[firstHalf[i]] = true
		}

		secondHalf := sack[middlePoint:]
		for i := 0; i < len(secondHalf); i++ {
			ch := secondHalf[i]
			if used[ch] {
				res += uint16(GetVal(secondHalf[i]))
				break
			}
		}
	}

	fmt.Println(res)
}

func GetVal(item uint8) uint8 {
	if item > 96 {
		return item - 96
	}

	return item - 64 + 26
}
