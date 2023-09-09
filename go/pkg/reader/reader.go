package reader

import (
	"os"
	"strings"
)

func ReadRows(useTest bool) []string {
	fileName := "input.txt"
	if useTest {
		fileName = "input-s.txt"
	}

	dat, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(dat), "\r\n")
}
