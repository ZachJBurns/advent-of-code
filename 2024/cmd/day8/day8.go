package main

import (
	"aoc2024/internal/helper"
	"fmt"
	"image"
	"strings"
	"unicode"
)

func findAntinodes(p1, p2, diff image.Point) (r1, r2 image.Point) {
	return subtractPoints(p1, diff), addPoints(p2, diff)
}

func subtractPoints(p1, p2 image.Point) image.Point {
	return image.Point{p1.X - p2.X, p1.Y - p2.Y}
}

func addPoints(p1, p2 image.Point) image.Point {
	return p1.Add(p2)
}

func orderPoints(p1, p2 image.Point) (image.Point, image.Point) {
	if p1.X < p2.X || (p1.X == p2.X && p1.Y < p2.Y) {
		return p1, p2
	}
	return p2, p1
}

func pointDifference(p1, p2 image.Point) image.Point {
	return image.Point{p2.X - p1.X, p2.Y - p1.Y}
}

func inBounds(point image.Point, i, j int) bool {
	if point.X < i && point.X >= 0 && point.Y < j && point.Y >= 0 {
		return true
	}
	return false
}

func findPart2Antinodes(p1, p2, diff image.Point, i, j int) []image.Point {
	// close node
	ret := []image.Point{}
	p1 = subtractPoints(p1, diff)
	for inBounds(p1, i, j) {
		ret = append(ret, p1)
		p1 = subtractPoints(p1, diff)

	}

	// far node
	p2 = addPoints(p2, diff)
	for inBounds(p2, i, j) {
		ret = append(ret, p2)
		p2 = addPoints(p2, diff)

	}
	return ret
}

func main() {
	file := helper.ReadFile("input.txt")
	lines := strings.Split(file, "\n")
	lines = lines[:len(lines)-1]
	antennaMap := map[string][]image.Point{}
	pointMap := map[image.Point]string{}

	for i, row := range lines {
		for j := range row {
			if unicode.IsDigit(rune(row[j])) || unicode.IsLetter(rune(row[j])) {

				antenna := string(lines[i][j])
				pointMap[image.Point{i, j}] = antenna
				if val, ok := antennaMap[antenna]; ok {
					arr := val
					arr = append(arr, image.Point{i, j})
					antennaMap[antenna] = arr
				} else {
					antennaMap[antenna] = []image.Point{{i, j}}
				}
			}
		}
	}

	uniqueAntinodes := make(map[image.Point]bool)
	uniqueAntinodes2 := make(map[image.Point]bool)
	for _, points := range antennaMap {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				p1, p2 := orderPoints(points[i], points[j])
				diff := pointDifference(p1, p2)
				a1, a2 := findAntinodes(p1, p2, diff)

				if inBounds(a1, len(lines), len(lines[0])) {
					uniqueAntinodes[a1] = true
				}
				if inBounds(a2, len(lines), len(lines[0])) {
					uniqueAntinodes[a2] = true
				}

				part2Antinodes := findPart2Antinodes(p1, p2, diff, len(lines), len(lines[0]))
				for _, val := range part2Antinodes {
					if _, ok := pointMap[val]; !ok {
						uniqueAntinodes2[val] = true
					}
				}
			}
		}
	}

	fmt.Println("Part 1 Answer:", len(uniqueAntinodes))
	fmt.Println("Part 2 Answer:", len(uniqueAntinodes2)+len(pointMap))
}
