package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("day_one/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSuffix(string(input), "\n"), "\n")

	// partOne(lines)
	partTwo(lines)
}

func partOne(lines []string) {
	var lists [2][]int

	for _, line := range lines {
		pair := strings.Split(line, "   ")

		first, err := strconv.Atoi(pair[0])
		if err != nil {
			panic(err)
		}

		second, secerr := strconv.Atoi(pair[1])
		if secerr != nil {
			panic(err)
		}

		lists[0] = append(lists[0], first)
		lists[1] = append(lists[1], second)
	}

	sort.Ints(lists[0])
	sort.Ints(lists[1])

	var distance int
	for i := 0; i < len(lists[0]); i++ {
		if lists[1][i] > lists[0][i] {
			distance += lists[1][i] - lists[0][i]
		} else {
			distance += lists[0][i] - lists[1][i]
		}
	}

	fmt.Printf("Total distance: %d\n", distance)
}

func partTwo(lines []string) {
	var lists [2][]int

	for _, line := range lines {
		pair := strings.Split(line, "   ")

		first, err := strconv.Atoi(pair[0])
		if err != nil {
			panic(err)
		}

		second, secerr := strconv.Atoi(pair[1])
		if secerr != nil {
			panic(err)
		}

		lists[0] = append(lists[0], first)
		lists[1] = append(lists[1], second)
	}

	var simScore int
	for _, firstNum := range lists[0] {
		var count int
		for _, secondNum := range lists[1] {
			if firstNum == secondNum {
				count += 1
			}
		}

		simScore += firstNum * count
	}

	fmt.Printf("Total similarity score: %d\n", simScore)
}
