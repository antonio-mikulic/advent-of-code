package main

import (
	"bufio"
	"fmt"
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
	groups := make([]string, 0)

	for scanner.Scan() {
		sack := scanner.Text()
		groups = append(groups, sack)

		if len(groups) != 3 {
			fmt.Printf("Next row %d \n", len(groups))
			continue
		}

		fmt.Printf("Calc for row %d \n", len(groups))

		// Transfer into uint8 arrays
		g1, g2, g3 := strings.Split(groups[0], ""), strings.Split(groups[1], ""), strings.Split(groups[2], "")
		inTwo := intersects(g1, g2)

		inThree := intersects(inTwo, g3)[0]

		res += uint16(GetValB(inThree[0]))

		groups = nil
	}

	fmt.Println(res)
}

func intersects[T comparable](a []T, b []T) []T {
	set := make([]T, 0)
	hash := make(map[T]struct{})

	for _, v := range a {
		hash[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := hash[v]; ok {
			set = append(set, v)
		}
	}

	return set
}

func GetValB(item uint8) uint8 {
	if item > 96 {
		return item - 96
	}

	return item - 64 + 26
}
