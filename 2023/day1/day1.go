package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
)

func part1() {
	file, _ := os.Open("day1.input")
	scanner := bufio.NewScanner(file)
	total := 0
	var matcher = regexp.MustCompile("([1-9])")
	for scanner.Scan() {
		text := scanner.Text()
		matches := matcher.FindAll([]byte(text), math.MaxInt64)

		if text == "" {
			continue
		}
		total += int(rune(matches[0][0])-'0')*10 + int(rune(matches[len(matches)-1][0])-'0')
	}

	fmt.Println(total)
}

func part2() {
	file, _ := os.Open("day1.input")
	scanner := bufio.NewScanner(file)
	total := 0
	spelled_digits := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

	var matcher = regexp.MustCompile("([1-9]|sevenine|eighthree|eightwo|threeight|oneight|twone|one|two|three|four|five|six|seven|eight|nine)")

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}

		matches := matcher.FindAllString(text, math.MaxInt64)

		last_element := len(matches) - 1

		if len(matches[0]) > 1 {
			switch matches[0] {
			case "sevenine":
				total += 70
				break
			case "eighthree":
				total += 80
				break
			case "twone":
				total += 20
				break
			case "eightwo":
				total += 80
				break
			case "threeight":
				total += 30
				break
			case "oneight":
				total += 10
				break
			default:
				total += spelled_digits[matches[0]] * 10
				break
			}
		} else {
			total += int([]rune(matches[0])[0]-'0') * 10
		}

		if len(matches[last_element]) > 1 {
			switch matches[last_element] {
			case "sevenine":
				total += 9
				break
			case "eighthree":
				total += 3
				break
			case "twone":
				total += 1
				break
			case "eightwo":
				total += 2
				break
			case "threeight":
				total += 8
				break
			case "oneight":
				total += 8
				break
			default:
				total += spelled_digits[matches[last_element]]
				break
			}
		} else {
			total += int([]rune(matches[last_element])[0] - '0')

		}
	}
	fmt.Println(total)
}

func main() {
	part1()
	part2()
}
