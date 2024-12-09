package main

import (
	"aoc2024/internal/helper"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	file := helper.ReadFile("input.txt")
	blocks := []string{}
	curr, p1, p2 := 0, 0, 1

	// number => [startPos, endPos]
	locationMapping := map[int][]int{}
	for p2 < len(file) {
		files, _ := strconv.Atoi(string(file[p1]))
		freeSpace, _ := strconv.Atoi(string(file[p2]))
		tmp := []int{len(blocks)}
		for j := 0; j < files; j++ {
			blocks = append(blocks, fmt.Sprintf("%d", curr))
		}
		tmp = append(tmp, len(blocks)-1)
		locationMapping[curr] = tmp
		for j := 0; j < freeSpace; j++ {
			blocks = append(blocks, ".")
		}

		curr++
		p1 += 2
		p2 += 2
	}

	left, right := 0, len(blocks)-1
	part2Blocks := make([]string, len(blocks))
	copy(part2Blocks, blocks)

	for left < right {
		if blocks[left] != "." {
			left++
			continue
		}
		if blocks[right] == "." {
			right--
			continue
		}
		blocks[left], blocks[right] = blocks[right], blocks[left]
		left++
		right--
	}

	fmt.Println("Part 1 Answer:", checkSum(blocks))

	// Part2
	freeSpacePointer := 0
	freeSpaceLength := 0
	keys := make([]int, 0, len(locationMapping))
	for key := range locationMapping {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	freeSpaceMapping := [][]int{}
	for freeSpacePointer < len(part2Blocks) {
		if part2Blocks[freeSpacePointer] == "." {
			freeSpacePointer++
			freeSpaceLength++
		} else {
			if freeSpaceLength > 0 {
				freeSpaceMapping = append(freeSpaceMapping, []int{freeSpacePointer - freeSpaceLength, freeSpaceLength})
				freeSpaceLength = 0
			}
			freeSpacePointer++
		}
	}

	for i := len(keys) - 1; i > 0; i-- {
		strKey := strconv.Itoa(keys[i])
		location := locationMapping[keys[i]]
		start := location[0]
		end := location[1]
		moveBlock(strKey, start, end, part2Blocks, freeSpaceMapping)

	}

	fmt.Println("Part 2 Answer:", checkSum(part2Blocks))
}

func checkSum(blocks []string) int {
	counter := 0
	total := 0
	for i := 0; i < len(blocks); i++ {
		if blocks[i] == "." {
			counter++
			continue
		}
		tmp, _ := strconv.Atoi(blocks[i])
		total += (counter * tmp)
		counter++
	}
	return total
}

func moveBlock(id string, start int, swapEnd int, blocks []string, mapping [][]int) {
	amount := swapEnd - start + 1
	for idx, blockMapping := range mapping {
		if amount <= blockMapping[1] && start > blockMapping[0] {
			for i := 0; i < amount; i++ {
				blocks[blockMapping[0]+i], blocks[swapEnd] = id, "."
				swapEnd--
			}
			mapping[idx][1] = blockMapping[1] - amount
			mapping[idx][0] = blockMapping[0] + amount
			return
		}
	}
}
