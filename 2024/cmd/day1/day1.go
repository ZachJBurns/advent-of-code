package main

import (
	"aoc2024/internal/helper"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file := helper.ReadFile("input.txt")
	left := []int{}
	right := []int{}
	for _, val := range strings.Split(file, "\n") {
		values := strings.Split(val, "   ")
		if len(values) != 2 {
			continue
		}
		left_num, _ := strconv.Atoi(values[0])
		right_num, _ := strconv.Atoi(values[1])
		left = append(left, left_num)
		right = append(right, right_num)
	}

	slices.Sort(left)
	slices.Sort(right)
	out := int64(0)
	for i := 0; i < len(left); i++ {
		out += int64(math.Abs(float64(left[i] - right[i])))
	}

	fmt.Println("Part 1 Answer: ", out)

	set := map[int]int{}
	for _, val := range right {
		if _, exists := set[val]; exists {
			set[val] += 1
		} else {
			set[val] = 1
		}
	}
	out = int64(0)
	for _, left_val := range left {
		if count, exists := set[left_val]; exists {
			out += int64(left_val * count)
		}
	}

	fmt.Println("Part 2 Answer: ", out)

}
