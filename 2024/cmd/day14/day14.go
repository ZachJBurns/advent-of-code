package main

import (
	"aoc2024/internal/helper"
	"fmt"
	"image"
	"regexp"
	"strconv"
	"strings"
)

type Robot struct {
	position image.Point
	velocity image.Point
}

func calculateFuturePosition(robots []Robot, seconds int, puzzle [2]int) []Robot {
	robotsCopy := make([]Robot, len(robots))
	copy(robotsCopy, robots)
	for i := 0; i < len(robotsCopy); i++ {
		currRobot := robotsCopy[i]
		pX := (currRobot.position.X + currRobot.velocity.X*seconds) % puzzle[0]
		if pX < 0 {
			pX += puzzle[0]
		}
		pY := (currRobot.position.Y + currRobot.velocity.Y*seconds) % puzzle[1]
		if pY < 0 {
			pY += puzzle[1]
		}
		currRobot.position = image.Point{pX, pY}
		robotsCopy[i] = currRobot
	}
	return robotsCopy
}

func findMaxGroupings(robots []Robot, bathroomSize [2]int, seconds int) [][2]int{
    counts := [][2]int{}
	for i := 1; i < seconds+1; i++ {
		updatedRobots := calculateFuturePosition(robots, i, bathroomSize)
		for _, p := range updatedRobots {
			count := 0

			for _, n := range updatedRobots {
				if p.position == n.position {
					continue
				}
				if distance(p.position, n.position) <= 2 {
					count++
				}
			}

            // Just a random grouping number
			if count > 10 {
                counts = append(counts, [2]int{i, count})
			}
		}
	}
    return counts
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

func printGrid(robots []Robot, bathroomSize [2]int, seconds int) {
	updatedRobots := calculateFuturePosition(robots, seconds, bathroomSize)

	for y := 0; y < bathroomSize[1]; y++ {
		for x := 0; x < bathroomSize[0]; x++ {
			found := false
			for _, robot := range updatedRobots {
				if robot.position.X == x && robot.position.Y == y {
					found = true
					break
				}
			}

			if found {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func getQuadrants(robots []Robot, puzzle [2]int) int {
	hMidpoint := (puzzle[0]) / 2
	vMidpoint := (puzzle[1]) / 2
	robotCount := [4]int{0, 0, 0, 0}

	getQuadrant := func(point image.Point) int {
		if point.X == hMidpoint || point.Y == vMidpoint {
			return -1
		}
		if point.X < hMidpoint && point.Y < vMidpoint {
			return 0
		}
		if point.X >= hMidpoint && point.Y < vMidpoint {
			return 1
		}
		if point.X < hMidpoint && point.Y >= vMidpoint {
			return 2
		}
		if point.X >= hMidpoint && point.Y >= vMidpoint {
			return 3
		}
		return -1
	}
	for _, robot := range robots {
		x := getQuadrant(robot.position)
		if x >= 0 {
			robotCount[x]++
		}
	}
	return robotCount[0] * robotCount[1] * robotCount[2] * robotCount[3]
}

func selectTime(possibilities [][2]int) int {
    counter := map[int]int{}
    for _,p := range possibilities {
        counter[p[0]]+=p[1]
    }

    max := 0
    bestTime := 0
    for k, v := range counter {
        if v > max {
            max = v
            bestTime = k
        }
    }
    return bestTime
}

func main() {
	file := helper.ReadFile("input.txt")
	r := regexp.MustCompile(`(-?\d+),(-?\d+)`)
	matches := r.FindAllString(file, -1)
	bathroomSize := [2]int{101, 103}
	robots := []Robot{}

	for i := 0; i < len(matches)-1; i += 2 {
		positions := strings.Split(matches[i], ",")
		x, _ := strconv.Atoi(positions[0])
		y, _ := strconv.Atoi(positions[1])
		velocities := strings.Split(matches[i+1], ",")
		vX, _ := strconv.Atoi(velocities[0])
		vY, _ := strconv.Atoi(velocities[1])
		robots = append(robots, Robot{image.Point{x, y}, image.Point{vX, vY}})
	}

	robots1 := calculateFuturePosition(robots, 100, bathroomSize)
	fmt.Println("Part 1 Answer:", getQuadrants(robots1, bathroomSize))
    ret := findMaxGroupings(robots, bathroomSize, 10000)
    time := selectTime(ret)
	printGrid(robots, bathroomSize, time)
	fmt.Println("Part 2 Answer:",time)
}
