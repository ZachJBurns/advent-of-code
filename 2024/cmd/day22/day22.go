package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	numbers := []int{}

	for scanner.Scan() {
		strNum := scanner.Text()

		num, _ := strconv.Atoi(strNum)
		numbers = append(numbers, num)
	}
	total := 0
	changeArr := [][]int{}
	for idx, num := range numbers {
		arr := []int{num % 10}
		for i := 0; i < 2000; i++ {
			ret := num * 64
			num = num ^ ret
			num = num % 16777216
			ret = num / 32
			num = num ^ ret
			num = num % 16777216
			ret = num * 2048
			num = num ^ ret
			numbers[idx] = num % 16777216
			arr = append(arr, (num%16777216)%10)
		}
		total += numbers[idx]
		changeArr = append(changeArr, arr)

	}

	changeMap := map[string]int{}
    maxNum := 0
	for _, arr := range changeArr {
		firstSeen := map[string]bool{}

		for j := 0; j < len(arr)-4; j += 1 {
			c1, c2, c3, c4 := arr[j+1]-arr[j], arr[j+2]-arr[j+1], arr[j+3]-arr[j+2], arr[j+4]-arr[j+3]
			num := arr[j+4]
			key := formatChange(c1, c2, c3, c4)
			if firstSeen[key] {
				continue
			}
			firstSeen[key] = true
			changeMap[key] += num
            if changeMap[key] > maxNum {
                maxNum = changeMap[key]
            }
		}
	}

	fmt.Println("Part 1 Answer:", total)
	fmt.Println("Part 2 Answer:", maxNum)
}

func formatChange(c1, c2, c3, c4 int) string {
	return fmt.Sprintf("%d,%d,%d,%d", c1, c2, c3, c4)
}
