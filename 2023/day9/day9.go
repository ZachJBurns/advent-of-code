package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1() {
	file, _ := os.ReadFile("day9.input")
	puzzles := strings.Split(string(file), "\n")
	matcher := regexp.MustCompile(`(-?\d+)`)
	total := 0
	for _, puzzle := range puzzles {
		if puzzle == "" {
			continue
		}
		local_puzzle := [][]int{}
		match := matcher.FindAllString(puzzle, -1)

		// lets make these all ints to make it easier
		convert_arr := []int{}
		for _, c := range match {
			convert, _ := strconv.Atoi(string(c))
			convert_arr = append(convert_arr, convert)
		}

		local_puzzle = append(local_puzzle, convert_arr)
		idx := 0
		all_zero := false
		for !all_zero {
			new_row := []int{}
			for i := 0; i < len(local_puzzle[idx])-1; i += 1 {
				new_row = append(new_row, local_puzzle[idx][i+1]-local_puzzle[idx][i])
			}
			local_puzzle = append(local_puzzle, new_row)
			zero := true
			for _, n := range new_row {
				if n != 0 {
					zero = false
				}
			}
			if zero {
				all_zero = true
			}
			idx++
		}
		// traverse backwards and total up
		local_total := 0
		for i := len(local_puzzle) - 1; i >= 0; i-- {
			local_total += local_puzzle[i][len(local_puzzle[i])-1]

		}
		total += local_total
	}
	fmt.Println(total)

}

func part2() {

	file, _ := os.ReadFile("day9.input")
	puzzles := strings.Split(string(file), "\n")
	matcher := regexp.MustCompile(`(-?\d+)`)
	total := 0
	for _, puzzle := range puzzles {
		if puzzle == "" {
			continue
		}
		local_puzzle := [][]int{}
		match := matcher.FindAllString(puzzle, -1)

		// lets make these all ints to make it easier
		convert_arr := []int{}
		for _, c := range match {
			convert, _ := strconv.Atoi(string(c))
			convert_arr = append(convert_arr, convert)
		}

		local_puzzle = append(local_puzzle, convert_arr)
		idx := 0
		all_zero := false
		for !all_zero {
			new_row := []int{}
			for i := 0; i < len(local_puzzle[idx])-1; i += 1 {
				new_row = append(new_row, local_puzzle[idx][i+1]-local_puzzle[idx][i])
			}
			local_puzzle = append(local_puzzle, new_row)
			zero := true
			for _, n := range new_row {
				if n != 0 {
					zero = false
				}
			}
			if zero {
				all_zero = true
			}
			idx++
		}
		// traverse backwards and total up
		local_total := 0
		for i := len(local_puzzle) - 1; i >= 0; i-- {
			local_total = local_puzzle[i][0] - local_total

		}
		total += local_total
	}
	fmt.Println(total)

}

func main() {
	part1()
	part2()
}
