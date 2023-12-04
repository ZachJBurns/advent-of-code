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
	file, _ := os.Open("day4.input")
	scanner := bufio.NewScanner(file)
	matcher := regexp.MustCompile(`([\d\|]+)`)
	total := 0

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		scratchoff_matches := map[string]bool{}
		matches := matcher.FindAllString(text, -1)
		// The first number is the game number
		matches = matches[1:]
		my_card := true
		winning_matches := 0
		for _, match := range matches {
			if match == "|" {
				my_card = false
			} else if my_card {
				scratchoff_matches[match] = true
			} else {
				if scratchoff_matches[match] {
					winning_matches++
				}
			}
		}
		if winning_matches > 1 {
			total += powInt(2, winning_matches-1)
		} else if winning_matches == 1 {
			total += 1
		}

	}
	fmt.Println(total)
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func part2() {
	file, _ := os.Open("day4.input")
	scanner := bufio.NewScanner(file)
	matcher := regexp.MustCompile(`([\d\|]+)`)
	total := 0
	scratchoff_copies := map[int]int{}

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}

		scratchoff_matches := map[string]bool{}
		matches := matcher.FindAllString(text, -1)
		game := matches[0]
		game_number, _ := strconv.Atoi(game)

		// Add the original for the current game
		scratchoff_copies[game_number] += 1

		matches = matches[1:]
		my_card := true
		winning_matches := 0
		for _, match := range matches {
			if match == "|" {
				my_card = false
			} else if my_card {
				scratchoff_matches[match] = true
			} else {
				if scratchoff_matches[match] {
					winning_matches++
				}
			}
		}

		if winning_matches > 0 {
			// Add for the copies
			for i := 1; i <= winning_matches; i++ {
				scratchoff_copies[game_number+i] += (scratchoff_copies[game_number])
			}
		}
	}

	for _, v := range scratchoff_copies {
		total += v
	}

	fmt.Println(total)
}

func main() {
	part1()
	part2()
}
