package main

import (
	"fmt"
	"os"
	"regexp"
)

func part1() {
	file, _ := os.ReadFile("day8.input")
	matcher := regexp.MustCompile(`(\w+)`)
	// Should probably just do this in one pass
	matches := matcher.FindAllString(string(file), -1)

	sequence := matches[0]
	matches = matches[1:]
	paths := map[string][]string{}

	// loop through array 3 at a time
	for idx := 0; idx < len(matches)-3; idx += 3 {
		paths[matches[idx]] = append(paths[matches[idx]], matches[idx+1], matches[idx+2])
	}

	path := "AAA"
	total_path := 0
	for path != "ZZZ" {
		for _, choice := range sequence {
			idx := 0
			if choice == 'L' {
				idx = 0
			} else {
				idx = 1
			}
			path = paths[path][idx]
			total_path++
			if path == "ZZZ" {
				break
			}
		}
	}
	fmt.Println(total_path)
}

func part2() {
	file, _ := os.ReadFile("day8.input")
	matcher := regexp.MustCompile(`(\w+)`)
	// Should probably just do this in one pass
	matches := matcher.FindAllString(string(file), -1)

	sequence := matches[0]
	matches = matches[1:]
	paths := map[string][]string{}

	// loop through array 3 at a time
	for idx := 0; idx < len(matches)-2; idx += 3 {
		paths[matches[idx]] = append(paths[matches[idx]], matches[idx+1], matches[idx+2])
	}

	path := []string{}
	for k := range paths {
		if k[len(k)-1] == 'A' {
			path = append(path, k)
		}
	}

	total_path := 0
	all_z := false
	// lets find out how long until I get back to the start of this path
	length_of_path := make([]uint64, len(path))
	for i := range length_of_path {
		length_of_path[i] = 0
	}

	for !all_z {
		for _, choice := range sequence {
			// Get left or right index
			idx := 0
			if choice == 'L' {
				idx = 0
			} else {
				idx = 1
			}

			// Update all paths with L or R
			for i, p := range path {
				path[i] = string(paths[p][idx])
				if string(path[i][len(path[i])-1]) == "Z" {
					if length_of_path[i] == 0 {
						length_of_path[i] = uint64(total_path + 1)
						filled := true
						for _, l := range length_of_path {
							if l == 0 {
								filled = false
							}
						}
						if filled {
							all_z = true
						}
					}
				}
			}
			total_path++
		}
	}
	fmt.Println(lcm(length_of_path))
}

func gcd(a, b uint64) uint64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)

}

func lcm(nums []uint64) uint64 {
	lcm := nums[0]
	for i := 1; i < len(nums); i++ {
		lcm = lcm * nums[i] / gcd(lcm, nums[i])
	}
	return lcm
}

func main() {
	part1()
	part2()
}
