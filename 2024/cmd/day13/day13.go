package main

import (
	"aoc2024/internal/helper"
	"fmt"
	"image"
	"math"
	"regexp"
	"strconv"
)

func solveLinearEquation(total1, total2, x1, x2, y1, y2 int) (float64, float64) {
    // Just good ol solving systems of linear equations
    // Who knew I actually had to go back and learn things from school
	num1, den1 := float64(total1), float64(x1)
	coefB1 := float64(x2)
	num2, den2 := float64(total2), float64(y1)
	coefB2 := float64(y2)

	left := den2 * num1
	right := den1 * num2

	coefB := den2*coefB1 - den1*coefB2

	b := (left - right) / coefB
	a := (num1 - coefB1*b) / den1

	return a, b
}

func BFS(buttons []image.Point, target image.Point) int {
	type State struct {
		curr  image.Point
		count [2]int
	}

	queue := []State{{image.Point{}, [2]int{}}}
	visited := map[image.Point]bool{
		queue[0].curr: true,
	}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		for i, button := range buttons {
			next := item.curr.Add(button)
			if next == target {
				item.count[i]++
				return item.count[0]*3 + item.count[1]
			}
			if !visited[next] && next.X <= target.X && next.Y <= target.Y {
				tmpCount := item.count
				tmpCount[i]++
				queue = append(queue, State{next, tmpCount})
				visited[next] = true
			}
		}
	}
	return -1
}

func main() {
	input := helper.ReadFile("input.txt")
	r := regexp.MustCompile(`\d+`)
	matches := r.FindAllString(input, -1)

	puzzles := [][]int{}
	for i := 0; i < len(matches); i += 6 {
		match := matches[i:i+6]
		puzzle := make([]int, 6)
		for j, group := range match {
			puzzle[j], _ = strconv.Atoi(group)
		}
		puzzles = append(puzzles, puzzle)
	}

	part1, part2, offset := 0, 0, 10000000000000

	for _, puzzle := range puzzles {
		A := image.Point{puzzle[0], puzzle[1]}
		B := image.Point{puzzle[2], puzzle[3]}
		target := image.Point{puzzle[4], puzzle[5]}

		total := BFS([]image.Point{A, B}, target)
		if total > 0 {
			part1 += total
		}

		a, b := solveLinearEquation(target.X+offset, target.Y+offset, A.X, B.X, A.Y, B.Y)
		if math.Ceil(a) == a && math.Ceil(b) == b {
			part2 += int(a)*3 + int(b)
		}
	}

	fmt.Println("Part 1 Answer:", part1)
	fmt.Println("Part 2 Answer:", part2)
}

