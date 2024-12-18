package main

import (
	"aoc2024/internal/helper"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func getOperand(op int, register [3]int) int {
	if op < 4 {
		return op
	}

	if op == 7 {
		fmt.Println("we have an invalid num")
	}

	return register[op-4]
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func compute(pointer int, instructions []int, register [3]int) []int {
	jump := false
	out := []int{}

	for pointer < len(instructions) {
		operand := instructions[pointer+1]
		switch instructions[pointer] {
		case 0:
			register[0] = register[0] / powInt(2, getOperand(operand, register))
		case 1:
			register[1] = register[1] ^ operand
		case 2:
			register[1] = getOperand(operand, register) % 8
		case 3:
			if register[0] != 0 {
				pointer = operand
				jump = true
			}
		case 4:
			register[1] = register[1] ^ register[2]
		case 5:
			out = append(out, getOperand(operand, register)%8)
		case 6:
			register[1] = register[0] / powInt(2, getOperand(operand, register))
		case 7:
			register[2] = register[0] / powInt(2, getOperand(operand, register))

		}
		if jump {
			jump = false
		} else {
			pointer += 2
		}
	}
	return out
}

func intArrayToStringArray(intArr []int) []string {
	strArr := make([]string, len(intArr))
	for i, val := range intArr {
		strArr[i] = strconv.Itoa(val) // Convert int to string
	}
	return strArr
}

func main() {
	file := helper.ReadFile("input.txt")
	parts := strings.Split(file, "\n\n")
	register := [3]int{}
	instructions := []int{}
	pointer := 0
	for i, p := range strings.Split(parts[0], "\n") {
		split := strings.Split(p, ":")
		numStr := strings.TrimSpace(split[1])
		num, _ := strconv.Atoi(numStr)
		register[i] = num

	}

	parts[1] = parts[1][9:]
	for _, p := range strings.Split(parts[1], ",") {
		num, _ := strconv.Atoi(strings.TrimSpace(p))
		instructions = append(instructions, num)
	}

	out := compute(pointer, instructions, register)

	fmt.Println("Part 1 Answer:", strings.Join(intArrayToStringArray(out), ","))

	a := 0
	register[0], register[1], register[2] = a, 0, 0

    // This took about 50 minutes to brute force
	for i := len(instructions) - 1; i >= 0; i-- {
		a <<= 3
		for !slices.Equal(compute(pointer, instructions, register), instructions[i:]) {
			a++
			register[0], register[1], register[2] = a, 0, 0
		}
	}
	fmt.Println("Part 2 Answer", a)
}
