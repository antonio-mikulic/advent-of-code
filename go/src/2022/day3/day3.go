package main

import (
	"fmt"
	"github.com/antonio-mikulic/advent-of-code/pkg/reader"
	"github.com/antonio-mikulic/advent-of-code/pkg/utils"
)

func main() {
	var res uint16 = 0

	for _, sack := range reader.ReadRows(false) {
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
				res += uint16(utils.GetCharFromShortInt(secondHalf[i]))
				break
			}
		}
	}

	fmt.Println(res)
}
