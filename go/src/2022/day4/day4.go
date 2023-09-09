package main

import (
	"bufio"
	"fmt"
	"github.com/antonio-mikulic/advent-of-code/pkg/utils"
	"os"
	"strings"
)

func main() {
	dat, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var res uint16 = 0

	scanner := bufio.NewScanner(dat)

	for scanner.Scan() {
		group := strings.Split(scanner.Text(), ",")
		firstElf := strings.Split(group[0], "-")
		secondElf := strings.Split(group[1], "-")

		first := []int{
			utils.GetInt(firstElf[0]), utils.GetInt(firstElf[1]),
		}
		second := []int{
			utils.GetInt(secondElf[0]), utils.GetInt(secondElf[1]),
		}

		if utils.IsFirstPartOfSecond(first, second) || utils.IsFirstPartOfSecond(second, first) {
			res += 1
		}
	}

	fmt.Println(res)
}
