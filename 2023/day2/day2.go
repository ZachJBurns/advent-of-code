package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func part1() {
	file, _ := os.Open("day2.input")
	scanner := bufio.NewScanner(file)
	total := 0
	var matcher = regexp.MustCompile("([0-9]+|;|blue|red|green)")

	RED_LIMIT := 12
	GREEN_LIMIT := 13
	BLUE_LIMIT := 14

	for scanner.Scan() {
		game := scanner.Text()
		reds, greens, blues := 0, 0, 0
		cubes := 0
		game_over := false

		if game == "" {
			continue
		}

		matches := matcher.FindAllString(game, math.MaxInt64)

		game_id, _ := strconv.Atoi(matches[0])
		pulls := matches[1:]

		for _, pull := range pulls {
			switch pull {
			case ";":
				reds, greens, blues = 0, 0, 0
				break
			case "red":
				reds += cubes
				break
			case "green":
				greens += cubes
				break
			case "blue":
				blues += cubes
				break
			default:
				cubes, _ = strconv.Atoi(pull)
				break
			}

			if reds > RED_LIMIT || greens > GREEN_LIMIT || blues > BLUE_LIMIT {
				game_over = true
				break
			}
		}
		if !game_over {
			total += game_id
		}
	}
	fmt.Println(total)
}

func part2() {
	file, _ := os.Open("day2.input")
	scanner := bufio.NewScanner(file)
	total := 0
	var matcher = regexp.MustCompile("([0-9]+|;|blue|red|green)")

	for scanner.Scan() {
		game := scanner.Text()
		cubes := 0
		reds, greens, blues := 0, 0, 0
		max_reds, max_greens, max_blues := 0, 0, 0

		if game == "" {
			continue
		}

		matches := matcher.FindAllString(game, math.MaxInt64)

		pulls := matches[1:]

		for _, pull := range pulls {
			switch pull {
			case ";":
				reds, greens, blues = 0, 0, 0
				break
			case "red":
				reds += cubes
				max_reds = Max(reds, max_reds)
				break
			case "green":
				greens += cubes
				max_greens = Max(greens, max_greens)
				break
			case "blue":
				blues += cubes
				max_blues = Max(blues, max_blues)
				break
			default:
				cubes, _ = strconv.Atoi(pull)
				break
			}
		}

		total += max_reds * max_greens * max_blues
	}

	fmt.Println(total)
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	part1()
	part2()
}
