package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("day_three/input.txt")
	if err != nil {
		panic(err)
	}

	// partOne(string(input))
	partTwo(string(input))
}

func partOne(input string) {
	split := strings.Split(input, "mul(")
	var res int

	for _, part := range split {
		commaIdx := strings.Index(part, ",")
		if commaIdx == -1 {
			continue
		}

		braceIdx := strings.Index(part, ")")
		if braceIdx == -1 {
			continue
		}

		nrOne, err := strconv.Atoi(part[:commaIdx])
		if err != nil {
			continue
		}

		nrTwo, errTwo := strconv.Atoi(part[commaIdx+1 : braceIdx])
		if errTwo != nil {
			continue
		}

		res += mul(nrOne, nrTwo)
	}

	fmt.Printf("The result is: %d\n", res)
}

func partTwo(input string) {
	split := strings.Split(input, "mul(")
	canMul := true
	var res int

	for _, part := range split {

		commaIdx := strings.Index(part, ",")
		if commaIdx == -1 {
			continue
		}

		braceIdx := strings.Index(part, ")")
		if braceIdx == -1 {
			continue
		}

		nrOne, err := strconv.Atoi(part[:commaIdx])
		if err != nil {
			continue
		}

		nrTwo, errTwo := strconv.Atoi(part[commaIdx+1 : braceIdx])
		if errTwo != nil {
			continue
		}

		if canMul {
			res += mul(nrOne, nrTwo)
		}

		if strings.Contains(part, "do()") {
			cmds := strings.Split(part, "do()")
			for _, cmd := range cmds {
				if strings.Contains(cmd, "don't()") {
					canMul = false
				} else {
					canMul = true
				}
			}
		} else if strings.Contains(part, "don't()") {
			canMul = false
		}

	}

	fmt.Printf("The result is: %d\n", res)
}

func mul(a int, b int) int {
	return a * b
}
