package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("day_five/input.txt")
	if err != nil {
		panic(err)
	}

	split := strings.Split(strings.TrimSuffix(string(input), "\n"), "\n\n")

	rules := strings.Split(split[0], "\n")
	pages := strings.Split(split[1], "\n")

	// partOne(rules, pages)
	partTwo(rules, pages)
}

func partOne(rulesArr []string, pageLists []string) {
	rules := populateRules(rulesArr)
	var res int

outer:
	for _, pageList := range pageLists {
		pages := strings.Split(strings.TrimSuffix(pageList, "\n"), ",")
		for i, page := range pages {
			rule := rules[page]
			if rule == nil {
				continue
			}

			for j := 0; j < i; j++ {
				if slices.Index(rule, pages[j]) != -1 {
					continue outer
				}
			}
		}

		midIdx := math.Floor(float64((len(pages) - 1) / 2))
		mid, err := strconv.Atoi(pages[int32(midIdx)])
		if err != nil {
			panic(err)
		}
		res += mid
	}

	fmt.Printf("The result is: %d\n", res)
}

func partTwo(rulesArr []string, pageLists []string) {
	rules := populateRules(rulesArr)
	var res int

outer:
	for _, pageList := range pageLists {
		pages := strings.Split(strings.TrimSuffix(pageList, "\n"), ",")
		for i, page := range pages {
			rule := rules[page]
			if rule == nil {
				continue
			}

			for j := 0; j < i; j++ {
				if slices.Index(rule, pages[j]) != -1 {
					res += midAfterOrdered(rules, pages)
					continue outer
				}
			}
		}
	}

	fmt.Printf("The result is: %d\n", res)
}

func populateRules(rulesArr []string) map[string][]string {
	rules := make(map[string][]string)

	for _, rule := range rulesArr {
		split := strings.Split(rule, "|")

		if rules[split[0]] != nil {
			rules[split[0]] = append(rules[split[0]], split[1])
		} else {
			rules[split[0]] = []string{split[1]}
		}
	}

	return rules
}

func midAfterOrdered(rules map[string][]string, pages []string) int {
	ordered := make([]string, len(pages))

	for i, page := range pages {
		if i == 0 {
			ordered = append(ordered, page)
			continue
		}

		rule := rules[page]
		if rule == nil {
			ordered = append(ordered, page)
			continue
		}

		problemIdx := earliest(ordered, rule)

		if problemIdx != -1 {
			ordered = insert(ordered, problemIdx, page)
		} else {
			ordered = append(ordered, page)
		}
	}

	var trimmed []string
	for _, o := range ordered {
		if o != "" {
			trimmed = append(trimmed, o)
		}
	}

	midIdx := math.Floor(float64((len(trimmed) - 1) / 2))
	mid, err := strconv.Atoi(trimmed[int32(midIdx)])
	if err != nil {
		panic(err)
	}

	return mid
}

func earliest(pages []string, rule []string) int {
	idx := -1

	for i := 0; i < len(rule); i++ {
		problem := slices.Index(pages, rule[i])

		if problem != -1 {
			if idx == -1 || problem < idx {
				idx = problem
			}
		}
	}

	return idx
}

func insert(slice []string, idx int, val string) []string {
	if len(slice) == idx {
		return append(slice, val)
	}

	slice = append(slice[:idx+1], slice[idx:]...)
	slice[idx] = val
	return slice
}
