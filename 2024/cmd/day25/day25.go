package main

import (
	"aoc2024/internal/helper"
	"fmt"
	"strings"
)

func main() {
	file := helper.ReadFile("input.txt")
	puzzles := strings.Split(file, "\n\n")
	locks, keys := [][5]int{}, [][5]int{}
	for _, puzzle := range puzzles {
		split := strings.SplitN(puzzle, "\n", 7)
		if split[0] == "#####" {
			lock := [5]int{0, 0, 0, 0, 0}
			split = split[:len(split)-1]
			for _, i := range split {
				for idx, j := range i {
					if string(j) == "." {
						lock[idx]++
					}
				}
			}
			locks = append(locks, lock)

		} else {
				key := [5]int{0, 0, 0, 0, 0}
				split = split[:len(split)-1]
				for _, i := range split {
					for idx, j := range i {
						if string(j) == "#" {
							key[idx]++
						}
					}
				}
				keys = append(keys, key)

		}

	}

    total := 0
    for _, lock := range locks {
        for _, key := range keys {
            if keyFits(lock, key){
                total++
            }
        }
    }
    fmt.Println("Part 1 Answer:", total)

}
func keyFits(arr1, arr2 [5]int) bool {
    if len(arr1) != len(arr2) {
        return false
    }

    for i := 0; i < len(arr1); i++ {
        if arr1[i] <  arr2[i] {
            return false
        }
    }

    return true
}
