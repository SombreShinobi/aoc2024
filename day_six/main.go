package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Dir string

type Guard struct {
	row, col int
	dir      Dir
	turns    []string
	stuck    bool
}

const (
	up    Dir = "up"
	down  Dir = "down"
	left  Dir = "left"
	right Dir = "right"
)

func (self *Guard) changeDir() {
	turn := fmt.Sprintf("%d,%d,%s", self.col, self.row, self.dir)
	idx := slices.Index(self.turns, turn)
	if idx == -1 {
		self.turns = append(self.turns, turn)
	} else {
		self.stuck = true
	}

	switch self.dir {
	case "up":
		self.dir = right
	case "right":
		self.dir = down
	case "down":
		self.dir = left
	case "left":
		self.dir = up
	}
}

func (self *Guard) moveVert(newRow int, rows *[]string, steps *int) bool {
	if len(*rows) == newRow || newRow == -1 {
		return false
	}

	char := (*rows)[newRow][self.col]
	if char == '#' {
		self.changeDir()
	} else if char == '.' {
		self.row = newRow
		*steps += 1
		(*rows)[newRow] = markSpot((*rows)[newRow], self.col, 'X')
	} else {
		self.row = newRow
	}

	return true
}

func (self *Guard) moveHoriz(newCol int, rows *[]string, steps *int) bool {
	if len((*rows)[self.row]) == newCol || newCol == -1 {
		return false
	}

	char := (*rows)[self.row][newCol]
	if char == '#' {
		self.changeDir()
	} else if char == '.' {
		self.col = newCol
		*steps += 1
		(*rows)[self.row] = markSpot((*rows)[self.row], newCol, 'X')
	} else {
		self.col = newCol
	}

	return true
}

func (self *Guard) move(rows *[]string, steps *int) bool {
	switch self.dir {
	case "up":
		return self.moveVert(self.row-1, rows, steps)
	case "right":
		return self.moveHoriz(self.col+1, rows, steps)
	case "down":
		return self.moveVert(self.row+1, rows, steps)
	case "left":
		return self.moveHoriz(self.col-1, rows, steps)
	}

	return false
}

func main() {
	input, err := os.ReadFile("day_six/input.txt")
	if err != nil {
		panic(err)
	}

	rows := strings.Split(strings.TrimSuffix(string(input), "\n"), "\n")

	// partOne(rows)
	partTwo(rows)
}

func partOne(rows []string) {
	row, col := locateGuard(rows)

	guard := Guard{row: row, col: col, dir: up, stuck: false, turns: []string{}}
	var steps int

	for guard.move(&rows, &steps) {
	}

	steps += 1

	fmt.Printf("Number of steps %d\n", steps)
}

func partTwo(rows []string) {
	var loops int
	for i, row := range rows {
		for j, tile := range row {
			if tile == '#' || tile == '^' {
				continue
			}

			if hasLoop(rows, [2]int{i, j}) {
				loops += 1
			}
		}
	}

	fmt.Printf("Number of possible obstructions: %d\n", loops)
}

func locateGuard(rows []string) (int, int) {
	for i, row := range rows {
		for j, char := range row {
			if char == '^' {
				return i, j
			}
		}
	}

	return -1, -1
}

func markSpot(str string, col int, r rune) string {
	runes := []rune(str)
	runes[col] = r
	return string(runes)
}

func hasLoop(grid []string, obstruction [2]int) bool {
	row, col := locateGuard(grid)

	guard := Guard{row: row, col: col, dir: up, stuck: false, turns: []string{}}
	// don't need this here but i'm too lazy to refactor
	var steps int

	grid[obstruction[0]] = markSpot(grid[obstruction[0]], obstruction[1], '#')

	for guard.move(&grid, &steps) {
		if guard.stuck {
			grid[obstruction[0]] = markSpot(grid[obstruction[0]], obstruction[1], '.')
			return true
		}
	}

	grid[obstruction[0]] = markSpot(grid[obstruction[0]], obstruction[1], '.')
	return false
}
