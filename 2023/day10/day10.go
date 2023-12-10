package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Point struct {
	x, y int
}

func part1() {
	file, _ := os.ReadFile("day10.input")
	text := strings.Split(string(file), "\n")
	text = text[:len(text)-1]

	var start_idx Point
	for i := range text {
		if f := strings.Index(text[i], "S"); f >= 0 {
			start_idx = Point{i, f}
		}
	}

	adjacency_list := map[Point][]Point{}
	// What if I create an adj list of all the points?
	for i := range text {
		for j := range text[i] {
			// add N S E W to adj list using indexes
			points := []Point{}
			switch text[i][j] {
			case 'S':
				points = append(points, Point{i, j + 1}, Point{i + 1, j}, Point{i - 1, j}, Point{i, j - 1})
				break
			case '|':
				// north south case
				points = append(points, Point{i + 1, j}, Point{i - 1, j})
				break
			case '-':
				// east west case
				points = append(points, Point{i, j + 1}, Point{i, j - 1})
				break
			case 'L':
				// north east case
				points = append(points, Point{i - 1, j}, Point{i, j + 1})
				break
			case 'J':
				// north west case
				points = append(points, Point{i, j - 1}, Point{i - 1, j})
				break
			case '7':
				// south west case
				points = append(points, Point{i + 1, j}, Point{i, j - 1})
				break
			case 'F':
				// south east case
				points = append(points, Point{i, j + 1}, Point{i + 1, j})
				break
			}
			for _, point := range points {
				// check bounds before adding to adj list
				if point.x >= 0 && point.x < len(text) && point.y >= 0 && point.y < len(text[0]) && text[point.x][point.y] != '.' {
					adjacency_list[Point{i, j}] = append(adjacency_list[Point{i, j}], point)
				}
			}
		}
	}
	points := DFS(adjacency_list, start_idx, start_idx)
	fmt.Println((len(points)) / 2)

}

func DFS(graph map[Point][]Point, start Point, target Point) []Point {
	points := []Point{}
	stack := []Point{start}
	visited := make(map[Point]bool)
	num := 0

	for len(stack) > 0 {
		num++
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[node] {
			continue
		}

		visited[node] = true
		points = append(points, node)

		// Reverse the order of neighbors to prioritize the first neighbor pushed onto the stack
		neighbors := graph[node]
		for i := len(neighbors) - 1; i >= 0; i-- {
			neighbor := neighbors[i]
			if !visited[neighbor] {
				stack = append(stack, neighbor)
			}
		}
	}

	return points
}

func part2() {
	file, _ := os.ReadFile("day10.input")
	text := strings.Split(string(file), "\n")
	text = text[:len(text)-1]

	var start_idx Point
	for i := range text {
		if f := strings.Index(text[i], "S"); f >= 0 {
			start_idx = Point{i, f}
		}
	}

	adjacency_list := map[Point][]Point{}
	// What if I create an adj list of all the points?
	for i := range text {
		for j := range text[i] {
			// add N S E W to adj list using indexes
			points := []Point{}
			switch text[i][j] {
			case 'S':
				points = append(points, Point{i, j + 1}, Point{i + 1, j}, Point{i - 1, j}, Point{i, j - 1})
				break
			case '|':
				// north south case
				points = append(points, Point{i + 1, j}, Point{i - 1, j})
				break
			case '-':
				// east west case
				points = append(points, Point{i, j + 1}, Point{i, j - 1})
				break
			case 'L':
				// north east case
				points = append(points, Point{i - 1, j}, Point{i, j + 1})
				break
			case 'J':
				// north west case
				points = append(points, Point{i, j - 1}, Point{i - 1, j})
				break
			case '7':
				// south west case
				points = append(points, Point{i + 1, j}, Point{i, j - 1})
				break
			case 'F':
				// south east case
				points = append(points, Point{i, j + 1}, Point{i + 1, j})
				break
			}
			for _, point := range points {
				// check bounds before adding to adj list
				if point.x >= 0 && point.x < len(text) && point.y >= 0 && point.y < len(text[0]) && text[point.x][point.y] != '.' {
					adjacency_list[Point{i, j}] = append(adjacency_list[Point{i, j}], point)
				}
			}
		}
	}

	points := DFS(adjacency_list, start_idx, start_idx)
	total := 0
	for i := range text {
		for j := range text[0] {
			// Filter out points on the path so the polygon check doesn't overlap
			if !slices.Contains(points, Point{i, j}) {
				if point_in_polygon(text, Point{i, j}, points) {
					total++
				}
			}
		}
	}
	fmt.Println(total)
}

// TIL
func point_in_polygon(points []string, point Point, path []Point) bool {
	intersectCount := 0

	for i := 0; i < len(path); i++ {
		next := (i + 1) % len(path)

		if (path[i].y >= point.y) != (path[next].y >= point.y) &&
			(point.x < (path[next].x-path[i].x)*(point.y-path[i].y)/(path[next].y-path[i].y)+path[i].x) {
			intersectCount++
		}
	}
	return intersectCount%2 == 1
}

func main() {
	part1()
	part2()
}
