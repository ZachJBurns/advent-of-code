package main

import (
	"aoc2024/internal/helper"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func CheckAndCalculate(substring string) (int, error) {
	if !strings.HasPrefix(substring, "mul(") || !strings.HasSuffix(substring, ")") {
		return 0, errors.New("does not contain substrings")
	}

	content := substring[4 : len(substring)-1]

	if strings.Contains(content, " ") {
		return 0, errors.New("contains whitespace")
	}

	nums := strings.Split(content, ",")

	if len(nums) != 2 {
		return 0, errors.New("invalid length")
	}

	total := 1
	for _, part := range nums {
		if num, err := strconv.Atoi(part); err == nil {
			total *= num
			continue
		}
		return 0, errors.New("")
	}
	return total, nil

}

func IsDo(substring string) bool {
	return strings.HasPrefix(substring, "do()")
}

func IsUndo(substring string) bool {
	return strings.HasPrefix(substring, "don't()")
}

func main() {
	file := helper.ReadFile("test.txt")

	start := 0
	total := 0
	disabled := false
	for start < len(file) {
		found := false

		for end := start; end < len(file); end++ {
			sub := file[start : end+1]
			if IsDo(sub) {
				disabled = false
			}
			if IsUndo(sub) {
				disabled = true
			}

			ret, err := CheckAndCalculate(sub)
			if err == nil {
				if !disabled {
					total += ret
				}
				start = end + 1
				found = true
				break

			}
		}

		if !found {
			start++
		}
	}
	fmt.Println("Part 1 Answer:", total)

}
