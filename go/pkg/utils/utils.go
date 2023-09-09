package utils

import "strconv"

func IsOverlap(first []int, second []int) bool {
	if first[0] < second[0] {
		return first[1] >= second[0]
	} else {
		return second[1] >= first[0]
	}
}

func GetInt(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return val
}

func IsFirstPartOfSecond(first []int, second []int) bool {
	return first[0] >= second[0] && first[1] <= second[1]
}

func GetCharFromShortInt(item uint8) uint8 {
	if item > 96 {
		return item - 96
	}

	return item - 64 + 26
}
