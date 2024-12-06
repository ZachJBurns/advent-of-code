package main

import (
	"aoc2024/internal/helper"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func ValidateAndGetMiddle(order []string, rules_map map[string]map[string]bool) int {
	if len(order) == 0 {
		return 0
	}

	for i := 0; i < len(order)-1; i++ {
		if _, ok := rules_map[string(order[i])][string(order[i+1])]; !ok {
			return 0
		}
	}

	val, _ := strconv.Atoi(string(order[len(order)/2]))
	return val
}

func main() {
	file := helper.ReadFile("input.txt")

	out := strings.SplitN(file, "\n\n", 2)
	rules := out[0]
	ordering := out[1]

	rules_map := map[string]map[string]bool{}

	for _, val := range strings.Split(rules, "\n") {
		split_rules := strings.Split(val, "|")

		if _, ok := rules_map[split_rules[0]]; !ok {
			rules_map[split_rules[0]] = make(map[string]bool)
		}

		rules_map[split_rules[0]][split_rules[1]] = true
	}

	total_1 := 0
	total_2 := 0

	for _, order := range strings.Split(ordering, "\n") {
		if len(order) == 0 {
			continue
		}

		order := strings.Split(order, ",")
		val := ValidateAndGetMiddle(order, rules_map)

		if val == 0 {
			sort.Slice(order, func(i, j int) bool {
				return rules_map[order[i]][order[j]]
			})
			total_2 += ValidateAndGetMiddle(order, rules_map)
		} else {
			total_1 += val
		}

	}

	fmt.Println("Part 1 Total:", total_1)
	fmt.Println("Part 2 Total:", total_2)
}
