package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func backtrack(position int, total int, numbers []int, matches *bool, target int) {
	if *matches {
		return
	}

	if position == len(numbers) {
		if total == target {
			*matches = true
		}
		return
	}

	backtrack(position+1, total+numbers[position], numbers, matches, target)
	backtrack(position+1, total*numbers[position], numbers, matches, target)
}

func backtrackWithConcat(position int, total int, numbers []int, matches *bool, target int) {
	if *matches {
		return
	}

	if position == len(numbers) {
		if total == target {
			*matches = true
		}
		return
	}

	backtrackWithConcat(position+1, total+numbers[position], numbers, matches, target)
	backtrackWithConcat(position+1, total*numbers[position], numbers, matches, target)

	if position < len(numbers) {
		concatenated, _ := strconv.Atoi(fmt.Sprintf("%d%d", total, numbers[position]))
		backtrackWithConcat(position+1, concatenated, numbers, matches, target)
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	matches := false
	part1Total := 0
	part2Total := 0

	for scanner.Scan() {
		lineParts := strings.Split(scanner.Text(), ": ")
		target, _ := strconv.Atoi(lineParts[0])
		values := strings.Split(lineParts[1], " ")
		intValues := []int{}

		for _, value := range values {
			num, _ := strconv.Atoi(value)
			intValues = append(intValues, num)
		}

		backtrack(1, intValues[0], intValues, &matches, target)

		if matches {
			part1Total += target
		} else {
			backtrackWithConcat(0, 0, intValues, &matches, target)
			if matches {
				part2Total += target
			}
		}
		matches = false
	}

	fmt.Println("Part 1 Answer:", part1Total)
	fmt.Println("Part 2 Answer:", part1Total+part2Total)
}

