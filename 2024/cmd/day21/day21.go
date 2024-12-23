package main

import (
	"aoc2024/internal/helper"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	keypad := [4][3]string{
		{"7", "8", "9"},
		{"4", "5", "6"},
		{"1", "2", "3"},
		{"", "0", "A"},
	}

	var changeMap = map[string]map[string]string{
		"A": {
			"A": "A",
			"^": "^",
			"<": "<",
			">": ">",
			"v": "v",
		},
		">": {
			"A": "A",
			"^": "A",
			"<": "v",
			">": ">",
			"v": ">",
		},
		"^": {
			"A": "A",
			"^": "^",
			"<": "<",
			"v": "^",
			">": "A",
		},
		"<": {
			"A": "^",
			"^": "^",
			"<": "<",
			"v": "<",
			">": "v",
		},
		"v": {
			"A": ">",
			"^": "v",
			"<": "<",
			"v": "v",
			">": ">",
		},
	}

	file := helper.ReadFile("input.txt")
	codes := strings.Split(file, "\n")
	codes = codes[:len(codes)-1]

	fmt.Println(codes)
	total := 0
	for _, code := range codes {
		startPos := [2]string{"A", "A"}
		startKeypad := [2]int{3, 2}
		shortestPath := findShortestPath(keypad, changeMap, startPos, startKeypad, code)
		codeNum, _ := strconv.Atoi(code[:len(code)-1])
		fmt.Println(shortestPath)
		fmt.Println("<v<A>>^AA<vA<A>>^AAvAA<^A>A<vA>^A<A>A<vA>^A<A>A<v<A>A>^AAvA<^A>A")
		fmt.Println(codeNum, len(shortestPath))
		total += (codeNum * len(shortestPath))
		fmt.Println("Shortest Input Path:", shortestPath, len(shortestPath))
	}
	fmt.Println(total)
}

func findShortestPath(keypad [4][3]string, changeMap map[string]map[string]string, startPos [2]string, startKeypad [2]int, targetOutput string) string {
	type State struct {
		puzzlePos [2]string
		keypadPos [2]int
		out       string
		path      string
	}

	queue := []State{
		{startPos, startKeypad, "", ""},
	}

	visited := map[string]bool{}

	// BFS Loop
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if len(curr.out) > len(targetOutput) {
			continue
		}
		// Check if we've reached the target output
		if curr.out == targetOutput {
			return curr.path
		}

		// Create a unique key to track visited states
		stateKey := fmt.Sprintf("%v-%v-%s", curr.puzzlePos, curr.keypadPos, curr.out)
		if visited[stateKey] {
			continue
		}
		visited[stateKey] = true

		// Explore all possible inputs
		for _, val := range []string{"<", ">", "^", "v", "A"} {
			newPuzzlePos := curr.puzzlePos
			newKeypadPos := curr.keypadPos
			newOut := curr.out
			newPath := curr.path + val

			puzzle2 := changeMap[val]
			newPuzzlePos[0] = puzzle2[newPuzzlePos[0]]
			if val == "A" {
				x := puzzle2[newPuzzlePos[0]]
				if x == "A" {
					moveVal := newPuzzlePos[1]

					switch moveVal {
					case "^": // Move up
						if newKeypadPos[0] > 0 && keypad[newKeypadPos[0]-1][newKeypadPos[1]] != "" {
							newKeypadPos[0]--
						}
					case "v": // Move down
						if newKeypadPos[0] < len(keypad)-1 && keypad[newKeypadPos[0]+1][newKeypadPos[1]] != "" {
							newKeypadPos[0]++
						}
					case "<": // Move left
						if newKeypadPos[1] > 0 && keypad[newKeypadPos[0]][newKeypadPos[1]-1] != "" {
							newKeypadPos[1]--
						}
					case ">": // Move right
						if newKeypadPos[1] < len(keypad[0])-1 && keypad[newKeypadPos[0]][newKeypadPos[1]+1] != "" {
							newKeypadPos[1]++
						}
					case "A": // Select current key
						// Ensure the current position is not an empty space before adding to the output
						if keypad[newKeypadPos[0]][newKeypadPos[1]] != "" {
							newOut += keypad[newKeypadPos[0]][newKeypadPos[1]]
						}
					}

				} else {
					newPuzzlePos[1] = changeMap[x][newPuzzlePos[1]]
				}
			}
			// Prune branches where the output grows too long or mismatches the target
			if len(newOut) > len(targetOutput) || !strings.HasPrefix(targetOutput, newOut) {
				continue
			}

			// Add the new state to the queue
			queue = append(queue, State{newPuzzlePos, newKeypadPos, newOut, newPath})
		}
	}

	return "No Path Found"
}
