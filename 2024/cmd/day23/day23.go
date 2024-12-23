package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	network := map[string]map[string]bool{}

	for scanner.Scan() {
		connection := scanner.Text()
		splitConnection := strings.Split(connection, "-")

		if _, ok := network[splitConnection[0]]; !ok {
			network[splitConnection[0]] = make(map[string]bool)
		}
		if _, ok := network[splitConnection[1]]; !ok {
			network[splitConnection[1]] = make(map[string]bool)
		}

		network[splitConnection[0]][splitConnection[1]] = true
		network[splitConnection[1]][splitConnection[0]] = true
	}

	part1 := findClique(network, 3)
	total := 0
	for _, arr := range part1 {
		for _, v := range arr {
			if strings.HasPrefix(v, "t") {
				total++
				break
			}
		}
	}
	fmt.Println("Part 1 Answer:", total)
    
    // Just incremented by 1 each time until I found the answer.
    // Lots of graph research today...
    part2 := findClique(network, 13)
    out := part2[0]
    sort.Strings(out)
    fmt.Println("Part 2 Answer", strings.Join(out, ","))
}

func findClique(graph map[string]map[string]bool, size int) [][]string {
	cliques := [][]string{}

	var dfs func(path []string, candidates []string)
	dfs = func(path []string, candidates []string) {
		if len(path)+len(candidates) < size {
			return
		}

		if len(path) == size {
			cliques = append(cliques, path)
			return
		}

		for i, candidate := range candidates {
			isValid := true
			for _, node := range path {
				if !graph[node][candidate] {
					isValid = false
					break
				}
			}

			if isValid {
				dfs(append(path, candidate), candidates[i+1:])
			}
		}
	}

	nodes := make([]string, 0, len(graph))
	for node := range graph {
		nodes = append(nodes, node)
	}
	sort.Strings(nodes)

	dfs([]string{}, nodes)
	return cliques
}
