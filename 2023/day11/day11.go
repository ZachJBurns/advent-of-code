package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Point struct{ x, y int }

func part1() {
	file, _ := os.ReadFile("day11.input")
	text := strings.Split(string(file), "\n")
	// remove the last empty line
	text = text[:len(text)-1]
	expanded_rows := make([]bool, len(text))
	expanded_cols := make([]bool, len(text[0]))

	// set default values
	for i := range expanded_rows {
		expanded_rows[i] = true
	}
	for i := range expanded_cols {
		expanded_cols[i] = true
	}

	for i := range text {
		for j := range text[i] {
			if text[i][j] != '.' {
				expanded_rows[i] = false
				expanded_cols[j] = false
			}
		}
	}
	for i := len(expanded_rows) - 1; i >= 0; i-- {
		if expanded_rows[i] {
			text = append(text[:i], append([]string{strings.Repeat(".", len(text[0]))}, text[i:]...)...)
		}
	}

	for i := len(expanded_rows) - 1; i >= 0; i-- {
		if expanded_cols[i] {
			// insert a new column at index i and fill with "."
			for j := range text {
				left := text[j][:i]
				right := text[j][i:]
				text[j] = left + "." + right
			}
		}
	}

	galaxies := []Point{}
	for i := range text {
		for j := range text[i] {
			if text[i][j] == '#' {
				galaxies = append(galaxies, Point{i, j})
			}
		}
	}
	total := 0
	visited := make(map[Point]bool)
	for i := range galaxies {
		for j := range galaxies {
			if i != j && !visited[Point{i, j}] && !visited[Point{j, i}] {
				total += int(math.Abs(float64(galaxies[i].x-galaxies[j].x)) + math.Abs(float64(galaxies[i].y-galaxies[j].y)))
				visited[Point{i, j}] = true
				visited[Point{j, i}] = true
			}
		}
	}
	fmt.Println(total)
}

func part2() {
	file, _ := os.ReadFile("day11.input")
	text := strings.Split(string(file), "\n")
	// remove the last empty line
	text = text[:len(text)-1]
	expanded_rows := make([]bool, len(text))
	expanded_cols := make([]bool, len(text[0]))

	// set default values
	for i := range expanded_rows {
		expanded_rows[i] = true
	}
	for i := range expanded_cols {
		expanded_cols[i] = true
	}

	for i := range text {
		for j := range text[i] {
			if text[i][j] != '.' {
				expanded_rows[i] = false
				expanded_cols[j] = false
			}
		}
	}

	galaxies := []Point{}
	for i := range text {
		for j := range text[i] {
			if text[i][j] == '#' {
				galaxies = append(galaxies, Point{i, j})
			}
		}
	}

	total := 0
	visited := make(map[Point]bool)
	for i := range galaxies {
		for j := range galaxies {
			if i != j && !visited[Point{i, j}] && !visited[Point{j, i}] {
				x_expansions := 0
				y_expansions := 0
				for idx := range expanded_rows {
					// count how many boundaraies are crossed from our current galaxy to the target galaxy
					if expanded_rows[idx] {
						if galaxies[i].x < galaxies[j].x || galaxies[i].x > galaxies[j].x {
							if idx > galaxies[i].x && idx < galaxies[j].x {
								x_expansions++
							}
						}
					}
					if expanded_cols[idx] {
						if galaxies[i].y < galaxies[j].y || galaxies[i].y > galaxies[j].y {
							if idx > galaxies[i].y && idx < galaxies[j].y {
								y_expansions++
							} else if idx < galaxies[i].y && idx > galaxies[j].y {
								y_expansions++
							}
						}
					}
				}
				if x_expansions > 0 {
					x_expansions *= 999_999
				}
				if y_expansions > 0 {
					y_expansions *= 999_999
				}
				total += int(math.Abs(float64(galaxies[i].x-galaxies[j].x)) + float64(x_expansions) + math.Abs(float64(galaxies[i].y-galaxies[j].y)) + float64(y_expansions))
				visited[Point{i, j}] = true
				visited[Point{j, i}] = true
			}
		}
	}
	fmt.Println(total)
}

func main() {
	part1()
	part2()
}
