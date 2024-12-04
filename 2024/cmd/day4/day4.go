package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func CheckIndex(i, j int, puzzle [][]rune) int {
	i_max := len(puzzle) - 1
	j_max := len(puzzle[0]) - 1

	words := []string{"XMAS", "SAMX"}
	word_len := len(words[0]) - 1
	matches := 0

	// Check if we would be out of bounds in any direction
	// Horizontal
	if j+word_len <= j_max {
		for _, word := range words {
			match := true
			tmp_j := j
			for _, letter := range word {
				if letter != puzzle[i][tmp_j] {
					match = false
					break
				}
				tmp_j++
			}
			if match {
				matches++
			}
		}
	}

	// Vertical
	if i+word_len <= i_max {
		for _, word := range words {
			match := true
			tmp_i := i
			for _, letter := range word {
				if letter != puzzle[tmp_i][j] {
					match = false
					break
				}
				tmp_i++
			}
			if match {
				matches++
			}
		}
	}

	// Forward diagonal
	if i+word_len <= i_max && j+word_len <= j_max {
		for _, word := range words {
			match := true
			tmp_i := i
			tmp_j := j
			for _, letter := range word {
				if letter != puzzle[tmp_i][tmp_j] {
					match = false
					break
				}
				tmp_i++
				tmp_j++
			}
			if match {
				matches++
			}
		}
	}

	// Reverse diagonal
	if i+word_len <= i_max && j-word_len >= 0 {
		for _, word := range words {
			match := true
			tmp_i := i
			tmp_j := j
			for _, letter := range word {
				if letter != puzzle[tmp_i][tmp_j] {
					match = false
					break
				}
				tmp_i++
				tmp_j--
			}
			if match {
				matches++
			}
		}
	}

	return matches
}

func CheckIndex2(i, j int, puzzle [][]rune) int {
	// Return early if we arent in the middle of our X
	if puzzle[i][j] != 'A' {
		return 0
	}

	i_max := len(puzzle) - 1
	j_max := len(puzzle[0]) - 1

	matches := 0

	if i+1 <= i_max && j+1 <= j_max && i-1 >= 0 && j-1 >= 0 {
		if (puzzle[i-1][j-1] == 'M' && puzzle[i+1][j+1] == 'S' || puzzle[i-1][j-1] == 'S' && puzzle[i+1][j+1] == 'M') && (puzzle[i-1][j+1] == 'M' && puzzle[i+1][j-1] == 'S' || puzzle[i-1][j+1] == 'S' && puzzle[i+1][j-1] == 'M') {
			matches++
		}
	}

	return matches
}

func main() {
	file, err := os.Open("input.txt")
	puzzle := [][]rune{}
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := []rune{}
		for _, val := range scanner.Text() {
			line = append(line, val)
		}
		puzzle = append(puzzle, line)
	}

	total_1 := 0
	total_2 := 0
	for i := range puzzle {
		for j := range puzzle[0] {
			total_1 += CheckIndex(i, j, puzzle)
			total_2 += CheckIndex2(i, j, puzzle)
		}
	}
	fmt.Println("Part 1 Answer", total_1)
	fmt.Println("Part 2 Answer", total_2)

}
