package main

import (
	"fmt"
	"github.com/antonio-mikulic/advent-of-code/pkg/reader"
	"github.com/antonio-mikulic/advent-of-code/pkg/utils"
	"strings"
)

func main() {
	var res uint16 = 0

	for _, text := range reader.ReadRows(false) {
		group := strings.Split(text, ",")
		firstElf := strings.Split(group[0], "-")
		secondElf := strings.Split(group[1], "-")

		first := []int{
			utils.GetInt(firstElf[0]), utils.GetInt(firstElf[1]),
		}
		second := []int{
			utils.GetInt(secondElf[0]), utils.GetInt(secondElf[1]),
		}

		if utils.IsOverlap(first, second) {
			res += 1
		}
	}

	fmt.Println(res)
}
