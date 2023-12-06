package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func part1() {
	file, _ := os.ReadFile("day6.input")
	file_string := string(file)
	matcher := regexp.MustCompile(`(\d+)`)
	matches := matcher.FindAllString(file_string, -1)

	time := matches[:len(matches)/2]
	distance := matches[len(matches)/2:]
	total := 1
	for idx := 0; idx < len(time); idx++ {
		local := 0
		race_time, _ := strconv.Atoi(time[idx])
		record_distance, _ := strconv.Atoi(distance[idx])

		for button := 1; button < race_time; button++ {
			if (race_time-button)*(button) > record_distance {
				local++
			}
		}
		total *= local
	}
	fmt.Println(total)
}

func part2() {
	file, _ := os.ReadFile("day6.input")
	file_string := string(file)
	matcher := regexp.MustCompile(`(\d+)`)
	matches := matcher.FindAllString(file_string, -1)

	time_arr := matches[:len(matches)/2]
	time_str := ""
	for _, t := range time_arr {
		time_str += t
	}

	distance_arr := matches[len(matches)/2:]
	distance_str := ""
	for _, d := range distance_arr {
		distance_str += d
	}

	race_time, _ := strconv.Atoi(time_str)
	record_distance, _ := strconv.Atoi(distance_str)

	total := 0
	for button := 1; button < race_time; button++ {

		if (race_time-button)*(button) > record_distance {
			total++
		}
	}
	fmt.Println(total)
}

func main() {
	part1()
	part2()
}
