package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isSafe(nums []int) bool {
	is_increasing := false
	found_increase := false
	for i := 0; i < len(nums)-1; i++ {
		difference := nums[i+1] - nums[i]

		if !found_increase {
			if difference > 0 {
				is_increasing = true
			}
			found_increase = true
		}

		if difference > 0 && !is_increasing || difference < 0 && is_increasing {
			return false
		}

		if difference < 0 {
			difference = -difference
		}
		if difference != 1 && difference != 2 && difference != 3 {
			return false
		}

	}
	return true

}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	out_1 := 0
	out_2 := 0
	for scanner.Scan() {
		levels := []int{}
		split := strings.Split(scanner.Text(), " ")
		for _, val := range split {
			converted, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println(err)
			}
			levels = append(levels, converted)
		}
		if isSafe(levels) {
			out_1 += 1
		} else {
			// Lets try to brute force if we're 1 off of being safe
			for i := range levels {
				subset := append([]int{}, levels[:i]...)
				subset = append(subset, levels[i+1:]...)
				if isSafe(subset) {
					out_2 += 1
					break
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1 Answer:", out_1)
	fmt.Println("Part 2 Answer:", out_1+out_2)

}
