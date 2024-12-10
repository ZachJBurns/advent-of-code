package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)
	trailMap := [][]int{}
	trailHeads := []image.Point{}
	for scanner.Scan() {
		row := []int{}
		for idx, val := range scanner.Text() {
			trailPoint := val
			trailPointInt, err := strconv.Atoi(string(trailPoint))

			// Catch for test inputs
			if err != nil {
				trailPointInt = -1
			}
			if trailPointInt == 0 {
				trailHeads = append(trailHeads, image.Point{len(trailMap), idx})
			}
			row = append(row, trailPointInt)
		}
		trailMap = append(trailMap, row)
	}

	total1 := 0
	total2 := 0
	visited := make([][]bool, len(trailMap))
	for _, val := range trailHeads {
		for i := range visited {
			visited[i] = make([]bool, len(trailMap[0]))
		}
		visitedMap := visited
		total2 += track(val, 0, trailMap, visitedMap, false)
		total1 += track(val, 0, trailMap, visitedMap, true)

	}

	fmt.Println("Part 1 Answer:", total1)
	fmt.Println("Part 2 Answer:", total2)
}

func track(pos image.Point, count int, trailMap [][]int, visited [][]bool, enableVisiting bool) int {
	i := len(trailMap)
	j := len(trailMap[0])
	if pos.X >= i || pos.Y >= j || pos.X < 0 || pos.Y < 0 {
		return 0
	}

	if visited[pos.X][pos.Y] || trailMap[pos.X][pos.Y] != count {
		return 0
	}

	if enableVisiting {
		visited[pos.X][pos.Y] = true
	}

	if trailMap[pos.X][pos.Y] == 9 {
		return 1
	}

	return track(pos.Add(image.Point{-1, 0}), count+1, trailMap, visited, enableVisiting) + track(pos.Add(image.Point{0, 1}), count+1, trailMap, visited, enableVisiting) + track(pos.Add(image.Point{1, 0}), count+1, trailMap, visited, enableVisiting) + track(pos.Add(image.Point{0, -1}), count+1, trailMap, visited, enableVisiting)

}
