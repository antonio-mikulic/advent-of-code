package main

import (
	"bufio"
	"fmt"
	utils2 "github.com/antonio-mikulic/advent-of-code/pkg/utils"
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
			continue
		}

		// Transfer into uint8 arrays
		g1, g2, g3 := strings.Split(groups[0], ""), strings.Split(groups[1], ""), strings.Split(groups[2], "")
		inThree := utils2.Intersects(utils2.Intersects(g1, g2), g3)[0]

		res += uint16(utils2.GetCharFromShortInt(inThree[0]))
		groups = nil
	}

	fmt.Println(res)
}
