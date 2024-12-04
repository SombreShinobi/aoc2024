package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("day_four/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSuffix(string(input), "\n"), "\n")

	// partOne(lines)
	partTwo(lines)
}

func partOne(lines []string) {
	var count int

	for i, line := range lines {
		idxs := allIdx(line, 'X')
		if len(idxs) == 0 {
			continue
		}

		for _, idx := range idxs {
			count += east(line, idx)
			count += west(line, idx)
			count += north(lines, i, idx)
			count += south(lines, i, idx)
			count += nw(lines, i, idx)
			count += ne(lines, i, idx)
			count += sw(lines, i, idx)
			count += se(lines, i, idx)
		}

	}

	fmt.Printf("Number of XMAS': %d\n", count)
}

func partTwo(lines []string) {
	var count int

	for i, line := range lines {
		idxs := allIdx(line, 'A')
		if len(idxs) == 0 {
			continue
		}

		for _, idx := range idxs {
			count += mas(lines, i, idx)
		}
	}

	fmt.Printf("Number of X-MAS': %d\n", count)
}

func east(line string, idx int) int {
	if len(line)-idx-1 < 3 {
		return 0
	}

	if line[idx+1] == 'M' {
		if line[idx+2] == 'A' {
			if line[idx+3] == 'S' {
				return 1
			}
		}
	}

	return 0
}

func west(line string, idx int) int {
	if idx < 3 {
		return 0
	}

	if line[idx-1] == 'M' {
		if line[idx-2] == 'A' {
			if line[idx-3] == 'S' {
				return 1
			}
		}
	}

	return 0
}

func north(lines []string, lineIdx int, charIdx int) int {
	if lineIdx < 3 {
		return 0
	}

	if lines[lineIdx-1][charIdx] == 'M' {
		if lines[lineIdx-2][charIdx] == 'A' {
			if lines[lineIdx-3][charIdx] == 'S' {
				return 1
			}
		}
	}

	return 0
}

func south(lines []string, lineIdx int, charIdx int) int {
	if len(lines)-lineIdx-1 < 3 {
		return 0
	}

	if lines[lineIdx+1][charIdx] == 'M' {
		if lines[lineIdx+2][charIdx] == 'A' {
			if lines[lineIdx+3][charIdx] == 'S' {
				return 1
			}
		}
	}

	return 0
}

func nw(lines []string, lineIdx int, charIdx int) int {
	if lineIdx < 3 || charIdx < 3 {
		return 0
	}

	if lines[lineIdx-1][charIdx-1] == 'M' {
		if lines[lineIdx-2][charIdx-2] == 'A' {
			if lines[lineIdx-3][charIdx-3] == 'S' {
				return 1
			}
		}
	}

	return 0
}

func ne(lines []string, lineIdx int, charIdx int) int {
	if lineIdx < 3 || len(lines[lineIdx])-charIdx-1 < 3 {
		return 0
	}

	if lines[lineIdx-1][charIdx+1] == 'M' {
		if lines[lineIdx-2][charIdx+2] == 'A' {
			if lines[lineIdx-3][charIdx+3] == 'S' {
				return 1
			}
		}
	}

	return 0
}

func sw(lines []string, lineIdx int, charIdx int) int {
	if len(lines)-lineIdx-1 < 3 || charIdx < 3 {
		return 0
	}

	if lines[lineIdx+1][charIdx-1] == 'M' {
		if lines[lineIdx+2][charIdx-2] == 'A' {
			if lines[lineIdx+3][charIdx-3] == 'S' {
				return 1
			}
		}
	}

	return 0
}

func se(lines []string, lineIdx int, charIdx int) int {
	if len(lines)-lineIdx-1 < 3 || len(lines[lineIdx])-charIdx-1 < 3 {
		return 0
	}

	if lines[lineIdx+1][charIdx+1] == 'M' {
		if lines[lineIdx+2][charIdx+2] == 'A' {
			if lines[lineIdx+3][charIdx+3] == 'S' {
				return 1
			}
		}
	}

	return 0
}

func allIdx(line string, ch rune) []int {
	var idxs []int
	for i, c := range line {
		if c == ch {
			idxs = append(idxs, i)
		}
	}
	return idxs
}

func mas(lines []string, lineIdx int, charIdx int) int {
	if lineIdx == 0 || len(lines)-lineIdx-1 == 0 || charIdx == 0 || len(lines[lineIdx])-charIdx-1 == 0 {
		return 0
	}

	first := firstDiagonal(lines, lineIdx, charIdx)
	second := secondDiagonal(lines, lineIdx, charIdx)

	if first && second {
		return 1
	}

	return 0
}

func firstDiagonal(lines []string, lineIdx int, charIdx int) bool {
	if lines[lineIdx-1][charIdx-1] == 'M' {
		if lines[lineIdx+1][charIdx+1] == 'S' {
			return true
		}
	} else if lines[lineIdx-1][charIdx-1] == 'S' {
		if lines[lineIdx+1][charIdx+1] == 'M' {
			return true
		}
	}

	return false
}

func secondDiagonal(lines []string, lineIdx int, charIdx int) bool {
	if lines[lineIdx-1][charIdx+1] == 'M' {
		if lines[lineIdx+1][charIdx-1] == 'S' {
			return true
		}
	} else if lines[lineIdx-1][charIdx+1] == 'S' {
		if lines[lineIdx+1][charIdx-1] == 'M' {
			return true
		}
	}

	return false
}
