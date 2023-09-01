package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	dat, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(dat)
	fileScanner.Split(bufio.ScanLines)

	elfArray := make([]int, 0)
	currentElfSum := 0

	for fileScanner.Scan() {
		val, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			fmt.Println("Adding %i", currentElfSum)
			elfArray = append(elfArray, currentElfSum)
			currentElfSum = 0
		}

		currentElfSum += val
	}

	slices.Sort(elfArray)

	sum := 0
	for i := len(elfArray) - 1; i >= len(elfArray)-3; i-- {
		sum += elfArray[i]
	}

	fmt.Println(sum)
}
