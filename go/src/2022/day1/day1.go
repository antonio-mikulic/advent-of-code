package main

import (
	"fmt"
	"github.com/antonio-mikulic/advent-of-code/pkg/reader"
	"slices"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	elfArray := make([]int, 0)
	currentElfSum := 0

	for _, row := range reader.ReadRows(false) {
		val, err := strconv.Atoi(row)
		if err != nil {
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
	fmt.Println(time.Since(start))
}
