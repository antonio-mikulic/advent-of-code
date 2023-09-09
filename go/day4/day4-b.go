package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
			getIntB(firstElf[0]), getIntB(firstElf[1]),
		}
		second := []int{
			getIntB(secondElf[0]), getIntB(secondElf[1]),
		}

		if isOverlap(first, second) {
			res += 1
		}
	}

	fmt.Println(res)
}

func isOverlap(first []int, second []int) bool {
	if first[0] < second[0] {
		return first[1] >= second[0]
	} else {
		return second[1] >= first[0]
	}
}

func getIntB(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return val
}
