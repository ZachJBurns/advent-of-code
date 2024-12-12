package main

import (
	"aoc2024/internal/helper"
	"fmt"
	"strconv"
	"strings"
)

func countStones(freq map[string]int) int {
	totalStones := 0
	for _, count := range freq {
		totalStones += count
	}
    return totalStones
}

func main() {
	file := helper.ReadFile("input.txt")
	file = strings.ReplaceAll(file, "\n", "")

	stones := strings.Split(file, " ")
	stoneFrequency := make(map[string]int)
	for _, stone := range stones {
		stoneFrequency[stone]++
	}

	for i := 0; i < 75; i++ {
		updatedFrequency := make(map[string]int)
		for stone, count := range stoneFrequency {
			if stone == "0" {
				updatedFrequency["1"] += count
			} else if len(stone)%2 == 0 {
				first := stone[:len(stone)/2]
				second := stone[len(stone)/2:]

				if len(second) > 1 && second[0] == '0' {
					stoneInt, _ := strconv.Atoi(second)
					second = strconv.Itoa(stoneInt)
				}
				updatedFrequency[first] += count
				updatedFrequency[second] += count
			} else {
				stoneInt, _ := strconv.Atoi(stone)
				updatedFrequency[strconv.Itoa(stoneInt*2024)] += count
			}
		}

		stoneFrequency = updatedFrequency
		if i == 24 {
			fmt.Println("Part 1 Answer:", countStones(stoneFrequency))
		}
	}

	fmt.Println("Part 2 Answer:", countStones(stoneFrequency))
}
