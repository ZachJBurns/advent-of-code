package main

import (
	"bufio"
	"fmt"
	"image"
	"log"
	"os"
)

func GetNextDirection(dir int) int {
	if dir == 3 {
		return 0
	}
	dir++
	return dir
}

func main() {
	file, err := os.Open("input.txt")
	puzzle := [][]rune{}

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	pos := image.Point{0, 0}
	for scanner.Scan() {
		line := []rune{}
		for idx, val := range scanner.Text() {
			if val == '^' {
				pos.Y = idx
				pos.X = len(puzzle)
			}
			line = append(line, val)
		}
		puzzle = append(puzzle, line)
	}



	directions := []image.Point{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	direction := 0

	// Part 1
	makeKey := func(pos image.Point) string {
		return fmt.Sprintf("%d,%d", pos.X, pos.Y)
	}

	distinct_spots := 1
	part_1_pos := pos
	visited := map[string]bool{}
	visited[makeKey(part_1_pos)] = true

	for {
		next := part_1_pos.Add(directions[direction])

		if next.X >= len(puzzle) || next.X < 0 || next.Y >= len(puzzle[0]) || next.Y < 0 {
			break
		}

		if puzzle[next.X][next.Y] == '#' {
			direction = GetNextDirection(direction)
			continue
		}

		part_1_pos = next

		if puzzle[part_1_pos.X][part_1_pos.Y] == '.' {
			if _, ok := visited[makeKey(part_1_pos)]; !ok {
				distinct_spots++
				visited[makeKey(part_1_pos)] = true
			}
		}
	}


	// Part 2
	makeKeyWithDir := func(pos image.Point, dirIndex int) string {
		return fmt.Sprintf("%d,%d,%d", pos.X, pos.Y, dirIndex)
	}

	total_loops := 0
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[0]); j++ {
			visited := map[string]bool{}
			direction = 0
			part_2_pos := pos

			for {
				if visited[makeKeyWithDir(part_2_pos, direction)] {
					total_loops++
					break

				}
				visited[makeKeyWithDir(part_2_pos, direction)] = true

				next := part_2_pos.Add(directions[direction])

				if next.X >= len(puzzle) || next.X < 0 || next.Y >= len(puzzle[0]) || next.Y < 0 {
					break
				}

				if puzzle[next.X][next.Y] == '#' || (next.X == i && next.Y == j) {
					direction = GetNextDirection(direction)
					continue
				}

				part_2_pos = next

			}
		}
	}

	fmt.Println("Part 1 Answer:", distinct_spots)
	fmt.Println("Part 2 Answer:", total_loops)
}
