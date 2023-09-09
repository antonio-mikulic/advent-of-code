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
			getInt(firstElf[0]), getInt(firstElf[1]),
		}
		second := []int{
			getInt(secondElf[0]), getInt(secondElf[1]),
		}

		if isFirstPartOfSecond(first, second) || isFirstPartOfSecond(second, first) {
			res += 1
		}
	}

	fmt.Println(res)
}

func isFirstPartOfSecond(first []int, second []int) bool {
	return first[0] >= second[0] && first[1] <= second[1]
}

func getInt(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return val
}
