package main

import (
	"aoc2024/internal/helper"
	"fmt"
	"image"
	"strings"
)

var directions = []image.Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

type State struct {
	position  image.Point
	direction image.Point
	score     int
	path      []image.Point
}

func formatPoint(pos, dir image.Point) string {
	return fmt.Sprintf("%d/%d/%d/%d", pos.X, pos.Y, dir.X, dir.Y)
}

func BFS(start, end, dir image.Point, puzzle [][]string) (int, int) {
	// point&dir -> current score
	visited := map[string]int{}
	queue := []State{{start, dir, 0, []image.Point{start}}}
	pointsMap := make(map[int][]image.Point)
	minScore := 1<<(64-1) - 1

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.position == end {
			if curr.score <= minScore {
				minScore = curr.score
				pointsMap[minScore] = append(pointsMap[minScore], curr.path...)
			}
			continue
		}

		for _, dir := range directions {
			nextPos := curr.position.Add(dir)
			if nextPos.X < 0 || nextPos.Y < 0 || nextPos.X >= len(puzzle) || nextPos.Y >= len(puzzle[0]) || puzzle[nextPos.X][nextPos.Y] == "#" {
				continue
			}

			nextScore := curr.score + 1
			if dir != curr.direction {
				nextScore += 1000
			}
			pointsIndex := formatPoint(nextPos, dir)
			if prev, ok := visited[pointsIndex]; ok {
				if prev < nextScore {
					continue
				}
			}

			visited[pointsIndex] = nextScore
			path := make([]image.Point, len(curr.path))
			copy(path, curr.path)
			queue = append(queue, State{nextPos, dir, nextScore, append(path, nextPos)})

		}

	}

	count := map[image.Point]bool{}
	for _, val := range pointsMap[minScore] {
		count[val] = true
	}

	return minScore, len(count)

}
func main() {
	maze := [][]string{}
	file := helper.ReadFile("input.txt")
	start := image.Point{}
	end := image.Point{}

	for i, row := range strings.Split(file, "\n") {
		mazeRow := []string{}
		for j, ob := range row {
			obStr := string(ob)
			if obStr == "S" {
				start = image.Point{i, j}
			}
			if obStr == "E" {

				end = image.Point{i, j}
			}
			mazeRow = append(mazeRow, obStr)
		}
		maze = append(maze, mazeRow)
	}
	// get rid of empty line
	maze = maze[:len(maze)-1]

	minScore, total := BFS(start, end, image.Point{0, 1}, maze)

	fmt.Println("Part 1 Answer:", minScore)
	fmt.Println("Part 2 Answer:", total)
}

