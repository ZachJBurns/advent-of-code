package main

import (
	"aoc2024/internal/helper"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func correctOrder(order []string, rules_map map[string]map[string]bool) int {
	if len(order) == 0 {
		return 0
	}
	match := true
	for i := 0; i < len(order)-1; i++ {
		if _, ok := rules_map[string(order[i])][string(order[i+1])]; !ok {
			match = false
		}
	}
	if match {
		val, _ := strconv.Atoi(string(order[len(order)/2]))
		return val
	}
	return 0
}

func main() {
	file := helper.ReadFile("input.txt")

	out := strings.Split(file, "\n\n")
	rules := out[0]
	ordering := out[1]

	rules_map := map[string]map[string]bool{}

	for _, val := range strings.Split(rules, "\n") {
		split_rules := strings.Split(val, "|")

		if _, ok := rules_map[split_rules[0]]; !ok {
			rules_map[split_rules[0]] = map[string]bool{split_rules[1]: true}
		} else {
			holder := rules_map[split_rules[0]]
			holder[split_rules[1]] = true
			rules_map[split_rules[0]] = holder
		}
	}

	total_1 := 0
	total_2 := 0
	for _, order := range strings.Split(ordering, "\n") {
		if len(order) == 0 {
			continue
		}
		order := strings.Split(order, ",")
		val := correctOrder(order, rules_map)

		if val == 0 {
			sort.Slice(order, func(i, j int) bool {
				if rules_map[order[i]][order[j]] {
					return true
				} else if rules_map[order[j]][order[i]] {
					return false
				} else {
					return false
				}
			})
            total_2 += correctOrder(order, rules_map)
		} else {
			total_1 += val
		}

	}
	fmt.Println("Part 1 Total:", total_1)
	fmt.Println("Part 2 Total:", total_2)

}
