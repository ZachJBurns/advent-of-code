package main

import (
	"aoc2024/internal/helper"
	"fmt"
	"image"
	"strings"
)

func BFS(start, end image.Point, puzzle [][]string) map[image.Point]int {
	var directions = []image.Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	distance := map[image.Point]int{}
	queue := []image.Point{start}
	distance[start] = 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr == end {
			return distance
		}

		for _, dir := range directions {
			nextPos := curr.Add(dir)
			if nextPos.X < 0 || nextPos.Y < 0 || nextPos.X >= len(puzzle) || nextPos.Y >= len(puzzle[0]) || puzzle[nextPos.X][nextPos.Y] == "#" {
				continue
			}
			if _, exists := distance[nextPos]; exists {
				continue
			}
			distance[nextPos] = distance[curr] + 1
			queue = append(queue, nextPos)
		}
	}
	return map[image.Point]int{}
}

func distance(a, b image.Point) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	puzzle := [][]string{}
	file := helper.ReadFile("input.txt")
	start := image.Point{}
	end := image.Point{}
	points := []image.Point{}

	for i, r := range strings.Split(file, "\n") {
		row := []string{}
		for j, p := range r {
			part := string(p)
			row = append(row, part)
			if part == "S" {
				start = image.Point{i, j}
			}
			if part == "E" {
				end = image.Point{i, j}
			}
			if part != "#" {
				points = append(points, image.Point{i, j})
			}
		}
		puzzle = append(puzzle, row)
	}

	puzzle = puzzle[:len(puzzle)-1]

	distances := BFS(start, end, puzzle)
	p1, p2 := 0, 0
	p1Distance, p2Distance := 2, 20
	for i, d1 := range distances {
		for j, d2 := range distances {
			d := distance(i, j)
			if d == p1Distance {
				if d1-d2-d >= 100 {
					p1++
				}
			}
			if d <= p2Distance {
				if d1-d2-d >= 100 {
					p2++
				}
			}
		}
	}

	fmt.Println("Part 1 Answer:", p1)
	fmt.Println("Part 2 Answer:", p2)
}
