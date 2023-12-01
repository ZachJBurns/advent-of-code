package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func part1() {
	file, _ := os.Open("day1.input")
	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			continue
		}

		left := 0
		right := len(text) - 1
		found1 := false
		found2 := false

		for !found1 || !found2 {
			if unicode.IsDigit(rune(text[left])) && !found1 {
				found1 = true
				total += int(text[left]-'0') * 10
			}

			if unicode.IsDigit(rune(text[right])) && !found2 {
				found2 = true
				total += int(text[right] - '0')
			}
			left++
			right--
		}
	}

	fmt.Println(total)
}

func part2() {
	file, _ := os.Open("day1.input")
	scanner := bufio.NewScanner(file)
	total := 0
	spelled_digits := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}

		left := 0
		right := len(text) - 1

		left_window := ""
		right_window := ""

		found1 := false
		found2 := false

		for !found1 || !found2 {
			if !found1 {
				if unicode.IsDigit(rune(text[left])) {
					found1 = true
					total += int(text[left]-'0') * 10
					left_window = ""
				} else {
					if len(left_window) > 5 {
						left_window = left_window[1:]
					}

					for k, v := range spelled_digits {
						if strings.Contains(left_window, k) {
							total += v * 10
							found1 = true
							break
						}
					}
				}
			}

			if !found2 {
				if unicode.IsDigit(rune(text[right])) {
					found2 = true
					total += int(text[right] - '0')
					right_window = ""
				} else {
					right_window = string(text[right]) + right_window
					if len(right_window) > 5 {
						right_window = right_window[:len(right_window)-1]
					}

					for k, v := range spelled_digits {
						if strings.Contains(right_window, k) {
							total += v
							found2 = true
							break
						}
					}
				}
			}

			left++
			right--
		}

	}

	fmt.Println(total)
}

func main() {
	part1()
	part2()
}
