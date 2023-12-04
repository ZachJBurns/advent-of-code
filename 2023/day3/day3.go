package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func find_symbols_part_1(grid []string, row int, col int) bool {
	max_col := len(grid[0]) - 1
	max_row := len(grid) - 1
	ret := false
	//Check in a box around current index.
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			// check bounds of index first to ensure we're looking at a valid area.
			if i < 0 || i > max_row || j < 0 || j > max_col {
				continue
			}
			// if the index is a special symbol and not a . return true
			if grid[i][j] != '.' && !unicode.IsDigit(rune(grid[i][j])) {
				ret = true
			}
		}
	}

	return ret
}

func part1() {
	file, _ := os.ReadFile("day3.input")
	total := 0
	grid := strings.Split(string(file), "\n")

	// Remove the empty string
	grid = grid[:len(grid)-2]
	for i, row := range grid {
		if len(row) == 0 {
			continue
		}
		num := ""
		symbol_found := false
		for j, letter := range row {
			if unicode.IsDigit(rune(letter)) {
				num += string(letter)
				if find_symbols_part_1(grid, i, j) {
					symbol_found = true

				}
			} else {
				if num != "" {
					// convert string to int and add to total
					num_int, _ := strconv.Atoi(num)
					num = ""
					if symbol_found {
						total += num_int
						symbol_found = false
					}
				}
			}
		}
		// We might have ended on a number
		if num != "" {
			num_int, _ := strconv.Atoi(num)
			if symbol_found {
				total += num_int
			}
		}
	}
	fmt.Println(total)
}

type Index struct{ r, c int }

func find_symbols_part_2(grid []string, row int, col int) int {
	max_col := len(grid[0]) - 1
	max_row := len(grid) - 1
	close_numbers := map[int]bool{}
	//Check in a box around current index.
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if i < 0 || i > max_row || j < 0 || j > max_col {
				continue
			}
			if unicode.IsDigit(rune(grid[i][j])) {
				num := string(grid[i][j])
				// Build int string by checking the left and the right
				for right := 1; j+right <= max_col && unicode.IsDigit(rune(grid[i][j+right])); right++ {
					num += string(grid[i][j+right])
				}

				for left := 1; j-left >= 0 && unicode.IsDigit(rune(grid[i][j-left])); left++ {
					num = string(grid[i][j-left]) + num
				}
				num_int, _ := strconv.Atoi(string(num))
				close_numbers[num_int] = true

			}
		}
	}
	total := 1
	if len(close_numbers) == 2 {
		for k := range close_numbers {
			total *= k
		}
		return total
	}

	return 0
}

func part2() {
	file, _ := os.ReadFile("day3.input")
	total := 0
	grid := strings.Split(string(file), "\n")

	// Remove the empty string
	grid = grid[:len(grid)-2]
	for i, row := range grid {
		if len(row) == 0 {
			continue
		}
		for j, letter := range row {
			if letter == '*' {
				total += find_symbols_part_2(grid, i, j)
			}
		}
	}
	fmt.Println(total)
}

func main() {
	part1()
	part2()
}
