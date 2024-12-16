package main

import (
	"aoc2024/internal/helper"
	"fmt"
	"image"
	"strings"
)

func copyPuzzle(puzzle [][]string) [][]string {
	puzzleCopy := make([][]string, len(puzzle))
	for i := range puzzle {
		puzzleCopy[i] = make([]string, len(puzzle[i]))
		copy(puzzleCopy[i], puzzle[i])
	}
	return puzzleCopy
}

func calculateDistances(puzzle [][]string) int {
	total := 0
	for y, row := range puzzle {
		for x, pos := range row {
			if pos == "O" || pos == "[" {
				total += (100 * y) + x
			}

		}
	}
	return total
}

func translatePuzzle(puzzle [][]string) ([][]string, image.Point) {
	copyPuzzle := [][]string{}
	start := image.Point{}
	for _, row := range puzzle {
		copyPuzzleRow := []string{}
		for _, pos := range row {
			if pos == "#" {
				copyPuzzleRow = append(copyPuzzleRow, "#", "#")
			}
			if pos == "O" {
				copyPuzzleRow = append(copyPuzzleRow, "[", "]")
			}
			if pos == "." {
				copyPuzzleRow = append(copyPuzzleRow, ".", ".")
			}
			if pos == "@" {
				copyPuzzleRow = append(copyPuzzleRow, "@", ".")
				start = image.Point{len(copyPuzzle), len(copyPuzzleRow) - 2}
			}
		}
		copyPuzzle = append(copyPuzzle, copyPuzzleRow)
	}
	return copyPuzzle, start
}

func BFS(pos, dir image.Point, puzzle [][]string) (bool, [][]string) {
	puzzleCopy := copyPuzzle(puzzle)
	state := []image.Point{pos}
	queue := []image.Point{pos}
	isValid := true
	visited := map[image.Point]bool{}

	visited[pos] = true
	for len(queue) > 0 {
		move := queue[0]
		queue = queue[1:]

		next := move.Add(dir)
		if visited[next] {
			continue
		}

		visited[next] = true
		if next.X < 0 || next.Y < 0 || next.X >= len(puzzle) || next.Y >= len(puzzle[0]) {
			isValid = false
			break
		}
		if puzzle[next.X][next.Y] == "#" {
			isValid = false
			break
		}

		if puzzle[next.X][next.Y] == "." {
			continue
		}

		queue = append(queue, next)
		state = append(state, next)

		if puzzle[next.X][next.Y] == "]" && (dir == image.Point{-1, 0} || dir == image.Point{1, 0}) {
			state = append(state, next.Add(image.Point{0, -1}))
			queue = append(queue, next.Add(image.Point{0, -1}))
		}

		if puzzle[next.X][next.Y] == "[" && (dir == image.Point{-1, 0} || dir == image.Point{1, 0}) {
			state = append(state, next.Add(image.Point{0, 1}))
			queue = append(queue, next.Add(image.Point{0, 1}))
		}
	}

	if isValid {
		for i := 0; i < len(state); i++ {
			curr := state[i]
			next:= state[i].Add(dir)
			puzzleCopy[next.X][next.Y] = "."
			puzzleCopy[curr.X][curr.Y] = "."
		}
		for i := 0; i < len(state); i++ {
			next := state[i].Add(dir)
			curr := state[i]
			puzzleCopy[next.X][next.Y] = puzzle[curr.X][curr.Y]
		}

	}
	return isValid, puzzleCopy

}

func main() {
	file := helper.ReadFile("input.txt")
	split := strings.Split(file, "\n\n")
	puzzleStr, moves := split[0], strings.ReplaceAll(split[1], "\n", "")
	puzzle := [][]string{}
	start := image.Point{}
	for i, row := range strings.Split(puzzleStr, "\n") {
		puzzleRow := []string{}
		for j, pos := range row {
			strPos := string(pos)
			if strPos == "@" {
				start = image.Point{i, j}
			}
			puzzleRow = append(puzzleRow, strPos)
		}
		puzzle = append(puzzle, puzzleRow)

	}
	copyPuzzle, copyStart := translatePuzzle(puzzle)

	for _, m := range moves {
		isValid := false
		dir := image.Point{}
		switch m {
		case '<':
			dir = image.Point{0, -1}
		case '>':
			dir = image.Point{0, 1}
		case '^':
			dir = image.Point{-1, 0}
		case 'v':
			dir = image.Point{1, 0}
		}
		isValid, puzzle = BFS(start, dir, puzzle)
		if isValid {
			start = start.Add(dir)
		}
		isValid, copyPuzzle = BFS(copyStart, dir, copyPuzzle)
		if isValid {
			copyStart= copyStart.Add(dir)
		}
	}
	fmt.Println("Part 1 Answer:",calculateDistances(puzzle))
	fmt.Println("Part 2 Answer:",calculateDistances(copyPuzzle))
}
