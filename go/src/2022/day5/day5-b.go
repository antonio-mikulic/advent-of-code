package main

import (
	"fmt"
	"github.com/antonio-mikulic/advent-of-code/pkg/data"
	"github.com/antonio-mikulic/advent-of-code/pkg/reader"
	"github.com/antonio-mikulic/advent-of-code/pkg/utils"
	"strings"
)

func main() {
	stacks := make([]data.Stack[uint8], 0)
	containerSpacing := 4

	for _, line := range reader.ReadRows(false) {
		fmt.Println(line)

		// When setup is done rows
		if line == "" {
			fmt.Println("Finished setup with stacks: ", len(stacks))

			for _, stack := range stacks {
				stack.Reverse()
				fmt.Printf("Reversed stack, top is %c \n", stack.Peek())
			}
			continue
		}

		if line[0] != 'm' && !strings.Contains(line, "[") {
			continue
		}

		// Setup
		if line[0] != 'm' && strings.Contains(line, "[") {
			for i := 1; i < len(line); i += containerSpacing {
				if line[i] == 32 {
					continue
				}

				currStack := i / containerSpacing

				if currStack >= len(stacks) {
					for j := len(stacks); j < currStack+1; j++ {
						fmt.Printf("Adding new stack to support stack size of %d \n", currStack)
						stacks = append(stacks, data.NewStack[uint8]())
					}
				}

				fmt.Printf("Adding '%c' to stack %d \n", line[i], currStack)
				stacks[currStack].Push(line[i])
			}
			continue
		}

		commands := strings.Split(line, " ")
		fmt.Println()

		amount := utils.GetInt(commands[1])
		from := utils.GetInt(commands[3]) - 1
		to := utils.GetInt(commands[5]) - 1

		fmt.Printf("Moving %d from %d to %d\n", amount, from, to)

		popped := make([]uint8, 0)

		for i := 0; i < amount; i++ {
			popped = append(popped, stacks[from].Pop())
		}

		for i := len(popped); i >= 1; i-- {
			stacks[to].Push(popped[i-1])
		}
	}

	for _, stack := range stacks {
		fmt.Printf("%c", stack.Peek())
	}
}
