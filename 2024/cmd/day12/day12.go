package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
)

var directions = []image.Point{
	{0, 1},  // Right
	{1, 0},  // Down
	{0, -1}, // Left
	{-1, 0}, // Up
}

// Check if a point exists in the set
func pointExists(p image.Point, pointSet map[image.Point]bool) bool {
	return pointSet[p]
}

func countCorners(points []image.Point) int {
	pointSet := make(map[image.Point]bool) // Use a map for fast lookups
	for _, p := range points {
		pointSet[p] = true
	}

	cornerCount := 0

	for _, p := range points {
		isCorner := false
        // Thanks random reddit comment https://www.reddit.com/r/adventofcode/comments/1hchskj/comment/m1po5k6/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button
		for i, dir := range directions {
			neighbor := p.Add(dir)
            nextDirection := getNextDirection(i)

			if pointExists(neighbor, pointSet) {
				continue
			}

			clockwiseNeighbor := p.Add(nextDirection)

			if !pointExists(clockwiseNeighbor, pointSet) {
				isCorner = true
			}

			// Check the diagonal neighbor
			diagonal := p.Add(image.Point{X: dir.X + nextDirection.X, Y: dir.Y + nextDirection.Y})
			if pointExists(diagonal, pointSet) {
				isCorner = true
			}

			if isCorner {
				cornerCount++
                isCorner = false
			}
		}
	}

	return cornerCount
}

func getNextDirection(i int) image.Point{
    return directions[(i+1)%len(directions)]
}

func BFS(pos image.Point, trailMap [][]string, visited map[image.Point]bool) (int, int, int) {
	type QueueItem struct {
		pos   image.Point
		count int
	}
	queue := []QueueItem{}
	queue = append(queue, QueueItem{pos, 0})

	area := 0
	perimeter := 0
	points := []image.Point{}

	if visited[pos] {
		return 0, 0, 0
	}

	visited[pos] = true

	for len(queue) > 0 {
		queueItem := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		point := queueItem.pos
		count := queueItem.count
		area++
		points = append(points, point)

		for _, dir := range directions {
			newPoint := point.Add(dir)

			if newPoint.X >= len(trailMap) || newPoint.Y >= len(trailMap[0]) || newPoint.X < 0 || newPoint.Y < 0 || trailMap[newPoint.X][newPoint.Y] != trailMap[pos.X][pos.Y] {
				perimeter++
				continue
			}

			if visited[newPoint] {
				continue
			}

			visited[newPoint] = true

			queue = append(queue, QueueItem{newPoint, count + 1})
		}
	}
	return area, perimeter, countCorners(points)
}

func main() {
	plots := [][]string{}

	file, _ := os.Open("input.txt")
	total1 := 0
	total2 := 0

	scanner := bufio.NewScanner(file)
	visited := map[image.Point]bool{}

	for scanner.Scan() {
		plants := scanner.Text()
		row := []string{}

		for _, plant := range plants {
			row = append(row, string(plant))
		}
		plots = append(plots, row)
	}
	for i := 0; i < len(plots); i++ {
		for j := 0; j < len(plots[0]); j++ {
			area, perimeter, sides := BFS(image.Point{i, j}, plots, visited)
			total1 += (area * perimeter)
			total2 += (area * sides)
		}
	}

	fmt.Println("Part 1 Answer:", total1)
	fmt.Println("Part 2 Answer:", total2)
}
