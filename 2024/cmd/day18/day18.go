package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	bytes := []image.Point{}
	badPoints := map[image.Point]bool{}

	endPoint := image.Point{70, 70}
	numBytes := 1024

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		byteStr := scanner.Text()
		byteStrSplit := strings.Split(byteStr, ",")

		x, _ := strconv.Atoi(byteStrSplit[0])
		y, _ := strconv.Atoi(byteStrSplit[1])
		bytes = append(bytes, image.Point{x, y})
	}

	for i := 0; i < numBytes; i++ {
		badPoints[bytes[i]] = true
	}

	fmt.Println("Part 1 Answer:", BFS(image.Point{0, 0}, endPoint, badPoints))

	clear(badPoints)
	low, high := 0, len(bytes)-1
	result := image.Point{}
	for low <= high {
		mid := (low + high) / 2

		for i := 0; i <= mid; i++ {
			badPoints[bytes[i]] = true
		}

		if BFS(image.Point{0, 0}, endPoint, badPoints) == -1 {
			result = bytes[mid]
			high = mid - 1
		} else {
			low = mid + 1
		}
		clear(badPoints)
	}
	fmt.Println("Part 2 Answer:", result)
}

func BFS(start, end image.Point, badPoints map[image.Point]bool) int {
	type State struct {
		point    image.Point
		distance int
	}
	directions := []image.Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	queue := []State{{start, 0}}
	visited := map[image.Point]bool{}

	for len(queue) > 0 {
		pop := queue[0]
		queue = queue[1:]

		if pop.point == end {
			return pop.distance
		}
		if visited[pop.point] {
			continue
		}
		visited[pop.point] = true

		for _, dir := range directions {
			next := pop.point.Add(dir)

			if next.X > end.X || next.X < 0 || next.Y > end.Y || next.Y < 0 || badPoints[next] {
				continue
			}
			if visited[next] {
				continue
			}
			queue = append(queue, State{next, pop.distance + 1})
		}
	}
	return -1
}
