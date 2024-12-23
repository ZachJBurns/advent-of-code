package main

import (
	"aoc2024/internal/helper"
	"fmt"
	"sort"
	"strings"
)

func BFS(design string, allowedPatterns []string) bool {
	queue := []int{0}
	visited := make(map[int]bool)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == len(design) {
			return true
		}

		if visited[current] {
			continue
		}
		visited[current] = true

		for _, pattern := range allowedPatterns {
			next := current + len(pattern)

			if next <= len(design) && design[current:next] == pattern {
				queue = append(queue, next)
			}
		}
	}

	return false
}

func main() {
	file := helper.ReadFile("input.txt")

	parts := strings.Split(file, "\n\n")
	allowedPatterns := strings.Split(parts[0], ", ")

	designs := strings.Split(parts[1], "\n")
	designs = designs[:len(designs)-1]

	// Sort patterns by length (descending) to prioritize longer matches
	sort.Slice(allowedPatterns, func(i, j int) bool {
		return len(allowedPatterns[i]) > len(allowedPatterns[j])
	})

	total := 0

	for _, design := range designs {
		if BFS(design, allowedPatterns) {
			total++
		}
	}

	total2 := 0
	for _, design := range designs {
		total2 += DFS(design, allowedPatterns)
	}
	fmt.Println("Part 1 Answer:", total)
	fmt.Println("Part 2 Answer:", total2)
}

func DFS(design string, allowedPatterns []string) int {
	cache := map[int]int{}
	var dfs func(int) int
	dfs = func(start int) int {
		if start == len(design) {
			return 1
		}

		if val, exists := cache[start]; exists {
			return val
		}

		totalMatches := 0
		for _, pattern := range allowedPatterns {
			next := start + len(pattern)

			if next <= len(design) && design[start:next] == pattern {
				totalMatches += dfs(next)
			}
		}

		cache[start] = totalMatches
		return totalMatches
	}
	return dfs(0)
}
