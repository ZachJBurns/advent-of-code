package main

import (
	"aoc2024/internal/helper"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	file := helper.ReadFile("input.txt")
	file = strings.ReplaceAll(file, "\n", "")

	stones := strings.Split(file, " ")
	puzzle := [][]int{{1, 25}, {2, 75}}
	for _, round := range puzzle {
		stoneFrequency := make(map[string]int)
		for _, stone := range stones {
			stoneFrequency[stone]++
		}

		for i := 0; i < round[1]; i++ {
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
		}

		totalStones := 0
		for _, count := range stoneFrequency {
			totalStones += count
		}

		fmt.Printf("Part %d Answer: %d\n", round[0], totalStones)
	}

}
