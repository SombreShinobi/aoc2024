package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("day_two/input.txt")
	if err != nil {
		panic(err)
	}

	reports := strings.Split(strings.TrimSuffix(string(input), "\n"), "\n")

	// partOne(reports)
	partTwo(reports)
}

func partOne(reports []string) {
	var safeReps int
	for _, rep := range reports {
		levls := strings.Split(rep, " ")

		if isRepSafe(levls) {
			safeReps += 1
		}
	}

	fmt.Printf("Number of safe reports: %d\n", safeReps)
}

func isRepSafe(levls []string) bool {
	for i, lvl := range levls {
		if i == 0 {
			continue
		}

		lvlNr, err := strconv.Atoi(lvl)
		if err != nil {
			panic(err)
		}

		prevLvl, secerr := strconv.Atoi(levls[i-1])
		if secerr != nil {
			panic(err)
		}

		if i+1 <= len(levls)-1 {
			nextLvl, thirderr := strconv.Atoi(levls[i+1])
			if thirderr != nil {
				panic(err)
			}

			if prevLvl < lvlNr && nextLvl < lvlNr {
				return false
			}

			if prevLvl > lvlNr && nextLvl > lvlNr {
				return false
			}
		}

		dif := math.Abs(float64(lvlNr - prevLvl))
		if dif < 1 || dif > 3 {
			return false
		}

	}

	return true
}

func partTwo(reports []string) {
	var safeReps int
outer:
	for _, rep := range reports {
		levls := strings.Split(rep, " ")
		fmt.Println(rep)

		if isRepSafe(levls) {
			safeReps += 1
		} else {
			for i := 0; i < len(levls); i++ {
				newLvls := removeIdx(levls, i)
				fmt.Println(newLvls)
				if isRepSafe(newLvls) {
					fmt.Println("Safe ", newLvls)
					safeReps += 1
					continue outer
				}
			}
		}
	}

	fmt.Printf("Number of safe reports: %d\n", safeReps)
}

func removeIdx(s []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
